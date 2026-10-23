# AdaptiveTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AdaptiveTriggersPaginatedLinks**](AdaptiveTriggersPaginatedLinks.md) |  | [optional] 
**Items** | [**[]AdaptiveTrigger**](AdaptiveTrigger.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAdaptiveTriggers

`func NewAdaptiveTriggers(items []AdaptiveTrigger, totalCount int32, ) *AdaptiveTriggers`

NewAdaptiveTriggers instantiates a new AdaptiveTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdaptiveTriggersWithDefaults

`func NewAdaptiveTriggersWithDefaults() *AdaptiveTriggers`

NewAdaptiveTriggersWithDefaults instantiates a new AdaptiveTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AdaptiveTriggers) GetLinks() AdaptiveTriggersPaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AdaptiveTriggers) GetLinksOk() (*AdaptiveTriggersPaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AdaptiveTriggers) SetLinks(v AdaptiveTriggersPaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AdaptiveTriggers) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *AdaptiveTriggers) GetItems() []AdaptiveTrigger`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdaptiveTriggers) GetItemsOk() (*[]AdaptiveTrigger, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdaptiveTriggers) SetItems(v []AdaptiveTrigger)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AdaptiveTriggers) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AdaptiveTriggers) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AdaptiveTriggers) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


