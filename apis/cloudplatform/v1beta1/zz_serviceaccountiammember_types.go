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

type ServiceAccountIAMMemberConditionObservation struct {
}

type ServiceAccountIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type ServiceAccountIAMMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceAccountIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ServiceAccountIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

// ServiceAccountIAMMemberSpec defines the desired state of ServiceAccountIAMMember
type ServiceAccountIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountIAMMemberParameters `json:"forProvider"`
}

// ServiceAccountIAMMemberStatus defines the observed state of ServiceAccountIAMMember.
type ServiceAccountIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIAMMember is the Schema for the ServiceAccountIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceAccountIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIAMMemberSpec   `json:"spec"`
	Status            ServiceAccountIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIAMMemberList contains a list of ServiceAccountIAMMembers
type ServiceAccountIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountIAMMember `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountIAMMember_Kind             = "ServiceAccountIAMMember"
	ServiceAccountIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountIAMMember_Kind}.String()
	ServiceAccountIAMMember_KindAPIVersion   = ServiceAccountIAMMember_Kind + "." + CRDGroupVersion.String()
	ServiceAccountIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountIAMMember{}, &ServiceAccountIAMMemberList{})
}
