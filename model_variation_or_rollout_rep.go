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

// VariationOrRolloutRep struct for VariationOrRolloutRep
type VariationOrRolloutRep struct {
	Variation *int32 `json:"variation,omitempty"`
	Rollout *RolloutRep `json:"rollout,omitempty"`
}

// NewVariationOrRolloutRep instantiates a new VariationOrRolloutRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationOrRolloutRep() *VariationOrRolloutRep {
	this := VariationOrRolloutRep{}
	return &this
}

// NewVariationOrRolloutRepWithDefaults instantiates a new VariationOrRolloutRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationOrRolloutRepWithDefaults() *VariationOrRolloutRep {
	this := VariationOrRolloutRep{}
	return &this
}

// GetVariation returns the Variation field value if set, zero value otherwise.
func (o *VariationOrRolloutRep) GetVariation() int32 {
	if o == nil || o.Variation == nil {
		var ret int32
		return ret
	}
	return *o.Variation
}

// GetVariationOk returns a tuple with the Variation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationOrRolloutRep) GetVariationOk() (*int32, bool) {
	if o == nil || o.Variation == nil {
		return nil, false
	}
	return o.Variation, true
}

// HasVariation returns a boolean if a field has been set.
func (o *VariationOrRolloutRep) HasVariation() bool {
	if o != nil && o.Variation != nil {
		return true
	}

	return false
}

// SetVariation gets a reference to the given int32 and assigns it to the Variation field.
func (o *VariationOrRolloutRep) SetVariation(v int32) {
	o.Variation = &v
}

// GetRollout returns the Rollout field value if set, zero value otherwise.
func (o *VariationOrRolloutRep) GetRollout() RolloutRep {
	if o == nil || o.Rollout == nil {
		var ret RolloutRep
		return ret
	}
	return *o.Rollout
}

// GetRolloutOk returns a tuple with the Rollout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationOrRolloutRep) GetRolloutOk() (*RolloutRep, bool) {
	if o == nil || o.Rollout == nil {
		return nil, false
	}
	return o.Rollout, true
}

// HasRollout returns a boolean if a field has been set.
func (o *VariationOrRolloutRep) HasRollout() bool {
	if o != nil && o.Rollout != nil {
		return true
	}

	return false
}

// SetRollout gets a reference to the given RolloutRep and assigns it to the Rollout field.
func (o *VariationOrRolloutRep) SetRollout(v RolloutRep) {
	o.Rollout = &v
}

func (o VariationOrRolloutRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Variation != nil {
		toSerialize["variation"] = o.Variation
	}
	if o.Rollout != nil {
		toSerialize["rollout"] = o.Rollout
	}
	return json.Marshal(toSerialize)
}

type NullableVariationOrRolloutRep struct {
	value *VariationOrRolloutRep
	isSet bool
}

func (v NullableVariationOrRolloutRep) Get() *VariationOrRolloutRep {
	return v.value
}

func (v *NullableVariationOrRolloutRep) Set(val *VariationOrRolloutRep) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationOrRolloutRep) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationOrRolloutRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationOrRolloutRep(val *VariationOrRolloutRep) *NullableVariationOrRolloutRep {
	return &NullableVariationOrRolloutRep{value: val, isSet: true}
}

func (v NullableVariationOrRolloutRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationOrRolloutRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


