# ExpiringTargetGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExpiringTarget**](ExpiringTarget.md) | A list of expiring targets | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewExpiringTargetGetResponse

`func NewExpiringTargetGetResponse(items []ExpiringTarget, ) *ExpiringTargetGetResponse`

NewExpiringTargetGetResponse instantiates a new ExpiringTargetGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringTargetGetResponseWithDefaults

`func NewExpiringTargetGetResponseWithDefaults() *ExpiringTargetGetResponse`

NewExpiringTargetGetResponseWithDefaults instantiates a new ExpiringTargetGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExpiringTargetGetResponse) GetItems() []ExpiringTarget`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExpiringTargetGetResponse) GetItemsOk() (*[]ExpiringTarget, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExpiringTargetGetResponse) SetItems(v []ExpiringTarget)`

SetItems sets Items field to given value.


### GetLinks

`func (o *ExpiringTargetGetResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpiringTargetGetResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpiringTargetGetResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpiringTargetGetResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


