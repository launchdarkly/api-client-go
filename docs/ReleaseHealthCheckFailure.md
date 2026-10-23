# ReleaseHealthCheckFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ReleaseHealthCheckKind**](ReleaseHealthCheckKind.md) |  | 
**MetricKey** | Pointer to **string** | The metric key associated with a metric event check. | [optional] 
**Message** | **string** | A description of the failed health check. | 

## Methods

### NewReleaseHealthCheckFailure

`func NewReleaseHealthCheckFailure(kind ReleaseHealthCheckKind, message string, ) *ReleaseHealthCheckFailure`

NewReleaseHealthCheckFailure instantiates a new ReleaseHealthCheckFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseHealthCheckFailureWithDefaults

`func NewReleaseHealthCheckFailureWithDefaults() *ReleaseHealthCheckFailure`

NewReleaseHealthCheckFailureWithDefaults instantiates a new ReleaseHealthCheckFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ReleaseHealthCheckFailure) GetKind() ReleaseHealthCheckKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReleaseHealthCheckFailure) GetKindOk() (*ReleaseHealthCheckKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReleaseHealthCheckFailure) SetKind(v ReleaseHealthCheckKind)`

SetKind sets Kind field to given value.


### GetMetricKey

`func (o *ReleaseHealthCheckFailure) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *ReleaseHealthCheckFailure) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *ReleaseHealthCheckFailure) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.

### HasMetricKey

`func (o *ReleaseHealthCheckFailure) HasMetricKey() bool`

HasMetricKey returns a boolean if a field has been set.

### GetMessage

`func (o *ReleaseHealthCheckFailure) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReleaseHealthCheckFailure) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReleaseHealthCheckFailure) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


