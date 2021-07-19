# ExperimentStatsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PValue** | Pointer to **float32** |  | [optional] 
**Chi2** | Pointer to **float32** |  | [optional] 
**WinningVariationIdx** | Pointer to **int32** |  | [optional] 
**MinSampleSizeMet** | Pointer to **bool** |  | [optional] 

## Methods

### NewExperimentStatsRep

`func NewExperimentStatsRep() *ExperimentStatsRep`

NewExperimentStatsRep instantiates a new ExperimentStatsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentStatsRepWithDefaults

`func NewExperimentStatsRepWithDefaults() *ExperimentStatsRep`

NewExperimentStatsRepWithDefaults instantiates a new ExperimentStatsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPValue

`func (o *ExperimentStatsRep) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ExperimentStatsRep) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ExperimentStatsRep) SetPValue(v float32)`

SetPValue sets PValue field to given value.

### HasPValue

`func (o *ExperimentStatsRep) HasPValue() bool`

HasPValue returns a boolean if a field has been set.

### GetChi2

`func (o *ExperimentStatsRep) GetChi2() float32`

GetChi2 returns the Chi2 field if non-nil, zero value otherwise.

### GetChi2Ok

`func (o *ExperimentStatsRep) GetChi2Ok() (*float32, bool)`

GetChi2Ok returns a tuple with the Chi2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChi2

`func (o *ExperimentStatsRep) SetChi2(v float32)`

SetChi2 sets Chi2 field to given value.

### HasChi2

`func (o *ExperimentStatsRep) HasChi2() bool`

HasChi2 returns a boolean if a field has been set.

### GetWinningVariationIdx

`func (o *ExperimentStatsRep) GetWinningVariationIdx() int32`

GetWinningVariationIdx returns the WinningVariationIdx field if non-nil, zero value otherwise.

### GetWinningVariationIdxOk

`func (o *ExperimentStatsRep) GetWinningVariationIdxOk() (*int32, bool)`

GetWinningVariationIdxOk returns a tuple with the WinningVariationIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningVariationIdx

`func (o *ExperimentStatsRep) SetWinningVariationIdx(v int32)`

SetWinningVariationIdx sets WinningVariationIdx field to given value.

### HasWinningVariationIdx

`func (o *ExperimentStatsRep) HasWinningVariationIdx() bool`

HasWinningVariationIdx returns a boolean if a field has been set.

### GetMinSampleSizeMet

`func (o *ExperimentStatsRep) GetMinSampleSizeMet() bool`

GetMinSampleSizeMet returns the MinSampleSizeMet field if non-nil, zero value otherwise.

### GetMinSampleSizeMetOk

`func (o *ExperimentStatsRep) GetMinSampleSizeMetOk() (*bool, bool)`

GetMinSampleSizeMetOk returns a tuple with the MinSampleSizeMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSampleSizeMet

`func (o *ExperimentStatsRep) SetMinSampleSizeMet(v bool)`

SetMinSampleSizeMet sets MinSampleSizeMet field to given value.

### HasMinSampleSizeMet

`func (o *ExperimentStatsRep) HasMinSampleSizeMet() bool`

HasMinSampleSizeMet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


