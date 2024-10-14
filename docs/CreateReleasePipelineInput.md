# CreateReleasePipelineInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The release pipeline description | [optional] 
**Key** | **string** | The unique identifier of this release pipeline | 
**Name** | **string** | The name of the release pipeline | 
**Phases** | [**[]CreatePhaseInput**](CreatePhaseInput.md) | A logical grouping of one or more environments that share attributes for rolling out changes | 
**Tags** | Pointer to **[]string** | A list of tags for this release pipeline | [optional] 
**IsProjectDefault** | Pointer to **bool** | Whether or not the newly created pipeline should be set as the default pipeline for this project | [optional] 
**IsLegacy** | Pointer to **bool** | Whether or not the pipeline is enabled for Release Automation. | [optional] 

## Methods

### NewCreateReleasePipelineInput

`func NewCreateReleasePipelineInput(key string, name string, phases []CreatePhaseInput, ) *CreateReleasePipelineInput`

NewCreateReleasePipelineInput instantiates a new CreateReleasePipelineInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleasePipelineInputWithDefaults

`func NewCreateReleasePipelineInputWithDefaults() *CreateReleasePipelineInput`

NewCreateReleasePipelineInputWithDefaults instantiates a new CreateReleasePipelineInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateReleasePipelineInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReleasePipelineInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReleasePipelineInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateReleasePipelineInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *CreateReleasePipelineInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateReleasePipelineInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateReleasePipelineInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateReleasePipelineInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReleasePipelineInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReleasePipelineInput) SetName(v string)`

SetName sets Name field to given value.


### GetPhases

`func (o *CreateReleasePipelineInput) GetPhases() []CreatePhaseInput`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *CreateReleasePipelineInput) GetPhasesOk() (*[]CreatePhaseInput, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *CreateReleasePipelineInput) SetPhases(v []CreatePhaseInput)`

SetPhases sets Phases field to given value.


### GetTags

`func (o *CreateReleasePipelineInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateReleasePipelineInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateReleasePipelineInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateReleasePipelineInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsProjectDefault

`func (o *CreateReleasePipelineInput) GetIsProjectDefault() bool`

GetIsProjectDefault returns the IsProjectDefault field if non-nil, zero value otherwise.

### GetIsProjectDefaultOk

`func (o *CreateReleasePipelineInput) GetIsProjectDefaultOk() (*bool, bool)`

GetIsProjectDefaultOk returns a tuple with the IsProjectDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectDefault

`func (o *CreateReleasePipelineInput) SetIsProjectDefault(v bool)`

SetIsProjectDefault sets IsProjectDefault field to given value.

### HasIsProjectDefault

`func (o *CreateReleasePipelineInput) HasIsProjectDefault() bool`

HasIsProjectDefault returns a boolean if a field has been set.

### GetIsLegacy

`func (o *CreateReleasePipelineInput) GetIsLegacy() bool`

GetIsLegacy returns the IsLegacy field if non-nil, zero value otherwise.

### GetIsLegacyOk

`func (o *CreateReleasePipelineInput) GetIsLegacyOk() (*bool, bool)`

GetIsLegacyOk returns a tuple with the IsLegacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegacy

`func (o *CreateReleasePipelineInput) SetIsLegacy(v bool)`

SetIsLegacy sets IsLegacy field to given value.

### HasIsLegacy

`func (o *CreateReleasePipelineInput) HasIsLegacy() bool`

HasIsLegacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


