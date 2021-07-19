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

// ExperimentMetadataRep struct for ExperimentMetadataRep
type ExperimentMetadataRep struct {
	Key interface{} `json:"key,omitempty"`
}

// NewExperimentMetadataRep instantiates a new ExperimentMetadataRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentMetadataRep() *ExperimentMetadataRep {
	this := ExperimentMetadataRep{}
	return &this
}

// NewExperimentMetadataRepWithDefaults instantiates a new ExperimentMetadataRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentMetadataRepWithDefaults() *ExperimentMetadataRep {
	this := ExperimentMetadataRep{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentMetadataRep) GetKey() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExperimentMetadataRep) GetKeyOk() (*interface{}, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ExperimentMetadataRep) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *ExperimentMetadataRep) SetKey(v interface{}) {
	o.Key = v
}

func (o ExperimentMetadataRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableExperimentMetadataRep struct {
	value *ExperimentMetadataRep
	isSet bool
}

func (v NullableExperimentMetadataRep) Get() *ExperimentMetadataRep {
	return v.value
}

func (v *NullableExperimentMetadataRep) Set(val *ExperimentMetadataRep) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentMetadataRep) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentMetadataRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentMetadataRep(val *ExperimentMetadataRep) *NullableExperimentMetadataRep {
	return &NullableExperimentMetadataRep{value: val, isSet: true}
}

func (v NullableExperimentMetadataRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentMetadataRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


