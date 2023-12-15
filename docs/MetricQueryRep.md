# MetricQueryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** |  | 
**GroupBy** | Pointer to [**MetricQueryGroupByRep**](MetricQueryGroupByRep.md) |  | [optional] 
**Filters** | Pointer to [**[]MetricQueryFilterRep**](MetricQueryFilterRep.md) |  | [optional] 

## Methods

### NewMetricQueryRep

`func NewMetricQueryRep(metricKey string, ) *MetricQueryRep`

NewMetricQueryRep instantiates a new MetricQueryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricQueryRepWithDefaults

`func NewMetricQueryRepWithDefaults() *MetricQueryRep`

NewMetricQueryRepWithDefaults instantiates a new MetricQueryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *MetricQueryRep) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *MetricQueryRep) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *MetricQueryRep) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetGroupBy

`func (o *MetricQueryRep) GetGroupBy() MetricQueryGroupByRep`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricQueryRep) GetGroupByOk() (*MetricQueryGroupByRep, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetricQueryRep) SetGroupBy(v MetricQueryGroupByRep)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetricQueryRep) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilters

`func (o *MetricQueryRep) GetFilters() []MetricQueryFilterRep`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricQueryRep) GetFiltersOk() (*[]MetricQueryFilterRep, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricQueryRep) SetFilters(v []MetricQueryFilterRep)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricQueryRep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


