//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	in.EtcdBackupSpec.DeepCopyInto(&out.EtcdBackupSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImagePolicy) DeepCopyInto(out *ClusterImagePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImagePolicy.
func (in *ClusterImagePolicy) DeepCopy() *ClusterImagePolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterImagePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterImagePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImagePolicyList) DeepCopyInto(out *ClusterImagePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterImagePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImagePolicyList.
func (in *ClusterImagePolicyList) DeepCopy() *ClusterImagePolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterImagePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterImagePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImagePolicySpec) DeepCopyInto(out *ClusterImagePolicySpec) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]ImageScope, len(*in))
		copy(*out, *in)
	}
	in.Policy.DeepCopyInto(&out.Policy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImagePolicySpec.
func (in *ClusterImagePolicySpec) DeepCopy() *ClusterImagePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterImagePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImagePolicyStatus) DeepCopyInto(out *ClusterImagePolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImagePolicyStatus.
func (in *ClusterImagePolicyStatus) DeepCopy() *ClusterImagePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterImagePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMonitoring) DeepCopyInto(out *ClusterMonitoring) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMonitoring.
func (in *ClusterMonitoring) DeepCopy() *ClusterMonitoring {
	if in == nil {
		return nil
	}
	out := new(ClusterMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterMonitoring) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMonitoringList) DeepCopyInto(out *ClusterMonitoringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterMonitoring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMonitoringList.
func (in *ClusterMonitoringList) DeepCopy() *ClusterMonitoringList {
	if in == nil {
		return nil
	}
	out := new(ClusterMonitoringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterMonitoringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMonitoringSpec) DeepCopyInto(out *ClusterMonitoringSpec) {
	*out = *in
	out.UserDefined = in.UserDefined
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMonitoringSpec.
func (in *ClusterMonitoringSpec) DeepCopy() *ClusterMonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterMonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMonitoringStatus) DeepCopyInto(out *ClusterMonitoringStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMonitoringStatus.
func (in *ClusterMonitoringStatus) DeepCopy() *ClusterMonitoringStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterMonitoringStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdBackupSpec) DeepCopyInto(out *EtcdBackupSpec) {
	*out = *in
	in.RetentionPolicy.DeepCopyInto(&out.RetentionPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdBackupSpec.
func (in *EtcdBackupSpec) DeepCopy() *EtcdBackupSpec {
	if in == nil {
		return nil
	}
	out := new(EtcdBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulcioCAWithRekor) DeepCopyInto(out *FulcioCAWithRekor) {
	*out = *in
	if in.FulcioCAData != nil {
		in, out := &in.FulcioCAData, &out.FulcioCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RekorKeyData != nil {
		in, out := &in.RekorKeyData, &out.RekorKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	out.FulcioSubject = in.FulcioSubject
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulcioCAWithRekor.
func (in *FulcioCAWithRekor) DeepCopy() *FulcioCAWithRekor {
	if in == nil {
		return nil
	}
	out := new(FulcioCAWithRekor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatherConfig) DeepCopyInto(out *GatherConfig) {
	*out = *in
	if in.DisabledGatherers != nil {
		in, out := &in.DisabledGatherers, &out.DisabledGatherers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageSpec != nil {
		in, out := &in.StorageSpec, &out.StorageSpec
		*out = new(StorageSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatherConfig.
func (in *GatherConfig) DeepCopy() *GatherConfig {
	if in == nil {
		return nil
	}
	out := new(GatherConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicy) DeepCopyInto(out *ImagePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicy.
func (in *ImagePolicy) DeepCopy() *ImagePolicy {
	if in == nil {
		return nil
	}
	out := new(ImagePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyList) DeepCopyInto(out *ImagePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImagePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyList.
func (in *ImagePolicyList) DeepCopy() *ImagePolicyList {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicySpec) DeepCopyInto(out *ImagePolicySpec) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]ImageScope, len(*in))
		copy(*out, *in)
	}
	in.Policy.DeepCopyInto(&out.Policy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicySpec.
func (in *ImagePolicySpec) DeepCopy() *ImagePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ImagePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyStatus) DeepCopyInto(out *ImagePolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyStatus.
func (in *ImagePolicyStatus) DeepCopy() *ImagePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGather) DeepCopyInto(out *InsightsDataGather) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGather.
func (in *InsightsDataGather) DeepCopy() *InsightsDataGather {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGather)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InsightsDataGather) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGatherList) DeepCopyInto(out *InsightsDataGatherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InsightsDataGather, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGatherList.
func (in *InsightsDataGatherList) DeepCopy() *InsightsDataGatherList {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGatherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InsightsDataGatherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGatherSpec) DeepCopyInto(out *InsightsDataGatherSpec) {
	*out = *in
	in.GatherConfig.DeepCopyInto(&out.GatherConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGatherSpec.
func (in *InsightsDataGatherSpec) DeepCopy() *InsightsDataGatherSpec {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGatherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsDataGatherStatus) DeepCopyInto(out *InsightsDataGatherStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsDataGatherStatus.
func (in *InsightsDataGatherStatus) DeepCopy() *InsightsDataGatherStatus {
	if in == nil {
		return nil
	}
	out := new(InsightsDataGatherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKI) DeepCopyInto(out *PKI) {
	*out = *in
	if in.CertificateAuthorityRootsData != nil {
		in, out := &in.CertificateAuthorityRootsData, &out.CertificateAuthorityRootsData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CertificateAuthorityIntermediatesData != nil {
		in, out := &in.CertificateAuthorityIntermediatesData, &out.CertificateAuthorityIntermediatesData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	out.PKICertificateSubject = in.PKICertificateSubject
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKI.
func (in *PKI) DeepCopy() *PKI {
	if in == nil {
		return nil
	}
	out := new(PKI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKICertificateSubject) DeepCopyInto(out *PKICertificateSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKICertificateSubject.
func (in *PKICertificateSubject) DeepCopy() *PKICertificateSubject {
	if in == nil {
		return nil
	}
	out := new(PKICertificateSubject)
	in.DeepCopyInto(out)
	return out
}

func (in *PersistentVolumeClaimReference) DeepCopyInto(out *PersistentVolumeClaimReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimReference.
func (in *PersistentVolumeClaimReference) DeepCopy() *PersistentVolumeClaimReference {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	in.RootOfTrust.DeepCopyInto(&out.RootOfTrust)
	in.SignedIdentity.DeepCopyInto(&out.SignedIdentity)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFulcioSubject) DeepCopyInto(out *PolicyFulcioSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFulcioSubject.
func (in *PolicyFulcioSubject) DeepCopy() *PolicyFulcioSubject {
	if in == nil {
		return nil
	}
	out := new(PolicyFulcioSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIdentity) DeepCopyInto(out *PolicyIdentity) {
	*out = *in
	if in.PolicyMatchExactRepository != nil {
		in, out := &in.PolicyMatchExactRepository, &out.PolicyMatchExactRepository
		*out = new(PolicyMatchExactRepository)
		**out = **in
	}
	if in.PolicyMatchRemapIdentity != nil {
		in, out := &in.PolicyMatchRemapIdentity, &out.PolicyMatchRemapIdentity
		*out = new(PolicyMatchRemapIdentity)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIdentity.
func (in *PolicyIdentity) DeepCopy() *PolicyIdentity {
	if in == nil {
		return nil
	}
	out := new(PolicyIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyMatchExactRepository) DeepCopyInto(out *PolicyMatchExactRepository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyMatchExactRepository.
func (in *PolicyMatchExactRepository) DeepCopy() *PolicyMatchExactRepository {
	if in == nil {
		return nil
	}
	out := new(PolicyMatchExactRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyMatchRemapIdentity) DeepCopyInto(out *PolicyMatchRemapIdentity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyMatchRemapIdentity.
func (in *PolicyMatchRemapIdentity) DeepCopy() *PolicyMatchRemapIdentity {
	if in == nil {
		return nil
	}
	out := new(PolicyMatchRemapIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRootOfTrust) DeepCopyInto(out *PolicyRootOfTrust) {
	*out = *in
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(PublicKey)
		(*in).DeepCopyInto(*out)
	}
	if in.FulcioCAWithRekor != nil {
		in, out := &in.FulcioCAWithRekor, &out.FulcioCAWithRekor
		*out = new(FulcioCAWithRekor)
		(*in).DeepCopyInto(*out)
	}
	if in.PKI != nil {
		in, out := &in.PKI, &out.PKI
		*out = new(PKI)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRootOfTrust.
func (in *PolicyRootOfTrust) DeepCopy() *PolicyRootOfTrust {
	if in == nil {
		return nil
	}
	out := new(PolicyRootOfTrust)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKey) DeepCopyInto(out *PublicKey) {
	*out = *in
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RekorKeyData != nil {
		in, out := &in.RekorKeyData, &out.RekorKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKey.
func (in *PublicKey) DeepCopy() *PublicKey {
	if in == nil {
		return nil
	}
	out := new(PublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionNumberConfig) DeepCopyInto(out *RetentionNumberConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionNumberConfig.
func (in *RetentionNumberConfig) DeepCopy() *RetentionNumberConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionNumberConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicy) DeepCopyInto(out *RetentionPolicy) {
	*out = *in
	if in.RetentionNumber != nil {
		in, out := &in.RetentionNumber, &out.RetentionNumber
		*out = new(RetentionNumberConfig)
		**out = **in
	}
	if in.RetentionSize != nil {
		in, out := &in.RetentionSize, &out.RetentionSize
		*out = new(RetentionSizeConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicy.
func (in *RetentionPolicy) DeepCopy() *RetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionSizeConfig) DeepCopyInto(out *RetentionSizeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionSizeConfig.
func (in *RetentionSizeConfig) DeepCopy() *RetentionSizeConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionSizeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	out.PersistentVolumeClaim = in.PersistentVolumeClaim
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefinedMonitoring) DeepCopyInto(out *UserDefinedMonitoring) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefinedMonitoring.
func (in *UserDefinedMonitoring) DeepCopy() *UserDefinedMonitoring {
	if in == nil {
		return nil
	}
	out := new(UserDefinedMonitoring)
	in.DeepCopyInto(out)
	return out
}
