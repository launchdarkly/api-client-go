# AIConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Description** | **string** |  | 
**Key** | **string** |  | 
**Maintainer** | Pointer to [**AIConfigMaintainer**](AIConfigMaintainer.md) |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "completion"]
**Name** | **string** |  | 
**Tags** | **[]string** |  | 
**Version** | **int32** |  | 
**Variations** | [**[]AIConfigVariation**](AIConfigVariation.md) |  | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**EvaluationMetricKeys** | Pointer to **[]string** | List of evaluation metric keys for this AI config | [optional] 

## Methods

### NewAIConfig

`func NewAIConfig(description string, key string, name string, tags []string, version int32, variations []AIConfigVariation, createdAt int64, updatedAt int64, ) *AIConfig`

NewAIConfig instantiates a new AIConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigWithDefaults

`func NewAIConfigWithDefaults() *AIConfig`

NewAIConfigWithDefaults instantiates a new AIConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AIConfig) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AIConfig) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AIConfig) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AIConfig) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *AIConfig) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AIConfig) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AIConfig) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AIConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *AIConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetKey

`func (o *AIConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIConfig) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaintainer

`func (o *AIConfig) GetMaintainer() AIConfigMaintainer`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *AIConfig) GetMaintainerOk() (*AIConfigMaintainer, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *AIConfig) SetMaintainer(v AIConfigMaintainer)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *AIConfig) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetMode

`func (o *AIConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AIConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AIConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AIConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *AIConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfig) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *AIConfig) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AIConfig) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AIConfig) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *AIConfig) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AIConfig) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AIConfig) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVariations

`func (o *AIConfig) GetVariations() []AIConfigVariation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *AIConfig) GetVariationsOk() (*[]AIConfigVariation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *AIConfig) SetVariations(v []AIConfigVariation)`

SetVariations sets Variations field to given value.


### GetCreatedAt

`func (o *AIConfig) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIConfig) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIConfig) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AIConfig) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AIConfig) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AIConfig) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEvaluationMetricKeys

`func (o *AIConfig) GetEvaluationMetricKeys() []string`

GetEvaluationMetricKeys returns the EvaluationMetricKeys field if non-nil, zero value otherwise.

### GetEvaluationMetricKeysOk

`func (o *AIConfig) GetEvaluationMetricKeysOk() (*[]string, bool)`

GetEvaluationMetricKeysOk returns a tuple with the EvaluationMetricKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMetricKeys

`func (o *AIConfig) SetEvaluationMetricKeys(v []string)`

SetEvaluationMetricKeys sets EvaluationMetricKeys field to given value.

### HasEvaluationMetricKeys

`func (o *AIConfig) HasEvaluationMetricKeys() bool`

HasEvaluationMetricKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


