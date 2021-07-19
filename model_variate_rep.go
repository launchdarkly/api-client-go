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

// VariateRep struct for VariateRep
type VariateRep struct {
	Id *string `json:"_id,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewVariateRep instantiates a new VariateRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariateRep() *VariateRep {
	this := VariateRep{}
	return &this
}

// NewVariateRepWithDefaults instantiates a new VariateRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariateRepWithDefaults() *VariateRep {
	this := VariateRep{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VariateRep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariateRep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VariateRep) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VariateRep) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VariateRep) GetValue() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VariateRep) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VariateRep) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *VariateRep) SetValue(v interface{}) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VariateRep) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariateRep) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VariateRep) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VariateRep) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VariateRep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariateRep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VariateRep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VariateRep) SetName(v string) {
	o.Name = &v
}

func (o VariateRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableVariateRep struct {
	value *VariateRep
	isSet bool
}

func (v NullableVariateRep) Get() *VariateRep {
	return v.value
}

func (v *NullableVariateRep) Set(val *VariateRep) {
	v.value = val
	v.isSet = true
}

func (v NullableVariateRep) IsSet() bool {
	return v.isSet
}

func (v *NullableVariateRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariateRep(val *VariateRep) *NullableVariateRep {
	return &NullableVariateRep{value: val, isSet: true}
}

func (v NullableVariateRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariateRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


