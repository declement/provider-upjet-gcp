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

type ApplicationInitParameters struct {

	// The domain to authenticate users with when using App Engine's User API.
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	// Can be CLOUD_FIRESTORE or CLOUD_DATASTORE_COMPATIBILITY for new
	// instances.  To support old instances, the value CLOUD_DATASTORE is accepted by the provider, but will be rejected by the API.
	// To create a Cloud Firestore database without creating an App Engine application, use the
	// google_firestore_database
	// resource instead.
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// A block of optional settings to configure specific App Engine features:
	FeatureSettings []FeatureSettingsInitParameters `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`

	// Settings for enabling Cloud Identity Aware Proxy
	Iap []IapInitParameters `json:"iap,omitempty" tf:"iap,omitempty"`

	// The location
	// to serve the app from.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// The serving status of the app.
	ServingStatus *string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
}

type ApplicationObservation struct {

	// Identifier of the app, usually {PROJECT_ID}
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// The domain to authenticate users with when using App Engine's User API.
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// The GCS bucket code is being stored in for this app.
	CodeBucket *string `json:"codeBucket,omitempty" tf:"code_bucket,omitempty"`

	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	// Can be CLOUD_FIRESTORE or CLOUD_DATASTORE_COMPATIBILITY for new
	// instances.  To support old instances, the value CLOUD_DATASTORE is accepted by the provider, but will be rejected by the API.
	// To create a Cloud Firestore database without creating an App Engine application, use the
	// google_firestore_database
	// resource instead.
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// The GCS bucket content is being stored in for this app.
	DefaultBucket *string `json:"defaultBucket,omitempty" tf:"default_bucket,omitempty"`

	// The default hostname for this app.
	DefaultHostname *string `json:"defaultHostname,omitempty" tf:"default_hostname,omitempty"`

	// A block of optional settings to configure specific App Engine features:
	FeatureSettings []FeatureSettingsObservation `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`

	// The GCR domain used for storing managed Docker images for this app.
	GcrDomain *string `json:"gcrDomain,omitempty" tf:"gcr_domain,omitempty"`

	// an identifier for the resource with format {{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Settings for enabling Cloud Identity Aware Proxy
	Iap []IapObservation `json:"iap,omitempty" tf:"iap,omitempty"`

	// The location
	// to serve the app from.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Unique name of the app, usually apps/{PROJECT_ID}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project ID to create the application under.
	// ~>NOTE: GCP only accepts project ID, not project number. If you are using number,
	// you may get a "Permission denied" error.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The serving status of the app.
	ServingStatus *string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`

	// A list of dispatch rule blocks. Each block has a domain, path, and service field.
	URLDispatchRule []URLDispatchRuleObservation `json:"urlDispatchRule,omitempty" tf:"url_dispatch_rule,omitempty"`
}

type ApplicationParameters struct {

	// The domain to authenticate users with when using App Engine's User API.
	// +kubebuilder:validation:Optional
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
	// Can be CLOUD_FIRESTORE or CLOUD_DATASTORE_COMPATIBILITY for new
	// instances.  To support old instances, the value CLOUD_DATASTORE is accepted by the provider, but will be rejected by the API.
	// To create a Cloud Firestore database without creating an App Engine application, use the
	// google_firestore_database
	// resource instead.
	// +kubebuilder:validation:Optional
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// A block of optional settings to configure specific App Engine features:
	// +kubebuilder:validation:Optional
	FeatureSettings []FeatureSettingsParameters `json:"featureSettings,omitempty" tf:"feature_settings,omitempty"`

	// Settings for enabling Cloud Identity Aware Proxy
	// +kubebuilder:validation:Optional
	Iap []IapParameters `json:"iap,omitempty" tf:"iap,omitempty"`

	// The location
	// to serve the app from.
	// +kubebuilder:validation:Optional
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// The project ID to create the application under.
	// ~>NOTE: GCP only accepts project ID, not project number. If you are using number,
	// you may get a "Permission denied" error.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// The serving status of the app.
	// +kubebuilder:validation:Optional
	ServingStatus *string `json:"servingStatus,omitempty" tf:"serving_status,omitempty"`
}

type FeatureSettingsInitParameters struct {

	// Set to false to use the legacy health check instead of the readiness
	// and liveness checks.
	SplitHealthChecks *bool `json:"splitHealthChecks,omitempty" tf:"split_health_checks,omitempty"`
}

type FeatureSettingsObservation struct {

	// Set to false to use the legacy health check instead of the readiness
	// and liveness checks.
	SplitHealthChecks *bool `json:"splitHealthChecks,omitempty" tf:"split_health_checks,omitempty"`
}

type FeatureSettingsParameters struct {

	// Set to false to use the legacy health check instead of the readiness
	// and liveness checks.
	// +kubebuilder:validation:Optional
	SplitHealthChecks *bool `json:"splitHealthChecks" tf:"split_health_checks,omitempty"`
}

type IapInitParameters struct {

	// Whether the serving infrastructure will authenticate and authorize all incoming requests.
	// (default is false)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// OAuth2 client ID to use for the authentication flow.
	Oauth2ClientID *string `json:"oauth2ClientId,omitempty" tf:"oauth2_client_id,omitempty"`
}

type IapObservation struct {

	// Whether the serving infrastructure will authenticate and authorize all incoming requests.
	// (default is false)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// OAuth2 client ID to use for the authentication flow.
	Oauth2ClientID *string `json:"oauth2ClientId,omitempty" tf:"oauth2_client_id,omitempty"`
}

type IapParameters struct {

	// Whether the serving infrastructure will authenticate and authorize all incoming requests.
	// (default is false)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// OAuth2 client ID to use for the authentication flow.
	// +kubebuilder:validation:Optional
	Oauth2ClientID *string `json:"oauth2ClientId" tf:"oauth2_client_id,omitempty"`

	// OAuth2 client secret to use for the authentication flow.
	// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	// +kubebuilder:validation:Required
	Oauth2ClientSecretSecretRef v1.SecretKeySelector `json:"oauth2ClientSecretSecretRef" tf:"-"`
}

type URLDispatchRuleInitParameters struct {
}

type URLDispatchRuleObservation struct {
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type URLDispatchRuleParameters struct {
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Allows management of an App Engine application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.locationId) || (has(self.initProvider) && has(self.initProvider.locationId))",message="spec.forProvider.locationId is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
