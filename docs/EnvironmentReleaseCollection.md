# EnvironmentReleaseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EnvironmentRelease**](EnvironmentRelease.md) | The releases on this page. | 
**TotalCount** | **int32** | The total number of releases that match the request, before &#x60;limit&#x60; and &#x60;offset&#x60; are applied. | 
**Links** | Pointer to [**map[string]EnvironmentReleaseLink**](EnvironmentReleaseLink.md) | Links to this page (&#x60;self&#x60;), the flag (&#x60;parent&#x60;), and, where they apply, the &#x60;first&#x60;, &#x60;prev&#x60;, &#x60;next&#x60;, and &#x60;last&#x60; pages. | [optional] 

## Methods

### NewEnvironmentReleaseCollection

`func NewEnvironmentReleaseCollection(items []EnvironmentRelease, totalCount int32, ) *EnvironmentReleaseCollection`

NewEnvironmentReleaseCollection instantiates a new EnvironmentReleaseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleaseCollectionWithDefaults

`func NewEnvironmentReleaseCollectionWithDefaults() *EnvironmentReleaseCollection`

NewEnvironmentReleaseCollectionWithDefaults instantiates a new EnvironmentReleaseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EnvironmentReleaseCollection) GetItems() []EnvironmentRelease`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentReleaseCollection) GetItemsOk() (*[]EnvironmentRelease, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentReleaseCollection) SetItems(v []EnvironmentRelease)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *EnvironmentReleaseCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EnvironmentReleaseCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EnvironmentReleaseCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *EnvironmentReleaseCollection) GetLinks() map[string]EnvironmentReleaseLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentReleaseCollection) GetLinksOk() (*map[string]EnvironmentReleaseLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentReleaseCollection) SetLinks(v map[string]EnvironmentReleaseLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentReleaseCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


