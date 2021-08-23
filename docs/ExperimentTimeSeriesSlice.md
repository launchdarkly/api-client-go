# ExperimentTimeSeriesSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int64** |  | [optional] 
**VariationData** | Pointer to [**[]ExperimentTimeSeriesVariationSlice**](ExperimentTimeSeriesVariationSlice.md) |  | [optional] 

## Methods

### NewExperimentTimeSeriesSlice

`func NewExperimentTimeSeriesSlice() *ExperimentTimeSeriesSlice`

NewExperimentTimeSeriesSlice instantiates a new ExperimentTimeSeriesSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentTimeSeriesSliceWithDefaults

`func NewExperimentTimeSeriesSliceWithDefaults() *ExperimentTimeSeriesSlice`

NewExperimentTimeSeriesSliceWithDefaults instantiates a new ExperimentTimeSeriesSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ExperimentTimeSeriesSlice) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ExperimentTimeSeriesSlice) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ExperimentTimeSeriesSlice) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ExperimentTimeSeriesSlice) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetVariationData

`func (o *ExperimentTimeSeriesSlice) GetVariationData() []ExperimentTimeSeriesVariationSlice`

GetVariationData returns the VariationData field if non-nil, zero value otherwise.

### GetVariationDataOk

`func (o *ExperimentTimeSeriesSlice) GetVariationDataOk() (*[]ExperimentTimeSeriesVariationSlice, bool)`

GetVariationDataOk returns a tuple with the VariationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationData

`func (o *ExperimentTimeSeriesSlice) SetVariationData(v []ExperimentTimeSeriesVariationSlice)`

SetVariationData sets VariationData field to given value.

### HasVariationData

`func (o *ExperimentTimeSeriesSlice) HasVariationData() bool`

HasVariationData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


