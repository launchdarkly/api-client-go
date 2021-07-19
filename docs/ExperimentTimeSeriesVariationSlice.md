# ExperimentTimeSeriesVariationSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**CumulativeValue** | Pointer to **float32** |  | [optional] 
**CumulativeCount** | Pointer to **int64** |  | [optional] 
**ConversionRate** | Pointer to **float32** |  | [optional] 
**CumulativeConversionRate** | Pointer to **float32** |  | [optional] 
**ConfidenceInterval** | Pointer to [**ConfidenceIntervalRep**](ConfidenceIntervalRep.md) |  | [optional] 
**CumulativeConfidenceInterval** | Pointer to [**ConfidenceIntervalRep**](ConfidenceIntervalRep.md) |  | [optional] 

## Methods

### NewExperimentTimeSeriesVariationSlice

`func NewExperimentTimeSeriesVariationSlice() *ExperimentTimeSeriesVariationSlice`

NewExperimentTimeSeriesVariationSlice instantiates a new ExperimentTimeSeriesVariationSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentTimeSeriesVariationSliceWithDefaults

`func NewExperimentTimeSeriesVariationSliceWithDefaults() *ExperimentTimeSeriesVariationSlice`

NewExperimentTimeSeriesVariationSliceWithDefaults instantiates a new ExperimentTimeSeriesVariationSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ExperimentTimeSeriesVariationSlice) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExperimentTimeSeriesVariationSlice) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExperimentTimeSeriesVariationSlice) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ExperimentTimeSeriesVariationSlice) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCount

`func (o *ExperimentTimeSeriesVariationSlice) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ExperimentTimeSeriesVariationSlice) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ExperimentTimeSeriesVariationSlice) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ExperimentTimeSeriesVariationSlice) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCumulativeValue

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeValue() float32`

GetCumulativeValue returns the CumulativeValue field if non-nil, zero value otherwise.

### GetCumulativeValueOk

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeValueOk() (*float32, bool)`

GetCumulativeValueOk returns a tuple with the CumulativeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeValue

`func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeValue(v float32)`

SetCumulativeValue sets CumulativeValue field to given value.

### HasCumulativeValue

`func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeValue() bool`

HasCumulativeValue returns a boolean if a field has been set.

### GetCumulativeCount

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeCount() int64`

GetCumulativeCount returns the CumulativeCount field if non-nil, zero value otherwise.

### GetCumulativeCountOk

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeCountOk() (*int64, bool)`

GetCumulativeCountOk returns a tuple with the CumulativeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCount

`func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeCount(v int64)`

SetCumulativeCount sets CumulativeCount field to given value.

### HasCumulativeCount

`func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeCount() bool`

HasCumulativeCount returns a boolean if a field has been set.

### GetConversionRate

`func (o *ExperimentTimeSeriesVariationSlice) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *ExperimentTimeSeriesVariationSlice) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *ExperimentTimeSeriesVariationSlice) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *ExperimentTimeSeriesVariationSlice) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCumulativeConversionRate

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConversionRate() float32`

GetCumulativeConversionRate returns the CumulativeConversionRate field if non-nil, zero value otherwise.

### GetCumulativeConversionRateOk

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConversionRateOk() (*float32, bool)`

GetCumulativeConversionRateOk returns a tuple with the CumulativeConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeConversionRate

`func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeConversionRate(v float32)`

SetCumulativeConversionRate sets CumulativeConversionRate field to given value.

### HasCumulativeConversionRate

`func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeConversionRate() bool`

HasCumulativeConversionRate returns a boolean if a field has been set.

### GetConfidenceInterval

`func (o *ExperimentTimeSeriesVariationSlice) GetConfidenceInterval() ConfidenceIntervalRep`

GetConfidenceInterval returns the ConfidenceInterval field if non-nil, zero value otherwise.

### GetConfidenceIntervalOk

`func (o *ExperimentTimeSeriesVariationSlice) GetConfidenceIntervalOk() (*ConfidenceIntervalRep, bool)`

GetConfidenceIntervalOk returns a tuple with the ConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceInterval

`func (o *ExperimentTimeSeriesVariationSlice) SetConfidenceInterval(v ConfidenceIntervalRep)`

SetConfidenceInterval sets ConfidenceInterval field to given value.

### HasConfidenceInterval

`func (o *ExperimentTimeSeriesVariationSlice) HasConfidenceInterval() bool`

HasConfidenceInterval returns a boolean if a field has been set.

### GetCumulativeConfidenceInterval

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConfidenceInterval() ConfidenceIntervalRep`

GetCumulativeConfidenceInterval returns the CumulativeConfidenceInterval field if non-nil, zero value otherwise.

### GetCumulativeConfidenceIntervalOk

`func (o *ExperimentTimeSeriesVariationSlice) GetCumulativeConfidenceIntervalOk() (*ConfidenceIntervalRep, bool)`

GetCumulativeConfidenceIntervalOk returns a tuple with the CumulativeConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeConfidenceInterval

`func (o *ExperimentTimeSeriesVariationSlice) SetCumulativeConfidenceInterval(v ConfidenceIntervalRep)`

SetCumulativeConfidenceInterval sets CumulativeConfidenceInterval field to given value.

### HasCumulativeConfidenceInterval

`func (o *ExperimentTimeSeriesVariationSlice) HasCumulativeConfidenceInterval() bool`

HasCumulativeConfidenceInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


