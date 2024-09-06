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

package repo_manager_pulp

import (
	"context"

	"github.com/go-logr/logr"
	repomanagerpulpprojectorgv1beta2 "github.com/pulp/pulp-operator/apis/repo-manager.pulpproject.org/v1beta2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// RepoManagerPulpResourceReconciler reconciles a PulpResource object
type RepoManagerPulpResourceReconciler struct {
	client.Client
	RawLogger logr.Logger
	Scheme    *runtime.Scheme
}

//+kubebuilder:rbac:groups=repo-manager.pulpproject.org,namespace=pulp-operator-system,resources=pulpresources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=repo-manager.pulpproject.org,namespace=pulp-operator-system,resources=pulpresources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=repo-manager.pulpproject.org,namespace=pulp-operator-system,resources=pulpresources/finalizers,verbs=update
//+kubebuilder:rbac:groups=repo-manager.pulpproject.org,namespace=pulp-operator-system,resources=pulps,verbs=get;list;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *RepoManagerPulpResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.RawLogger
	pulpResource := &repomanagerpulpprojectorgv1beta2.PulpResource{}
	if err := r.Get(ctx, req.NamespacedName, pulpResource); err != nil {
		if errors.IsNotFound(err) {
			log.Info("PulpResource resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get PulpResource")
		return ctrl.Result{}, err
	}

	pulpClient := getClientConfig(ctx, *r, *pulpResource)
	if err := validatePulpConnection(pulpClient); err != nil {
		log.Error(err, "PulpResource "+pulpResource.Name+" failed to connet to Pulp")
		return ctrl.Result{}, nil
	}
	r.createPulpResources(pulpResource, pulpClient)
	r.syncPulpResources(pulpResource, pulpClient)

	r.pulpResourceFinalizer(ctx, *pulpResource, pulpClient)
	if reconcile := r.updateStatusFields(pulpResource, pulpClient); reconcile != nil {
		return *reconcile, nil
	}
	log.Info("PulpResource " + pulpResource.Name + " synced")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RepoManagerPulpResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {

	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &repomanagerpulpprojectorgv1beta2.PulpResource{}, "pulpresources-objects", indexerFunc); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&repomanagerpulpprojectorgv1beta2.PulpResource{}).
		Watches(
			&corev1.Secret{},
			handler.EnqueueRequestsFromMapFunc(r.findPulpResourceDependentObjects),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}
