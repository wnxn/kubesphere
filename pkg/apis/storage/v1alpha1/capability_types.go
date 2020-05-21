/*
Copyright 2019 The KubeSphere authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:noStatus
// +genclient:nonNamespaced

// ProvisionerCapability is the schema for the provisionercapability API
// +k8s:openapi-gen=true
type ProvisionerCapability struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProvisionerCapabilitySpec `json:"spec"`
}

// ProvisionerCapabilitySpec defines the desired state of ProvisionerCapability
type ProvisionerCapabilitySpec struct {
	PluginInfo ProvisionerCapabilitySpecPluginInfo `json:"pluginInfo"`
	Features   ProvisionerCapabilitySpecFeatures   `json:"features"`
}

// ProvisionerCapabilitySpecPluginInfo describes plugin info
type ProvisionerCapabilitySpecPluginInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ProvisionerCapabilitySpecFeatures describes plugin features
type ProvisionerCapabilitySpecFeatures struct {
	Topology bool                                      `json:"topology"`
	Volume   ProvisionerCapabilitySpecFeaturesVolume   `json:"volume"`
	Snapshot ProvisionerCapabilitySpecFeaturesSnapshot `json:"snapshot"`
}

// ProvisionerCapabilitySpecFeaturesVolume describe volume features
type ProvisionerCapabilitySpecFeaturesVolume struct {
	Create bool       `json:"create"`
	Attach bool       `json:"attach"`
	List   bool       `json:"list"`
	Clone  bool       `json:"clone"`
	Stats  bool       `json:"stats"`
	Expand ExpandMode `json:"expandMode"`
}

// ProvisionerCapabilitySpecFeaturesSnapshot describe snapshot features
type ProvisionerCapabilitySpecFeaturesSnapshot struct {
	Create bool `json:"create"`
	List   bool `json:"list"`
}

type ExpandMode string

const (
	ExpandModeUnknown ExpandMode = "UNKNOWN"
	ExpandModeOffline ExpandMode = "OFFLINE"
	ExpandModeOnline  ExpandMode = "ONLINE"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ProvisionerCapabilityList contains a list of ProvisionerCapability
type ProvisionerCapabilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ProvisionerCapability `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:noStatus
// +genclient:nonNamespaced

// StorageClassCapability is the Schema for the storage class capability API
// +k8s:openapi-gen=true
type StorageClassCapability struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec StorageClassCapabilitySpec `json:"spec"`
}

// StorageClassCapabilitySpec defines the desired state of StorageClassCapability
type StorageClassCapabilitySpec struct {
	Provisioner string                             `json:"provisioner"`
	Features    StorageClassCapabilitySpecFeatures `json:"features"`
}

// StorageClassCapabilitySpecFeatures describe storage class features
type StorageClassCapabilitySpecFeatures struct {
	Topology bool                                      `json:"topology"`
	Volume   ProvisionerCapabilitySpecFeaturesVolume   `json:"volume"`
	Snapshot ProvisionerCapabilitySpecFeaturesSnapshot `json:"snapshot"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// StorageClassCapabilityList contains a list of StorageClassCapability
type StorageClassCapabilityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []StorageClassCapability `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ProvisionerCapability{},
		&ProvisionerCapabilityList{},
		&StorageClassCapability{},
		&StorageClassCapabilityList{})
}
