# FlagEventImpactRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** | The size of the flag event impact. Sizes are defined as: none (0%), small (0-20%), medium (20-80%), large (&gt;80%) | [optional] 
**Percentage** | Pointer to **float32** | The percentage of the flag event impact | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**EvaluationsSummary** | Pointer to [**EvaluationsSummary**](EvaluationsSummary.md) |  | [optional] 

## Methods

### NewFlagEventImpactRep

`func NewFlagEventImpactRep() *FlagEventImpactRep`

NewFlagEventImpactRep instantiates a new FlagEventImpactRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagEventImpactRepWithDefaults

`func NewFlagEventImpactRepWithDefaults() *FlagEventImpactRep`

NewFlagEventImpactRepWithDefaults instantiates a new FlagEventImpactRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *FlagEventImpactRep) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FlagEventImpactRep) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FlagEventImpactRep) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *FlagEventImpactRep) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPercentage

`func (o *FlagEventImpactRep) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *FlagEventImpactRep) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *FlagEventImpactRep) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *FlagEventImpactRep) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetReason

`func (o *FlagEventImpactRep) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FlagEventImpactRep) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FlagEventImpactRep) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *FlagEventImpactRep) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetEvaluationsSummary

`func (o *FlagEventImpactRep) GetEvaluationsSummary() EvaluationsSummary`

GetEvaluationsSummary returns the EvaluationsSummary field if non-nil, zero value otherwise.

### GetEvaluationsSummaryOk

`func (o *FlagEventImpactRep) GetEvaluationsSummaryOk() (*EvaluationsSummary, bool)`

GetEvaluationsSummaryOk returns a tuple with the EvaluationsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationsSummary

`func (o *FlagEventImpactRep) SetEvaluationsSummary(v EvaluationsSummary)`

SetEvaluationsSummary sets EvaluationsSummary field to given value.

### HasEvaluationsSummary

`func (o *FlagEventImpactRep) HasEvaluationsSummary() bool`

HasEvaluationsSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


