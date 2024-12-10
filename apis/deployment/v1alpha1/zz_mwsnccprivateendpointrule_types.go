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

type MwsNccPrivateEndpointRuleInitParameters struct {

	// The current status of this private endpoint. The private endpoint rules are effective only if the connection state is ESTABLISHED. Remember that you must approve new endpoints on your resources in the Azure portal before they take effect.
	// The possible values are:
	ConnectionState *string `json:"connectionState,omitempty" tf:"connection_state,omitempty"`

	// Time in epoch milliseconds when this object was created.
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Whether this private endpoint is deactivated.
	Deactivated *bool `json:"deactivated,omitempty" tf:"deactivated,omitempty"`

	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt *float64 `json:"deactivatedAt,omitempty" tf:"deactivated_at,omitempty"`

	// The name of the Azure private endpoint resource, e.g. "databricks-088781b3-77fa-4132-b429-1af0d91bc593-pe-3cb31234"
	EndpointName *string `json:"endpointName,omitempty" tf:"endpoint_name,omitempty"`

	// The sub-resource type (group ID) of the target resource. Must be one of supported resource types (i.e., blob, dfs, sqlServer , etc. Consult the Azure documentation for full list of supported resources). Note that to connect to workspace root storage (root DBFS), you need two endpoints, one for blob and one for dfs. Change forces creation of a new resource.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Canonical unique identifier of Network Connectivity Config in Databricks Account. Change forces creation of a new resource.
	NetworkConnectivityConfigID *string `json:"networkConnectivityConfigId,omitempty" tf:"network_connectivity_config_id,omitempty"`

	// The Azure resource ID of the target resource. Change forces creation of a new resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// the ID of a private endpoint rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Time in epoch milliseconds when this object was updated.
	UpdatedTime *float64 `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`
}

type MwsNccPrivateEndpointRuleObservation struct {

	// The current status of this private endpoint. The private endpoint rules are effective only if the connection state is ESTABLISHED. Remember that you must approve new endpoints on your resources in the Azure portal before they take effect.
	// The possible values are:
	ConnectionState *string `json:"connectionState,omitempty" tf:"connection_state,omitempty"`

	// Time in epoch milliseconds when this object was created.
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Whether this private endpoint is deactivated.
	Deactivated *bool `json:"deactivated,omitempty" tf:"deactivated,omitempty"`

	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt *float64 `json:"deactivatedAt,omitempty" tf:"deactivated_at,omitempty"`

	// The name of the Azure private endpoint resource, e.g. "databricks-088781b3-77fa-4132-b429-1af0d91bc593-pe-3cb31234"
	EndpointName *string `json:"endpointName,omitempty" tf:"endpoint_name,omitempty"`

	// The sub-resource type (group ID) of the target resource. Must be one of supported resource types (i.e., blob, dfs, sqlServer , etc. Consult the Azure documentation for full list of supported resources). Note that to connect to workspace root storage (root DBFS), you need two endpoints, one for blob and one for dfs. Change forces creation of a new resource.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Canonical unique identifier of Network Connectivity Config in Databricks Account. Change forces creation of a new resource.
	NetworkConnectivityConfigID *string `json:"networkConnectivityConfigId,omitempty" tf:"network_connectivity_config_id,omitempty"`

	// The Azure resource ID of the target resource. Change forces creation of a new resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// the ID of a private endpoint rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Time in epoch milliseconds when this object was updated.
	UpdatedTime *float64 `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`
}

type MwsNccPrivateEndpointRuleParameters struct {

	// The current status of this private endpoint. The private endpoint rules are effective only if the connection state is ESTABLISHED. Remember that you must approve new endpoints on your resources in the Azure portal before they take effect.
	// The possible values are:
	// +kubebuilder:validation:Optional
	ConnectionState *string `json:"connectionState,omitempty" tf:"connection_state,omitempty"`

	// Time in epoch milliseconds when this object was created.
	// +kubebuilder:validation:Optional
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// Whether this private endpoint is deactivated.
	// +kubebuilder:validation:Optional
	Deactivated *bool `json:"deactivated,omitempty" tf:"deactivated,omitempty"`

	// Time in epoch milliseconds when this object was deactivated.
	// +kubebuilder:validation:Optional
	DeactivatedAt *float64 `json:"deactivatedAt,omitempty" tf:"deactivated_at,omitempty"`

	// The name of the Azure private endpoint resource, e.g. "databricks-088781b3-77fa-4132-b429-1af0d91bc593-pe-3cb31234"
	// +kubebuilder:validation:Optional
	EndpointName *string `json:"endpointName,omitempty" tf:"endpoint_name,omitempty"`

	// The sub-resource type (group ID) of the target resource. Must be one of supported resource types (i.e., blob, dfs, sqlServer , etc. Consult the Azure documentation for full list of supported resources). Note that to connect to workspace root storage (root DBFS), you need two endpoints, one for blob and one for dfs. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Canonical unique identifier of Network Connectivity Config in Databricks Account. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	NetworkConnectivityConfigID *string `json:"networkConnectivityConfigId,omitempty" tf:"network_connectivity_config_id,omitempty"`

	// The Azure resource ID of the target resource. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// the ID of a private endpoint rule.
	// +kubebuilder:validation:Optional
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Time in epoch milliseconds when this object was updated.
	// +kubebuilder:validation:Optional
	UpdatedTime *float64 `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`
}

// MwsNccPrivateEndpointRuleSpec defines the desired state of MwsNccPrivateEndpointRule
type MwsNccPrivateEndpointRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MwsNccPrivateEndpointRuleParameters `json:"forProvider"`
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
	InitProvider MwsNccPrivateEndpointRuleInitParameters `json:"initProvider,omitempty"`
}

// MwsNccPrivateEndpointRuleStatus defines the observed state of MwsNccPrivateEndpointRule.
type MwsNccPrivateEndpointRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MwsNccPrivateEndpointRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MwsNccPrivateEndpointRule is the Schema for the MwsNccPrivateEndpointRules API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type MwsNccPrivateEndpointRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupId) || (has(self.initProvider) && has(self.initProvider.groupId))",message="spec.forProvider.groupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkConnectivityConfigId) || (has(self.initProvider) && has(self.initProvider.networkConnectivityConfigId))",message="spec.forProvider.networkConnectivityConfigId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceId) || (has(self.initProvider) && has(self.initProvider.resourceId))",message="spec.forProvider.resourceId is a required parameter"
	Spec   MwsNccPrivateEndpointRuleSpec   `json:"spec"`
	Status MwsNccPrivateEndpointRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MwsNccPrivateEndpointRuleList contains a list of MwsNccPrivateEndpointRules
type MwsNccPrivateEndpointRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MwsNccPrivateEndpointRule `json:"items"`
}

// Repository type metadata.
var (
	MwsNccPrivateEndpointRule_Kind             = "MwsNccPrivateEndpointRule"
	MwsNccPrivateEndpointRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MwsNccPrivateEndpointRule_Kind}.String()
	MwsNccPrivateEndpointRule_KindAPIVersion   = MwsNccPrivateEndpointRule_Kind + "." + CRDGroupVersion.String()
	MwsNccPrivateEndpointRule_GroupVersionKind = CRDGroupVersion.WithKind(MwsNccPrivateEndpointRule_Kind)
)

func init() {
	SchemeBuilder.Register(&MwsNccPrivateEndpointRule{}, &MwsNccPrivateEndpointRuleList{})
}
