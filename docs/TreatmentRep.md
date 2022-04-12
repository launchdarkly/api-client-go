# TreatmentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**AllocationPercent** | [**NullDecimal**](NullDecimal.md) |  | 
**Baseline** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**[]ParameterRep**](ParameterRep.md) |  | [optional] 

## Methods

### NewTreatmentRep

`func NewTreatmentRep(name string, allocationPercent NullDecimal, ) *TreatmentRep`

NewTreatmentRep instantiates a new TreatmentRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreatmentRepWithDefaults

`func NewTreatmentRepWithDefaults() *TreatmentRep`

NewTreatmentRepWithDefaults instantiates a new TreatmentRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TreatmentRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TreatmentRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TreatmentRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TreatmentRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TreatmentRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TreatmentRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TreatmentRep) SetName(v string)`

SetName sets Name field to given value.


### GetAllocationPercent

`func (o *TreatmentRep) GetAllocationPercent() NullDecimal`

GetAllocationPercent returns the AllocationPercent field if non-nil, zero value otherwise.

### GetAllocationPercentOk

`func (o *TreatmentRep) GetAllocationPercentOk() (*NullDecimal, bool)`

GetAllocationPercentOk returns a tuple with the AllocationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationPercent

`func (o *TreatmentRep) SetAllocationPercent(v NullDecimal)`

SetAllocationPercent sets AllocationPercent field to given value.


### GetBaseline

`func (o *TreatmentRep) GetBaseline() bool`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *TreatmentRep) GetBaselineOk() (*bool, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *TreatmentRep) SetBaseline(v bool)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *TreatmentRep) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetParameters

`func (o *TreatmentRep) GetParameters() []ParameterRep`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TreatmentRep) GetParametersOk() (*[]ParameterRep, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TreatmentRep) SetParameters(v []ParameterRep)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TreatmentRep) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


