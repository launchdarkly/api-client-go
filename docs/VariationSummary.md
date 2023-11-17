# VariationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | **int32** |  | 
**NullRules** | **int32** |  | 
**Targets** | **int32** |  | 
**ContextTargets** | **int32** |  | 
**IsFallthrough** | Pointer to **bool** |  | [optional] 
**IsOff** | Pointer to **bool** |  | [optional] 
**Rollout** | Pointer to **int32** |  | [optional] 
**BucketBy** | Pointer to **string** |  | [optional] 

## Methods

### NewVariationSummary

`func NewVariationSummary(rules int32, nullRules int32, targets int32, contextTargets int32, ) *VariationSummary`

NewVariationSummary instantiates a new VariationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationSummaryWithDefaults

`func NewVariationSummaryWithDefaults() *VariationSummary`

NewVariationSummaryWithDefaults instantiates a new VariationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *VariationSummary) GetRules() int32`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *VariationSummary) GetRulesOk() (*int32, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *VariationSummary) SetRules(v int32)`

SetRules sets Rules field to given value.


### GetNullRules

`func (o *VariationSummary) GetNullRules() int32`

GetNullRules returns the NullRules field if non-nil, zero value otherwise.

### GetNullRulesOk

`func (o *VariationSummary) GetNullRulesOk() (*int32, bool)`

GetNullRulesOk returns a tuple with the NullRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullRules

`func (o *VariationSummary) SetNullRules(v int32)`

SetNullRules sets NullRules field to given value.


### GetTargets

`func (o *VariationSummary) GetTargets() int32`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *VariationSummary) GetTargetsOk() (*int32, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *VariationSummary) SetTargets(v int32)`

SetTargets sets Targets field to given value.


### GetContextTargets

`func (o *VariationSummary) GetContextTargets() int32`

GetContextTargets returns the ContextTargets field if non-nil, zero value otherwise.

### GetContextTargetsOk

`func (o *VariationSummary) GetContextTargetsOk() (*int32, bool)`

GetContextTargetsOk returns a tuple with the ContextTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextTargets

`func (o *VariationSummary) SetContextTargets(v int32)`

SetContextTargets sets ContextTargets field to given value.


### GetIsFallthrough

`func (o *VariationSummary) GetIsFallthrough() bool`

GetIsFallthrough returns the IsFallthrough field if non-nil, zero value otherwise.

### GetIsFallthroughOk

`func (o *VariationSummary) GetIsFallthroughOk() (*bool, bool)`

GetIsFallthroughOk returns a tuple with the IsFallthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFallthrough

`func (o *VariationSummary) SetIsFallthrough(v bool)`

SetIsFallthrough sets IsFallthrough field to given value.

### HasIsFallthrough

`func (o *VariationSummary) HasIsFallthrough() bool`

HasIsFallthrough returns a boolean if a field has been set.

### GetIsOff

`func (o *VariationSummary) GetIsOff() bool`

GetIsOff returns the IsOff field if non-nil, zero value otherwise.

### GetIsOffOk

`func (o *VariationSummary) GetIsOffOk() (*bool, bool)`

GetIsOffOk returns a tuple with the IsOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOff

`func (o *VariationSummary) SetIsOff(v bool)`

SetIsOff sets IsOff field to given value.

### HasIsOff

`func (o *VariationSummary) HasIsOff() bool`

HasIsOff returns a boolean if a field has been set.

### GetRollout

`func (o *VariationSummary) GetRollout() int32`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *VariationSummary) GetRolloutOk() (*int32, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *VariationSummary) SetRollout(v int32)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *VariationSummary) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetBucketBy

`func (o *VariationSummary) GetBucketBy() string`

GetBucketBy returns the BucketBy field if non-nil, zero value otherwise.

### GetBucketByOk

`func (o *VariationSummary) GetBucketByOk() (*string, bool)`

GetBucketByOk returns a tuple with the BucketBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBy

`func (o *VariationSummary) SetBucketBy(v string)`

SetBucketBy sets BucketBy field to given value.

### HasBucketBy

`func (o *VariationSummary) HasBucketBy() bool`

HasBucketBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


