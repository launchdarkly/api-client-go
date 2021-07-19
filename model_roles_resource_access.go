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

// RolesResourceAccess struct for RolesResourceAccess
type RolesResourceAccess struct {
	Action *string `json:"action,omitempty"`
	Resource interface{} `json:"resource,omitempty"`
}

// NewRolesResourceAccess instantiates a new RolesResourceAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesResourceAccess() *RolesResourceAccess {
	this := RolesResourceAccess{}
	return &this
}

// NewRolesResourceAccessWithDefaults instantiates a new RolesResourceAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesResourceAccessWithDefaults() *RolesResourceAccess {
	this := RolesResourceAccess{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RolesResourceAccess) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesResourceAccess) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RolesResourceAccess) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RolesResourceAccess) SetAction(v string) {
	o.Action = &v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesResourceAccess) GetResource() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesResourceAccess) GetResourceOk() (*interface{}, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *RolesResourceAccess) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given interface{} and assigns it to the Resource field.
func (o *RolesResourceAccess) SetResource(v interface{}) {
	o.Resource = v
}

func (o RolesResourceAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableRolesResourceAccess struct {
	value *RolesResourceAccess
	isSet bool
}

func (v NullableRolesResourceAccess) Get() *RolesResourceAccess {
	return v.value
}

func (v *NullableRolesResourceAccess) Set(val *RolesResourceAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesResourceAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesResourceAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesResourceAccess(val *RolesResourceAccess) *NullableRolesResourceAccess {
	return &NullableRolesResourceAccess{value: val, isSet: true}
}

func (v NullableRolesResourceAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesResourceAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


