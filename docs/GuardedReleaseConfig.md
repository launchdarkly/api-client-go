# GuardedReleaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutContextKindKey** | Pointer to **string** | Context kind key to use as the randomization unit for the rollout | [optional] 
**MinSampleSize** | Pointer to **int32** | The minimum number of samples required to make a decision | [optional] 
**RollbackOnRegression** | Pointer to **bool** | Whether to roll back on regression | [optional] 
**MetricKeys** | Pointer to **[]string** | List of metric keys | [optional] 
**MetricGroupKeys** | Pointer to **[]string** | List of metric group keys | [optional] 
**Stages** | Pointer to [**[]ReleasePolicyStage**](ReleasePolicyStage.md) | List of stages | [optional] 

## Methods

### NewGuardedReleaseConfig

`func NewGuardedReleaseConfig() *GuardedReleaseConfig`

NewGuardedReleaseConfig instantiates a new GuardedReleaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardedReleaseConfigWithDefaults

`func NewGuardedReleaseConfigWithDefaults() *GuardedReleaseConfig`

NewGuardedReleaseConfigWithDefaults instantiates a new GuardedReleaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutContextKindKey

`func (o *GuardedReleaseConfig) GetRolloutContextKindKey() string`

GetRolloutContextKindKey returns the RolloutContextKindKey field if non-nil, zero value otherwise.

### GetRolloutContextKindKeyOk

`func (o *GuardedReleaseConfig) GetRolloutContextKindKeyOk() (*string, bool)`

GetRolloutContextKindKeyOk returns a tuple with the RolloutContextKindKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutContextKindKey

`func (o *GuardedReleaseConfig) SetRolloutContextKindKey(v string)`

SetRolloutContextKindKey sets RolloutContextKindKey field to given value.

### HasRolloutContextKindKey

`func (o *GuardedReleaseConfig) HasRolloutContextKindKey() bool`

HasRolloutContextKindKey returns a boolean if a field has been set.

### GetMinSampleSize

`func (o *GuardedReleaseConfig) GetMinSampleSize() int32`

GetMinSampleSize returns the MinSampleSize field if non-nil, zero value otherwise.

### GetMinSampleSizeOk

`func (o *GuardedReleaseConfig) GetMinSampleSizeOk() (*int32, bool)`

GetMinSampleSizeOk returns a tuple with the MinSampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSampleSize

`func (o *GuardedReleaseConfig) SetMinSampleSize(v int32)`

SetMinSampleSize sets MinSampleSize field to given value.

### HasMinSampleSize

`func (o *GuardedReleaseConfig) HasMinSampleSize() bool`

HasMinSampleSize returns a boolean if a field has been set.

### GetRollbackOnRegression

`func (o *GuardedReleaseConfig) GetRollbackOnRegression() bool`

GetRollbackOnRegression returns the RollbackOnRegression field if non-nil, zero value otherwise.

### GetRollbackOnRegressionOk

`func (o *GuardedReleaseConfig) GetRollbackOnRegressionOk() (*bool, bool)`

GetRollbackOnRegressionOk returns a tuple with the RollbackOnRegression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackOnRegression

`func (o *GuardedReleaseConfig) SetRollbackOnRegression(v bool)`

SetRollbackOnRegression sets RollbackOnRegression field to given value.

### HasRollbackOnRegression

`func (o *GuardedReleaseConfig) HasRollbackOnRegression() bool`

HasRollbackOnRegression returns a boolean if a field has been set.

### GetMetricKeys

`func (o *GuardedReleaseConfig) GetMetricKeys() []string`

GetMetricKeys returns the MetricKeys field if non-nil, zero value otherwise.

### GetMetricKeysOk

`func (o *GuardedReleaseConfig) GetMetricKeysOk() (*[]string, bool)`

GetMetricKeysOk returns a tuple with the MetricKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKeys

`func (o *GuardedReleaseConfig) SetMetricKeys(v []string)`

SetMetricKeys sets MetricKeys field to given value.

### HasMetricKeys

`func (o *GuardedReleaseConfig) HasMetricKeys() bool`

HasMetricKeys returns a boolean if a field has been set.

### GetMetricGroupKeys

`func (o *GuardedReleaseConfig) GetMetricGroupKeys() []string`

GetMetricGroupKeys returns the MetricGroupKeys field if non-nil, zero value otherwise.

### GetMetricGroupKeysOk

`func (o *GuardedReleaseConfig) GetMetricGroupKeysOk() (*[]string, bool)`

GetMetricGroupKeysOk returns a tuple with the MetricGroupKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroupKeys

`func (o *GuardedReleaseConfig) SetMetricGroupKeys(v []string)`

SetMetricGroupKeys sets MetricGroupKeys field to given value.

### HasMetricGroupKeys

`func (o *GuardedReleaseConfig) HasMetricGroupKeys() bool`

HasMetricGroupKeys returns a boolean if a field has been set.

### GetStages

`func (o *GuardedReleaseConfig) GetStages() []ReleasePolicyStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *GuardedReleaseConfig) GetStagesOk() (*[]ReleasePolicyStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *GuardedReleaseConfig) SetStages(v []ReleasePolicyStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *GuardedReleaseConfig) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


