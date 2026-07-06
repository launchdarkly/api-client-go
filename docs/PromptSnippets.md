# PromptSnippets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**Items** | [**[]PromptSnippet**](PromptSnippet.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPromptSnippets

`func NewPromptSnippets(items []PromptSnippet, totalCount int32, ) *PromptSnippets`

NewPromptSnippets instantiates a new PromptSnippets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptSnippetsWithDefaults

`func NewPromptSnippetsWithDefaults() *PromptSnippets`

NewPromptSnippetsWithDefaults instantiates a new PromptSnippets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PromptSnippets) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PromptSnippets) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PromptSnippets) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PromptSnippets) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *PromptSnippets) GetItems() []PromptSnippet`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PromptSnippets) GetItemsOk() (*[]PromptSnippet, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PromptSnippets) SetItems(v []PromptSnippet)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *PromptSnippets) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PromptSnippets) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PromptSnippets) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


