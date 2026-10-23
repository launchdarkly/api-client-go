# PipelineReleaseGuardianConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoringWindowMilliseconds** | **int64** | The monitoring window in milliseconds | 
**RolloutWeight** | **int32** | The rollout weight percentage | 
**RollbackOnRegression** | **bool** | Whether or not to roll back on regression | 
**RandomizationUnit** | Pointer to **string** | The randomization unit for the measured rollout | [optional] 

## Methods

### NewPipelineReleaseGuardianConfiguration

`func NewPipelineReleaseGuardianConfiguration(monitoringWindowMilliseconds int64, rolloutWeight int32, rollbackOnRegression bool, ) *PipelineReleaseGuardianConfiguration`

NewPipelineReleaseGuardianConfiguration instantiates a new PipelineReleaseGuardianConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReleaseGuardianConfigurationWithDefaults

`func NewPipelineReleaseGuardianConfigurationWithDefaults() *PipelineReleaseGuardianConfiguration`

NewPipelineReleaseGuardianConfigurationWithDefaults instantiates a new PipelineReleaseGuardianConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoringWindowMilliseconds

`func (o *PipelineReleaseGuardianConfiguration) GetMonitoringWindowMilliseconds() int64`

GetMonitoringWindowMilliseconds returns the MonitoringWindowMilliseconds field if non-nil, zero value otherwise.

### GetMonitoringWindowMillisecondsOk

`func (o *PipelineReleaseGuardianConfiguration) GetMonitoringWindowMillisecondsOk() (*int64, bool)`

GetMonitoringWindowMillisecondsOk returns a tuple with the MonitoringWindowMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringWindowMilliseconds

`func (o *PipelineReleaseGuardianConfiguration) SetMonitoringWindowMilliseconds(v int64)`

SetMonitoringWindowMilliseconds sets MonitoringWindowMilliseconds field to given value.


### GetRolloutWeight

`func (o *PipelineReleaseGuardianConfiguration) GetRolloutWeight() int32`

GetRolloutWeight returns the RolloutWeight field if non-nil, zero value otherwise.

### GetRolloutWeightOk

`func (o *PipelineReleaseGuardianConfiguration) GetRolloutWeightOk() (*int32, bool)`

GetRolloutWeightOk returns a tuple with the RolloutWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutWeight

`func (o *PipelineReleaseGuardianConfiguration) SetRolloutWeight(v int32)`

SetRolloutWeight sets RolloutWeight field to given value.


### GetRollbackOnRegression

`func (o *PipelineReleaseGuardianConfiguration) GetRollbackOnRegression() bool`

GetRollbackOnRegression returns the RollbackOnRegression field if non-nil, zero value otherwise.

### GetRollbackOnRegressionOk

`func (o *PipelineReleaseGuardianConfiguration) GetRollbackOnRegressionOk() (*bool, bool)`

GetRollbackOnRegressionOk returns a tuple with the RollbackOnRegression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackOnRegression

`func (o *PipelineReleaseGuardianConfiguration) SetRollbackOnRegression(v bool)`

SetRollbackOnRegression sets RollbackOnRegression field to given value.


### GetRandomizationUnit

`func (o *PipelineReleaseGuardianConfiguration) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *PipelineReleaseGuardianConfiguration) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *PipelineReleaseGuardianConfiguration) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *PipelineReleaseGuardianConfiguration) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


