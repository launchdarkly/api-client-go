# DesignExpandableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Treatments** | Pointer to [**[]TreatmentRep**](TreatmentRep.md) | Details on the variations you are testing in the experiment | [optional] 
**SecondaryMetrics** | Pointer to [**[]MetricV2Rep**](MetricV2Rep.md) | Details on the secondary metrics for this experiment | [optional] 

## Methods

### NewDesignExpandableProperties

`func NewDesignExpandableProperties() *DesignExpandableProperties`

NewDesignExpandableProperties instantiates a new DesignExpandableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesignExpandablePropertiesWithDefaults

`func NewDesignExpandablePropertiesWithDefaults() *DesignExpandableProperties`

NewDesignExpandablePropertiesWithDefaults instantiates a new DesignExpandableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTreatments

`func (o *DesignExpandableProperties) GetTreatments() []TreatmentRep`

GetTreatments returns the Treatments field if non-nil, zero value otherwise.

### GetTreatmentsOk

`func (o *DesignExpandableProperties) GetTreatmentsOk() (*[]TreatmentRep, bool)`

GetTreatmentsOk returns a tuple with the Treatments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatments

`func (o *DesignExpandableProperties) SetTreatments(v []TreatmentRep)`

SetTreatments sets Treatments field to given value.

### HasTreatments

`func (o *DesignExpandableProperties) HasTreatments() bool`

HasTreatments returns a boolean if a field has been set.

### GetSecondaryMetrics

`func (o *DesignExpandableProperties) GetSecondaryMetrics() []MetricV2Rep`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *DesignExpandableProperties) GetSecondaryMetricsOk() (*[]MetricV2Rep, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *DesignExpandableProperties) SetSecondaryMetrics(v []MetricV2Rep)`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *DesignExpandableProperties) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


