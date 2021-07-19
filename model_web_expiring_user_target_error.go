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

// WebExpiringUserTargetError struct for WebExpiringUserTargetError
type WebExpiringUserTargetError struct {
	InstructionIndex *int32 `json:"instructionIndex,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewWebExpiringUserTargetError instantiates a new WebExpiringUserTargetError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebExpiringUserTargetError() *WebExpiringUserTargetError {
	this := WebExpiringUserTargetError{}
	return &this
}

// NewWebExpiringUserTargetErrorWithDefaults instantiates a new WebExpiringUserTargetError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebExpiringUserTargetErrorWithDefaults() *WebExpiringUserTargetError {
	this := WebExpiringUserTargetError{}
	return &this
}

// GetInstructionIndex returns the InstructionIndex field value if set, zero value otherwise.
func (o *WebExpiringUserTargetError) GetInstructionIndex() int32 {
	if o == nil || o.InstructionIndex == nil {
		var ret int32
		return ret
	}
	return *o.InstructionIndex
}

// GetInstructionIndexOk returns a tuple with the InstructionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebExpiringUserTargetError) GetInstructionIndexOk() (*int32, bool) {
	if o == nil || o.InstructionIndex == nil {
		return nil, false
	}
	return o.InstructionIndex, true
}

// HasInstructionIndex returns a boolean if a field has been set.
func (o *WebExpiringUserTargetError) HasInstructionIndex() bool {
	if o != nil && o.InstructionIndex != nil {
		return true
	}

	return false
}

// SetInstructionIndex gets a reference to the given int32 and assigns it to the InstructionIndex field.
func (o *WebExpiringUserTargetError) SetInstructionIndex(v int32) {
	o.InstructionIndex = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WebExpiringUserTargetError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebExpiringUserTargetError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WebExpiringUserTargetError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WebExpiringUserTargetError) SetMessage(v string) {
	o.Message = &v
}

func (o WebExpiringUserTargetError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstructionIndex != nil {
		toSerialize["instructionIndex"] = o.InstructionIndex
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableWebExpiringUserTargetError struct {
	value *WebExpiringUserTargetError
	isSet bool
}

func (v NullableWebExpiringUserTargetError) Get() *WebExpiringUserTargetError {
	return v.value
}

func (v *NullableWebExpiringUserTargetError) Set(val *WebExpiringUserTargetError) {
	v.value = val
	v.isSet = true
}

func (v NullableWebExpiringUserTargetError) IsSet() bool {
	return v.isSet
}

func (v *NullableWebExpiringUserTargetError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebExpiringUserTargetError(val *WebExpiringUserTargetError) *NullableWebExpiringUserTargetError {
	return &NullableWebExpiringUserTargetError{value: val, isSet: true}
}

func (v NullableWebExpiringUserTargetError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebExpiringUserTargetError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


