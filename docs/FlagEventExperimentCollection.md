# FlagEventExperimentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of experiments | 
**Items** | [**[]FlagEventExperiment**](FlagEventExperiment.md) | A list of experiments | 

## Methods

### NewFlagEventExperimentCollection

`func NewFlagEventExperimentCollection(totalCount int32, items []FlagEventExperiment, ) *FlagEventExperimentCollection`

NewFlagEventExperimentCollection instantiates a new FlagEventExperimentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagEventExperimentCollectionWithDefaults

`func NewFlagEventExperimentCollectionWithDefaults() *FlagEventExperimentCollection`

NewFlagEventExperimentCollectionWithDefaults instantiates a new FlagEventExperimentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *FlagEventExperimentCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FlagEventExperimentCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FlagEventExperimentCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *FlagEventExperimentCollection) GetItems() []FlagEventExperiment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlagEventExperimentCollection) GetItemsOk() (*[]FlagEventExperiment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlagEventExperimentCollection) SetItems(v []FlagEventExperiment)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


