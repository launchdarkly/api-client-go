# AgentSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Maintainer** | Pointer to [**AIConfigMaintainer**](AIConfigMaintainer.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Markdown** | **string** | The verbatim SKILL.md content of the agent skill | 
**Tags** | **[]string** |  | 
**Version** | **int32** |  | 
**CreatedAt** | **int64** |  | 

## Methods

### NewAgentSkill

`func NewAgentSkill(key string, name string, markdown string, tags []string, version int32, createdAt int64, ) *AgentSkill`

NewAgentSkill instantiates a new AgentSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSkillWithDefaults

`func NewAgentSkillWithDefaults() *AgentSkill`

NewAgentSkillWithDefaults instantiates a new AgentSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentSkill) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentSkill) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentSkill) SetKey(v string)`

SetKey sets Key field to given value.


### GetAccess

`func (o *AgentSkill) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AgentSkill) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AgentSkill) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AgentSkill) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *AgentSkill) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentSkill) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentSkill) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentSkill) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaintainer

`func (o *AgentSkill) GetMaintainer() AIConfigMaintainer`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *AgentSkill) GetMaintainerOk() (*AIConfigMaintainer, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *AgentSkill) SetMaintainer(v AIConfigMaintainer)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *AgentSkill) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetName

`func (o *AgentSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentSkill) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentSkill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentSkill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkdown

`func (o *AgentSkill) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *AgentSkill) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *AgentSkill) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetTags

`func (o *AgentSkill) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentSkill) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentSkill) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *AgentSkill) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentSkill) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentSkill) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *AgentSkill) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentSkill) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentSkill) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


