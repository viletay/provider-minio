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

type PolicyInitParameters struct {

	// (Block List, Min: 1) (see below for nested schema)
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type PolicyObservation struct {

	// (String)
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Min: 1) (see below for nested schema)
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type PolicyParameters struct {

	// (String)
	// +crossplane:generate:reference:type=github.com/viletay/provider-minio/apis/s3/v1alpha1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// (Block List, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleInitParameters struct {

	// (String)
	// Value may be duration (5d), date (1970-01-01), or "DeleteMarker" to expire delete markers if `noncurrent_version_expiration_days` is used
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Int)
	NoncurrentVersionExpirationDays *float64 `json:"noncurrentVersionExpirationDays,omitempty" tf:"noncurrent_version_expiration_days,omitempty"`

	// (Map of String)
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleObservation struct {

	// (String)
	// Value may be duration (5d), date (1970-01-01), or "DeleteMarker" to expire delete markers if `noncurrent_version_expiration_days` is used
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// (String)
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Int)
	NoncurrentVersionExpirationDays *float64 `json:"noncurrentVersionExpirationDays,omitempty" tf:"noncurrent_version_expiration_days,omitempty"`

	// (String)
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Map of String)
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RuleParameters struct {

	// (String)
	// Value may be duration (5d), date (1970-01-01), or "DeleteMarker" to expire delete markers if `noncurrent_version_expiration_days` is used
	// +kubebuilder:validation:Optional
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// (String) The ID of this resource.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// (Int)
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpirationDays *float64 `json:"noncurrentVersionExpirationDays,omitempty" tf:"noncurrent_version_expiration_days,omitempty"`

	// (Map of String)
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
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
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. minio_ilm_policy handles lifecycle settings for a given minio_s3_bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,minio}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || (has(self.initProvider) && has(self.initProvider.rule))",message="spec.forProvider.rule is a required parameter"
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