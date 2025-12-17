# AgentGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key for the agent graph | 
**Name** | **string** | A human-readable name for the agent graph | 
**Description** | Pointer to **string** | A description of the agent graph | [optional] 
**RootConfigKey** | Pointer to **string** | The AI Config key of the root node | [optional] 
**Edges** | Pointer to [**[]AgentGraphEdge**](AgentGraphEdge.md) | The edges in the graph | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewAgentGraph

`func NewAgentGraph(key string, name string, createdAt int64, updatedAt int64, ) *AgentGraph`

NewAgentGraph instantiates a new AgentGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGraphWithDefaults

`func NewAgentGraphWithDefaults() *AgentGraph`

NewAgentGraphWithDefaults instantiates a new AgentGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentGraph) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentGraph) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentGraph) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AgentGraph) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentGraph) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentGraph) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentGraph) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentGraph) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentGraph) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentGraph) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRootConfigKey

`func (o *AgentGraph) GetRootConfigKey() string`

GetRootConfigKey returns the RootConfigKey field if non-nil, zero value otherwise.

### GetRootConfigKeyOk

`func (o *AgentGraph) GetRootConfigKeyOk() (*string, bool)`

GetRootConfigKeyOk returns a tuple with the RootConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootConfigKey

`func (o *AgentGraph) SetRootConfigKey(v string)`

SetRootConfigKey sets RootConfigKey field to given value.

### HasRootConfigKey

`func (o *AgentGraph) HasRootConfigKey() bool`

HasRootConfigKey returns a boolean if a field has been set.

### GetEdges

`func (o *AgentGraph) GetEdges() []AgentGraphEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *AgentGraph) GetEdgesOk() (*[]AgentGraphEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *AgentGraph) SetEdges(v []AgentGraphEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *AgentGraph) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentGraph) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentGraph) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentGraph) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgentGraph) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentGraph) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentGraph) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


