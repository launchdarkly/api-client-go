# AgentOptimizationResultPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AgentOptimizationResultStatus**](AgentOptimizationResultStatus.md) |  | [optional] 
**Activity** | Pointer to [**AgentOptimizationResultActivity**](AgentOptimizationResultActivity.md) |  | [optional] 
**CompletionResponse** | Pointer to **string** |  | [optional] 
**Variation** | Pointer to **map[string]interface{}** |  | [optional] 
**Scores** | Pointer to **map[string]interface{}** |  | [optional] 
**GenerationTokens** | Pointer to **map[string]interface{}** |  | [optional] 
**EvaluationTokens** | Pointer to **map[string]interface{}** |  | [optional] 
**GenerationLatency** | Pointer to **int32** |  | [optional] 
**EvaluationLatencies** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedVariationKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentOptimizationResultPatch

`func NewAgentOptimizationResultPatch() *AgentOptimizationResultPatch`

NewAgentOptimizationResultPatch instantiates a new AgentOptimizationResultPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationResultPatchWithDefaults

`func NewAgentOptimizationResultPatchWithDefaults() *AgentOptimizationResultPatch`

NewAgentOptimizationResultPatchWithDefaults instantiates a new AgentOptimizationResultPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AgentOptimizationResultPatch) GetStatus() AgentOptimizationResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentOptimizationResultPatch) GetStatusOk() (*AgentOptimizationResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentOptimizationResultPatch) SetStatus(v AgentOptimizationResultStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentOptimizationResultPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivity

`func (o *AgentOptimizationResultPatch) GetActivity() AgentOptimizationResultActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *AgentOptimizationResultPatch) GetActivityOk() (*AgentOptimizationResultActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *AgentOptimizationResultPatch) SetActivity(v AgentOptimizationResultActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *AgentOptimizationResultPatch) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetCompletionResponse

`func (o *AgentOptimizationResultPatch) GetCompletionResponse() string`

GetCompletionResponse returns the CompletionResponse field if non-nil, zero value otherwise.

### GetCompletionResponseOk

`func (o *AgentOptimizationResultPatch) GetCompletionResponseOk() (*string, bool)`

GetCompletionResponseOk returns a tuple with the CompletionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionResponse

`func (o *AgentOptimizationResultPatch) SetCompletionResponse(v string)`

SetCompletionResponse sets CompletionResponse field to given value.

### HasCompletionResponse

`func (o *AgentOptimizationResultPatch) HasCompletionResponse() bool`

HasCompletionResponse returns a boolean if a field has been set.

### GetVariation

`func (o *AgentOptimizationResultPatch) GetVariation() map[string]interface{}`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *AgentOptimizationResultPatch) GetVariationOk() (*map[string]interface{}, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *AgentOptimizationResultPatch) SetVariation(v map[string]interface{})`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *AgentOptimizationResultPatch) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetScores

`func (o *AgentOptimizationResultPatch) GetScores() map[string]interface{}`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *AgentOptimizationResultPatch) GetScoresOk() (*map[string]interface{}, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *AgentOptimizationResultPatch) SetScores(v map[string]interface{})`

SetScores sets Scores field to given value.

### HasScores

`func (o *AgentOptimizationResultPatch) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetGenerationTokens

`func (o *AgentOptimizationResultPatch) GetGenerationTokens() map[string]interface{}`

GetGenerationTokens returns the GenerationTokens field if non-nil, zero value otherwise.

### GetGenerationTokensOk

`func (o *AgentOptimizationResultPatch) GetGenerationTokensOk() (*map[string]interface{}, bool)`

GetGenerationTokensOk returns a tuple with the GenerationTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTokens

`func (o *AgentOptimizationResultPatch) SetGenerationTokens(v map[string]interface{})`

SetGenerationTokens sets GenerationTokens field to given value.

### HasGenerationTokens

`func (o *AgentOptimizationResultPatch) HasGenerationTokens() bool`

HasGenerationTokens returns a boolean if a field has been set.

### GetEvaluationTokens

`func (o *AgentOptimizationResultPatch) GetEvaluationTokens() map[string]interface{}`

GetEvaluationTokens returns the EvaluationTokens field if non-nil, zero value otherwise.

### GetEvaluationTokensOk

`func (o *AgentOptimizationResultPatch) GetEvaluationTokensOk() (*map[string]interface{}, bool)`

GetEvaluationTokensOk returns a tuple with the EvaluationTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTokens

`func (o *AgentOptimizationResultPatch) SetEvaluationTokens(v map[string]interface{})`

SetEvaluationTokens sets EvaluationTokens field to given value.

### HasEvaluationTokens

`func (o *AgentOptimizationResultPatch) HasEvaluationTokens() bool`

HasEvaluationTokens returns a boolean if a field has been set.

### GetGenerationLatency

`func (o *AgentOptimizationResultPatch) GetGenerationLatency() int32`

GetGenerationLatency returns the GenerationLatency field if non-nil, zero value otherwise.

### GetGenerationLatencyOk

`func (o *AgentOptimizationResultPatch) GetGenerationLatencyOk() (*int32, bool)`

GetGenerationLatencyOk returns a tuple with the GenerationLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationLatency

`func (o *AgentOptimizationResultPatch) SetGenerationLatency(v int32)`

SetGenerationLatency sets GenerationLatency field to given value.

### HasGenerationLatency

`func (o *AgentOptimizationResultPatch) HasGenerationLatency() bool`

HasGenerationLatency returns a boolean if a field has been set.

### GetEvaluationLatencies

`func (o *AgentOptimizationResultPatch) GetEvaluationLatencies() map[string]interface{}`

GetEvaluationLatencies returns the EvaluationLatencies field if non-nil, zero value otherwise.

### GetEvaluationLatenciesOk

`func (o *AgentOptimizationResultPatch) GetEvaluationLatenciesOk() (*map[string]interface{}, bool)`

GetEvaluationLatenciesOk returns a tuple with the EvaluationLatencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationLatencies

`func (o *AgentOptimizationResultPatch) SetEvaluationLatencies(v map[string]interface{})`

SetEvaluationLatencies sets EvaluationLatencies field to given value.

### HasEvaluationLatencies

`func (o *AgentOptimizationResultPatch) HasEvaluationLatencies() bool`

HasEvaluationLatencies returns a boolean if a field has been set.

### GetCreatedVariationKey

`func (o *AgentOptimizationResultPatch) GetCreatedVariationKey() string`

GetCreatedVariationKey returns the CreatedVariationKey field if non-nil, zero value otherwise.

### GetCreatedVariationKeyOk

`func (o *AgentOptimizationResultPatch) GetCreatedVariationKeyOk() (*string, bool)`

GetCreatedVariationKeyOk returns a tuple with the CreatedVariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedVariationKey

`func (o *AgentOptimizationResultPatch) SetCreatedVariationKey(v string)`

SetCreatedVariationKey sets CreatedVariationKey field to given value.

### HasCreatedVariationKey

`func (o *AgentOptimizationResultPatch) HasCreatedVariationKey() bool`

HasCreatedVariationKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


