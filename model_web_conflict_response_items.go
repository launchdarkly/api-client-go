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

// WebConflictResponseItems struct for WebConflictResponseItems
type WebConflictResponseItems struct {
	Kind *string `json:"kind,omitempty"`
	Conflicts *[]WebConflictResponseConflicts `json:"conflicts,omitempty"`
}

// NewWebConflictResponseItems instantiates a new WebConflictResponseItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebConflictResponseItems() *WebConflictResponseItems {
	this := WebConflictResponseItems{}
	return &this
}

// NewWebConflictResponseItemsWithDefaults instantiates a new WebConflictResponseItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebConflictResponseItemsWithDefaults() *WebConflictResponseItems {
	this := WebConflictResponseItems{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *WebConflictResponseItems) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebConflictResponseItems) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *WebConflictResponseItems) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *WebConflictResponseItems) SetKind(v string) {
	o.Kind = &v
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *WebConflictResponseItems) GetConflicts() []WebConflictResponseConflicts {
	if o == nil || o.Conflicts == nil {
		var ret []WebConflictResponseConflicts
		return ret
	}
	return *o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebConflictResponseItems) GetConflictsOk() (*[]WebConflictResponseConflicts, bool) {
	if o == nil || o.Conflicts == nil {
		return nil, false
	}
	return o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *WebConflictResponseItems) HasConflicts() bool {
	if o != nil && o.Conflicts != nil {
		return true
	}

	return false
}

// SetConflicts gets a reference to the given []WebConflictResponseConflicts and assigns it to the Conflicts field.
func (o *WebConflictResponseItems) SetConflicts(v []WebConflictResponseConflicts) {
	o.Conflicts = &v
}

func (o WebConflictResponseItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Conflicts != nil {
		toSerialize["conflicts"] = o.Conflicts
	}
	return json.Marshal(toSerialize)
}

type NullableWebConflictResponseItems struct {
	value *WebConflictResponseItems
	isSet bool
}

func (v NullableWebConflictResponseItems) Get() *WebConflictResponseItems {
	return v.value
}

func (v *NullableWebConflictResponseItems) Set(val *WebConflictResponseItems) {
	v.value = val
	v.isSet = true
}

func (v NullableWebConflictResponseItems) IsSet() bool {
	return v.isSet
}

func (v *NullableWebConflictResponseItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebConflictResponseItems(val *WebConflictResponseItems) *NullableWebConflictResponseItems {
	return &NullableWebConflictResponseItems{value: val, isSet: true}
}

func (v NullableWebConflictResponseItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebConflictResponseItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


