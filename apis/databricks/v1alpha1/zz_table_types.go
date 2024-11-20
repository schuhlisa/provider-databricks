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

type TableColumnInitParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Nullable *bool `json:"nullable,omitempty" tf:"nullable,omitempty"`

	PartitionIndex *float64 `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	TypeIntervalType *string `json:"typeIntervalType,omitempty" tf:"type_interval_type,omitempty"`

	TypeJSON *string `json:"typeJson,omitempty" tf:"type_json,omitempty"`

	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	TypePrecision *float64 `json:"typePrecision,omitempty" tf:"type_precision,omitempty"`

	TypeScale *float64 `json:"typeScale,omitempty" tf:"type_scale,omitempty"`

	TypeText *string `json:"typeText,omitempty" tf:"type_text,omitempty"`
}

type TableColumnObservation struct {
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Nullable *bool `json:"nullable,omitempty" tf:"nullable,omitempty"`

	PartitionIndex *float64 `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	TypeIntervalType *string `json:"typeIntervalType,omitempty" tf:"type_interval_type,omitempty"`

	TypeJSON *string `json:"typeJson,omitempty" tf:"type_json,omitempty"`

	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	TypePrecision *float64 `json:"typePrecision,omitempty" tf:"type_precision,omitempty"`

	TypeScale *float64 `json:"typeScale,omitempty" tf:"type_scale,omitempty"`

	TypeText *string `json:"typeText,omitempty" tf:"type_text,omitempty"`
}

type TableColumnParameters struct {

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Nullable *bool `json:"nullable,omitempty" tf:"nullable,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionIndex *float64 `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	// +kubebuilder:validation:Optional
	Position *float64 `json:"position" tf:"position,omitempty"`

	// +kubebuilder:validation:Optional
	TypeIntervalType *string `json:"typeIntervalType,omitempty" tf:"type_interval_type,omitempty"`

	// +kubebuilder:validation:Optional
	TypeJSON *string `json:"typeJson,omitempty" tf:"type_json,omitempty"`

	// +kubebuilder:validation:Optional
	TypeName *string `json:"typeName" tf:"type_name,omitempty"`

	// +kubebuilder:validation:Optional
	TypePrecision *float64 `json:"typePrecision,omitempty" tf:"type_precision,omitempty"`

	// +kubebuilder:validation:Optional
	TypeScale *float64 `json:"typeScale,omitempty" tf:"type_scale,omitempty"`

	// +kubebuilder:validation:Optional
	TypeText *string `json:"typeText" tf:"type_text,omitempty"`
}

type TableInitParameters_2 struct {
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	Column []TableColumnInitParameters `json:"column,omitempty" tf:"column,omitempty"`

	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	DataSourceFormat *string `json:"dataSourceFormat,omitempty" tf:"data_source_format,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	StorageCredentialName *string `json:"storageCredentialName,omitempty" tf:"storage_credential_name,omitempty"`

	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	ViewDefinition *string `json:"viewDefinition,omitempty" tf:"view_definition,omitempty"`
}

type TableObservation_2 struct {
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	Column []TableColumnObservation `json:"column,omitempty" tf:"column,omitempty"`

	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	DataSourceFormat *string `json:"dataSourceFormat,omitempty" tf:"data_source_format,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	StorageCredentialName *string `json:"storageCredentialName,omitempty" tf:"storage_credential_name,omitempty"`

	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	ViewDefinition *string `json:"viewDefinition,omitempty" tf:"view_definition,omitempty"`
}

type TableParameters_2 struct {

	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// +kubebuilder:validation:Optional
	Column []TableColumnParameters `json:"column,omitempty" tf:"column,omitempty"`

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	DataSourceFormat *string `json:"dataSourceFormat,omitempty" tf:"data_source_format,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageCredentialName *string `json:"storageCredentialName,omitempty" tf:"storage_credential_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageLocation *string `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// +kubebuilder:validation:Optional
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// +kubebuilder:validation:Optional
	ViewDefinition *string `json:"viewDefinition,omitempty" tf:"view_definition,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters_2 `json:"forProvider"`
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
	InitProvider TableInitParameters_2 `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Table is the Schema for the Tables API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catalogName) || (has(self.initProvider) && has(self.initProvider.catalogName))",message="spec.forProvider.catalogName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.column) || (has(self.initProvider) && has(self.initProvider.column))",message="spec.forProvider.column is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataSourceFormat) || (has(self.initProvider) && has(self.initProvider.dataSourceFormat))",message="spec.forProvider.dataSourceFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemaName) || (has(self.initProvider) && has(self.initProvider.schemaName))",message="spec.forProvider.schemaName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tableType) || (has(self.initProvider) && has(self.initProvider.tableType))",message="spec.forProvider.tableType is a required parameter"
	Spec   TableSpec   `json:"spec"`
	Status TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
