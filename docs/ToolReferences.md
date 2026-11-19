# ToolReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**ResourceKey** | **string** | The key of the AI tool. | 
**ResourceType** | **string** | The type of the resource being referenced. | 
**Items** | [**[]ToolReference**](ToolReference.md) | The resources that reference this tool | 
**TotalCount** | **int32** | The total number of references. | 
**Skills** | Pointer to [**[]ToolSkillReference**](ToolSkillReference.md) | One page of the agent skill versions that pin a version of this tool, ordered by skill key then skill version, newest first. It is paged by &#x60;limit&#x60; and &#x60;offset&#x60; like &#x60;items&#x60;. A tool cannot be deleted while any skill version pins it. | [optional] 
**SkillsTotalCount** | Pointer to **int32** | The total number of agent skill versions that pin this tool, across all pages. | [optional] 

## Methods

### NewToolReferences

`func NewToolReferences(resourceKey string, resourceType string, items []ToolReference, totalCount int32, ) *ToolReferences`

NewToolReferences instantiates a new ToolReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolReferencesWithDefaults

`func NewToolReferencesWithDefaults() *ToolReferences`

NewToolReferencesWithDefaults instantiates a new ToolReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ToolReferences) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ToolReferences) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ToolReferences) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ToolReferences) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceKey

`func (o *ToolReferences) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ToolReferences) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ToolReferences) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *ToolReferences) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ToolReferences) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ToolReferences) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetItems

`func (o *ToolReferences) GetItems() []ToolReference`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ToolReferences) GetItemsOk() (*[]ToolReference, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ToolReferences) SetItems(v []ToolReference)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ToolReferences) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ToolReferences) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ToolReferences) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetSkills

`func (o *ToolReferences) GetSkills() []ToolSkillReference`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ToolReferences) GetSkillsOk() (*[]ToolSkillReference, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ToolReferences) SetSkills(v []ToolSkillReference)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ToolReferences) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetSkillsTotalCount

`func (o *ToolReferences) GetSkillsTotalCount() int32`

GetSkillsTotalCount returns the SkillsTotalCount field if non-nil, zero value otherwise.

### GetSkillsTotalCountOk

`func (o *ToolReferences) GetSkillsTotalCountOk() (*int32, bool)`

GetSkillsTotalCountOk returns a tuple with the SkillsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillsTotalCount

`func (o *ToolReferences) SetSkillsTotalCount(v int32)`

SetSkillsTotalCount sets SkillsTotalCount field to given value.

### HasSkillsTotalCount

`func (o *ToolReferences) HasSkillsTotalCount() bool`

HasSkillsTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


