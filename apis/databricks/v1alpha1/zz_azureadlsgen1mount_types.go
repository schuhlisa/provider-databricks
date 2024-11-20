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

type AzureAdlsGen1MountInitParameters struct {
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	ClientSecretKey *string `json:"clientSecretKey,omitempty" tf:"client_secret_key,omitempty"`

	ClientSecretScope *string `json:"clientSecretScope,omitempty" tf:"client_secret_scope,omitempty"`

	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	Directory *string `json:"directory,omitempty" tf:"directory,omitempty"`

	MountName *string `json:"mountName,omitempty" tf:"mount_name,omitempty"`

	SparkConfPrefix *string `json:"sparkConfPrefix,omitempty" tf:"spark_conf_prefix,omitempty"`

	StorageResourceName *string `json:"storageResourceName,omitempty" tf:"storage_resource_name,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureAdlsGen1MountObservation struct {
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	ClientSecretKey *string `json:"clientSecretKey,omitempty" tf:"client_secret_key,omitempty"`

	ClientSecretScope *string `json:"clientSecretScope,omitempty" tf:"client_secret_scope,omitempty"`

	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	Directory *string `json:"directory,omitempty" tf:"directory,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MountName *string `json:"mountName,omitempty" tf:"mount_name,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	SparkConfPrefix *string `json:"sparkConfPrefix,omitempty" tf:"spark_conf_prefix,omitempty"`

	StorageResourceName *string `json:"storageResourceName,omitempty" tf:"storage_resource_name,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureAdlsGen1MountParameters struct {

	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClientSecretKey *string `json:"clientSecretKey,omitempty" tf:"client_secret_key,omitempty"`

	// +kubebuilder:validation:Optional
	ClientSecretScope *string `json:"clientSecretScope,omitempty" tf:"client_secret_scope,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	Directory *string `json:"directory,omitempty" tf:"directory,omitempty"`

	// +kubebuilder:validation:Optional
	MountName *string `json:"mountName,omitempty" tf:"mount_name,omitempty"`

	// +kubebuilder:validation:Optional
	SparkConfPrefix *string `json:"sparkConfPrefix,omitempty" tf:"spark_conf_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	StorageResourceName *string `json:"storageResourceName,omitempty" tf:"storage_resource_name,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// AzureAdlsGen1MountSpec defines the desired state of AzureAdlsGen1Mount
type AzureAdlsGen1MountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureAdlsGen1MountParameters `json:"forProvider"`
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
	InitProvider AzureAdlsGen1MountInitParameters `json:"initProvider,omitempty"`
}

// AzureAdlsGen1MountStatus defines the observed state of AzureAdlsGen1Mount.
type AzureAdlsGen1MountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureAdlsGen1MountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AzureAdlsGen1Mount is the Schema for the AzureAdlsGen1Mounts API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type AzureAdlsGen1Mount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretKey) || (has(self.initProvider) && has(self.initProvider.clientSecretKey))",message="spec.forProvider.clientSecretKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretScope) || (has(self.initProvider) && has(self.initProvider.clientSecretScope))",message="spec.forProvider.clientSecretScope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mountName) || (has(self.initProvider) && has(self.initProvider.mountName))",message="spec.forProvider.mountName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageResourceName) || (has(self.initProvider) && has(self.initProvider.storageResourceName))",message="spec.forProvider.storageResourceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   AzureAdlsGen1MountSpec   `json:"spec"`
	Status AzureAdlsGen1MountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureAdlsGen1MountList contains a list of AzureAdlsGen1Mounts
type AzureAdlsGen1MountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureAdlsGen1Mount `json:"items"`
}

// Repository type metadata.
var (
	AzureAdlsGen1Mount_Kind             = "AzureAdlsGen1Mount"
	AzureAdlsGen1Mount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AzureAdlsGen1Mount_Kind}.String()
	AzureAdlsGen1Mount_KindAPIVersion   = AzureAdlsGen1Mount_Kind + "." + CRDGroupVersion.String()
	AzureAdlsGen1Mount_GroupVersionKind = CRDGroupVersion.WithKind(AzureAdlsGen1Mount_Kind)
)

func init() {
	SchemeBuilder.Register(&AzureAdlsGen1Mount{}, &AzureAdlsGen1MountList{})
}
