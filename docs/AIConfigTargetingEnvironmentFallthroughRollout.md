# AIConfigTargetingEnvironmentFallthroughRollout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketBy** | Pointer to **string** |  | [optional] 
**ContextKind** | **string** |  | 
**ExperimentAllocation** | Pointer to [**AIConfigTargetingEnvironmentFallthroughRolloutExperimentationAllocation**](AIConfigTargetingEnvironmentFallthroughRolloutExperimentationAllocation.md) |  | [optional] 
**Seed** | Pointer to **int32** |  | [optional] 
**Variations** | [**[]AIConfigTargetingEnvironmentFallthroughRolloutVariation**](AIConfigTargetingEnvironmentFallthroughRolloutVariation.md) |  | 

## Methods

### NewAIConfigTargetingEnvironmentFallthroughRollout

`func NewAIConfigTargetingEnvironmentFallthroughRollout(contextKind string, variations []AIConfigTargetingEnvironmentFallthroughRolloutVariation, ) *AIConfigTargetingEnvironmentFallthroughRollout`

NewAIConfigTargetingEnvironmentFallthroughRollout instantiates a new AIConfigTargetingEnvironmentFallthroughRollout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigTargetingEnvironmentFallthroughRolloutWithDefaults

`func NewAIConfigTargetingEnvironmentFallthroughRolloutWithDefaults() *AIConfigTargetingEnvironmentFallthroughRollout`

NewAIConfigTargetingEnvironmentFallthroughRolloutWithDefaults instantiates a new AIConfigTargetingEnvironmentFallthroughRollout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketBy

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.

### GetContextKind

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.


### GetExperimentAllocation

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetExperimentAllocation() AIConfigTargetingEnvironmentFallthroughRolloutExperimentationAllocation`

GetExperimentAllocation returns the ExperimentAllocation field if non-nil, zero value otherwise.

### GetExperimentAllocationOk

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetExperimentAllocationOk() (*AIConfigTargetingEnvironmentFallthroughRolloutExperimentationAllocation, bool)`

GetExperimentAllocationOk returns a tuple with the ExperimentAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentAllocation

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) SetExperimentAllocation(v AIConfigTargetingEnvironmentFallthroughRolloutExperimentationAllocation)`

SetExperimentAllocation sets ExperimentAllocation field to given value.

### HasExperimentAllocation

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) HasExperimentAllocation() bool`

HasExperimentAllocation returns a boolean if a field has been set.

### GetSeed

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetVariations

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetVariations() []AIConfigTargetingEnvironmentFallthroughRolloutVariation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) GetVariationsOk() (*[]AIConfigTargetingEnvironmentFallthroughRolloutVariation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *AIConfigTargetingEnvironmentFallthroughRollout) SetVariations(v []AIConfigTargetingEnvironmentFallthroughRolloutVariation)`

SetVariations sets Variations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


