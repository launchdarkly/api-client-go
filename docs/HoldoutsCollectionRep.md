# HoldoutsCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SimpleHoldoutRep**](SimpleHoldoutRep.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TotalCount** | Pointer to **int32** | The total number of holdouts in this project and environment. | [optional] 

## Methods

### NewHoldoutsCollectionRep

`func NewHoldoutsCollectionRep() *HoldoutsCollectionRep`

NewHoldoutsCollectionRep instantiates a new HoldoutsCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutsCollectionRepWithDefaults

`func NewHoldoutsCollectionRepWithDefaults() *HoldoutsCollectionRep`

NewHoldoutsCollectionRepWithDefaults instantiates a new HoldoutsCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HoldoutsCollectionRep) GetItems() []SimpleHoldoutRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HoldoutsCollectionRep) GetItemsOk() (*[]SimpleHoldoutRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HoldoutsCollectionRep) SetItems(v []SimpleHoldoutRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *HoldoutsCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *HoldoutsCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HoldoutsCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HoldoutsCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HoldoutsCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *HoldoutsCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *HoldoutsCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *HoldoutsCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *HoldoutsCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


