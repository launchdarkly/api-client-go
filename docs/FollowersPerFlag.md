# FollowersPerFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlagKey** | Pointer to **string** |  | [optional] 
**Followers** | Pointer to [**[]FollowFlagMember**](FollowFlagMember.md) |  | [optional] 

## Methods

### NewFollowersPerFlag

`func NewFollowersPerFlag() *FollowersPerFlag`

NewFollowersPerFlag instantiates a new FollowersPerFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowersPerFlagWithDefaults

`func NewFollowersPerFlagWithDefaults() *FollowersPerFlag`

NewFollowersPerFlagWithDefaults instantiates a new FollowersPerFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlagKey

`func (o *FollowersPerFlag) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *FollowersPerFlag) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *FollowersPerFlag) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *FollowersPerFlag) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetFollowers

`func (o *FollowersPerFlag) GetFollowers() []FollowFlagMember`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *FollowersPerFlag) GetFollowersOk() (*[]FollowFlagMember, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *FollowersPerFlag) SetFollowers(v []FollowFlagMember)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *FollowersPerFlag) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


