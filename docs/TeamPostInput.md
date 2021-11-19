# TeamPostInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRoleKeys** | Pointer to **[]string** | List of custom role keys the team will access | [optional] 
**Description** | Pointer to **string** | A description of the team | [optional] 
**Key** | **string** | The team&#39;s key or ID | 
**MemberIDs** | Pointer to **[]string** | A list of member IDs who belong to the team | [optional] 
**Name** | **string** | A human-friendly name for the team | 
**PermissionGrants** | Pointer to [**[]PermissionGrantInput**](PermissionGrantInput.md) | A list of permission grants to apply to the team. Can use \&quot;maintainTeam\&quot; to add team maintainers | [optional] 

## Methods

### NewTeamPostInput

`func NewTeamPostInput(key string, name string, ) *TeamPostInput`

NewTeamPostInput instantiates a new TeamPostInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPostInputWithDefaults

`func NewTeamPostInputWithDefaults() *TeamPostInput`

NewTeamPostInputWithDefaults instantiates a new TeamPostInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRoleKeys

`func (o *TeamPostInput) GetCustomRoleKeys() []string`

GetCustomRoleKeys returns the CustomRoleKeys field if non-nil, zero value otherwise.

### GetCustomRoleKeysOk

`func (o *TeamPostInput) GetCustomRoleKeysOk() (*[]string, bool)`

GetCustomRoleKeysOk returns a tuple with the CustomRoleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleKeys

`func (o *TeamPostInput) SetCustomRoleKeys(v []string)`

SetCustomRoleKeys sets CustomRoleKeys field to given value.

### HasCustomRoleKeys

`func (o *TeamPostInput) HasCustomRoleKeys() bool`

HasCustomRoleKeys returns a boolean if a field has been set.

### GetDescription

`func (o *TeamPostInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamPostInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamPostInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamPostInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *TeamPostInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TeamPostInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TeamPostInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetMemberIDs

`func (o *TeamPostInput) GetMemberIDs() []string`

GetMemberIDs returns the MemberIDs field if non-nil, zero value otherwise.

### GetMemberIDsOk

`func (o *TeamPostInput) GetMemberIDsOk() (*[]string, bool)`

GetMemberIDsOk returns a tuple with the MemberIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIDs

`func (o *TeamPostInput) SetMemberIDs(v []string)`

SetMemberIDs sets MemberIDs field to given value.

### HasMemberIDs

`func (o *TeamPostInput) HasMemberIDs() bool`

HasMemberIDs returns a boolean if a field has been set.

### GetName

`func (o *TeamPostInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamPostInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamPostInput) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionGrants

`func (o *TeamPostInput) GetPermissionGrants() []PermissionGrantInput`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *TeamPostInput) GetPermissionGrantsOk() (*[]PermissionGrantInput, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *TeamPostInput) SetPermissionGrants(v []PermissionGrantInput)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *TeamPostInput) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


