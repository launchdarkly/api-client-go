# ModelConfigPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Immutable provider model identifier. To use a different model identifier, create a new model config. | [optional] 
**Provider** | Pointer to **string** | Immutable model provider. To use a different provider, create a new model config. | [optional] 
**Name** | Pointer to **string** | Human-readable name of the model | [optional] 
**CostPerInputToken** | Pointer to **float64** | Cost per input token in USD | [optional] 
**CostPerOutputToken** | Pointer to **float64** | Cost per output token in USD | [optional] 
**CostPerCachedInputToken** | Pointer to **float64** | Cost per cached input token in USD | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**MaintainerId** | Pointer to **string** |  | [optional] 
**MaintainerTeamKey** | Pointer to **string** |  | [optional] 

## Methods

### NewModelConfigPatch

`func NewModelConfigPatch() *ModelConfigPatch`

NewModelConfigPatch instantiates a new ModelConfigPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConfigPatchWithDefaults

`func NewModelConfigPatchWithDefaults() *ModelConfigPatch`

NewModelConfigPatchWithDefaults instantiates a new ModelConfigPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelConfigPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelConfigPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelConfigPatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelConfigPatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelConfigPatch) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelConfigPatch) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelConfigPatch) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelConfigPatch) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetName

`func (o *ModelConfigPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelConfigPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelConfigPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelConfigPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCostPerInputToken

`func (o *ModelConfigPatch) GetCostPerInputToken() float64`

GetCostPerInputToken returns the CostPerInputToken field if non-nil, zero value otherwise.

### GetCostPerInputTokenOk

`func (o *ModelConfigPatch) GetCostPerInputTokenOk() (*float64, bool)`

GetCostPerInputTokenOk returns a tuple with the CostPerInputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerInputToken

`func (o *ModelConfigPatch) SetCostPerInputToken(v float64)`

SetCostPerInputToken sets CostPerInputToken field to given value.

### HasCostPerInputToken

`func (o *ModelConfigPatch) HasCostPerInputToken() bool`

HasCostPerInputToken returns a boolean if a field has been set.

### GetCostPerOutputToken

`func (o *ModelConfigPatch) GetCostPerOutputToken() float64`

GetCostPerOutputToken returns the CostPerOutputToken field if non-nil, zero value otherwise.

### GetCostPerOutputTokenOk

`func (o *ModelConfigPatch) GetCostPerOutputTokenOk() (*float64, bool)`

GetCostPerOutputTokenOk returns a tuple with the CostPerOutputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerOutputToken

`func (o *ModelConfigPatch) SetCostPerOutputToken(v float64)`

SetCostPerOutputToken sets CostPerOutputToken field to given value.

### HasCostPerOutputToken

`func (o *ModelConfigPatch) HasCostPerOutputToken() bool`

HasCostPerOutputToken returns a boolean if a field has been set.

### GetCostPerCachedInputToken

`func (o *ModelConfigPatch) GetCostPerCachedInputToken() float64`

GetCostPerCachedInputToken returns the CostPerCachedInputToken field if non-nil, zero value otherwise.

### GetCostPerCachedInputTokenOk

`func (o *ModelConfigPatch) GetCostPerCachedInputTokenOk() (*float64, bool)`

GetCostPerCachedInputTokenOk returns a tuple with the CostPerCachedInputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCachedInputToken

`func (o *ModelConfigPatch) SetCostPerCachedInputToken(v float64)`

SetCostPerCachedInputToken sets CostPerCachedInputToken field to given value.

### HasCostPerCachedInputToken

`func (o *ModelConfigPatch) HasCostPerCachedInputToken() bool`

HasCostPerCachedInputToken returns a boolean if a field has been set.

### GetParams

`func (o *ModelConfigPatch) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModelConfigPatch) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModelConfigPatch) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ModelConfigPatch) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetCustomParams

`func (o *ModelConfigPatch) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *ModelConfigPatch) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *ModelConfigPatch) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *ModelConfigPatch) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetTags

`func (o *ModelConfigPatch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelConfigPatch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelConfigPatch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelConfigPatch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMaintainerId

`func (o *ModelConfigPatch) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ModelConfigPatch) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ModelConfigPatch) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ModelConfigPatch) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *ModelConfigPatch) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *ModelConfigPatch) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *ModelConfigPatch) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *ModelConfigPatch) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


