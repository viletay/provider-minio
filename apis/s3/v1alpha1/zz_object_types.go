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

type ObjectInitParameters struct {

	// (String)
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String)
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// (String)
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// (String)
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// (String)
	ObjectName *string `json:"objectName,omitempty" tf:"object_name,omitempty"`

	// (String)
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String)
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type ObjectObservation struct {

	// (String)
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String)
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String)
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// (String)
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// (String)
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	ObjectName *string `json:"objectName,omitempty" tf:"object_name,omitempty"`

	// (String)
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String)
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type ObjectParameters struct {

	// (String)
	// +crossplane:generate:reference:type=Bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ObjectName *string `json:"objectName,omitempty" tf:"object_name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
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
	InitProvider ObjectInitParameters `json:"initProvider,omitempty"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Object is the Schema for the Objects API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,minio}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectName) || (has(self.initProvider) && has(self.initProvider.objectName))",message="spec.forProvider.objectName is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}