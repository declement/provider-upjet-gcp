// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type SSLCertInitParameters struct {

	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type SSLCertObservation struct {

	// The actual certificate data for this client certificate.
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	// The serial number extracted from the certificate data.
	CertSerialNumber *string `json:"certSerialNumber,omitempty" tf:"cert_serial_number,omitempty"`

	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The time when the certificate was created in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The time when the certificate expires in RFC 3339 format,
	// for example 2012-11-15T16:19:00.094Z.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The CA cert of the server this client cert was generated from.
	ServerCACert *string `json:"serverCaCert,omitempty" tf:"server_ca_cert,omitempty"`

	// The SHA1 Fingerprint of the certificate.
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty"`
}

type SSLCertParameters struct {

	// The common name to be used in the certificate to identify the
	// client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The name of the Cloud SQL instance. Changing this
	// forces a new resource to be created.
	// +crossplane:generate:reference:type=DatabaseInstance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a DatabaseInstance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// SSLCertSpec defines the desired state of SSLCert
type SSLCertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSLCertParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SSLCertInitParameters `json:"initProvider,omitempty"`
}

// SSLCertStatus defines the observed state of SSLCert.
type SSLCertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSLCertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SSLCert is the Schema for the SSLCerts API. Creates a new SQL Ssl Cert in Google Cloud SQL.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SSLCert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.commonName) || (has(self.initProvider) && has(self.initProvider.commonName))",message="spec.forProvider.commonName is a required parameter"
	Spec   SSLCertSpec   `json:"spec"`
	Status SSLCertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSLCertList contains a list of SSLCerts
type SSLCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSLCert `json:"items"`
}

// Repository type metadata.
var (
	SSLCert_Kind             = "SSLCert"
	SSLCert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSLCert_Kind}.String()
	SSLCert_KindAPIVersion   = SSLCert_Kind + "." + CRDGroupVersion.String()
	SSLCert_GroupVersionKind = CRDGroupVersion.WithKind(SSLCert_Kind)
)

func init() {
	SchemeBuilder.Register(&SSLCert{}, &SSLCertList{})
}
