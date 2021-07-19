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

// FlagsFlagCopyConfigEnvironment struct for FlagsFlagCopyConfigEnvironment
type FlagsFlagCopyConfigEnvironment struct {
	Key *string `json:"key,omitempty"`
	CurrentVersion *int32 `json:"currentVersion,omitempty"`
}

// NewFlagsFlagCopyConfigEnvironment instantiates a new FlagsFlagCopyConfigEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagsFlagCopyConfigEnvironment() *FlagsFlagCopyConfigEnvironment {
	this := FlagsFlagCopyConfigEnvironment{}
	return &this
}

// NewFlagsFlagCopyConfigEnvironmentWithDefaults instantiates a new FlagsFlagCopyConfigEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagsFlagCopyConfigEnvironmentWithDefaults() *FlagsFlagCopyConfigEnvironment {
	this := FlagsFlagCopyConfigEnvironment{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FlagsFlagCopyConfigEnvironment) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagsFlagCopyConfigEnvironment) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FlagsFlagCopyConfigEnvironment) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FlagsFlagCopyConfigEnvironment) SetKey(v string) {
	o.Key = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *FlagsFlagCopyConfigEnvironment) GetCurrentVersion() int32 {
	if o == nil || o.CurrentVersion == nil {
		var ret int32
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagsFlagCopyConfigEnvironment) GetCurrentVersionOk() (*int32, bool) {
	if o == nil || o.CurrentVersion == nil {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *FlagsFlagCopyConfigEnvironment) HasCurrentVersion() bool {
	if o != nil && o.CurrentVersion != nil {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given int32 and assigns it to the CurrentVersion field.
func (o *FlagsFlagCopyConfigEnvironment) SetCurrentVersion(v int32) {
	o.CurrentVersion = &v
}

func (o FlagsFlagCopyConfigEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.CurrentVersion != nil {
		toSerialize["currentVersion"] = o.CurrentVersion
	}
	return json.Marshal(toSerialize)
}

type NullableFlagsFlagCopyConfigEnvironment struct {
	value *FlagsFlagCopyConfigEnvironment
	isSet bool
}

func (v NullableFlagsFlagCopyConfigEnvironment) Get() *FlagsFlagCopyConfigEnvironment {
	return v.value
}

func (v *NullableFlagsFlagCopyConfigEnvironment) Set(val *FlagsFlagCopyConfigEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagsFlagCopyConfigEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagsFlagCopyConfigEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagsFlagCopyConfigEnvironment(val *FlagsFlagCopyConfigEnvironment) *NullableFlagsFlagCopyConfigEnvironment {
	return &NullableFlagsFlagCopyConfigEnvironment{value: val, isSet: true}
}

func (v NullableFlagsFlagCopyConfigEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagsFlagCopyConfigEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


