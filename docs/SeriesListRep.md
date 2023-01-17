# SeriesListRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | **map[string]interface{}** | The location and content type of related resources | 
**Metadata** | **[]map[string]interface{}** | Metadata about each series | 
**Series** | **[]map[string]int32** | An array of data points with timestamps | 

## Methods

### NewSeriesListRep

`func NewSeriesListRep(links map[string]interface{}, metadata []map[string]interface{}, series []map[string]int32, ) *SeriesListRep`

NewSeriesListRep instantiates a new SeriesListRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesListRepWithDefaults

`func NewSeriesListRepWithDefaults() *SeriesListRep`

NewSeriesListRepWithDefaults instantiates a new SeriesListRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SeriesListRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SeriesListRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SeriesListRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *SeriesListRep) GetMetadata() []map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SeriesListRep) GetMetadataOk() (*[]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SeriesListRep) SetMetadata(v []map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetSeries

`func (o *SeriesListRep) GetSeries() []map[string]int32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SeriesListRep) GetSeriesOk() (*[]map[string]int32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SeriesListRep) SetSeries(v []map[string]int32)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


