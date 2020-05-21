// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapability) DeepCopyInto(out *ProvisionerCapability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapability.
func (in *ProvisionerCapability) DeepCopy() *ProvisionerCapability {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionerCapability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapabilityList) DeepCopyInto(out *ProvisionerCapabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProvisionerCapability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapabilityList.
func (in *ProvisionerCapabilityList) DeepCopy() *ProvisionerCapabilityList {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionerCapabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapabilitySpec) DeepCopyInto(out *ProvisionerCapabilitySpec) {
	*out = *in
	out.PluginInfo = in.PluginInfo
	out.Features = in.Features
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapabilitySpec.
func (in *ProvisionerCapabilitySpec) DeepCopy() *ProvisionerCapabilitySpec {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapabilitySpecFeatures) DeepCopyInto(out *ProvisionerCapabilitySpecFeatures) {
	*out = *in
	out.Volume = in.Volume
	out.Snapshot = in.Snapshot
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapabilitySpecFeatures.
func (in *ProvisionerCapabilitySpecFeatures) DeepCopy() *ProvisionerCapabilitySpecFeatures {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapabilitySpecFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapabilitySpecFeaturesSnapshot) DeepCopyInto(out *ProvisionerCapabilitySpecFeaturesSnapshot) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapabilitySpecFeaturesSnapshot.
func (in *ProvisionerCapabilitySpecFeaturesSnapshot) DeepCopy() *ProvisionerCapabilitySpecFeaturesSnapshot {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapabilitySpecFeaturesSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapabilitySpecFeaturesVolume) DeepCopyInto(out *ProvisionerCapabilitySpecFeaturesVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapabilitySpecFeaturesVolume.
func (in *ProvisionerCapabilitySpecFeaturesVolume) DeepCopy() *ProvisionerCapabilitySpecFeaturesVolume {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapabilitySpecFeaturesVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerCapabilitySpecPluginInfo) DeepCopyInto(out *ProvisionerCapabilitySpecPluginInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerCapabilitySpecPluginInfo.
func (in *ProvisionerCapabilitySpecPluginInfo) DeepCopy() *ProvisionerCapabilitySpecPluginInfo {
	if in == nil {
		return nil
	}
	out := new(ProvisionerCapabilitySpecPluginInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassCapability) DeepCopyInto(out *StorageClassCapability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassCapability.
func (in *StorageClassCapability) DeepCopy() *StorageClassCapability {
	if in == nil {
		return nil
	}
	out := new(StorageClassCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClassCapability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassCapabilityList) DeepCopyInto(out *StorageClassCapabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageClassCapability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassCapabilityList.
func (in *StorageClassCapabilityList) DeepCopy() *StorageClassCapabilityList {
	if in == nil {
		return nil
	}
	out := new(StorageClassCapabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClassCapabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassCapabilitySpec) DeepCopyInto(out *StorageClassCapabilitySpec) {
	*out = *in
	out.Features = in.Features
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassCapabilitySpec.
func (in *StorageClassCapabilitySpec) DeepCopy() *StorageClassCapabilitySpec {
	if in == nil {
		return nil
	}
	out := new(StorageClassCapabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassCapabilitySpecFeatures) DeepCopyInto(out *StorageClassCapabilitySpecFeatures) {
	*out = *in
	out.Volume = in.Volume
	out.Snapshot = in.Snapshot
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassCapabilitySpecFeatures.
func (in *StorageClassCapabilitySpecFeatures) DeepCopy() *StorageClassCapabilitySpecFeatures {
	if in == nil {
		return nil
	}
	out := new(StorageClassCapabilitySpecFeatures)
	in.DeepCopyInto(out)
	return out
}
