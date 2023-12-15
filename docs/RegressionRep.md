# RegressionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IterationId** | **string** |  | 
**MetricKey** | **string** |  | 
**Timestamp** | **int64** |  | 

## Methods

### NewRegressionRep

`func NewRegressionRep(iterationId string, metricKey string, timestamp int64, ) *RegressionRep`

NewRegressionRep instantiates a new RegressionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegressionRepWithDefaults

`func NewRegressionRepWithDefaults() *RegressionRep`

NewRegressionRepWithDefaults instantiates a new RegressionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIterationId

`func (o *RegressionRep) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *RegressionRep) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *RegressionRep) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.


### GetMetricKey

`func (o *RegressionRep) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *RegressionRep) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *RegressionRep) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetTimestamp

`func (o *RegressionRep) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RegressionRep) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RegressionRep) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


