# ExperimentInfoRepItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to [**MetricListingRep**](MetricListingRep.md) |  | [optional] 
**Environments** | Pointer to **[]string** |  | [optional] 
**EnvironmentSettings** | Pointer to [**map[string]ExperimentInfoRepEnvironmentSettings**](ExperimentInfoRepEnvironmentSettings.md) |  | [optional] 

## Methods

### NewExperimentInfoRepItems

`func NewExperimentInfoRepItems() *ExperimentInfoRepItems`

NewExperimentInfoRepItems instantiates a new ExperimentInfoRepItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentInfoRepItemsWithDefaults

`func NewExperimentInfoRepItemsWithDefaults() *ExperimentInfoRepItems`

NewExperimentInfoRepItemsWithDefaults instantiates a new ExperimentInfoRepItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *ExperimentInfoRepItems) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *ExperimentInfoRepItems) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *ExperimentInfoRepItems) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.

### HasMetricKey

`func (o *ExperimentInfoRepItems) HasMetricKey() bool`

HasMetricKey returns a boolean if a field has been set.

### GetMetric

`func (o *ExperimentInfoRepItems) GetMetric() MetricListingRep`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExperimentInfoRepItems) GetMetricOk() (*MetricListingRep, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExperimentInfoRepItems) SetMetric(v MetricListingRep)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ExperimentInfoRepItems) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetEnvironments

`func (o *ExperimentInfoRepItems) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ExperimentInfoRepItems) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ExperimentInfoRepItems) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ExperimentInfoRepItems) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetEnvironmentSettings

`func (o *ExperimentInfoRepItems) GetEnvironmentSettings() map[string]ExperimentInfoRepEnvironmentSettings`

GetEnvironmentSettings returns the EnvironmentSettings field if non-nil, zero value otherwise.

### GetEnvironmentSettingsOk

`func (o *ExperimentInfoRepItems) GetEnvironmentSettingsOk() (*map[string]ExperimentInfoRepEnvironmentSettings, bool)`

GetEnvironmentSettingsOk returns a tuple with the EnvironmentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSettings

`func (o *ExperimentInfoRepItems) SetEnvironmentSettings(v map[string]ExperimentInfoRepEnvironmentSettings)`

SetEnvironmentSettings sets EnvironmentSettings field to given value.

### HasEnvironmentSettings

`func (o *ExperimentInfoRepItems) HasEnvironmentSettings() bool`

HasEnvironmentSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


