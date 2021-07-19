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

// DestinationCollectionRep struct for DestinationCollectionRep
type DestinationCollectionRep struct {
	Links *map[string]InlineResponse200 `json:"_links,omitempty"`
	Items *[]DestinationRep `json:"items,omitempty"`
}

// NewDestinationCollectionRep instantiates a new DestinationCollectionRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationCollectionRep() *DestinationCollectionRep {
	this := DestinationCollectionRep{}
	return &this
}

// NewDestinationCollectionRepWithDefaults instantiates a new DestinationCollectionRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationCollectionRepWithDefaults() *DestinationCollectionRep {
	this := DestinationCollectionRep{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DestinationCollectionRep) GetLinks() map[string]InlineResponse200 {
	if o == nil || o.Links == nil {
		var ret map[string]InlineResponse200
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationCollectionRep) GetLinksOk() (*map[string]InlineResponse200, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DestinationCollectionRep) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]InlineResponse200 and assigns it to the Links field.
func (o *DestinationCollectionRep) SetLinks(v map[string]InlineResponse200) {
	o.Links = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DestinationCollectionRep) GetItems() []DestinationRep {
	if o == nil || o.Items == nil {
		var ret []DestinationRep
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationCollectionRep) GetItemsOk() (*[]DestinationRep, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DestinationCollectionRep) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DestinationRep and assigns it to the Items field.
func (o *DestinationCollectionRep) SetItems(v []DestinationRep) {
	o.Items = &v
}

func (o DestinationCollectionRep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationCollectionRep struct {
	value *DestinationCollectionRep
	isSet bool
}

func (v NullableDestinationCollectionRep) Get() *DestinationCollectionRep {
	return v.value
}

func (v *NullableDestinationCollectionRep) Set(val *DestinationCollectionRep) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationCollectionRep) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationCollectionRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationCollectionRep(val *DestinationCollectionRep) *NullableDestinationCollectionRep {
	return &NullableDestinationCollectionRep{value: val, isSet: true}
}

func (v NullableDestinationCollectionRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationCollectionRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


