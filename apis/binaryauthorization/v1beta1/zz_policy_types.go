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

type AdmissionWhitelistPatternsInitParameters struct {

	// An image name pattern to whitelist, in the form
	// registry/path/to/image. This supports a trailing * as a
	// wildcard, but this is allowed only in text after the registry/
	// part.
	NamePattern *string `json:"namePattern,omitempty" tf:"name_pattern,omitempty"`
}

type AdmissionWhitelistPatternsObservation struct {

	// An image name pattern to whitelist, in the form
	// registry/path/to/image. This supports a trailing * as a
	// wildcard, but this is allowed only in text after the registry/
	// part.
	NamePattern *string `json:"namePattern,omitempty" tf:"name_pattern,omitempty"`
}

type AdmissionWhitelistPatternsParameters struct {

	// An image name pattern to whitelist, in the form
	// registry/path/to/image. This supports a trailing * as a
	// wildcard, but this is allowed only in text after the registry/
	// part.
	// +kubebuilder:validation:Optional
	NamePattern *string `json:"namePattern" tf:"name_pattern,omitempty"`
}

type ClusterAdmissionRulesInitParameters struct {

	// The identifier for this object. Format specified above.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The action when a pod creation is denied by the admission rule.
	// Possible values are: ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY.
	EnforcementMode *string `json:"enforcementMode,omitempty" tf:"enforcement_mode,omitempty"`

	// How this admission rule will be evaluated.
	// Possible values are: ALWAYS_ALLOW, REQUIRE_ATTESTATION, ALWAYS_DENY.
	EvaluationMode *string `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	// The resource names of the attestors that must attest to a
	// container image. If the attestor is in a different project from the
	// policy, it should be specified in the format projects/*/attestors/*.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	RequireAttestationsBy []*string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type ClusterAdmissionRulesObservation struct {

	// The identifier for this object. Format specified above.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The action when a pod creation is denied by the admission rule.
	// Possible values are: ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY.
	EnforcementMode *string `json:"enforcementMode,omitempty" tf:"enforcement_mode,omitempty"`

	// How this admission rule will be evaluated.
	// Possible values are: ALWAYS_ALLOW, REQUIRE_ATTESTATION, ALWAYS_DENY.
	EvaluationMode *string `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	// The resource names of the attestors that must attest to a
	// container image. If the attestor is in a different project from the
	// policy, it should be specified in the format projects/*/attestors/*.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	RequireAttestationsBy []*string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type ClusterAdmissionRulesParameters struct {

	// The identifier for this object. Format specified above.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster" tf:"cluster,omitempty"`

	// The action when a pod creation is denied by the admission rule.
	// Possible values are: ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY.
	// +kubebuilder:validation:Optional
	EnforcementMode *string `json:"enforcementMode" tf:"enforcement_mode,omitempty"`

	// How this admission rule will be evaluated.
	// Possible values are: ALWAYS_ALLOW, REQUIRE_ATTESTATION, ALWAYS_DENY.
	// +kubebuilder:validation:Optional
	EvaluationMode *string `json:"evaluationMode" tf:"evaluation_mode,omitempty"`

	// The resource names of the attestors that must attest to a
	// container image. If the attestor is in a different project from the
	// policy, it should be specified in the format projects/*/attestors/*.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	// +kubebuilder:validation:Optional
	RequireAttestationsBy []*string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type DefaultAdmissionRuleInitParameters struct {

	// The action when a pod creation is denied by the admission rule.
	// Possible values are: ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY.
	EnforcementMode *string `json:"enforcementMode,omitempty" tf:"enforcement_mode,omitempty"`

	// How this admission rule will be evaluated.
	// Possible values are: ALWAYS_ALLOW, REQUIRE_ATTESTATION, ALWAYS_DENY.
	EvaluationMode *string `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	// The resource names of the attestors that must attest to a
	// container image. If the attestor is in a different project from the
	// policy, it should be specified in the format projects/*/attestors/*.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	RequireAttestationsBy []*string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type DefaultAdmissionRuleObservation struct {

	// The action when a pod creation is denied by the admission rule.
	// Possible values are: ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY.
	EnforcementMode *string `json:"enforcementMode,omitempty" tf:"enforcement_mode,omitempty"`

	// How this admission rule will be evaluated.
	// Possible values are: ALWAYS_ALLOW, REQUIRE_ATTESTATION, ALWAYS_DENY.
	EvaluationMode *string `json:"evaluationMode,omitempty" tf:"evaluation_mode,omitempty"`

	// The resource names of the attestors that must attest to a
	// container image. If the attestor is in a different project from the
	// policy, it should be specified in the format projects/*/attestors/*.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	RequireAttestationsBy []*string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type DefaultAdmissionRuleParameters struct {

	// The action when a pod creation is denied by the admission rule.
	// Possible values are: ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY.
	// +kubebuilder:validation:Optional
	EnforcementMode *string `json:"enforcementMode" tf:"enforcement_mode,omitempty"`

	// How this admission rule will be evaluated.
	// Possible values are: ALWAYS_ALLOW, REQUIRE_ATTESTATION, ALWAYS_DENY.
	// +kubebuilder:validation:Optional
	EvaluationMode *string `json:"evaluationMode" tf:"evaluation_mode,omitempty"`

	// The resource names of the attestors that must attest to a
	// container image. If the attestor is in a different project from the
	// policy, it should be specified in the format projects/*/attestors/*.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	// +kubebuilder:validation:Optional
	RequireAttestationsBy []*string `json:"requireAttestationsBy,omitempty" tf:"require_attestations_by,omitempty"`
}

type PolicyInitParameters struct {

	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns []AdmissionWhitelistPatternsInitParameters `json:"admissionWhitelistPatterns,omitempty" tf:"admission_whitelist_patterns,omitempty"`

	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules []ClusterAdmissionRulesInitParameters `json:"clusterAdmissionRules,omitempty" tf:"cluster_admission_rules,omitempty"`

	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule []DefaultAdmissionRuleInitParameters `json:"defaultAdmissionRule,omitempty" tf:"default_admission_rule,omitempty"`

	// A descriptive comment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are: ENABLE, DISABLE.
	GlobalPolicyEvaluationMode *string `json:"globalPolicyEvaluationMode,omitempty" tf:"global_policy_evaluation_mode,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type PolicyObservation struct {

	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns []AdmissionWhitelistPatternsObservation `json:"admissionWhitelistPatterns,omitempty" tf:"admission_whitelist_patterns,omitempty"`

	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules []ClusterAdmissionRulesObservation `json:"clusterAdmissionRules,omitempty" tf:"cluster_admission_rules,omitempty"`

	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule []DefaultAdmissionRuleObservation `json:"defaultAdmissionRule,omitempty" tf:"default_admission_rule,omitempty"`

	// A descriptive comment.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are: ENABLE, DISABLE.
	GlobalPolicyEvaluationMode *string `json:"globalPolicyEvaluationMode,omitempty" tf:"global_policy_evaluation_mode,omitempty"`

	// an identifier for the resource with format projects/{{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type PolicyParameters struct {

	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AdmissionWhitelistPatterns []AdmissionWhitelistPatternsParameters `json:"admissionWhitelistPatterns,omitempty" tf:"admission_whitelist_patterns,omitempty"`

	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	// +kubebuilder:validation:Optional
	ClusterAdmissionRules []ClusterAdmissionRulesParameters `json:"clusterAdmissionRules,omitempty" tf:"cluster_admission_rules,omitempty"`

	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DefaultAdmissionRule []DefaultAdmissionRuleParameters `json:"defaultAdmissionRule,omitempty" tf:"default_admission_rule,omitempty"`

	// A descriptive comment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are: ENABLE, DISABLE.
	// +kubebuilder:validation:Optional
	GlobalPolicyEvaluationMode *string `json:"globalPolicyEvaluationMode,omitempty" tf:"global_policy_evaluation_mode,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
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
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. A policy for container image binary authorization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultAdmissionRule) || (has(self.initProvider) && has(self.initProvider.defaultAdmissionRule))",message="spec.forProvider.defaultAdmissionRule is a required parameter"
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
