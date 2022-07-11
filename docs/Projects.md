# Projects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | A link to this resource. | 
**Items** | [**[]Project**](Project.md) | List of projects. | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewProjects

`func NewProjects(links map[string]Link, items []Project, ) *Projects`

NewProjects instantiates a new Projects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsWithDefaults

`func NewProjectsWithDefaults() *Projects`

NewProjectsWithDefaults instantiates a new Projects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Projects) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Projects) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Projects) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *Projects) GetItems() []Project`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Projects) GetItemsOk() (*[]Project, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Projects) SetItems(v []Project)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *Projects) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Projects) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Projects) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Projects) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


