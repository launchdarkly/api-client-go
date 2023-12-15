# DiscoveredMetricEventWithMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventKey** | **string** |  | 
**IsNumeric** | **bool** |  | 
**Count** | **int32** |  | 
**FirstSeen** | **int64** |  | 
**LastSeen** | **int64** |  | 
**Metrics** | [**[]MetricListingRep**](MetricListingRep.md) |  | 

## Methods

### NewDiscoveredMetricEventWithMetrics

`func NewDiscoveredMetricEventWithMetrics(eventKey string, isNumeric bool, count int32, firstSeen int64, lastSeen int64, metrics []MetricListingRep, ) *DiscoveredMetricEventWithMetrics`

NewDiscoveredMetricEventWithMetrics instantiates a new DiscoveredMetricEventWithMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredMetricEventWithMetricsWithDefaults

`func NewDiscoveredMetricEventWithMetricsWithDefaults() *DiscoveredMetricEventWithMetrics`

NewDiscoveredMetricEventWithMetricsWithDefaults instantiates a new DiscoveredMetricEventWithMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventKey

`func (o *DiscoveredMetricEventWithMetrics) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *DiscoveredMetricEventWithMetrics) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *DiscoveredMetricEventWithMetrics) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.


### GetIsNumeric

`func (o *DiscoveredMetricEventWithMetrics) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *DiscoveredMetricEventWithMetrics) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *DiscoveredMetricEventWithMetrics) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.


### GetCount

`func (o *DiscoveredMetricEventWithMetrics) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DiscoveredMetricEventWithMetrics) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DiscoveredMetricEventWithMetrics) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFirstSeen

`func (o *DiscoveredMetricEventWithMetrics) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *DiscoveredMetricEventWithMetrics) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *DiscoveredMetricEventWithMetrics) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.


### GetLastSeen

`func (o *DiscoveredMetricEventWithMetrics) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DiscoveredMetricEventWithMetrics) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DiscoveredMetricEventWithMetrics) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.


### GetMetrics

`func (o *DiscoveredMetricEventWithMetrics) GetMetrics() []MetricListingRep`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DiscoveredMetricEventWithMetrics) GetMetricsOk() (*[]MetricListingRep, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DiscoveredMetricEventWithMetrics) SetMetrics(v []MetricListingRep)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


