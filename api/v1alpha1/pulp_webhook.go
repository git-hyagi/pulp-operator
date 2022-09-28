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

package v1alpha1

import (
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var pulplog = logf.Log.WithName("pulp-resource")

func (r *Pulp) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-repo-manager-pulpproject-org-v1alpha1-pulp,mutating=true,failurePolicy=fail,sideEffects=None,groups=repo-manager.pulpproject.org,resources=pulps,verbs=create;update,versions=v1alpha1,name=mpulp.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Pulp{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Pulp) Default() {
	//pulplog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-repo-manager-pulpproject-org-v1alpha1-pulp,mutating=false,failurePolicy=fail,sideEffects=None,groups=repo-manager.pulpproject.org,resources=pulps,verbs=create;update,versions=v1alpha1,name=vpulp.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Pulp{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Pulp) ValidateCreate() error {
	//pulplog.Info("validate create", "name", r.Name)
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Pulp) ValidateUpdate(old runtime.Object) error {
	pulplog.Info("[validatingwebhook] checking update ...", "name", r.Name)

	// asserts that old is of type *Pulp
	oldPulp, _ := old.(*Pulp)

	// check immutable fields
	immutableFields := []string{"DeploymentType", "FileStorageSize", "FileStorageAccessMode"}
	for _, field := range immutableFields {
		newField := reflect.ValueOf(r.Spec).FieldByName(field)
		oldField := reflect.ValueOf(oldPulp.Spec).FieldByName(field)
		if newField.Interface() != oldField.Interface() {
			pulplog.Info("Error trying to modify an immutable field!", "field", field)
			return fmt.Errorf("%s is an immutable field", field)
		}
	}

	// non-nil values (once defined, these fields should not be deleted but can be modified)
	// we allow to use another SC or PVC or update credentials to access the Object Storage
	// but we should not allow to delete the field (which would deploy emptyDir) once it was defined
	nonNilFields := []string{"FileStorageClass", "ObjectStorageAzureSecret", "ObjectStorageS3Secret", "PVC"}
	for _, field := range nonNilFields {
		newField := reflect.ValueOf(r.Spec).FieldByName(field)
		oldField := reflect.ValueOf(oldPulp.Spec).FieldByName(field)
		if len(oldField.Interface().(string)) > 0 && len(newField.Interface().(string)) == 0 {
			pulplog.Info("Error trying to remove field!", "field", field)
			return fmt.Errorf("%s should not be removed because doing so would lose all cluster data", field)
		}
	}

	// check database field removal
	// if the old CR wasn't empty and we are trying to delete its content
	// we should not allow that
	if !reflect.DeepEqual(oldPulp.Spec.Database, Database{}) && reflect.DeepEqual(r.Spec.Database, Database{}) {
		pulplog.Info("Error trying to remove database field!")
		return fmt.Errorf("unable to remove database field... this field should not be deleted")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Pulp) ValidateDelete() error {
	//pulplog.Info("validate delete", "name", r.Name)
	return nil
}
