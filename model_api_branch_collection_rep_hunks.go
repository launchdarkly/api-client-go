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

// ApiBranchCollectionRepHunks struct for ApiBranchCollectionRepHunks
type ApiBranchCollectionRepHunks struct {
	StartingLineNumber *int32 `json:"startingLineNumber,omitempty"`
	Lines *string `json:"lines,omitempty"`
	ProjKey *string `json:"projKey,omitempty"`
	FlagKey *string `json:"flagKey,omitempty"`
	Aliases *[]string `json:"aliases,omitempty"`
}

// NewApiBranchCollectionRepHunks instantiates a new ApiBranchCollectionRepHunks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBranchCollectionRepHunks() *ApiBranchCollectionRepHunks {
	this := ApiBranchCollectionRepHunks{}
	return &this
}

// NewApiBranchCollectionRepHunksWithDefaults instantiates a new ApiBranchCollectionRepHunks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBranchCollectionRepHunksWithDefaults() *ApiBranchCollectionRepHunks {
	this := ApiBranchCollectionRepHunks{}
	return &this
}

// GetStartingLineNumber returns the StartingLineNumber field value if set, zero value otherwise.
func (o *ApiBranchCollectionRepHunks) GetStartingLineNumber() int32 {
	if o == nil || o.StartingLineNumber == nil {
		var ret int32
		return ret
	}
	return *o.StartingLineNumber
}

// GetStartingLineNumberOk returns a tuple with the StartingLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBranchCollectionRepHunks) GetStartingLineNumberOk() (*int32, bool) {
	if o == nil || o.StartingLineNumber == nil {
		return nil, false
	}
	return o.StartingLineNumber, true
}

// HasStartingLineNumber returns a boolean if a field has been set.
func (o *ApiBranchCollectionRepHunks) HasStartingLineNumber() bool {
	if o != nil && o.StartingLineNumber != nil {
		return true
	}

	return false
}

// SetStartingLineNumber gets a reference to the given int32 and assigns it to the StartingLineNumber field.
func (o *ApiBranchCollectionRepHunks) SetStartingLineNumber(v int32) {
	o.StartingLineNumber = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *ApiBranchCollectionRepHunks) GetLines() string {
	if o == nil || o.Lines == nil {
		var ret string
		return ret
	}
	return *o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBranchCollectionRepHunks) GetLinesOk() (*string, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *ApiBranchCollectionRepHunks) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given string and assigns it to the Lines field.
func (o *ApiBranchCollectionRepHunks) SetLines(v string) {
	o.Lines = &v
}

// GetProjKey returns the ProjKey field value if set, zero value otherwise.
func (o *ApiBranchCollectionRepHunks) GetProjKey() string {
	if o == nil || o.ProjKey == nil {
		var ret string
		return ret
	}
	return *o.ProjKey
}

// GetProjKeyOk returns a tuple with the ProjKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBranchCollectionRepHunks) GetProjKeyOk() (*string, bool) {
	if o == nil || o.ProjKey == nil {
		return nil, false
	}
	return o.ProjKey, true
}

// HasProjKey returns a boolean if a field has been set.
func (o *ApiBranchCollectionRepHunks) HasProjKey() bool {
	if o != nil && o.ProjKey != nil {
		return true
	}

	return false
}

// SetProjKey gets a reference to the given string and assigns it to the ProjKey field.
func (o *ApiBranchCollectionRepHunks) SetProjKey(v string) {
	o.ProjKey = &v
}

// GetFlagKey returns the FlagKey field value if set, zero value otherwise.
func (o *ApiBranchCollectionRepHunks) GetFlagKey() string {
	if o == nil || o.FlagKey == nil {
		var ret string
		return ret
	}
	return *o.FlagKey
}

// GetFlagKeyOk returns a tuple with the FlagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBranchCollectionRepHunks) GetFlagKeyOk() (*string, bool) {
	if o == nil || o.FlagKey == nil {
		return nil, false
	}
	return o.FlagKey, true
}

// HasFlagKey returns a boolean if a field has been set.
func (o *ApiBranchCollectionRepHunks) HasFlagKey() bool {
	if o != nil && o.FlagKey != nil {
		return true
	}

	return false
}

// SetFlagKey gets a reference to the given string and assigns it to the FlagKey field.
func (o *ApiBranchCollectionRepHunks) SetFlagKey(v string) {
	o.FlagKey = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *ApiBranchCollectionRepHunks) GetAliases() []string {
	if o == nil || o.Aliases == nil {
		var ret []string
		return ret
	}
	return *o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBranchCollectionRepHunks) GetAliasesOk() (*[]string, bool) {
	if o == nil || o.Aliases == nil {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *ApiBranchCollectionRepHunks) HasAliases() bool {
	if o != nil && o.Aliases != nil {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *ApiBranchCollectionRepHunks) SetAliases(v []string) {
	o.Aliases = &v
}

func (o ApiBranchCollectionRepHunks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartingLineNumber != nil {
		toSerialize["startingLineNumber"] = o.StartingLineNumber
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	if o.ProjKey != nil {
		toSerialize["projKey"] = o.ProjKey
	}
	if o.FlagKey != nil {
		toSerialize["flagKey"] = o.FlagKey
	}
	if o.Aliases != nil {
		toSerialize["aliases"] = o.Aliases
	}
	return json.Marshal(toSerialize)
}

type NullableApiBranchCollectionRepHunks struct {
	value *ApiBranchCollectionRepHunks
	isSet bool
}

func (v NullableApiBranchCollectionRepHunks) Get() *ApiBranchCollectionRepHunks {
	return v.value
}

func (v *NullableApiBranchCollectionRepHunks) Set(val *ApiBranchCollectionRepHunks) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBranchCollectionRepHunks) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBranchCollectionRepHunks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBranchCollectionRepHunks(val *ApiBranchCollectionRepHunks) *NullableApiBranchCollectionRepHunks {
	return &NullableApiBranchCollectionRepHunks{value: val, isSet: true}
}

func (v NullableApiBranchCollectionRepHunks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBranchCollectionRepHunks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


