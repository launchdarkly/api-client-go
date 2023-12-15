# DiscoveredMetricEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventKey** | **string** |  | 
**IsNumeric** | **bool** |  | 
**Count** | **int32** |  | 
**FirstSeen** | **int64** |  | 
**LastSeen** | **int64** |  | 

## Methods

### NewDiscoveredMetricEvent

`func NewDiscoveredMetricEvent(eventKey string, isNumeric bool, count int32, firstSeen int64, lastSeen int64, ) *DiscoveredMetricEvent`

NewDiscoveredMetricEvent instantiates a new DiscoveredMetricEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredMetricEventWithDefaults

`func NewDiscoveredMetricEventWithDefaults() *DiscoveredMetricEvent`

NewDiscoveredMetricEventWithDefaults instantiates a new DiscoveredMetricEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventKey

`func (o *DiscoveredMetricEvent) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *DiscoveredMetricEvent) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *DiscoveredMetricEvent) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.


### GetIsNumeric

`func (o *DiscoveredMetricEvent) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *DiscoveredMetricEvent) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *DiscoveredMetricEvent) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.


### GetCount

`func (o *DiscoveredMetricEvent) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DiscoveredMetricEvent) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DiscoveredMetricEvent) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFirstSeen

`func (o *DiscoveredMetricEvent) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *DiscoveredMetricEvent) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *DiscoveredMetricEvent) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.


### GetLastSeen

`func (o *DiscoveredMetricEvent) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DiscoveredMetricEvent) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DiscoveredMetricEvent) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


