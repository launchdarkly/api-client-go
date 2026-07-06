# AgentOptimizationResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**Items** | [**[]AgentOptimizationResult**](AgentOptimizationResult.md) |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAgentOptimizationResults

`func NewAgentOptimizationResults(items []AgentOptimizationResult, ) *AgentOptimizationResults`

NewAgentOptimizationResults instantiates a new AgentOptimizationResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentOptimizationResultsWithDefaults

`func NewAgentOptimizationResultsWithDefaults() *AgentOptimizationResults`

NewAgentOptimizationResultsWithDefaults instantiates a new AgentOptimizationResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AgentOptimizationResults) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentOptimizationResults) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentOptimizationResults) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentOptimizationResults) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *AgentOptimizationResults) GetItems() []AgentOptimizationResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentOptimizationResults) GetItemsOk() (*[]AgentOptimizationResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentOptimizationResults) SetItems(v []AgentOptimizationResult)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *AgentOptimizationResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AgentOptimizationResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AgentOptimizationResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AgentOptimizationResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


