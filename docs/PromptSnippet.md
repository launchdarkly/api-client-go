# PromptSnippet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Maintainer** | Pointer to [**AIConfigMaintainer**](AIConfigMaintainer.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Text** | **string** | The text content of the prompt snippet | 
**Tags** | **[]string** |  | 
**Version** | **int32** |  | 
**CreatedAt** | **int64** |  | 

## Methods

### NewPromptSnippet

`func NewPromptSnippet(key string, name string, text string, tags []string, version int32, createdAt int64, ) *PromptSnippet`

NewPromptSnippet instantiates a new PromptSnippet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptSnippetWithDefaults

`func NewPromptSnippetWithDefaults() *PromptSnippet`

NewPromptSnippetWithDefaults instantiates a new PromptSnippet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PromptSnippet) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PromptSnippet) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PromptSnippet) SetKey(v string)`

SetKey sets Key field to given value.


### GetAccess

`func (o *PromptSnippet) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *PromptSnippet) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *PromptSnippet) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *PromptSnippet) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *PromptSnippet) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PromptSnippet) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PromptSnippet) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PromptSnippet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaintainer

`func (o *PromptSnippet) GetMaintainer() AIConfigMaintainer`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *PromptSnippet) GetMaintainerOk() (*AIConfigMaintainer, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *PromptSnippet) SetMaintainer(v AIConfigMaintainer)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *PromptSnippet) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetName

`func (o *PromptSnippet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptSnippet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptSnippet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PromptSnippet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PromptSnippet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PromptSnippet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PromptSnippet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetText

`func (o *PromptSnippet) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PromptSnippet) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PromptSnippet) SetText(v string)`

SetText sets Text field to given value.


### GetTags

`func (o *PromptSnippet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptSnippet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptSnippet) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *PromptSnippet) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptSnippet) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptSnippet) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *PromptSnippet) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptSnippet) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptSnippet) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


