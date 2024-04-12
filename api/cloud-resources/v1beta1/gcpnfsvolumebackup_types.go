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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type GcpNfsBackupState string

// Valid GcpNfsBackup States.
const (
	// GcpNfsBackupReady signifies Backup is completed and is Ready for use.
	GcpNfsBackupReady GcpNfsBackupState = "Ready"

	// GcpNfsBackupInProgress signifies nfs backup is in-progress.
	GcpNfsBackupInProgress GcpNfsBackupState = "InProgress"

	// GcpNfsBackupError signifies the backup operation resulted in error.
	GcpNfsBackupError GcpNfsBackupState = "Error"
)

type GcpNfsVolumeBackupSource struct {
	// GcpNfsVolumeRef specifies the GcpNfsVolume resource that a backup has to be made of.
	// +kubebuilder:validation:Required
	Volume GcpNfsVolumeRef `json:"volume"`
}

type GcpNfsVolumeRef struct {
	// Name specifies the name of the GcpNfsVolume resource that a backup has to be made of.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Namespace specified the namespace of the GcpNfsVolume resource that a backup has to be made of.
	// If not specified then namespace of the GcpNfsVolumeBackup is used.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

func (v GcpNfsVolumeRef) ToNamespacedName(fallbackNamespace string) types.NamespacedName {
	ns := v.Namespace
	if len(ns) == 0 {
		ns = fallbackNamespace
	}
	return types.NamespacedName{
		Namespace: ns,
		Name:      v.Name,
	}
}

// GcpNfsVolumeBackupSpec defines the desired state of GcpNfsVolumeBackup
type GcpNfsVolumeBackupSpec struct {

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf), message="source is immutable."
	Source GcpNfsVolumeBackupSource `json:"source"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf), message="Location is immutable."
	Location string `json:"location"`
}

// GcpNfsVolumeBackupStatus defines the observed state of GcpNfsVolumeBackup
type GcpNfsVolumeBackupStatus struct {
	State GcpNfsBackupState `json:"state,omitempty"`

	// List of status conditions
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GcpNfsVolumeBackup is the Schema for the gcpnfsvolumebackups API
type GcpNfsVolumeBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GcpNfsVolumeBackupSpec   `json:"spec,omitempty"`
	Status GcpNfsVolumeBackupStatus `json:"status,omitempty"`
}

func (in *GcpNfsVolumeBackup) State() GcpNfsBackupState {
	return in.Status.State
}

func (in *GcpNfsVolumeBackup) SetState(v GcpNfsBackupState) {
	in.Status.State = v
}

func (in *GcpNfsVolumeBackup) Conditions() *[]metav1.Condition {
	return &in.Status.Conditions
}

func (in *GcpNfsVolumeBackup) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

//+kubebuilder:object:root=true

// GcpNfsVolumeBackupList contains a list of GcpNfsVolumeBackup
type GcpNfsVolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GcpNfsVolumeBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GcpNfsVolumeBackup{}, &GcpNfsVolumeBackupList{})
}
