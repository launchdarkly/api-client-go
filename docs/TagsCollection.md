# TagsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]string** | List of tags | 
**Links** | [**map[string]TagsLink**](TagsLink.md) |  | 
**TotalCount** | Pointer to **int32** | The total number of tags | [optional] 

## Methods

### NewTagsCollection

`func NewTagsCollection(items []string, links map[string]TagsLink, ) *TagsCollection`

NewTagsCollection instantiates a new TagsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsCollectionWithDefaults

`func NewTagsCollectionWithDefaults() *TagsCollection`

NewTagsCollectionWithDefaults instantiates a new TagsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TagsCollection) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TagsCollection) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TagsCollection) SetItems(v []string)`

SetItems sets Items field to given value.


### GetLinks

`func (o *TagsCollection) GetLinks() map[string]TagsLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TagsCollection) GetLinksOk() (*map[string]TagsLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TagsCollection) SetLinks(v map[string]TagsLink)`

SetLinks sets Links field to given value.


### GetTotalCount

`func (o *TagsCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TagsCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TagsCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TagsCollection) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


