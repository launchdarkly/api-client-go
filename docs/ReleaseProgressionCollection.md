# ReleaseProgressionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCount** | **int32** | The number of active releases | 
**CompletedCount** | **int32** | The number of completed releases | 
**Items** | [**[]ReleaseProgression**](ReleaseProgression.md) | A list of details for each release, across all flags, for this release pipeline | 
**Phases** | [**[]PhaseInfo**](PhaseInfo.md) | A list of details for each phase, across all releases, for this release pipeline | 
**TotalCount** | **int32** | The total number of releases for this release pipeline | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewReleaseProgressionCollection

`func NewReleaseProgressionCollection(activeCount int32, completedCount int32, items []ReleaseProgression, phases []PhaseInfo, totalCount int32, links map[string]Link, ) *ReleaseProgressionCollection`

NewReleaseProgressionCollection instantiates a new ReleaseProgressionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseProgressionCollectionWithDefaults

`func NewReleaseProgressionCollectionWithDefaults() *ReleaseProgressionCollection`

NewReleaseProgressionCollectionWithDefaults instantiates a new ReleaseProgressionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCount

`func (o *ReleaseProgressionCollection) GetActiveCount() int32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *ReleaseProgressionCollection) GetActiveCountOk() (*int32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *ReleaseProgressionCollection) SetActiveCount(v int32)`

SetActiveCount sets ActiveCount field to given value.


### GetCompletedCount

`func (o *ReleaseProgressionCollection) GetCompletedCount() int32`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *ReleaseProgressionCollection) GetCompletedCountOk() (*int32, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *ReleaseProgressionCollection) SetCompletedCount(v int32)`

SetCompletedCount sets CompletedCount field to given value.


### GetItems

`func (o *ReleaseProgressionCollection) GetItems() []ReleaseProgression`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ReleaseProgressionCollection) GetItemsOk() (*[]ReleaseProgression, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ReleaseProgressionCollection) SetItems(v []ReleaseProgression)`

SetItems sets Items field to given value.


### GetPhases

`func (o *ReleaseProgressionCollection) GetPhases() []PhaseInfo`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *ReleaseProgressionCollection) GetPhasesOk() (*[]PhaseInfo, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *ReleaseProgressionCollection) SetPhases(v []PhaseInfo)`

SetPhases sets Phases field to given value.


### GetTotalCount

`func (o *ReleaseProgressionCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReleaseProgressionCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReleaseProgressionCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *ReleaseProgressionCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReleaseProgressionCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReleaseProgressionCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


