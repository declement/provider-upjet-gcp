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

type AutoscalingConfigInitParameters struct {

	// The target CPU utilization for autoscaling, in percentage. Must be between 10 and 80.
	CPUTarget *float64 `json:"cpuTarget,omitempty" tf:"cpu_target,omitempty"`

	// The maximum number of nodes for autoscaling.
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// The minimum number of nodes for autoscaling.
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The target storage utilization for autoscaling, in GB, for each node in a cluster. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters.
	StorageTarget *float64 `json:"storageTarget,omitempty" tf:"storage_target,omitempty"`
}

type AutoscalingConfigObservation struct {

	// The target CPU utilization for autoscaling, in percentage. Must be between 10 and 80.
	CPUTarget *float64 `json:"cpuTarget,omitempty" tf:"cpu_target,omitempty"`

	// The maximum number of nodes for autoscaling.
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// The minimum number of nodes for autoscaling.
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The target storage utilization for autoscaling, in GB, for each node in a cluster. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters.
	StorageTarget *float64 `json:"storageTarget,omitempty" tf:"storage_target,omitempty"`
}

type AutoscalingConfigParameters struct {

	// The target CPU utilization for autoscaling, in percentage. Must be between 10 and 80.
	// +kubebuilder:validation:Optional
	CPUTarget *float64 `json:"cpuTarget" tf:"cpu_target,omitempty"`

	// The maximum number of nodes for autoscaling.
	// +kubebuilder:validation:Optional
	MaxNodes *float64 `json:"maxNodes" tf:"max_nodes,omitempty"`

	// The minimum number of nodes for autoscaling.
	// +kubebuilder:validation:Optional
	MinNodes *float64 `json:"minNodes" tf:"min_nodes,omitempty"`

	// The target storage utilization for autoscaling, in GB, for each node in a cluster. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters.
	// +kubebuilder:validation:Optional
	StorageTarget *float64 `json:"storageTarget,omitempty" tf:"storage_target,omitempty"`
}

type ClusterInitParameters struct {

	// Autoscaling config for the cluster, contains the following arguments:
	AutoscalingConfig []AutoscalingConfigInitParameters `json:"autoscalingConfig,omitempty" tf:"autoscaling_config,omitempty"`

	// The ID of the Cloud Bigtable cluster. Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster. The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The number of nodes in the cluster.
	// If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization.
	NumNodes *float64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// The storage type to use. One of "SSD" or
	// "HDD". Defaults to "SSD".
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// The zone to create the Cloud Bigtable cluster in. If it not
	// specified, the provider zone is used. Each cluster must have a different zone in the same region. Zones that support
	// Bigtable instances are noted on the Cloud Bigtable locations page.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ClusterObservation struct {

	// Autoscaling config for the cluster, contains the following arguments:
	AutoscalingConfig []AutoscalingConfigObservation `json:"autoscalingConfig,omitempty" tf:"autoscaling_config,omitempty"`

	// The ID of the Cloud Bigtable cluster. Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster. The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The number of nodes in the cluster.
	// If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization.
	NumNodes *float64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// The storage type to use. One of "SSD" or
	// "HDD". Defaults to "SSD".
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// The zone to create the Cloud Bigtable cluster in. If it not
	// specified, the provider zone is used. Each cluster must have a different zone in the same region. Zones that support
	// Bigtable instances are noted on the Cloud Bigtable locations page.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ClusterParameters struct {

	// Autoscaling config for the cluster, contains the following arguments:
	// +kubebuilder:validation:Optional
	AutoscalingConfig []AutoscalingConfigParameters `json:"autoscalingConfig,omitempty" tf:"autoscaling_config,omitempty"`

	// The ID of the Cloud Bigtable cluster. Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster. The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The number of nodes in the cluster.
	// If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization.
	// +kubebuilder:validation:Optional
	NumNodes *float64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// The storage type to use. One of "SSD" or
	// "HDD". Defaults to "SSD".
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// The zone to create the Cloud Bigtable cluster in. If it not
	// specified, the provider zone is used. Each cluster must have a different zone in the same region. Zones that support
	// Bigtable instances are noted on the Cloud Bigtable locations page.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceInitParameters struct {

	// A block of cluster configuration options. This can be specified at least once, and up
	// to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
	// to default to the backend value. See structure below.
	Cluster []ClusterInitParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`

	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The human-readable display name of the Bigtable instance. Defaults to the instance name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION".
	// It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away,
	// and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to
	// "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance
	// is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type InstanceObservation struct {

	// A block of cluster configuration options. This can be specified at least once, and up
	// to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
	// to default to the backend value. See structure below.
	Cluster []ClusterObservation `json:"cluster,omitempty" tf:"cluster,omitempty"`

	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The human-readable display name of the Bigtable instance. Defaults to the instance name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION".
	// It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away,
	// and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to
	// "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance
	// is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type InstanceParameters struct {

	// A block of cluster configuration options. This can be specified at least once, and up
	// to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
	// to default to the backend value. See structure below.
	// +kubebuilder:validation:Optional
	Cluster []ClusterParameters `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The human-readable display name of the Bigtable instance. Defaults to the instance name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The instance type to create. One of "DEVELOPMENT" or "PRODUCTION". Defaults to "PRODUCTION".
	// It is recommended to leave this field unspecified since the distinction between "DEVELOPMENT" and "PRODUCTION" instances is going away,
	// and all instances will become "PRODUCTION" instances. This means that new and existing "DEVELOPMENT" instances will be converted to
	// "PRODUCTION" instances. It is recommended for users to use "PRODUCTION" instances in any case, since a 1-node "PRODUCTION" instance
	// is functionally identical to a "DEVELOPMENT" instance, but without the accompanying restrictions.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
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
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Creates a Google Bigtable instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
