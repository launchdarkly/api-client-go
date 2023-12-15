# PostProductAnalyticsQueryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimestamp** | **int64** |  | 
**EndTimestamp** | **int64** |  | 
**MetricQueries** | [**[]MetricQueryRep**](MetricQueryRep.md) |  | 

## Methods

### NewPostProductAnalyticsQueryRep

`func NewPostProductAnalyticsQueryRep(startTimestamp int64, endTimestamp int64, metricQueries []MetricQueryRep, ) *PostProductAnalyticsQueryRep`

NewPostProductAnalyticsQueryRep instantiates a new PostProductAnalyticsQueryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProductAnalyticsQueryRepWithDefaults

`func NewPostProductAnalyticsQueryRepWithDefaults() *PostProductAnalyticsQueryRep`

NewPostProductAnalyticsQueryRepWithDefaults instantiates a new PostProductAnalyticsQueryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimestamp

`func (o *PostProductAnalyticsQueryRep) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *PostProductAnalyticsQueryRep) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *PostProductAnalyticsQueryRep) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *PostProductAnalyticsQueryRep) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *PostProductAnalyticsQueryRep) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *PostProductAnalyticsQueryRep) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetMetricQueries

`func (o *PostProductAnalyticsQueryRep) GetMetricQueries() []MetricQueryRep`

GetMetricQueries returns the MetricQueries field if non-nil, zero value otherwise.

### GetMetricQueriesOk

`func (o *PostProductAnalyticsQueryRep) GetMetricQueriesOk() (*[]MetricQueryRep, bool)`

GetMetricQueriesOk returns a tuple with the MetricQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricQueries

`func (o *PostProductAnalyticsQueryRep) SetMetricQueries(v []MetricQueryRep)`

SetMetricQueries sets MetricQueries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


