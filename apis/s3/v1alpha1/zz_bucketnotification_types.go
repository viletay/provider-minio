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

type BucketNotificationInitParameters struct {
	Queue []QueueInitParameters `json:"queue,omitempty" tf:"queue,omitempty"`
}

type BucketNotificationObservation struct {
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Queue []QueueObservation `json:"queue,omitempty" tf:"queue,omitempty"`
}

type BucketNotificationParameters struct {

	// +crossplane:generate:reference:type=Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Queue []QueueParameters `json:"queue,omitempty" tf:"queue,omitempty"`
}

type QueueInitParameters struct {
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	FilterPrefix *string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`

	FilterSuffix *string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	QueueArn *string `json:"queueArn,omitempty" tf:"queue_arn,omitempty"`
}

type QueueObservation struct {
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	FilterPrefix *string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`

	FilterSuffix *string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	QueueArn *string `json:"queueArn,omitempty" tf:"queue_arn,omitempty"`
}

type QueueParameters struct {

	// +kubebuilder:validation:Optional
	Events []*string `json:"events" tf:"events,omitempty"`

	// +kubebuilder:validation:Optional
	FilterPrefix *string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	FilterSuffix *string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	QueueArn *string `json:"queueArn" tf:"queue_arn,omitempty"`
}

// BucketNotificationSpec defines the desired state of BucketNotification
type BucketNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketNotificationParameters `json:"forProvider"`
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
	InitProvider BucketNotificationInitParameters `json:"initProvider,omitempty"`
}

// BucketNotificationStatus defines the observed state of BucketNotification.
type BucketNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketNotification is the Schema for the BucketNotifications API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,minio}
type BucketNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketNotificationSpec   `json:"spec"`
	Status            BucketNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketNotificationList contains a list of BucketNotifications
type BucketNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketNotification `json:"items"`
}

// Repository type metadata.
var (
	BucketNotification_Kind             = "BucketNotification"
	BucketNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketNotification_Kind}.String()
	BucketNotification_KindAPIVersion   = BucketNotification_Kind + "." + CRDGroupVersion.String()
	BucketNotification_GroupVersionKind = CRDGroupVersion.WithKind(BucketNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketNotification{}, &BucketNotificationList{})
}
