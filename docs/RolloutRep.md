# RolloutRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variations** | [**[]WeightedVariationRep**](WeightedVariationRep.md) |  | 
**ExperimentAllocation** | Pointer to [**ExperimentAllocationRep**](ExperimentAllocationRep.md) |  | [optional] 
**Seed** | Pointer to **int32** |  | [optional] 
**BucketBy** | Pointer to **string** |  | [optional] 

## Methods

### NewRolloutRep

`func NewRolloutRep(variations []WeightedVariationRep, ) *RolloutRep`

NewRolloutRep instantiates a new RolloutRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRepWithDefaults

`func NewRolloutRepWithDefaults() *RolloutRep`

NewRolloutRepWithDefaults instantiates a new RolloutRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariations

`func (o *RolloutRep) GetVariations() []WeightedVariationRep`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *RolloutRep) GetVariationsOk() (*[]WeightedVariationRep, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *RolloutRep) SetVariations(v []WeightedVariationRep)`

SetVariations sets Variations field to given value.


### GetExperimentAllocation

`func (o *RolloutRep) GetExperimentAllocation() ExperimentAllocationRep`

GetExperimentAllocation returns the ExperimentAllocation field if non-nil, zero value otherwise.

### GetExperimentAllocationOk

`func (o *RolloutRep) GetExperimentAllocationOk() (*ExperimentAllocationRep, bool)`

GetExperimentAllocationOk returns a tuple with the ExperimentAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentAllocation

`func (o *RolloutRep) SetExperimentAllocation(v ExperimentAllocationRep)`

SetExperimentAllocation sets ExperimentAllocation field to given value.

### HasExperimentAllocation

`func (o *RolloutRep) HasExperimentAllocation() bool`

HasExperimentAllocation returns a boolean if a field has been set.

### GetSeed

`func (o *RolloutRep) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *RolloutRep) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *RolloutRep) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *RolloutRep) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetBucketBy

`func (o *RolloutRep) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *RolloutRep) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *RolloutRep) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *RolloutRep) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


