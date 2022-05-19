# ExperimentBayesianResultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**TreatmentResults** | Pointer to [**[]TreatmentResultRep**](TreatmentResultRep.md) |  | [optional] 
**MetricSeen** | Pointer to [**MetricSeen**](MetricSeen.md) |  | [optional] 

## Methods

### NewExperimentBayesianResultsRep

`func NewExperimentBayesianResultsRep() *ExperimentBayesianResultsRep`

NewExperimentBayesianResultsRep instantiates a new ExperimentBayesianResultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentBayesianResultsRepWithDefaults

`func NewExperimentBayesianResultsRepWithDefaults() *ExperimentBayesianResultsRep`

NewExperimentBayesianResultsRepWithDefaults instantiates a new ExperimentBayesianResultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExperimentBayesianResultsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentBayesianResultsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentBayesianResultsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentBayesianResultsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTreatmentResults

`func (o *ExperimentBayesianResultsRep) GetTreatmentResults() []TreatmentResultRep`

GetTreatmentResults returns the TreatmentResults field if non-nil, zero value otherwise.

### GetTreatmentResultsOk

`func (o *ExperimentBayesianResultsRep) GetTreatmentResultsOk() (*[]TreatmentResultRep, bool)`

GetTreatmentResultsOk returns a tuple with the TreatmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentResults

`func (o *ExperimentBayesianResultsRep) SetTreatmentResults(v []TreatmentResultRep)`

SetTreatmentResults sets TreatmentResults field to given value.

### HasTreatmentResults

`func (o *ExperimentBayesianResultsRep) HasTreatmentResults() bool`

HasTreatmentResults returns a boolean if a field has been set.

### GetMetricSeen

`func (o *ExperimentBayesianResultsRep) GetMetricSeen() MetricSeen`

GetMetricSeen returns the MetricSeen field if non-nil, zero value otherwise.

### GetMetricSeenOk

`func (o *ExperimentBayesianResultsRep) GetMetricSeenOk() (*MetricSeen, bool)`

GetMetricSeenOk returns a tuple with the MetricSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSeen

`func (o *ExperimentBayesianResultsRep) SetMetricSeen(v MetricSeen)`

SetMetricSeen sets MetricSeen field to given value.

### HasMetricSeen

`func (o *ExperimentBayesianResultsRep) HasMetricSeen() bool`

HasMetricSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


