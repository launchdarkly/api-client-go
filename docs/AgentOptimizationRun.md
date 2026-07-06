# AgentOptimizationRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** |  | 
**OptimizationKey** | **string** |  | 
**AgentOptimizationId** | **string** |  | 
**AgentOptimizationVersion** | **int32** |  | 
**Status** | [**AgentOptimizationResultStatus**](AgentOptimizationResultStatus.md) |  | 
**Activity** | [**AgentOptimizationResultActivity**](AgentOptimizationResultActivity.md) |  | 
**CreatedAt** | **int64** |  | 
**CompletedAt** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewAgentOptimizationRun

`func NewAgentOptimizationRun(runId string, optimizationKey string, agentOptimizationId string, agentOptimizationVersion int32, status AgentOptimizationResultStatus, activity AgentOptimizationResultActivity, createdAt int64, ) *AgentOptimizationRun`

NewAgentOptimizationRun instantiates a new AgentOptimizationRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationRunWithDefaults

`func NewAgentOptimizationRunWithDefaults() *AgentOptimizationRun`

NewAgentOptimizationRunWithDefaults instantiates a new AgentOptimizationRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *AgentOptimizationRun) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AgentOptimizationRun) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AgentOptimizationRun) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetOptimizationKey

`func (o *AgentOptimizationRun) GetOptimizationKey() string`

GetOptimizationKey returns the OptimizationKey field if non-nil, zero value otherwise.

### GetOptimizationKeyOk

`func (o *AgentOptimizationRun) GetOptimizationKeyOk() (*string, bool)`

GetOptimizationKeyOk returns a tuple with the OptimizationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationKey

`func (o *AgentOptimizationRun) SetOptimizationKey(v string)`

SetOptimizationKey sets OptimizationKey field to given value.


### GetAgentOptimizationId

`func (o *AgentOptimizationRun) GetAgentOptimizationId() string`

GetAgentOptimizationId returns the AgentOptimizationId field if non-nil, zero value otherwise.

### GetAgentOptimizationIdOk

`func (o *AgentOptimizationRun) GetAgentOptimizationIdOk() (*string, bool)`

GetAgentOptimizationIdOk returns a tuple with the AgentOptimizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOptimizationId

`func (o *AgentOptimizationRun) SetAgentOptimizationId(v string)`

SetAgentOptimizationId sets AgentOptimizationId field to given value.


### GetAgentOptimizationVersion

`func (o *AgentOptimizationRun) GetAgentOptimizationVersion() int32`

GetAgentOptimizationVersion returns the AgentOptimizationVersion field if non-nil, zero value otherwise.

### GetAgentOptimizationVersionOk

`func (o *AgentOptimizationRun) GetAgentOptimizationVersionOk() (*int32, bool)`

GetAgentOptimizationVersionOk returns a tuple with the AgentOptimizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOptimizationVersion

`func (o *AgentOptimizationRun) SetAgentOptimizationVersion(v int32)`

SetAgentOptimizationVersion sets AgentOptimizationVersion field to given value.


### GetStatus

`func (o *AgentOptimizationRun) GetStatus() AgentOptimizationResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentOptimizationRun) GetStatusOk() (*AgentOptimizationResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentOptimizationRun) SetStatus(v AgentOptimizationResultStatus)`

SetStatus sets Status field to given value.


### GetActivity

`func (o *AgentOptimizationRun) GetActivity() AgentOptimizationResultActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *AgentOptimizationRun) GetActivityOk() (*AgentOptimizationResultActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *AgentOptimizationRun) SetActivity(v AgentOptimizationResultActivity)`

SetActivity sets Activity field to given value.


### GetCreatedAt

`func (o *AgentOptimizationRun) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentOptimizationRun) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentOptimizationRun) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *AgentOptimizationRun) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *AgentOptimizationRun) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *AgentOptimizationRun) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *AgentOptimizationRun) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *AgentOptimizationRun) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *AgentOptimizationRun) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


