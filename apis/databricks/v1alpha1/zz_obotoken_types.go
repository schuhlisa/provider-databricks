// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OboTokenInitParameters struct {

	// Application ID of databricks_service_principal to create a PAT token for.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String, Optional) Comment that describes the purpose of the token.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (Integer, Optional) The number of seconds before the token expires. Token resource is re-created when it expires. If no lifetime is specified, the token remains valid indefinitely.
	LifetimeSeconds *float64 `json:"lifetimeSeconds,omitempty" tf:"lifetime_seconds,omitempty"`
}

type OboTokenObservation struct {

	// Application ID of databricks_service_principal to create a PAT token for.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String, Optional) Comment that describes the purpose of the token.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Canonical unique identifier for the token.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Integer, Optional) The number of seconds before the token expires. Token resource is re-created when it expires. If no lifetime is specified, the token remains valid indefinitely.
	LifetimeSeconds *float64 `json:"lifetimeSeconds,omitempty" tf:"lifetime_seconds,omitempty"`
}

type OboTokenParameters struct {

	// Application ID of databricks_service_principal to create a PAT token for.
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// (String, Optional) Comment that describes the purpose of the token.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (Integer, Optional) The number of seconds before the token expires. Token resource is re-created when it expires. If no lifetime is specified, the token remains valid indefinitely.
	// +kubebuilder:validation:Optional
	LifetimeSeconds *float64 `json:"lifetimeSeconds,omitempty" tf:"lifetime_seconds,omitempty"`
}

// OboTokenSpec defines the desired state of OboToken
type OboTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OboTokenParameters `json:"forProvider"`
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
	InitProvider OboTokenInitParameters `json:"initProvider,omitempty"`
}

// OboTokenStatus defines the observed state of OboToken.
type OboTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OboTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OboToken is the Schema for the OboTokens API. ""subcategory: "Security"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type OboToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationId) || (has(self.initProvider) && has(self.initProvider.applicationId))",message="spec.forProvider.applicationId is a required parameter"
	Spec   OboTokenSpec   `json:"spec"`
	Status OboTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OboTokenList contains a list of OboTokens
type OboTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OboToken `json:"items"`
}

// Repository type metadata.
var (
	OboToken_Kind             = "OboToken"
	OboToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OboToken_Kind}.String()
	OboToken_KindAPIVersion   = OboToken_Kind + "." + CRDGroupVersion.String()
	OboToken_GroupVersionKind = CRDGroupVersion.WithKind(OboToken_Kind)
)

func init() {
	SchemeBuilder.Register(&OboToken{}, &OboTokenList{})
}
