# StatementPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]string** | Resource specifier strings | [optional] 
**NotResources** | Pointer to **[]string** | Targeted resources are the resources NOT in this list. The &lt;code&gt;resources&lt;/code&gt; field must be empty to use this field. | [optional] 
**Actions** | Pointer to **[]string** | Actions to perform on a resource | [optional] 
**NotActions** | Pointer to **[]string** | Targeted actions are the actions NOT in this list. The &lt;code&gt;actions&lt;/code&gt; field must be empty to use this field. | [optional] 
**Effect** | **string** | Whether this statement should allow or deny actions on the resources. | 

## Methods

### NewStatementPost

`func NewStatementPost(effect string, ) *StatementPost`

NewStatementPost instantiates a new StatementPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementPostWithDefaults

`func NewStatementPostWithDefaults() *StatementPost`

NewStatementPostWithDefaults instantiates a new StatementPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *StatementPost) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *StatementPost) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *StatementPost) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *StatementPost) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetNotResources

`func (o *StatementPost) GetNotResources() []string`

GetNotResources returns the NotResources field if non-nil, zero value otherwise.

### GetNotResourcesOk

`func (o *StatementPost) GetNotResourcesOk() (*[]string, bool)`

GetNotResourcesOk returns a tuple with the NotResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResources

`func (o *StatementPost) SetNotResources(v []string)`

SetNotResources sets NotResources field to given value.

### HasNotResources

`func (o *StatementPost) HasNotResources() bool`

HasNotResources returns a boolean if a field has been set.

### GetActions

`func (o *StatementPost) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *StatementPost) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *StatementPost) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *StatementPost) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNotActions

`func (o *StatementPost) GetNotActions() []string`

GetNotActions returns the NotActions field if non-nil, zero value otherwise.

### GetNotActionsOk

`func (o *StatementPost) GetNotActionsOk() (*[]string, bool)`

GetNotActionsOk returns a tuple with the NotActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotActions

`func (o *StatementPost) SetNotActions(v []string)`

SetNotActions sets NotActions field to given value.

### HasNotActions

`func (o *StatementPost) HasNotActions() bool`

HasNotActions returns a boolean if a field has been set.

### GetEffect

`func (o *StatementPost) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *StatementPost) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *StatementPost) SetEffect(v string)`

SetEffect sets Effect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


