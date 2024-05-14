// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupPolicyInitParameters struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// (String)
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type GroupPolicyObservation struct {

	// (String)
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// (String)
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type GroupPolicyParameters struct {

	// (String)
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Reference to a Group to populate group.
	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// Selector for a Group to populate group.
	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

// GroupPolicySpec defines the desired state of GroupPolicy
type GroupPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupPolicyParameters `json:"forProvider"`
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
	InitProvider GroupPolicyInitParameters `json:"initProvider,omitempty"`
}

// GroupPolicyStatus defines the observed state of GroupPolicy.
type GroupPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicy is the Schema for the GroupPolicys API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,minio}
type GroupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   GroupPolicySpec   `json:"spec"`
	Status GroupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicyList contains a list of GroupPolicys
type GroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupPolicy `json:"items"`
}

// Repository type metadata.
var (
	GroupPolicy_Kind             = "GroupPolicy"
	GroupPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupPolicy_Kind}.String()
	GroupPolicy_KindAPIVersion   = GroupPolicy_Kind + "." + CRDGroupVersion.String()
	GroupPolicy_GroupVersionKind = CRDGroupVersion.WithKind(GroupPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupPolicy{}, &GroupPolicyList{})
}