# PipelineReleaseGuardianConfigurationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoringWindowMilliseconds** | Pointer to **int64** | The monitoring window in milliseconds | [optional] 
**RolloutWeight** | Pointer to **int32** | The rollout weight | [optional] 
**RollbackOnRegression** | Pointer to **bool** | Whether or not to rollback on regression | [optional] 
**RandomizationUnit** | Pointer to **string** | The randomization unit for the measured rollout | [optional] 

## Methods

### NewPipelineReleaseGuardianConfigurationInput

`func NewPipelineReleaseGuardianConfigurationInput() *PipelineReleaseGuardianConfigurationInput`

NewPipelineReleaseGuardianConfigurationInput instantiates a new PipelineReleaseGuardianConfigurationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReleaseGuardianConfigurationInputWithDefaults

`func NewPipelineReleaseGuardianConfigurationInputWithDefaults() *PipelineReleaseGuardianConfigurationInput`

NewPipelineReleaseGuardianConfigurationInputWithDefaults instantiates a new PipelineReleaseGuardianConfigurationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoringWindowMilliseconds

`func (o *PipelineReleaseGuardianConfigurationInput) GetMonitoringWindowMilliseconds() int64`

GetMonitoringWindowMilliseconds returns the MonitoringWindowMilliseconds field if non-nil, zero value otherwise.

### GetMonitoringWindowMillisecondsOk

`func (o *PipelineReleaseGuardianConfigurationInput) GetMonitoringWindowMillisecondsOk() (*int64, bool)`

GetMonitoringWindowMillisecondsOk returns a tuple with the MonitoringWindowMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringWindowMilliseconds

`func (o *PipelineReleaseGuardianConfigurationInput) SetMonitoringWindowMilliseconds(v int64)`

SetMonitoringWindowMilliseconds sets MonitoringWindowMilliseconds field to given value.

### HasMonitoringWindowMilliseconds

`func (o *PipelineReleaseGuardianConfigurationInput) HasMonitoringWindowMilliseconds() bool`

HasMonitoringWindowMilliseconds returns a boolean if a field has been set.

### GetRolloutWeight

`func (o *PipelineReleaseGuardianConfigurationInput) GetRolloutWeight() int32`

GetRolloutWeight returns the RolloutWeight field if non-nil, zero value otherwise.

### GetRolloutWeightOk

`func (o *PipelineReleaseGuardianConfigurationInput) GetRolloutWeightOk() (*int32, bool)`

GetRolloutWeightOk returns a tuple with the RolloutWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutWeight

`func (o *PipelineReleaseGuardianConfigurationInput) SetRolloutWeight(v int32)`

SetRolloutWeight sets RolloutWeight field to given value.

### HasRolloutWeight

`func (o *PipelineReleaseGuardianConfigurationInput) HasRolloutWeight() bool`

HasRolloutWeight returns a boolean if a field has been set.

### GetRollbackOnRegression

`func (o *PipelineReleaseGuardianConfigurationInput) GetRollbackOnRegression() bool`

GetRollbackOnRegression returns the RollbackOnRegression field if non-nil, zero value otherwise.

### GetRollbackOnRegressionOk

`func (o *PipelineReleaseGuardianConfigurationInput) GetRollbackOnRegressionOk() (*bool, bool)`

GetRollbackOnRegressionOk returns a tuple with the RollbackOnRegression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackOnRegression

`func (o *PipelineReleaseGuardianConfigurationInput) SetRollbackOnRegression(v bool)`

SetRollbackOnRegression sets RollbackOnRegression field to given value.

### HasRollbackOnRegression

`func (o *PipelineReleaseGuardianConfigurationInput) HasRollbackOnRegression() bool`

HasRollbackOnRegression returns a boolean if a field has been set.

### GetRandomizationUnit

`func (o *PipelineReleaseGuardianConfigurationInput) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *PipelineReleaseGuardianConfigurationInput) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *PipelineReleaseGuardianConfigurationInput) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *PipelineReleaseGuardianConfigurationInput) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


