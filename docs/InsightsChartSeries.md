# InsightsChartSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**InsightsChartSeriesMetadata**](InsightsChartSeriesMetadata.md) |  | 
**Data** | [**[]InsightsChartSeriesDataPoint**](InsightsChartSeriesDataPoint.md) | Data points for the series | 

## Methods

### NewInsightsChartSeries

`func NewInsightsChartSeries(metadata InsightsChartSeriesMetadata, data []InsightsChartSeriesDataPoint, ) *InsightsChartSeries`

NewInsightsChartSeries instantiates a new InsightsChartSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsChartSeriesWithDefaults

`func NewInsightsChartSeriesWithDefaults() *InsightsChartSeries`

NewInsightsChartSeriesWithDefaults instantiates a new InsightsChartSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *InsightsChartSeries) GetMetadata() InsightsChartSeriesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InsightsChartSeries) GetMetadataOk() (*InsightsChartSeriesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InsightsChartSeries) SetMetadata(v InsightsChartSeriesMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *InsightsChartSeries) GetData() []InsightsChartSeriesDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InsightsChartSeries) GetDataOk() (*[]InsightsChartSeriesDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InsightsChartSeries) SetData(v []InsightsChartSeriesDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


