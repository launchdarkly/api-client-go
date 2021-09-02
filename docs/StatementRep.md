# StatementRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]map[string]interface{}** | Resource specifier strings | [optional] 
**NotResources** | Pointer to **[]map[string]interface{}** | Targeted resources are the resources NOT in this list. The \&quot;resources\&quot; field must be empty to use this field. | [optional] 
**Actions** | Pointer to **[]string** | Actions to perform on a resource | [optional] 
**NotActions** | Pointer to **[]string** | Targeted actions are the actions NOT in this list. The \&quot;actions\&quot; field must be empty to use this field. | [optional] 
**Effect** | Pointer to **string** |  | [optional] 

## Methods

### NewStatementRep

`func NewStatementRep() *StatementRep`

NewStatementRep instantiates a new StatementRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementRepWithDefaults

`func NewStatementRepWithDefaults() *StatementRep`

NewStatementRepWithDefaults instantiates a new StatementRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *StatementRep) GetResources() []map[string]interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *StatementRep) GetResourcesOk() (*[]map[string]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *StatementRep) SetResources(v []map[string]interface{})`

SetResources sets Resources field to given value.

### HasResources

`func (o *StatementRep) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetNotResources

`func (o *StatementRep) GetNotResources() []map[string]interface{}`

GetNotResources returns the NotResources field if non-nil, zero value otherwise.

### GetNotResourcesOk

`func (o *StatementRep) GetNotResourcesOk() (*[]map[string]interface{}, bool)`

GetNotResourcesOk returns a tuple with the NotResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResources

`func (o *StatementRep) SetNotResources(v []map[string]interface{})`

SetNotResources sets NotResources field to given value.

### HasNotResources

`func (o *StatementRep) HasNotResources() bool`

HasNotResources returns a boolean if a field has been set.

### GetActions

`func (o *StatementRep) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *StatementRep) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *StatementRep) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *StatementRep) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNotActions

`func (o *StatementRep) GetNotActions() []string`

GetNotActions returns the NotActions field if non-nil, zero value otherwise.

### GetNotActionsOk

`func (o *StatementRep) GetNotActionsOk() (*[]string, bool)`

GetNotActionsOk returns a tuple with the NotActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotActions

`func (o *StatementRep) SetNotActions(v []string)`

SetNotActions sets NotActions field to given value.

### HasNotActions

`func (o *StatementRep) HasNotActions() bool`

HasNotActions returns a boolean if a field has been set.

### GetEffect

`func (o *StatementRep) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *StatementRep) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *StatementRep) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *StatementRep) HasEffect() bool`

HasEffect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


