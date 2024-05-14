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

type UserInitParameters struct {

	// (Boolean) Disable user
	// Disable user
	DisableUser *bool `json:"disableUser,omitempty" tf:"disable_user,omitempty"`

	// provider-managed iam access keys
	// delete user even if it has non-provider-managed iam access keys
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Map of String)
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Boolean) Rotate Minio User Secret Key
	// Rotate Minio User Secret Key
	UpdateSecret *bool `json:"updateSecret,omitempty" tf:"update_secret,omitempty"`
}

type UserObservation struct {

	// (Boolean) Disable user
	// Disable user
	DisableUser *bool `json:"disableUser,omitempty" tf:"disable_user,omitempty"`

	// provider-managed iam access keys
	// delete user even if it has non-provider-managed iam access keys
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Map of String)
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Boolean) Rotate Minio User Secret Key
	// Rotate Minio User Secret Key
	UpdateSecret *bool `json:"updateSecret,omitempty" tf:"update_secret,omitempty"`
}

type UserParameters struct {

	// (Boolean) Disable user
	// Disable user
	// +kubebuilder:validation:Optional
	DisableUser *bool `json:"disableUser,omitempty" tf:"disable_user,omitempty"`

	// provider-managed iam access keys
	// delete user even if it has non-provider-managed iam access keys
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// (Map of String)
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Boolean) Rotate Minio User Secret Key
	// Rotate Minio User Secret Key
	// +kubebuilder:validation:Optional
	UpdateSecret *bool `json:"updateSecret,omitempty" tf:"update_secret,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,minio}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
