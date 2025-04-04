# AIConfigVariationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]Message**](Message.md) |  | [optional] 
**Model** | Pointer to **map[string]interface{}** |  | [optional] 
**ModelConfigKey** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** | One of &#39;archived&#39;, &#39;published&#39; | [optional] 

## Methods

### NewAIConfigVariationPatch

`func NewAIConfigVariationPatch() *AIConfigVariationPatch`

NewAIConfigVariationPatch instantiates a new AIConfigVariationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigVariationPatchWithDefaults

`func NewAIConfigVariationPatchWithDefaults() *AIConfigVariationPatch`

NewAIConfigVariationPatchWithDefaults instantiates a new AIConfigVariationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *AIConfigVariationPatch) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AIConfigVariationPatch) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AIConfigVariationPatch) SetMessages(v []Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AIConfigVariationPatch) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetModel

`func (o *AIConfigVariationPatch) GetModel() map[string]interface{}`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AIConfigVariationPatch) GetModelOk() (*map[string]interface{}, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AIConfigVariationPatch) SetModel(v map[string]interface{})`

SetModel sets Model field to given value.

### HasModel

`func (o *AIConfigVariationPatch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelConfigKey

`func (o *AIConfigVariationPatch) GetModelConfigKey() string`

GetModelConfigKey returns the ModelConfigKey field if non-nil, zero value otherwise.

### GetModelConfigKeyOk

`func (o *AIConfigVariationPatch) GetModelConfigKeyOk() (*string, bool)`

GetModelConfigKeyOk returns a tuple with the ModelConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConfigKey

`func (o *AIConfigVariationPatch) SetModelConfigKey(v string)`

SetModelConfigKey sets ModelConfigKey field to given value.

### HasModelConfigKey

`func (o *AIConfigVariationPatch) HasModelConfigKey() bool`

HasModelConfigKey returns a boolean if a field has been set.

### GetName

`func (o *AIConfigVariationPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfigVariationPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfigVariationPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AIConfigVariationPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublished

`func (o *AIConfigVariationPatch) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AIConfigVariationPatch) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AIConfigVariationPatch) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AIConfigVariationPatch) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetState

`func (o *AIConfigVariationPatch) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AIConfigVariationPatch) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AIConfigVariationPatch) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AIConfigVariationPatch) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


