# AgentGraphPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name for the agent graph | [optional] 
**Description** | Pointer to **string** | A description of the agent graph | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this agent graph. Pass an empty string to remove maintainer. | [optional] 
**MaintainerTeamKey** | Pointer to **string** | The key of the team that maintains this agent graph. Pass an empty string to remove maintainer. | [optional] 
**RootConfigKey** | Pointer to **string** | The AI Config key of the root node. If present, edges must also be present. | [optional] 
**Edges** | Pointer to [**[]AgentGraphEdge**](AgentGraphEdge.md) | The edges in the graph. If present, rootConfigKey must also be present. Replaces all existing edges. | [optional] 

## Methods

### NewAgentGraphPatch

`func NewAgentGraphPatch() *AgentGraphPatch`

NewAgentGraphPatch instantiates a new AgentGraphPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGraphPatchWithDefaults

`func NewAgentGraphPatchWithDefaults() *AgentGraphPatch`

NewAgentGraphPatchWithDefaults instantiates a new AgentGraphPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentGraphPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentGraphPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentGraphPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentGraphPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AgentGraphPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentGraphPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentGraphPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentGraphPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainerId

`func (o *AgentGraphPatch) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AgentGraphPatch) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AgentGraphPatch) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AgentGraphPatch) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *AgentGraphPatch) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *AgentGraphPatch) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *AgentGraphPatch) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *AgentGraphPatch) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetRootConfigKey

`func (o *AgentGraphPatch) GetRootConfigKey() string`

GetRootConfigKey returns the RootConfigKey field if non-nil, zero value otherwise.

### GetRootConfigKeyOk

`func (o *AgentGraphPatch) GetRootConfigKeyOk() (*string, bool)`

GetRootConfigKeyOk returns a tuple with the RootConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootConfigKey

`func (o *AgentGraphPatch) SetRootConfigKey(v string)`

SetRootConfigKey sets RootConfigKey field to given value.

### HasRootConfigKey

`func (o *AgentGraphPatch) HasRootConfigKey() bool`

HasRootConfigKey returns a boolean if a field has been set.

### GetEdges

`func (o *AgentGraphPatch) GetEdges() []AgentGraphEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *AgentGraphPatch) GetEdgesOk() (*[]AgentGraphEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *AgentGraphPatch) SetEdges(v []AgentGraphEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *AgentGraphPatch) HasEdges() bool`

HasEdges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


