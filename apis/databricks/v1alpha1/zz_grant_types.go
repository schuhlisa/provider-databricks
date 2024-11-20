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

type GrantInitParameters struct {
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	ExternalLocation *string `json:"externalLocation,omitempty" tf:"external_location,omitempty"`

	ForeignConnection *string `json:"foreignConnection,omitempty" tf:"foreign_connection,omitempty"`

	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	Metastore *string `json:"metastore,omitempty" tf:"metastore,omitempty"`

	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// User name, group name or service principal application ID.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// One or more privileges that are specific to a securable type.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	Recipient *string `json:"recipient,omitempty" tf:"recipient,omitempty"`

	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	Share *string `json:"share,omitempty" tf:"share,omitempty"`

	StorageCredential *string `json:"storageCredential,omitempty" tf:"storage_credential,omitempty"`

	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type GrantObservation struct {
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	ExternalLocation *string `json:"externalLocation,omitempty" tf:"external_location,omitempty"`

	ForeignConnection *string `json:"foreignConnection,omitempty" tf:"foreign_connection,omitempty"`

	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Metastore *string `json:"metastore,omitempty" tf:"metastore,omitempty"`

	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// User name, group name or service principal application ID.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// One or more privileges that are specific to a securable type.
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	Recipient *string `json:"recipient,omitempty" tf:"recipient,omitempty"`

	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	Share *string `json:"share,omitempty" tf:"share,omitempty"`

	StorageCredential *string `json:"storageCredential,omitempty" tf:"storage_credential,omitempty"`

	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type GrantParameters struct {

	// +kubebuilder:validation:Optional
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalLocation *string `json:"externalLocation,omitempty" tf:"external_location,omitempty"`

	// +kubebuilder:validation:Optional
	ForeignConnection *string `json:"foreignConnection,omitempty" tf:"foreign_connection,omitempty"`

	// +kubebuilder:validation:Optional
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// +kubebuilder:validation:Optional
	Metastore *string `json:"metastore,omitempty" tf:"metastore,omitempty"`

	// +kubebuilder:validation:Optional
	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	// +kubebuilder:validation:Optional
	Pipeline *string `json:"pipeline,omitempty" tf:"pipeline,omitempty"`

	// User name, group name or service principal application ID.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// One or more privileges that are specific to a securable type.
	// +kubebuilder:validation:Optional
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// +kubebuilder:validation:Optional
	Recipient *string `json:"recipient,omitempty" tf:"recipient,omitempty"`

	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// +kubebuilder:validation:Optional
	Share *string `json:"share,omitempty" tf:"share,omitempty"`

	// +kubebuilder:validation:Optional
	StorageCredential *string `json:"storageCredential,omitempty" tf:"storage_credential,omitempty"`

	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// +kubebuilder:validation:Optional
	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

// GrantSpec defines the desired state of Grant
type GrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrantParameters `json:"forProvider"`
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
	InitProvider GrantInitParameters `json:"initProvider,omitempty"`
}

// GrantStatus defines the observed state of Grant.
type GrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Grant is the Schema for the Grants API. ""subcategory: "Unity Catalog"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Grant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principal) || (has(self.initProvider) && has(self.initProvider.principal))",message="spec.forProvider.principal is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privileges) || (has(self.initProvider) && has(self.initProvider.privileges))",message="spec.forProvider.privileges is a required parameter"
	Spec   GrantSpec   `json:"spec"`
	Status GrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrantList contains a list of Grants
type GrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grant `json:"items"`
}

// Repository type metadata.
var (
	Grant_Kind             = "Grant"
	Grant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Grant_Kind}.String()
	Grant_KindAPIVersion   = Grant_Kind + "." + CRDGroupVersion.String()
	Grant_GroupVersionKind = CRDGroupVersion.WithKind(Grant_Kind)
)

func init() {
	SchemeBuilder.Register(&Grant{}, &GrantList{})
}
