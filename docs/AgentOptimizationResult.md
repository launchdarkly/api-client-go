# AgentOptimizationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Id** | **string** |  | 
**RunId** | **string** |  | 
**AgentOptimizationId** | **string** |  | 
**AgentOptimizationVersion** | **int32** |  | 
**Status** | [**AgentOptimizationResultStatus**](AgentOptimizationResultStatus.md) |  | 
**Activity** | [**AgentOptimizationResultActivity**](AgentOptimizationResultActivity.md) |  | 
**Iteration** | **int32** |  | 
**Instructions** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**UserInput** | **string** |  | 
**CompletionResponse** | Pointer to **string** |  | [optional] 
**Variation** | Pointer to **map[string]interface{}** |  | [optional] 
**Scores** | Pointer to **map[string]interface{}** |  | [optional] 
**GenerationTokens** | Pointer to **map[string]interface{}** |  | [optional] 
**EvaluationTokens** | Pointer to **map[string]interface{}** |  | [optional] 
**GenerationLatency** | Pointer to **int32** |  | [optional] 
**EvaluationLatencies** | Pointer to **map[string]interface{}** |  | [optional] 
**CompletedAt** | Pointer to **int64** |  | [optional] 
**CreatedVariationKey** | Pointer to **string** |  | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewAgentOptimizationResult

`func NewAgentOptimizationResult(id string, runId string, agentOptimizationId string, agentOptimizationVersion int32, status AgentOptimizationResultStatus, activity AgentOptimizationResultActivity, iteration int32, instructions string, userInput string, createdAt int64, updatedAt int64, ) *AgentOptimizationResult`

NewAgentOptimizationResult instantiates a new AgentOptimizationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationResultWithDefaults

`func NewAgentOptimizationResultWithDefaults() *AgentOptimizationResult`

NewAgentOptimizationResultWithDefaults instantiates a new AgentOptimizationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AgentOptimizationResult) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AgentOptimizationResult) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AgentOptimizationResult) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AgentOptimizationResult) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *AgentOptimizationResult) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentOptimizationResult) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentOptimizationResult) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentOptimizationResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AgentOptimizationResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentOptimizationResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentOptimizationResult) SetId(v string)`

SetId sets Id field to given value.


### GetRunId

`func (o *AgentOptimizationResult) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AgentOptimizationResult) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AgentOptimizationResult) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetAgentOptimizationId

`func (o *AgentOptimizationResult) GetAgentOptimizationId() string`

GetAgentOptimizationId returns the AgentOptimizationId field if non-nil, zero value otherwise.

### GetAgentOptimizationIdOk

`func (o *AgentOptimizationResult) GetAgentOptimizationIdOk() (*string, bool)`

GetAgentOptimizationIdOk returns a tuple with the AgentOptimizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOptimizationId

`func (o *AgentOptimizationResult) SetAgentOptimizationId(v string)`

SetAgentOptimizationId sets AgentOptimizationId field to given value.


### GetAgentOptimizationVersion

`func (o *AgentOptimizationResult) GetAgentOptimizationVersion() int32`

GetAgentOptimizationVersion returns the AgentOptimizationVersion field if non-nil, zero value otherwise.

### GetAgentOptimizationVersionOk

`func (o *AgentOptimizationResult) GetAgentOptimizationVersionOk() (*int32, bool)`

GetAgentOptimizationVersionOk returns a tuple with the AgentOptimizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOptimizationVersion

`func (o *AgentOptimizationResult) SetAgentOptimizationVersion(v int32)`

SetAgentOptimizationVersion sets AgentOptimizationVersion field to given value.


### GetStatus

`func (o *AgentOptimizationResult) GetStatus() AgentOptimizationResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentOptimizationResult) GetStatusOk() (*AgentOptimizationResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentOptimizationResult) SetStatus(v AgentOptimizationResultStatus)`

SetStatus sets Status field to given value.


### GetActivity

`func (o *AgentOptimizationResult) GetActivity() AgentOptimizationResultActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *AgentOptimizationResult) GetActivityOk() (*AgentOptimizationResultActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *AgentOptimizationResult) SetActivity(v AgentOptimizationResultActivity)`

SetActivity sets Activity field to given value.


### GetIteration

`func (o *AgentOptimizationResult) GetIteration() int32`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *AgentOptimizationResult) GetIterationOk() (*int32, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *AgentOptimizationResult) SetIteration(v int32)`

SetIteration sets Iteration field to given value.


### GetInstructions

`func (o *AgentOptimizationResult) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentOptimizationResult) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentOptimizationResult) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetParameters

`func (o *AgentOptimizationResult) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AgentOptimizationResult) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AgentOptimizationResult) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AgentOptimizationResult) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetUserInput

`func (o *AgentOptimizationResult) GetUserInput() string`

GetUserInput returns the UserInput field if non-nil, zero value otherwise.

### GetUserInputOk

`func (o *AgentOptimizationResult) GetUserInputOk() (*string, bool)`

GetUserInputOk returns a tuple with the UserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInput

`func (o *AgentOptimizationResult) SetUserInput(v string)`

SetUserInput sets UserInput field to given value.


### GetCompletionResponse

`func (o *AgentOptimizationResult) GetCompletionResponse() string`

GetCompletionResponse returns the CompletionResponse field if non-nil, zero value otherwise.

### GetCompletionResponseOk

`func (o *AgentOptimizationResult) GetCompletionResponseOk() (*string, bool)`

GetCompletionResponseOk returns a tuple with the CompletionResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionResponse

`func (o *AgentOptimizationResult) SetCompletionResponse(v string)`

SetCompletionResponse sets CompletionResponse field to given value.

### HasCompletionResponse

`func (o *AgentOptimizationResult) HasCompletionResponse() bool`

HasCompletionResponse returns a boolean if a field has been set.

### GetVariation

`func (o *AgentOptimizationResult) GetVariation() map[string]interface{}`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *AgentOptimizationResult) GetVariationOk() (*map[string]interface{}, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *AgentOptimizationResult) SetVariation(v map[string]interface{})`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *AgentOptimizationResult) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetScores

`func (o *AgentOptimizationResult) GetScores() map[string]interface{}`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *AgentOptimizationResult) GetScoresOk() (*map[string]interface{}, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *AgentOptimizationResult) SetScores(v map[string]interface{})`

SetScores sets Scores field to given value.

### HasScores

`func (o *AgentOptimizationResult) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetGenerationTokens

`func (o *AgentOptimizationResult) GetGenerationTokens() map[string]interface{}`

GetGenerationTokens returns the GenerationTokens field if non-nil, zero value otherwise.

### GetGenerationTokensOk

`func (o *AgentOptimizationResult) GetGenerationTokensOk() (*map[string]interface{}, bool)`

GetGenerationTokensOk returns a tuple with the GenerationTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTokens

`func (o *AgentOptimizationResult) SetGenerationTokens(v map[string]interface{})`

SetGenerationTokens sets GenerationTokens field to given value.

### HasGenerationTokens

`func (o *AgentOptimizationResult) HasGenerationTokens() bool`

HasGenerationTokens returns a boolean if a field has been set.

### GetEvaluationTokens

`func (o *AgentOptimizationResult) GetEvaluationTokens() map[string]interface{}`

GetEvaluationTokens returns the EvaluationTokens field if non-nil, zero value otherwise.

### GetEvaluationTokensOk

`func (o *AgentOptimizationResult) GetEvaluationTokensOk() (*map[string]interface{}, bool)`

GetEvaluationTokensOk returns a tuple with the EvaluationTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTokens

`func (o *AgentOptimizationResult) SetEvaluationTokens(v map[string]interface{})`

SetEvaluationTokens sets EvaluationTokens field to given value.

### HasEvaluationTokens

`func (o *AgentOptimizationResult) HasEvaluationTokens() bool`

HasEvaluationTokens returns a boolean if a field has been set.

### GetGenerationLatency

`func (o *AgentOptimizationResult) GetGenerationLatency() int32`

GetGenerationLatency returns the GenerationLatency field if non-nil, zero value otherwise.

### GetGenerationLatencyOk

`func (o *AgentOptimizationResult) GetGenerationLatencyOk() (*int32, bool)`

GetGenerationLatencyOk returns a tuple with the GenerationLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationLatency

`func (o *AgentOptimizationResult) SetGenerationLatency(v int32)`

SetGenerationLatency sets GenerationLatency field to given value.

### HasGenerationLatency

`func (o *AgentOptimizationResult) HasGenerationLatency() bool`

HasGenerationLatency returns a boolean if a field has been set.

### GetEvaluationLatencies

`func (o *AgentOptimizationResult) GetEvaluationLatencies() map[string]interface{}`

GetEvaluationLatencies returns the EvaluationLatencies field if non-nil, zero value otherwise.

### GetEvaluationLatenciesOk

`func (o *AgentOptimizationResult) GetEvaluationLatenciesOk() (*map[string]interface{}, bool)`

GetEvaluationLatenciesOk returns a tuple with the EvaluationLatencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationLatencies

`func (o *AgentOptimizationResult) SetEvaluationLatencies(v map[string]interface{})`

SetEvaluationLatencies sets EvaluationLatencies field to given value.

### HasEvaluationLatencies

`func (o *AgentOptimizationResult) HasEvaluationLatencies() bool`

HasEvaluationLatencies returns a boolean if a field has been set.

### GetCompletedAt

`func (o *AgentOptimizationResult) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *AgentOptimizationResult) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *AgentOptimizationResult) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *AgentOptimizationResult) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedVariationKey

`func (o *AgentOptimizationResult) GetCreatedVariationKey() string`

GetCreatedVariationKey returns the CreatedVariationKey field if non-nil, zero value otherwise.

### GetCreatedVariationKeyOk

`func (o *AgentOptimizationResult) GetCreatedVariationKeyOk() (*string, bool)`

GetCreatedVariationKeyOk returns a tuple with the CreatedVariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedVariationKey

`func (o *AgentOptimizationResult) SetCreatedVariationKey(v string)`

SetCreatedVariationKey sets CreatedVariationKey field to given value.

### HasCreatedVariationKey

`func (o *AgentOptimizationResult) HasCreatedVariationKey() bool`

HasCreatedVariationKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentOptimizationResult) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentOptimizationResult) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentOptimizationResult) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgentOptimizationResult) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentOptimizationResult) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentOptimizationResult) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


