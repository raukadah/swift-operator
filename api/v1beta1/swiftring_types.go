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

package v1beta1

import (
	condition "github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	RingCreateHash = "ringcreate"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SwiftRingSpec defines the desired state of SwiftRing
type SwiftRingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// Number of Swift object replicas (=copies)
	Replicas int64 `json:"replicas"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// Number of devices/pods/nodes. Must be same or larger than replicas
	Devices int64 `json:"devices"`

	// +kubebuilder:validation:Required
	// Image URL for Swift proxy service
	ContainerImage string `json:"containerImage,omitempty"`

	// +kubebuilder:validation:Required
	// Storage Pod prefix
	StoragePodPrefix string `json:"storagePodPrefix,omitempty"`

	// +kubebuilder:validation:Required
	// Storage Service name
	StorageServiceName string `json:"storageServiceName,omitempty"`
}

// SwiftRingStatus defines the observed state of SwiftRing
type SwiftRingStatus struct {
	// Conditions
	Conditions condition.Conditions `json:"conditions,omitempty" optional:"true"`

	// Map of hashes to track e.g. job status
	Hash map[string]string `json:"hash,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[0].status",description="Status"
//+kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[0].message",description="Message"

// SwiftRing is the Schema for the swiftrings API
type SwiftRing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SwiftRingSpec   `json:"spec,omitempty"`
	Status SwiftRingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SwiftRingList contains a list of SwiftRing
type SwiftRingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SwiftRing `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SwiftRing{}, &SwiftRingList{})
}