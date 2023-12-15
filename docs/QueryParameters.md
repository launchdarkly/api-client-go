# QueryParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** |  | 
**MetricName** | **string** |  | 
**GroupBy** | Pointer to [**MetricQueryGroupByRep**](MetricQueryGroupByRep.md) |  | [optional] 
**Filters** | Pointer to [**[]MetricQueryFilterRep**](MetricQueryFilterRep.md) |  | [optional] 
**Aggregation** | **string** |  | 
**IsHigherBetter** | **bool** |  | 
**UnitOfAnalysis** | **string** |  | 
**AnalysisType** | **string** |  | 

## Methods

### NewQueryParameters

`func NewQueryParameters(metricKey string, metricName string, aggregation string, isHigherBetter bool, unitOfAnalysis string, analysisType string, ) *QueryParameters`

NewQueryParameters instantiates a new QueryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParametersWithDefaults

`func NewQueryParametersWithDefaults() *QueryParameters`

NewQueryParametersWithDefaults instantiates a new QueryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *QueryParameters) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *QueryParameters) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *QueryParameters) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *QueryParameters) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *QueryParameters) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *QueryParameters) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetGroupBy

`func (o *QueryParameters) GetGroupBy() MetricQueryGroupByRep`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *QueryParameters) GetGroupByOk() (*MetricQueryGroupByRep, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *QueryParameters) SetGroupBy(v MetricQueryGroupByRep)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *QueryParameters) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilters

`func (o *QueryParameters) GetFilters() []MetricQueryFilterRep`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryParameters) GetFiltersOk() (*[]MetricQueryFilterRep, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryParameters) SetFilters(v []MetricQueryFilterRep)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryParameters) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetAggregation

`func (o *QueryParameters) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *QueryParameters) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *QueryParameters) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetIsHigherBetter

`func (o *QueryParameters) GetIsHigherBetter() bool`

GetIsHigherBetter returns the IsHigherBetter field if non-nil, zero value otherwise.

### GetIsHigherBetterOk

`func (o *QueryParameters) GetIsHigherBetterOk() (*bool, bool)`

GetIsHigherBetterOk returns a tuple with the IsHigherBetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHigherBetter

`func (o *QueryParameters) SetIsHigherBetter(v bool)`

SetIsHigherBetter sets IsHigherBetter field to given value.


### GetUnitOfAnalysis

`func (o *QueryParameters) GetUnitOfAnalysis() string`

GetUnitOfAnalysis returns the UnitOfAnalysis field if non-nil, zero value otherwise.

### GetUnitOfAnalysisOk

`func (o *QueryParameters) GetUnitOfAnalysisOk() (*string, bool)`

GetUnitOfAnalysisOk returns a tuple with the UnitOfAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfAnalysis

`func (o *QueryParameters) SetUnitOfAnalysis(v string)`

SetUnitOfAnalysis sets UnitOfAnalysis field to given value.


### GetAnalysisType

`func (o *QueryParameters) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *QueryParameters) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *QueryParameters) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


