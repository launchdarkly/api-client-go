# CustomRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the custom role | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the custom role | [optional] 
**Key** | **string** | The key of the custom role | 
**Name** | **string** | The name of the custom role | 
**Policy** | [**[]Statement**](Statement.md) | An array of the policies that comprise this custom role | 
**BasePermissions** | Pointer to **string** |  | [optional] 
**ResourceCategory** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to [**AssignedToRep**](AssignedToRep.md) |  | [optional] 
**PresetBundleVersion** | Pointer to **int32** | If created from a preset, the preset bundle version | [optional] 
**PresetStatements** | Pointer to [**[]Statement**](Statement.md) | If created from a preset, the read-only statements copied from the preset | [optional] 

## Methods

### NewCustomRole

`func NewCustomRole(id string, links map[string]Link, key string, name string, policy []Statement, ) *CustomRole`

NewCustomRole instantiates a new CustomRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleWithDefaults

`func NewCustomRoleWithDefaults() *CustomRole`

NewCustomRoleWithDefaults instantiates a new CustomRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRole) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *CustomRole) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomRole) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomRole) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetAccess

`func (o *CustomRole) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CustomRole) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CustomRole) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CustomRole) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetDescription

`func (o *CustomRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *CustomRole) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomRole) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomRole) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CustomRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRole) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *CustomRole) GetPolicy() []Statement`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CustomRole) GetPolicyOk() (*[]Statement, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CustomRole) SetPolicy(v []Statement)`

SetPolicy sets Policy field to given value.


### GetBasePermissions

`func (o *CustomRole) GetBasePermissions() string`

GetBasePermissions returns the BasePermissions field if non-nil, zero value otherwise.

### GetBasePermissionsOk

`func (o *CustomRole) GetBasePermissionsOk() (*string, bool)`

GetBasePermissionsOk returns a tuple with the BasePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePermissions

`func (o *CustomRole) SetBasePermissions(v string)`

SetBasePermissions sets BasePermissions field to given value.

### HasBasePermissions

`func (o *CustomRole) HasBasePermissions() bool`

HasBasePermissions returns a boolean if a field has been set.

### GetResourceCategory

`func (o *CustomRole) GetResourceCategory() string`

GetResourceCategory returns the ResourceCategory field if non-nil, zero value otherwise.

### GetResourceCategoryOk

`func (o *CustomRole) GetResourceCategoryOk() (*string, bool)`

GetResourceCategoryOk returns a tuple with the ResourceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCategory

`func (o *CustomRole) SetResourceCategory(v string)`

SetResourceCategory sets ResourceCategory field to given value.

### HasResourceCategory

`func (o *CustomRole) HasResourceCategory() bool`

HasResourceCategory returns a boolean if a field has been set.

### GetAssignedTo

`func (o *CustomRole) GetAssignedTo() AssignedToRep`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *CustomRole) GetAssignedToOk() (*AssignedToRep, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *CustomRole) SetAssignedTo(v AssignedToRep)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *CustomRole) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetPresetBundleVersion

`func (o *CustomRole) GetPresetBundleVersion() int32`

GetPresetBundleVersion returns the PresetBundleVersion field if non-nil, zero value otherwise.

### GetPresetBundleVersionOk

`func (o *CustomRole) GetPresetBundleVersionOk() (*int32, bool)`

GetPresetBundleVersionOk returns a tuple with the PresetBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetBundleVersion

`func (o *CustomRole) SetPresetBundleVersion(v int32)`

SetPresetBundleVersion sets PresetBundleVersion field to given value.

### HasPresetBundleVersion

`func (o *CustomRole) HasPresetBundleVersion() bool`

HasPresetBundleVersion returns a boolean if a field has been set.

### GetPresetStatements

`func (o *CustomRole) GetPresetStatements() []Statement`

GetPresetStatements returns the PresetStatements field if non-nil, zero value otherwise.

### GetPresetStatementsOk

`func (o *CustomRole) GetPresetStatementsOk() (*[]Statement, bool)`

GetPresetStatementsOk returns a tuple with the PresetStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetStatements

`func (o *CustomRole) SetPresetStatements(v []Statement)`

SetPresetStatements sets PresetStatements field to given value.

### HasPresetStatements

`func (o *CustomRole) HasPresetStatements() bool`

HasPresetStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


