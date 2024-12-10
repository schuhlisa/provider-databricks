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

type EncryptionDetailsInitParameters struct {

	// The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
	SseEncryptionDetails []SseEncryptionDetailsInitParameters `json:"sseEncryptionDetails,omitempty" tf:"sse_encryption_details,omitempty"`
}

type EncryptionDetailsObservation struct {

	// The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
	SseEncryptionDetails []SseEncryptionDetailsObservation `json:"sseEncryptionDetails,omitempty" tf:"sse_encryption_details,omitempty"`
}

type EncryptionDetailsParameters struct {

	// The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
	// +kubebuilder:validation:Optional
	SseEncryptionDetails []SseEncryptionDetailsParameters `json:"sseEncryptionDetails,omitempty" tf:"sse_encryption_details,omitempty"`
}

type ExternalLocationInitParameters struct {

	// The ARN of the s3 access point to use with the external location (AWS).
	AccessPoint *string `json:"accessPoint,omitempty" tf:"access_point,omitempty"`

	// User-supplied free-form text.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the databricks_storage_credential to use with this external location.
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
	EncryptionDetails []EncryptionDetailsInitParameters `json:"encryptionDetails,omitempty" tf:"encryption_details,omitempty"`

	// Destroy external location regardless of its dependents.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Update external location regardless of its dependents.
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`

	// Whether the external location is accessible from all workspaces or a specific set of workspaces. Can be ISOLATION_MODE_ISOLATED or ISOLATION_MODE_OPEN. Setting the external location to ISOLATION_MODE_ISOLATED will automatically allow access from the current workspace.
	IsolationMode *string `json:"isolationMode,omitempty" tf:"isolation_mode,omitempty"`

	// ID of this external location - same as name.
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of External Location, which must be unique within the databricks_metastore. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Username/groupname/sp application_id of the external location owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Indicates whether the external location is read-only.
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Suppress validation errors if any & force save the external location
	SkipValidation *bool `json:"skipValidation,omitempty" tf:"skip_validation,omitempty"`

	// Path URL in cloud storage, of the form: s3://[bucket-host]/[bucket-dir] (AWS), abfss://[user]@[host]/[path] (Azure), gs://[bucket-host]/[bucket-dir] (GCP).
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ExternalLocationObservation struct {

	// The ARN of the s3 access point to use with the external location (AWS).
	AccessPoint *string `json:"accessPoint,omitempty" tf:"access_point,omitempty"`

	// User-supplied free-form text.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the databricks_storage_credential to use with this external location.
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
	EncryptionDetails []EncryptionDetailsObservation `json:"encryptionDetails,omitempty" tf:"encryption_details,omitempty"`

	// Destroy external location regardless of its dependents.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Update external location regardless of its dependents.
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`

	// ID of this external location - same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the external location is accessible from all workspaces or a specific set of workspaces. Can be ISOLATION_MODE_ISOLATED or ISOLATION_MODE_OPEN. Setting the external location to ISOLATION_MODE_ISOLATED will automatically allow access from the current workspace.
	IsolationMode *string `json:"isolationMode,omitempty" tf:"isolation_mode,omitempty"`

	// ID of this external location - same as name.
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of External Location, which must be unique within the databricks_metastore. Change forces creation of a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Username/groupname/sp application_id of the external location owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Indicates whether the external location is read-only.
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Suppress validation errors if any & force save the external location
	SkipValidation *bool `json:"skipValidation,omitempty" tf:"skip_validation,omitempty"`

	// Path URL in cloud storage, of the form: s3://[bucket-host]/[bucket-dir] (AWS), abfss://[user]@[host]/[path] (Azure), gs://[bucket-host]/[bucket-dir] (GCP).
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ExternalLocationParameters struct {

	// The ARN of the s3 access point to use with the external location (AWS).
	// +kubebuilder:validation:Optional
	AccessPoint *string `json:"accessPoint,omitempty" tf:"access_point,omitempty"`

	// User-supplied free-form text.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the databricks_storage_credential to use with this external location.
	// +kubebuilder:validation:Optional
	CredentialName *string `json:"credentialName,omitempty" tf:"credential_name,omitempty"`

	// The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
	// +kubebuilder:validation:Optional
	EncryptionDetails []EncryptionDetailsParameters `json:"encryptionDetails,omitempty" tf:"encryption_details,omitempty"`

	// Destroy external location regardless of its dependents.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Update external location regardless of its dependents.
	// +kubebuilder:validation:Optional
	ForceUpdate *bool `json:"forceUpdate,omitempty" tf:"force_update,omitempty"`

	// Whether the external location is accessible from all workspaces or a specific set of workspaces. Can be ISOLATION_MODE_ISOLATED or ISOLATION_MODE_OPEN. Setting the external location to ISOLATION_MODE_ISOLATED will automatically allow access from the current workspace.
	// +kubebuilder:validation:Optional
	IsolationMode *string `json:"isolationMode,omitempty" tf:"isolation_mode,omitempty"`

	// ID of this external location - same as name.
	// +kubebuilder:validation:Optional
	MetastoreID *string `json:"metastoreId,omitempty" tf:"metastore_id,omitempty"`

	// Name of External Location, which must be unique within the databricks_metastore. Change forces creation of a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Username/groupname/sp application_id of the external location owner.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Indicates whether the external location is read-only.
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Suppress validation errors if any & force save the external location
	// +kubebuilder:validation:Optional
	SkipValidation *bool `json:"skipValidation,omitempty" tf:"skip_validation,omitempty"`

	// Path URL in cloud storage, of the form: s3://[bucket-host]/[bucket-dir] (AWS), abfss://[user]@[host]/[path] (Azure), gs://[bucket-host]/[bucket-dir] (GCP).
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SseEncryptionDetailsInitParameters struct {
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	AwsKMSKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`
}

type SseEncryptionDetailsObservation struct {
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	AwsKMSKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`
}

type SseEncryptionDetailsParameters struct {

	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	AwsKMSKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`
}

// ExternalLocationSpec defines the desired state of ExternalLocation
type ExternalLocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExternalLocationParameters `json:"forProvider"`
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
	InitProvider ExternalLocationInitParameters `json:"initProvider,omitempty"`
}

// ExternalLocationStatus defines the observed state of ExternalLocation.
type ExternalLocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExternalLocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ExternalLocation is the Schema for the ExternalLocations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type ExternalLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialName) || (has(self.initProvider) && has(self.initProvider.credentialName))",message="spec.forProvider.credentialName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   ExternalLocationSpec   `json:"spec"`
	Status ExternalLocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalLocationList contains a list of ExternalLocations
type ExternalLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalLocation `json:"items"`
}

// Repository type metadata.
var (
	ExternalLocation_Kind             = "ExternalLocation"
	ExternalLocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExternalLocation_Kind}.String()
	ExternalLocation_KindAPIVersion   = ExternalLocation_Kind + "." + CRDGroupVersion.String()
	ExternalLocation_GroupVersionKind = CRDGroupVersion.WithKind(ExternalLocation_Kind)
)

func init() {
	SchemeBuilder.Register(&ExternalLocation{}, &ExternalLocationList{})
}
