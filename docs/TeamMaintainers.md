# TeamMaintainers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The number of maintainers of the team | [optional] 
**Items** | Pointer to [**[]MemberSummary**](MemberSummary.md) | Details on the members that have been assigned as maintainers of the team | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewTeamMaintainers

`func NewTeamMaintainers() *TeamMaintainers`

NewTeamMaintainers instantiates a new TeamMaintainers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMaintainersWithDefaults

`func NewTeamMaintainersWithDefaults() *TeamMaintainers`

NewTeamMaintainersWithDefaults instantiates a new TeamMaintainers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *TeamMaintainers) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamMaintainers) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamMaintainers) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamMaintainers) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *TeamMaintainers) GetItems() []MemberSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamMaintainers) GetItemsOk() (*[]MemberSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamMaintainers) SetItems(v []MemberSummary)`

SetItems sets Items field to given value.

### HasItems

`func (o *TeamMaintainers) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *TeamMaintainers) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamMaintainers) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamMaintainers) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamMaintainers) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


