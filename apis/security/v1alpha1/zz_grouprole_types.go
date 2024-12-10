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

type GroupRoleInitParameters struct {

	// This is the id of the group resource.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Either a role name or the ARN/ID of the instance profile resource.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type GroupRoleObservation struct {

	// This is the id of the group resource.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The id for the databricks_group_role object which is in the format <group_id>|<role>.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Either a role name or the ARN/ID of the instance profile resource.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type GroupRoleParameters struct {

	// This is the id of the group resource.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Either a role name or the ARN/ID of the instance profile resource.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// GroupRoleSpec defines the desired state of GroupRole
type GroupRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupRoleParameters `json:"forProvider"`
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
	InitProvider GroupRoleInitParameters `json:"initProvider,omitempty"`
}

// GroupRoleStatus defines the observed state of GroupRole.
type GroupRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GroupRole is the Schema for the GroupRoles API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type GroupRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupId) || (has(self.initProvider) && has(self.initProvider.groupId))",message="spec.forProvider.groupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   GroupRoleSpec   `json:"spec"`
	Status GroupRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupRoleList contains a list of GroupRoles
type GroupRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupRole `json:"items"`
}

// Repository type metadata.
var (
	GroupRole_Kind             = "GroupRole"
	GroupRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupRole_Kind}.String()
	GroupRole_KindAPIVersion   = GroupRole_Kind + "." + CRDGroupVersion.String()
	GroupRole_GroupVersionKind = CRDGroupVersion.WithKind(GroupRole_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupRole{}, &GroupRoleList{})
}
