# IterationRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The iteration ID | [optional] 
**Hypothesis** | **string** | The expected outcome of this experiment | 
**Status** | **string** | The status of the iteration: &lt;code&gt;not_started&lt;/code&gt;, &lt;code&gt;running&lt;/code&gt;, &lt;code&gt;stopped&lt;/code&gt; | 
**CreatedAt** | **int64** |  | 
**StartedAt** | Pointer to **int64** |  | [optional] 
**EndedAt** | Pointer to **int64** |  | [optional] 
**WinningTreatmentId** | Pointer to **string** | The ID of the treatment chosen when the experiment stopped | [optional] 
**WinningReason** | Pointer to **string** | The reason you stopped the experiment | [optional] 
**CanReshuffleTraffic** | Pointer to **bool** | Whether the experiment may reassign traffic to different variations when the experiment audience changes (true) or must keep all traffic assigned to its initial variation (false). | [optional] 
**Flags** | Pointer to [**map[string]FlagRep**](FlagRep.md) | Details on the flag used in this experiment | [optional] 
**PrimaryMetric** | Pointer to [**DependentMetricOrMetricGroupRep**](DependentMetricOrMetricGroupRep.md) |  | [optional] 
**RandomizationUnit** | Pointer to **string** | The unit of randomization for this iteration | [optional] 
**Treatments** | Pointer to [**[]TreatmentRep**](TreatmentRep.md) | Details on the variations you are testing in the experiment | [optional] 
**SecondaryMetrics** | Pointer to [**[]MetricV2Rep**](MetricV2Rep.md) | Details on the secondary metrics for this experiment | [optional] 

## Methods

### NewIterationRep

`func NewIterationRep(hypothesis string, status string, createdAt int64, ) *IterationRep`

NewIterationRep instantiates a new IterationRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationRepWithDefaults

`func NewIterationRepWithDefaults() *IterationRep`

NewIterationRepWithDefaults instantiates a new IterationRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IterationRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IterationRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IterationRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IterationRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHypothesis

`func (o *IterationRep) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *IterationRep) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *IterationRep) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.


### GetStatus

`func (o *IterationRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IterationRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IterationRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *IterationRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IterationRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IterationRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *IterationRep) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IterationRep) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IterationRep) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IterationRep) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *IterationRep) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *IterationRep) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *IterationRep) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *IterationRep) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetWinningTreatmentId

`func (o *IterationRep) GetWinningTreatmentId() string`

GetWinningTreatmentId returns the WinningTreatmentId field if non-nil, zero value otherwise.

### GetWinningTreatmentIdOk

`func (o *IterationRep) GetWinningTreatmentIdOk() (*string, bool)`

GetWinningTreatmentIdOk returns a tuple with the WinningTreatmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningTreatmentId

`func (o *IterationRep) SetWinningTreatmentId(v string)`

SetWinningTreatmentId sets WinningTreatmentId field to given value.

### HasWinningTreatmentId

`func (o *IterationRep) HasWinningTreatmentId() bool`

HasWinningTreatmentId returns a boolean if a field has been set.

### GetWinningReason

`func (o *IterationRep) GetWinningReason() string`

GetWinningReason returns the WinningReason field if non-nil, zero value otherwise.

### GetWinningReasonOk

`func (o *IterationRep) GetWinningReasonOk() (*string, bool)`

GetWinningReasonOk returns a tuple with the WinningReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningReason

`func (o *IterationRep) SetWinningReason(v string)`

SetWinningReason sets WinningReason field to given value.

### HasWinningReason

`func (o *IterationRep) HasWinningReason() bool`

HasWinningReason returns a boolean if a field has been set.

### GetCanReshuffleTraffic

`func (o *IterationRep) GetCanReshuffleTraffic() bool`

GetCanReshuffleTraffic returns the CanReshuffleTraffic field if non-nil, zero value otherwise.

### GetCanReshuffleTrafficOk

`func (o *IterationRep) GetCanReshuffleTrafficOk() (*bool, bool)`

GetCanReshuffleTrafficOk returns a tuple with the CanReshuffleTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReshuffleTraffic

`func (o *IterationRep) SetCanReshuffleTraffic(v bool)`

SetCanReshuffleTraffic sets CanReshuffleTraffic field to given value.

### HasCanReshuffleTraffic

`func (o *IterationRep) HasCanReshuffleTraffic() bool`

HasCanReshuffleTraffic returns a boolean if a field has been set.

### GetFlags

`func (o *IterationRep) GetFlags() map[string]FlagRep`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *IterationRep) GetFlagsOk() (*map[string]FlagRep, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *IterationRep) SetFlags(v map[string]FlagRep)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *IterationRep) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetPrimaryMetric

`func (o *IterationRep) GetPrimaryMetric() DependentMetricOrMetricGroupRep`

GetPrimaryMetric returns the PrimaryMetric field if non-nil, zero value otherwise.

### GetPrimaryMetricOk

`func (o *IterationRep) GetPrimaryMetricOk() (*DependentMetricOrMetricGroupRep, bool)`

GetPrimaryMetricOk returns a tuple with the PrimaryMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetric

`func (o *IterationRep) SetPrimaryMetric(v DependentMetricOrMetricGroupRep)`

SetPrimaryMetric sets PrimaryMetric field to given value.

### HasPrimaryMetric

`func (o *IterationRep) HasPrimaryMetric() bool`

HasPrimaryMetric returns a boolean if a field has been set.

### GetRandomizationUnit

`func (o *IterationRep) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *IterationRep) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *IterationRep) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *IterationRep) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetTreatments

`func (o *IterationRep) GetTreatments() []TreatmentRep`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *IterationRep) GetTreatmentsOk() (*[]TreatmentRep, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatments

`func (o *IterationRep) SetTreatments(v []TreatmentRep)`

SetTreatments sets Treatments field to given value.

### HasTreatments

`func (o *IterationRep) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *IterationRep) GetSecondaryMetrics() []MetricV2Rep`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *IterationRep) GetSecondaryMetricsOk() (*[]MetricV2Rep, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *IterationRep) SetSecondaryMetrics(v []MetricV2Rep)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *IterationRep) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


