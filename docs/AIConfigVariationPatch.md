# AIConfigVariationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifiers** | Pointer to [**[]Classifier**](Classifier.md) | Replaces classifier rules on a judge variation. An empty array clears them. When this array is non-empty, messages must be empty.  | [optional] 
**Comment** | Pointer to **string** | Human-readable description of what this patch changes | [optional] 
**Description** | Pointer to **string** | Description for agent when the config is in agent mode. | [optional] 
**Instructions** | Pointer to **string** | Instructions for agent when the config is in agent mode. | [optional] 
**Messages** | Pointer to [**[]Message**](Message.md) | The prompt messages for this variation | [optional] 
**Model** | Pointer to **map[string]interface{}** | Freeform JSON object. LaunchDarkly does not validate the keys. In a model config, &#x60;params&#x60; and &#x60;customParams&#x60; hold parameters that the provider&#39;s API accepts, for example &#x60;{\&quot;temperature\&quot;: 0.7, \&quot;maxTokens\&quot;: 4096}&#x60;. | [optional] 
**ModelConfigKey** | Pointer to **string** | The key that GET /api/v2/projects/{projectKey}/ai-configs/model-configs returns for the model config. Do not use a custom model config&#39;s ID, which is the provider model ID.  | [optional] 
**ModelConfigVersion** | Pointer to **int32** | The exact model config version to set this variation to. Requires modelConfigKey to be set (either in this patch or already on the variation), and the version must exist for that model config. When omitted, the variation keeps the version it is currently set to, unless the patch changes modelConfigKey, in which case the variation records the new model config&#39;s current (latest) version.  | [optional] 
**Name** | Pointer to **string** | A human-readable name for the variation | [optional] 
**OutputFormat** | Pointer to **map[string]interface{}** | JSON Schema defining the structured output format for the variation. | [optional] 
**Published** | Pointer to **bool** | Whether to publish this version of the variation | [optional] 
**State** | Pointer to **string** | The state to move the variation to. Omit this field to keep the current state. - &#x60;archived&#x60;: Archive a published variation. You cannot archive a variation that is already archived. You also cannot archive the only variation of an AI Config. An AI Config itself cannot be archived. - &#x60;published&#x60;: Restore an archived variation. You cannot publish a variation that is already published.  | [optional] 
**Tools** | Pointer to [**[]VariationToolPost**](VariationToolPost.md) | List of tools to use for this variation. The latest version of the tool will be used. | [optional] 
**ToolKeys** | Pointer to **[]string** | List of tool keys to use for this variation. The latest version of the tool will be used. | [optional] 
**Skills** | Pointer to [**[]VariationSkillPost**](VariationSkillPost.md) | List of agent skills to attach to this variation. Replaces the current attachments. | [optional] 
**JudgeConfiguration** | Pointer to [**JudgeConfiguration**](JudgeConfiguration.md) |  | [optional] 

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

### GetClassifiers

`func (o *AIConfigVariationPatch) GetClassifiers() []Classifier`

GetClassifiers returns the Classifiers field if non-nil, zero value otherwise.

### GetClassifiersOk

`func (o *AIConfigVariationPatch) GetClassifiersOk() (*[]Classifier, bool)`

GetClassifiersOk returns a tuple with the Classifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifiers

`func (o *AIConfigVariationPatch) SetClassifiers(v []Classifier)`

SetClassifiers sets Classifiers field to given value.

### HasClassifiers

`func (o *AIConfigVariationPatch) HasClassifiers() bool`

HasClassifiers returns a boolean if a field has been set.

### GetComment

`func (o *AIConfigVariationPatch) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AIConfigVariationPatch) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AIConfigVariationPatch) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AIConfigVariationPatch) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *AIConfigVariationPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIConfigVariationPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIConfigVariationPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIConfigVariationPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstructions

`func (o *AIConfigVariationPatch) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AIConfigVariationPatch) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AIConfigVariationPatch) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AIConfigVariationPatch) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

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

### GetModelConfigVersion

`func (o *AIConfigVariationPatch) GetModelConfigVersion() int32`

GetModelConfigVersion returns the ModelConfigVersion field if non-nil, zero value otherwise.

### GetModelConfigVersionOk

`func (o *AIConfigVariationPatch) GetModelConfigVersionOk() (*int32, bool)`

GetModelConfigVersionOk returns a tuple with the ModelConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConfigVersion

`func (o *AIConfigVariationPatch) SetModelConfigVersion(v int32)`

SetModelConfigVersion sets ModelConfigVersion field to given value.

### HasModelConfigVersion

`func (o *AIConfigVariationPatch) HasModelConfigVersion() bool`

HasModelConfigVersion returns a boolean if a field has been set.

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

### GetOutputFormat

`func (o *AIConfigVariationPatch) GetOutputFormat() map[string]interface{}`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *AIConfigVariationPatch) GetOutputFormatOk() (*map[string]interface{}, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *AIConfigVariationPatch) SetOutputFormat(v map[string]interface{})`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *AIConfigVariationPatch) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

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

### GetTools

`func (o *AIConfigVariationPatch) GetTools() []VariationToolPost`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AIConfigVariationPatch) GetToolsOk() (*[]VariationToolPost, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AIConfigVariationPatch) SetTools(v []VariationToolPost)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AIConfigVariationPatch) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolKeys

`func (o *AIConfigVariationPatch) GetToolKeys() []string`

GetToolKeys returns the ToolKeys field if non-nil, zero value otherwise.

### GetToolKeysOk

`func (o *AIConfigVariationPatch) GetToolKeysOk() (*[]string, bool)`

GetToolKeysOk returns a tuple with the ToolKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolKeys

`func (o *AIConfigVariationPatch) SetToolKeys(v []string)`

SetToolKeys sets ToolKeys field to given value.

### HasToolKeys

`func (o *AIConfigVariationPatch) HasToolKeys() bool`

HasToolKeys returns a boolean if a field has been set.

### GetSkills

`func (o *AIConfigVariationPatch) GetSkills() []VariationSkillPost`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *AIConfigVariationPatch) GetSkillsOk() (*[]VariationSkillPost, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *AIConfigVariationPatch) SetSkills(v []VariationSkillPost)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *AIConfigVariationPatch) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetJudgeConfiguration

`func (o *AIConfigVariationPatch) GetJudgeConfiguration() JudgeConfiguration`

GetJudgeConfiguration returns the JudgeConfiguration field if non-nil, zero value otherwise.

### GetJudgeConfigurationOk

`func (o *AIConfigVariationPatch) GetJudgeConfigurationOk() (*JudgeConfiguration, bool)`

GetJudgeConfigurationOk returns a tuple with the JudgeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeConfiguration

`func (o *AIConfigVariationPatch) SetJudgeConfiguration(v JudgeConfiguration)`

SetJudgeConfiguration sets JudgeConfiguration field to given value.

### HasJudgeConfiguration

`func (o *AIConfigVariationPatch) HasJudgeConfiguration() bool`

HasJudgeConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


