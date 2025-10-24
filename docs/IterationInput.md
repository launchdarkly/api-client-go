# IterationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hypothesis** | **string** | The expected outcome of this experiment | 
**CanReshuffleTraffic** | Pointer to **bool** | Whether to allow the experiment to reassign traffic to different variations when you increase or decrease the traffic in your experiment audience (true) or keep all traffic assigned to its initial variation (false). Defaults to true. | [optional] 
**Metrics** | [**[]MetricInput**](MetricInput.md) |  | 
**PrimarySingleMetricKey** | Pointer to **string** | The key of the primary metric for this experiment. Either &lt;code&gt;primarySingleMetricKey&lt;/code&gt; or &lt;code&gt;primaryFunnelKey&lt;/code&gt; must be present. | [optional] 
**PrimaryFunnelKey** | Pointer to **string** | The key of the primary funnel group for this experiment. Either &lt;code&gt;primarySingleMetricKey&lt;/code&gt; or &lt;code&gt;primaryFunnelKey&lt;/code&gt; must be present. | [optional] 
**Treatments** | [**[]TreatmentInput**](TreatmentInput.md) |  | 
**Flags** | [**map[string]FlagInput**](FlagInput.md) |  | 
**RandomizationUnit** | Pointer to **string** | The unit of randomization for this iteration. Defaults to user. | [optional] 
**CovarianceId** | Pointer to **string** | The ID of the covariance CSV | [optional] 
**Attributes** | Pointer to **[]string** | The attributes that this iteration&#39;s results can be sliced by | [optional] 

## Methods

### NewIterationInput

`func NewIterationInput(hypothesis string, metrics []MetricInput, treatments []TreatmentInput, flags map[string]FlagInput, ) *IterationInput`

NewIterationInput instantiates a new IterationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationInputWithDefaults

`func NewIterationInputWithDefaults() *IterationInput`

NewIterationInputWithDefaults instantiates a new IterationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypothesis

`func (o *IterationInput) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *IterationInput) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *IterationInput) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.


### GetCanReshuffleTraffic

`func (o *IterationInput) GetCanReshuffleTraffic() bool`

GetCanReshuffleTraffic returns the CanReshuffleTraffic field if non-nil, zero value otherwise.

### GetCanReshuffleTrafficOk

`func (o *IterationInput) GetCanReshuffleTrafficOk() (*bool, bool)`

GetCanReshuffleTrafficOk returns a tuple with the CanReshuffleTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReshuffleTraffic

`func (o *IterationInput) SetCanReshuffleTraffic(v bool)`

SetCanReshuffleTraffic sets CanReshuffleTraffic field to given value.

### HasCanReshuffleTraffic

`func (o *IterationInput) HasCanReshuffleTraffic() bool`

HasCanReshuffleTraffic returns a boolean if a field has been set.

### GetMetrics

`func (o *IterationInput) GetMetrics() []MetricInput`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *IterationInput) GetMetricsOk() (*[]MetricInput, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *IterationInput) SetMetrics(v []MetricInput)`

SetMetrics sets Metrics field to given value.


### GetPrimarySingleMetricKey

`func (o *IterationInput) GetPrimarySingleMetricKey() string`

GetPrimarySingleMetricKey returns the PrimarySingleMetricKey field if non-nil, zero value otherwise.

### GetPrimarySingleMetricKeyOk

`func (o *IterationInput) GetPrimarySingleMetricKeyOk() (*string, bool)`

GetPrimarySingleMetricKeyOk returns a tuple with the PrimarySingleMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySingleMetricKey

`func (o *IterationInput) SetPrimarySingleMetricKey(v string)`

SetPrimarySingleMetricKey sets PrimarySingleMetricKey field to given value.

### HasPrimarySingleMetricKey

`func (o *IterationInput) HasPrimarySingleMetricKey() bool`

HasPrimarySingleMetricKey returns a boolean if a field has been set.

### GetPrimaryFunnelKey

`func (o *IterationInput) GetPrimaryFunnelKey() string`

GetPrimaryFunnelKey returns the PrimaryFunnelKey field if non-nil, zero value otherwise.

### GetPrimaryFunnelKeyOk

`func (o *IterationInput) GetPrimaryFunnelKeyOk() (*string, bool)`

GetPrimaryFunnelKeyOk returns a tuple with the PrimaryFunnelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFunnelKey

`func (o *IterationInput) SetPrimaryFunnelKey(v string)`

SetPrimaryFunnelKey sets PrimaryFunnelKey field to given value.

### HasPrimaryFunnelKey

`func (o *IterationInput) HasPrimaryFunnelKey() bool`

HasPrimaryFunnelKey returns a boolean if a field has been set.

### GetTreatments

`func (o *IterationInput) GetTreatments() []TreatmentInput`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *IterationInput) GetTreatmentsOk() (*[]TreatmentInput, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatments

`func (o *IterationInput) SetTreatments(v []TreatmentInput)`

SetTreatments sets Treatments field to given value.


### GetFlags

`func (o *IterationInput) GetFlags() map[string]FlagInput`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *IterationInput) GetFlagsOk() (*map[string]FlagInput, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *IterationInput) SetFlags(v map[string]FlagInput)`

SetFlags sets Flags field to given value.


### GetRandomizationUnit

`func (o *IterationInput) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *IterationInput) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *IterationInput) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *IterationInput) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetCovarianceId

`func (o *IterationInput) GetCovarianceId() string`

GetCovarianceId returns the CovarianceId field if non-nil, zero value otherwise.

### GetCovarianceIdOk

`func (o *IterationInput) GetCovarianceIdOk() (*string, bool)`

GetCovarianceIdOk returns a tuple with the CovarianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovarianceId

`func (o *IterationInput) SetCovarianceId(v string)`

SetCovarianceId sets CovarianceId field to given value.

### HasCovarianceId

`func (o *IterationInput) HasCovarianceId() bool`

HasCovarianceId returns a boolean if a field has been set.

### GetAttributes

`func (o *IterationInput) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IterationInput) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IterationInput) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IterationInput) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


