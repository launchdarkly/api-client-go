# TeamCustomRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**TeamProjects**](TeamProjects.md) |  | [optional] 
**AppliedOn** | Pointer to **int64** |  | [optional] 

## Methods

### NewTeamCustomRole

`func NewTeamCustomRole() *TeamCustomRole`

NewTeamCustomRole instantiates a new TeamCustomRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCustomRoleWithDefaults

`func NewTeamCustomRoleWithDefaults() *TeamCustomRole`

NewTeamCustomRoleWithDefaults instantiates a new TeamCustomRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TeamCustomRole) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TeamCustomRole) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TeamCustomRole) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TeamCustomRole) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TeamCustomRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamCustomRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamCustomRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamCustomRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjects

`func (o *TeamCustomRole) GetProjects() TeamProjects`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *TeamCustomRole) GetProjectsOk() (*TeamProjects, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *TeamCustomRole) SetProjects(v TeamProjects)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *TeamCustomRole) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetAppliedOn

`func (o *TeamCustomRole) GetAppliedOn() int64`

GetAppliedOn returns the AppliedOn field if non-nil, zero value otherwise.

### GetAppliedOnOk

`func (o *TeamCustomRole) GetAppliedOnOk() (*int64, bool)`

GetAppliedOnOk returns a tuple with the AppliedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOn

`func (o *TeamCustomRole) SetAppliedOn(v int64)`

SetAppliedOn sets AppliedOn field to given value.

### HasAppliedOn

`func (o *TeamCustomRole) HasAppliedOn() bool`

HasAppliedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


