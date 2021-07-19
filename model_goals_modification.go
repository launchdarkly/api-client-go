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
	"time"
)

// GoalsModification struct for GoalsModification
type GoalsModification struct {
	Date *time.Time `json:"date,omitempty"`
}

// NewGoalsModification instantiates a new GoalsModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoalsModification() *GoalsModification {
	this := GoalsModification{}
	return &this
}

// NewGoalsModificationWithDefaults instantiates a new GoalsModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoalsModificationWithDefaults() *GoalsModification {
	this := GoalsModification{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GoalsModification) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoalsModification) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GoalsModification) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *GoalsModification) SetDate(v time.Time) {
	o.Date = &v
}

func (o GoalsModification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableGoalsModification struct {
	value *GoalsModification
	isSet bool
}

func (v NullableGoalsModification) Get() *GoalsModification {
	return v.value
}

func (v *NullableGoalsModification) Set(val *GoalsModification) {
	v.value = val
	v.isSet = true
}

func (v NullableGoalsModification) IsSet() bool {
	return v.isSet
}

func (v *NullableGoalsModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoalsModification(val *GoalsModification) *NullableGoalsModification {
	return &NullableGoalsModification{value: val, isSet: true}
}

func (v NullableGoalsModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoalsModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


