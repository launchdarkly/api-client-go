# FlagFollowersGetRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Items** | [**[]FollowFlagMember**](FollowFlagMember.md) | An array of members who are following this flag | 

## Methods

### NewFlagFollowersGetRep

`func NewFlagFollowersGetRep(links map[string]Link, items []FollowFlagMember, ) *FlagFollowersGetRep`

NewFlagFollowersGetRep instantiates a new FlagFollowersGetRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagFollowersGetRepWithDefaults

`func NewFlagFollowersGetRepWithDefaults() *FlagFollowersGetRep`

NewFlagFollowersGetRepWithDefaults instantiates a new FlagFollowersGetRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagFollowersGetRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagFollowersGetRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagFollowersGetRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *FlagFollowersGetRep) GetItems() []FollowFlagMember`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlagFollowersGetRep) GetItemsOk() (*[]FollowFlagMember, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlagFollowersGetRep) SetItems(v []FollowFlagMember)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


