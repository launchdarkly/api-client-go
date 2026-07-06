# QuickStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveConfigs** | **int32** |  | 
**ActiveExperiments** | **int32** |  | 
**AverageSatisfaction7D** | Pointer to **NullableFloat32** |  | [optional] 
**SpendMonthToDate** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewQuickStats

`func NewQuickStats(activeConfigs int32, activeExperiments int32, ) *QuickStats`

NewQuickStats instantiates a new QuickStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickStatsWithDefaults

`func NewQuickStatsWithDefaults() *QuickStats`

NewQuickStatsWithDefaults instantiates a new QuickStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveConfigs

`func (o *QuickStats) GetActiveConfigs() int32`

GetActiveConfigs returns the ActiveConfigs field if non-nil, zero value otherwise.

### GetActiveConfigsOk

`func (o *QuickStats) GetActiveConfigsOk() (*int32, bool)`

GetActiveConfigsOk returns a tuple with the ActiveConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConfigs

`func (o *QuickStats) SetActiveConfigs(v int32)`

SetActiveConfigs sets ActiveConfigs field to given value.


### GetActiveExperiments

`func (o *QuickStats) GetActiveExperiments() int32`

GetActiveExperiments returns the ActiveExperiments field if non-nil, zero value otherwise.

### GetActiveExperimentsOk

`func (o *QuickStats) GetActiveExperimentsOk() (*int32, bool)`

GetActiveExperimentsOk returns a tuple with the ActiveExperiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExperiments

`func (o *QuickStats) SetActiveExperiments(v int32)`

SetActiveExperiments sets ActiveExperiments field to given value.


### GetAverageSatisfaction7D

`func (o *QuickStats) GetAverageSatisfaction7D() float32`

GetAverageSatisfaction7D returns the AverageSatisfaction7D field if non-nil, zero value otherwise.

### GetAverageSatisfaction7DOk

`func (o *QuickStats) GetAverageSatisfaction7DOk() (*float32, bool)`

GetAverageSatisfaction7DOk returns a tuple with the AverageSatisfaction7D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSatisfaction7D

`func (o *QuickStats) SetAverageSatisfaction7D(v float32)`

SetAverageSatisfaction7D sets AverageSatisfaction7D field to given value.

### HasAverageSatisfaction7D

`func (o *QuickStats) HasAverageSatisfaction7D() bool`

HasAverageSatisfaction7D returns a boolean if a field has been set.

### SetAverageSatisfaction7DNil

`func (o *QuickStats) SetAverageSatisfaction7DNil(b bool)`

 SetAverageSatisfaction7DNil sets the value for AverageSatisfaction7D to be an explicit nil

### UnsetAverageSatisfaction7D
`func (o *QuickStats) UnsetAverageSatisfaction7D()`

UnsetAverageSatisfaction7D ensures that no value is present for AverageSatisfaction7D, not even an explicit nil
### GetSpendMonthToDate

`func (o *QuickStats) GetSpendMonthToDate() float64`

GetSpendMonthToDate returns the SpendMonthToDate field if non-nil, zero value otherwise.

### GetSpendMonthToDateOk

`func (o *QuickStats) GetSpendMonthToDateOk() (*float64, bool)`

GetSpendMonthToDateOk returns a tuple with the SpendMonthToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMonthToDate

`func (o *QuickStats) SetSpendMonthToDate(v float64)`

SetSpendMonthToDate sets SpendMonthToDate field to given value.

### HasSpendMonthToDate

`func (o *QuickStats) HasSpendMonthToDate() bool`

HasSpendMonthToDate returns a boolean if a field has been set.

### SetSpendMonthToDateNil

`func (o *QuickStats) SetSpendMonthToDateNil(b bool)`

 SetSpendMonthToDateNil sets the value for SpendMonthToDate to be an explicit nil

### UnsetSpendMonthToDate
`func (o *QuickStats) UnsetSpendMonthToDate()`

UnsetSpendMonthToDate ensures that no value is present for SpendMonthToDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


