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

type DefaultNamespaceSettingInitParameters struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The configuration details.
	Namespace []NamespaceInitParameters `json:"namespace,omitempty" tf:"namespace,omitempty"`

	SettingName *string `json:"settingName,omitempty" tf:"setting_name,omitempty"`
}

type DefaultNamespaceSettingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The configuration details.
	Namespace []NamespaceObservation `json:"namespace,omitempty" tf:"namespace,omitempty"`

	SettingName *string `json:"settingName,omitempty" tf:"setting_name,omitempty"`
}

type DefaultNamespaceSettingParameters struct {

	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The configuration details.
	// +kubebuilder:validation:Optional
	Namespace []NamespaceParameters `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	SettingName *string `json:"settingName,omitempty" tf:"setting_name,omitempty"`
}

type NamespaceInitParameters struct {

	// The value for the setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NamespaceObservation struct {

	// The value for the setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NamespaceParameters struct {

	// The value for the setting.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// DefaultNamespaceSettingSpec defines the desired state of DefaultNamespaceSetting
type DefaultNamespaceSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultNamespaceSettingParameters `json:"forProvider"`
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
	InitProvider DefaultNamespaceSettingInitParameters `json:"initProvider,omitempty"`
}

// DefaultNamespaceSettingStatus defines the observed state of DefaultNamespaceSetting.
type DefaultNamespaceSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultNamespaceSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DefaultNamespaceSetting is the Schema for the DefaultNamespaceSettings API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type DefaultNamespaceSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	Spec   DefaultNamespaceSettingSpec   `json:"spec"`
	Status DefaultNamespaceSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNamespaceSettingList contains a list of DefaultNamespaceSettings
type DefaultNamespaceSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultNamespaceSetting `json:"items"`
}

// Repository type metadata.
var (
	DefaultNamespaceSetting_Kind             = "DefaultNamespaceSetting"
	DefaultNamespaceSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultNamespaceSetting_Kind}.String()
	DefaultNamespaceSetting_KindAPIVersion   = DefaultNamespaceSetting_Kind + "." + CRDGroupVersion.String()
	DefaultNamespaceSetting_GroupVersionKind = CRDGroupVersion.WithKind(DefaultNamespaceSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultNamespaceSetting{}, &DefaultNamespaceSettingList{})
}
