# StatisticRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The repository name | 
**Type** | **string** | The type of repository | 
**SourceLink** | **string** | A URL to access the repository | 
**DefaultBranch** | **string** | The repository&#39;s default branch | 
**Enabled** | **bool** | Whether or not a repository is enabled for code reference scanning | 
**Version** | **int32** | The version of the repository&#39;s saved information | 
**HunkCount** | **int32** | The number of code reference hunks in which the flag appears in this repository | 
**FileCount** | **int32** | The number of files in which the flag appears in this repository | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**LatestCommitTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewStatisticRep

`func NewStatisticRep(name string, type_ string, sourceLink string, defaultBranch string, enabled bool, version int32, hunkCount int32, fileCount int32, links map[string]Link, ) *StatisticRep`

NewStatisticRep instantiates a new StatisticRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticRepWithDefaults

`func NewStatisticRepWithDefaults() *StatisticRep`

NewStatisticRepWithDefaults instantiates a new StatisticRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StatisticRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatisticRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatisticRep) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *StatisticRep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatisticRep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatisticRep) SetType(v string)`

SetType sets Type field to given value.


### GetSourceLink

`func (o *StatisticRep) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *StatisticRep) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *StatisticRep) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.


### GetDefaultBranch

`func (o *StatisticRep) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *StatisticRep) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *StatisticRep) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetEnabled

`func (o *StatisticRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StatisticRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StatisticRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVersion

`func (o *StatisticRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatisticRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatisticRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetHunkCount

`func (o *StatisticRep) GetHunkCount() int32`

GetHunkCount returns the HunkCount field if non-nil, zero value otherwise.

### GetHunkCountOk

`func (o *StatisticRep) GetHunkCountOk() (*int32, bool)`

GetHunkCountOk returns a tuple with the HunkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunkCount

`func (o *StatisticRep) SetHunkCount(v int32)`

SetHunkCount sets HunkCount field to given value.


### GetFileCount

`func (o *StatisticRep) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *StatisticRep) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *StatisticRep) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.


### GetLinks

`func (o *StatisticRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StatisticRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StatisticRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetLatestCommitTime

`func (o *StatisticRep) GetLatestCommitTime() int64`

GetLatestCommitTime returns the LatestCommitTime field if non-nil, zero value otherwise.

### GetLatestCommitTimeOk

`func (o *StatisticRep) GetLatestCommitTimeOk() (*int64, bool)`

GetLatestCommitTimeOk returns a tuple with the LatestCommitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCommitTime

`func (o *StatisticRep) SetLatestCommitTime(v int64)`

SetLatestCommitTime sets LatestCommitTime field to given value.

### HasLatestCommitTime

`func (o *StatisticRep) HasLatestCommitTime() bool`

HasLatestCommitTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


