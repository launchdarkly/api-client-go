# TeamRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRoleKeys** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**MemberIDs** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PermissionGrants** | Pointer to [**[]PermissionGrantRep**](PermissionGrantRep.md) |  | [optional] 
**ProjectKeys** | Pointer to **[]string** |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewTeamRep

`func NewTeamRep() *TeamRep`

NewTeamRep instantiates a new TeamRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRepWithDefaults

`func NewTeamRepWithDefaults() *TeamRep`

NewTeamRepWithDefaults instantiates a new TeamRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRoleKeys

`func (o *TeamRep) GetCustomRoleKeys() []string`

GetCustomRoleKeys returns the CustomRoleKeys field if non-nil, zero value otherwise.

### GetCustomRoleKeysOk

`func (o *TeamRep) GetCustomRoleKeysOk() (*[]string, bool)`

GetCustomRoleKeysOk returns a tuple with the CustomRoleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleKeys

`func (o *TeamRep) SetCustomRoleKeys(v []string)`

SetCustomRoleKeys sets CustomRoleKeys field to given value.

### HasCustomRoleKeys

`func (o *TeamRep) HasCustomRoleKeys() bool`

HasCustomRoleKeys returns a boolean if a field has been set.

### GetDescription

`func (o *TeamRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *TeamRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TeamRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TeamRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TeamRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMemberIDs

`func (o *TeamRep) GetMemberIDs() []string`

GetMemberIDs returns the MemberIDs field if non-nil, zero value otherwise.

### GetMemberIDsOk

`func (o *TeamRep) GetMemberIDsOk() (*[]string, bool)`

GetMemberIDsOk returns a tuple with the MemberIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIDs

`func (o *TeamRep) SetMemberIDs(v []string)`

SetMemberIDs sets MemberIDs field to given value.

### HasMemberIDs

`func (o *TeamRep) HasMemberIDs() bool`

HasMemberIDs returns a boolean if a field has been set.

### GetName

`func (o *TeamRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionGrants

`func (o *TeamRep) GetPermissionGrants() []PermissionGrantRep`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *TeamRep) GetPermissionGrantsOk() (*[]PermissionGrantRep, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *TeamRep) SetPermissionGrants(v []PermissionGrantRep)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *TeamRep) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.

### GetProjectKeys

`func (o *TeamRep) GetProjectKeys() []string`

GetProjectKeys returns the ProjectKeys field if non-nil, zero value otherwise.

### GetProjectKeysOk

`func (o *TeamRep) GetProjectKeysOk() (*[]string, bool)`

GetProjectKeysOk returns a tuple with the ProjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKeys

`func (o *TeamRep) SetProjectKeys(v []string)`

SetProjectKeys sets ProjectKeys field to given value.

### HasProjectKeys

`func (o *TeamRep) HasProjectKeys() bool`

HasProjectKeys returns a boolean if a field has been set.

### GetAccess

`func (o *TeamRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *TeamRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *TeamRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *TeamRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TeamRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TeamRep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *TeamRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TeamRep) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TeamRep) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TeamRep) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TeamRep) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *TeamRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TeamRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TeamRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TeamRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


