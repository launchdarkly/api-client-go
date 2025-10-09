# ModelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Name** | **string** | Human readable name of the model | 
**Key** | **string** | Unique key for the model | 
**Id** | **string** | Identifier for the model, for use with third party providers | 
**Icon** | Pointer to **string** | Icon for the model | [optional] 
**Provider** | Pointer to **string** | Provider for the model | [optional] 
**Global** | **bool** | Whether the model is global | 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | **[]string** |  | 
**Version** | **int32** |  | 
**CostPerInputToken** | Pointer to **float64** | Cost per input token in USD | [optional] 
**CostPerOutputToken** | Pointer to **float64** | Cost per output token in USD | [optional] 
**IsRestricted** | **bool** | Whether the model is restricted | 

## Methods

### NewModelConfig

`func NewModelConfig(name string, key string, id string, global bool, tags []string, version int32, isRestricted bool, ) *ModelConfig`

NewModelConfig instantiates a new ModelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConfigWithDefaults

`func NewModelConfigWithDefaults() *ModelConfig`

NewModelConfigWithDefaults instantiates a new ModelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ModelConfig) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ModelConfig) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ModelConfig) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ModelConfig) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetName

`func (o *ModelConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelConfig) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ModelConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelConfig) SetKey(v string)`

SetKey sets Key field to given value.


### GetId

`func (o *ModelConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelConfig) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *ModelConfig) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ModelConfig) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ModelConfig) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ModelConfig) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetProvider

`func (o *ModelConfig) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelConfig) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelConfig) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelConfig) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetGlobal

`func (o *ModelConfig) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ModelConfig) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ModelConfig) SetGlobal(v bool)`

SetGlobal sets Global field to given value.


### GetParams

`func (o *ModelConfig) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModelConfig) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModelConfig) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ModelConfig) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetCustomParams

`func (o *ModelConfig) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *ModelConfig) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *ModelConfig) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *ModelConfig) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetTags

`func (o *ModelConfig) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelConfig) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelConfig) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *ModelConfig) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelConfig) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelConfig) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCostPerInputToken

`func (o *ModelConfig) GetCostPerInputToken() float64`

GetCostPerInputToken returns the CostPerInputToken field if non-nil, zero value otherwise.

### GetCostPerInputTokenOk

`func (o *ModelConfig) GetCostPerInputTokenOk() (*float64, bool)`

GetCostPerInputTokenOk returns a tuple with the CostPerInputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerInputToken

`func (o *ModelConfig) SetCostPerInputToken(v float64)`

SetCostPerInputToken sets CostPerInputToken field to given value.

### HasCostPerInputToken

`func (o *ModelConfig) HasCostPerInputToken() bool`

HasCostPerInputToken returns a boolean if a field has been set.

### GetCostPerOutputToken

`func (o *ModelConfig) GetCostPerOutputToken() float64`

GetCostPerOutputToken returns the CostPerOutputToken field if non-nil, zero value otherwise.

### GetCostPerOutputTokenOk

`func (o *ModelConfig) GetCostPerOutputTokenOk() (*float64, bool)`

GetCostPerOutputTokenOk returns a tuple with the CostPerOutputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerOutputToken

`func (o *ModelConfig) SetCostPerOutputToken(v float64)`

SetCostPerOutputToken sets CostPerOutputToken field to given value.

### HasCostPerOutputToken

`func (o *ModelConfig) HasCostPerOutputToken() bool`

HasCostPerOutputToken returns a boolean if a field has been set.

### GetIsRestricted

`func (o *ModelConfig) GetIsRestricted() bool`

GetIsRestricted returns the IsRestricted field if non-nil, zero value otherwise.

### GetIsRestrictedOk

`func (o *ModelConfig) GetIsRestrictedOk() (*bool, bool)`

GetIsRestrictedOk returns a tuple with the IsRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestricted

`func (o *ModelConfig) SetIsRestricted(v bool)`

SetIsRestricted sets IsRestricted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


