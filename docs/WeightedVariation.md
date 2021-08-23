# WeightedVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variation** | **int32** |  | 
**Weight** | **int32** |  | 
**Untracked** | Pointer to **bool** |  | [optional] 

## Methods

### NewWeightedVariation

`func NewWeightedVariation(variation int32, weight int32, ) *WeightedVariation`

NewWeightedVariation instantiates a new WeightedVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeightedVariationWithDefaults

`func NewWeightedVariationWithDefaults() *WeightedVariation`

NewWeightedVariationWithDefaults instantiates a new WeightedVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariation

`func (o *WeightedVariation) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *WeightedVariation) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *WeightedVariation) SetVariation(v int32)`

SetVariation sets Variation field to given value.


### GetWeight

`func (o *WeightedVariation) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WeightedVariation) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WeightedVariation) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetUntracked

`func (o *WeightedVariation) GetUntracked() bool`

GetUntracked returns the Untracked field if non-nil, zero value otherwise.

### GetUntrackedOk

`func (o *WeightedVariation) GetUntrackedOk() (*bool, bool)`

GetUntrackedOk returns a tuple with the Untracked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntracked

`func (o *WeightedVariation) SetUntracked(v bool)`

SetUntracked sets Untracked field to given value.

### HasUntracked

`func (o *WeightedVariation) HasUntracked() bool`

HasUntracked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


