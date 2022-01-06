# ExpandedTeamRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRoles** | Pointer to [**[]CustomRolesRep**](CustomRolesRep.md) |  | [optional] 
**TeamMaintainers** | Pointer to [**[]MemberSummaryRep**](MemberSummaryRep.md) |  | [optional] 
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

### NewExpandedTeamRep

`func NewExpandedTeamRep() *ExpandedTeamRep`

NewExpandedTeamRep instantiates a new ExpandedTeamRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedTeamRepWithDefaults

`func NewExpandedTeamRepWithDefaults() *ExpandedTeamRep`

NewExpandedTeamRepWithDefaults instantiates a new ExpandedTeamRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRoles

`func (o *ExpandedTeamRep) GetCustomRoles() []CustomRolesRep`

GetCustomRoles returns the CustomRoles field if non-nil, zero value otherwise.

### GetCustomRolesOk

`func (o *ExpandedTeamRep) GetCustomRolesOk() (*[]CustomRolesRep, bool)`

GetCustomRolesOk returns a tuple with the CustomRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoles

`func (o *ExpandedTeamRep) SetCustomRoles(v []CustomRolesRep)`

SetCustomRoles sets CustomRoles field to given value.

### HasCustomRoles

`func (o *ExpandedTeamRep) HasCustomRoles() bool`

HasCustomRoles returns a boolean if a field has been set.

### GetTeamMaintainers

`func (o *ExpandedTeamRep) GetTeamMaintainers() []MemberSummaryRep`

GetTeamMaintainers returns the TeamMaintainers field if non-nil, zero value otherwise.

### GetTeamMaintainersOk

`func (o *ExpandedTeamRep) GetTeamMaintainersOk() (*[]MemberSummaryRep, bool)`

GetTeamMaintainersOk returns a tuple with the TeamMaintainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMaintainers

`func (o *ExpandedTeamRep) SetTeamMaintainers(v []MemberSummaryRep)`

SetTeamMaintainers sets TeamMaintainers field to given value.

### HasTeamMaintainers

`func (o *ExpandedTeamRep) HasTeamMaintainers() bool`

HasTeamMaintainers returns a boolean if a field has been set.

### GetCustomRoleKeys

`func (o *ExpandedTeamRep) GetCustomRoleKeys() []string`

GetCustomRoleKeys returns the CustomRoleKeys field if non-nil, zero value otherwise.

### GetCustomRoleKeysOk

`func (o *ExpandedTeamRep) GetCustomRoleKeysOk() (*[]string, bool)`

GetCustomRoleKeysOk returns a tuple with the CustomRoleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleKeys

`func (o *ExpandedTeamRep) SetCustomRoleKeys(v []string)`

SetCustomRoleKeys sets CustomRoleKeys field to given value.

### HasCustomRoleKeys

`func (o *ExpandedTeamRep) HasCustomRoleKeys() bool`

HasCustomRoleKeys returns a boolean if a field has been set.

### GetDescription

`func (o *ExpandedTeamRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandedTeamRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandedTeamRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandedTeamRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *ExpandedTeamRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedTeamRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedTeamRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ExpandedTeamRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMemberIDs

`func (o *ExpandedTeamRep) GetMemberIDs() []string`

GetMemberIDs returns the MemberIDs field if non-nil, zero value otherwise.

### GetMemberIDsOk

`func (o *ExpandedTeamRep) GetMemberIDsOk() (*[]string, bool)`

GetMemberIDsOk returns a tuple with the MemberIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIDs

`func (o *ExpandedTeamRep) SetMemberIDs(v []string)`

SetMemberIDs sets MemberIDs field to given value.

### HasMemberIDs

`func (o *ExpandedTeamRep) HasMemberIDs() bool`

HasMemberIDs returns a boolean if a field has been set.

### GetName

`func (o *ExpandedTeamRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedTeamRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedTeamRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpandedTeamRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionGrants

`func (o *ExpandedTeamRep) GetPermissionGrants() []PermissionGrantRep`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *ExpandedTeamRep) GetPermissionGrantsOk() (*[]PermissionGrantRep, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *ExpandedTeamRep) SetPermissionGrants(v []PermissionGrantRep)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *ExpandedTeamRep) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.

### GetProjectKeys

`func (o *ExpandedTeamRep) GetProjectKeys() []string`

GetProjectKeys returns the ProjectKeys field if non-nil, zero value otherwise.

### GetProjectKeysOk

`func (o *ExpandedTeamRep) GetProjectKeysOk() (*[]string, bool)`

GetProjectKeysOk returns a tuple with the ProjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKeys

`func (o *ExpandedTeamRep) SetProjectKeys(v []string)`

SetProjectKeys sets ProjectKeys field to given value.

### HasProjectKeys

`func (o *ExpandedTeamRep) HasProjectKeys() bool`

HasProjectKeys returns a boolean if a field has been set.

### GetAccess

`func (o *ExpandedTeamRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ExpandedTeamRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ExpandedTeamRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ExpandedTeamRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExpandedTeamRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExpandedTeamRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExpandedTeamRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExpandedTeamRep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *ExpandedTeamRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandedTeamRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandedTeamRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpandedTeamRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExpandedTeamRep) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExpandedTeamRep) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExpandedTeamRep) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExpandedTeamRep) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *ExpandedTeamRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpandedTeamRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpandedTeamRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExpandedTeamRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


