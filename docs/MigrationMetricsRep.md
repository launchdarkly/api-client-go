# MigrationMetricsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overview** | [**OverviewMetricsRep**](OverviewMetricsRep.md) |  | 
**Rules** | Pointer to [**[]RuleMetricsRep**](RuleMetricsRep.md) |  | [optional] 

## Methods

### NewMigrationMetricsRep

`func NewMigrationMetricsRep(overview OverviewMetricsRep, ) *MigrationMetricsRep`

NewMigrationMetricsRep instantiates a new MigrationMetricsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationMetricsRepWithDefaults

`func NewMigrationMetricsRepWithDefaults() *MigrationMetricsRep`

NewMigrationMetricsRepWithDefaults instantiates a new MigrationMetricsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverview

`func (o *MigrationMetricsRep) GetOverview() OverviewMetricsRep`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *MigrationMetricsRep) GetOverviewOk() (*OverviewMetricsRep, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *MigrationMetricsRep) SetOverview(v OverviewMetricsRep)`

SetOverview sets Overview field to given value.


### GetRules

`func (o *MigrationMetricsRep) GetRules() []RuleMetricsRep`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *MigrationMetricsRep) GetRulesOk() (*[]RuleMetricsRep, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *MigrationMetricsRep) SetRules(v []RuleMetricsRep)`

SetRules sets Rules field to given value.

### HasRules

`func (o *MigrationMetricsRep) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


