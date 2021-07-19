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

// StatementStatementPost struct for StatementStatementPost
type StatementStatementPost struct {
	Resources *[]string `json:"resources,omitempty"`
	NotResources *[]string `json:"notResources,omitempty"`
	Actions *[]string `json:"actions,omitempty"`
	NotActions *[]string `json:"notActions,omitempty"`
	Effect *string `json:"effect,omitempty"`
}

// NewStatementStatementPost instantiates a new StatementStatementPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementStatementPost() *StatementStatementPost {
	this := StatementStatementPost{}
	return &this
}

// NewStatementStatementPostWithDefaults instantiates a new StatementStatementPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementStatementPostWithDefaults() *StatementStatementPost {
	this := StatementStatementPost{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *StatementStatementPost) GetResources() []string {
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementStatementPost) GetResourcesOk() (*[]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *StatementStatementPost) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *StatementStatementPost) SetResources(v []string) {
	o.Resources = &v
}

// GetNotResources returns the NotResources field value if set, zero value otherwise.
func (o *StatementStatementPost) GetNotResources() []string {
	if o == nil || o.NotResources == nil {
		var ret []string
		return ret
	}
	return *o.NotResources
}

// GetNotResourcesOk returns a tuple with the NotResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementStatementPost) GetNotResourcesOk() (*[]string, bool) {
	if o == nil || o.NotResources == nil {
		return nil, false
	}
	return o.NotResources, true
}

// HasNotResources returns a boolean if a field has been set.
func (o *StatementStatementPost) HasNotResources() bool {
	if o != nil && o.NotResources != nil {
		return true
	}

	return false
}

// SetNotResources gets a reference to the given []string and assigns it to the NotResources field.
func (o *StatementStatementPost) SetNotResources(v []string) {
	o.NotResources = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *StatementStatementPost) GetActions() []string {
	if o == nil || o.Actions == nil {
		var ret []string
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementStatementPost) GetActionsOk() (*[]string, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *StatementStatementPost) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *StatementStatementPost) SetActions(v []string) {
	o.Actions = &v
}

// GetNotActions returns the NotActions field value if set, zero value otherwise.
func (o *StatementStatementPost) GetNotActions() []string {
	if o == nil || o.NotActions == nil {
		var ret []string
		return ret
	}
	return *o.NotActions
}

// GetNotActionsOk returns a tuple with the NotActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementStatementPost) GetNotActionsOk() (*[]string, bool) {
	if o == nil || o.NotActions == nil {
		return nil, false
	}
	return o.NotActions, true
}

// HasNotActions returns a boolean if a field has been set.
func (o *StatementStatementPost) HasNotActions() bool {
	if o != nil && o.NotActions != nil {
		return true
	}

	return false
}

// SetNotActions gets a reference to the given []string and assigns it to the NotActions field.
func (o *StatementStatementPost) SetNotActions(v []string) {
	o.NotActions = &v
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *StatementStatementPost) GetEffect() string {
	if o == nil || o.Effect == nil {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementStatementPost) GetEffectOk() (*string, bool) {
	if o == nil || o.Effect == nil {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *StatementStatementPost) HasEffect() bool {
	if o != nil && o.Effect != nil {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *StatementStatementPost) SetEffect(v string) {
	o.Effect = &v
}

func (o StatementStatementPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.NotResources != nil {
		toSerialize["notResources"] = o.NotResources
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.NotActions != nil {
		toSerialize["notActions"] = o.NotActions
	}
	if o.Effect != nil {
		toSerialize["effect"] = o.Effect
	}
	return json.Marshal(toSerialize)
}

type NullableStatementStatementPost struct {
	value *StatementStatementPost
	isSet bool
}

func (v NullableStatementStatementPost) Get() *StatementStatementPost {
	return v.value
}

func (v *NullableStatementStatementPost) Set(val *StatementStatementPost) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementStatementPost) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementStatementPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementStatementPost(val *StatementStatementPost) *NullableStatementStatementPost {
	return &NullableStatementStatementPost{value: val, isSet: true}
}

func (v NullableStatementStatementPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementStatementPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


