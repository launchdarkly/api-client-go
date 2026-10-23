# DataExportProjectRankRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectKey** | **string** | The canonical project key. | 
**ProjectName** | **string** | The current project name. | 
**Rank** | **int32** | The one-based rank within all matching projects. | 
**Usage** | **int64** | Successful Streaming transfers in the requested range. | 
**Components** | [**DataExportProjectRankComponentsRep**](DataExportProjectRankComponentsRep.md) |  | 

## Methods

### NewDataExportProjectRankRep

`func NewDataExportProjectRankRep(projectKey string, projectName string, rank int32, usage int64, components DataExportProjectRankComponentsRep, ) *DataExportProjectRankRep`

NewDataExportProjectRankRep instantiates a new DataExportProjectRankRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExportProjectRankRepWithDefaults

`func NewDataExportProjectRankRepWithDefaults() *DataExportProjectRankRep`

NewDataExportProjectRankRepWithDefaults instantiates a new DataExportProjectRankRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectKey

`func (o *DataExportProjectRankRep) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *DataExportProjectRankRep) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *DataExportProjectRankRep) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetProjectName

`func (o *DataExportProjectRankRep) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DataExportProjectRankRep) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DataExportProjectRankRep) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRank

`func (o *DataExportProjectRankRep) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *DataExportProjectRankRep) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *DataExportProjectRankRep) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetUsage

`func (o *DataExportProjectRankRep) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DataExportProjectRankRep) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DataExportProjectRankRep) SetUsage(v int64)`

SetUsage sets Usage field to given value.


### GetComponents

`func (o *DataExportProjectRankRep) GetComponents() DataExportProjectRankComponentsRep`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DataExportProjectRankRep) GetComponentsOk() (*DataExportProjectRankComponentsRep, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DataExportProjectRankRep) SetComponents(v DataExportProjectRankComponentsRep)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


