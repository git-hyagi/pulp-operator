//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Api) DeepCopyInto(out *Api) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.PDB != nil {
		in, out := &in.PDB, &out.PDB
		*out = new(policyv1.PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	in.InitContainer.DeepCopyInto(&out.InitContainer)
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Api.
func (in *Api) DeepCopy() *Api {
	if in == nil {
		return nil
	}
	out := new(Api)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cache) DeepCopyInto(out *Cache) {
	*out = *in
	in.RedisResourceRequirements.DeepCopyInto(&out.RedisResourceRequirements)
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cache.
func (in *Cache) DeepCopy() *Cache {
	if in == nil {
		return nil
	}
	out := new(Cache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Content) DeepCopyInto(out *Content) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.PDB != nil {
		in, out := &in.PDB, &out.PDB
		*out = new(policyv1.PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	in.InitContainer.DeepCopyInto(&out.InitContainer)
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Content.
func (in *Content) DeepCopy() *Content {
	if in == nil {
		return nil
	}
	out := new(Content)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	if in.PostgresExtraArgs != nil {
		in, out := &in.PostgresExtraArgs, &out.PostgresExtraArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostgresStorageClass != nil {
		in, out := &in.PostgresStorageClass, &out.PostgresStorageClass
		*out = new(string)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAP) DeepCopyInto(out *LDAP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAP.
func (in *LDAP) DeepCopy() *LDAP {
	if in == nil {
		return nil
	}
	out := new(LDAP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pulp) DeepCopyInto(out *Pulp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pulp.
func (in *Pulp) DeepCopy() *Pulp {
	if in == nil {
		return nil
	}
	out := new(Pulp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pulp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpBackup) DeepCopyInto(out *PulpBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpBackup.
func (in *PulpBackup) DeepCopy() *PulpBackup {
	if in == nil {
		return nil
	}
	out := new(PulpBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpBackupList) DeepCopyInto(out *PulpBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulpBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpBackupList.
func (in *PulpBackupList) DeepCopy() *PulpBackupList {
	if in == nil {
		return nil
	}
	out := new(PulpBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpBackupSpec) DeepCopyInto(out *PulpBackupSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpBackupSpec.
func (in *PulpBackupSpec) DeepCopy() *PulpBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PulpBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpBackupStatus) DeepCopyInto(out *PulpBackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpBackupStatus.
func (in *PulpBackupStatus) DeepCopy() *PulpBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PulpBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpContainer) DeepCopyInto(out *PulpContainer) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpContainer.
func (in *PulpContainer) DeepCopy() *PulpContainer {
	if in == nil {
		return nil
	}
	out := new(PulpContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpContainerDistribution) DeepCopyInto(out *PulpContainerDistribution) {
	*out = *in
	in.PulpCoreDistribution.DeepCopyInto(&out.PulpCoreDistribution)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpContainerDistribution.
func (in *PulpContainerDistribution) DeepCopy() *PulpContainerDistribution {
	if in == nil {
		return nil
	}
	out := new(PulpContainerDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpContainerPlugin) DeepCopyInto(out *PulpContainerPlugin) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(PulpContainerRepository)
		(*in).DeepCopyInto(*out)
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(PulpContainerDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(PulpContainerRemote)
		(*in).DeepCopyInto(*out)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpContainerPlugin.
func (in *PulpContainerPlugin) DeepCopy() *PulpContainerPlugin {
	if in == nil {
		return nil
	}
	out := new(PulpContainerPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpContainerRemote) DeepCopyInto(out *PulpContainerRemote) {
	*out = *in
	in.PulpCoreRemote.DeepCopyInto(&out.PulpCoreRemote)
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeTags != nil {
		in, out := &in.ExcludeTags, &out.ExcludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpContainerRemote.
func (in *PulpContainerRemote) DeepCopy() *PulpContainerRemote {
	if in == nil {
		return nil
	}
	out := new(PulpContainerRemote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpContainerRepository) DeepCopyInto(out *PulpContainerRepository) {
	*out = *in
	in.PulpCoreRepository.DeepCopyInto(&out.PulpCoreRepository)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpContainerRepository.
func (in *PulpContainerRepository) DeepCopy() *PulpContainerRepository {
	if in == nil {
		return nil
	}
	out := new(PulpContainerRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpCoreDistribution) DeepCopyInto(out *PulpCoreDistribution) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpCoreDistribution.
func (in *PulpCoreDistribution) DeepCopy() *PulpCoreDistribution {
	if in == nil {
		return nil
	}
	out := new(PulpCoreDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpCoreRemote) DeepCopyInto(out *PulpCoreRemote) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpCoreRemote.
func (in *PulpCoreRemote) DeepCopy() *PulpCoreRemote {
	if in == nil {
		return nil
	}
	out := new(PulpCoreRemote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpCoreRepository) DeepCopyInto(out *PulpCoreRepository) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpCoreRepository.
func (in *PulpCoreRepository) DeepCopy() *PulpCoreRepository {
	if in == nil {
		return nil
	}
	out := new(PulpCoreRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpFileDistribution) DeepCopyInto(out *PulpFileDistribution) {
	*out = *in
	in.PulpCoreDistribution.DeepCopyInto(&out.PulpCoreDistribution)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpFileDistribution.
func (in *PulpFileDistribution) DeepCopy() *PulpFileDistribution {
	if in == nil {
		return nil
	}
	out := new(PulpFileDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpFilePlugin) DeepCopyInto(out *PulpFilePlugin) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(PulpFileRepository)
		(*in).DeepCopyInto(*out)
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(PulpFileDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(PulpFileRemote)
		(*in).DeepCopyInto(*out)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpFilePlugin.
func (in *PulpFilePlugin) DeepCopy() *PulpFilePlugin {
	if in == nil {
		return nil
	}
	out := new(PulpFilePlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpFileRemote) DeepCopyInto(out *PulpFileRemote) {
	*out = *in
	in.PulpCoreRemote.DeepCopyInto(&out.PulpCoreRemote)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpFileRemote.
func (in *PulpFileRemote) DeepCopy() *PulpFileRemote {
	if in == nil {
		return nil
	}
	out := new(PulpFileRemote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpFileRepository) DeepCopyInto(out *PulpFileRepository) {
	*out = *in
	in.PulpCoreRepository.DeepCopyInto(&out.PulpCoreRepository)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpFileRepository.
func (in *PulpFileRepository) DeepCopy() *PulpFileRepository {
	if in == nil {
		return nil
	}
	out := new(PulpFileRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpJob) DeepCopyInto(out *PulpJob) {
	*out = *in
	in.PulpContainer.DeepCopyInto(&out.PulpContainer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpJob.
func (in *PulpJob) DeepCopy() *PulpJob {
	if in == nil {
		return nil
	}
	out := new(PulpJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpList) DeepCopyInto(out *PulpList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pulp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpList.
func (in *PulpList) DeepCopy() *PulpList {
	if in == nil {
		return nil
	}
	out := new(PulpList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRPMDistribution) DeepCopyInto(out *PulpRPMDistribution) {
	*out = *in
	in.PulpCoreDistribution.DeepCopyInto(&out.PulpCoreDistribution)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRPMDistribution.
func (in *PulpRPMDistribution) DeepCopy() *PulpRPMDistribution {
	if in == nil {
		return nil
	}
	out := new(PulpRPMDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRPMPlugin) DeepCopyInto(out *PulpRPMPlugin) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(PulpRPMRepository)
		(*in).DeepCopyInto(*out)
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(PulpRPMDistribution)
		(*in).DeepCopyInto(*out)
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(PulpRPMRemote)
		(*in).DeepCopyInto(*out)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRPMPlugin.
func (in *PulpRPMPlugin) DeepCopy() *PulpRPMPlugin {
	if in == nil {
		return nil
	}
	out := new(PulpRPMPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRPMRemote) DeepCopyInto(out *PulpRPMRemote) {
	*out = *in
	in.PulpCoreRemote.DeepCopyInto(&out.PulpCoreRemote)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRPMRemote.
func (in *PulpRPMRemote) DeepCopy() *PulpRPMRemote {
	if in == nil {
		return nil
	}
	out := new(PulpRPMRemote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRPMRepository) DeepCopyInto(out *PulpRPMRepository) {
	*out = *in
	in.PulpCoreRepository.DeepCopyInto(&out.PulpCoreRepository)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRPMRepository.
func (in *PulpRPMRepository) DeepCopy() *PulpRPMRepository {
	if in == nil {
		return nil
	}
	out := new(PulpRPMRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpResource) DeepCopyInto(out *PulpResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PulpResourceStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpResource.
func (in *PulpResource) DeepCopy() *PulpResource {
	if in == nil {
		return nil
	}
	out := new(PulpResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpResourceList) DeepCopyInto(out *PulpResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulpResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpResourceList.
func (in *PulpResourceList) DeepCopy() *PulpResourceList {
	if in == nil {
		return nil
	}
	out := new(PulpResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpResourceSpec) DeepCopyInto(out *PulpResourceSpec) {
	*out = *in
	if in.ClientConfig != nil {
		in, out := &in.ClientConfig, &out.ClientConfig
		*out = new(string)
		**out = **in
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(PulpContainerPlugin)
		(*in).DeepCopyInto(*out)
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(PulpFilePlugin)
		(*in).DeepCopyInto(*out)
	}
	if in.RPM != nil {
		in, out := &in.RPM, &out.RPM
		*out = new(PulpRPMPlugin)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpResourceSpec.
func (in *PulpResourceSpec) DeepCopy() *PulpResourceSpec {
	if in == nil {
		return nil
	}
	out := new(PulpResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpResourceStatus) DeepCopyInto(out *PulpResourceStatus) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(PulpContainerPlugin)
		(*in).DeepCopyInto(*out)
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(PulpFilePlugin)
		(*in).DeepCopyInto(*out)
	}
	if in.RPM != nil {
		in, out := &in.RPM, &out.RPM
		*out = new(PulpRPMPlugin)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpResourceStatus.
func (in *PulpResourceStatus) DeepCopy() *PulpResourceStatus {
	if in == nil {
		return nil
	}
	out := new(PulpResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRestore) DeepCopyInto(out *PulpRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRestore.
func (in *PulpRestore) DeepCopy() *PulpRestore {
	if in == nil {
		return nil
	}
	out := new(PulpRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRestoreList) DeepCopyInto(out *PulpRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PulpRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRestoreList.
func (in *PulpRestoreList) DeepCopy() *PulpRestoreList {
	if in == nil {
		return nil
	}
	out := new(PulpRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PulpRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRestoreSpec) DeepCopyInto(out *PulpRestoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRestoreSpec.
func (in *PulpRestoreSpec) DeepCopy() *PulpRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(PulpRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpRestoreStatus) DeepCopyInto(out *PulpRestoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpRestoreStatus.
func (in *PulpRestoreStatus) DeepCopy() *PulpRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(PulpRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpSpec) DeepCopyInto(out *PulpSpec) {
	*out = *in
	if in.IngressAnnotations != nil {
		in, out := &in.IngressAnnotations, &out.IngressAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RouteLabels != nil {
		in, out := &in.RouteLabels, &out.RouteLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RouteAnnotations != nil {
		in, out := &in.RouteAnnotations, &out.RouteAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Api.DeepCopyInto(&out.Api)
	in.Database.DeepCopyInto(&out.Database)
	in.Content.DeepCopyInto(&out.Content)
	in.Worker.DeepCopyInto(&out.Worker)
	in.Web.DeepCopyInto(&out.Web)
	in.Cache.DeepCopyInto(&out.Cache)
	in.PulpSettings.DeepCopyInto(&out.PulpSettings)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SAAnnotations != nil {
		in, out := &in.SAAnnotations, &out.SAAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SALabels != nil {
		in, out := &in.SALabels, &out.SALabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AdminPasswordJob.DeepCopyInto(&out.AdminPasswordJob)
	in.MigrationJob.DeepCopyInto(&out.MigrationJob)
	in.SigningJob.DeepCopyInto(&out.SigningJob)
	if in.AllowedContentChecksums != nil {
		in, out := &in.AllowedContentChecksums, &out.AllowedContentChecksums
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Telemetry.DeepCopyInto(&out.Telemetry)
	out.LDAP = in.LDAP
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpSpec.
func (in *PulpSpec) DeepCopy() *PulpSpec {
	if in == nil {
		return nil
	}
	out := new(PulpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpStatus) DeepCopyInto(out *PulpStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpStatus.
func (in *PulpStatus) DeepCopy() *PulpStatus {
	if in == nil {
		return nil
	}
	out := new(PulpStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulpSync) DeepCopyInto(out *PulpSync) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulpSync.
func (in *PulpSync) DeepCopy() *PulpSync {
	if in == nil {
		return nil
	}
	out := new(PulpSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Telemetry) DeepCopyInto(out *Telemetry) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Telemetry.
func (in *Telemetry) DeepCopy() *Telemetry {
	if in == nil {
		return nil
	}
	out := new(Telemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Web) DeepCopyInto(out *Web) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PDB != nil {
		in, out := &in.PDB, &out.PDB
		*out = new(policyv1.PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Web.
func (in *Web) DeepCopy() *Web {
	if in == nil {
		return nil
	}
	out := new(Web)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Worker) DeepCopyInto(out *Worker) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.PDB != nil {
		in, out := &in.PDB, &out.PDB
		*out = new(policyv1.PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Strategy.DeepCopyInto(&out.Strategy)
	in.InitContainer.DeepCopyInto(&out.InitContainer)
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentAnnotations != nil {
		in, out := &in.DeploymentAnnotations, &out.DeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Worker.
func (in *Worker) DeepCopy() *Worker {
	if in == nil {
		return nil
	}
	out := new(Worker)
	in.DeepCopyInto(out)
	return out
}
