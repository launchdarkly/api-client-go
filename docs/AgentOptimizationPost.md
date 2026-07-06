# AgentOptimizationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**AiConfigKey** | **string** |  | 
**MaxAttempts** | **int32** |  | 
**ModelChoices** | Pointer to **[]string** |  | [optional] 
**JudgeModel** | **string** |  | 
**VariableChoices** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AcceptanceStatements** | Pointer to [**[]AgentOptimizationAcceptanceStatement**](AgentOptimizationAcceptanceStatement.md) |  | [optional] 
**Judges** | Pointer to [**[]AgentOptimizationJudge**](AgentOptimizationJudge.md) |  | [optional] 
**UserInputOptions** | Pointer to **[]string** |  | [optional] 
**GroundTruthResponses** | Pointer to **[]string** |  | [optional] 
**MetricKey** | Pointer to **string** |  | [optional] 
**TokenLimit** | Pointer to **int32** |  | [optional] 
**VariationKey** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LatencyOptimization** | Pointer to **bool** |  | [optional] 
**TokenOptimization** | Pointer to **bool** |  | [optional] 
**AutoCommit** | Pointer to **bool** |  | [optional] 

## Methods

### NewAgentOptimizationPost

`func NewAgentOptimizationPost(key string, aiConfigKey string, maxAttempts int32, judgeModel string, ) *AgentOptimizationPost`

NewAgentOptimizationPost instantiates a new AgentOptimizationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationPostWithDefaults

`func NewAgentOptimizationPostWithDefaults() *AgentOptimizationPost`

NewAgentOptimizationPostWithDefaults instantiates a new AgentOptimizationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentOptimizationPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentOptimizationPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentOptimizationPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetAiConfigKey

`func (o *AgentOptimizationPost) GetAiConfigKey() string`

GetAiConfigKey returns the AiConfigKey field if non-nil, zero value otherwise.

### GetAiConfigKeyOk

`func (o *AgentOptimizationPost) GetAiConfigKeyOk() (*string, bool)`

GetAiConfigKeyOk returns a tuple with the AiConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigKey

`func (o *AgentOptimizationPost) SetAiConfigKey(v string)`

SetAiConfigKey sets AiConfigKey field to given value.


### GetMaxAttempts

`func (o *AgentOptimizationPost) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AgentOptimizationPost) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AgentOptimizationPost) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.


### GetModelChoices

`func (o *AgentOptimizationPost) GetModelChoices() []string`

GetModelChoices returns the ModelChoices field if non-nil, zero value otherwise.

### GetModelChoicesOk

`func (o *AgentOptimizationPost) GetModelChoicesOk() (*[]string, bool)`

GetModelChoicesOk returns a tuple with the ModelChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelChoices

`func (o *AgentOptimizationPost) SetModelChoices(v []string)`

SetModelChoices sets ModelChoices field to given value.

### HasModelChoices

`func (o *AgentOptimizationPost) HasModelChoices() bool`

HasModelChoices returns a boolean if a field has been set.

### GetJudgeModel

`func (o *AgentOptimizationPost) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *AgentOptimizationPost) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *AgentOptimizationPost) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.


### GetVariableChoices

`func (o *AgentOptimizationPost) GetVariableChoices() []map[string]interface{}`

GetVariableChoices returns the VariableChoices field if non-nil, zero value otherwise.

### GetVariableChoicesOk

`func (o *AgentOptimizationPost) GetVariableChoicesOk() (*[]map[string]interface{}, bool)`

GetVariableChoicesOk returns a tuple with the VariableChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableChoices

`func (o *AgentOptimizationPost) SetVariableChoices(v []map[string]interface{})`

SetVariableChoices sets VariableChoices field to given value.

### HasVariableChoices

`func (o *AgentOptimizationPost) HasVariableChoices() bool`

HasVariableChoices returns a boolean if a field has been set.

### GetAcceptanceStatements

`func (o *AgentOptimizationPost) GetAcceptanceStatements() []AgentOptimizationAcceptanceStatement`

GetAcceptanceStatements returns the AcceptanceStatements field if non-nil, zero value otherwise.

### GetAcceptanceStatementsOk

`func (o *AgentOptimizationPost) GetAcceptanceStatementsOk() (*[]AgentOptimizationAcceptanceStatement, bool)`

GetAcceptanceStatementsOk returns a tuple with the AcceptanceStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatements

`func (o *AgentOptimizationPost) SetAcceptanceStatements(v []AgentOptimizationAcceptanceStatement)`

SetAcceptanceStatements sets AcceptanceStatements field to given value.

### HasAcceptanceStatements

`func (o *AgentOptimizationPost) HasAcceptanceStatements() bool`

HasAcceptanceStatements returns a boolean if a field has been set.

### GetJudges

`func (o *AgentOptimizationPost) GetJudges() []AgentOptimizationJudge`

GetJudges returns the Judges field if non-nil, zero value otherwise.

### GetJudgesOk

`func (o *AgentOptimizationPost) GetJudgesOk() (*[]AgentOptimizationJudge, bool)`

GetJudgesOk returns a tuple with the Judges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudges

`func (o *AgentOptimizationPost) SetJudges(v []AgentOptimizationJudge)`

SetJudges sets Judges field to given value.

### HasJudges

`func (o *AgentOptimizationPost) HasJudges() bool`

HasJudges returns a boolean if a field has been set.

### GetUserInputOptions

`func (o *AgentOptimizationPost) GetUserInputOptions() []string`

GetUserInputOptions returns the UserInputOptions field if non-nil, zero value otherwise.

### GetUserInputOptionsOk

`func (o *AgentOptimizationPost) GetUserInputOptionsOk() (*[]string, bool)`

GetUserInputOptionsOk returns a tuple with the UserInputOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputOptions

`func (o *AgentOptimizationPost) SetUserInputOptions(v []string)`

SetUserInputOptions sets UserInputOptions field to given value.

### HasUserInputOptions

`func (o *AgentOptimizationPost) HasUserInputOptions() bool`

HasUserInputOptions returns a boolean if a field has been set.

### GetGroundTruthResponses

`func (o *AgentOptimizationPost) GetGroundTruthResponses() []string`

GetGroundTruthResponses returns the GroundTruthResponses field if non-nil, zero value otherwise.

### GetGroundTruthResponsesOk

`func (o *AgentOptimizationPost) GetGroundTruthResponsesOk() (*[]string, bool)`

GetGroundTruthResponsesOk returns a tuple with the GroundTruthResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroundTruthResponses

`func (o *AgentOptimizationPost) SetGroundTruthResponses(v []string)`

SetGroundTruthResponses sets GroundTruthResponses field to given value.

### HasGroundTruthResponses

`func (o *AgentOptimizationPost) HasGroundTruthResponses() bool`

HasGroundTruthResponses returns a boolean if a field has been set.

### GetMetricKey

`func (o *AgentOptimizationPost) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *AgentOptimizationPost) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *AgentOptimizationPost) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.

### HasMetricKey

`func (o *AgentOptimizationPost) HasMetricKey() bool`

HasMetricKey returns a boolean if a field has been set.

### GetTokenLimit

`func (o *AgentOptimizationPost) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *AgentOptimizationPost) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *AgentOptimizationPost) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *AgentOptimizationPost) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### GetVariationKey

`func (o *AgentOptimizationPost) GetVariationKey() string`

GetVariationKey returns the VariationKey field if non-nil, zero value otherwise.

### GetVariationKeyOk

`func (o *AgentOptimizationPost) GetVariationKeyOk() (*string, bool)`

GetVariationKeyOk returns a tuple with the VariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationKey

`func (o *AgentOptimizationPost) SetVariationKey(v string)`

SetVariationKey sets VariationKey field to given value.

### HasVariationKey

`func (o *AgentOptimizationPost) HasVariationKey() bool`

HasVariationKey returns a boolean if a field has been set.

### GetLabel

`func (o *AgentOptimizationPost) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentOptimizationPost) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentOptimizationPost) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentOptimizationPost) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLatencyOptimization

`func (o *AgentOptimizationPost) GetLatencyOptimization() bool`

GetLatencyOptimization returns the LatencyOptimization field if non-nil, zero value otherwise.

### GetLatencyOptimizationOk

`func (o *AgentOptimizationPost) GetLatencyOptimizationOk() (*bool, bool)`

GetLatencyOptimizationOk returns a tuple with the LatencyOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyOptimization

`func (o *AgentOptimizationPost) SetLatencyOptimization(v bool)`

SetLatencyOptimization sets LatencyOptimization field to given value.

### HasLatencyOptimization

`func (o *AgentOptimizationPost) HasLatencyOptimization() bool`

HasLatencyOptimization returns a boolean if a field has been set.

### GetTokenOptimization

`func (o *AgentOptimizationPost) GetTokenOptimization() bool`

GetTokenOptimization returns the TokenOptimization field if non-nil, zero value otherwise.

### GetTokenOptimizationOk

`func (o *AgentOptimizationPost) GetTokenOptimizationOk() (*bool, bool)`

GetTokenOptimizationOk returns a tuple with the TokenOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOptimization

`func (o *AgentOptimizationPost) SetTokenOptimization(v bool)`

SetTokenOptimization sets TokenOptimization field to given value.

### HasTokenOptimization

`func (o *AgentOptimizationPost) HasTokenOptimization() bool`

HasTokenOptimization returns a boolean if a field has been set.

### GetAutoCommit

`func (o *AgentOptimizationPost) GetAutoCommit() bool`

GetAutoCommit returns the AutoCommit field if non-nil, zero value otherwise.

### GetAutoCommitOk

`func (o *AgentOptimizationPost) GetAutoCommitOk() (*bool, bool)`

GetAutoCommitOk returns a tuple with the AutoCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommit

`func (o *AgentOptimizationPost) SetAutoCommit(v bool)`

SetAutoCommit sets AutoCommit field to given value.

### HasAutoCommit

`func (o *AgentOptimizationPost) HasAutoCommit() bool`

HasAutoCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


