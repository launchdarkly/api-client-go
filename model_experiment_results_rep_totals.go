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

// ExperimentResultsRepTotals struct for ExperimentResultsRepTotals
type ExperimentResultsRepTotals struct {
	CumulativeValue *float32 `json:"cumulativeValue,omitempty"`
	CumulativeCount *int64 `json:"cumulativeCount,omitempty"`
	CumulativeImpressionCount *int64 `json:"cumulativeImpressionCount,omitempty"`
	CumulativeConversionRate *float32 `json:"cumulativeConversionRate,omitempty"`
	CumulativeConfidenceInterval *ConfidenceIntervalRep `json:"cumulativeConfidenceInterval,omitempty"`
	PValue *float32 `json:"pValue,omitempty"`
	Improvement *float32 `json:"improvement,omitempty"`
	MinSampleSize *int64 `json:"minSampleSize,omitempty"`
}

// NewExperimentResultsRepTotals instantiates a new ExperimentResultsRepTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentResultsRepTotals() *ExperimentResultsRepTotals {
	this := ExperimentResultsRepTotals{}
	return &this
}

// NewExperimentResultsRepTotalsWithDefaults instantiates a new ExperimentResultsRepTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentResultsRepTotalsWithDefaults() *ExperimentResultsRepTotals {
	this := ExperimentResultsRepTotals{}
	return &this
}

// GetCumulativeValue returns the CumulativeValue field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetCumulativeValue() float32 {
	if o == nil || o.CumulativeValue == nil {
		var ret float32
		return ret
	}
	return *o.CumulativeValue
}

// GetCumulativeValueOk returns a tuple with the CumulativeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetCumulativeValueOk() (*float32, bool) {
	if o == nil || o.CumulativeValue == nil {
		return nil, false
	}
	return o.CumulativeValue, true
}

// HasCumulativeValue returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasCumulativeValue() bool {
	if o != nil && o.CumulativeValue != nil {
		return true
	}

	return false
}

// SetCumulativeValue gets a reference to the given float32 and assigns it to the CumulativeValue field.
func (o *ExperimentResultsRepTotals) SetCumulativeValue(v float32) {
	o.CumulativeValue = &v
}

// GetCumulativeCount returns the CumulativeCount field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetCumulativeCount() int64 {
	if o == nil || o.CumulativeCount == nil {
		var ret int64
		return ret
	}
	return *o.CumulativeCount
}

// GetCumulativeCountOk returns a tuple with the CumulativeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetCumulativeCountOk() (*int64, bool) {
	if o == nil || o.CumulativeCount == nil {
		return nil, false
	}
	return o.CumulativeCount, true
}

// HasCumulativeCount returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasCumulativeCount() bool {
	if o != nil && o.CumulativeCount != nil {
		return true
	}

	return false
}

// SetCumulativeCount gets a reference to the given int64 and assigns it to the CumulativeCount field.
func (o *ExperimentResultsRepTotals) SetCumulativeCount(v int64) {
	o.CumulativeCount = &v
}

// GetCumulativeImpressionCount returns the CumulativeImpressionCount field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetCumulativeImpressionCount() int64 {
	if o == nil || o.CumulativeImpressionCount == nil {
		var ret int64
		return ret
	}
	return *o.CumulativeImpressionCount
}

// GetCumulativeImpressionCountOk returns a tuple with the CumulativeImpressionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetCumulativeImpressionCountOk() (*int64, bool) {
	if o == nil || o.CumulativeImpressionCount == nil {
		return nil, false
	}
	return o.CumulativeImpressionCount, true
}

// HasCumulativeImpressionCount returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasCumulativeImpressionCount() bool {
	if o != nil && o.CumulativeImpressionCount != nil {
		return true
	}

	return false
}

// SetCumulativeImpressionCount gets a reference to the given int64 and assigns it to the CumulativeImpressionCount field.
func (o *ExperimentResultsRepTotals) SetCumulativeImpressionCount(v int64) {
	o.CumulativeImpressionCount = &v
}

// GetCumulativeConversionRate returns the CumulativeConversionRate field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetCumulativeConversionRate() float32 {
	if o == nil || o.CumulativeConversionRate == nil {
		var ret float32
		return ret
	}
	return *o.CumulativeConversionRate
}

// GetCumulativeConversionRateOk returns a tuple with the CumulativeConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetCumulativeConversionRateOk() (*float32, bool) {
	if o == nil || o.CumulativeConversionRate == nil {
		return nil, false
	}
	return o.CumulativeConversionRate, true
}

// HasCumulativeConversionRate returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasCumulativeConversionRate() bool {
	if o != nil && o.CumulativeConversionRate != nil {
		return true
	}

	return false
}

// SetCumulativeConversionRate gets a reference to the given float32 and assigns it to the CumulativeConversionRate field.
func (o *ExperimentResultsRepTotals) SetCumulativeConversionRate(v float32) {
	o.CumulativeConversionRate = &v
}

// GetCumulativeConfidenceInterval returns the CumulativeConfidenceInterval field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetCumulativeConfidenceInterval() ConfidenceIntervalRep {
	if o == nil || o.CumulativeConfidenceInterval == nil {
		var ret ConfidenceIntervalRep
		return ret
	}
	return *o.CumulativeConfidenceInterval
}

// GetCumulativeConfidenceIntervalOk returns a tuple with the CumulativeConfidenceInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetCumulativeConfidenceIntervalOk() (*ConfidenceIntervalRep, bool) {
	if o == nil || o.CumulativeConfidenceInterval == nil {
		return nil, false
	}
	return o.CumulativeConfidenceInterval, true
}

// HasCumulativeConfidenceInterval returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasCumulativeConfidenceInterval() bool {
	if o != nil && o.CumulativeConfidenceInterval != nil {
		return true
	}

	return false
}

// SetCumulativeConfidenceInterval gets a reference to the given ConfidenceIntervalRep and assigns it to the CumulativeConfidenceInterval field.
func (o *ExperimentResultsRepTotals) SetCumulativeConfidenceInterval(v ConfidenceIntervalRep) {
	o.CumulativeConfidenceInterval = &v
}

// GetPValue returns the PValue field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetPValue() float32 {
	if o == nil || o.PValue == nil {
		var ret float32
		return ret
	}
	return *o.PValue
}

// GetPValueOk returns a tuple with the PValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetPValueOk() (*float32, bool) {
	if o == nil || o.PValue == nil {
		return nil, false
	}
	return o.PValue, true
}

// HasPValue returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasPValue() bool {
	if o != nil && o.PValue != nil {
		return true
	}

	return false
}

// SetPValue gets a reference to the given float32 and assigns it to the PValue field.
func (o *ExperimentResultsRepTotals) SetPValue(v float32) {
	o.PValue = &v
}

// GetImprovement returns the Improvement field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetImprovement() float32 {
	if o == nil || o.Improvement == nil {
		var ret float32
		return ret
	}
	return *o.Improvement
}

// GetImprovementOk returns a tuple with the Improvement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetImprovementOk() (*float32, bool) {
	if o == nil || o.Improvement == nil {
		return nil, false
	}
	return o.Improvement, true
}

// HasImprovement returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasImprovement() bool {
	if o != nil && o.Improvement != nil {
		return true
	}

	return false
}

// SetImprovement gets a reference to the given float32 and assigns it to the Improvement field.
func (o *ExperimentResultsRepTotals) SetImprovement(v float32) {
	o.Improvement = &v
}

// GetMinSampleSize returns the MinSampleSize field value if set, zero value otherwise.
func (o *ExperimentResultsRepTotals) GetMinSampleSize() int64 {
	if o == nil || o.MinSampleSize == nil {
		var ret int64
		return ret
	}
	return *o.MinSampleSize
}

// GetMinSampleSizeOk returns a tuple with the MinSampleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentResultsRepTotals) GetMinSampleSizeOk() (*int64, bool) {
	if o == nil || o.MinSampleSize == nil {
		return nil, false
	}
	return o.MinSampleSize, true
}

// HasMinSampleSize returns a boolean if a field has been set.
func (o *ExperimentResultsRepTotals) HasMinSampleSize() bool {
	if o != nil && o.MinSampleSize != nil {
		return true
	}

	return false
}

// SetMinSampleSize gets a reference to the given int64 and assigns it to the MinSampleSize field.
func (o *ExperimentResultsRepTotals) SetMinSampleSize(v int64) {
	o.MinSampleSize = &v
}

func (o ExperimentResultsRepTotals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CumulativeValue != nil {
		toSerialize["cumulativeValue"] = o.CumulativeValue
	}
	if o.CumulativeCount != nil {
		toSerialize["cumulativeCount"] = o.CumulativeCount
	}
	if o.CumulativeImpressionCount != nil {
		toSerialize["cumulativeImpressionCount"] = o.CumulativeImpressionCount
	}
	if o.CumulativeConversionRate != nil {
		toSerialize["cumulativeConversionRate"] = o.CumulativeConversionRate
	}
	if o.CumulativeConfidenceInterval != nil {
		toSerialize["cumulativeConfidenceInterval"] = o.CumulativeConfidenceInterval
	}
	if o.PValue != nil {
		toSerialize["pValue"] = o.PValue
	}
	if o.Improvement != nil {
		toSerialize["improvement"] = o.Improvement
	}
	if o.MinSampleSize != nil {
		toSerialize["minSampleSize"] = o.MinSampleSize
	}
	return json.Marshal(toSerialize)
}

type NullableExperimentResultsRepTotals struct {
	value *ExperimentResultsRepTotals
	isSet bool
}

func (v NullableExperimentResultsRepTotals) Get() *ExperimentResultsRepTotals {
	return v.value
}

func (v *NullableExperimentResultsRepTotals) Set(val *ExperimentResultsRepTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentResultsRepTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentResultsRepTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentResultsRepTotals(val *ExperimentResultsRepTotals) *NullableExperimentResultsRepTotals {
	return &NullableExperimentResultsRepTotals{value: val, isSet: true}
}

func (v NullableExperimentResultsRepTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentResultsRepTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


