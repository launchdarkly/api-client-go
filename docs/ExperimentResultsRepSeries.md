# ExperimentResultsRepSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int64** |  | [optional] 
**VariationData** | Pointer to [**[]ExperimentTimeSeriesVariationSlice**](ExperimentTimeSeriesVariationSlice.md) |  | [optional] 

## Methods

### NewExperimentResultsRepSeries

`func NewExperimentResultsRepSeries() *ExperimentResultsRepSeries`

NewExperimentResultsRepSeries instantiates a new ExperimentResultsRepSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentResultsRepSeriesWithDefaults

`func NewExperimentResultsRepSeriesWithDefaults() *ExperimentResultsRepSeries`

NewExperimentResultsRepSeriesWithDefaults instantiates a new ExperimentResultsRepSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ExperimentResultsRepSeries) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ExperimentResultsRepSeries) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ExperimentResultsRepSeries) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ExperimentResultsRepSeries) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetVariationData

`func (o *ExperimentResultsRepSeries) GetVariationData() []ExperimentTimeSeriesVariationSlice`

GetVariationData returns the VariationData field if non-nil, zero value otherwise.

### GetVariationDataOk

`func (o *ExperimentResultsRepSeries) GetVariationDataOk() (*[]ExperimentTimeSeriesVariationSlice, bool)`

GetVariationDataOk returns a tuple with the VariationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationData

`func (o *ExperimentResultsRepSeries) SetVariationData(v []ExperimentTimeSeriesVariationSlice)`

SetVariationData sets VariationData field to given value.

### HasVariationData

`func (o *ExperimentResultsRepSeries) HasVariationData() bool`

HasVariationData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


