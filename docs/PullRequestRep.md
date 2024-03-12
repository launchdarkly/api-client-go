# PullRequestRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The pull request internal ID | 
**ExternalId** | **string** | The pull request number | 
**Title** | **string** | The pull request title | 
**Status** | **string** | The pull request status | 
**Author** | **string** | The pull request author | 
**CreateTime** | **int64** |  | 
**MergeTime** | Pointer to **int64** |  | [optional] 
**MergeCommitKey** | Pointer to **string** | The pull request merge commit key | [optional] 
**BaseCommitKey** | **string** | The pull request base commit key | 
**HeadCommitKey** | **string** | The pull request head commit key | 
**FilesChanged** | **int32** | The number of files changed | 
**LinesAdded** | **int32** | The number of lines added | 
**LinesDeleted** | **int32** | The number of lines deleted | 
**Url** | **string** | The pull request URL | 
**Deployments** | Pointer to [**DeploymentCollectionRep**](DeploymentCollectionRep.md) |  | [optional] 
**FlagReferences** | Pointer to [**FlagReferenceCollectionRep**](FlagReferenceCollectionRep.md) |  | [optional] 
**LeadTime** | Pointer to [**PullRequestLeadTimeRep**](PullRequestLeadTimeRep.md) |  | [optional] 

## Methods

### NewPullRequestRep

`func NewPullRequestRep(id string, externalId string, title string, status string, author string, createTime int64, baseCommitKey string, headCommitKey string, filesChanged int32, linesAdded int32, linesDeleted int32, url string, ) *PullRequestRep`

NewPullRequestRep instantiates a new PullRequestRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestRepWithDefaults

`func NewPullRequestRepWithDefaults() *PullRequestRep`

NewPullRequestRepWithDefaults instantiates a new PullRequestRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequestRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestRep) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *PullRequestRep) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PullRequestRep) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PullRequestRep) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetTitle

`func (o *PullRequestRep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequestRep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequestRep) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *PullRequestRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PullRequestRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PullRequestRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAuthor

`func (o *PullRequestRep) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PullRequestRep) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PullRequestRep) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetCreateTime

`func (o *PullRequestRep) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PullRequestRep) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PullRequestRep) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.


### GetMergeTime

`func (o *PullRequestRep) GetMergeTime() int64`

GetMergeTime returns the MergeTime field if non-nil, zero value otherwise.

### GetMergeTimeOk

`func (o *PullRequestRep) GetMergeTimeOk() (*int64, bool)`

GetMergeTimeOk returns a tuple with the MergeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeTime

`func (o *PullRequestRep) SetMergeTime(v int64)`

SetMergeTime sets MergeTime field to given value.

### HasMergeTime

`func (o *PullRequestRep) HasMergeTime() bool`

HasMergeTime returns a boolean if a field has been set.

### GetMergeCommitKey

`func (o *PullRequestRep) GetMergeCommitKey() string`

GetMergeCommitKey returns the MergeCommitKey field if non-nil, zero value otherwise.

### GetMergeCommitKeyOk

`func (o *PullRequestRep) GetMergeCommitKeyOk() (*string, bool)`

GetMergeCommitKeyOk returns a tuple with the MergeCommitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitKey

`func (o *PullRequestRep) SetMergeCommitKey(v string)`

SetMergeCommitKey sets MergeCommitKey field to given value.

### HasMergeCommitKey

`func (o *PullRequestRep) HasMergeCommitKey() bool`

HasMergeCommitKey returns a boolean if a field has been set.

### GetBaseCommitKey

`func (o *PullRequestRep) GetBaseCommitKey() string`

GetBaseCommitKey returns the BaseCommitKey field if non-nil, zero value otherwise.

### GetBaseCommitKeyOk

`func (o *PullRequestRep) GetBaseCommitKeyOk() (*string, bool)`

GetBaseCommitKeyOk returns a tuple with the BaseCommitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommitKey

`func (o *PullRequestRep) SetBaseCommitKey(v string)`

SetBaseCommitKey sets BaseCommitKey field to given value.


### GetHeadCommitKey

`func (o *PullRequestRep) GetHeadCommitKey() string`

GetHeadCommitKey returns the HeadCommitKey field if non-nil, zero value otherwise.

### GetHeadCommitKeyOk

`func (o *PullRequestRep) GetHeadCommitKeyOk() (*string, bool)`

GetHeadCommitKeyOk returns a tuple with the HeadCommitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommitKey

`func (o *PullRequestRep) SetHeadCommitKey(v string)`

SetHeadCommitKey sets HeadCommitKey field to given value.


### GetFilesChanged

`func (o *PullRequestRep) GetFilesChanged() int32`

GetFilesChanged returns the FilesChanged field if non-nil, zero value otherwise.

### GetFilesChangedOk

`func (o *PullRequestRep) GetFilesChangedOk() (*int32, bool)`

GetFilesChangedOk returns a tuple with the FilesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesChanged

`func (o *PullRequestRep) SetFilesChanged(v int32)`

SetFilesChanged sets FilesChanged field to given value.


### GetLinesAdded

`func (o *PullRequestRep) GetLinesAdded() int32`

GetLinesAdded returns the LinesAdded field if non-nil, zero value otherwise.

### GetLinesAddedOk

`func (o *PullRequestRep) GetLinesAddedOk() (*int32, bool)`

GetLinesAddedOk returns a tuple with the LinesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesAdded

`func (o *PullRequestRep) SetLinesAdded(v int32)`

SetLinesAdded sets LinesAdded field to given value.


### GetLinesDeleted

`func (o *PullRequestRep) GetLinesDeleted() int32`

GetLinesDeleted returns the LinesDeleted field if non-nil, zero value otherwise.

### GetLinesDeletedOk

`func (o *PullRequestRep) GetLinesDeletedOk() (*int32, bool)`

GetLinesDeletedOk returns a tuple with the LinesDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesDeleted

`func (o *PullRequestRep) SetLinesDeleted(v int32)`

SetLinesDeleted sets LinesDeleted field to given value.


### GetUrl

`func (o *PullRequestRep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestRep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestRep) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDeployments

`func (o *PullRequestRep) GetDeployments() DeploymentCollectionRep`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *PullRequestRep) GetDeploymentsOk() (*DeploymentCollectionRep, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *PullRequestRep) SetDeployments(v DeploymentCollectionRep)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *PullRequestRep) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetFlagReferences

`func (o *PullRequestRep) GetFlagReferences() FlagReferenceCollectionRep`

GetFlagReferences returns the FlagReferences field if non-nil, zero value otherwise.

### GetFlagReferencesOk

`func (o *PullRequestRep) GetFlagReferencesOk() (*FlagReferenceCollectionRep, bool)`

GetFlagReferencesOk returns a tuple with the FlagReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagReferences

`func (o *PullRequestRep) SetFlagReferences(v FlagReferenceCollectionRep)`

SetFlagReferences sets FlagReferences field to given value.

### HasFlagReferences

`func (o *PullRequestRep) HasFlagReferences() bool`

HasFlagReferences returns a boolean if a field has been set.

### GetLeadTime

`func (o *PullRequestRep) GetLeadTime() PullRequestLeadTimeRep`

GetLeadTime returns the LeadTime field if non-nil, zero value otherwise.

### GetLeadTimeOk

`func (o *PullRequestRep) GetLeadTimeOk() (*PullRequestLeadTimeRep, bool)`

GetLeadTimeOk returns a tuple with the LeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTime

`func (o *PullRequestRep) SetLeadTime(v PullRequestLeadTimeRep)`

SetLeadTime sets LeadTime field to given value.

### HasLeadTime

`func (o *PullRequestRep) HasLeadTime() bool`

HasLeadTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


