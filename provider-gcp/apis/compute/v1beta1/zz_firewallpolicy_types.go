/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallPolicyObservation struct {

	// Creation timestamp in RFC3339 text format.
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Fingerprint of the resource. This field is used internally during updates of this resource.
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	FirewallPolicyID *string `json:"firewallPolicyId,omitempty" tf:"firewall_policy_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *float64 `json:"ruleTupleCount,omitempty" tf:"rule_tuple_count,omitempty"`

	// Server-defined URL for the resource.
	// Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Server-defined URL for this resource with the resource id.
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithID *string `json:"selfLinkWithId,omitempty" tf:"self_link_with_id,omitempty"`
}

type FirewallPolicyParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	// An optional description of this resource. Provide this property when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parent of the firewall policy.
	// The parent of the firewall policy.
	// +kubebuilder:validation:Required
	Parent *string `json:"parent" tf:"parent,omitempty"`

	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	ShortName *string `json:"shortName" tf:"short_name,omitempty"`
}

// FirewallPolicySpec defines the desired state of FirewallPolicy
type FirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyParameters `json:"forProvider"`
}

// FirewallPolicyStatus defines the observed state of FirewallPolicy.
type FirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicy is the Schema for the FirewallPolicys API. Creates a hierarchical firewall policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicySpec   `json:"spec"`
	Status            FirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyList contains a list of FirewallPolicys
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicy_Kind             = "FirewallPolicy"
	FirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicy_Kind}.String()
	FirewallPolicy_KindAPIVersion   = FirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicy{}, &FirewallPolicyList{})
}