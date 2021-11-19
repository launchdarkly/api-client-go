# CustomRolePostData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the custom role | 
**Key** | **string** | The custom role key | 
**Description** | Pointer to **string** | Description of custom role | [optional] 
**Policy** | [**[]StatementPost**](StatementPost.md) |  | 
**BasePermissions** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomRolePostData

`func NewCustomRolePostData(name string, key string, policy []StatementPost, ) *CustomRolePostData`

NewCustomRolePostData instantiates a new CustomRolePostData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRolePostDataWithDefaults

`func NewCustomRolePostDataWithDefaults() *CustomRolePostData`

NewCustomRolePostDataWithDefaults instantiates a new CustomRolePostData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomRolePostData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRolePostData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRolePostData) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *CustomRolePostData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomRolePostData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomRolePostData) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *CustomRolePostData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomRolePostData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomRolePostData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomRolePostData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicy

`func (o *CustomRolePostData) GetPolicy() []StatementPost`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CustomRolePostData) GetPolicyOk() (*[]StatementPost, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CustomRolePostData) SetPolicy(v []StatementPost)`

SetPolicy sets Policy field to given value.


### GetBasePermissions

`func (o *CustomRolePostData) GetBasePermissions() string`

GetBasePermissions returns the BasePermissions field if non-nil, zero value otherwise.

### GetBasePermissionsOk

`func (o *CustomRolePostData) GetBasePermissionsOk() (*string, bool)`

GetBasePermissionsOk returns a tuple with the BasePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePermissions

`func (o *CustomRolePostData) SetBasePermissions(v string)`

SetBasePermissions sets BasePermissions field to given value.

### HasBasePermissions

`func (o *CustomRolePostData) HasBasePermissions() bool`

HasBasePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


