# UserListRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]UserListRepItems**](UserListRepItems.md) |  | [optional] 

## Methods

### NewUserListRep

`func NewUserListRep() *UserListRep`

NewUserListRep instantiates a new UserListRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListRepWithDefaults

`func NewUserListRepWithDefaults() *UserListRep`

NewUserListRepWithDefaults instantiates a new UserListRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UserListRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserListRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserListRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserListRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *UserListRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UserListRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UserListRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *UserListRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *UserListRep) GetItems() []UserListRepItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserListRep) GetItemsOk() (*[]UserListRepItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserListRep) SetItems(v []UserListRepItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserListRep) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


