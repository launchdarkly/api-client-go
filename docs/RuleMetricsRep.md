# RuleMetricsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MigrationMetricRep**](MigrationMetricRep.md) |  | [optional] 

## Methods

### NewRuleMetricsRep

`func NewRuleMetricsRep() *RuleMetricsRep`

NewRuleMetricsRep instantiates a new RuleMetricsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleMetricsRepWithDefaults

`func NewRuleMetricsRepWithDefaults() *RuleMetricsRep`

NewRuleMetricsRepWithDefaults instantiates a new RuleMetricsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *RuleMetricsRep) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleMetricsRep) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleMetricsRep) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleMetricsRep) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSummary

`func (o *RuleMetricsRep) GetSummary() map[string]interface{}`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RuleMetricsRep) GetSummaryOk() (*map[string]interface{}, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RuleMetricsRep) SetSummary(v map[string]interface{})`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RuleMetricsRep) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetData

`func (o *RuleMetricsRep) GetData() []MigrationMetricRep`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RuleMetricsRep) GetDataOk() (*[]MigrationMetricRep, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RuleMetricsRep) SetData(v []MigrationMetricRep)`

SetData sets Data field to given value.

### HasData

`func (o *RuleMetricsRep) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


