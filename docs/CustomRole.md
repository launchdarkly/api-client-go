# CustomRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Links** | [**map[string]Link**](Link.md) |  | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Policy** | [**[]Statement**](Statement.md) |  | 
**BasePermissions** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


