# InsightsRepositoryProjectCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | Total number of repository project associations | 
**Items** | [**[]InsightsRepositoryProject**](InsightsRepositoryProject.md) | List of repository project associations | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewInsightsRepositoryProjectCollection

`func NewInsightsRepositoryProjectCollection(totalCount int32, items []InsightsRepositoryProject, ) *InsightsRepositoryProjectCollection`

NewInsightsRepositoryProjectCollection instantiates a new InsightsRepositoryProjectCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsRepositoryProjectCollectionWithDefaults

`func NewInsightsRepositoryProjectCollectionWithDefaults() *InsightsRepositoryProjectCollection`

NewInsightsRepositoryProjectCollectionWithDefaults instantiates a new InsightsRepositoryProjectCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *InsightsRepositoryProjectCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *InsightsRepositoryProjectCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *InsightsRepositoryProjectCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *InsightsRepositoryProjectCollection) GetItems() []InsightsRepositoryProject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InsightsRepositoryProjectCollection) GetItemsOk() (*[]InsightsRepositoryProject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InsightsRepositoryProjectCollection) SetItems(v []InsightsRepositoryProject)`

SetItems sets Items field to given value.


### GetLinks

`func (o *InsightsRepositoryProjectCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InsightsRepositoryProjectCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InsightsRepositoryProjectCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InsightsRepositoryProjectCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


