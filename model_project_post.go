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

// ProjectPost struct for ProjectPost
type ProjectPost struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	IncludeInSnippetByDefault *bool `json:"includeInSnippetByDefault,omitempty"`
	DefaultClientSideAvailability *DefaultClientSideAvailabilityPost `json:"defaultClientSideAvailability,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Environments *[]map[string]interface{} `json:"environments,omitempty"`
}

// NewProjectPost instantiates a new ProjectPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectPost() *ProjectPost {
	this := ProjectPost{}
	return &this
}

// NewProjectPostWithDefaults instantiates a new ProjectPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectPostWithDefaults() *ProjectPost {
	this := ProjectPost{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectPost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectPost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectPost) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProjectPost) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPost) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectPost) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ProjectPost) SetKey(v string) {
	o.Key = &v
}

// GetIncludeInSnippetByDefault returns the IncludeInSnippetByDefault field value if set, zero value otherwise.
func (o *ProjectPost) GetIncludeInSnippetByDefault() bool {
	if o == nil || o.IncludeInSnippetByDefault == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInSnippetByDefault
}

// GetIncludeInSnippetByDefaultOk returns a tuple with the IncludeInSnippetByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPost) GetIncludeInSnippetByDefaultOk() (*bool, bool) {
	if o == nil || o.IncludeInSnippetByDefault == nil {
		return nil, false
	}
	return o.IncludeInSnippetByDefault, true
}

// HasIncludeInSnippetByDefault returns a boolean if a field has been set.
func (o *ProjectPost) HasIncludeInSnippetByDefault() bool {
	if o != nil && o.IncludeInSnippetByDefault != nil {
		return true
	}

	return false
}

// SetIncludeInSnippetByDefault gets a reference to the given bool and assigns it to the IncludeInSnippetByDefault field.
func (o *ProjectPost) SetIncludeInSnippetByDefault(v bool) {
	o.IncludeInSnippetByDefault = &v
}

// GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field value if set, zero value otherwise.
func (o *ProjectPost) GetDefaultClientSideAvailability() DefaultClientSideAvailabilityPost {
	if o == nil || o.DefaultClientSideAvailability == nil {
		var ret DefaultClientSideAvailabilityPost
		return ret
	}
	return *o.DefaultClientSideAvailability
}

// GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPost) GetDefaultClientSideAvailabilityOk() (*DefaultClientSideAvailabilityPost, bool) {
	if o == nil || o.DefaultClientSideAvailability == nil {
		return nil, false
	}
	return o.DefaultClientSideAvailability, true
}

// HasDefaultClientSideAvailability returns a boolean if a field has been set.
func (o *ProjectPost) HasDefaultClientSideAvailability() bool {
	if o != nil && o.DefaultClientSideAvailability != nil {
		return true
	}

	return false
}

// SetDefaultClientSideAvailability gets a reference to the given DefaultClientSideAvailabilityPost and assigns it to the DefaultClientSideAvailability field.
func (o *ProjectPost) SetDefaultClientSideAvailability(v DefaultClientSideAvailabilityPost) {
	o.DefaultClientSideAvailability = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectPost) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPost) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectPost) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ProjectPost) SetTags(v []string) {
	o.Tags = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ProjectPost) GetEnvironments() []map[string]interface{} {
	if o == nil || o.Environments == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPost) GetEnvironmentsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ProjectPost) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []map[string]interface{} and assigns it to the Environments field.
func (o *ProjectPost) SetEnvironments(v []map[string]interface{}) {
	o.Environments = &v
}

func (o ProjectPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.IncludeInSnippetByDefault != nil {
		toSerialize["includeInSnippetByDefault"] = o.IncludeInSnippetByDefault
	}
	if o.DefaultClientSideAvailability != nil {
		toSerialize["defaultClientSideAvailability"] = o.DefaultClientSideAvailability
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	return json.Marshal(toSerialize)
}

type NullableProjectPost struct {
	value *ProjectPost
	isSet bool
}

func (v NullableProjectPost) Get() *ProjectPost {
	return v.value
}

func (v *NullableProjectPost) Set(val *ProjectPost) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectPost) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectPost(val *ProjectPost) *NullableProjectPost {
	return &NullableProjectPost{value: val, isSet: true}
}

func (v NullableProjectPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


