# AgentOptimizationResultPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** |  | 
**AgentOptimizationVersion** | **int32** |  | 
**Iteration** | **int32** |  | 
**Instructions** | **string** |  | 
**UserInput** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAgentOptimizationResultPost

`func NewAgentOptimizationResultPost(runId string, agentOptimizationVersion int32, iteration int32, instructions string, userInput string, ) *AgentOptimizationResultPost`

NewAgentOptimizationResultPost instantiates a new AgentOptimizationResultPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationResultPostWithDefaults

`func NewAgentOptimizationResultPostWithDefaults() *AgentOptimizationResultPost`

NewAgentOptimizationResultPostWithDefaults instantiates a new AgentOptimizationResultPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *AgentOptimizationResultPost) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AgentOptimizationResultPost) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AgentOptimizationResultPost) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetAgentOptimizationVersion

`func (o *AgentOptimizationResultPost) GetAgentOptimizationVersion() int32`

GetAgentOptimizationVersion returns the AgentOptimizationVersion field if non-nil, zero value otherwise.

### GetAgentOptimizationVersionOk

`func (o *AgentOptimizationResultPost) GetAgentOptimizationVersionOk() (*int32, bool)`

GetAgentOptimizationVersionOk returns a tuple with the AgentOptimizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOptimizationVersion

`func (o *AgentOptimizationResultPost) SetAgentOptimizationVersion(v int32)`

SetAgentOptimizationVersion sets AgentOptimizationVersion field to given value.


### GetIteration

`func (o *AgentOptimizationResultPost) GetIteration() int32`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *AgentOptimizationResultPost) GetIterationOk() (*int32, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *AgentOptimizationResultPost) SetIteration(v int32)`

SetIteration sets Iteration field to given value.


### GetInstructions

`func (o *AgentOptimizationResultPost) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentOptimizationResultPost) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentOptimizationResultPost) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetUserInput

`func (o *AgentOptimizationResultPost) GetUserInput() string`

GetUserInput returns the UserInput field if non-nil, zero value otherwise.

### GetUserInputOk

`func (o *AgentOptimizationResultPost) GetUserInputOk() (*string, bool)`

GetUserInputOk returns a tuple with the UserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInput

`func (o *AgentOptimizationResultPost) SetUserInput(v string)`

SetUserInput sets UserInput field to given value.


### GetParameters

`func (o *AgentOptimizationResultPost) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AgentOptimizationResultPost) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AgentOptimizationResultPost) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AgentOptimizationResultPost) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


