// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LakeIAMPolicyInitParameters struct {

	// The policy data generated by
	// a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta2.Lake
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project",false)
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Lake in dataplex to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Lake in dataplex to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

type LakeIAMPolicyObservation struct {

	// (Computed) The etag of the IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Used to find the parent resource to bind the IAM policy to
	Lake *string `json:"lake,omitempty" tf:"lake,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The policy data generated by
	// a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type LakeIAMPolicyParameters struct {

	// Used to find the parent resource to bind the IAM policy to
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta2.Lake
	// +kubebuilder:validation:Optional
	Lake *string `json:"lake,omitempty" tf:"lake,omitempty"`

	// Reference to a Lake in dataplex to populate lake.
	// +kubebuilder:validation:Optional
	LakeRef *v1.Reference `json:"lakeRef,omitempty" tf:"-"`

	// Selector for a Lake in dataplex to populate lake.
	// +kubebuilder:validation:Optional
	LakeSelector *v1.Selector `json:"lakeSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta2.Lake
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("location",false)
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Reference to a Lake in dataplex to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Lake in dataplex to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// The policy data generated by
	// a google_iam_policy data source.
	// +kubebuilder:validation:Optional
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dataplex/v1beta2.Lake
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Lake in dataplex to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Lake in dataplex to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

// LakeIAMPolicySpec defines the desired state of LakeIAMPolicy
type LakeIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LakeIAMPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LakeIAMPolicyInitParameters `json:"initProvider,omitempty"`
}

// LakeIAMPolicyStatus defines the observed state of LakeIAMPolicy.
type LakeIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LakeIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LakeIAMPolicy is the Schema for the LakeIAMPolicys API. Collection of resources to manage IAM policy for Dataplex Lake
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type LakeIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyData) || (has(self.initProvider) && has(self.initProvider.policyData))",message="spec.forProvider.policyData is a required parameter"
	Spec   LakeIAMPolicySpec   `json:"spec"`
	Status LakeIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LakeIAMPolicyList contains a list of LakeIAMPolicys
type LakeIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LakeIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	LakeIAMPolicy_Kind             = "LakeIAMPolicy"
	LakeIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LakeIAMPolicy_Kind}.String()
	LakeIAMPolicy_KindAPIVersion   = LakeIAMPolicy_Kind + "." + CRDGroupVersion.String()
	LakeIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LakeIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LakeIAMPolicy{}, &LakeIAMPolicyList{})
}
