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

type MwsNccBindingInitParameters struct {

	// Canonical unique identifier of Network Connectivity Config in Databricks Account.
	NetworkConnectivityConfigID *string `json:"networkConnectivityConfigId,omitempty" tf:"network_connectivity_config_id,omitempty"`

	// Identifier of the workspace to attach the NCC to. Change forces creation of a new resource.
	WorkspaceID *float64 `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type MwsNccBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Canonical unique identifier of Network Connectivity Config in Databricks Account.
	NetworkConnectivityConfigID *string `json:"networkConnectivityConfigId,omitempty" tf:"network_connectivity_config_id,omitempty"`

	// Identifier of the workspace to attach the NCC to. Change forces creation of a new resource.
	WorkspaceID *float64 `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type MwsNccBindingParameters struct {

	// Canonical unique identifier of Network Connectivity Config in Databricks Account.
	// +kubebuilder:validation:Optional
	NetworkConnectivityConfigID *string `json:"networkConnectivityConfigId,omitempty" tf:"network_connectivity_config_id,omitempty"`

	// Identifier of the workspace to attach the NCC to. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	WorkspaceID *float64 `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

// MwsNccBindingSpec defines the desired state of MwsNccBinding
type MwsNccBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MwsNccBindingParameters `json:"forProvider"`
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
	InitProvider MwsNccBindingInitParameters `json:"initProvider,omitempty"`
}

// MwsNccBindingStatus defines the observed state of MwsNccBinding.
type MwsNccBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MwsNccBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MwsNccBinding is the Schema for the MwsNccBindings API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type MwsNccBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkConnectivityConfigId) || (has(self.initProvider) && has(self.initProvider.networkConnectivityConfigId))",message="spec.forProvider.networkConnectivityConfigId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.workspaceId) || (has(self.initProvider) && has(self.initProvider.workspaceId))",message="spec.forProvider.workspaceId is a required parameter"
	Spec   MwsNccBindingSpec   `json:"spec"`
	Status MwsNccBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MwsNccBindingList contains a list of MwsNccBindings
type MwsNccBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MwsNccBinding `json:"items"`
}

// Repository type metadata.
var (
	MwsNccBinding_Kind             = "MwsNccBinding"
	MwsNccBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MwsNccBinding_Kind}.String()
	MwsNccBinding_KindAPIVersion   = MwsNccBinding_Kind + "." + CRDGroupVersion.String()
	MwsNccBinding_GroupVersionKind = CRDGroupVersion.WithKind(MwsNccBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&MwsNccBinding{}, &MwsNccBindingList{})
}
