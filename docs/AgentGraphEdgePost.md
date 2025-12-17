# AgentGraphEdgePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key for this edge within the graph | 
**SourceConfig** | **string** | The AI Config key that is the source of this edge | 
**TargetConfig** | **string** | The AI Config key that is the target of this edge | 
**Handoff** | Pointer to **map[string]interface{}** | The handoff options from the source AI Config to the target AI Config | [optional] 

## Methods

### NewAgentGraphEdgePost

`func NewAgentGraphEdgePost(key string, sourceConfig string, targetConfig string, ) *AgentGraphEdgePost`

NewAgentGraphEdgePost instantiates a new AgentGraphEdgePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGraphEdgePostWithDefaults

`func NewAgentGraphEdgePostWithDefaults() *AgentGraphEdgePost`

NewAgentGraphEdgePostWithDefaults instantiates a new AgentGraphEdgePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentGraphEdgePost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentGraphEdgePost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentGraphEdgePost) SetKey(v string)`

SetKey sets Key field to given value.


### GetSourceConfig

`func (o *AgentGraphEdgePost) GetSourceConfig() string`

GetSourceConfig returns the SourceConfig field if non-nil, zero value otherwise.

### GetSourceConfigOk

`func (o *AgentGraphEdgePost) GetSourceConfigOk() (*string, bool)`

GetSourceConfigOk returns a tuple with the SourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfig

`func (o *AgentGraphEdgePost) SetSourceConfig(v string)`

SetSourceConfig sets SourceConfig field to given value.


### GetTargetConfig

`func (o *AgentGraphEdgePost) GetTargetConfig() string`

GetTargetConfig returns the TargetConfig field if non-nil, zero value otherwise.

### GetTargetConfigOk

`func (o *AgentGraphEdgePost) GetTargetConfigOk() (*string, bool)`

GetTargetConfigOk returns a tuple with the TargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfig

`func (o *AgentGraphEdgePost) SetTargetConfig(v string)`

SetTargetConfig sets TargetConfig field to given value.


### GetHandoff

`func (o *AgentGraphEdgePost) GetHandoff() map[string]interface{}`

GetHandoff returns the Handoff field if non-nil, zero value otherwise.

### GetHandoffOk

`func (o *AgentGraphEdgePost) GetHandoffOk() (*map[string]interface{}, bool)`

GetHandoffOk returns a tuple with the Handoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoff

`func (o *AgentGraphEdgePost) SetHandoff(v map[string]interface{})`

SetHandoff sets Handoff field to given value.

### HasHandoff

`func (o *AgentGraphEdgePost) HasHandoff() bool`

HasHandoff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


