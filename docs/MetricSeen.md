# MetricSeen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ever** | Pointer to **bool** | Whether the metric has received an event for this iteration | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp of when the metric most recently received an event for this iteration | [optional] 

## Methods

### NewMetricSeen

`func NewMetricSeen() *MetricSeen`

NewMetricSeen instantiates a new MetricSeen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSeenWithDefaults

`func NewMetricSeenWithDefaults() *MetricSeen`

NewMetricSeenWithDefaults instantiates a new MetricSeen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEver

`func (o *MetricSeen) GetEver() bool`

GetEver returns the Ever field if non-nil, zero value otherwise.

### GetEverOk

`func (o *MetricSeen) GetEverOk() (*bool, bool)`

GetEverOk returns a tuple with the Ever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEver

`func (o *MetricSeen) SetEver(v bool)`

SetEver sets Ever field to given value.

### HasEver

`func (o *MetricSeen) HasEver() bool`

HasEver returns a boolean if a field has been set.

### GetTimestamp

`func (o *MetricSeen) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetricSeen) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetricSeen) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetricSeen) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


