# TreatmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The treatment name | 
**Baseline** | **bool** | Whether this treatment is the baseline to compare other treatments against | 
**AllocationPercent** | **string** | The percentage of traffic allocated to this treatment during the iteration | 
**Parameters** | [**[]TreatmentParameterInput**](TreatmentParameterInput.md) | Details on the flag and variation to use for this treatment | 

## Methods

### NewTreatmentInput

`func NewTreatmentInput(name string, baseline bool, allocationPercent string, parameters []TreatmentParameterInput, ) *TreatmentInput`

NewTreatmentInput instantiates a new TreatmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreatmentInputWithDefaults

`func NewTreatmentInputWithDefaults() *TreatmentInput`

NewTreatmentInputWithDefaults instantiates a new TreatmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TreatmentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TreatmentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TreatmentInput) SetName(v string)`

SetName sets Name field to given value.


### GetBaseline

`func (o *TreatmentInput) GetBaseline() bool`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *TreatmentInput) GetBaselineOk() (*bool, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *TreatmentInput) SetBaseline(v bool)`

SetBaseline sets Baseline field to given value.


### GetAllocationPercent

`func (o *TreatmentInput) GetAllocationPercent() string`

GetAllocationPercent returns the AllocationPercent field if non-nil, zero value otherwise.

### GetAllocationPercentOk

`func (o *TreatmentInput) GetAllocationPercentOk() (*string, bool)`

GetAllocationPercentOk returns a tuple with the AllocationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationPercent

`func (o *TreatmentInput) SetAllocationPercent(v string)`

SetAllocationPercent sets AllocationPercent field to given value.


### GetParameters

`func (o *TreatmentInput) GetParameters() []TreatmentParameterInput`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TreatmentInput) GetParametersOk() (*[]TreatmentParameterInput, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TreatmentInput) SetParameters(v []TreatmentParameterInput)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


