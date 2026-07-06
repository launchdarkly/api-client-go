# SnippetReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**ResourceKey** | **string** | The key of the prompt snippet. | 
**ResourceType** | **string** | The type of the resource being referenced. | 
**Items** | [**[]SnippetReference**](SnippetReference.md) |  | 
**TotalCount** | **int32** | The total number of references. | 

## Methods

### NewSnippetReferences

`func NewSnippetReferences(resourceKey string, resourceType string, items []SnippetReference, totalCount int32, ) *SnippetReferences`

NewSnippetReferences instantiates a new SnippetReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnippetReferencesWithDefaults

`func NewSnippetReferencesWithDefaults() *SnippetReferences`

NewSnippetReferencesWithDefaults instantiates a new SnippetReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SnippetReferences) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SnippetReferences) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SnippetReferences) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SnippetReferences) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResourceKey

`func (o *SnippetReferences) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *SnippetReferences) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *SnippetReferences) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *SnippetReferences) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *SnippetReferences) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *SnippetReferences) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetItems

`func (o *SnippetReferences) GetItems() []SnippetReference`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnippetReferences) GetItemsOk() (*[]SnippetReference, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnippetReferences) SetItems(v []SnippetReference)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *SnippetReferences) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SnippetReferences) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SnippetReferences) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


