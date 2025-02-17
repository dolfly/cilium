// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories={cilium},singular="ciliumgatewayclassconfig",path="ciliumgatewayclassconfigs",scope="Cluster",shortName={cgcc}
// +kubebuilder:printcolumn:name="Accepted",type=string,JSONPath=`.status.conditions[?(@.type=="Accepted")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`,priority=1
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CiliumGatewayClassConfig is a Kubernetes third-party resource which
// is used to configure Gateways owned by GatewayClass.
//
// This is a cluster-scoped resource.
type CiliumGatewayClassConfig struct {
	// +deepequal-gen=false
	metav1.TypeMeta `json:",inline"`
	// +deepequal-gen=false
	metav1.ObjectMeta `json:"metadata"`

	// Spec is a human-readable of a GatewayClass configuration.
	//
	// +kubebuilder:validation:Optional
	Spec CiliumGatewayClassConfigSpec `json:"spec,omitempty"`

	// Status is the status of the policy.
	//
	// +deepequal-gen=false
	// +kubebuilder:validation:Optional
	Status CiliumGatewayClassConfigStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +deepequal-gen=false

// CiliumGatewayClassConfigList is a list of
// CiliumGatewayClassConfig objects.
type CiliumGatewayClassConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is a list of CiliumL2AnnouncementPolicies.
	Items []CiliumGatewayClassConfig `json:"items"`
}

// +deepequal-gen=true

type ServiceConfig struct {
	// Type specifies the type of the Service. Defaults to LoadBalancer.
	//
	// +kubebuilder:default=LoadBalancer
	Type *string `json:"type,omitempty"`
}

// CiliumGatewayClassConfigSpec specifies all the configuration options for a
// Cilium managed GatewayClass.
type CiliumGatewayClassConfigSpec struct {
	// Description helps describe a GatewayClass configuration with more details.
	//
	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Service specifies the configuration for the generated Service.
	// Note that not all fields from upstream Service.Spec are supported
	//
	// +kubebuilder:default=LoadBalancer
	// +kubebuilder:validation:Optional
	Service *ServiceConfig `json:"service,omitempty"`
}

// +deepequal-gen=false

// CiliumGatewayClassConfigStatus contains the status of a CiliumGatewayClassConfig.
type CiliumGatewayClassConfigStatus struct {
	// Current service state
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}
