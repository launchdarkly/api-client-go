# ReleaseGuardianConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoringWindowMilliseconds** | **int64** | The monitoring window in milliseconds | 
**RolloutWeight** | **int32** | The rollout weight percentage | 
**RollbackOnRegression** | **bool** | Whether or not to roll back on regression | 
**RandomizationUnit** | Pointer to **string** | The randomization unit for the measured rollout | [optional] 

## Methods

### NewReleaseGuardianConfiguration

`func NewReleaseGuardianConfiguration(monitoringWindowMilliseconds int64, rolloutWeight int32, rollbackOnRegression bool, ) *ReleaseGuardianConfiguration`

NewReleaseGuardianConfiguration instantiates a new ReleaseGuardianConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseGuardianConfigurationWithDefaults

`func NewReleaseGuardianConfigurationWithDefaults() *ReleaseGuardianConfiguration`

NewReleaseGuardianConfigurationWithDefaults instantiates a new ReleaseGuardianConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoringWindowMilliseconds

`func (o *ReleaseGuardianConfiguration) GetMonitoringWindowMilliseconds() int64`

GetMonitoringWindowMilliseconds returns the MonitoringWindowMilliseconds field if non-nil, zero value otherwise.

### GetMonitoringWindowMillisecondsOk

`func (o *ReleaseGuardianConfiguration) GetMonitoringWindowMillisecondsOk() (*int64, bool)`

GetMonitoringWindowMillisecondsOk returns a tuple with the MonitoringWindowMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringWindowMilliseconds

`func (o *ReleaseGuardianConfiguration) SetMonitoringWindowMilliseconds(v int64)`

SetMonitoringWindowMilliseconds sets MonitoringWindowMilliseconds field to given value.


### GetRolloutWeight

`func (o *ReleaseGuardianConfiguration) GetRolloutWeight() int32`

GetRolloutWeight returns the RolloutWeight field if non-nil, zero value otherwise.

### GetRolloutWeightOk

`func (o *ReleaseGuardianConfiguration) GetRolloutWeightOk() (*int32, bool)`

GetRolloutWeightOk returns a tuple with the RolloutWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutWeight

`func (o *ReleaseGuardianConfiguration) SetRolloutWeight(v int32)`

SetRolloutWeight sets RolloutWeight field to given value.


### GetRollbackOnRegression

`func (o *ReleaseGuardianConfiguration) GetRollbackOnRegression() bool`

GetRollbackOnRegression returns the RollbackOnRegression field if non-nil, zero value otherwise.

### GetRollbackOnRegressionOk

`func (o *ReleaseGuardianConfiguration) GetRollbackOnRegressionOk() (*bool, bool)`

GetRollbackOnRegressionOk returns a tuple with the RollbackOnRegression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackOnRegression

`func (o *ReleaseGuardianConfiguration) SetRollbackOnRegression(v bool)`

SetRollbackOnRegression sets RollbackOnRegression field to given value.


### GetRandomizationUnit

`func (o *ReleaseGuardianConfiguration) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *ReleaseGuardianConfiguration) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *ReleaseGuardianConfiguration) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *ReleaseGuardianConfiguration) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


