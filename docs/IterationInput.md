# IterationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hypothesis** | **string** | The expected outcome of this experiment | 
**CanReshuffleTraffic** | Pointer to **bool** | Whether to allow the experiment to reassign traffic to different variations when you increase or decrease the traffic in your experiment audience (true) or keep all traffic assigned to its initial variation (false). Defaults to true. | [optional] 
**Metrics** | [**[]MetricInput**](MetricInput.md) |  | 
**Treatments** | [**[]TreatmentInput**](TreatmentInput.md) |  | 
**Flags** | [**map[string]FlagInput**](FlagInput.md) |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


