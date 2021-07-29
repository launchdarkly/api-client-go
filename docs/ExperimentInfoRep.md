# ExperimentInfoRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaselineIdx** | **int32** |  | 
**Items** | [**[]ExperimentInfoRepItems**](ExperimentInfoRepItems.md) |  | 

## Methods

### NewExperimentInfoRep

`func NewExperimentInfoRep(baselineIdx int32, items []ExperimentInfoRepItems, ) *ExperimentInfoRep`

NewExperimentInfoRep instantiates a new ExperimentInfoRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentInfoRepWithDefaults

`func NewExperimentInfoRepWithDefaults() *ExperimentInfoRep`

NewExperimentInfoRepWithDefaults instantiates a new ExperimentInfoRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaselineIdx

`func (o *ExperimentInfoRep) GetBaselineIdx() int32`

GetBaselineIdx returns the BaselineIdx field if non-nil, zero value otherwise.

### GetBaselineIdxOk

`func (o *ExperimentInfoRep) GetBaselineIdxOk() (*int32, bool)`

GetBaselineIdxOk returns a tuple with the BaselineIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineIdx

`func (o *ExperimentInfoRep) SetBaselineIdx(v int32)`

SetBaselineIdx sets BaselineIdx field to given value.


### GetItems

`func (o *ExperimentInfoRep) GetItems() []ExperimentInfoRepItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExperimentInfoRep) GetItemsOk() (*[]ExperimentInfoRepItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExperimentInfoRep) SetItems(v []ExperimentInfoRepItems)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


