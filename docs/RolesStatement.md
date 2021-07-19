# RolesStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to [**RolesActionList**](RolesActionList.md) |  | [optional] 
**Resources** | Pointer to [**RolesResourceList**](RolesResourceList.md) |  | [optional] 

## Methods

### NewRolesStatement

`func NewRolesStatement() *RolesStatement`

NewRolesStatement instantiates a new RolesStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesStatementWithDefaults

`func NewRolesStatementWithDefaults() *RolesStatement`

NewRolesStatementWithDefaults instantiates a new RolesStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *RolesStatement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *RolesStatement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *RolesStatement) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *RolesStatement) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetActions

`func (o *RolesStatement) GetActions() RolesActionList`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RolesStatement) GetActionsOk() (*RolesActionList, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RolesStatement) SetActions(v RolesActionList)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RolesStatement) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetResources

`func (o *RolesStatement) GetResources() RolesResourceList`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RolesStatement) GetResourcesOk() (*RolesResourceList, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RolesStatement) SetResources(v RolesResourceList)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RolesStatement) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


