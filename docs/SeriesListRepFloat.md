# SeriesListRepFloat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | **map[string]interface{}** | The location and content type of related resources | 
**Metadata** | **[]map[string]interface{}** | Metadata about each series | 
**Series** | **[]map[string]float32** | An array of data points with timestamps. Each element of the array is an object with a &#39;time&#39; field, whose value is the timestamp, and one or more key fields. If there are multiple key fields, they are labeled &#39;0&#39;, &#39;1&#39;, and so on, and are explained in the &lt;code&gt;metadata&lt;/code&gt;. | 

## Methods

### NewSeriesListRepFloat

`func NewSeriesListRepFloat(links map[string]interface{}, metadata []map[string]interface{}, series []map[string]float32, ) *SeriesListRepFloat`

NewSeriesListRepFloat instantiates a new SeriesListRepFloat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesListRepFloatWithDefaults

`func NewSeriesListRepFloatWithDefaults() *SeriesListRepFloat`

NewSeriesListRepFloatWithDefaults instantiates a new SeriesListRepFloat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SeriesListRepFloat) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SeriesListRepFloat) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SeriesListRepFloat) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *SeriesListRepFloat) GetMetadata() []map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SeriesListRepFloat) GetMetadataOk() (*[]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SeriesListRepFloat) SetMetadata(v []map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetSeries

`func (o *SeriesListRepFloat) GetSeries() []map[string]float32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SeriesListRepFloat) GetSeriesOk() (*[]map[string]float32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SeriesListRepFloat) SetSeries(v []map[string]float32)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


