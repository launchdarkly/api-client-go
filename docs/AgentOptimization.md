# AgentOptimization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Id** | **string** |  | 
**Key** | **string** |  | 
**AiConfigKey** | **string** |  | 
**MaxAttempts** | **int32** |  | 
**ModelChoices** | **[]string** |  | 
**JudgeModel** | **string** |  | 
**VariableChoices** | **[]map[string]interface{}** |  | 
**AcceptanceStatements** | [**[]AgentOptimizationAcceptanceStatement**](AgentOptimizationAcceptanceStatement.md) |  | 
**Judges** | [**[]AgentOptimizationJudge**](AgentOptimizationJudge.md) |  | 
**UserInputOptions** | **[]string** |  | 
**GroundTruthResponses** | **[]string** |  | 
**MetricKey** | Pointer to **string** |  | [optional] 
**TokenLimit** | Pointer to **int32** |  | [optional] 
**VariationKey** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LatencyOptimization** | Pointer to **bool** |  | [optional] 
**TokenOptimization** | Pointer to **bool** |  | [optional] 
**AutoCommit** | Pointer to **bool** |  | [optional] 
**Version** | **int32** |  | 
**CreatedAt** | **int64** |  | 

## Methods

### NewAgentOptimization

`func NewAgentOptimization(id string, key string, aiConfigKey string, maxAttempts int32, modelChoices []string, judgeModel string, variableChoices []map[string]interface{}, acceptanceStatements []AgentOptimizationAcceptanceStatement, judges []AgentOptimizationJudge, userInputOptions []string, groundTruthResponses []string, version int32, createdAt int64, ) *AgentOptimization`

NewAgentOptimization instantiates a new AgentOptimization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationWithDefaults

`func NewAgentOptimizationWithDefaults() *AgentOptimization`

NewAgentOptimizationWithDefaults instantiates a new AgentOptimization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AgentOptimization) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AgentOptimization) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AgentOptimization) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AgentOptimization) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *AgentOptimization) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentOptimization) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentOptimization) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentOptimization) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AgentOptimization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentOptimization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentOptimization) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *AgentOptimization) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentOptimization) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentOptimization) SetKey(v string)`

SetKey sets Key field to given value.


### GetAiConfigKey

`func (o *AgentOptimization) GetAiConfigKey() string`

GetAiConfigKey returns the AiConfigKey field if non-nil, zero value otherwise.

### GetAiConfigKeyOk

`func (o *AgentOptimization) GetAiConfigKeyOk() (*string, bool)`

GetAiConfigKeyOk returns a tuple with the AiConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigKey

`func (o *AgentOptimization) SetAiConfigKey(v string)`

SetAiConfigKey sets AiConfigKey field to given value.


### GetMaxAttempts

`func (o *AgentOptimization) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AgentOptimization) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AgentOptimization) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.


### GetModelChoices

`func (o *AgentOptimization) GetModelChoices() []string`

GetModelChoices returns the ModelChoices field if non-nil, zero value otherwise.

### GetModelChoicesOk

`func (o *AgentOptimization) GetModelChoicesOk() (*[]string, bool)`

GetModelChoicesOk returns a tuple with the ModelChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelChoices

`func (o *AgentOptimization) SetModelChoices(v []string)`

SetModelChoices sets ModelChoices field to given value.


### GetJudgeModel

`func (o *AgentOptimization) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *AgentOptimization) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *AgentOptimization) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.


### GetVariableChoices

`func (o *AgentOptimization) GetVariableChoices() []map[string]interface{}`

GetVariableChoices returns the VariableChoices field if non-nil, zero value otherwise.

### GetVariableChoicesOk

`func (o *AgentOptimization) GetVariableChoicesOk() (*[]map[string]interface{}, bool)`

GetVariableChoicesOk returns a tuple with the VariableChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableChoices

`func (o *AgentOptimization) SetVariableChoices(v []map[string]interface{})`

SetVariableChoices sets VariableChoices field to given value.


### GetAcceptanceStatements

`func (o *AgentOptimization) GetAcceptanceStatements() []AgentOptimizationAcceptanceStatement`

GetAcceptanceStatements returns the AcceptanceStatements field if non-nil, zero value otherwise.

### GetAcceptanceStatementsOk

`func (o *AgentOptimization) GetAcceptanceStatementsOk() (*[]AgentOptimizationAcceptanceStatement, bool)`

GetAcceptanceStatementsOk returns a tuple with the AcceptanceStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatements

`func (o *AgentOptimization) SetAcceptanceStatements(v []AgentOptimizationAcceptanceStatement)`

SetAcceptanceStatements sets AcceptanceStatements field to given value.


### GetJudges

`func (o *AgentOptimization) GetJudges() []AgentOptimizationJudge`

GetJudges returns the Judges field if non-nil, zero value otherwise.

### GetJudgesOk

`func (o *AgentOptimization) GetJudgesOk() (*[]AgentOptimizationJudge, bool)`

GetJudgesOk returns a tuple with the Judges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudges

`func (o *AgentOptimization) SetJudges(v []AgentOptimizationJudge)`

SetJudges sets Judges field to given value.


### GetUserInputOptions

`func (o *AgentOptimization) GetUserInputOptions() []string`

GetUserInputOptions returns the UserInputOptions field if non-nil, zero value otherwise.

### GetUserInputOptionsOk

`func (o *AgentOptimization) GetUserInputOptionsOk() (*[]string, bool)`

GetUserInputOptionsOk returns a tuple with the UserInputOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputOptions

`func (o *AgentOptimization) SetUserInputOptions(v []string)`

SetUserInputOptions sets UserInputOptions field to given value.


### GetGroundTruthResponses

`func (o *AgentOptimization) GetGroundTruthResponses() []string`

GetGroundTruthResponses returns the GroundTruthResponses field if non-nil, zero value otherwise.

### GetGroundTruthResponsesOk

`func (o *AgentOptimization) GetGroundTruthResponsesOk() (*[]string, bool)`

GetGroundTruthResponsesOk returns a tuple with the GroundTruthResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroundTruthResponses

`func (o *AgentOptimization) SetGroundTruthResponses(v []string)`

SetGroundTruthResponses sets GroundTruthResponses field to given value.


### GetMetricKey

`func (o *AgentOptimization) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *AgentOptimization) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *AgentOptimization) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.

### HasMetricKey

`func (o *AgentOptimization) HasMetricKey() bool`

HasMetricKey returns a boolean if a field has been set.

### GetTokenLimit

`func (o *AgentOptimization) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *AgentOptimization) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *AgentOptimization) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *AgentOptimization) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### GetVariationKey

`func (o *AgentOptimization) GetVariationKey() string`

GetVariationKey returns the VariationKey field if non-nil, zero value otherwise.

### GetVariationKeyOk

`func (o *AgentOptimization) GetVariationKeyOk() (*string, bool)`

GetVariationKeyOk returns a tuple with the VariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationKey

`func (o *AgentOptimization) SetVariationKey(v string)`

SetVariationKey sets VariationKey field to given value.

### HasVariationKey

`func (o *AgentOptimization) HasVariationKey() bool`

HasVariationKey returns a boolean if a field has been set.

### GetLabel

`func (o *AgentOptimization) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentOptimization) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentOptimization) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentOptimization) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLatencyOptimization

`func (o *AgentOptimization) GetLatencyOptimization() bool`

GetLatencyOptimization returns the LatencyOptimization field if non-nil, zero value otherwise.

### GetLatencyOptimizationOk

`func (o *AgentOptimization) GetLatencyOptimizationOk() (*bool, bool)`

GetLatencyOptimizationOk returns a tuple with the LatencyOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyOptimization

`func (o *AgentOptimization) SetLatencyOptimization(v bool)`

SetLatencyOptimization sets LatencyOptimization field to given value.

### HasLatencyOptimization

`func (o *AgentOptimization) HasLatencyOptimization() bool`

HasLatencyOptimization returns a boolean if a field has been set.

### GetTokenOptimization

`func (o *AgentOptimization) GetTokenOptimization() bool`

GetTokenOptimization returns the TokenOptimization field if non-nil, zero value otherwise.

### GetTokenOptimizationOk

`func (o *AgentOptimization) GetTokenOptimizationOk() (*bool, bool)`

GetTokenOptimizationOk returns a tuple with the TokenOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOptimization

`func (o *AgentOptimization) SetTokenOptimization(v bool)`

SetTokenOptimization sets TokenOptimization field to given value.

### HasTokenOptimization

`func (o *AgentOptimization) HasTokenOptimization() bool`

HasTokenOptimization returns a boolean if a field has been set.

### GetAutoCommit

`func (o *AgentOptimization) GetAutoCommit() bool`

GetAutoCommit returns the AutoCommit field if non-nil, zero value otherwise.

### GetAutoCommitOk

`func (o *AgentOptimization) GetAutoCommitOk() (*bool, bool)`

GetAutoCommitOk returns a tuple with the AutoCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommit

`func (o *AgentOptimization) SetAutoCommit(v bool)`

SetAutoCommit sets AutoCommit field to given value.

### HasAutoCommit

`func (o *AgentOptimization) HasAutoCommit() bool`

HasAutoCommit returns a boolean if a field has been set.

### GetVersion

`func (o *AgentOptimization) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentOptimization) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentOptimization) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *AgentOptimization) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentOptimization) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentOptimization) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


