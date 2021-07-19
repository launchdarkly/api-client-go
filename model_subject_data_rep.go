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

// SubjectDataRep struct for SubjectDataRep
type SubjectDataRep struct {
	Links *map[string]InlineResponse200 `json:"_links,omitempty"`
	Name *string `json:"name,omitempty"`
	AvatarUrl *string `json:"avatarUrl,omitempty"`
}

// NewSubjectDataRep instantiates a new SubjectDataRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectDataRep() *SubjectDataRep {
	this := SubjectDataRep{}
	return &this
}

// NewSubjectDataRepWithDefaults instantiates a new SubjectDataRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectDataRepWithDefaults() *SubjectDataRep {
	this := SubjectDataRep{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubjectDataRep) GetLinks() map[string]InlineResponse200 {
	if o == nil || o.Links == nil {
		var ret map[string]InlineResponse200
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectDataRep) GetLinksOk() (*map[string]InlineResponse200, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubjectDataRep) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]InlineResponse200 and assigns it to the Links field.
func (o *SubjectDataRep) SetLinks(v map[string]InlineResponse200) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubjectDataRep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectDataRep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubjectDataRep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubjectDataRep) SetName(v string) {
	o.Name = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *SubjectDataRep) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectDataRep) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *SubjectDataRep) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *SubjectDataRep) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

func (o SubjectDataRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AvatarUrl != nil {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSubjectDataRep struct {
	value *SubjectDataRep
	isSet bool
}

func (v NullableSubjectDataRep) Get() *SubjectDataRep {
	return v.value
}

func (v *NullableSubjectDataRep) Set(val *SubjectDataRep) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectDataRep) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectDataRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectDataRep(val *SubjectDataRep) *NullableSubjectDataRep {
	return &NullableSubjectDataRep{value: val, isSet: true}
}

func (v NullableSubjectDataRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectDataRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


