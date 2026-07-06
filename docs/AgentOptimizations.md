# AgentOptimizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**Items** | [**[]AgentOptimization**](AgentOptimization.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAgentOptimizations

`func NewAgentOptimizations(items []AgentOptimization, totalCount int32, ) *AgentOptimizations`

NewAgentOptimizations instantiates a new AgentOptimizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationsWithDefaults

`func NewAgentOptimizationsWithDefaults() *AgentOptimizations`

NewAgentOptimizationsWithDefaults instantiates a new AgentOptimizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AgentOptimizations) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentOptimizations) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentOptimizations) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentOptimizations) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *AgentOptimizations) GetItems() []AgentOptimization`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentOptimizations) GetItemsOk() (*[]AgentOptimization, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentOptimizations) SetItems(v []AgentOptimization)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AgentOptimizations) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AgentOptimizations) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AgentOptimizations) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


