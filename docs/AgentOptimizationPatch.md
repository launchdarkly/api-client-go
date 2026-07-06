# AgentOptimizationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAttempts** | Pointer to **int32** |  | [optional] 
**ModelChoices** | Pointer to **[]string** |  | [optional] 
**JudgeModel** | Pointer to **string** |  | [optional] 
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

### NewAgentOptimizationPatch

`func NewAgentOptimizationPatch() *AgentOptimizationPatch`

NewAgentOptimizationPatch instantiates a new AgentOptimizationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationPatchWithDefaults

`func NewAgentOptimizationPatchWithDefaults() *AgentOptimizationPatch`

NewAgentOptimizationPatchWithDefaults instantiates a new AgentOptimizationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAttempts

`func (o *AgentOptimizationPatch) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AgentOptimizationPatch) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AgentOptimizationPatch) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *AgentOptimizationPatch) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetModelChoices

`func (o *AgentOptimizationPatch) GetModelChoices() []string`

GetModelChoices returns the ModelChoices field if non-nil, zero value otherwise.

### GetModelChoicesOk

`func (o *AgentOptimizationPatch) GetModelChoicesOk() (*[]string, bool)`

GetModelChoicesOk returns a tuple with the ModelChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelChoices

`func (o *AgentOptimizationPatch) SetModelChoices(v []string)`

SetModelChoices sets ModelChoices field to given value.

### HasModelChoices

`func (o *AgentOptimizationPatch) HasModelChoices() bool`

HasModelChoices returns a boolean if a field has been set.

### GetJudgeModel

`func (o *AgentOptimizationPatch) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *AgentOptimizationPatch) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *AgentOptimizationPatch) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *AgentOptimizationPatch) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.

### GetVariableChoices

`func (o *AgentOptimizationPatch) GetVariableChoices() []map[string]interface{}`

GetVariableChoices returns the VariableChoices field if non-nil, zero value otherwise.

### GetVariableChoicesOk

`func (o *AgentOptimizationPatch) GetVariableChoicesOk() (*[]map[string]interface{}, bool)`

GetVariableChoicesOk returns a tuple with the VariableChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableChoices

`func (o *AgentOptimizationPatch) SetVariableChoices(v []map[string]interface{})`

SetVariableChoices sets VariableChoices field to given value.

### HasVariableChoices

`func (o *AgentOptimizationPatch) HasVariableChoices() bool`

HasVariableChoices returns a boolean if a field has been set.

### GetAcceptanceStatements

`func (o *AgentOptimizationPatch) GetAcceptanceStatements() []AgentOptimizationAcceptanceStatement`

GetAcceptanceStatements returns the AcceptanceStatements field if non-nil, zero value otherwise.

### GetAcceptanceStatementsOk

`func (o *AgentOptimizationPatch) GetAcceptanceStatementsOk() (*[]AgentOptimizationAcceptanceStatement, bool)`

GetAcceptanceStatementsOk returns a tuple with the AcceptanceStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatements

`func (o *AgentOptimizationPatch) SetAcceptanceStatements(v []AgentOptimizationAcceptanceStatement)`

SetAcceptanceStatements sets AcceptanceStatements field to given value.

### HasAcceptanceStatements

`func (o *AgentOptimizationPatch) HasAcceptanceStatements() bool`

HasAcceptanceStatements returns a boolean if a field has been set.

### GetJudges

`func (o *AgentOptimizationPatch) GetJudges() []AgentOptimizationJudge`

GetJudges returns the Judges field if non-nil, zero value otherwise.

### GetJudgesOk

`func (o *AgentOptimizationPatch) GetJudgesOk() (*[]AgentOptimizationJudge, bool)`

GetJudgesOk returns a tuple with the Judges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudges

`func (o *AgentOptimizationPatch) SetJudges(v []AgentOptimizationJudge)`

SetJudges sets Judges field to given value.

### HasJudges

`func (o *AgentOptimizationPatch) HasJudges() bool`

HasJudges returns a boolean if a field has been set.

### GetUserInputOptions

`func (o *AgentOptimizationPatch) GetUserInputOptions() []string`

GetUserInputOptions returns the UserInputOptions field if non-nil, zero value otherwise.

### GetUserInputOptionsOk

`func (o *AgentOptimizationPatch) GetUserInputOptionsOk() (*[]string, bool)`

GetUserInputOptionsOk returns a tuple with the UserInputOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputOptions

`func (o *AgentOptimizationPatch) SetUserInputOptions(v []string)`

SetUserInputOptions sets UserInputOptions field to given value.

### HasUserInputOptions

`func (o *AgentOptimizationPatch) HasUserInputOptions() bool`

HasUserInputOptions returns a boolean if a field has been set.

### GetGroundTruthResponses

`func (o *AgentOptimizationPatch) GetGroundTruthResponses() []string`

GetGroundTruthResponses returns the GroundTruthResponses field if non-nil, zero value otherwise.

### GetGroundTruthResponsesOk

`func (o *AgentOptimizationPatch) GetGroundTruthResponsesOk() (*[]string, bool)`

GetGroundTruthResponsesOk returns a tuple with the GroundTruthResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroundTruthResponses

`func (o *AgentOptimizationPatch) SetGroundTruthResponses(v []string)`

SetGroundTruthResponses sets GroundTruthResponses field to given value.

### HasGroundTruthResponses

`func (o *AgentOptimizationPatch) HasGroundTruthResponses() bool`

HasGroundTruthResponses returns a boolean if a field has been set.

### GetMetricKey

`func (o *AgentOptimizationPatch) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *AgentOptimizationPatch) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *AgentOptimizationPatch) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.

### HasMetricKey

`func (o *AgentOptimizationPatch) HasMetricKey() bool`

HasMetricKey returns a boolean if a field has been set.

### GetTokenLimit

`func (o *AgentOptimizationPatch) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *AgentOptimizationPatch) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *AgentOptimizationPatch) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *AgentOptimizationPatch) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### GetVariationKey

`func (o *AgentOptimizationPatch) GetVariationKey() string`

GetVariationKey returns the VariationKey field if non-nil, zero value otherwise.

### GetVariationKeyOk

`func (o *AgentOptimizationPatch) GetVariationKeyOk() (*string, bool)`

GetVariationKeyOk returns a tuple with the VariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationKey

`func (o *AgentOptimizationPatch) SetVariationKey(v string)`

SetVariationKey sets VariationKey field to given value.

### HasVariationKey

`func (o *AgentOptimizationPatch) HasVariationKey() bool`

HasVariationKey returns a boolean if a field has been set.

### GetLabel

`func (o *AgentOptimizationPatch) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentOptimizationPatch) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentOptimizationPatch) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentOptimizationPatch) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLatencyOptimization

`func (o *AgentOptimizationPatch) GetLatencyOptimization() bool`

GetLatencyOptimization returns the LatencyOptimization field if non-nil, zero value otherwise.

### GetLatencyOptimizationOk

`func (o *AgentOptimizationPatch) GetLatencyOptimizationOk() (*bool, bool)`

GetLatencyOptimizationOk returns a tuple with the LatencyOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyOptimization

`func (o *AgentOptimizationPatch) SetLatencyOptimization(v bool)`

SetLatencyOptimization sets LatencyOptimization field to given value.

### HasLatencyOptimization

`func (o *AgentOptimizationPatch) HasLatencyOptimization() bool`

HasLatencyOptimization returns a boolean if a field has been set.

### GetTokenOptimization

`func (o *AgentOptimizationPatch) GetTokenOptimization() bool`

GetTokenOptimization returns the TokenOptimization field if non-nil, zero value otherwise.

### GetTokenOptimizationOk

`func (o *AgentOptimizationPatch) GetTokenOptimizationOk() (*bool, bool)`

GetTokenOptimizationOk returns a tuple with the TokenOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOptimization

`func (o *AgentOptimizationPatch) SetTokenOptimization(v bool)`

SetTokenOptimization sets TokenOptimization field to given value.

### HasTokenOptimization

`func (o *AgentOptimizationPatch) HasTokenOptimization() bool`

HasTokenOptimization returns a boolean if a field has been set.

### GetAutoCommit

`func (o *AgentOptimizationPatch) GetAutoCommit() bool`

GetAutoCommit returns the AutoCommit field if non-nil, zero value otherwise.

### GetAutoCommitOk

`func (o *AgentOptimizationPatch) GetAutoCommitOk() (*bool, bool)`

GetAutoCommitOk returns a tuple with the AutoCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommit

`func (o *AgentOptimizationPatch) SetAutoCommit(v bool)`

SetAutoCommit sets AutoCommit field to given value.

### HasAutoCommit

`func (o *AgentOptimizationPatch) HasAutoCommit() bool`

HasAutoCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


