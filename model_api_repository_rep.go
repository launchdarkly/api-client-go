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

// ApiRepositoryRep struct for ApiRepositoryRep
type ApiRepositoryRep struct {
	Name *string `json:"name,omitempty"`
	SourceLink *string `json:"sourceLink,omitempty"`
	CommitUrlTemplate *string `json:"commitUrlTemplate,omitempty"`
	HunkUrlTemplate *string `json:"hunkUrlTemplate,omitempty"`
	Type *string `json:"type,omitempty"`
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Version *int32 `json:"version,omitempty"`
	Branches *[]ApiBranchRep `json:"branches,omitempty"`
	Links *map[string]interface{} `json:"_links,omitempty"`
	Access *AccessRep `json:"_access,omitempty"`
}

// NewApiRepositoryRep instantiates a new ApiRepositoryRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryRep() *ApiRepositoryRep {
	this := ApiRepositoryRep{}
	return &this
}

// NewApiRepositoryRepWithDefaults instantiates a new ApiRepositoryRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryRepWithDefaults() *ApiRepositoryRep {
	this := ApiRepositoryRep{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiRepositoryRep) SetName(v string) {
	o.Name = &v
}

// GetSourceLink returns the SourceLink field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetSourceLink() string {
	if o == nil || o.SourceLink == nil {
		var ret string
		return ret
	}
	return *o.SourceLink
}

// GetSourceLinkOk returns a tuple with the SourceLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetSourceLinkOk() (*string, bool) {
	if o == nil || o.SourceLink == nil {
		return nil, false
	}
	return o.SourceLink, true
}

// HasSourceLink returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasSourceLink() bool {
	if o != nil && o.SourceLink != nil {
		return true
	}

	return false
}

// SetSourceLink gets a reference to the given string and assigns it to the SourceLink field.
func (o *ApiRepositoryRep) SetSourceLink(v string) {
	o.SourceLink = &v
}

// GetCommitUrlTemplate returns the CommitUrlTemplate field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetCommitUrlTemplate() string {
	if o == nil || o.CommitUrlTemplate == nil {
		var ret string
		return ret
	}
	return *o.CommitUrlTemplate
}

// GetCommitUrlTemplateOk returns a tuple with the CommitUrlTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetCommitUrlTemplateOk() (*string, bool) {
	if o == nil || o.CommitUrlTemplate == nil {
		return nil, false
	}
	return o.CommitUrlTemplate, true
}

// HasCommitUrlTemplate returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasCommitUrlTemplate() bool {
	if o != nil && o.CommitUrlTemplate != nil {
		return true
	}

	return false
}

// SetCommitUrlTemplate gets a reference to the given string and assigns it to the CommitUrlTemplate field.
func (o *ApiRepositoryRep) SetCommitUrlTemplate(v string) {
	o.CommitUrlTemplate = &v
}

// GetHunkUrlTemplate returns the HunkUrlTemplate field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetHunkUrlTemplate() string {
	if o == nil || o.HunkUrlTemplate == nil {
		var ret string
		return ret
	}
	return *o.HunkUrlTemplate
}

// GetHunkUrlTemplateOk returns a tuple with the HunkUrlTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetHunkUrlTemplateOk() (*string, bool) {
	if o == nil || o.HunkUrlTemplate == nil {
		return nil, false
	}
	return o.HunkUrlTemplate, true
}

// HasHunkUrlTemplate returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasHunkUrlTemplate() bool {
	if o != nil && o.HunkUrlTemplate != nil {
		return true
	}

	return false
}

// SetHunkUrlTemplate gets a reference to the given string and assigns it to the HunkUrlTemplate field.
func (o *ApiRepositoryRep) SetHunkUrlTemplate(v string) {
	o.HunkUrlTemplate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiRepositoryRep) SetType(v string) {
	o.Type = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetDefaultBranch() string {
	if o == nil || o.DefaultBranch == nil {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetDefaultBranchOk() (*string, bool) {
	if o == nil || o.DefaultBranch == nil {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch != nil {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *ApiRepositoryRep) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiRepositoryRep) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ApiRepositoryRep) SetVersion(v int32) {
	o.Version = &v
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetBranches() []ApiBranchRep {
	if o == nil || o.Branches == nil {
		var ret []ApiBranchRep
		return ret
	}
	return *o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetBranchesOk() (*[]ApiBranchRep, bool) {
	if o == nil || o.Branches == nil {
		return nil, false
	}
	return o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasBranches() bool {
	if o != nil && o.Branches != nil {
		return true
	}

	return false
}

// SetBranches gets a reference to the given []ApiBranchRep and assigns it to the Branches field.
func (o *ApiRepositoryRep) SetBranches(v []ApiBranchRep) {
	o.Branches = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetLinksOk() (*map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *ApiRepositoryRep) SetLinks(v map[string]interface{}) {
	o.Links = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ApiRepositoryRep) GetAccess() AccessRep {
	if o == nil || o.Access == nil {
		var ret AccessRep
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryRep) GetAccessOk() (*AccessRep, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ApiRepositoryRep) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AccessRep and assigns it to the Access field.
func (o *ApiRepositoryRep) SetAccess(v AccessRep) {
	o.Access = &v
}

func (o ApiRepositoryRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SourceLink != nil {
		toSerialize["sourceLink"] = o.SourceLink
	}
	if o.CommitUrlTemplate != nil {
		toSerialize["commitUrlTemplate"] = o.CommitUrlTemplate
	}
	if o.HunkUrlTemplate != nil {
		toSerialize["hunkUrlTemplate"] = o.HunkUrlTemplate
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DefaultBranch != nil {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Branches != nil {
		toSerialize["branches"] = o.Branches
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Access != nil {
		toSerialize["_access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableApiRepositoryRep struct {
	value *ApiRepositoryRep
	isSet bool
}

func (v NullableApiRepositoryRep) Get() *ApiRepositoryRep {
	return v.value
}

func (v *NullableApiRepositoryRep) Set(val *ApiRepositoryRep) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryRep) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryRep(val *ApiRepositoryRep) *NullableApiRepositoryRep {
	return &NullableApiRepositoryRep{value: val, isSet: true}
}

func (v NullableApiRepositoryRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


