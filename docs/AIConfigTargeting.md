# AIConfigTargeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**Defaults** | Pointer to [**AIConfigTargetingDefaults**](AIConfigTargetingDefaults.md) |  | [optional] 
**Environments** | [**map[string]AIConfigTargetingEnvironment**](AIConfigTargetingEnvironment.md) |  | 
**Experiments** | [**AiConfigsExperimentInfoRep**](AiConfigsExperimentInfoRep.md) |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Tags** | **[]string** |  | 
**Variations** | [**[]AIConfigTargetingVariation**](AIConfigTargetingVariation.md) |  | 
**Version** | **int32** |  | 

## Methods

### NewAIConfigTargeting

`func NewAIConfigTargeting(createdAt int64, environments map[string]AIConfigTargetingEnvironment, experiments AiConfigsExperimentInfoRep, key string, name string, tags []string, variations []AIConfigTargetingVariation, version int32, ) *AIConfigTargeting`

NewAIConfigTargeting instantiates a new AIConfigTargeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigTargetingWithDefaults

`func NewAIConfigTargetingWithDefaults() *AIConfigTargeting`

NewAIConfigTargetingWithDefaults instantiates a new AIConfigTargeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AIConfigTargeting) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIConfigTargeting) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIConfigTargeting) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaults

`func (o *AIConfigTargeting) GetDefaults() AIConfigTargetingDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *AIConfigTargeting) GetDefaultsOk() (*AIConfigTargetingDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *AIConfigTargeting) SetDefaults(v AIConfigTargetingDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *AIConfigTargeting) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetEnvironments

`func (o *AIConfigTargeting) GetEnvironments() map[string]AIConfigTargetingEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AIConfigTargeting) GetEnvironmentsOk() (*map[string]AIConfigTargetingEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AIConfigTargeting) SetEnvironments(v map[string]AIConfigTargetingEnvironment)`

SetEnvironments sets Environments field to given value.


### GetExperiments

`func (o *AIConfigTargeting) GetExperiments() AiConfigsExperimentInfoRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *AIConfigTargeting) GetExperimentsOk() (*AiConfigsExperimentInfoRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *AIConfigTargeting) SetExperiments(v AiConfigsExperimentInfoRep)`

SetExperiments sets Experiments field to given value.


### GetKey

`func (o *AIConfigTargeting) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIConfigTargeting) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIConfigTargeting) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AIConfigTargeting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfigTargeting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfigTargeting) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *AIConfigTargeting) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AIConfigTargeting) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AIConfigTargeting) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVariations

`func (o *AIConfigTargeting) GetVariations() []AIConfigTargetingVariation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *AIConfigTargeting) GetVariationsOk() (*[]AIConfigTargetingVariation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *AIConfigTargeting) SetVariations(v []AIConfigTargetingVariation)`

SetVariations sets Variations field to given value.


### GetVersion

`func (o *AIConfigTargeting) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AIConfigTargeting) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AIConfigTargeting) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


