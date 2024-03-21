# InsightsMetricScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** | The score for the metric | 
**AggregateOf** | Pointer to **[]string** | The keys of the metrics that were aggregated to calculate this score | [optional] 
**DiffVsLastPeriod** | Pointer to **int32** |  | [optional] 
**Indicator** | **string** |  | 
**IndicatorRange** | [**InsightsMetricIndicatorRange**](InsightsMetricIndicatorRange.md) |  | 
**LastPeriod** | Pointer to [**InsightsMetricScore**](InsightsMetricScore.md) |  | [optional] 

## Methods

### NewInsightsMetricScore

`func NewInsightsMetricScore(score int32, indicator string, indicatorRange InsightsMetricIndicatorRange, ) *InsightsMetricScore`

NewInsightsMetricScore instantiates a new InsightsMetricScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsMetricScoreWithDefaults

`func NewInsightsMetricScoreWithDefaults() *InsightsMetricScore`

NewInsightsMetricScoreWithDefaults instantiates a new InsightsMetricScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *InsightsMetricScore) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *InsightsMetricScore) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *InsightsMetricScore) SetScore(v int32)`

SetScore sets Score field to given value.


### GetAggregateOf

`func (o *InsightsMetricScore) GetAggregateOf() []string`

GetAggregateOf returns the AggregateOf field if non-nil, zero value otherwise.

### GetAggregateOfOk

`func (o *InsightsMetricScore) GetAggregateOfOk() (*[]string, bool)`

GetAggregateOfOk returns a tuple with the AggregateOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateOf

`func (o *InsightsMetricScore) SetAggregateOf(v []string)`

SetAggregateOf sets AggregateOf field to given value.

### HasAggregateOf

`func (o *InsightsMetricScore) HasAggregateOf() bool`

HasAggregateOf returns a boolean if a field has been set.

### GetDiffVsLastPeriod

`func (o *InsightsMetricScore) GetDiffVsLastPeriod() int32`

GetDiffVsLastPeriod returns the DiffVsLastPeriod field if non-nil, zero value otherwise.

### GetDiffVsLastPeriodOk

`func (o *InsightsMetricScore) GetDiffVsLastPeriodOk() (*int32, bool)`

GetDiffVsLastPeriodOk returns a tuple with the DiffVsLastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffVsLastPeriod

`func (o *InsightsMetricScore) SetDiffVsLastPeriod(v int32)`

SetDiffVsLastPeriod sets DiffVsLastPeriod field to given value.

### HasDiffVsLastPeriod

`func (o *InsightsMetricScore) HasDiffVsLastPeriod() bool`

HasDiffVsLastPeriod returns a boolean if a field has been set.

### GetIndicator

`func (o *InsightsMetricScore) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *InsightsMetricScore) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *InsightsMetricScore) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetIndicatorRange

`func (o *InsightsMetricScore) GetIndicatorRange() InsightsMetricIndicatorRange`

GetIndicatorRange returns the IndicatorRange field if non-nil, zero value otherwise.

### GetIndicatorRangeOk

`func (o *InsightsMetricScore) GetIndicatorRangeOk() (*InsightsMetricIndicatorRange, bool)`

GetIndicatorRangeOk returns a tuple with the IndicatorRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorRange

`func (o *InsightsMetricScore) SetIndicatorRange(v InsightsMetricIndicatorRange)`

SetIndicatorRange sets IndicatorRange field to given value.


### GetLastPeriod

`func (o *InsightsMetricScore) GetLastPeriod() InsightsMetricScore`

GetLastPeriod returns the LastPeriod field if non-nil, zero value otherwise.

### GetLastPeriodOk

`func (o *InsightsMetricScore) GetLastPeriodOk() (*InsightsMetricScore, bool)`

GetLastPeriodOk returns a tuple with the LastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPeriod

`func (o *InsightsMetricScore) SetLastPeriod(v InsightsMetricScore)`

SetLastPeriod sets LastPeriod field to given value.

### HasLastPeriod

`func (o *InsightsMetricScore) HasLastPeriod() bool`

HasLastPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


