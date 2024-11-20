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

type UserRoleInitParameters struct {

	// Either a role name or the ARN/ID of the instance profile resource.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// This is the id of the user resource.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserRoleObservation struct {

	// The id in the format <user_id>|<role>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Either a role name or the ARN/ID of the instance profile resource.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// This is the id of the user resource.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserRoleParameters struct {

	// Either a role name or the ARN/ID of the instance profile resource.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// This is the id of the user resource.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// UserRoleSpec defines the desired state of UserRole
type UserRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserRoleParameters `json:"forProvider"`
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
	InitProvider UserRoleInitParameters `json:"initProvider,omitempty"`
}

// UserRoleStatus defines the observed state of UserRole.
type UserRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserRole is the Schema for the UserRoles API. ""subcategory: "Security"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type UserRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userId) || (has(self.initProvider) && has(self.initProvider.userId))",message="spec.forProvider.userId is a required parameter"
	Spec   UserRoleSpec   `json:"spec"`
	Status UserRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserRoleList contains a list of UserRoles
type UserRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserRole `json:"items"`
}

// Repository type metadata.
var (
	UserRole_Kind             = "UserRole"
	UserRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserRole_Kind}.String()
	UserRole_KindAPIVersion   = UserRole_Kind + "." + CRDGroupVersion.String()
	UserRole_GroupVersionKind = CRDGroupVersion.WithKind(UserRole_Kind)
)

func init() {
	SchemeBuilder.Register(&UserRole{}, &UserRoleList{})
}
