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

// WebReportFlagScheduledChangesInput struct for WebReportFlagScheduledChangesInput
type WebReportFlagScheduledChangesInput struct {
	ExecutionDate *int64 `json:"ExecutionDate,omitempty"`
	ExistingScheduledChangeId *string `json:"ExistingScheduledChangeId,omitempty"`
	Instructions *[]interface{} `json:"Instructions,omitempty"`
}

// NewWebReportFlagScheduledChangesInput instantiates a new WebReportFlagScheduledChangesInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebReportFlagScheduledChangesInput() *WebReportFlagScheduledChangesInput {
	this := WebReportFlagScheduledChangesInput{}
	return &this
}

// NewWebReportFlagScheduledChangesInputWithDefaults instantiates a new WebReportFlagScheduledChangesInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebReportFlagScheduledChangesInputWithDefaults() *WebReportFlagScheduledChangesInput {
	this := WebReportFlagScheduledChangesInput{}
	return &this
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *WebReportFlagScheduledChangesInput) GetExecutionDate() int64 {
	if o == nil || o.ExecutionDate == nil {
		var ret int64
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebReportFlagScheduledChangesInput) GetExecutionDateOk() (*int64, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *WebReportFlagScheduledChangesInput) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given int64 and assigns it to the ExecutionDate field.
func (o *WebReportFlagScheduledChangesInput) SetExecutionDate(v int64) {
	o.ExecutionDate = &v
}

// GetExistingScheduledChangeId returns the ExistingScheduledChangeId field value if set, zero value otherwise.
func (o *WebReportFlagScheduledChangesInput) GetExistingScheduledChangeId() string {
	if o == nil || o.ExistingScheduledChangeId == nil {
		var ret string
		return ret
	}
	return *o.ExistingScheduledChangeId
}

// GetExistingScheduledChangeIdOk returns a tuple with the ExistingScheduledChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebReportFlagScheduledChangesInput) GetExistingScheduledChangeIdOk() (*string, bool) {
	if o == nil || o.ExistingScheduledChangeId == nil {
		return nil, false
	}
	return o.ExistingScheduledChangeId, true
}

// HasExistingScheduledChangeId returns a boolean if a field has been set.
func (o *WebReportFlagScheduledChangesInput) HasExistingScheduledChangeId() bool {
	if o != nil && o.ExistingScheduledChangeId != nil {
		return true
	}

	return false
}

// SetExistingScheduledChangeId gets a reference to the given string and assigns it to the ExistingScheduledChangeId field.
func (o *WebReportFlagScheduledChangesInput) SetExistingScheduledChangeId(v string) {
	o.ExistingScheduledChangeId = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *WebReportFlagScheduledChangesInput) GetInstructions() []interface{} {
	if o == nil || o.Instructions == nil {
		var ret []interface{}
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebReportFlagScheduledChangesInput) GetInstructionsOk() (*[]interface{}, bool) {
	if o == nil || o.Instructions == nil {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *WebReportFlagScheduledChangesInput) HasInstructions() bool {
	if o != nil && o.Instructions != nil {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given []interface{} and assigns it to the Instructions field.
func (o *WebReportFlagScheduledChangesInput) SetInstructions(v []interface{}) {
	o.Instructions = &v
}

func (o WebReportFlagScheduledChangesInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExecutionDate != nil {
		toSerialize["ExecutionDate"] = o.ExecutionDate
	}
	if o.ExistingScheduledChangeId != nil {
		toSerialize["ExistingScheduledChangeId"] = o.ExistingScheduledChangeId
	}
	if o.Instructions != nil {
		toSerialize["Instructions"] = o.Instructions
	}
	return json.Marshal(toSerialize)
}

type NullableWebReportFlagScheduledChangesInput struct {
	value *WebReportFlagScheduledChangesInput
	isSet bool
}

func (v NullableWebReportFlagScheduledChangesInput) Get() *WebReportFlagScheduledChangesInput {
	return v.value
}

func (v *NullableWebReportFlagScheduledChangesInput) Set(val *WebReportFlagScheduledChangesInput) {
	v.value = val
	v.isSet = true
}

func (v NullableWebReportFlagScheduledChangesInput) IsSet() bool {
	return v.isSet
}

func (v *NullableWebReportFlagScheduledChangesInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebReportFlagScheduledChangesInput(val *WebReportFlagScheduledChangesInput) *NullableWebReportFlagScheduledChangesInput {
	return &NullableWebReportFlagScheduledChangesInput{value: val, isSet: true}
}

func (v NullableWebReportFlagScheduledChangesInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebReportFlagScheduledChangesInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


