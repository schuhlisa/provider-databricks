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

type MetastoreInitParameters struct {
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	DefaultDataAccessConfigID *string `json:"defaultDataAccessConfigId,omitempty" tf:"default_data_access_config_id,omitempty"`

	// The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing. Once this is set it cannot be removed and can only be modified to another valid value. To delete this value please taint and recreate the resource.
	DeltaSharingOrganizationName *string `json:"deltaSharingOrganizationName,omitempty" tf:"delta_sharing_organization_name,omitempty"`

	// Required along with delta_sharing_scope. Used to set expiration duration in seconds on recipient data access tokens. Set to 0 for unlimited duration.
	DeltaSharingRecipientTokenLifetimeInSeconds *float64 `json:"deltaSharingRecipientTokenLifetimeInSeconds,omitempty" tf:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`

	// Required along with delta_sharing_recipient_token_lifetime_in_seconds. Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL.  INTERNAL only allows sharing within the same account, and INTERNAL_AND_EXTERNAL allows cross account sharing and token based sharing.
	DeltaSharingScope *string `json:"deltaSharingScope,omitempty" tf:"delta_sharing_scope,omitempty"`

	// Destroy metastore regardless of its contents.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	GlobalMetastoreID *string `json:"globalMetastoreId,omitempty" tf:"global_metastore_id,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of metastore.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Username/groupname/sp application_id of the metastore owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// (Mandatory for account-level) The region of the metastore
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Path on cloud storage account, where managed databricks_table are stored. Change forces creation of a new resource. If no storage_root is defined for the metastore, each catalog must have a storage_root defined.
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	StorageRootCredentialID *string `json:"storageRootCredentialId,omitempty" tf:"storage_root_credential_id,omitempty"`

	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

type MetastoreObservation struct {
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	DefaultDataAccessConfigID *string `json:"defaultDataAccessConfigId,omitempty" tf:"default_data_access_config_id,omitempty"`

	// The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing. Once this is set it cannot be removed and can only be modified to another valid value. To delete this value please taint and recreate the resource.
	DeltaSharingOrganizationName *string `json:"deltaSharingOrganizationName,omitempty" tf:"delta_sharing_organization_name,omitempty"`

	// Required along with delta_sharing_scope. Used to set expiration duration in seconds on recipient data access tokens. Set to 0 for unlimited duration.
	DeltaSharingRecipientTokenLifetimeInSeconds *float64 `json:"deltaSharingRecipientTokenLifetimeInSeconds,omitempty" tf:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`

	// Required along with delta_sharing_recipient_token_lifetime_in_seconds. Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL.  INTERNAL only allows sharing within the same account, and INTERNAL_AND_EXTERNAL allows cross account sharing and token based sharing.
	DeltaSharingScope *string `json:"deltaSharingScope,omitempty" tf:"delta_sharing_scope,omitempty"`

	// Destroy metastore regardless of its contents.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	GlobalMetastoreID *string `json:"globalMetastoreId,omitempty" tf:"global_metastore_id,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of metastore.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Username/groupname/sp application_id of the metastore owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// (Mandatory for account-level) The region of the metastore
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Path on cloud storage account, where managed databricks_table are stored. Change forces creation of a new resource. If no storage_root is defined for the metastore, each catalog must have a storage_root defined.
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	StorageRootCredentialID *string `json:"storageRootCredentialId,omitempty" tf:"storage_root_credential_id,omitempty"`

	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

type MetastoreParameters struct {

	// +kubebuilder:validation:Optional
	Cloud *string `json:"cloud,omitempty" tf:"cloud,omitempty"`

	// +kubebuilder:validation:Optional
	CreatedAt *float64 `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	// +kubebuilder:validation:Optional
	DefaultDataAccessConfigID *string `json:"defaultDataAccessConfigId,omitempty" tf:"default_data_access_config_id,omitempty"`

	// The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing. Once this is set it cannot be removed and can only be modified to another valid value. To delete this value please taint and recreate the resource.
	// +kubebuilder:validation:Optional
	DeltaSharingOrganizationName *string `json:"deltaSharingOrganizationName,omitempty" tf:"delta_sharing_organization_name,omitempty"`

	// Required along with delta_sharing_scope. Used to set expiration duration in seconds on recipient data access tokens. Set to 0 for unlimited duration.
	// +kubebuilder:validation:Optional
	DeltaSharingRecipientTokenLifetimeInSeconds *float64 `json:"deltaSharingRecipientTokenLifetimeInSeconds,omitempty" tf:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`

	// Required along with delta_sharing_recipient_token_lifetime_in_seconds. Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL.  INTERNAL only allows sharing within the same account, and INTERNAL_AND_EXTERNAL allows cross account sharing and token based sharing.
	// +kubebuilder:validation:Optional
	DeltaSharingScope *string `json:"deltaSharingScope,omitempty" tf:"delta_sharing_scope,omitempty"`

	// Destroy metastore regardless of its contents.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	// +kubebuilder:validation:Optional
	GlobalMetastoreID *string `json:"globalMetastoreId,omitempty" tf:"global_metastore_id,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	// +kubebuilder:validation:Optional
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of metastore.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Username/groupname/sp application_id of the metastore owner.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// (Mandatory for account-level) The region of the metastore
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Path on cloud storage account, where managed databricks_table are stored. Change forces creation of a new resource. If no storage_root is defined for the metastore, each catalog must have a storage_root defined.
	// +kubebuilder:validation:Optional
	StorageRoot *string `json:"storageRoot,omitempty" tf:"storage_root,omitempty"`

	// system-generated ID of this Unity Catalog Metastore.
	// +kubebuilder:validation:Optional
	StorageRootCredentialID *string `json:"storageRootCredentialId,omitempty" tf:"storage_root_credential_id,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatedAt *float64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

// MetastoreSpec defines the desired state of Metastore
type MetastoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetastoreParameters `json:"forProvider"`
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
	InitProvider MetastoreInitParameters `json:"initProvider,omitempty"`
}

// MetastoreStatus defines the observed state of Metastore.
type MetastoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetastoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Metastore is the Schema for the Metastores API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Metastore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MetastoreSpec   `json:"spec"`
	Status MetastoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetastoreList contains a list of Metastores
type MetastoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Metastore `json:"items"`
}

// Repository type metadata.
var (
	Metastore_Kind             = "Metastore"
	Metastore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Metastore_Kind}.String()
	Metastore_KindAPIVersion   = Metastore_Kind + "." + CRDGroupVersion.String()
	Metastore_GroupVersionKind = CRDGroupVersion.WithKind(Metastore_Kind)
)

func init() {
	SchemeBuilder.Register(&Metastore{}, &MetastoreList{})
}
