# StatisticsRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]Link**](Link.md) | The location and content type of all projects that have code references | [optional] 
**Self** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewStatisticsRoot

`func NewStatisticsRoot() *StatisticsRoot`

NewStatisticsRoot instantiates a new StatisticsRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsRootWithDefaults

`func NewStatisticsRootWithDefaults() *StatisticsRoot`

NewStatisticsRootWithDefaults instantiates a new StatisticsRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *StatisticsRoot) GetProjects() []Link`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *StatisticsRoot) GetProjectsOk() (*[]Link, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *StatisticsRoot) SetProjects(v []Link)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *StatisticsRoot) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSelf

`func (o *StatisticsRoot) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *StatisticsRoot) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *StatisticsRoot) SetSelf(v Link)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *StatisticsRoot) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


