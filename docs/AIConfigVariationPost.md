# AIConfigVariationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Messages** | [**[]Message**](Message.md) |  | 
**Model** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**ModelConfigKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAIConfigVariationPost

`func NewAIConfigVariationPost(key string, messages []Message, model map[string]interface{}, name string, ) *AIConfigVariationPost`

NewAIConfigVariationPost instantiates a new AIConfigVariationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigVariationPostWithDefaults

`func NewAIConfigVariationPostWithDefaults() *AIConfigVariationPost`

NewAIConfigVariationPostWithDefaults instantiates a new AIConfigVariationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AIConfigVariationPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIConfigVariationPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIConfigVariationPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetMessages

`func (o *AIConfigVariationPost) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AIConfigVariationPost) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AIConfigVariationPost) SetMessages(v []Message)`

SetMessages sets Messages field to given value.


### GetModel

`func (o *AIConfigVariationPost) GetModel() map[string]interface{}`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AIConfigVariationPost) GetModelOk() (*map[string]interface{}, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AIConfigVariationPost) SetModel(v map[string]interface{})`

SetModel sets Model field to given value.


### GetName

`func (o *AIConfigVariationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfigVariationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfigVariationPost) SetName(v string)`

SetName sets Name field to given value.


### GetModelConfigKey

`func (o *AIConfigVariationPost) GetModelConfigKey() string`

GetModelConfigKey returns the ModelConfigKey field if non-nil, zero value otherwise.

### GetModelConfigKeyOk

`func (o *AIConfigVariationPost) GetModelConfigKeyOk() (*string, bool)`

GetModelConfigKeyOk returns a tuple with the ModelConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConfigKey

`func (o *AIConfigVariationPost) SetModelConfigKey(v string)`

SetModelConfigKey sets ModelConfigKey field to given value.

### HasModelConfigKey

`func (o *AIConfigVariationPost) HasModelConfigKey() bool`

HasModelConfigKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


