# ExperimentQuantileResultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TreatmentResults** | Pointer to [**[]TreatmentQuantileResultRep**](TreatmentQuantileResultRep.md) | A list of the results for each treatment | [optional] 
**Percentile** | Pointer to **int64** | The target percentile | [optional] 
**Confidence** | Pointer to **int64** | The confidence level of percentile interval | [optional] 
**MetricSeen** | Pointer to [**MetricSeen**](MetricSeen.md) |  | [optional] 
**ProbabilityOfMismatch** | Pointer to **float32** | The probability of a Sample Ratio Mismatch | [optional] 

## Methods

### NewExperimentQuantileResultsRep

`func NewExperimentQuantileResultsRep() *ExperimentQuantileResultsRep`

NewExperimentQuantileResultsRep instantiates a new ExperimentQuantileResultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentQuantileResultsRepWithDefaults

`func NewExperimentQuantileResultsRepWithDefaults() *ExperimentQuantileResultsRep`

NewExperimentQuantileResultsRepWithDefaults instantiates a new ExperimentQuantileResultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExperimentQuantileResultsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentQuantileResultsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentQuantileResultsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentQuantileResultsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTreatmentResults

`func (o *ExperimentQuantileResultsRep) GetTreatmentResults() []TreatmentQuantileResultRep`

GetTreatmentResults returns the TreatmentResults field if non-nil, zero value otherwise.

### GetTreatmentResultsOk

`func (o *ExperimentQuantileResultsRep) GetTreatmentResultsOk() (*[]TreatmentQuantileResultRep, bool)`

GetTreatmentResultsOk returns a tuple with the TreatmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentResults

`func (o *ExperimentQuantileResultsRep) SetTreatmentResults(v []TreatmentQuantileResultRep)`

SetTreatmentResults sets TreatmentResults field to given value.

### HasTreatmentResults

`func (o *ExperimentQuantileResultsRep) HasTreatmentResults() bool`

HasTreatmentResults returns a boolean if a field has been set.

### GetPercentile

`func (o *ExperimentQuantileResultsRep) GetPercentile() int64`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *ExperimentQuantileResultsRep) GetPercentileOk() (*int64, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *ExperimentQuantileResultsRep) SetPercentile(v int64)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *ExperimentQuantileResultsRep) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### GetConfidence

`func (o *ExperimentQuantileResultsRep) GetConfidence() int64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExperimentQuantileResultsRep) GetConfidenceOk() (*int64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExperimentQuantileResultsRep) SetConfidence(v int64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ExperimentQuantileResultsRep) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetMetricSeen

`func (o *ExperimentQuantileResultsRep) GetMetricSeen() MetricSeen`

GetMetricSeen returns the MetricSeen field if non-nil, zero value otherwise.

### GetMetricSeenOk

`func (o *ExperimentQuantileResultsRep) GetMetricSeenOk() (*MetricSeen, bool)`

GetMetricSeenOk returns a tuple with the MetricSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSeen

`func (o *ExperimentQuantileResultsRep) SetMetricSeen(v MetricSeen)`

SetMetricSeen sets MetricSeen field to given value.

### HasMetricSeen

`func (o *ExperimentQuantileResultsRep) HasMetricSeen() bool`

HasMetricSeen returns a boolean if a field has been set.

### GetProbabilityOfMismatch

`func (o *ExperimentQuantileResultsRep) GetProbabilityOfMismatch() float32`

GetProbabilityOfMismatch returns the ProbabilityOfMismatch field if non-nil, zero value otherwise.

### GetProbabilityOfMismatchOk

`func (o *ExperimentQuantileResultsRep) GetProbabilityOfMismatchOk() (*float32, bool)`

GetProbabilityOfMismatchOk returns a tuple with the ProbabilityOfMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilityOfMismatch

`func (o *ExperimentQuantileResultsRep) SetProbabilityOfMismatch(v float32)`

SetProbabilityOfMismatch sets ProbabilityOfMismatch field to given value.

### HasProbabilityOfMismatch

`func (o *ExperimentQuantileResultsRep) HasProbabilityOfMismatch() bool`

HasProbabilityOfMismatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


