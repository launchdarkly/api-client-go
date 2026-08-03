# AgentSkillReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**ResourceKey** | **string** | The key of the agent skill. | 
**ResourceType** | **string** | The type of the resource being referenced. | 
**Items** | [**[]AgentSkillReference**](AgentSkillReference.md) |  | 
**TotalCount** | **int32** | The total number of references. | 

## Methods

### NewAgentSkillReferences

`func NewAgentSkillReferences(resourceKey string, resourceType string, items []AgentSkillReference, totalCount int32, ) *AgentSkillReferences`

NewAgentSkillReferences instantiates a new AgentSkillReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentSkillReferencesWithDefaults

`func NewAgentSkillReferencesWithDefaults() *AgentSkillReferences`

NewAgentSkillReferencesWithDefaults instantiates a new AgentSkillReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AgentSkillReferences) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentSkillReferences) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentSkillReferences) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentSkillReferences) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceKey

`func (o *AgentSkillReferences) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *AgentSkillReferences) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *AgentSkillReferences) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *AgentSkillReferences) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AgentSkillReferences) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AgentSkillReferences) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetItems

`func (o *AgentSkillReferences) GetItems() []AgentSkillReference`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentSkillReferences) GetItemsOk() (*[]AgentSkillReference, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentSkillReferences) SetItems(v []AgentSkillReference)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AgentSkillReferences) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AgentSkillReferences) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AgentSkillReferences) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


