CR_KIND ?= Pulp
LOWER_CR_KIND = $(shell echo $(CR_KIND) | tr A-Z a-z)
CR_PLURAL ?= pulps
CR_DOMAIN ?= pulpproject.org
APP_IMAGE ?= quay.io/pulp/pulp-minimal
WEB_IMAGE ?= quay.io/pulp/pulp-web

# VERSION defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
# To re-generate a bundle for another specific version without changing the standard setup, you can:
# - use the VERSION as arg of the bundle target (e.g make bundle VERSION=0.0.2)
# - use environment variables to overwrite this value (e.g export VERSION=0.0.2)
VERSION ?= 1.0.0-beta.5

# CHANNELS define the bundle channels used in the bundle.
CHANNELS = "beta"
# Add a new line here if you would like to change its default config. (E.g CHANNELS = "candidate,fast,stable")
# To re-generate a bundle for other specific channels without changing the standard setup, you can:
# - use the CHANNELS as arg of the bundle target (e.g make bundle CHANNELS=candidate,fast,stable)
# - use environment variables to overwrite this value (e.g export CHANNELS="candidate,fast,stable")
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif

# DEFAULT_CHANNEL defines the default channel used in the bundle.
# Add a new line here if you would like to change its default config. (E.g DEFAULT_CHANNEL = "stable")
DEFAULT_CHANNEL = "beta"
# To re-generate a bundle for any other default channel without changing the default setup, you can:
# - use the DEFAULT_CHANNEL as arg of the bundle target (e.g make bundle DEFAULT_CHANNEL=stable)
# - use environment variables to overwrite this value (e.g export DEFAULT_CHANNEL="stable")
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# IMAGE_TAG_BASE defines the docker.io namespace and part of the image name for remote images.
# This variable is used to construct full image tags for bundle and catalog images.
#
# For example, running 'make bundle-build bundle-push catalog-build catalog-push' will build and push both
# pulpproject.org/pulp-operator-bundle:$VERSION and pulpproject.org/pulp-operator-catalog:$VERSION.
#IMAGE_TAG_BASE ?= pulpproject.org/pulp-operator
IMAGE_TAG_BASE ?= quay.io/pulp/$(LOWER_CR_KIND)-operator

# BUNDLE_IMG defines the image:tag used for the bundle.
# You can use it as an arg. (E.g make bundle-build BUNDLE_IMG=<some-registry>/<project-name-bundle>:<tag>)
BUNDLE_IMG ?= $(IMAGE_TAG_BASE)-bundle:v$(VERSION)

# BUNDLE_GEN_FLAGS are the flags passed to the operator-sdk generate bundle command
BUNDLE_GEN_FLAGS ?= -q --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)

# USE_IMAGE_DIGESTS defines if images are resolved via tags or digests
# You can enable this value if you would like to use SHA Based Digests
# To enable set flag to true
USE_IMAGE_DIGESTS ?= false
ifeq ($(USE_IMAGE_DIGESTS), true)
	BUNDLE_GEN_FLAGS += --use-image-digests
endif

# Image URL to use all building/pushing image targets
IMG ?= $(IMAGE_TAG_BASE):v$(VERSION)
NAMESPACE ?= $(LOWER_CR_KIND)-operator-system
WATCH_NAMESPACE ?= $(NAMESPACE)

# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.24.2

GOLANG_VERSION=1.20.10
GOLANG_ARCH=linux-amd64
GOLANG_INSTALL_PATH=/tmp

export CONTAINER_TOOL ?= docker

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ifeq ($(shell go env GOROOT 2>/dev/null),)
PATH=$(shell echo $$PATH:$(GOLANG_INSTALL_PATH)/go/bin)
endif


# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: tidy controller-gen crd-to-markdown ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
	$(CRD_MARKDOWN) -f apis/repo-manager.$(CR_DOMAIN)/v1beta2/pulp_types.go -n $(CR_KIND) > controllers/repo_manager/README.md
	$(CRD_MARKDOWN) -f apis/repo-manager.$(CR_DOMAIN)/v1beta2/pulp_backup_types.go -n $(CR_KIND)Backup > controllers/backup/README.md
	$(CRD_MARKDOWN) -f apis/repo-manager.$(CR_DOMAIN)/v1beta2/pulp_restore_types.go -n $(CR_KIND)Restore > controllers/restore/README.md
	$(CRD_MARKDOWN) -f apis/repo-manager.$(CR_DOMAIN)/v1beta2/pulp_resource_types.go -n $(CR_KIND)Resource > controllers/pulp_resources/README.md

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: manifests generate fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) --bin-dir $(LOCALBIN) -p path)" go test -v ./... -coverprofile cover.out

.PHONY: testbin
testbin:  ## Ensure envtest bins.
# https://kubebuilder.io/reference/envtest.html
ifneq ($(wildcard /usr/local/kubebuilder), )
	@echo "/usr/local/kubebuilder already exists"
else
	curl -sSLo envtest-bins.tar.gz "https://go.kubebuilder.io/test-tools/$(ENVTEST_K8S_VERSION)/$(shell go env GOOS)/$(shell go env GOARCH)"
	sudo mkdir -p /usr/local/kubebuilder
	sudo tar -C /usr/local/kubebuilder --strip-components=1 -zvxf envtest-bins.tar.gz
endif

.PHONY: golang
golang: ## Ensure golang is installed
ifeq ($(shell which go 2>/dev/null), )
	curl -sSLo golang.tar.gz "https://go.dev/dl/go$(GOLANG_VERSION).$(GOLANG_ARCH).tar.gz"
	tar -C $(GOLANG_INSTALL_PATH) -xzf golang.tar.gz
endif

.PHONY: docs
docs: ## Build unified docs
	pulp-docs build

.PHONY: servedocs
servedocs: ## Build unified docs
	pulp-docs serve

##@ Build

.PHONY: build
build: generate fmt vet ## Build manager binary.
	go build -o bin/manager main.go

.PHONY: rename
rename: ## Replace Custom Resource name
	find apis/repo-manager.pulpproject.org/*/*_types.go -exec sed -i "s/Pulp/${CR_KIND}/g" {} \;
	find controllers/*.go -exec sed -i "s/Pulp/${CR_KIND}/g" {} \;
	find controllers/*/*.go -exec sed -i "s/Pulp/${CR_KIND}/g" {} \;
	find config/samples/*.yaml -exec sed -i "s/Pulp/${CR_KIND}/g" {} \;
	sed -i "s/Pulp/${CR_KIND}/g" PROJECT
	sed -i "s/pulpproject.org/${CR_DOMAIN}/g" PROJECT main.go
	sed -i "s/RepoManagerPulpResourceReconciler/RepoManager${CR_KIND}ResourceReconciler/g" main.go
	find apis/repo-manager.pulpproject.org/*/* -exec sed -i "s/repo-manager.pulpproject.org/repo-manager.${CR_DOMAIN}/g" {} \;
	find bundle/manifests/* -exec sed -i "s/repo-manager.pulpproject.org/repo-manager.${CR_DOMAIN}/g" {} \;
	find config/*/* -exec sed -i "s/pulps/${CR_PLURAL}/g" {} \;
	find config/*/* -exec sed -i "s/pulprestores/${LOWER_CR_KIND}restores/g" {} \;
	find config/*/* -exec sed -i "s/pulpbackups/${LOWER_CR_KIND}backups/g" {} \;
	find config/*/* -exec sed -i "s/pulpresources/${LOWER_CR_KIND}resources/g" {} \;
	sed -i "s/pulp/${LOWER_CR_KIND}/g" config/default/kustomization.yaml
	find config/*/* -exec sed -i "s/3b5210cd.pulpproject.org/3b5210cd.${CR_DOMAIN}/g" {} \;
	find config/*/* -exec sed -i "s/repo-manager.pulpproject.org/repo-manager.${CR_DOMAIN}/g" {} \;
	find controllers/*/*.go -exec sed -i "s/repo-manager.pulpproject.org/repo-manager.${CR_DOMAIN}/g" {} \;
	find controllers/*.go -exec sed -i "s/repo-manager.pulpproject.org/repo-manager.${CR_DOMAIN}/g" {} \;
	find controllers/*/*.go -exec sed -i "s/pulps/${CR_PLURAL}/g" {} \;
	find controllers/*.go -exec sed -i "s/pulps/${CR_PLURAL}/g" {} \;
	find controllers/*/*.go -exec sed -i "s/pulpbackup/${LOWER_CR_KIND}backup/g" {} \;
	find controllers/*/*.go -exec sed -i "s/pulprestore/${LOWER_CR_KIND}restore/g" {} \;
	find controllers/*/*.go -exec sed -i "s/pulpresources/${LOWER_CR_KIND}resources/g" {} \;
	sed -i "s/default:=\"pulp\"/default:=\"${LOWER_CR_KIND}\"/" apis/repo-manager.pulpproject.org/v1beta2/pulp_types.go
	sed -i "s|quay.io/pulp/pulp-minimal|${APP_IMAGE}|g" apis/repo-manager.pulpproject.org/v1beta2/pulp_types.go
	sed -i "s|quay.io/pulp/pulp-web|${WEB_IMAGE}|g" apis/repo-manager.pulpproject.org/v1beta2/pulp_types.go
	sed -i '0,/OperatorType/s/OperatorType.*/OperatorType  = "${LOWER_CR_KIND}"/' controllers/repo_manager/controller_test.go
	sed -i "s|quay.io/pulp/pulp-minimal|${APP_IMAGE}|g" controllers/repo_manager/controller_test.go
	mv config/crd/bases/repo-manager.pulpproject.org_pulps.yaml config/crd/bases/repo-manager.${CR_DOMAIN}_${CR_PLURAL}.yaml
	mv config/crd/bases/repo-manager.pulpproject.org_pulpbackups.yaml config/crd/bases/repo-manager.${CR_DOMAIN}_${LOWER_CR_KIND}backups.yaml
	mv config/crd/bases/repo-manager.pulpproject.org_pulprestores.yaml config/crd/bases/repo-manager.${CR_DOMAIN}_${LOWER_CR_KIND}restores.yaml
	mv config/crd/bases/repo-manager.pulpproject.org_pulpresources.yaml config/crd/bases/repo-manager.${CR_DOMAIN}_${LOWER_CR_KIND}resources.yaml
	mv bundle/manifests/repo-manager.pulpproject.org_pulps.yaml bundle/manifests/repo-manager.${CR_DOMAIN}_${CR_PLURAL}.yaml
	mv bundle/manifests/repo-manager.pulpproject.org_pulpbackups.yaml bundle/manifests/repo-manager.${CR_DOMAIN}_${LOWER_CR_KIND}backups.yaml
	mv bundle/manifests/repo-manager.pulpproject.org_pulprestores.yaml bundle/manifests/repo-manager.${CR_DOMAIN}_${LOWER_CR_KIND}restores.yaml
	mv bundle/manifests/repo-manager.pulpproject.org_pulpresources.yaml bundle/manifests/repo-manager.${CR_DOMAIN}_${LOWER_CR_KIND}resources.yaml
	mv apis/repo-manager.pulpproject.org apis/repo-manager.${CR_DOMAIN}
	sed -i "s/repo-manager.pulpproject.org/repo-manager.${CR_DOMAIN}/g" controllers/utils.go
	mv config/samples/repo-manager.pulpproject.org_v1beta2_pulp.yaml config/samples/repo-manager.${CR_DOMAIN}_v1beta2_pulp.yaml

.PHONY: run
run: manifests generate fmt vet ## Run a controller from your host.
	go run ./main.go

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	$(CONTAINER_TOOL) build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	$(CONTAINER_TOOL) push ${IMG}

# PLATFORMS defines the target platforms for  the manager image be build to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - able to use docker buildx . More info: https://docs.docker.com/build/buildx/
# - have enable BuildKit, More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image for your registry (i.e. if you do not inform a valid value via IMG=<myregistry/image:<tag>> than the export will fail)
# To properly provided solutions that supports more than one platform you should use this option.
PLATFORMS ?= linux/arm64,linux/amd64,linux/s390x,linux/ppc64le
.PHONY: docker-buildx
docker-buildx: ## Build and push docker image for the manager for cross-platform support
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	- docker buildx create --name project-v3-builder
	docker buildx use project-v3-builder
	- docker buildx build --push --platform=$(PLATFORMS) --tag ${IMG} -f Dockerfile.cross .
	- docker buildx rm project-v3-builder
	rm Dockerfile.cross

##@ Deployment

ifndef ignore-not-found
  ignore-not-found = false
endif

.PHONY: install
install: manifests kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl apply --server-side=true -f -

.PHONY: uninstall
uninstall: manifests kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/crd | kubectl delete --ignore-not-found=$(ignore-not-found) -f -

.PHONY: local
local: kustomize ## Run controller in the K8s cluster specified in ~/.kube/config.
	.ci/scripts/local.sh $(CR_KIND) $(CR_DOMAIN) $(CR_PLURAL) $(APP_IMAGE) $(WEB_IMAGE)

.PHONY: deploy
deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	cd config/default && $(KUSTOMIZE) edit set namespace ${NAMESPACE}
	$(KUSTOMIZE) build config/default | kubectl apply --server-side=true -f -

.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
	$(KUSTOMIZE) build config/default | kubectl delete --ignore-not-found=$(ignore-not-found) -f -

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
KUSTOMIZE ?= $(LOCALBIN)/kustomize
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
ENVTEST ?= $(LOCALBIN)/setup-envtest
CRD_MARKDOWN ?= $(LOCALBIN)/crd-to-markdown
SDK_BIN = $(LOCALBIN)/operator-sdk

## Tool Versions
KUSTOMIZE_VERSION ?= v3.8.7
CONTROLLER_TOOLS_VERSION ?= v0.14.0
CRD_MARKDOWN_VERSION ?= v0.0.3
OPERATOR_SDK_VERSION ?= v1.29.0

KUSTOMIZE_INSTALL_SCRIPT ?= "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"
.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary.
$(KUSTOMIZE): $(LOCALBIN)
	test -s $(LOCALBIN)/kustomize || { curl -s $(KUSTOMIZE_INSTALL_SCRIPT) | bash -s -- $(subst v,,$(KUSTOMIZE_VERSION)) $(LOCALBIN); }

.PHONY: controller-gen
controller-gen: golang $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: crd-to-markdown
crd-to-markdown: $(CRD_MARKDOWN) ## Download crd-to-markdown locally if necessary.
$(CRD_MARKDOWN): $(LOCALBIN)
	test -s $(LOCALBIN)/crd-to-markdown || GOBIN=$(LOCALBIN) go install github.com/clamoriniere/crd-to-markdown@$(CRD_MARKDOWN_VERSION)

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	test -s $(LOCALBIN)/setup-envtest || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest

.PHONY: sdkbin
sdkbin: ## Download operator-sdk locally if necessary, preferring the $(pwd)/bin path over global if both exist.
ifeq (,$(wildcard $(SDK_BIN)))
ifeq (,$(shell which operator-sdk 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(SDK_BIN)) ;\
	OS=$(shell go env GOOS) && ARCH=$(shell go env GOARCH) && \
	curl -sSLo $(SDK_BIN) https://github.com/operator-framework/operator-sdk/releases/download/$(OPERATOR_SDK_VERSION)/operator-sdk_$${OS}_$${ARCH} ;\
	chmod +x $(SDK_BIN) ;\
	}
else
SDK_BIN = $(shell which operator-sdk)
endif
endif

.PHONY: bundle
bundle: sdkbin manifests kustomize ## Generate bundle manifests and metadata, then validate generated files.
	$(SDK_BIN) generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	$(KUSTOMIZE) build config/manifests | $(SDK_BIN) generate bundle $(BUNDLE_GEN_FLAGS)
	$(SDK_BIN) bundle validate ./bundle

.PHONY: bundle-build
bundle-build: ## Build the bundle image.
	$(CONTAINER_TOOL) build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

.PHONY: bundle-push
bundle-push: ## Push the bundle image.
	$(MAKE) docker-push IMG=$(BUNDLE_IMG)


.PHONY: bundle-check
bundle-check: sdkbin kustomize
	$(KUSTOMIZE) build config/manifests | $(SDK_BIN) generate bundle -q --output-dir=/tmp/bundle --overwrite=false --version $(VERSION) $(BUNDLE_METADATA_OPTS)

.PHONY: opm
OPM = ./bin/opm
opm: ## Download opm locally if necessary.
ifeq (,$(wildcard $(OPM)))
ifeq (,$(shell which opm 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OPM)) ;\
	OS=$(shell go env GOOS) && ARCH=$(shell go env GOARCH) && \
	curl -sSLo $(OPM) https://github.com/operator-framework/operator-registry/releases/download/v1.23.0/$${OS}-$${ARCH}-opm ;\
	chmod +x $(OPM) ;\
	}
else
OPM = $(shell which opm)
endif
endif

# A comma-separated list of bundle images (e.g. make catalog-build BUNDLE_IMGS=example.com/operator-bundle:v0.1.0,example.com/operator-bundle:v0.2.0).
# These images MUST exist in a registry and be pull-able.
BUNDLE_IMGS ?= $(BUNDLE_IMG)

# The image tag given to the resulting catalog image (e.g. make catalog-build CATALOG_IMG=example.com/operator-catalog:v0.2.0).
CATALOG_IMG ?= $(IMAGE_TAG_BASE)-catalog:v$(VERSION)

# Set CATALOG_BASE_IMG to an existing catalog image tag to add $BUNDLE_IMGS to that image.
ifneq ($(origin CATALOG_BASE_IMG), undefined)
FROM_INDEX_OPT := --from-index $(CATALOG_BASE_IMG)
endif

# Build a catalog image by adding bundle images to an empty catalog using the operator package manager tool, 'opm'.
# This recipe invokes 'opm' in 'semver' bundle add mode. For more information on add modes, see:
# https://github.com/operator-framework/community-operators/blob/7f1438c/docs/packaging-operator.md#updating-your-existing-operator
.PHONY: catalog-build
catalog-build: opm ## Build a catalog image.
	$(OPM) index add --container-tool $(CONTAINER_TOOL) --mode semver --tag $(CATALOG_IMG) --bundles $(BUNDLE_IMGS) $(FROM_INDEX_OPT)

# Push the catalog image.
.PHONY: catalog-push
catalog-push: ## Push a catalog image.
	$(MAKE) docker-push IMG=$(CATALOG_IMG)
