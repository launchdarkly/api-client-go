# AgentSkillReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiConfigKey** | **string** | The key of the config with a variation that references this skill. | 
**AiConfigName** | **string** | The name of the config with a variation that references this skill. | 
**VariationId** | **string** | The ID of the variation that references this skill. | 
**VariationKey** | **string** | The key of the config variation that references this skill. | 
**VariationName** | **string** | The name of the variation that references this skill. | 
**ResourceVersion** | **int32** | The version of the skill being referenced. | 

## Methods

### NewAgentSkillReference

`func NewAgentSkillReference(aiConfigKey string, aiConfigName string, variationId string, variationKey string, variationName string, resourceVersion int32, ) *AgentSkillReference`

NewAgentSkillReference instantiates a new AgentSkillReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSkillReferenceWithDefaults

`func NewAgentSkillReferenceWithDefaults() *AgentSkillReference`

NewAgentSkillReferenceWithDefaults instantiates a new AgentSkillReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiConfigKey

`func (o *AgentSkillReference) GetAiConfigKey() string`

GetAiConfigKey returns the AiConfigKey field if non-nil, zero value otherwise.

### GetAiConfigKeyOk

`func (o *AgentSkillReference) GetAiConfigKeyOk() (*string, bool)`

GetAiConfigKeyOk returns a tuple with the AiConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigKey

`func (o *AgentSkillReference) SetAiConfigKey(v string)`

SetAiConfigKey sets AiConfigKey field to given value.


### GetAiConfigName

`func (o *AgentSkillReference) GetAiConfigName() string`

GetAiConfigName returns the AiConfigName field if non-nil, zero value otherwise.

### GetAiConfigNameOk

`func (o *AgentSkillReference) GetAiConfigNameOk() (*string, bool)`

GetAiConfigNameOk returns a tuple with the AiConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigName

`func (o *AgentSkillReference) SetAiConfigName(v string)`

SetAiConfigName sets AiConfigName field to given value.


### GetVariationId

`func (o *AgentSkillReference) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *AgentSkillReference) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *AgentSkillReference) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.


### GetVariationKey

`func (o *AgentSkillReference) GetVariationKey() string`

GetVariationKey returns the VariationKey field if non-nil, zero value otherwise.

### GetVariationKeyOk

`func (o *AgentSkillReference) GetVariationKeyOk() (*string, bool)`

GetVariationKeyOk returns a tuple with the VariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationKey

`func (o *AgentSkillReference) SetVariationKey(v string)`

SetVariationKey sets VariationKey field to given value.


### GetVariationName

`func (o *AgentSkillReference) GetVariationName() string`

GetVariationName returns the VariationName field if non-nil, zero value otherwise.

### GetVariationNameOk

`func (o *AgentSkillReference) GetVariationNameOk() (*string, bool)`

GetVariationNameOk returns a tuple with the VariationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationName

`func (o *AgentSkillReference) SetVariationName(v string)`

SetVariationName sets VariationName field to given value.


### GetResourceVersion

`func (o *AgentSkillReference) GetResourceVersion() int32`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *AgentSkillReference) GetResourceVersionOk() (*int32, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *AgentSkillReference) SetResourceVersion(v int32)`

SetResourceVersion sets ResourceVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


