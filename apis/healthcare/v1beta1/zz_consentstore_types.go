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

type ConsentStoreObservation struct {

	// an identifier for the resource with format {{dataset}}/consentStores/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConsentStoreParameters struct {

	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/healthcare/v1beta1.Dataset
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Reference to a Dataset in healthcare to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetRef *v1.Reference `json:"datasetRef,omitempty" tf:"-"`

	// Selector for a Dataset in healthcare to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetSelector *v1.Selector `json:"datasetSelector,omitempty" tf:"-"`

	// Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	DefaultConsentTTL *string `json:"defaultConsentTtl,omitempty" tf:"default_consent_ttl,omitempty"`

	// If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
	// +kubebuilder:validation:Optional
	EnableConsentCreateOnUpdate *bool `json:"enableConsentCreateOnUpdate,omitempty" tf:"enable_consent_create_on_update,omitempty"`

	// User-supplied key-value pairs used to organize Consent stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of this ConsentStore, for example:
	// "consent1"
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ConsentStoreSpec defines the desired state of ConsentStore
type ConsentStoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConsentStoreParameters `json:"forProvider"`
}

// ConsentStoreStatus defines the observed state of ConsentStore.
type ConsentStoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConsentStoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConsentStore is the Schema for the ConsentStores API. The Consent Management API is a tool for tracking user consents and the documentation associated with the consents.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ConsentStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConsentStoreSpec   `json:"spec"`
	Status            ConsentStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsentStoreList contains a list of ConsentStores
type ConsentStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsentStore `json:"items"`
}

// Repository type metadata.
var (
	ConsentStore_Kind             = "ConsentStore"
	ConsentStore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConsentStore_Kind}.String()
	ConsentStore_KindAPIVersion   = ConsentStore_Kind + "." + CRDGroupVersion.String()
	ConsentStore_GroupVersionKind = CRDGroupVersion.WithKind(ConsentStore_Kind)
)

func init() {
	SchemeBuilder.Register(&ConsentStore{}, &ConsentStoreList{})
}
