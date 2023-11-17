# AccessDeniedReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]string** | Resource specifier strings | [optional] 
**NotResources** | Pointer to **[]string** | Targeted resources are the resources NOT in this list. The &lt;code&gt;resources&lt;/code&gt; and &lt;code&gt;notActions&lt;/code&gt; fields must be empty to use this field. | [optional] 
**Actions** | Pointer to **[]string** | Actions to perform on a resource | [optional] 
**NotActions** | Pointer to **[]string** | Targeted actions are the actions NOT in this list. The &lt;code&gt;actions&lt;/code&gt; and &lt;code&gt;notResources&lt;/code&gt; fields must be empty to use this field. | [optional] 
**Effect** | **string** | Whether this statement should allow or deny actions on the resources. | 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessDeniedReason

`func NewAccessDeniedReason(effect string, ) *AccessDeniedReason`

NewAccessDeniedReason instantiates a new AccessDeniedReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDeniedReasonWithDefaults

`func NewAccessDeniedReasonWithDefaults() *AccessDeniedReason`

NewAccessDeniedReasonWithDefaults instantiates a new AccessDeniedReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *AccessDeniedReason) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AccessDeniedReason) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AccessDeniedReason) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AccessDeniedReason) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetNotResources

`func (o *AccessDeniedReason) GetNotResources() []string`

GetNotResources returns the NotResources field if non-nil, zero value otherwise.

### GetNotResourcesOk

`func (o *AccessDeniedReason) GetNotResourcesOk() (*[]string, bool)`

GetNotResourcesOk returns a tuple with the NotResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResources

`func (o *AccessDeniedReason) SetNotResources(v []string)`

SetNotResources sets NotResources field to given value.

### HasNotResources

`func (o *AccessDeniedReason) HasNotResources() bool`

HasNotResources returns a boolean if a field has been set.

### GetActions

`func (o *AccessDeniedReason) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AccessDeniedReason) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AccessDeniedReason) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AccessDeniedReason) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNotActions

`func (o *AccessDeniedReason) GetNotActions() []string`

GetNotActions returns the NotActions field if non-nil, zero value otherwise.

### GetNotActionsOk

`func (o *AccessDeniedReason) GetNotActionsOk() (*[]string, bool)`

GetNotActionsOk returns a tuple with the NotActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotActions

`func (o *AccessDeniedReason) SetNotActions(v []string)`

SetNotActions sets NotActions field to given value.

### HasNotActions

`func (o *AccessDeniedReason) HasNotActions() bool`

HasNotActions returns a boolean if a field has been set.

### GetEffect

`func (o *AccessDeniedReason) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AccessDeniedReason) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AccessDeniedReason) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetRoleName

`func (o *AccessDeniedReason) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AccessDeniedReason) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AccessDeniedReason) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AccessDeniedReason) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


