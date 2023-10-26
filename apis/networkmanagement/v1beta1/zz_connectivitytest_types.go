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

type ConnectivityTestInitParameters struct {

	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination []DestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Resource labels to represent user-provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Unique name for the connectivity test.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects []*string `json:"relatedProjects,omitempty" tf:"related_projects,omitempty"`

	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type ConnectivityTestObservation struct {

	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/connectivityTests/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user-provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Unique name for the connectivity test.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects []*string `json:"relatedProjects,omitempty" tf:"related_projects,omitempty"`

	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`
}

type ConnectivityTestParameters struct {

	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Unique name for the connectivity test.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// IP Protocol of the test. When not provided, "TCP" is assumed.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	// +kubebuilder:validation:Optional
	RelatedProjects []*string `json:"relatedProjects,omitempty" tf:"related_projects,omitempty"`

	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type DestinationInitParameters struct {

	// The IP protocol port of the endpoint. Only applicable when
	// protocol is TCP or UDP.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type DestinationObservation struct {

	// The IP address of the endpoint, which can be an external or
	// internal IP. An IPv6 address is only allowed when the test's
	// destination is a global load balancer VIP.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A Compute Engine instance URI.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// A Compute Engine network URI.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The IP protocol port of the endpoint. Only applicable when
	// protocol is TCP or UDP.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Project ID where the endpoint is located. The Project ID can be
	// derived from the URI if you provide a VM instance or network URI.
	// The following are two cases where you must provide the project ID:
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type DestinationParameters struct {

	// The IP address of the endpoint, which can be an external or
	// internal IP. An IPv6 address is only allowed when the test's
	// destination is a global load balancer VIP.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("address",false)
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Reference to a Address in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressRef *v1.Reference `json:"ipAddressRef,omitempty" tf:"-"`

	// Selector for a Address in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressSelector *v1.Selector `json:"ipAddressSelector,omitempty" tf:"-"`

	// A Compute Engine instance URI.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// A Compute Engine network URI.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The IP protocol port of the endpoint. Only applicable when
	// protocol is TCP or UDP.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Project ID where the endpoint is located. The Project ID can be
	// derived from the URI if you provide a VM instance or network URI.
	// The following are two cases where you must provide the project ID:
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Address in compute to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Address in compute to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type SourceInitParameters struct {

	// Type of the network where the endpoint is located.
	// Possible values are: GCP_NETWORK, NON_GCP_NETWORK.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The IP protocol port of the endpoint. Only applicable when
	// protocol is TCP or UDP.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type SourceObservation struct {

	// The IP address of the endpoint, which can be an external or
	// internal IP. An IPv6 address is only allowed when the test's
	// destination is a global load balancer VIP.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A Compute Engine instance URI.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// A Compute Engine network URI.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Type of the network where the endpoint is located.
	// Possible values are: GCP_NETWORK, NON_GCP_NETWORK.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The IP protocol port of the endpoint. Only applicable when
	// protocol is TCP or UDP.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Project ID where the endpoint is located. The Project ID can be
	// derived from the URI if you provide a VM instance or network URI.
	// The following are two cases where you must provide the project ID:
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type SourceParameters struct {

	// The IP address of the endpoint, which can be an external or
	// internal IP. An IPv6 address is only allowed when the test's
	// destination is a global load balancer VIP.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("address",false)
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Reference to a Address in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressRef *v1.Reference `json:"ipAddressRef,omitempty" tf:"-"`

	// Selector for a Address in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressSelector *v1.Selector `json:"ipAddressSelector,omitempty" tf:"-"`

	// A Compute Engine instance URI.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// A Compute Engine network URI.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// Type of the network where the endpoint is located.
	// Possible values are: GCP_NETWORK, NON_GCP_NETWORK.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The IP protocol port of the endpoint. Only applicable when
	// protocol is TCP or UDP.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Project ID where the endpoint is located. The Project ID can be
	// derived from the URI if you provide a VM instance or network URI.
	// The following are two cases where you must provide the project ID:
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Address
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Address in compute to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Address in compute to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

// ConnectivityTestSpec defines the desired state of ConnectivityTest
type ConnectivityTestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectivityTestParameters `json:"forProvider"`
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
	InitProvider ConnectivityTestInitParameters `json:"initProvider,omitempty"`
}

// ConnectivityTestStatus defines the observed state of ConnectivityTest.
type ConnectivityTestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectivityTestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectivityTest is the Schema for the ConnectivityTests API. A connectivity test are a static analysis of your resource configurations that enables you to evaluate connectivity to and from Google Cloud resources in your Virtual Private Cloud (VPC) network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ConnectivityTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   ConnectivityTestSpec   `json:"spec"`
	Status ConnectivityTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectivityTestList contains a list of ConnectivityTests
type ConnectivityTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectivityTest `json:"items"`
}

// Repository type metadata.
var (
	ConnectivityTest_Kind             = "ConnectivityTest"
	ConnectivityTest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectivityTest_Kind}.String()
	ConnectivityTest_KindAPIVersion   = ConnectivityTest_Kind + "." + CRDGroupVersion.String()
	ConnectivityTest_GroupVersionKind = CRDGroupVersion.WithKind(ConnectivityTest_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectivityTest{}, &ConnectivityTestList{})
}
