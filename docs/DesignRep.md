# DesignRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hypothesis** | **string** | The expected outcome of this experiment | 
**CanReshuffleTraffic** | Pointer to **bool** | Whether the experiment can reassign traffic to different variations when you increase or decrease the traffic in your experiment audience (true) or keep all traffic assigned to its initial variation (false). | [optional] 
**Flags** | Pointer to [**map[string]FlagRep**](FlagRep.md) | Details on the flag used in this experiment | [optional] 
**PrimaryMetric** | Pointer to [**DependentMetricOrMetricGroupRep**](DependentMetricOrMetricGroupRep.md) |  | [optional] 
**RandomizationUnit** | Pointer to **string** | The unit of randomization for this iteration | [optional] 
**Treatments** | Pointer to [**[]TreatmentRep**](TreatmentRep.md) | Details on the variations you are testing in the experiment | [optional] 
**SecondaryMetrics** | Pointer to [**[]MetricV2Rep**](MetricV2Rep.md) | Details on the secondary metrics for this experiment | [optional] 

## Methods

### NewDesignRep

`func NewDesignRep(hypothesis string, ) *DesignRep`

NewDesignRep instantiates a new DesignRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesignRepWithDefaults

`func NewDesignRepWithDefaults() *DesignRep`

NewDesignRepWithDefaults instantiates a new DesignRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypothesis

`func (o *DesignRep) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *DesignRep) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *DesignRep) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.


### GetCanReshuffleTraffic

`func (o *DesignRep) GetCanReshuffleTraffic() bool`

GetCanReshuffleTraffic returns the CanReshuffleTraffic field if non-nil, zero value otherwise.

### GetCanReshuffleTrafficOk

`func (o *DesignRep) GetCanReshuffleTrafficOk() (*bool, bool)`

GetCanReshuffleTrafficOk returns a tuple with the CanReshuffleTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReshuffleTraffic

`func (o *DesignRep) SetCanReshuffleTraffic(v bool)`

SetCanReshuffleTraffic sets CanReshuffleTraffic field to given value.

### HasCanReshuffleTraffic

`func (o *DesignRep) HasCanReshuffleTraffic() bool`

HasCanReshuffleTraffic returns a boolean if a field has been set.

### GetFlags

`func (o *DesignRep) GetFlags() map[string]FlagRep`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *DesignRep) GetFlagsOk() (*map[string]FlagRep, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *DesignRep) SetFlags(v map[string]FlagRep)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *DesignRep) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetPrimaryMetric

`func (o *DesignRep) GetPrimaryMetric() DependentMetricOrMetricGroupRep`

GetPrimaryMetric returns the PrimaryMetric field if non-nil, zero value otherwise.

### GetPrimaryMetricOk

`func (o *DesignRep) GetPrimaryMetricOk() (*DependentMetricOrMetricGroupRep, bool)`

GetPrimaryMetricOk returns a tuple with the PrimaryMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetric

`func (o *DesignRep) SetPrimaryMetric(v DependentMetricOrMetricGroupRep)`

SetPrimaryMetric sets PrimaryMetric field to given value.

### HasPrimaryMetric

`func (o *DesignRep) HasPrimaryMetric() bool`

HasPrimaryMetric returns a boolean if a field has been set.

### GetRandomizationUnit

`func (o *DesignRep) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *DesignRep) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *DesignRep) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *DesignRep) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetTreatments

`func (o *DesignRep) GetTreatments() []TreatmentRep`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *DesignRep) GetTreatmentsOk() (*[]TreatmentRep, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatments

`func (o *DesignRep) SetTreatments(v []TreatmentRep)`

SetTreatments sets Treatments field to given value.

### HasTreatments

`func (o *DesignRep) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *DesignRep) GetSecondaryMetrics() []MetricV2Rep`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *DesignRep) GetSecondaryMetricsOk() (*[]MetricV2Rep, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *DesignRep) SetSecondaryMetrics(v []MetricV2Rep)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *DesignRep) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


