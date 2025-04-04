# ModelConfigPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human readable name of the model | 
**Key** | **string** | Unique key for the model | 
**Id** | **string** | Identifier for the model, for use with third party providers | 
**Icon** | Pointer to **string** | Icon for the model | [optional] 
**Provider** | Pointer to **string** | Provider for the model | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomParams** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CostPerInputToken** | Pointer to **float64** | Cost per input token in USD | [optional] 
**CostPerOutputToken** | Pointer to **float64** | Cost per output token in USD | [optional] 

## Methods

### NewModelConfigPost

`func NewModelConfigPost(name string, key string, id string, ) *ModelConfigPost`

NewModelConfigPost instantiates a new ModelConfigPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConfigPostWithDefaults

`func NewModelConfigPostWithDefaults() *ModelConfigPost`

NewModelConfigPostWithDefaults instantiates a new ModelConfigPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelConfigPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelConfigPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelConfigPost) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ModelConfigPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelConfigPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelConfigPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetId

`func (o *ModelConfigPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelConfigPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelConfigPost) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *ModelConfigPost) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ModelConfigPost) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ModelConfigPost) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ModelConfigPost) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetProvider

`func (o *ModelConfigPost) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelConfigPost) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelConfigPost) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelConfigPost) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetParams

`func (o *ModelConfigPost) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModelConfigPost) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModelConfigPost) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ModelConfigPost) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetCustomParams

`func (o *ModelConfigPost) GetCustomParams() map[string]interface{}`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *ModelConfigPost) GetCustomParamsOk() (*map[string]interface{}, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *ModelConfigPost) SetCustomParams(v map[string]interface{})`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *ModelConfigPost) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetTags

`func (o *ModelConfigPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelConfigPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelConfigPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelConfigPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCostPerInputToken

`func (o *ModelConfigPost) GetCostPerInputToken() float64`

GetCostPerInputToken returns the CostPerInputToken field if non-nil, zero value otherwise.

### GetCostPerInputTokenOk

`func (o *ModelConfigPost) GetCostPerInputTokenOk() (*float64, bool)`

GetCostPerInputTokenOk returns a tuple with the CostPerInputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerInputToken

`func (o *ModelConfigPost) SetCostPerInputToken(v float64)`

SetCostPerInputToken sets CostPerInputToken field to given value.

### HasCostPerInputToken

`func (o *ModelConfigPost) HasCostPerInputToken() bool`

HasCostPerInputToken returns a boolean if a field has been set.

### GetCostPerOutputToken

`func (o *ModelConfigPost) GetCostPerOutputToken() float64`

GetCostPerOutputToken returns the CostPerOutputToken field if non-nil, zero value otherwise.

### GetCostPerOutputTokenOk

`func (o *ModelConfigPost) GetCostPerOutputTokenOk() (*float64, bool)`

GetCostPerOutputTokenOk returns a tuple with the CostPerOutputToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerOutputToken

`func (o *ModelConfigPost) SetCostPerOutputToken(v float64)`

SetCostPerOutputToken sets CostPerOutputToken field to given value.

### HasCostPerOutputToken

`func (o *ModelConfigPost) HasCostPerOutputToken() bool`

HasCostPerOutputToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


