# AgentOptimizationRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**Items** | [**[]AgentOptimizationRun**](AgentOptimizationRun.md) |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAgentOptimizationRuns

`func NewAgentOptimizationRuns(items []AgentOptimizationRun, ) *AgentOptimizationRuns`

NewAgentOptimizationRuns instantiates a new AgentOptimizationRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationRunsWithDefaults

`func NewAgentOptimizationRunsWithDefaults() *AgentOptimizationRuns`

NewAgentOptimizationRunsWithDefaults instantiates a new AgentOptimizationRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AgentOptimizationRuns) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentOptimizationRuns) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentOptimizationRuns) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentOptimizationRuns) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *AgentOptimizationRuns) GetItems() []AgentOptimizationRun`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentOptimizationRuns) GetItemsOk() (*[]AgentOptimizationRun, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentOptimizationRuns) SetItems(v []AgentOptimizationRun)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AgentOptimizationRuns) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AgentOptimizationRuns) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AgentOptimizationRuns) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AgentOptimizationRuns) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


