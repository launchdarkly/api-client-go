# AIConfigVariationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AIConfigVariation**](AIConfigVariation.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAIConfigVariationsResponse

`func NewAIConfigVariationsResponse(items []AIConfigVariation, totalCount int32, ) *AIConfigVariationsResponse`

NewAIConfigVariationsResponse instantiates a new AIConfigVariationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigVariationsResponseWithDefaults

`func NewAIConfigVariationsResponseWithDefaults() *AIConfigVariationsResponse`

NewAIConfigVariationsResponseWithDefaults instantiates a new AIConfigVariationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AIConfigVariationsResponse) GetItems() []AIConfigVariation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AIConfigVariationsResponse) GetItemsOk() (*[]AIConfigVariation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AIConfigVariationsResponse) SetItems(v []AIConfigVariation)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AIConfigVariationsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AIConfigVariationsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AIConfigVariationsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


