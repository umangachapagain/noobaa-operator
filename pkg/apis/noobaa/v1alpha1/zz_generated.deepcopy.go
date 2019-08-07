// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountsStatus) DeepCopyInto(out *AccountsStatus) {
	*out = *in
	out.Admin = in.Admin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountsStatus.
func (in *AccountsStatus) DeepCopy() *AccountsStatus {
	if in == nil {
		return nil
	}
	out := new(AccountsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStore) DeepCopyInto(out *BackingStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStore.
func (in *BackingStore) DeepCopy() *BackingStore {
	if in == nil {
		return nil
	}
	out := new(BackingStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackingStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreList) DeepCopyInto(out *BackingStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackingStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreList.
func (in *BackingStoreList) DeepCopy() *BackingStoreList {
	if in == nil {
		return nil
	}
	out := new(BackingStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackingStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreSpec) DeepCopyInto(out *BackingStoreSpec) {
	*out = *in
	out.Secret = in.Secret
	if in.S3Options != nil {
		in, out := &in.S3Options, &out.S3Options
		*out = new(S3Options)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreSpec.
func (in *BackingStoreSpec) DeepCopy() *BackingStoreSpec {
	if in == nil {
		return nil
	}
	out := new(BackingStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreStatus) DeepCopyInto(out *BackingStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreStatus.
func (in *BackingStoreStatus) DeepCopy() *BackingStoreStatus {
	if in == nil {
		return nil
	}
	out := new(BackingStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClass) DeepCopyInto(out *BucketClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClass.
func (in *BucketClass) DeepCopy() *BucketClass {
	if in == nil {
		return nil
	}
	out := new(BucketClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassList) DeepCopyInto(out *BucketClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassList.
func (in *BucketClassList) DeepCopy() *BucketClassList {
	if in == nil {
		return nil
	}
	out := new(BucketClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassSpec) DeepCopyInto(out *BucketClassSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassSpec.
func (in *BucketClassSpec) DeepCopy() *BucketClassSpec {
	if in == nil {
		return nil
	}
	out := new(BucketClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassStatus) DeepCopyInto(out *BucketClassStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassStatus.
func (in *BucketClassStatus) DeepCopy() *BucketClassStatus {
	if in == nil {
		return nil
	}
	out := new(BucketClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaa) DeepCopyInto(out *NooBaa) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaa.
func (in *NooBaa) DeepCopy() *NooBaa {
	if in == nil {
		return nil
	}
	out := new(NooBaa)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NooBaa) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaList) DeepCopyInto(out *NooBaaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NooBaa, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaList.
func (in *NooBaaList) DeepCopy() *NooBaaList {
	if in == nil {
		return nil
	}
	out := new(NooBaaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NooBaaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaSpec) DeepCopyInto(out *NooBaaSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.MongoImage != nil {
		in, out := &in.MongoImage, &out.MongoImage
		*out = new(string)
		**out = **in
	}
	if in.CoreResources != nil {
		in, out := &in.CoreResources, &out.CoreResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.MongoResources != nil {
		in, out := &in.MongoResources, &out.MongoResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaSpec.
func (in *NooBaaSpec) DeepCopy() *NooBaaSpec {
	if in == nil {
		return nil
	}
	out := new(NooBaaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaStatus) DeepCopyInto(out *NooBaaStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	out.Accounts = in.Accounts
	in.Services.DeepCopyInto(&out.Services)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaStatus.
func (in *NooBaaStatus) DeepCopy() *NooBaaStatus {
	if in == nil {
		return nil
	}
	out := new(NooBaaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Options) DeepCopyInto(out *S3Options) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Options.
func (in *S3Options) DeepCopy() *S3Options {
	if in == nil {
		return nil
	}
	out := new(S3Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
	if in.NodePorts != nil {
		in, out := &in.NodePorts, &out.NodePorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodPorts != nil {
		in, out := &in.PodPorts, &out.PodPorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InternalIP != nil {
		in, out := &in.InternalIP, &out.InternalIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InternalDNS != nil {
		in, out := &in.InternalDNS, &out.InternalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesStatus) DeepCopyInto(out *ServicesStatus) {
	*out = *in
	in.ServiceMgmt.DeepCopyInto(&out.ServiceMgmt)
	in.ServiceS3.DeepCopyInto(&out.ServiceS3)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesStatus.
func (in *ServicesStatus) DeepCopy() *ServicesStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}
