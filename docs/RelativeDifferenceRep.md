# RelativeDifferenceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upper** | Pointer to **float32** | An upper bound of the relative difference between the treatment and the &lt;code&gt;fromTreatmentId&lt;/code&gt; | [optional] 
**Lower** | Pointer to **float32** | A lower bound of the relative difference between the treatment and the &lt;code&gt;fromTreatmentId&lt;/code&gt; | [optional] 
**Estimate** | Pointer to **float32** | A point estimate of the relative difference between the treatment and the &lt;code&gt;fromTreatmentId&lt;/code&gt; | [optional] 
**FromTreatmentId** | Pointer to **string** | The treatment ID of the treatment against which the relative difference is calculated | [optional] 

## Methods

### NewRelativeDifferenceRep

`func NewRelativeDifferenceRep() *RelativeDifferenceRep`

NewRelativeDifferenceRep instantiates a new RelativeDifferenceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelativeDifferenceRepWithDefaults

`func NewRelativeDifferenceRepWithDefaults() *RelativeDifferenceRep`

NewRelativeDifferenceRepWithDefaults instantiates a new RelativeDifferenceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpper

`func (o *RelativeDifferenceRep) GetUpper() float32`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *RelativeDifferenceRep) GetUpperOk() (*float32, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *RelativeDifferenceRep) SetUpper(v float32)`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *RelativeDifferenceRep) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### GetLower

`func (o *RelativeDifferenceRep) GetLower() float32`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *RelativeDifferenceRep) GetLowerOk() (*float32, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *RelativeDifferenceRep) SetLower(v float32)`

SetLower sets Lower field to given value.

### HasLower

`func (o *RelativeDifferenceRep) HasLower() bool`

HasLower returns a boolean if a field has been set.

### GetEstimate

`func (o *RelativeDifferenceRep) GetEstimate() float32`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *RelativeDifferenceRep) GetEstimateOk() (*float32, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *RelativeDifferenceRep) SetEstimate(v float32)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *RelativeDifferenceRep) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetFromTreatmentId

`func (o *RelativeDifferenceRep) GetFromTreatmentId() string`

GetFromTreatmentId returns the FromTreatmentId field if non-nil, zero value otherwise.

### GetFromTreatmentIdOk

`func (o *RelativeDifferenceRep) GetFromTreatmentIdOk() (*string, bool)`

GetFromTreatmentIdOk returns a tuple with the FromTreatmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTreatmentId

`func (o *RelativeDifferenceRep) SetFromTreatmentId(v string)`

SetFromTreatmentId sets FromTreatmentId field to given value.

### HasFromTreatmentId

`func (o *RelativeDifferenceRep) HasFromTreatmentId() bool`

HasFromTreatmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


