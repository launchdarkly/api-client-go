# Rollout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variations** | [**[]WeightedVariation**](WeightedVariation.md) |  | 
**ExperimentAllocation** | Pointer to [**ExperimentAllocationRep**](ExperimentAllocationRep.md) |  | [optional] 
**Seed** | Pointer to **int32** |  | [optional] 
**BucketBy** | Pointer to **string** |  | [optional] 
**ContextKind** | Pointer to **string** |  | [optional] 

## Methods

### NewRollout

`func NewRollout(variations []WeightedVariation, ) *Rollout`

NewRollout instantiates a new Rollout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutWithDefaults

`func NewRolloutWithDefaults() *Rollout`

NewRolloutWithDefaults instantiates a new Rollout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariations

`func (o *Rollout) GetVariations() []WeightedVariation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *Rollout) GetVariationsOk() (*[]WeightedVariation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *Rollout) SetVariations(v []WeightedVariation)`

SetVariations sets Variations field to given value.


### GetExperimentAllocation

`func (o *Rollout) GetExperimentAllocation() ExperimentAllocationRep`

GetExperimentAllocation returns the ExperimentAllocation field if non-nil, zero value otherwise.

### GetExperimentAllocationOk

`func (o *Rollout) GetExperimentAllocationOk() (*ExperimentAllocationRep, bool)`

GetExperimentAllocationOk returns a tuple with the ExperimentAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentAllocation

`func (o *Rollout) SetExperimentAllocation(v ExperimentAllocationRep)`

SetExperimentAllocation sets ExperimentAllocation field to given value.

### HasExperimentAllocation

`func (o *Rollout) HasExperimentAllocation() bool`

HasExperimentAllocation returns a boolean if a field has been set.

### GetSeed

`func (o *Rollout) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *Rollout) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *Rollout) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *Rollout) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetBucketBy

`func (o *Rollout) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *Rollout) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *Rollout) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *Rollout) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.

### GetContextKind

`func (o *Rollout) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *Rollout) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *Rollout) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *Rollout) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


