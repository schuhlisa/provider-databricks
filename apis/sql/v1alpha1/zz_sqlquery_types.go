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

type ContinuousInitParameters struct {
	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type ContinuousObservation struct {
	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type ContinuousParameters struct {

	// +kubebuilder:validation:Optional
	IntervalSeconds *float64 `json:"intervalSeconds" tf:"interval_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type DailyInitParameters struct {
	IntervalDays *float64 `json:"intervalDays,omitempty" tf:"interval_days,omitempty"`

	TimeOfDay *string `json:"timeOfDay,omitempty" tf:"time_of_day,omitempty"`

	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type DailyObservation struct {
	IntervalDays *float64 `json:"intervalDays,omitempty" tf:"interval_days,omitempty"`

	TimeOfDay *string `json:"timeOfDay,omitempty" tf:"time_of_day,omitempty"`

	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type DailyParameters struct {

	// +kubebuilder:validation:Optional
	IntervalDays *float64 `json:"intervalDays" tf:"interval_days,omitempty"`

	// +kubebuilder:validation:Optional
	TimeOfDay *string `json:"timeOfDay" tf:"time_of_day,omitempty"`

	// +kubebuilder:validation:Optional
	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type DateInitParameters struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DateObservation struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DateParameters struct {

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DateRangeInitParameters struct {
	Range []RangeInitParameters `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DateRangeObservation struct {
	Range []RangeObservation `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DateRangeParameters struct {

	// +kubebuilder:validation:Optional
	Range []RangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimeInitParameters struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimeObservation struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimeParameters struct {

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DatetimeRangeInitParameters struct {
	Range []DatetimeRangeRangeInitParameters `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimeRangeObservation struct {
	Range []DatetimeRangeRangeObservation `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimeRangeParameters struct {

	// +kubebuilder:validation:Optional
	Range []DatetimeRangeRangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimeRangeRangeInitParameters struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type DatetimeRangeRangeObservation struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type DatetimeRangeRangeParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type DatetimesecInitParameters struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimesecObservation struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimesecParameters struct {

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DatetimesecRangeInitParameters struct {
	Range []DatetimesecRangeRangeInitParameters `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimesecRangeObservation struct {
	Range []DatetimesecRangeRangeObservation `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimesecRangeParameters struct {

	// +kubebuilder:validation:Optional
	Range []DatetimesecRangeRangeParameters `json:"range,omitempty" tf:"range,omitempty"`

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DatetimesecRangeRangeInitParameters struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type DatetimesecRangeRangeObservation struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type DatetimesecRangeRangeParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type EnumInitParameters struct {
	Multiple []MultipleInitParameters `json:"multiple,omitempty" tf:"multiple,omitempty"`

	Options []*string `json:"options,omitempty" tf:"options,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type EnumObservation struct {
	Multiple []MultipleObservation `json:"multiple,omitempty" tf:"multiple,omitempty"`

	Options []*string `json:"options,omitempty" tf:"options,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type EnumParameters struct {

	// +kubebuilder:validation:Optional
	Multiple []MultipleParameters `json:"multiple,omitempty" tf:"multiple,omitempty"`

	// +kubebuilder:validation:Optional
	Options []*string `json:"options" tf:"options,omitempty"`

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MultipleInitParameters struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	Separator *string `json:"separator,omitempty" tf:"separator,omitempty"`

	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type MultipleObservation struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	Separator *string `json:"separator,omitempty" tf:"separator,omitempty"`

	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type MultipleParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Separator *string `json:"separator" tf:"separator,omitempty"`

	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type NumberInitParameters struct {

	// The default value for this parameter.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type NumberObservation struct {

	// The default value for this parameter.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type NumberParameters struct {

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type ParameterQueryInitParameters struct {
	Multiple []QueryMultipleInitParameters `json:"multiple,omitempty" tf:"multiple,omitempty"`

	// the unique ID of the SQL Query.
	QueryID *string `json:"queryId,omitempty" tf:"query_id,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ParameterQueryObservation struct {
	Multiple []QueryMultipleObservation `json:"multiple,omitempty" tf:"multiple,omitempty"`

	// the unique ID of the SQL Query.
	QueryID *string `json:"queryId,omitempty" tf:"query_id,omitempty"`

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ParameterQueryParameters struct {

	// +kubebuilder:validation:Optional
	Multiple []QueryMultipleParameters `json:"multiple,omitempty" tf:"multiple,omitempty"`

	// the unique ID of the SQL Query.
	// +kubebuilder:validation:Optional
	QueryID *string `json:"queryId" tf:"query_id,omitempty"`

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type QueryMultipleInitParameters struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	Separator *string `json:"separator,omitempty" tf:"separator,omitempty"`

	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type QueryMultipleObservation struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	Separator *string `json:"separator,omitempty" tf:"separator,omitempty"`

	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type QueryMultipleParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Separator *string `json:"separator" tf:"separator,omitempty"`

	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type RangeInitParameters struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type RangeObservation struct {
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type RangeParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

type SQLQueryInitParameters struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Data source ID of a SQL warehouse
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// General description that conveys additional information about this query such as usage notes.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The title of this query that appears in list views, widget headings, and on the query page.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Parameter []SQLQueryParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The identifier of the workspace folder containing the object.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// The text of the query to be run.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Run as role. Possible values are viewer, owner.
	RunAsRole *string `json:"runAsRole,omitempty" tf:"run_as_role,omitempty"`

	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type SQLQueryObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Data source ID of a SQL warehouse
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// General description that conveys additional information about this query such as usage notes.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// the unique ID of the SQL Query.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The title of this query that appears in list views, widget headings, and on the query page.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Parameter []SQLQueryParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The identifier of the workspace folder containing the object.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// The text of the query to be run.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Run as role. Possible values are viewer, owner.
	RunAsRole *string `json:"runAsRole,omitempty" tf:"run_as_role,omitempty"`

	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type SQLQueryParameterInitParameters struct {
	Date []DateInitParameters `json:"date,omitempty" tf:"date,omitempty"`

	DateRange []DateRangeInitParameters `json:"dateRange,omitempty" tf:"date_range,omitempty"`

	Datetime []DatetimeInitParameters `json:"datetime,omitempty" tf:"datetime,omitempty"`

	DatetimeRange []DatetimeRangeInitParameters `json:"datetimeRange,omitempty" tf:"datetime_range,omitempty"`

	Datetimesec []DatetimesecInitParameters `json:"datetimesec,omitempty" tf:"datetimesec,omitempty"`

	DatetimesecRange []DatetimesecRangeInitParameters `json:"datetimesecRange,omitempty" tf:"datetimesec_range,omitempty"`

	Enum []EnumInitParameters `json:"enum,omitempty" tf:"enum,omitempty"`

	// The title of this query that appears in list views, widget headings, and on the query page.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Number []NumberInitParameters `json:"number,omitempty" tf:"number,omitempty"`

	// The text of the query to be run.
	Query []ParameterQueryInitParameters `json:"query,omitempty" tf:"query,omitempty"`

	Text []TextInitParameters `json:"text,omitempty" tf:"text,omitempty"`

	// The text displayed in a parameter picking widget.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SQLQueryParameterObservation struct {
	Date []DateObservation `json:"date,omitempty" tf:"date,omitempty"`

	DateRange []DateRangeObservation `json:"dateRange,omitempty" tf:"date_range,omitempty"`

	Datetime []DatetimeObservation `json:"datetime,omitempty" tf:"datetime,omitempty"`

	DatetimeRange []DatetimeRangeObservation `json:"datetimeRange,omitempty" tf:"datetime_range,omitempty"`

	Datetimesec []DatetimesecObservation `json:"datetimesec,omitempty" tf:"datetimesec,omitempty"`

	DatetimesecRange []DatetimesecRangeObservation `json:"datetimesecRange,omitempty" tf:"datetimesec_range,omitempty"`

	Enum []EnumObservation `json:"enum,omitempty" tf:"enum,omitempty"`

	// The title of this query that appears in list views, widget headings, and on the query page.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Number []NumberObservation `json:"number,omitempty" tf:"number,omitempty"`

	// The text of the query to be run.
	Query []ParameterQueryObservation `json:"query,omitempty" tf:"query,omitempty"`

	Text []TextObservation `json:"text,omitempty" tf:"text,omitempty"`

	// The text displayed in a parameter picking widget.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SQLQueryParameterParameters struct {

	// +kubebuilder:validation:Optional
	Date []DateParameters `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	DateRange []DateRangeParameters `json:"dateRange,omitempty" tf:"date_range,omitempty"`

	// +kubebuilder:validation:Optional
	Datetime []DatetimeParameters `json:"datetime,omitempty" tf:"datetime,omitempty"`

	// +kubebuilder:validation:Optional
	DatetimeRange []DatetimeRangeParameters `json:"datetimeRange,omitempty" tf:"datetime_range,omitempty"`

	// +kubebuilder:validation:Optional
	Datetimesec []DatetimesecParameters `json:"datetimesec,omitempty" tf:"datetimesec,omitempty"`

	// +kubebuilder:validation:Optional
	DatetimesecRange []DatetimesecRangeParameters `json:"datetimesecRange,omitempty" tf:"datetimesec_range,omitempty"`

	// +kubebuilder:validation:Optional
	Enum []EnumParameters `json:"enum,omitempty" tf:"enum,omitempty"`

	// The title of this query that appears in list views, widget headings, and on the query page.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Number []NumberParameters `json:"number,omitempty" tf:"number,omitempty"`

	// The text of the query to be run.
	// +kubebuilder:validation:Optional
	Query []ParameterQueryParameters `json:"query,omitempty" tf:"query,omitempty"`

	// +kubebuilder:validation:Optional
	Text []TextParameters `json:"text,omitempty" tf:"text,omitempty"`

	// The text displayed in a parameter picking widget.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SQLQueryParameters struct {

	// +kubebuilder:validation:Optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Data source ID of a SQL warehouse
	// +kubebuilder:validation:Optional
	DataSourceID *string `json:"dataSourceId,omitempty" tf:"data_source_id,omitempty"`

	// General description that conveys additional information about this query such as usage notes.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The title of this query that appears in list views, widget headings, and on the query page.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameter []SQLQueryParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// The identifier of the workspace folder containing the object.
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// The text of the query to be run.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// Run as role. Possible values are viewer, owner.
	// +kubebuilder:validation:Optional
	RunAsRole *string `json:"runAsRole,omitempty" tf:"run_as_role,omitempty"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ScheduleInitParameters struct {
	Continuous []ContinuousInitParameters `json:"continuous,omitempty" tf:"continuous,omitempty"`

	Daily []DailyInitParameters `json:"daily,omitempty" tf:"daily,omitempty"`

	Weekly []WeeklyInitParameters `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type ScheduleObservation struct {
	Continuous []ContinuousObservation `json:"continuous,omitempty" tf:"continuous,omitempty"`

	Daily []DailyObservation `json:"daily,omitempty" tf:"daily,omitempty"`

	Weekly []WeeklyObservation `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	Continuous []ContinuousParameters `json:"continuous,omitempty" tf:"continuous,omitempty"`

	// +kubebuilder:validation:Optional
	Daily []DailyParameters `json:"daily,omitempty" tf:"daily,omitempty"`

	// +kubebuilder:validation:Optional
	Weekly []WeeklyParameters `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type TextInitParameters struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TextObservation struct {

	// The default value for this parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TextParameters struct {

	// The default value for this parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type WeeklyInitParameters struct {
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	IntervalWeeks *float64 `json:"intervalWeeks,omitempty" tf:"interval_weeks,omitempty"`

	TimeOfDay *string `json:"timeOfDay,omitempty" tf:"time_of_day,omitempty"`

	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type WeeklyObservation struct {
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	IntervalWeeks *float64 `json:"intervalWeeks,omitempty" tf:"interval_weeks,omitempty"`

	TimeOfDay *string `json:"timeOfDay,omitempty" tf:"time_of_day,omitempty"`

	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

type WeeklyParameters struct {

	// +kubebuilder:validation:Optional
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	IntervalWeeks *float64 `json:"intervalWeeks" tf:"interval_weeks,omitempty"`

	// +kubebuilder:validation:Optional
	TimeOfDay *string `json:"timeOfDay" tf:"time_of_day,omitempty"`

	// +kubebuilder:validation:Optional
	UntilDate *string `json:"untilDate,omitempty" tf:"until_date,omitempty"`
}

// SQLQuerySpec defines the desired state of SQLQuery
type SQLQuerySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLQueryParameters `json:"forProvider"`
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
	InitProvider SQLQueryInitParameters `json:"initProvider,omitempty"`
}

// SQLQueryStatus defines the observed state of SQLQuery.
type SQLQueryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLQueryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLQuery is the Schema for the SQLQuerys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type SQLQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataSourceId) || (has(self.initProvider) && has(self.initProvider.dataSourceId))",message="spec.forProvider.dataSourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.query) || (has(self.initProvider) && has(self.initProvider.query))",message="spec.forProvider.query is a required parameter"
	Spec   SQLQuerySpec   `json:"spec"`
	Status SQLQueryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLQueryList contains a list of SQLQuerys
type SQLQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLQuery `json:"items"`
}

// Repository type metadata.
var (
	SQLQuery_Kind             = "SQLQuery"
	SQLQuery_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLQuery_Kind}.String()
	SQLQuery_KindAPIVersion   = SQLQuery_Kind + "." + CRDGroupVersion.String()
	SQLQuery_GroupVersionKind = CRDGroupVersion.WithKind(SQLQuery_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLQuery{}, &SQLQueryList{})
}
