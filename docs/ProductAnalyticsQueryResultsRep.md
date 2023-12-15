# ProductAnalyticsQueryResultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryParameters** | [**[]QueryParameters**](QueryParameters.md) |  | 
**Series** | **[]map[string]int32** |  | 
**Metadata** | [**[]MetricAnalyticsMetadata**](MetricAnalyticsMetadata.md) |  | 

## Methods

### NewProductAnalyticsQueryResultsRep

`func NewProductAnalyticsQueryResultsRep(queryParameters []QueryParameters, series []map[string]int32, metadata []MetricAnalyticsMetadata, ) *ProductAnalyticsQueryResultsRep`

NewProductAnalyticsQueryResultsRep instantiates a new ProductAnalyticsQueryResultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAnalyticsQueryResultsRepWithDefaults

`func NewProductAnalyticsQueryResultsRepWithDefaults() *ProductAnalyticsQueryResultsRep`

NewProductAnalyticsQueryResultsRepWithDefaults instantiates a new ProductAnalyticsQueryResultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryParameters

`func (o *ProductAnalyticsQueryResultsRep) GetQueryParameters() []QueryParameters`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *ProductAnalyticsQueryResultsRep) GetQueryParametersOk() (*[]QueryParameters, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *ProductAnalyticsQueryResultsRep) SetQueryParameters(v []QueryParameters)`

SetQueryParameters sets QueryParameters field to given value.


### GetSeries

`func (o *ProductAnalyticsQueryResultsRep) GetSeries() []map[string]int32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ProductAnalyticsQueryResultsRep) GetSeriesOk() (*[]map[string]int32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ProductAnalyticsQueryResultsRep) SetSeries(v []map[string]int32)`

SetSeries sets Series field to given value.


### GetMetadata

`func (o *ProductAnalyticsQueryResultsRep) GetMetadata() []MetricAnalyticsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductAnalyticsQueryResultsRep) GetMetadataOk() (*[]MetricAnalyticsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductAnalyticsQueryResultsRep) SetMetadata(v []MetricAnalyticsMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


