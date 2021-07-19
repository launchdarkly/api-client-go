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

// FlagsTarget struct for FlagsTarget
type FlagsTarget struct {
	Values *[]string `json:"values,omitempty"`
	Variation *int32 `json:"variation,omitempty"`
}

// NewFlagsTarget instantiates a new FlagsTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagsTarget() *FlagsTarget {
	this := FlagsTarget{}
	return &this
}

// NewFlagsTargetWithDefaults instantiates a new FlagsTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagsTargetWithDefaults() *FlagsTarget {
	this := FlagsTarget{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FlagsTarget) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagsTarget) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FlagsTarget) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *FlagsTarget) SetValues(v []string) {
	o.Values = &v
}

// GetVariation returns the Variation field value if set, zero value otherwise.
func (o *FlagsTarget) GetVariation() int32 {
	if o == nil || o.Variation == nil {
		var ret int32
		return ret
	}
	return *o.Variation
}

// GetVariationOk returns a tuple with the Variation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagsTarget) GetVariationOk() (*int32, bool) {
	if o == nil || o.Variation == nil {
		return nil, false
	}
	return o.Variation, true
}

// HasVariation returns a boolean if a field has been set.
func (o *FlagsTarget) HasVariation() bool {
	if o != nil && o.Variation != nil {
		return true
	}

	return false
}

// SetVariation gets a reference to the given int32 and assigns it to the Variation field.
func (o *FlagsTarget) SetVariation(v int32) {
	o.Variation = &v
}

func (o FlagsTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Variation != nil {
		toSerialize["variation"] = o.Variation
	}
	return json.Marshal(toSerialize)
}

type NullableFlagsTarget struct {
	value *FlagsTarget
	isSet bool
}

func (v NullableFlagsTarget) Get() *FlagsTarget {
	return v.value
}

func (v *NullableFlagsTarget) Set(val *FlagsTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagsTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagsTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagsTarget(val *FlagsTarget) *NullableFlagsTarget {
	return &NullableFlagsTarget{value: val, isSet: true}
}

func (v NullableFlagsTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagsTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


