/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO: Add kubebuilder/operator-sdk markers

// PulpResourceSpec defines the desired state of PulpResource
type PulpResourceSpec struct {
	ClientConfig *string              `json:"client_config,omitempty"`
	Container    *PulpContainerPlugin `json:"container,omitempty"`
	File         *PulpFilePlugin      `json:"file,omitempty"`
	RPM          *PulpRPMPlugin       `json:"rpm,omitempty"`
}

type PulpContainerPlugin struct {
	Repository   *PulpContainerRepository   `json:"repository,omitempty"`
	Distribution *PulpContainerDistribution `json:"distribution,omitempty"`
	Remote       *PulpContainerRemote       `json:"remote,omitempty"`
	Sync         *bool                      `json:"sync,omitempty"`
}

type PulpContainerRepository struct {
	PulpCoreRepository `json:",inline"`
}

type PulpContainerDistribution struct {
	PulpCoreDistribution `json:",inline"`
	Version              int  `json:"version,omitempty"`
	Private              bool `json:"private,omitempty"`
}

type PulpContainerRemote struct {
	PulpCoreRemote `json:",inline"`
	UpstreamName   string   `json:"upstream_name,omitempty"`
	IncludeTags    []string `json:"include_tags,omitempty"`
	ExcludeTags    []string `json:"exclude_tags,omitempty"`
	Sigstore       string   `json:"sigstore,omitempty"`
}

type PulpFilePlugin struct {
	Repository   *PulpFileRepository   `json:"repository,omitempty"`
	Distribution *PulpFileDistribution `json:"distribution,omitempty"`
	Remote       *PulpFileRemote       `json:"remote,omitempty"`
	Sync         *bool                 `json:"sync,omitempty"`
}

type PulpFileRepository struct {
	PulpCoreRepository `json:",inline"`
	Manifest           string `json:"manifest,omitempty"`
	AutoPublish        bool   `json:"autopublish,omitempty"`
}

type PulpFileDistribution struct {
	PulpCoreDistribution `json:",inline"`
	Publication          string `json:"publication,omitempty"`
}

type PulpFileRemote struct {
	PulpCoreRemote `json:",inline"`
}

type PulpRPMPlugin struct {
	Repository   *PulpRPMRepository   `json:"repository,omitempty"`
	Distribution *PulpRPMDistribution `json:"distribution,omitempty"`
	Remote       *PulpRPMRemote       `json:"remote,omitempty"`
	Sync         *bool                `json:"sync,omitempty"`
}

type PulpRPMRepository struct {
	PulpCoreRepository     `json:",inline"`
	RetainPackageVersions  int    `json:"retain-package-vesions,omitempty"`
	MetadataSigningService string `json:"metadata-signing-service,omitempty"`
	ChecksumType           string `json:"checksum-type,omitempty"`
	Manifest               string `json:"manifest,omitempty"`
	AutoPublish            bool   `json:"autopublish,omitempty"`
	RepoConfig             string `json:"repo-config,omitempty"`
}

type PulpRPMDistribution struct {
	PulpCoreDistribution `json:",inline"`
	Publication          string `json:"publication,omitempty"`
	GenerateRepoConfig   bool   `json:"generate-repo-config,omitempty"`
}

type PulpRPMRemote struct {
	PulpCoreRemote `json:",inline"`
	SLESAuthToken  string `json:"sles_auth_token,omitempty"`
}

type PulpCoreRepository struct {
	Description        string            `json:"description,omitempty"`
	Remote             string            `json:"remote,omitempty"`
	RetainRepoVersions int               `json:"retain-repo-versions,omitempty"`
	Labels             map[string]string `json:"pulp_labels,omitempty"`
	Name               string            `json:"name"`
}

type PulpCoreDistribution struct {
	Name         string            `json:"name"`
	BasePath     string            `json:"base_path,omitempty"`
	Repository   string            `json:"repository,omitempty"`
	ContentGuard string            `json:"content-guard,omitempty"`
	Labels       map[string]string `json:"pulp_labels,omitempty"`
}

type PulpCoreRemote struct {
	Name                string              `json:"name"`
	URL                 string              `json:"url,omitempty"`
	CaCert              string              `json:"ca_cert,omitempty"`
	ClientCert          string              `json:"client_cert,omitempty"`
	ClientKey           string              `json:"client_key,omitempty"`
	TLSValidtaion       bool                `json:"tls_validation,omitempty"`
	ProxyURL            string              `json:"proxy_url,omitempty"`
	ProxyUsername       string              `json:"proxy_username,omitempty"`
	ProxyPassword       string              `json:"proxy_password,omitempty"`
	Username            string              `json:"username,omitempty"`
	Password            string              `json:"password,omitempty"`
	ConnectionTimeout   string              `json:"connection_timeout,omitempty"`
	DownloadConcurrency int                 `json:"download_concurrency,omitempty"`
	RateLimit           int                 `json:"rate_limit,omitempty"`
	SockConnectTimeout  string              `json:"sock_connect_timeout,omitempty"`
	SockReadTimeout     string              `json:"sock_read_timeout,omitempty"`
	MaxRetries          int                 `json:"max_retries,omitempty"`
	Policy              string              `json:"policy,omitempty"`
	TotalTimeout        string              `json:"total_timeout,omitempty"`
	Headers             []map[string]string `json:"headers,omitempty"`
	Labels              map[string]string   `json:"pulp_labels,omitempty"`
}

type PulpSync struct {
	Remote     string `json:"remote,omitempty"`
	Mirror     bool   `json:"mirror,omitempty"`
	SignedOnly bool   `json:"signed_only,omitempty"`
}

// PulpResourceStatus defines the observed state of PulpResource
type PulpResourceStatus struct {
	Container *PulpContainerPlugin `json:"container,omitempty"`
	File      *PulpFilePlugin      `json:"file,omitempty"`
	RPM       *PulpRPMPlugin       `json:"rpm,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// PulpResource is the Schema for the pulpresource API
type PulpResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PulpResourceSpec    `json:"spec,omitempty"`
	Status *PulpResourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PulpResourceList contains a list of PulpResource
type PulpResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PulpResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PulpResource{}, &PulpResourceList{})
}
