# VariationOrRolloutRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variation** | Pointer to **int32** | The index of the variation, from the array of variations for this flag | [optional] 
**Rollout** | Pointer to [**Rollout**](Rollout.md) |  | [optional] 

## Methods

### NewVariationOrRolloutRep

`func NewVariationOrRolloutRep() *VariationOrRolloutRep`

NewVariationOrRolloutRep instantiates a new VariationOrRolloutRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationOrRolloutRepWithDefaults

`func NewVariationOrRolloutRepWithDefaults() *VariationOrRolloutRep`

NewVariationOrRolloutRepWithDefaults instantiates a new VariationOrRolloutRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariation

`func (o *VariationOrRolloutRep) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *VariationOrRolloutRep) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *VariationOrRolloutRep) SetVariation(v int32)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *VariationOrRolloutRep) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetRollout

`func (o *VariationOrRolloutRep) GetRollout() Rollout`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *VariationOrRolloutRep) GetRolloutOk() (*Rollout, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *VariationOrRolloutRep) SetRollout(v Rollout)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *VariationOrRolloutRep) HasRollout() bool`

HasRollout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


