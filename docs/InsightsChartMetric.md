# InsightsChartMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indicator** | **string** | Metric indicator tier | 
**Value** | **float32** | Metric value | 
**Unit** | **string** | Metric unit | 
**Modifier** | **string** | Metric modifier | 
**Tiers** | [**[]InsightsMetricTierDefinition**](InsightsMetricTierDefinition.md) | Metric indicator tiers | 

## Methods

### NewInsightsChartMetric

`func NewInsightsChartMetric(indicator string, value float32, unit string, modifier string, tiers []InsightsMetricTierDefinition, ) *InsightsChartMetric`

NewInsightsChartMetric instantiates a new InsightsChartMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsChartMetricWithDefaults

`func NewInsightsChartMetricWithDefaults() *InsightsChartMetric`

NewInsightsChartMetricWithDefaults instantiates a new InsightsChartMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicator

`func (o *InsightsChartMetric) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *InsightsChartMetric) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *InsightsChartMetric) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetValue

`func (o *InsightsChartMetric) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InsightsChartMetric) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InsightsChartMetric) SetValue(v float32)`

SetValue sets Value field to given value.


### GetUnit

`func (o *InsightsChartMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InsightsChartMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InsightsChartMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetModifier

`func (o *InsightsChartMetric) GetModifier() string`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *InsightsChartMetric) GetModifierOk() (*string, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *InsightsChartMetric) SetModifier(v string)`

SetModifier sets Modifier field to given value.


### GetTiers

`func (o *InsightsChartMetric) GetTiers() []InsightsMetricTierDefinition`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *InsightsChartMetric) GetTiersOk() (*[]InsightsMetricTierDefinition, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *InsightsChartMetric) SetTiers(v []InsightsMetricTierDefinition)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


