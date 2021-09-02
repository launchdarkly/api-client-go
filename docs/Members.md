# Members

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Member**](Member.md) |  | 
**Links** | [**map[string]Link**](Link.md) |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewMembers

`func NewMembers(items []Member, links map[string]Link, ) *Members`

NewMembers instantiates a new Members object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembersWithDefaults

`func NewMembersWithDefaults() *Members`

NewMembersWithDefaults instantiates a new Members object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Members) GetItems() []Member`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Members) GetItemsOk() (*[]Member, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Members) SetItems(v []Member)`

SetItems sets Items field to given value.


### GetLinks

`func (o *Members) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Members) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Members) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetTotalCount

`func (o *Members) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Members) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Members) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Members) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


