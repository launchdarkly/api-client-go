# ProgressiveReleaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutContextKindKey** | Pointer to **string** | Context kind key to use as the randomization unit for the rollout | [optional] 
**Stages** | Pointer to [**[]ReleasePolicyStage**](ReleasePolicyStage.md) | List of stages | [optional] 

## Methods

### NewProgressiveReleaseConfig

`func NewProgressiveReleaseConfig() *ProgressiveReleaseConfig`

NewProgressiveReleaseConfig instantiates a new ProgressiveReleaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressiveReleaseConfigWithDefaults

`func NewProgressiveReleaseConfigWithDefaults() *ProgressiveReleaseConfig`

NewProgressiveReleaseConfigWithDefaults instantiates a new ProgressiveReleaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutContextKindKey

`func (o *ProgressiveReleaseConfig) GetRolloutContextKindKey() string`

GetRolloutContextKindKey returns the RolloutContextKindKey field if non-nil, zero value otherwise.

### GetRolloutContextKindKeyOk

`func (o *ProgressiveReleaseConfig) GetRolloutContextKindKeyOk() (*string, bool)`

GetRolloutContextKindKeyOk returns a tuple with the RolloutContextKindKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutContextKindKey

`func (o *ProgressiveReleaseConfig) SetRolloutContextKindKey(v string)`

SetRolloutContextKindKey sets RolloutContextKindKey field to given value.

### HasRolloutContextKindKey

`func (o *ProgressiveReleaseConfig) HasRolloutContextKindKey() bool`

HasRolloutContextKindKey returns a boolean if a field has been set.

### GetStages

`func (o *ProgressiveReleaseConfig) GetStages() []ReleasePolicyStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *ProgressiveReleaseConfig) GetStagesOk() (*[]ReleasePolicyStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *ProgressiveReleaseConfig) SetStages(v []ReleasePolicyStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *ProgressiveReleaseConfig) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


