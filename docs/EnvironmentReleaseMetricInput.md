# EnvironmentReleaseMetricInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** |  | 
**AutoRollback** | Pointer to **bool** | Whether a regression on this metric automatically reverts the release. When &#x60;false&#x60;, a detected regression pauses the release instead: it reports &#x60;waiting&#x60; with reason &#x60;waiting_regression_detected&#x60; and does not advance until the regression is dismissed or the release is stopped.  | [optional] [default to false]

## Methods

### NewEnvironmentReleaseMetricInput

`func NewEnvironmentReleaseMetricInput(metricKey string, ) *EnvironmentReleaseMetricInput`

NewEnvironmentReleaseMetricInput instantiates a new EnvironmentReleaseMetricInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleaseMetricInputWithDefaults

`func NewEnvironmentReleaseMetricInputWithDefaults() *EnvironmentReleaseMetricInput`

NewEnvironmentReleaseMetricInputWithDefaults instantiates a new EnvironmentReleaseMetricInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *EnvironmentReleaseMetricInput) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *EnvironmentReleaseMetricInput) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *EnvironmentReleaseMetricInput) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetAutoRollback

`func (o *EnvironmentReleaseMetricInput) GetAutoRollback() bool`

GetAutoRollback returns the AutoRollback field if non-nil, zero value otherwise.

### GetAutoRollbackOk

`func (o *EnvironmentReleaseMetricInput) GetAutoRollbackOk() (*bool, bool)`

GetAutoRollbackOk returns a tuple with the AutoRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRollback

`func (o *EnvironmentReleaseMetricInput) SetAutoRollback(v bool)`

SetAutoRollback sets AutoRollback field to given value.

### HasAutoRollback

`func (o *EnvironmentReleaseMetricInput) HasAutoRollback() bool`

HasAutoRollback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


