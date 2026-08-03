# AgentSkillPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Markdown** | **string** | The verbatim SKILL.md content of the agent skill | 
**MaintainerId** | Pointer to **string** |  | [optional] 
**MaintainerTeamKey** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAgentSkillPost

`func NewAgentSkillPost(key string, name string, markdown string, ) *AgentSkillPost`

NewAgentSkillPost instantiates a new AgentSkillPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSkillPostWithDefaults

`func NewAgentSkillPostWithDefaults() *AgentSkillPost`

NewAgentSkillPostWithDefaults instantiates a new AgentSkillPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentSkillPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentSkillPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentSkillPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AgentSkillPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentSkillPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentSkillPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentSkillPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentSkillPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentSkillPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentSkillPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkdown

`func (o *AgentSkillPost) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *AgentSkillPost) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *AgentSkillPost) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetMaintainerId

`func (o *AgentSkillPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AgentSkillPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AgentSkillPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AgentSkillPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *AgentSkillPost) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *AgentSkillPost) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *AgentSkillPost) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *AgentSkillPost) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetTags

`func (o *AgentSkillPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentSkillPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentSkillPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentSkillPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


