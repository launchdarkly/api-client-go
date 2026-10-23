# DataExportSeriesListRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | **map[string]interface{}** | The location and content type of related resources | 
**Metadata** | **[]map[string]interface{}** | Metadata about each series | 
**Series** | **[]map[string]int64** | An array of data points with timestamps. Each element of the array is an object with a &#39;time&#39; field, whose value is the timestamp, and one or more key fields. If there are multiple key fields, they are labeled &#39;0&#39;, &#39;1&#39;, and so on, and are explained in the &lt;code&gt;metadata&lt;/code&gt;. | 
**DataThrough** | Pointer to **string** | The end of the latest hour written by this source&#39;s usage pipeline, as an RFC 3339 UTC timestamp. It is measured across all accounts, not only yours, and usage after this time is not counted yet. Omitted when the watermark is unavailable. | [optional] 
**Delayed** | Pointer to **bool** | Whether this source&#39;s usage pipeline is late: it wrote every hour of the 24 hours ending at &lt;code&gt;dataThrough&lt;/code&gt;, and more than three hours have passed since &lt;code&gt;dataThrough&lt;/code&gt;. Always &lt;code&gt;false&lt;/code&gt; for a source with gaps in those 24 hours, where gaps usually mean no usage rather than a stalled pipeline. Omitted when &lt;code&gt;dataThrough&lt;/code&gt; is omitted. | [optional] 

## Methods

### NewDataExportSeriesListRep

`func NewDataExportSeriesListRep(links map[string]interface{}, metadata []map[string]interface{}, series []map[string]int64, ) *DataExportSeriesListRep`

NewDataExportSeriesListRep instantiates a new DataExportSeriesListRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExportSeriesListRepWithDefaults

`func NewDataExportSeriesListRepWithDefaults() *DataExportSeriesListRep`

NewDataExportSeriesListRepWithDefaults instantiates a new DataExportSeriesListRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DataExportSeriesListRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataExportSeriesListRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataExportSeriesListRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *DataExportSeriesListRep) GetMetadata() []map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataExportSeriesListRep) GetMetadataOk() (*[]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataExportSeriesListRep) SetMetadata(v []map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetSeries

`func (o *DataExportSeriesListRep) GetSeries() []map[string]int64`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DataExportSeriesListRep) GetSeriesOk() (*[]map[string]int64, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DataExportSeriesListRep) SetSeries(v []map[string]int64)`

SetSeries sets Series field to given value.


### GetDataThrough

`func (o *DataExportSeriesListRep) GetDataThrough() string`

GetDataThrough returns the DataThrough field if non-nil, zero value otherwise.

### GetDataThroughOk

`func (o *DataExportSeriesListRep) GetDataThroughOk() (*string, bool)`

GetDataThroughOk returns a tuple with the DataThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataThrough

`func (o *DataExportSeriesListRep) SetDataThrough(v string)`

SetDataThrough sets DataThrough field to given value.

### HasDataThrough

`func (o *DataExportSeriesListRep) HasDataThrough() bool`

HasDataThrough returns a boolean if a field has been set.

### GetDelayed

`func (o *DataExportSeriesListRep) GetDelayed() bool`

GetDelayed returns the Delayed field if non-nil, zero value otherwise.

### GetDelayedOk

`func (o *DataExportSeriesListRep) GetDelayedOk() (*bool, bool)`

GetDelayedOk returns a tuple with the Delayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayed

`func (o *DataExportSeriesListRep) SetDelayed(v bool)`

SetDelayed sets Delayed field to given value.

### HasDelayed

`func (o *DataExportSeriesListRep) HasDelayed() bool`

HasDelayed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


