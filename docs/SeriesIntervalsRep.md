# SeriesIntervalsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**[]Series**](Series.md) | An array of timestamps and values for a given meter | 
**Links** | **map[string]interface{}** | The location and content type of related resources | 

## Methods

### NewSeriesIntervalsRep

`func NewSeriesIntervalsRep(series []Series, links map[string]interface{}, ) *SeriesIntervalsRep`

NewSeriesIntervalsRep instantiates a new SeriesIntervalsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesIntervalsRepWithDefaults

`func NewSeriesIntervalsRepWithDefaults() *SeriesIntervalsRep`

NewSeriesIntervalsRepWithDefaults instantiates a new SeriesIntervalsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *SeriesIntervalsRep) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SeriesIntervalsRep) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SeriesIntervalsRep) SetSeries(v []Series)`

SetSeries sets Series field to given value.


### GetLinks

`func (o *SeriesIntervalsRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SeriesIntervalsRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SeriesIntervalsRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


