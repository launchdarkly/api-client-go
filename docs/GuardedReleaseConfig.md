# GuardedReleaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinSampleSize** | Pointer to **int32** | The minimum number of samples required to make a decision | [optional] 
**RollbackOnRegression** | **bool** | Whether to roll back on regression | 

## Methods

### NewGuardedReleaseConfig

`func NewGuardedReleaseConfig(rollbackOnRegression bool, ) *GuardedReleaseConfig`

NewGuardedReleaseConfig instantiates a new GuardedReleaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardedReleaseConfigWithDefaults

`func NewGuardedReleaseConfigWithDefaults() *GuardedReleaseConfig`

NewGuardedReleaseConfigWithDefaults instantiates a new GuardedReleaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


