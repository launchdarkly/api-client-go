/*
 * LaunchDarkly REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0
 * Contact: support@launchdarkly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"encoding/json"
)

// ExperimentEnvironmentSettingRep struct for ExperimentEnvironmentSettingRep
type ExperimentEnvironmentSettingRep struct {
	StartDate *int64 `json:"startDate,omitempty"`
	StopDate *int64 `json:"stopDate,omitempty"`
}

// NewExperimentEnvironmentSettingRep instantiates a new ExperimentEnvironmentSettingRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentEnvironmentSettingRep() *ExperimentEnvironmentSettingRep {
	this := ExperimentEnvironmentSettingRep{}
	return &this
}

// NewExperimentEnvironmentSettingRepWithDefaults instantiates a new ExperimentEnvironmentSettingRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentEnvironmentSettingRepWithDefaults() *ExperimentEnvironmentSettingRep {
	this := ExperimentEnvironmentSettingRep{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ExperimentEnvironmentSettingRep) GetStartDate() int64 {
	if o == nil || o.StartDate == nil {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentEnvironmentSettingRep) GetStartDateOk() (*int64, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ExperimentEnvironmentSettingRep) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *ExperimentEnvironmentSettingRep) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetStopDate returns the StopDate field value if set, zero value otherwise.
func (o *ExperimentEnvironmentSettingRep) GetStopDate() int64 {
	if o == nil || o.StopDate == nil {
		var ret int64
		return ret
	}
	return *o.StopDate
}

// GetStopDateOk returns a tuple with the StopDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentEnvironmentSettingRep) GetStopDateOk() (*int64, bool) {
	if o == nil || o.StopDate == nil {
		return nil, false
	}
	return o.StopDate, true
}

// HasStopDate returns a boolean if a field has been set.
func (o *ExperimentEnvironmentSettingRep) HasStopDate() bool {
	if o != nil && o.StopDate != nil {
		return true
	}

	return false
}

// SetStopDate gets a reference to the given int64 and assigns it to the StopDate field.
func (o *ExperimentEnvironmentSettingRep) SetStopDate(v int64) {
	o.StopDate = &v
}

func (o ExperimentEnvironmentSettingRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.StopDate != nil {
		toSerialize["stopDate"] = o.StopDate
	}
	return json.Marshal(toSerialize)
}

type NullableExperimentEnvironmentSettingRep struct {
	value *ExperimentEnvironmentSettingRep
	isSet bool
}

func (v NullableExperimentEnvironmentSettingRep) Get() *ExperimentEnvironmentSettingRep {
	return v.value
}

func (v *NullableExperimentEnvironmentSettingRep) Set(val *ExperimentEnvironmentSettingRep) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentEnvironmentSettingRep) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentEnvironmentSettingRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentEnvironmentSettingRep(val *ExperimentEnvironmentSettingRep) *NullableExperimentEnvironmentSettingRep {
	return &NullableExperimentEnvironmentSettingRep{value: val, isSet: true}
}

func (v NullableExperimentEnvironmentSettingRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentEnvironmentSettingRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


