# AIConfigVariationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Human-readable description of this variation | [optional] 
**Description** | Pointer to **string** | Returns the description for the agent. This is only returned for agent variations. | [optional] 
**Instructions** | Pointer to **string** | Returns the instructions for the agent. This is only returned for agent variations. | [optional] 
**Key** | **string** |  | 
**Messages** | Pointer to [**[]Message**](Message.md) |  | [optional] 
**Model** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**ModelConfigKey** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to [**[]VariationToolPost**](VariationToolPost.md) | List of tools to use for this variation. The latest version of the tool will be used. | [optional] 
**ToolKeys** | Pointer to **[]string** | List of tool keys to use for this variation. The latest version of the tool will be used. | [optional] 
**JudgeConfiguration** | Pointer to [**JudgeConfiguration**](JudgeConfiguration.md) |  | [optional] 

## Methods

### NewAIConfigVariationPost

`func NewAIConfigVariationPost(key string, name string, ) *AIConfigVariationPost`

NewAIConfigVariationPost instantiates a new AIConfigVariationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigVariationPostWithDefaults

`func NewAIConfigVariationPostWithDefaults() *AIConfigVariationPost`

NewAIConfigVariationPostWithDefaults instantiates a new AIConfigVariationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *AIConfigVariationPost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AIConfigVariationPost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AIConfigVariationPost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AIConfigVariationPost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *AIConfigVariationPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIConfigVariationPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIConfigVariationPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIConfigVariationPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstructions

`func (o *AIConfigVariationPost) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AIConfigVariationPost) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AIConfigVariationPost) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AIConfigVariationPost) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

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

### HasMessages

`func (o *AIConfigVariationPost) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

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

### HasModel

`func (o *AIConfigVariationPost) HasModel() bool`

HasModel returns a boolean if a field has been set.

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

### GetTools

`func (o *AIConfigVariationPost) GetTools() []VariationToolPost`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AIConfigVariationPost) GetToolsOk() (*[]VariationToolPost, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AIConfigVariationPost) SetTools(v []VariationToolPost)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AIConfigVariationPost) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolKeys

`func (o *AIConfigVariationPost) GetToolKeys() []string`

GetToolKeys returns the ToolKeys field if non-nil, zero value otherwise.

### GetToolKeysOk

`func (o *AIConfigVariationPost) GetToolKeysOk() (*[]string, bool)`

GetToolKeysOk returns a tuple with the ToolKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolKeys

`func (o *AIConfigVariationPost) SetToolKeys(v []string)`

SetToolKeys sets ToolKeys field to given value.

### HasToolKeys

`func (o *AIConfigVariationPost) HasToolKeys() bool`

HasToolKeys returns a boolean if a field has been set.

### GetJudgeConfiguration

`func (o *AIConfigVariationPost) GetJudgeConfiguration() JudgeConfiguration`

GetJudgeConfiguration returns the JudgeConfiguration field if non-nil, zero value otherwise.

### GetJudgeConfigurationOk

`func (o *AIConfigVariationPost) GetJudgeConfigurationOk() (*JudgeConfiguration, bool)`

GetJudgeConfigurationOk returns a tuple with the JudgeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeConfiguration

`func (o *AIConfigVariationPost) SetJudgeConfiguration(v JudgeConfiguration)`

SetJudgeConfiguration sets JudgeConfiguration field to given value.

### HasJudgeConfiguration

`func (o *AIConfigVariationPost) HasJudgeConfiguration() bool`

HasJudgeConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


