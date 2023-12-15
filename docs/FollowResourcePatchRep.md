# FollowResourcePatchRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**FlagKey** | **string** |  | 
**ProjectKey** | Pointer to **string** |  | [optional] 
**EnvironmentKey** | Pointer to **string** |  | [optional] 
**Followers** | [**[]FollowFlagMember**](FollowFlagMember.md) |  | 
**AddedCount** | **int32** |  | 
**RemovedCount** | **int32** |  | 

## Methods

### NewFollowResourcePatchRep

`func NewFollowResourcePatchRep(links map[string]Link, flagKey string, followers []FollowFlagMember, addedCount int32, removedCount int32, ) *FollowResourcePatchRep`

NewFollowResourcePatchRep instantiates a new FollowResourcePatchRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowResourcePatchRepWithDefaults

`func NewFollowResourcePatchRepWithDefaults() *FollowResourcePatchRep`

NewFollowResourcePatchRepWithDefaults instantiates a new FollowResourcePatchRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FollowResourcePatchRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FollowResourcePatchRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FollowResourcePatchRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetFlagKey

`func (o *FollowResourcePatchRep) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *FollowResourcePatchRep) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *FollowResourcePatchRep) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetProjectKey

`func (o *FollowResourcePatchRep) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *FollowResourcePatchRep) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *FollowResourcePatchRep) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *FollowResourcePatchRep) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *FollowResourcePatchRep) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *FollowResourcePatchRep) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *FollowResourcePatchRep) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *FollowResourcePatchRep) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetFollowers

`func (o *FollowResourcePatchRep) GetFollowers() []FollowFlagMember`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *FollowResourcePatchRep) GetFollowersOk() (*[]FollowFlagMember, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *FollowResourcePatchRep) SetFollowers(v []FollowFlagMember)`

SetFollowers sets Followers field to given value.


### GetAddedCount

`func (o *FollowResourcePatchRep) GetAddedCount() int32`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *FollowResourcePatchRep) GetAddedCountOk() (*int32, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *FollowResourcePatchRep) SetAddedCount(v int32)`

SetAddedCount sets AddedCount field to given value.


### GetRemovedCount

`func (o *FollowResourcePatchRep) GetRemovedCount() int32`

GetRemovedCount returns the RemovedCount field if non-nil, zero value otherwise.

### GetRemovedCountOk

`func (o *FollowResourcePatchRep) GetRemovedCountOk() (*int32, bool)`

GetRemovedCountOk returns a tuple with the RemovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCount

`func (o *FollowResourcePatchRep) SetRemovedCount(v int32)`

SetRemovedCount sets RemovedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


