# AgentGraphPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key for the agent graph | 
**Name** | **string** | A human-readable name for the agent graph | 
**Description** | Pointer to **string** | A description of the agent graph | [optional] 
**RootConfigKey** | Pointer to **string** | The AI Config key of the root node. A missing root implies a newly created graph with metadata only. | [optional] 
**Edges** | Pointer to [**[]AgentGraphEdgePost**](AgentGraphEdgePost.md) | The edges in the graph. If edges or rootConfigKey is present, both must be present. | [optional] 

## Methods

### NewAgentGraphPost

`func NewAgentGraphPost(key string, name string, ) *AgentGraphPost`

NewAgentGraphPost instantiates a new AgentGraphPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGraphPostWithDefaults

`func NewAgentGraphPostWithDefaults() *AgentGraphPost`

NewAgentGraphPostWithDefaults instantiates a new AgentGraphPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentGraphPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentGraphPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentGraphPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AgentGraphPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentGraphPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentGraphPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentGraphPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentGraphPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentGraphPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentGraphPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRootConfigKey

`func (o *AgentGraphPost) GetRootConfigKey() string`

GetRootConfigKey returns the RootConfigKey field if non-nil, zero value otherwise.

### GetRootConfigKeyOk

`func (o *AgentGraphPost) GetRootConfigKeyOk() (*string, bool)`

GetRootConfigKeyOk returns a tuple with the RootConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootConfigKey

`func (o *AgentGraphPost) SetRootConfigKey(v string)`

SetRootConfigKey sets RootConfigKey field to given value.

### HasRootConfigKey

`func (o *AgentGraphPost) HasRootConfigKey() bool`

HasRootConfigKey returns a boolean if a field has been set.

### GetEdges

`func (o *AgentGraphPost) GetEdges() []AgentGraphEdgePost`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *AgentGraphPost) GetEdgesOk() (*[]AgentGraphEdgePost, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *AgentGraphPost) SetEdges(v []AgentGraphEdgePost)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *AgentGraphPost) HasEdges() bool`

HasEdges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


