# EnvironmentReleasePlanMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsGroup** | **bool** | true if this key identifies a metric group, false if it identifies a single metric | 
**AutoRollback** | Pointer to **bool** | Whether to automatically roll back the release if this metric/group regresses | [optional] [default to false]

## Methods

### NewEnvironmentReleasePlanMetric

`func NewEnvironmentReleasePlanMetric(isGroup bool, ) *EnvironmentReleasePlanMetric`

NewEnvironmentReleasePlanMetric instantiates a new EnvironmentReleasePlanMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanMetricWithDefaults

`func NewEnvironmentReleasePlanMetricWithDefaults() *EnvironmentReleasePlanMetric`

NewEnvironmentReleasePlanMetricWithDefaults instantiates a new EnvironmentReleasePlanMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsGroup

`func (o *EnvironmentReleasePlanMetric) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *EnvironmentReleasePlanMetric) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *EnvironmentReleasePlanMetric) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.


### GetAutoRollback

`func (o *EnvironmentReleasePlanMetric) GetAutoRollback() bool`

GetAutoRollback returns the AutoRollback field if non-nil, zero value otherwise.

### GetAutoRollbackOk

`func (o *EnvironmentReleasePlanMetric) GetAutoRollbackOk() (*bool, bool)`

GetAutoRollbackOk returns a tuple with the AutoRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRollback

`func (o *EnvironmentReleasePlanMetric) SetAutoRollback(v bool)`

SetAutoRollback sets AutoRollback field to given value.

### HasAutoRollback

`func (o *EnvironmentReleasePlanMetric) HasAutoRollback() bool`

HasAutoRollback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


