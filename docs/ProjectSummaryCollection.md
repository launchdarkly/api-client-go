# ProjectSummaryCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Items** | [**[]ProjectSummary**](ProjectSummary.md) |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewProjectSummaryCollection

`func NewProjectSummaryCollection(totalCount int32, items []ProjectSummary, ) *ProjectSummaryCollection`

NewProjectSummaryCollection instantiates a new ProjectSummaryCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSummaryCollectionWithDefaults

`func NewProjectSummaryCollectionWithDefaults() *ProjectSummaryCollection`

NewProjectSummaryCollectionWithDefaults instantiates a new ProjectSummaryCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ProjectSummaryCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectSummaryCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectSummaryCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *ProjectSummaryCollection) GetItems() []ProjectSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProjectSummaryCollection) GetItemsOk() (*[]ProjectSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProjectSummaryCollection) SetItems(v []ProjectSummary)`

SetItems sets Items field to given value.


### GetLinks

`func (o *ProjectSummaryCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectSummaryCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectSummaryCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectSummaryCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


