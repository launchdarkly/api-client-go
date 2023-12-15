# MetricAnalyticsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** |  | 
**MetricName** | **string** |  | 
**GroupBy** | Pointer to [**MetricQueryGroupByRep**](MetricQueryGroupByRep.md) |  | [optional] 
**GroupByValue** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]MetricQueryFilterRep**](MetricQueryFilterRep.md) |  | [optional] 
**UnitOfAnalysis** | **string** |  | 
**Aggregation** | **string** |  | 
**AnalysisType** | **string** |  | 
**PercentileValue** | Pointer to **int32** |  | [optional] 
**AnalysisValue** | Pointer to **float32** |  | [optional] 
**NumberOfEvents** | **int32** |  | 
**IsHigherBetter** | **bool** |  | 
**IsNumeric** | **bool** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**DefaultMissingDataToZero** | **bool** |  | 

## Methods

### NewMetricAnalyticsMetadata

`func NewMetricAnalyticsMetadata(metricKey string, metricName string, unitOfAnalysis string, aggregation string, analysisType string, numberOfEvents int32, isHigherBetter bool, isNumeric bool, defaultMissingDataToZero bool, ) *MetricAnalyticsMetadata`

NewMetricAnalyticsMetadata instantiates a new MetricAnalyticsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricAnalyticsMetadataWithDefaults

`func NewMetricAnalyticsMetadataWithDefaults() *MetricAnalyticsMetadata`

NewMetricAnalyticsMetadataWithDefaults instantiates a new MetricAnalyticsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *MetricAnalyticsMetadata) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *MetricAnalyticsMetadata) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *MetricAnalyticsMetadata) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *MetricAnalyticsMetadata) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricAnalyticsMetadata) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricAnalyticsMetadata) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetGroupBy

`func (o *MetricAnalyticsMetadata) GetGroupBy() MetricQueryGroupByRep`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricAnalyticsMetadata) GetGroupByOk() (*MetricQueryGroupByRep, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetricAnalyticsMetadata) SetGroupBy(v MetricQueryGroupByRep)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetricAnalyticsMetadata) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupByValue

`func (o *MetricAnalyticsMetadata) GetGroupByValue() string`

GetGroupByValue returns the GroupByValue field if non-nil, zero value otherwise.

### GetGroupByValueOk

`func (o *MetricAnalyticsMetadata) GetGroupByValueOk() (*string, bool)`

GetGroupByValueOk returns a tuple with the GroupByValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByValue

`func (o *MetricAnalyticsMetadata) SetGroupByValue(v string)`

SetGroupByValue sets GroupByValue field to given value.

### HasGroupByValue

`func (o *MetricAnalyticsMetadata) HasGroupByValue() bool`

HasGroupByValue returns a boolean if a field has been set.

### GetFilters

`func (o *MetricAnalyticsMetadata) GetFilters() []MetricQueryFilterRep`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricAnalyticsMetadata) GetFiltersOk() (*[]MetricQueryFilterRep, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricAnalyticsMetadata) SetFilters(v []MetricQueryFilterRep)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricAnalyticsMetadata) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUnitOfAnalysis

`func (o *MetricAnalyticsMetadata) GetUnitOfAnalysis() string`

GetUnitOfAnalysis returns the UnitOfAnalysis field if non-nil, zero value otherwise.

### GetUnitOfAnalysisOk

`func (o *MetricAnalyticsMetadata) GetUnitOfAnalysisOk() (*string, bool)`

GetUnitOfAnalysisOk returns a tuple with the UnitOfAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfAnalysis

`func (o *MetricAnalyticsMetadata) SetUnitOfAnalysis(v string)`

SetUnitOfAnalysis sets UnitOfAnalysis field to given value.


### GetAggregation

`func (o *MetricAnalyticsMetadata) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MetricAnalyticsMetadata) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MetricAnalyticsMetadata) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetAnalysisType

`func (o *MetricAnalyticsMetadata) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *MetricAnalyticsMetadata) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *MetricAnalyticsMetadata) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.


### GetPercentileValue

`func (o *MetricAnalyticsMetadata) GetPercentileValue() int32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *MetricAnalyticsMetadata) GetPercentileValueOk() (*int32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *MetricAnalyticsMetadata) SetPercentileValue(v int32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *MetricAnalyticsMetadata) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetAnalysisValue

`func (o *MetricAnalyticsMetadata) GetAnalysisValue() float32`

GetAnalysisValue returns the AnalysisValue field if non-nil, zero value otherwise.

### GetAnalysisValueOk

`func (o *MetricAnalyticsMetadata) GetAnalysisValueOk() (*float32, bool)`

GetAnalysisValueOk returns a tuple with the AnalysisValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisValue

`func (o *MetricAnalyticsMetadata) SetAnalysisValue(v float32)`

SetAnalysisValue sets AnalysisValue field to given value.

### HasAnalysisValue

`func (o *MetricAnalyticsMetadata) HasAnalysisValue() bool`

HasAnalysisValue returns a boolean if a field has been set.

### GetNumberOfEvents

`func (o *MetricAnalyticsMetadata) GetNumberOfEvents() int32`

GetNumberOfEvents returns the NumberOfEvents field if non-nil, zero value otherwise.

### GetNumberOfEventsOk

`func (o *MetricAnalyticsMetadata) GetNumberOfEventsOk() (*int32, bool)`

GetNumberOfEventsOk returns a tuple with the NumberOfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEvents

`func (o *MetricAnalyticsMetadata) SetNumberOfEvents(v int32)`

SetNumberOfEvents sets NumberOfEvents field to given value.


### GetIsHigherBetter

`func (o *MetricAnalyticsMetadata) GetIsHigherBetter() bool`

GetIsHigherBetter returns the IsHigherBetter field if non-nil, zero value otherwise.

### GetIsHigherBetterOk

`func (o *MetricAnalyticsMetadata) GetIsHigherBetterOk() (*bool, bool)`

GetIsHigherBetterOk returns a tuple with the IsHigherBetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHigherBetter

`func (o *MetricAnalyticsMetadata) SetIsHigherBetter(v bool)`

SetIsHigherBetter sets IsHigherBetter field to given value.


### GetIsNumeric

`func (o *MetricAnalyticsMetadata) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricAnalyticsMetadata) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricAnalyticsMetadata) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.


### GetUnit

`func (o *MetricAnalyticsMetadata) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricAnalyticsMetadata) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricAnalyticsMetadata) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricAnalyticsMetadata) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDefaultMissingDataToZero

`func (o *MetricAnalyticsMetadata) GetDefaultMissingDataToZero() bool`

GetDefaultMissingDataToZero returns the DefaultMissingDataToZero field if non-nil, zero value otherwise.

### GetDefaultMissingDataToZeroOk

`func (o *MetricAnalyticsMetadata) GetDefaultMissingDataToZeroOk() (*bool, bool)`

GetDefaultMissingDataToZeroOk returns a tuple with the DefaultMissingDataToZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMissingDataToZero

`func (o *MetricAnalyticsMetadata) SetDefaultMissingDataToZero(v bool)`

SetDefaultMissingDataToZero sets DefaultMissingDataToZero field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


