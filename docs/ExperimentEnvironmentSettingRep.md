# ExperimentEnvironmentSettingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **int64** |  | [optional] 
**StopDate** | Pointer to **int64** |  | [optional] 
**EnabledPeriods** | Pointer to [**[]ExperimentEnabledPeriodRep**](ExperimentEnabledPeriodRep.md) |  | [optional] 

## Methods

### NewExperimentEnvironmentSettingRep

`func NewExperimentEnvironmentSettingRep() *ExperimentEnvironmentSettingRep`

NewExperimentEnvironmentSettingRep instantiates a new ExperimentEnvironmentSettingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentEnvironmentSettingRepWithDefaults

`func NewExperimentEnvironmentSettingRepWithDefaults() *ExperimentEnvironmentSettingRep`

NewExperimentEnvironmentSettingRepWithDefaults instantiates a new ExperimentEnvironmentSettingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ExperimentEnvironmentSettingRep) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExperimentEnvironmentSettingRep) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExperimentEnvironmentSettingRep) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExperimentEnvironmentSettingRep) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStopDate

`func (o *ExperimentEnvironmentSettingRep) GetStopDate() int64`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *ExperimentEnvironmentSettingRep) GetStopDateOk() (*int64, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDate

`func (o *ExperimentEnvironmentSettingRep) SetStopDate(v int64)`

SetStopDate sets StopDate field to given value.

### HasStopDate

`func (o *ExperimentEnvironmentSettingRep) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.

### GetEnabledPeriods

`func (o *ExperimentEnvironmentSettingRep) GetEnabledPeriods() []ExperimentEnabledPeriodRep`

GetEnabledPeriods returns the EnabledPeriods field if non-nil, zero value otherwise.

### GetEnabledPeriodsOk

`func (o *ExperimentEnvironmentSettingRep) GetEnabledPeriodsOk() (*[]ExperimentEnabledPeriodRep, bool)`

GetEnabledPeriodsOk returns a tuple with the EnabledPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledPeriods

`func (o *ExperimentEnvironmentSettingRep) SetEnabledPeriods(v []ExperimentEnabledPeriodRep)`

SetEnabledPeriods sets EnabledPeriods field to given value.

### HasEnabledPeriods

`func (o *ExperimentEnvironmentSettingRep) HasEnabledPeriods() bool`

HasEnabledPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


