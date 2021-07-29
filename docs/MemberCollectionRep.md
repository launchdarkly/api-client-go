# MemberCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]MemberRep**](MemberRep.md) |  | 
**Links** | [**map[string]InlineResponse200**](InlineResponse200.md) |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewMemberCollectionRep

`func NewMemberCollectionRep(items []MemberRep, links map[string]InlineResponse200, ) *MemberCollectionRep`

NewMemberCollectionRep instantiates a new MemberCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberCollectionRepWithDefaults

`func NewMemberCollectionRepWithDefaults() *MemberCollectionRep`

NewMemberCollectionRepWithDefaults instantiates a new MemberCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MemberCollectionRep) GetItems() []MemberRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MemberCollectionRep) GetItemsOk() (*[]MemberRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MemberCollectionRep) SetItems(v []MemberRep)`

SetItems sets Items field to given value.


### GetLinks

`func (o *MemberCollectionRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberCollectionRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberCollectionRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.


### GetTotalCount

`func (o *MemberCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MemberCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MemberCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *MemberCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


