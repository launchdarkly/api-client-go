# IterationExpandableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Treatments** | Pointer to [**[]TreatmentRep**](TreatmentRep.md) |  | [optional] 
**SecondaryMetrics** | Pointer to [**[]MetricRep**](MetricRep.md) |  | [optional] 

## Methods

### NewIterationExpandableProperties

`func NewIterationExpandableProperties() *IterationExpandableProperties`

NewIterationExpandableProperties instantiates a new IterationExpandableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationExpandablePropertiesWithDefaults

`func NewIterationExpandablePropertiesWithDefaults() *IterationExpandableProperties`

NewIterationExpandablePropertiesWithDefaults instantiates a new IterationExpandableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTreatments

`func (o *IterationExpandableProperties) GetTreatments() []TreatmentRep`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *IterationExpandableProperties) GetTreatmentsOk() (*[]TreatmentRep, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatments

`func (o *IterationExpandableProperties) SetTreatments(v []TreatmentRep)`

SetTreatments sets Treatments field to given value.

### HasTreatments

`func (o *IterationExpandableProperties) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *IterationExpandableProperties) GetSecondaryMetrics() []MetricRep`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *IterationExpandableProperties) GetSecondaryMetricsOk() (*[]MetricRep, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *IterationExpandableProperties) SetSecondaryMetrics(v []MetricRep)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *IterationExpandableProperties) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


