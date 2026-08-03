# ModelConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**Items** | [**[]ModelConfig**](ModelConfig.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewModelConfigs

`func NewModelConfigs(items []ModelConfig, totalCount int32, ) *ModelConfigs`

NewModelConfigs instantiates a new ModelConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConfigsWithDefaults

`func NewModelConfigsWithDefaults() *ModelConfigs`

NewModelConfigsWithDefaults instantiates a new ModelConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ModelConfigs) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ModelConfigs) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ModelConfigs) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ModelConfigs) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *ModelConfigs) GetItems() []ModelConfig`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelConfigs) GetItemsOk() (*[]ModelConfig, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelConfigs) SetItems(v []ModelConfig)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ModelConfigs) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ModelConfigs) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ModelConfigs) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


