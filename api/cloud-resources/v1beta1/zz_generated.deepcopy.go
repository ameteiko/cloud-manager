//go:build !ignore_autogenerated

/*
Copyright 2023.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolume) DeepCopyInto(out *AwsNfsVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolume.
func (in *AwsNfsVolume) DeepCopy() *AwsNfsVolume {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsNfsVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeBackup) DeepCopyInto(out *AwsNfsVolumeBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeBackup.
func (in *AwsNfsVolumeBackup) DeepCopy() *AwsNfsVolumeBackup {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsNfsVolumeBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeBackupLifecycle) DeepCopyInto(out *AwsNfsVolumeBackupLifecycle) {
	*out = *in
	if in.DeleteAfterDays != nil {
		in, out := &in.DeleteAfterDays, &out.DeleteAfterDays
		*out = new(int64)
		**out = **in
	}
	if in.MoveToColdStorageAfterDays != nil {
		in, out := &in.MoveToColdStorageAfterDays, &out.MoveToColdStorageAfterDays
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeBackupLifecycle.
func (in *AwsNfsVolumeBackupLifecycle) DeepCopy() *AwsNfsVolumeBackupLifecycle {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeBackupLifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeBackupList) DeepCopyInto(out *AwsNfsVolumeBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AwsNfsVolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeBackupList.
func (in *AwsNfsVolumeBackupList) DeepCopy() *AwsNfsVolumeBackupList {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsNfsVolumeBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeBackupSource) DeepCopyInto(out *AwsNfsVolumeBackupSource) {
	*out = *in
	out.Volume = in.Volume
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeBackupSource.
func (in *AwsNfsVolumeBackupSource) DeepCopy() *AwsNfsVolumeBackupSource {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeBackupSpec) DeepCopyInto(out *AwsNfsVolumeBackupSpec) {
	*out = *in
	out.Source = in.Source
	in.Lifecycle.DeepCopyInto(&out.Lifecycle)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeBackupSpec.
func (in *AwsNfsVolumeBackupSpec) DeepCopy() *AwsNfsVolumeBackupSpec {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeBackupStatus) DeepCopyInto(out *AwsNfsVolumeBackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeBackupStatus.
func (in *AwsNfsVolumeBackupStatus) DeepCopy() *AwsNfsVolumeBackupStatus {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeList) DeepCopyInto(out *AwsNfsVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AwsNfsVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeList.
func (in *AwsNfsVolumeList) DeepCopy() *AwsNfsVolumeList {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsNfsVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumePvSpec) DeepCopyInto(out *AwsNfsVolumePvSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumePvSpec.
func (in *AwsNfsVolumePvSpec) DeepCopy() *AwsNfsVolumePvSpec {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumePvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeSpec) DeepCopyInto(out *AwsNfsVolumeSpec) {
	*out = *in
	out.IpRange = in.IpRange
	out.Capacity = in.Capacity.DeepCopy()
	if in.PersistentVolume != nil {
		in, out := &in.PersistentVolume, &out.PersistentVolume
		*out = new(AwsNfsVolumePvSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeSpec.
func (in *AwsNfsVolumeSpec) DeepCopy() *AwsNfsVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsNfsVolumeStatus) DeepCopyInto(out *AwsNfsVolumeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsNfsVolumeStatus.
func (in *AwsNfsVolumeStatus) DeepCopy() *AwsNfsVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(AwsNfsVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResources) DeepCopyInto(out *CloudResources) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResources.
func (in *CloudResources) DeepCopy() *CloudResources {
	if in == nil {
		return nil
	}
	out := new(CloudResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudResources) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourcesList) DeepCopyInto(out *CloudResourcesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourcesList.
func (in *CloudResourcesList) DeepCopy() *CloudResourcesList {
	if in == nil {
		return nil
	}
	out := new(CloudResourcesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudResourcesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourcesSpec) DeepCopyInto(out *CloudResourcesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourcesSpec.
func (in *CloudResourcesSpec) DeepCopy() *CloudResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(CloudResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudResourcesStatus) DeepCopyInto(out *CloudResourcesStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudResourcesStatus.
func (in *CloudResourcesStatus) DeepCopy() *CloudResourcesStatus {
	if in == nil {
		return nil
	}
	out := new(CloudResourcesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolume) DeepCopyInto(out *GcpNfsVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolume.
func (in *GcpNfsVolume) DeepCopy() *GcpNfsVolume {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GcpNfsVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeBackup) DeepCopyInto(out *GcpNfsVolumeBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeBackup.
func (in *GcpNfsVolumeBackup) DeepCopy() *GcpNfsVolumeBackup {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GcpNfsVolumeBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeBackupList) DeepCopyInto(out *GcpNfsVolumeBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GcpNfsVolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeBackupList.
func (in *GcpNfsVolumeBackupList) DeepCopy() *GcpNfsVolumeBackupList {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GcpNfsVolumeBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeBackupSource) DeepCopyInto(out *GcpNfsVolumeBackupSource) {
	*out = *in
	out.Volume = in.Volume
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeBackupSource.
func (in *GcpNfsVolumeBackupSource) DeepCopy() *GcpNfsVolumeBackupSource {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeBackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeBackupSpec) DeepCopyInto(out *GcpNfsVolumeBackupSpec) {
	*out = *in
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeBackupSpec.
func (in *GcpNfsVolumeBackupSpec) DeepCopy() *GcpNfsVolumeBackupSpec {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeBackupStatus) DeepCopyInto(out *GcpNfsVolumeBackupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeBackupStatus.
func (in *GcpNfsVolumeBackupStatus) DeepCopy() *GcpNfsVolumeBackupStatus {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeList) DeepCopyInto(out *GcpNfsVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GcpNfsVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeList.
func (in *GcpNfsVolumeList) DeepCopy() *GcpNfsVolumeList {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GcpNfsVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeRef) DeepCopyInto(out *GcpNfsVolumeRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeRef.
func (in *GcpNfsVolumeRef) DeepCopy() *GcpNfsVolumeRef {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeSpec) DeepCopyInto(out *GcpNfsVolumeSpec) {
	*out = *in
	out.IpRange = in.IpRange
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeSpec.
func (in *GcpNfsVolumeSpec) DeepCopy() *GcpNfsVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpNfsVolumeStatus) DeepCopyInto(out *GcpNfsVolumeStatus) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpNfsVolumeStatus.
func (in *GcpNfsVolumeStatus) DeepCopy() *GcpNfsVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(GcpNfsVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRange) DeepCopyInto(out *IpRange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRange.
func (in *IpRange) DeepCopy() *IpRange {
	if in == nil {
		return nil
	}
	out := new(IpRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpRange) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRangeList) DeepCopyInto(out *IpRangeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IpRange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRangeList.
func (in *IpRangeList) DeepCopy() *IpRangeList {
	if in == nil {
		return nil
	}
	out := new(IpRangeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IpRangeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRangeRef) DeepCopyInto(out *IpRangeRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRangeRef.
func (in *IpRangeRef) DeepCopy() *IpRangeRef {
	if in == nil {
		return nil
	}
	out := new(IpRangeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRangeSpec) DeepCopyInto(out *IpRangeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRangeSpec.
func (in *IpRangeSpec) DeepCopy() *IpRangeSpec {
	if in == nil {
		return nil
	}
	out := new(IpRangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpRangeStatus) DeepCopyInto(out *IpRangeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpRangeStatus.
func (in *IpRangeStatus) DeepCopy() *IpRangeStatus {
	if in == nil {
		return nil
	}
	out := new(IpRangeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeRef) DeepCopyInto(out *VolumeRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeRef.
func (in *VolumeRef) DeepCopy() *VolumeRef {
	if in == nil {
		return nil
	}
	out := new(VolumeRef)
	in.DeepCopyInto(out)
	return out
}
