# ExperimentResultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 
**Metadata** | Pointer to [**[]ExperimentMetadataRep**](ExperimentMetadataRep.md) |  | [optional] 
**Totals** | Pointer to [**[]ExperimentTotalsRep**](ExperimentTotalsRep.md) |  | [optional] 
**Series** | Pointer to [**[]ExperimentTimeSeriesSlice**](ExperimentTimeSeriesSlice.md) |  | [optional] 
**Stats** | Pointer to [**ExperimentStatsRep**](ExperimentStatsRep.md) |  | [optional] 
**Granularity** | Pointer to **string** |  | [optional] 

## Methods

### NewExperimentResultsRep

`func NewExperimentResultsRep() *ExperimentResultsRep`

NewExperimentResultsRep instantiates a new ExperimentResultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentResultsRepWithDefaults

`func NewExperimentResultsRepWithDefaults() *ExperimentResultsRep`

NewExperimentResultsRepWithDefaults instantiates a new ExperimentResultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExperimentResultsRep) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentResultsRep) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentResultsRep) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentResultsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ExperimentResultsRep) GetMetadata() []ExperimentMetadataRep`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExperimentResultsRep) GetMetadataOk() (*[]ExperimentMetadataRep, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExperimentResultsRep) SetMetadata(v []ExperimentMetadataRep)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExperimentResultsRep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTotals

`func (o *ExperimentResultsRep) GetTotals() []ExperimentTotalsRep`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *ExperimentResultsRep) GetTotalsOk() (*[]ExperimentTotalsRep, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *ExperimentResultsRep) SetTotals(v []ExperimentTotalsRep)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *ExperimentResultsRep) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetSeries

`func (o *ExperimentResultsRep) GetSeries() []ExperimentTimeSeriesSlice`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ExperimentResultsRep) GetSeriesOk() (*[]ExperimentTimeSeriesSlice, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ExperimentResultsRep) SetSeries(v []ExperimentTimeSeriesSlice)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ExperimentResultsRep) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStats

`func (o *ExperimentResultsRep) GetStats() ExperimentStatsRep`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ExperimentResultsRep) GetStatsOk() (*ExperimentStatsRep, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ExperimentResultsRep) SetStats(v ExperimentStatsRep)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ExperimentResultsRep) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetGranularity

`func (o *ExperimentResultsRep) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ExperimentResultsRep) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ExperimentResultsRep) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *ExperimentResultsRep) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


