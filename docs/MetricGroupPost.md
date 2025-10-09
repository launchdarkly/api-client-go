# MetricGroupPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A unique key to reference the metric group | [optional] 
**Name** | **string** | A human-friendly name for the metric group | 
**Kind** | **string** | The type of the metric group | 
**Description** | Pointer to **string** | Description of the metric group | [optional] 
**MaintainerId** | **string** | The ID of the member who maintains this metric group | 
**Tags** | **[]string** | Tags for the metric group | 
**Metrics** | [**[]MetricInMetricGroupInput**](MetricInMetricGroupInput.md) | An ordered list of the metrics in this metric group | 

## Methods

### NewMetricGroupPost

`func NewMetricGroupPost(name string, kind string, maintainerId string, tags []string, metrics []MetricInMetricGroupInput, ) *MetricGroupPost`

NewMetricGroupPost instantiates a new MetricGroupPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGroupPostWithDefaults

`func NewMetricGroupPostWithDefaults() *MetricGroupPost`

NewMetricGroupPostWithDefaults instantiates a new MetricGroupPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricGroupPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricGroupPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricGroupPost) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricGroupPost) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *MetricGroupPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricGroupPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricGroupPost) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *MetricGroupPost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricGroupPost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricGroupPost) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *MetricGroupPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricGroupPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricGroupPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricGroupPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainerId

`func (o *MetricGroupPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *MetricGroupPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *MetricGroupPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetTags

`func (o *MetricGroupPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricGroupPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricGroupPost) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetMetrics

`func (o *MetricGroupPost) GetMetrics() []MetricInMetricGroupInput`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricGroupPost) GetMetricsOk() (*[]MetricInMetricGroupInput, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricGroupPost) SetMetrics(v []MetricInMetricGroupInput)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


