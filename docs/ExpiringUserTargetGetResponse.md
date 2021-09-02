# ExpiringUserTargetGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExpiringUserTargetItem**](ExpiringUserTargetItem.md) |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewExpiringUserTargetGetResponse

`func NewExpiringUserTargetGetResponse(items []ExpiringUserTargetItem, ) *ExpiringUserTargetGetResponse`

NewExpiringUserTargetGetResponse instantiates a new ExpiringUserTargetGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringUserTargetGetResponseWithDefaults

`func NewExpiringUserTargetGetResponseWithDefaults() *ExpiringUserTargetGetResponse`

NewExpiringUserTargetGetResponseWithDefaults instantiates a new ExpiringUserTargetGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExpiringUserTargetGetResponse) GetItems() []ExpiringUserTargetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExpiringUserTargetGetResponse) GetItemsOk() (*[]ExpiringUserTargetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExpiringUserTargetGetResponse) SetItems(v []ExpiringUserTargetItem)`

SetItems sets Items field to given value.


### GetLinks

`func (o *ExpiringUserTargetGetResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpiringUserTargetGetResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpiringUserTargetGetResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpiringUserTargetGetResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


