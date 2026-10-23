# AIConfigPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVariationKey** | Pointer to **string** | Key of the variation this config serves by default | [optional] 
**Description** | Pointer to **string** | A description of the AI Config | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this resource | [optional] 
**MaintainerTeamKey** | Pointer to **string** | The key of the team that maintains this resource | [optional] 
**Name** | Pointer to **string** | A human-readable name for the AI Config | [optional] 
**Tags** | Pointer to **[]string** | Tags for this resource | [optional] 
**EvaluationMetricKey** | Pointer to **string** | Evaluation metric key for this config | [optional] 
**IsInverted** | Pointer to **bool** | Whether the evaluation metric is inverted, meaning a lower value is better if set as true | [optional] 

## Methods

### NewAIConfigPatch

`func NewAIConfigPatch() *AIConfigPatch`

NewAIConfigPatch instantiates a new AIConfigPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigPatchWithDefaults

`func NewAIConfigPatchWithDefaults() *AIConfigPatch`

NewAIConfigPatchWithDefaults instantiates a new AIConfigPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVariationKey

`func (o *AIConfigPatch) GetDefaultVariationKey() string`

GetDefaultVariationKey returns the DefaultVariationKey field if non-nil, zero value otherwise.

### GetDefaultVariationKeyOk

`func (o *AIConfigPatch) GetDefaultVariationKeyOk() (*string, bool)`

GetDefaultVariationKeyOk returns a tuple with the DefaultVariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVariationKey

`func (o *AIConfigPatch) SetDefaultVariationKey(v string)`

SetDefaultVariationKey sets DefaultVariationKey field to given value.

### HasDefaultVariationKey

`func (o *AIConfigPatch) HasDefaultVariationKey() bool`

HasDefaultVariationKey returns a boolean if a field has been set.

### GetDescription

`func (o *AIConfigPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIConfigPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIConfigPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIConfigPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainerId

`func (o *AIConfigPatch) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AIConfigPatch) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AIConfigPatch) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AIConfigPatch) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *AIConfigPatch) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *AIConfigPatch) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *AIConfigPatch) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *AIConfigPatch) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetName

`func (o *AIConfigPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfigPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfigPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AIConfigPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *AIConfigPatch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AIConfigPatch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AIConfigPatch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AIConfigPatch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvaluationMetricKey

`func (o *AIConfigPatch) GetEvaluationMetricKey() string`

GetEvaluationMetricKey returns the EvaluationMetricKey field if non-nil, zero value otherwise.

### GetEvaluationMetricKeyOk

`func (o *AIConfigPatch) GetEvaluationMetricKeyOk() (*string, bool)`

GetEvaluationMetricKeyOk returns a tuple with the EvaluationMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMetricKey

`func (o *AIConfigPatch) SetEvaluationMetricKey(v string)`

SetEvaluationMetricKey sets EvaluationMetricKey field to given value.

### HasEvaluationMetricKey

`func (o *AIConfigPatch) HasEvaluationMetricKey() bool`

HasEvaluationMetricKey returns a boolean if a field has been set.

### GetIsInverted

`func (o *AIConfigPatch) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *AIConfigPatch) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *AIConfigPatch) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.

### HasIsInverted

`func (o *AIConfigPatch) HasIsInverted() bool`

HasIsInverted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


