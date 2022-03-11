# TeamProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]ProjectSummary**](ProjectSummary.md) |  | [optional] 

## Methods

### NewTeamProjects

`func NewTeamProjects() *TeamProjects`

NewTeamProjects instantiates a new TeamProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamProjectsWithDefaults

`func NewTeamProjectsWithDefaults() *TeamProjects`

NewTeamProjectsWithDefaults instantiates a new TeamProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *TeamProjects) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamProjects) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamProjects) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamProjects) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *TeamProjects) GetItems() []ProjectSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamProjects) GetItemsOk() (*[]ProjectSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamProjects) SetItems(v []ProjectSummary)`

SetItems sets Items field to given value.

### HasItems

`func (o *TeamProjects) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


