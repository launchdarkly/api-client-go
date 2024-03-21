# InsightsChartMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | **map[string]interface{}** |  | 
**Name** | Pointer to **string** | Name of the chart | [optional] 
**Metrics** | Pointer to [**map[string]InsightsChartMetric**](InsightsChartMetric.md) |  | [optional] 
**XAxis** | [**InsightsChartSeriesMetadataAxis**](InsightsChartSeriesMetadataAxis.md) |  | 
**YAxis** | [**InsightsChartSeriesMetadataAxis**](InsightsChartSeriesMetadataAxis.md) |  | 

## Methods

### NewInsightsChartMetadata

`func NewInsightsChartMetadata(summary map[string]interface{}, xAxis InsightsChartSeriesMetadataAxis, yAxis InsightsChartSeriesMetadataAxis, ) *InsightsChartMetadata`

NewInsightsChartMetadata instantiates a new InsightsChartMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsChartMetadataWithDefaults

`func NewInsightsChartMetadataWithDefaults() *InsightsChartMetadata`

NewInsightsChartMetadataWithDefaults instantiates a new InsightsChartMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *InsightsChartMetadata) GetSummary() map[string]interface{}`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InsightsChartMetadata) GetSummaryOk() (*map[string]interface{}, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InsightsChartMetadata) SetSummary(v map[string]interface{})`

SetSummary sets Summary field to given value.


### GetName

`func (o *InsightsChartMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightsChartMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightsChartMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InsightsChartMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetrics

`func (o *InsightsChartMetadata) GetMetrics() map[string]InsightsChartMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InsightsChartMetadata) GetMetricsOk() (*map[string]InsightsChartMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InsightsChartMetadata) SetMetrics(v map[string]InsightsChartMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *InsightsChartMetadata) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetXAxis

`func (o *InsightsChartMetadata) GetXAxis() InsightsChartSeriesMetadataAxis`

GetXAxis returns the XAxis field if non-nil, zero value otherwise.

### GetXAxisOk

`func (o *InsightsChartMetadata) GetXAxisOk() (*InsightsChartSeriesMetadataAxis, bool)`

GetXAxisOk returns a tuple with the XAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxis

`func (o *InsightsChartMetadata) SetXAxis(v InsightsChartSeriesMetadataAxis)`

SetXAxis sets XAxis field to given value.


### GetYAxis

`func (o *InsightsChartMetadata) GetYAxis() InsightsChartSeriesMetadataAxis`

GetYAxis returns the YAxis field if non-nil, zero value otherwise.

### GetYAxisOk

`func (o *InsightsChartMetadata) GetYAxisOk() (*InsightsChartSeriesMetadataAxis, bool)`

GetYAxisOk returns a tuple with the YAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxis

`func (o *InsightsChartMetadata) SetYAxis(v InsightsChartSeriesMetadataAxis)`

SetYAxis sets YAxis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


