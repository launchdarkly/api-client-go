# MetricGroupRepExpandableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experiments** | Pointer to [**[]DependentExperimentRep**](DependentExperimentRep.md) |  | [optional] 
**ExperimentCount** | Pointer to **int32** | The number of experiments using this metric group | [optional] 

## Methods

### NewMetricGroupRepExpandableProperties

`func NewMetricGroupRepExpandableProperties() *MetricGroupRepExpandableProperties`

NewMetricGroupRepExpandableProperties instantiates a new MetricGroupRepExpandableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGroupRepExpandablePropertiesWithDefaults

`func NewMetricGroupRepExpandablePropertiesWithDefaults() *MetricGroupRepExpandableProperties`

NewMetricGroupRepExpandablePropertiesWithDefaults instantiates a new MetricGroupRepExpandableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperiments

`func (o *MetricGroupRepExpandableProperties) GetExperiments() []DependentExperimentRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *MetricGroupRepExpandableProperties) GetExperimentsOk() (*[]DependentExperimentRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *MetricGroupRepExpandableProperties) SetExperiments(v []DependentExperimentRep)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *MetricGroupRepExpandableProperties) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetExperimentCount

`func (o *MetricGroupRepExpandableProperties) GetExperimentCount() int32`

GetExperimentCount returns the ExperimentCount field if non-nil, zero value otherwise.

### GetExperimentCountOk

`func (o *MetricGroupRepExpandableProperties) GetExperimentCountOk() (*int32, bool)`

GetExperimentCountOk returns a tuple with the ExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentCount

`func (o *MetricGroupRepExpandableProperties) SetExperimentCount(v int32)`

SetExperimentCount sets ExperimentCount field to given value.

### HasExperimentCount

`func (o *MetricGroupRepExpandableProperties) HasExperimentCount() bool`

HasExperimentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


