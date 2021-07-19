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

// ProjectRep struct for ProjectRep
type ProjectRep struct {
	Links *map[string]InlineResponse200 `json:"_links,omitempty"`
	Id *string `json:"_id,omitempty"`
	Key *string `json:"key,omitempty"`
	IncludeInSnippetByDefault *bool `json:"includeInSnippetByDefault,omitempty"`
	DefaultClientSideAvailability *AccountsClientSideAvailability `json:"defaultClientSideAvailability,omitempty"`
	Name *string `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Environments *[]EnvironmentRep `json:"environments,omitempty"`
}

// NewProjectRep instantiates a new ProjectRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRep() *ProjectRep {
	this := ProjectRep{}
	return &this
}

// NewProjectRepWithDefaults instantiates a new ProjectRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRepWithDefaults() *ProjectRep {
	this := ProjectRep{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProjectRep) GetLinks() map[string]InlineResponse200 {
	if o == nil || o.Links == nil {
		var ret map[string]InlineResponse200
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetLinksOk() (*map[string]InlineResponse200, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProjectRep) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]InlineResponse200 and assigns it to the Links field.
func (o *ProjectRep) SetLinks(v map[string]InlineResponse200) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectRep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectRep) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectRep) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProjectRep) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectRep) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ProjectRep) SetKey(v string) {
	o.Key = &v
}

// GetIncludeInSnippetByDefault returns the IncludeInSnippetByDefault field value if set, zero value otherwise.
func (o *ProjectRep) GetIncludeInSnippetByDefault() bool {
	if o == nil || o.IncludeInSnippetByDefault == nil {
		var ret bool
		return ret
	}
	return *o.IncludeInSnippetByDefault
}

// GetIncludeInSnippetByDefaultOk returns a tuple with the IncludeInSnippetByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetIncludeInSnippetByDefaultOk() (*bool, bool) {
	if o == nil || o.IncludeInSnippetByDefault == nil {
		return nil, false
	}
	return o.IncludeInSnippetByDefault, true
}

// HasIncludeInSnippetByDefault returns a boolean if a field has been set.
func (o *ProjectRep) HasIncludeInSnippetByDefault() bool {
	if o != nil && o.IncludeInSnippetByDefault != nil {
		return true
	}

	return false
}

// SetIncludeInSnippetByDefault gets a reference to the given bool and assigns it to the IncludeInSnippetByDefault field.
func (o *ProjectRep) SetIncludeInSnippetByDefault(v bool) {
	o.IncludeInSnippetByDefault = &v
}

// GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field value if set, zero value otherwise.
func (o *ProjectRep) GetDefaultClientSideAvailability() AccountsClientSideAvailability {
	if o == nil || o.DefaultClientSideAvailability == nil {
		var ret AccountsClientSideAvailability
		return ret
	}
	return *o.DefaultClientSideAvailability
}

// GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetDefaultClientSideAvailabilityOk() (*AccountsClientSideAvailability, bool) {
	if o == nil || o.DefaultClientSideAvailability == nil {
		return nil, false
	}
	return o.DefaultClientSideAvailability, true
}

// HasDefaultClientSideAvailability returns a boolean if a field has been set.
func (o *ProjectRep) HasDefaultClientSideAvailability() bool {
	if o != nil && o.DefaultClientSideAvailability != nil {
		return true
	}

	return false
}

// SetDefaultClientSideAvailability gets a reference to the given AccountsClientSideAvailability and assigns it to the DefaultClientSideAvailability field.
func (o *ProjectRep) SetDefaultClientSideAvailability(v AccountsClientSideAvailability) {
	o.DefaultClientSideAvailability = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectRep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectRep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectRep) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProjectRep) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProjectRep) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ProjectRep) SetTags(v []string) {
	o.Tags = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ProjectRep) GetEnvironments() []EnvironmentRep {
	if o == nil || o.Environments == nil {
		var ret []EnvironmentRep
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRep) GetEnvironmentsOk() (*[]EnvironmentRep, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ProjectRep) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []EnvironmentRep and assigns it to the Environments field.
func (o *ProjectRep) SetEnvironments(v []EnvironmentRep) {
	o.Environments = &v
}

func (o ProjectRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	return json.Marshal(toSerialize)
}

type NullableProjectRep struct {
	value *ProjectRep
	isSet bool
}

func (v NullableProjectRep) Get() *ProjectRep {
	return v.value
}

func (v *NullableProjectRep) Set(val *ProjectRep) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRep) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRep(val *ProjectRep) *NullableProjectRep {
	return &NullableProjectRep{value: val, isSet: true}
}

func (v NullableProjectRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


