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

// ExperimentTimeSeriesVariationSlice struct for ExperimentTimeSeriesVariationSlice
type ExperimentTimeSeriesVariationSlice struct {
	Value *float32 `json:"value,omitempty"`
	Count *int64 `json:"count,omitempty"`
	CumulativeValue *float32 `json:"cumulativeValue,omitempty"`
	CumulativeCount *int64 `json:"cumulativeCount,omitempty"`
	ConversionRate *float32 `json:"conversionRate,omitempty"`
	CumulativeConversionRate *float32 `json:"cumulativeConversionRate,omitempty"`
	ConfidenceInterval *ConfidenceIntervalRep `json:"confidenceInterval,omitempty"`
	CumulativeConfidenceInterval *ConfidenceIntervalRep `json:"cumulativeConfidenceInterval,omitempty"`
}

// NewExperimentTimeSeriesVariationSlice instantiates a new ExperimentTimeSeriesVariationSlice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentTimeSeriesVariationSlice() *ExperimentTimeSeriesVariationSlice {
	this := ExperimentTimeSeriesVariationSlice{}
	return &this
}

// NewExperimentTimeSeriesVariationSliceWithDefaults instantiates a new ExperimentTimeSeriesVariationSlice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentTimeSeriesVariationSliceWithDefaults() *ExperimentTimeSeriesVariationSlice {
	this := ExperimentTimeSeriesVariationSlice{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ExperimentTimeSeriesVariationSlice) SetValue(v float32) {
	o.Value = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ExperimentTimeSeriesVariationSlice) SetCount(v int64) {
	o.Count = &v
}

// GetCumulativeValue returns the CumulativeValue field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeValue() float32 {
	if o == nil || o.CumulativeValue == nil {
		var ret float32
		return ret
	}
	return *o.CumulativeValue
}

// GetCumulativeValueOk returns a tuple with the CumulativeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeValueOk() (*float32, bool) {
	if o == nil || o.CumulativeValue == nil {
		return nil, false
	}
	return o.CumulativeValue, true
}

// HasCumulativeValue returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeValue() bool {
	if o != nil && o.CumulativeValue != nil {
		return true
	}

	return false
}

// SetCumulativeValue gets a reference to the given float32 and assigns it to the CumulativeValue field.
func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeValue(v float32) {
	o.CumulativeValue = &v
}

// GetCumulativeCount returns the CumulativeCount field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeCount() int64 {
	if o == nil || o.CumulativeCount == nil {
		var ret int64
		return ret
	}
	return *o.CumulativeCount
}

// GetCumulativeCountOk returns a tuple with the CumulativeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeCountOk() (*int64, bool) {
	if o == nil || o.CumulativeCount == nil {
		return nil, false
	}
	return o.CumulativeCount, true
}

// HasCumulativeCount returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeCount() bool {
	if o != nil && o.CumulativeCount != nil {
		return true
	}

	return false
}

// SetCumulativeCount gets a reference to the given int64 and assigns it to the CumulativeCount field.
func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeCount(v int64) {
	o.CumulativeCount = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetConversionRate() float32 {
	if o == nil || o.ConversionRate == nil {
		var ret float32
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetConversionRateOk() (*float32, bool) {
	if o == nil || o.ConversionRate == nil {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasConversionRate() bool {
	if o != nil && o.ConversionRate != nil {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given float32 and assigns it to the ConversionRate field.
func (o *ExperimentTimeSeriesVariationSlice) SetConversionRate(v float32) {
	o.ConversionRate = &v
}

// GetCumulativeConversionRate returns the CumulativeConversionRate field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConversionRate() float32 {
	if o == nil || o.CumulativeConversionRate == nil {
		var ret float32
		return ret
	}
	return *o.CumulativeConversionRate
}

// GetCumulativeConversionRateOk returns a tuple with the CumulativeConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConversionRateOk() (*float32, bool) {
	if o == nil || o.CumulativeConversionRate == nil {
		return nil, false
	}
	return o.CumulativeConversionRate, true
}

// HasCumulativeConversionRate returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeConversionRate() bool {
	if o != nil && o.CumulativeConversionRate != nil {
		return true
	}

	return false
}

// SetCumulativeConversionRate gets a reference to the given float32 and assigns it to the CumulativeConversionRate field.
func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeConversionRate(v float32) {
	o.CumulativeConversionRate = &v
}

// GetConfidenceInterval returns the ConfidenceInterval field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetConfidenceInterval() ConfidenceIntervalRep {
	if o == nil || o.ConfidenceInterval == nil {
		var ret ConfidenceIntervalRep
		return ret
	}
	return *o.ConfidenceInterval
}

// GetConfidenceIntervalOk returns a tuple with the ConfidenceInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetConfidenceIntervalOk() (*ConfidenceIntervalRep, bool) {
	if o == nil || o.ConfidenceInterval == nil {
		return nil, false
	}
	return o.ConfidenceInterval, true
}

// HasConfidenceInterval returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasConfidenceInterval() bool {
	if o != nil && o.ConfidenceInterval != nil {
		return true
	}

	return false
}

// SetConfidenceInterval gets a reference to the given ConfidenceIntervalRep and assigns it to the ConfidenceInterval field.
func (o *ExperimentTimeSeriesVariationSlice) SetConfidenceInterval(v ConfidenceIntervalRep) {
	o.ConfidenceInterval = &v
}

// GetCumulativeConfidenceInterval returns the CumulativeConfidenceInterval field value if set, zero value otherwise.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConfidenceInterval() ConfidenceIntervalRep {
	if o == nil || o.CumulativeConfidenceInterval == nil {
		var ret ConfidenceIntervalRep
		return ret
	}
	return *o.CumulativeConfidenceInterval
}

// GetCumulativeConfidenceIntervalOk returns a tuple with the CumulativeConfidenceInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConfidenceIntervalOk() (*ConfidenceIntervalRep, bool) {
	if o == nil || o.CumulativeConfidenceInterval == nil {
		return nil, false
	}
	return o.CumulativeConfidenceInterval, true
}

// HasCumulativeConfidenceInterval returns a boolean if a field has been set.
func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeConfidenceInterval() bool {
	if o != nil && o.CumulativeConfidenceInterval != nil {
		return true
	}

	return false
}

// SetCumulativeConfidenceInterval gets a reference to the given ConfidenceIntervalRep and assigns it to the CumulativeConfidenceInterval field.
func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeConfidenceInterval(v ConfidenceIntervalRep) {
	o.CumulativeConfidenceInterval = &v
}

func (o ExperimentTimeSeriesVariationSlice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.CumulativeValue != nil {
		toSerialize["cumulativeValue"] = o.CumulativeValue
	}
	if o.CumulativeCount != nil {
		toSerialize["cumulativeCount"] = o.CumulativeCount
	}
	if o.ConversionRate != nil {
		toSerialize["conversionRate"] = o.ConversionRate
	}
	if o.CumulativeConversionRate != nil {
		toSerialize["cumulativeConversionRate"] = o.CumulativeConversionRate
	}
	if o.ConfidenceInterval != nil {
		toSerialize["confidenceInterval"] = o.ConfidenceInterval
	}
	if o.CumulativeConfidenceInterval != nil {
		toSerialize["cumulativeConfidenceInterval"] = o.CumulativeConfidenceInterval
	}
	return json.Marshal(toSerialize)
}

type NullableExperimentTimeSeriesVariationSlice struct {
	value *ExperimentTimeSeriesVariationSlice
	isSet bool
}

func (v NullableExperimentTimeSeriesVariationSlice) Get() *ExperimentTimeSeriesVariationSlice {
	return v.value
}

func (v *NullableExperimentTimeSeriesVariationSlice) Set(val *ExperimentTimeSeriesVariationSlice) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentTimeSeriesVariationSlice) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentTimeSeriesVariationSlice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentTimeSeriesVariationSlice(val *ExperimentTimeSeriesVariationSlice) *NullableExperimentTimeSeriesVariationSlice {
	return &NullableExperimentTimeSeriesVariationSlice{value: val, isSet: true}
}

func (v NullableExperimentTimeSeriesVariationSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentTimeSeriesVariationSlice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


