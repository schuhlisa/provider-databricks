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

type EndpointStatusInitParameters struct {
}

type EndpointStatusObservation struct {

	// Additional status message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Current state of the endpoint. Currently following values are supported: PROVISIONING, ONLINE, and OFFLINE.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type EndpointStatusParameters struct {
}

type VectorSearchEndpointInitParameters struct {

	// Type of Mosaic AI Vector Search Endpoint.  Currently only accepting single value: STANDARD (See documentation for the list of currently supported values).
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Name of the Mosaic AI Vector Search Endpoint to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VectorSearchEndpointObservation struct {

	// Timestamp of endpoint creation (milliseconds).
	CreationTimestamp *float64 `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Creator of the endpoint.
	Creator *string `json:"creator,omitempty" tf:"creator,omitempty"`

	// Unique internal identifier of the endpoint (UUID).
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// Object describing the current status of the endpoint consisting of the following fields:
	EndpointStatus []EndpointStatusObservation `json:"endpointStatus,omitempty" tf:"endpoint_status,omitempty"`

	// Type of Mosaic AI Vector Search Endpoint.  Currently only accepting single value: STANDARD (See documentation for the list of currently supported values).
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// The same as the name of the endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Timestamp of the last update to the endpoint (milliseconds).
	LastUpdatedTimestamp *float64 `json:"lastUpdatedTimestamp,omitempty" tf:"last_updated_timestamp,omitempty"`

	// User who last updated the endpoint.
	LastUpdatedUser *string `json:"lastUpdatedUser,omitempty" tf:"last_updated_user,omitempty"`

	// Name of the Mosaic AI Vector Search Endpoint to create.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of indexes on the endpoint.
	NumIndexes *float64 `json:"numIndexes,omitempty" tf:"num_indexes,omitempty"`
}

type VectorSearchEndpointParameters struct {

	// Type of Mosaic AI Vector Search Endpoint.  Currently only accepting single value: STANDARD (See documentation for the list of currently supported values).
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Name of the Mosaic AI Vector Search Endpoint to create.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// VectorSearchEndpointSpec defines the desired state of VectorSearchEndpoint
type VectorSearchEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VectorSearchEndpointParameters `json:"forProvider"`
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
	InitProvider VectorSearchEndpointInitParameters `json:"initProvider,omitempty"`
}

// VectorSearchEndpointStatus defines the observed state of VectorSearchEndpoint.
type VectorSearchEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VectorSearchEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VectorSearchEndpoint is the Schema for the VectorSearchEndpoints API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type VectorSearchEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpointType) || (has(self.initProvider) && has(self.initProvider.endpointType))",message="spec.forProvider.endpointType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VectorSearchEndpointSpec   `json:"spec"`
	Status VectorSearchEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VectorSearchEndpointList contains a list of VectorSearchEndpoints
type VectorSearchEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VectorSearchEndpoint `json:"items"`
}

// Repository type metadata.
var (
	VectorSearchEndpoint_Kind             = "VectorSearchEndpoint"
	VectorSearchEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VectorSearchEndpoint_Kind}.String()
	VectorSearchEndpoint_KindAPIVersion   = VectorSearchEndpoint_Kind + "." + CRDGroupVersion.String()
	VectorSearchEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(VectorSearchEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&VectorSearchEndpoint{}, &VectorSearchEndpointList{})
}
