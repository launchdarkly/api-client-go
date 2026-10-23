# ReleaseMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** |  | 
**MetricVersionId** | Pointer to **string** |  | [optional] [readonly] 
**AutoRollback** | **bool** | Whether a regression automatically reverts the release. | 
**Definition** | Pointer to [**EnvironmentReleasesMetricListingRep**](EnvironmentReleasesMetricListingRep.md) | Present only when the caller requests &#x60;expand&#x3D;metrics&#x60; and the pinned metric version still exists. A metric whose pinned version has been deleted keeps its &#x60;metricKey&#x60; and &#x60;metricVersionId&#x60; but has no definition.  | [optional] [readonly] 

## Methods

### NewReleaseMetric

`func NewReleaseMetric(metricKey string, autoRollback bool, ) *ReleaseMetric`

NewReleaseMetric instantiates a new ReleaseMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseMetricWithDefaults

`func NewReleaseMetricWithDefaults() *ReleaseMetric`

NewReleaseMetricWithDefaults instantiates a new ReleaseMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *ReleaseMetric) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *ReleaseMetric) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *ReleaseMetric) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricVersionId

`func (o *ReleaseMetric) GetMetricVersionId() string`

GetMetricVersionId returns the MetricVersionId field if non-nil, zero value otherwise.

### GetMetricVersionIdOk

`func (o *ReleaseMetric) GetMetricVersionIdOk() (*string, bool)`

GetMetricVersionIdOk returns a tuple with the MetricVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricVersionId

`func (o *ReleaseMetric) SetMetricVersionId(v string)`

SetMetricVersionId sets MetricVersionId field to given value.

### HasMetricVersionId

`func (o *ReleaseMetric) HasMetricVersionId() bool`

HasMetricVersionId returns a boolean if a field has been set.

### GetAutoRollback

`func (o *ReleaseMetric) GetAutoRollback() bool`

GetAutoRollback returns the AutoRollback field if non-nil, zero value otherwise.

### GetAutoRollbackOk

`func (o *ReleaseMetric) GetAutoRollbackOk() (*bool, bool)`

GetAutoRollbackOk returns a tuple with the AutoRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRollback

`func (o *ReleaseMetric) SetAutoRollback(v bool)`

SetAutoRollback sets AutoRollback field to given value.


### GetDefinition

`func (o *ReleaseMetric) GetDefinition() EnvironmentReleasesMetricListingRep`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ReleaseMetric) GetDefinitionOk() (*EnvironmentReleasesMetricListingRep, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ReleaseMetric) SetDefinition(v EnvironmentReleasesMetricListingRep)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ReleaseMetric) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


