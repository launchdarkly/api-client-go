# GenerateVariationsPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **string** |  | [optional] 
**Examples** | Pointer to **[]string** |  | [optional] 
**NumVariations** | Pointer to **int32** |  | [optional] 

## Methods

### NewGenerateVariationsPost

`func NewGenerateVariationsPost() *GenerateVariationsPost`

NewGenerateVariationsPost instantiates a new GenerateVariationsPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateVariationsPostWithDefaults

`func NewGenerateVariationsPostWithDefaults() *GenerateVariationsPost`

NewGenerateVariationsPostWithDefaults instantiates a new GenerateVariationsPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *GenerateVariationsPost) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *GenerateVariationsPost) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *GenerateVariationsPost) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *GenerateVariationsPost) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetExamples

`func (o *GenerateVariationsPost) GetExamples() []string`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *GenerateVariationsPost) GetExamplesOk() (*[]string, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *GenerateVariationsPost) SetExamples(v []string)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *GenerateVariationsPost) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetNumVariations

`func (o *GenerateVariationsPost) GetNumVariations() int32`

GetNumVariations returns the NumVariations field if non-nil, zero value otherwise.

### GetNumVariationsOk

`func (o *GenerateVariationsPost) GetNumVariationsOk() (*int32, bool)`

GetNumVariationsOk returns a tuple with the NumVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVariations

`func (o *GenerateVariationsPost) SetNumVariations(v int32)`

SetNumVariations sets NumVariations field to given value.

### HasNumVariations

`func (o *GenerateVariationsPost) HasNumVariations() bool`

HasNumVariations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


