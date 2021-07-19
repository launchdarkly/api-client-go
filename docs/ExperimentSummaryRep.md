# ExperimentSummaryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProjKey** | Pointer to **string** |  | [optional] 
**EnvKey** | Pointer to **string** |  | [optional] 
**BaselineIdx** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**StopDate** | Pointer to **int64** |  | [optional] 
**Flag** | Pointer to [**ExperimentFlagRep**](ExperimentFlagRep.md) |  | [optional] 
**Metric** | Pointer to [**MetricListingRep**](MetricListingRep.md) |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 

## Methods

### NewExperimentSummaryRep

`func NewExperimentSummaryRep() *ExperimentSummaryRep`

NewExperimentSummaryRep instantiates a new ExperimentSummaryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentSummaryRepWithDefaults

`func NewExperimentSummaryRepWithDefaults() *ExperimentSummaryRep`

NewExperimentSummaryRepWithDefaults instantiates a new ExperimentSummaryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExperimentSummaryRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentSummaryRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentSummaryRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExperimentSummaryRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjKey

`func (o *ExperimentSummaryRep) GetProjKey() string`

GetProjKey returns the ProjKey field if non-nil, zero value otherwise.

### GetProjKeyOk

`func (o *ExperimentSummaryRep) GetProjKeyOk() (*string, bool)`

GetProjKeyOk returns a tuple with the ProjKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjKey

`func (o *ExperimentSummaryRep) SetProjKey(v string)`

SetProjKey sets ProjKey field to given value.

### HasProjKey

`func (o *ExperimentSummaryRep) HasProjKey() bool`

HasProjKey returns a boolean if a field has been set.

### GetEnvKey

`func (o *ExperimentSummaryRep) GetEnvKey() string`

GetEnvKey returns the EnvKey field if non-nil, zero value otherwise.

### GetEnvKeyOk

`func (o *ExperimentSummaryRep) GetEnvKeyOk() (*string, bool)`

GetEnvKeyOk returns a tuple with the EnvKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvKey

`func (o *ExperimentSummaryRep) SetEnvKey(v string)`

SetEnvKey sets EnvKey field to given value.

### HasEnvKey

`func (o *ExperimentSummaryRep) HasEnvKey() bool`

HasEnvKey returns a boolean if a field has been set.

### GetBaselineIdx

`func (o *ExperimentSummaryRep) GetBaselineIdx() int32`

GetBaselineIdx returns the BaselineIdx field if non-nil, zero value otherwise.

### GetBaselineIdxOk

`func (o *ExperimentSummaryRep) GetBaselineIdxOk() (*int32, bool)`

GetBaselineIdxOk returns a tuple with the BaselineIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineIdx

`func (o *ExperimentSummaryRep) SetBaselineIdx(v int32)`

SetBaselineIdx sets BaselineIdx field to given value.

### HasBaselineIdx

`func (o *ExperimentSummaryRep) HasBaselineIdx() bool`

HasBaselineIdx returns a boolean if a field has been set.

### GetStatus

`func (o *ExperimentSummaryRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExperimentSummaryRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExperimentSummaryRep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExperimentSummaryRep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *ExperimentSummaryRep) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExperimentSummaryRep) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExperimentSummaryRep) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExperimentSummaryRep) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStopDate

`func (o *ExperimentSummaryRep) GetStopDate() int64`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *ExperimentSummaryRep) GetStopDateOk() (*int64, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDate

`func (o *ExperimentSummaryRep) SetStopDate(v int64)`

SetStopDate sets StopDate field to given value.

### HasStopDate

`func (o *ExperimentSummaryRep) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.

### GetFlag

`func (o *ExperimentSummaryRep) GetFlag() ExperimentFlagRep`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ExperimentSummaryRep) GetFlagOk() (*ExperimentFlagRep, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ExperimentSummaryRep) SetFlag(v ExperimentFlagRep)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ExperimentSummaryRep) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetMetric

`func (o *ExperimentSummaryRep) GetMetric() MetricListingRep`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExperimentSummaryRep) GetMetricOk() (*MetricListingRep, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExperimentSummaryRep) SetMetric(v MetricListingRep)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ExperimentSummaryRep) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetLinks

`func (o *ExperimentSummaryRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentSummaryRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentSummaryRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentSummaryRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


