# InsightsRepositoryCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | Total number of repositories | 
**Items** | [**[]InsightsRepository**](InsightsRepository.md) | List of repositories | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewInsightsRepositoryCollection

`func NewInsightsRepositoryCollection(totalCount int32, items []InsightsRepository, ) *InsightsRepositoryCollection`

NewInsightsRepositoryCollection instantiates a new InsightsRepositoryCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsRepositoryCollectionWithDefaults

`func NewInsightsRepositoryCollectionWithDefaults() *InsightsRepositoryCollection`

NewInsightsRepositoryCollectionWithDefaults instantiates a new InsightsRepositoryCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *InsightsRepositoryCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *InsightsRepositoryCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *InsightsRepositoryCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *InsightsRepositoryCollection) GetItems() []InsightsRepository`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InsightsRepositoryCollection) GetItemsOk() (*[]InsightsRepository, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InsightsRepositoryCollection) SetItems(v []InsightsRepository)`

SetItems sets Items field to given value.


### GetLinks

`func (o *InsightsRepositoryCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InsightsRepositoryCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InsightsRepositoryCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InsightsRepositoryCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


