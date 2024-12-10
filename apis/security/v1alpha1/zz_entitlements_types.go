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

type EntitlementsInitParameters struct {

	// Allow the principal to have cluster create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and cluster_id argument. Everyone without allow_cluster_create argument set, but with permission to use Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
	AllowClusterCreate *bool `json:"allowClusterCreate,omitempty" tf:"allow_cluster_create,omitempty"`

	// Allow the principal to have instance pool create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and instance_pool_id argument.
	AllowInstancePoolCreate *bool `json:"allowInstancePoolCreate,omitempty" tf:"allow_instance_pool_create,omitempty"`

	// This is a field to allow the principal to have access to Databricks SQL feature in User Interface and through databricks_sql_endpoint.
	DatabricksSQLAccess *bool `json:"databricksSqlAccess,omitempty" tf:"databricks_sql_access,omitempty"`

	// Canonical unique identifier for the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Canonical unique identifier for the service principal.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Canonical unique identifier for the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// This is a field to allow the principal to have access to Databricks Workspace.
	WorkspaceAccess *bool `json:"workspaceAccess,omitempty" tf:"workspace_access,omitempty"`
}

type EntitlementsObservation struct {

	// Allow the principal to have cluster create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and cluster_id argument. Everyone without allow_cluster_create argument set, but with permission to use Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
	AllowClusterCreate *bool `json:"allowClusterCreate,omitempty" tf:"allow_cluster_create,omitempty"`

	// Allow the principal to have instance pool create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and instance_pool_id argument.
	AllowInstancePoolCreate *bool `json:"allowInstancePoolCreate,omitempty" tf:"allow_instance_pool_create,omitempty"`

	// This is a field to allow the principal to have access to Databricks SQL feature in User Interface and through databricks_sql_endpoint.
	DatabricksSQLAccess *bool `json:"databricksSqlAccess,omitempty" tf:"databricks_sql_access,omitempty"`

	// Canonical unique identifier for the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Canonical unique identifier for the service principal.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Canonical unique identifier for the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// This is a field to allow the principal to have access to Databricks Workspace.
	WorkspaceAccess *bool `json:"workspaceAccess,omitempty" tf:"workspace_access,omitempty"`
}

type EntitlementsParameters struct {

	// Allow the principal to have cluster create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and cluster_id argument. Everyone without allow_cluster_create argument set, but with permission to use Cluster Policy would be able to create clusters, but within boundaries of that specific policy.
	// +kubebuilder:validation:Optional
	AllowClusterCreate *bool `json:"allowClusterCreate,omitempty" tf:"allow_cluster_create,omitempty"`

	// Allow the principal to have instance pool create privileges. Defaults to false. More fine grained permissions could be assigned with databricks_permissions and instance_pool_id argument.
	// +kubebuilder:validation:Optional
	AllowInstancePoolCreate *bool `json:"allowInstancePoolCreate,omitempty" tf:"allow_instance_pool_create,omitempty"`

	// This is a field to allow the principal to have access to Databricks SQL feature in User Interface and through databricks_sql_endpoint.
	// +kubebuilder:validation:Optional
	DatabricksSQLAccess *bool `json:"databricksSqlAccess,omitempty" tf:"databricks_sql_access,omitempty"`

	// Canonical unique identifier for the group.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Canonical unique identifier for the service principal.
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Canonical unique identifier for the user.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// This is a field to allow the principal to have access to Databricks Workspace.
	// +kubebuilder:validation:Optional
	WorkspaceAccess *bool `json:"workspaceAccess,omitempty" tf:"workspace_access,omitempty"`
}

// EntitlementsSpec defines the desired state of Entitlements
type EntitlementsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntitlementsParameters `json:"forProvider"`
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
	InitProvider EntitlementsInitParameters `json:"initProvider,omitempty"`
}

// EntitlementsStatus defines the observed state of Entitlements.
type EntitlementsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntitlementsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Entitlements is the Schema for the Entitlementss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Entitlements struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EntitlementsSpec   `json:"spec"`
	Status            EntitlementsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntitlementsList contains a list of Entitlementss
type EntitlementsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Entitlements `json:"items"`
}

// Repository type metadata.
var (
	Entitlements_Kind             = "Entitlements"
	Entitlements_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Entitlements_Kind}.String()
	Entitlements_KindAPIVersion   = Entitlements_Kind + "." + CRDGroupVersion.String()
	Entitlements_GroupVersionKind = CRDGroupVersion.WithKind(Entitlements_Kind)
)

func init() {
	SchemeBuilder.Register(&Entitlements{}, &EntitlementsList{})
}
