# FlagFollowersByProjEnvGetRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Items** | Pointer to [**[]FollowersPerFlag**](FollowersPerFlag.md) | An array of flags and their followers | [optional] 

## Methods

### NewFlagFollowersByProjEnvGetRep

`func NewFlagFollowersByProjEnvGetRep(links map[string]Link, ) *FlagFollowersByProjEnvGetRep`

NewFlagFollowersByProjEnvGetRep instantiates a new FlagFollowersByProjEnvGetRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagFollowersByProjEnvGetRepWithDefaults

`func NewFlagFollowersByProjEnvGetRepWithDefaults() *FlagFollowersByProjEnvGetRep`

NewFlagFollowersByProjEnvGetRepWithDefaults instantiates a new FlagFollowersByProjEnvGetRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagFollowersByProjEnvGetRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagFollowersByProjEnvGetRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagFollowersByProjEnvGetRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *FlagFollowersByProjEnvGetRep) GetItems() []FollowersPerFlag`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlagFollowersByProjEnvGetRep) GetItemsOk() (*[]FollowersPerFlag, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlagFollowersByProjEnvGetRep) SetItems(v []FollowersPerFlag)`

SetItems sets Items field to given value.

### HasItems

`func (o *FlagFollowersByProjEnvGetRep) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


