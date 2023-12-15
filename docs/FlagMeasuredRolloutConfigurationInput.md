# FlagMeasuredRolloutConfigurationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKeys** | Pointer to **[]string** | The metrics to use for measured rollout | [optional] 
**RandomizationUnit** | Pointer to **string** | The randomization unit to use for measured rollout | [optional] 

## Methods

### NewFlagMeasuredRolloutConfigurationInput

`func NewFlagMeasuredRolloutConfigurationInput() *FlagMeasuredRolloutConfigurationInput`

NewFlagMeasuredRolloutConfigurationInput instantiates a new FlagMeasuredRolloutConfigurationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagMeasuredRolloutConfigurationInputWithDefaults

`func NewFlagMeasuredRolloutConfigurationInputWithDefaults() *FlagMeasuredRolloutConfigurationInput`

NewFlagMeasuredRolloutConfigurationInputWithDefaults instantiates a new FlagMeasuredRolloutConfigurationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKeys

`func (o *FlagMeasuredRolloutConfigurationInput) GetMetricKeys() []string`

GetMetricKeys returns the MetricKeys field if non-nil, zero value otherwise.

### GetMetricKeysOk

`func (o *FlagMeasuredRolloutConfigurationInput) GetMetricKeysOk() (*[]string, bool)`

GetMetricKeysOk returns a tuple with the MetricKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKeys

`func (o *FlagMeasuredRolloutConfigurationInput) SetMetricKeys(v []string)`

SetMetricKeys sets MetricKeys field to given value.

### HasMetricKeys

`func (o *FlagMeasuredRolloutConfigurationInput) HasMetricKeys() bool`

HasMetricKeys returns a boolean if a field has been set.

### GetRandomizationUnit

`func (o *FlagMeasuredRolloutConfigurationInput) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *FlagMeasuredRolloutConfigurationInput) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *FlagMeasuredRolloutConfigurationInput) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *FlagMeasuredRolloutConfigurationInput) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


