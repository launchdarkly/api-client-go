# MonitoringPropertiesRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoringWindowMilliseconds** | Pointer to **int64** | The monitoring window in milliseconds | [optional] 
**OnRegression** | [**OnRegressionRep**](OnRegressionRep.md) |  | 
**OnProgression** | [**OnProgressionRep**](OnProgressionRep.md) |  | 
**Regressions** | [**[]RegressionRep**](RegressionRep.md) | The regressions on the experiment | 
**Progressions** | [**[]ProgressionRep**](ProgressionRep.md) | The progressions on the experiment | 

## Methods

### NewMonitoringPropertiesRep

`func NewMonitoringPropertiesRep(onRegression OnRegressionRep, onProgression OnProgressionRep, regressions []RegressionRep, progressions []ProgressionRep, ) *MonitoringPropertiesRep`

NewMonitoringPropertiesRep instantiates a new MonitoringPropertiesRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringPropertiesRepWithDefaults

`func NewMonitoringPropertiesRepWithDefaults() *MonitoringPropertiesRep`

NewMonitoringPropertiesRepWithDefaults instantiates a new MonitoringPropertiesRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoringWindowMilliseconds

`func (o *MonitoringPropertiesRep) GetMonitoringWindowMilliseconds() int64`

GetMonitoringWindowMilliseconds returns the MonitoringWindowMilliseconds field if non-nil, zero value otherwise.

### GetMonitoringWindowMillisecondsOk

`func (o *MonitoringPropertiesRep) GetMonitoringWindowMillisecondsOk() (*int64, bool)`

GetMonitoringWindowMillisecondsOk returns a tuple with the MonitoringWindowMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringWindowMilliseconds

`func (o *MonitoringPropertiesRep) SetMonitoringWindowMilliseconds(v int64)`

SetMonitoringWindowMilliseconds sets MonitoringWindowMilliseconds field to given value.

### HasMonitoringWindowMilliseconds

`func (o *MonitoringPropertiesRep) HasMonitoringWindowMilliseconds() bool`

HasMonitoringWindowMilliseconds returns a boolean if a field has been set.

### GetOnRegression

`func (o *MonitoringPropertiesRep) GetOnRegression() OnRegressionRep`

GetOnRegression returns the OnRegression field if non-nil, zero value otherwise.

### GetOnRegressionOk

`func (o *MonitoringPropertiesRep) GetOnRegressionOk() (*OnRegressionRep, bool)`

GetOnRegressionOk returns a tuple with the OnRegression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRegression

`func (o *MonitoringPropertiesRep) SetOnRegression(v OnRegressionRep)`

SetOnRegression sets OnRegression field to given value.


### GetOnProgression

`func (o *MonitoringPropertiesRep) GetOnProgression() OnProgressionRep`

GetOnProgression returns the OnProgression field if non-nil, zero value otherwise.

### GetOnProgressionOk

`func (o *MonitoringPropertiesRep) GetOnProgressionOk() (*OnProgressionRep, bool)`

GetOnProgressionOk returns a tuple with the OnProgression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnProgression

`func (o *MonitoringPropertiesRep) SetOnProgression(v OnProgressionRep)`

SetOnProgression sets OnProgression field to given value.


### GetRegressions

`func (o *MonitoringPropertiesRep) GetRegressions() []RegressionRep`

GetRegressions returns the Regressions field if non-nil, zero value otherwise.

### GetRegressionsOk

`func (o *MonitoringPropertiesRep) GetRegressionsOk() (*[]RegressionRep, bool)`

GetRegressionsOk returns a tuple with the Regressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegressions

`func (o *MonitoringPropertiesRep) SetRegressions(v []RegressionRep)`

SetRegressions sets Regressions field to given value.


### GetProgressions

`func (o *MonitoringPropertiesRep) GetProgressions() []ProgressionRep`

GetProgressions returns the Progressions field if non-nil, zero value otherwise.

### GetProgressionsOk

`func (o *MonitoringPropertiesRep) GetProgressionsOk() (*[]ProgressionRep, bool)`

GetProgressionsOk returns a tuple with the Progressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressions

`func (o *MonitoringPropertiesRep) SetProgressions(v []ProgressionRep)`

SetProgressions sets Progressions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


