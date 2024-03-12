# InsightsChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**InsightsChartMetadata**](InsightsChartMetadata.md) |  | 
**Series** | [**[]InsightsChartSeries**](InsightsChartSeries.md) | Series data for the chart | 

## Methods

### NewInsightsChart

`func NewInsightsChart(metadata InsightsChartMetadata, series []InsightsChartSeries, ) *InsightsChart`

NewInsightsChart instantiates a new InsightsChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsChartWithDefaults

`func NewInsightsChartWithDefaults() *InsightsChart`

NewInsightsChartWithDefaults instantiates a new InsightsChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *InsightsChart) GetMetadata() InsightsChartMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InsightsChart) GetMetadataOk() (*InsightsChartMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InsightsChart) SetMetadata(v InsightsChartMetadata)`

SetMetadata sets Metadata field to given value.


### GetSeries

`func (o *InsightsChart) GetSeries() []InsightsChartSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InsightsChart) GetSeriesOk() (*[]InsightsChartSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InsightsChart) SetSeries(v []InsightsChartSeries)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


