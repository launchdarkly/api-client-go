# ExperimentExpandableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftIteration** | Pointer to [**IterationRep**](IterationRep.md) |  | [optional] 
**PreviousIterations** | Pointer to [**[]IterationRep**](IterationRep.md) | Details on the previous iterations for this experiment. | [optional] 
**MonitoringProperties** | Pointer to [**MonitoringPropertiesRep**](MonitoringPropertiesRep.md) |  | [optional] 

## Methods

### NewExperimentExpandableProperties

`func NewExperimentExpandableProperties() *ExperimentExpandableProperties`

NewExperimentExpandableProperties instantiates a new ExperimentExpandableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentExpandablePropertiesWithDefaults

`func NewExperimentExpandablePropertiesWithDefaults() *ExperimentExpandableProperties`

NewExperimentExpandablePropertiesWithDefaults instantiates a new ExperimentExpandableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftIteration

`func (o *ExperimentExpandableProperties) GetDraftIteration() IterationRep`

GetDraftIteration returns the DraftIteration field if non-nil, zero value otherwise.

### GetDraftIterationOk

`func (o *ExperimentExpandableProperties) GetDraftIterationOk() (*IterationRep, bool)`

GetDraftIterationOk returns a tuple with the DraftIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftIteration

`func (o *ExperimentExpandableProperties) SetDraftIteration(v IterationRep)`

SetDraftIteration sets DraftIteration field to given value.

### HasDraftIteration

`func (o *ExperimentExpandableProperties) HasDraftIteration() bool`

HasDraftIteration returns a boolean if a field has been set.

### GetPreviousIterations

`func (o *ExperimentExpandableProperties) GetPreviousIterations() []IterationRep`

GetPreviousIterations returns the PreviousIterations field if non-nil, zero value otherwise.

### GetPreviousIterationsOk

`func (o *ExperimentExpandableProperties) GetPreviousIterationsOk() (*[]IterationRep, bool)`

GetPreviousIterationsOk returns a tuple with the PreviousIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIterations

`func (o *ExperimentExpandableProperties) SetPreviousIterations(v []IterationRep)`

SetPreviousIterations sets PreviousIterations field to given value.

### HasPreviousIterations

`func (o *ExperimentExpandableProperties) HasPreviousIterations() bool`

HasPreviousIterations returns a boolean if a field has been set.

### GetMonitoringProperties

`func (o *ExperimentExpandableProperties) GetMonitoringProperties() MonitoringPropertiesRep`

GetMonitoringProperties returns the MonitoringProperties field if non-nil, zero value otherwise.

### GetMonitoringPropertiesOk

`func (o *ExperimentExpandableProperties) GetMonitoringPropertiesOk() (*MonitoringPropertiesRep, bool)`

GetMonitoringPropertiesOk returns a tuple with the MonitoringProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringProperties

`func (o *ExperimentExpandableProperties) SetMonitoringProperties(v MonitoringPropertiesRep)`

SetMonitoringProperties sets MonitoringProperties field to given value.

### HasMonitoringProperties

`func (o *ExperimentExpandableProperties) HasMonitoringProperties() bool`

HasMonitoringProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


