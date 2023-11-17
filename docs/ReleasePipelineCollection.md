# ReleasePipelineCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ReleasePipeline**](ReleasePipeline.md) | An array of release pipelines | 
**TotalCount** | **int32** | Total number of release pipelines | 

## Methods

### NewReleasePipelineCollection

`func NewReleasePipelineCollection(items []ReleasePipeline, totalCount int32, ) *ReleasePipelineCollection`

NewReleasePipelineCollection instantiates a new ReleasePipelineCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePipelineCollectionWithDefaults

`func NewReleasePipelineCollectionWithDefaults() *ReleasePipelineCollection`

NewReleasePipelineCollectionWithDefaults instantiates a new ReleasePipelineCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ReleasePipelineCollection) GetItems() []ReleasePipeline`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ReleasePipelineCollection) GetItemsOk() (*[]ReleasePipeline, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ReleasePipelineCollection) SetItems(v []ReleasePipeline)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ReleasePipelineCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReleasePipelineCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReleasePipelineCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


