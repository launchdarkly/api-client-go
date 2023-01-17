# UsersRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TotalCount** | **int32** | The total number of users in the environment | 
**Items** | [**[]UserRecord**](UserRecord.md) | Details on the users | 

## Methods

### NewUsersRep

`func NewUsersRep(totalCount int32, items []UserRecord, ) *UsersRep`

NewUsersRep instantiates a new UsersRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersRepWithDefaults

`func NewUsersRepWithDefaults() *UsersRep`

NewUsersRepWithDefaults instantiates a new UsersRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UsersRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UsersRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UsersRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UsersRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *UsersRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UsersRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UsersRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *UsersRep) GetItems() []UserRecord`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UsersRep) GetItemsOk() (*[]UserRecord, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UsersRep) SetItems(v []UserRecord)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


