# AccessDeniedReasonRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]string** | Resource specifier strings | [optional] 
**NotResources** | Pointer to **[]string** | Targeted resources are the resources NOT in this list. The \&quot;resources\&quot; field must be empty to use this field. | [optional] 
**Actions** | Pointer to **[]string** | Actions to perform on a resource | [optional] 
**NotActions** | Pointer to **[]string** | Targeted actions are the actions NOT in this list. The \&quot;actions\&quot; field must be empty to use this field. | [optional] 
**Effect** | **string** |  | 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessDeniedReasonRep

`func NewAccessDeniedReasonRep(effect string, ) *AccessDeniedReasonRep`

NewAccessDeniedReasonRep instantiates a new AccessDeniedReasonRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDeniedReasonRepWithDefaults

`func NewAccessDeniedReasonRepWithDefaults() *AccessDeniedReasonRep`

NewAccessDeniedReasonRepWithDefaults instantiates a new AccessDeniedReasonRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *AccessDeniedReasonRep) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AccessDeniedReasonRep) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AccessDeniedReasonRep) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AccessDeniedReasonRep) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetNotResources

`func (o *AccessDeniedReasonRep) GetNotResources() []string`

GetNotResources returns the NotResources field if non-nil, zero value otherwise.

### GetNotResourcesOk

`func (o *AccessDeniedReasonRep) GetNotResourcesOk() (*[]string, bool)`

GetNotResourcesOk returns a tuple with the NotResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResources

`func (o *AccessDeniedReasonRep) SetNotResources(v []string)`

SetNotResources sets NotResources field to given value.

### HasNotResources

`func (o *AccessDeniedReasonRep) HasNotResources() bool`

HasNotResources returns a boolean if a field has been set.

### GetActions

`func (o *AccessDeniedReasonRep) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AccessDeniedReasonRep) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AccessDeniedReasonRep) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AccessDeniedReasonRep) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNotActions

`func (o *AccessDeniedReasonRep) GetNotActions() []string`

GetNotActions returns the NotActions field if non-nil, zero value otherwise.

### GetNotActionsOk

`func (o *AccessDeniedReasonRep) GetNotActionsOk() (*[]string, bool)`

GetNotActionsOk returns a tuple with the NotActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotActions

`func (o *AccessDeniedReasonRep) SetNotActions(v []string)`

SetNotActions sets NotActions field to given value.

### HasNotActions

`func (o *AccessDeniedReasonRep) HasNotActions() bool`

HasNotActions returns a boolean if a field has been set.

### GetEffect

`func (o *AccessDeniedReasonRep) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AccessDeniedReasonRep) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AccessDeniedReasonRep) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetRoleName

`func (o *AccessDeniedReasonRep) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AccessDeniedReasonRep) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AccessDeniedReasonRep) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AccessDeniedReasonRep) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

