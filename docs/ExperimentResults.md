# ExperimentResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Metadata** | Pointer to [**[]ExperimentMetadataRep**](ExperimentMetadataRep.md) |  | [optional] 
**Totals** | Pointer to [**[]ExperimentTotalsRep**](ExperimentTotalsRep.md) |  | [optional] 
**Series** | Pointer to [**[]ExperimentTimeSeriesSlice**](ExperimentTimeSeriesSlice.md) |  | [optional] 
**Stats** | Pointer to [**ExperimentStatsRep**](ExperimentStatsRep.md) |  | [optional] 
**Granularity** | Pointer to **string** |  | [optional] 
**MetricSeen** | Pointer to [**MetricSeen**](MetricSeen.md) |  | [optional] 

## Methods

### NewExperimentResults

`func NewExperimentResults() *ExperimentResults`

NewExperimentResults instantiates a new ExperimentResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentResultsWithDefaults

`func NewExperimentResultsWithDefaults() *ExperimentResults`

NewExperimentResultsWithDefaults instantiates a new ExperimentResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExperimentResults) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentResults) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentResults) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentResults) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *ExperimentResults) GetMetadata() []ExperimentMetadataRep`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExperimentResults) GetMetadataOk() (*[]ExperimentMetadataRep, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExperimentResults) SetMetadata(v []ExperimentMetadataRep)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExperimentResults) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTotals

`func (o *ExperimentResults) GetTotals() []ExperimentTotalsRep`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *ExperimentResults) GetTotalsOk() (*[]ExperimentTotalsRep, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *ExperimentResults) SetTotals(v []ExperimentTotalsRep)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *ExperimentResults) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetSeries

`func (o *ExperimentResults) GetSeries() []ExperimentTimeSeriesSlice`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ExperimentResults) GetSeriesOk() (*[]ExperimentTimeSeriesSlice, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ExperimentResults) SetSeries(v []ExperimentTimeSeriesSlice)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ExperimentResults) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStats

`func (o *ExperimentResults) GetStats() ExperimentStatsRep`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ExperimentResults) GetStatsOk() (*ExperimentStatsRep, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ExperimentResults) SetStats(v ExperimentStatsRep)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ExperimentResults) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetGranularity

`func (o *ExperimentResults) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ExperimentResults) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ExperimentResults) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *ExperimentResults) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetMetricSeen

`func (o *ExperimentResults) GetMetricSeen() MetricSeen`

GetMetricSeen returns the MetricSeen field if non-nil, zero value otherwise.

### GetMetricSeenOk

`func (o *ExperimentResults) GetMetricSeenOk() (*MetricSeen, bool)`

GetMetricSeenOk returns a tuple with the MetricSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSeen

`func (o *ExperimentResults) SetMetricSeen(v MetricSeen)`

SetMetricSeen sets MetricSeen field to given value.

### HasMetricSeen

`func (o *ExperimentResults) HasMetricSeen() bool`

HasMetricSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


