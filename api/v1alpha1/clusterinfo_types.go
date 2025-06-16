/*
Copyright 2025.

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

// NodeInfo holds information about a node
type NodeInfo struct {
	Name           string       `json:"name,omitempty"`
	InternalIP     string       `json:"internalIP,omitempty"`
	ExternalIP     string       `json:"externalIP,omitempty"`
	Capacity       NodeCapacity `json:"capacity,omitempty"`
	Allocatable    NodeCapacity `json:"allocatable,omitempty"`
	OSImage        string       `json:"osImage,omitempty"`
	KubeletVersion string       `json:"kubeletVersion,omitempty"`
}

// NodeCapacity describes resource capacity or allocatable
type NodeCapacity struct {
	CPU     string `json:"cpu,omitempty"`
	Memory  string `json:"memory,omitempty"`
	Pods    string `json:"pods,omitempty"`
	Storage string `json:"storage,omitempty"`
}

// GPUInfo describes GPU resources on a node
type GPUInfo struct {
	NodeName string `json:"nodeName,omitempty"`
	Model    string `json:"model,omitempty"`
	Count    int    `json:"count,omitempty"`
}

// ClusterInfoSpec defines the desired state of ClusterInfo.
type ClusterInfoSpec struct{}

type ClusterInfoStatus struct {
	KubernetesVersion string      `json:"kubernetesVersion,omitempty"`
	NodeCount         int         `json:"nodeCount,omitempty"`
	Nodes             []NodeInfo  `json:"nodes,omitempty"`
	PodCount          int         `json:"podCount,omitempty"`
	GPUs              []GPUInfo   `json:"gpus,omitempty"`
	LastUpdated       metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterInfo is the Schema for the clusterinfoes API.
type ClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterInfoSpec   `json:"spec,omitempty"`
	Status ClusterInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInfoList contains a list of ClusterInfo.
type ClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterInfo{}, &ClusterInfoList{})
}
