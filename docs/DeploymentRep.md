# DeploymentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The deployment ID | 
**ApplicationKey** | **string** | The application key | 
**ApplicationVersion** | **string** | The application version | 
**StartedAt** | **int64** |  | 
**EndedAt** | Pointer to **int64** |  | [optional] 
**DurationMs** | Pointer to **int64** | The duration of the deployment in milliseconds | [optional] 
**Status** | **string** |  | 
**Kind** | **string** |  | 
**Active** | **bool** | Whether the deployment is active | 
**Metadata** | Pointer to **map[string]interface{}** | The metadata associated with the deployment | [optional] 
**Archived** | **bool** | Whether the deployment is archived | 
**EnvironmentKey** | **string** | The environment key | 
**NumberOfContributors** | **int32** | The number of contributors | 
**NumberOfPullRequests** | **int32** | The number of pull requests | 
**LinesAdded** | **int64** | The number of lines added | 
**LinesDeleted** | **int64** | The number of lines deleted | 
**LeadTime** | **int64** | The total lead time from first commit to deployment end in milliseconds | 
**PullRequests** | Pointer to [**PullRequestCollectionRep**](PullRequestCollectionRep.md) |  | [optional] 
**FlagReferences** | Pointer to [**FlagReferenceCollectionRep**](FlagReferenceCollectionRep.md) |  | [optional] 
**LeadTimeStages** | Pointer to [**LeadTimeStagesRep**](LeadTimeStagesRep.md) |  | [optional] 

## Methods

### NewDeploymentRep

`func NewDeploymentRep(id string, applicationKey string, applicationVersion string, startedAt int64, status string, kind string, active bool, archived bool, environmentKey string, numberOfContributors int32, numberOfPullRequests int32, linesAdded int64, linesDeleted int64, leadTime int64, ) *DeploymentRep`

NewDeploymentRep instantiates a new DeploymentRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRepWithDefaults

`func NewDeploymentRepWithDefaults() *DeploymentRep`

NewDeploymentRepWithDefaults instantiates a new DeploymentRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentRep) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationKey

`func (o *DeploymentRep) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *DeploymentRep) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *DeploymentRep) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.


### GetApplicationVersion

`func (o *DeploymentRep) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *DeploymentRep) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *DeploymentRep) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.


### GetStartedAt

`func (o *DeploymentRep) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DeploymentRep) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DeploymentRep) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *DeploymentRep) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *DeploymentRep) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *DeploymentRep) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *DeploymentRep) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetDurationMs

`func (o *DeploymentRep) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *DeploymentRep) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *DeploymentRep) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *DeploymentRep) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetKind

`func (o *DeploymentRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeploymentRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeploymentRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetActive

`func (o *DeploymentRep) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DeploymentRep) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DeploymentRep) SetActive(v bool)`

SetActive sets Active field to given value.


### GetMetadata

`func (o *DeploymentRep) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeploymentRep) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeploymentRep) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeploymentRep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetArchived

`func (o *DeploymentRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *DeploymentRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *DeploymentRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetEnvironmentKey

`func (o *DeploymentRep) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *DeploymentRep) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *DeploymentRep) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetNumberOfContributors

`func (o *DeploymentRep) GetNumberOfContributors() int32`

GetNumberOfContributors returns the NumberOfContributors field if non-nil, zero value otherwise.

### GetNumberOfContributorsOk

`func (o *DeploymentRep) GetNumberOfContributorsOk() (*int32, bool)`

GetNumberOfContributorsOk returns a tuple with the NumberOfContributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfContributors

`func (o *DeploymentRep) SetNumberOfContributors(v int32)`

SetNumberOfContributors sets NumberOfContributors field to given value.


### GetNumberOfPullRequests

`func (o *DeploymentRep) GetNumberOfPullRequests() int32`

GetNumberOfPullRequests returns the NumberOfPullRequests field if non-nil, zero value otherwise.

### GetNumberOfPullRequestsOk

`func (o *DeploymentRep) GetNumberOfPullRequestsOk() (*int32, bool)`

GetNumberOfPullRequestsOk returns a tuple with the NumberOfPullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPullRequests

`func (o *DeploymentRep) SetNumberOfPullRequests(v int32)`

SetNumberOfPullRequests sets NumberOfPullRequests field to given value.


### GetLinesAdded

`func (o *DeploymentRep) GetLinesAdded() int64`

GetLinesAdded returns the LinesAdded field if non-nil, zero value otherwise.

### GetLinesAddedOk

`func (o *DeploymentRep) GetLinesAddedOk() (*int64, bool)`

GetLinesAddedOk returns a tuple with the LinesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesAdded

`func (o *DeploymentRep) SetLinesAdded(v int64)`

SetLinesAdded sets LinesAdded field to given value.


### GetLinesDeleted

`func (o *DeploymentRep) GetLinesDeleted() int64`

GetLinesDeleted returns the LinesDeleted field if non-nil, zero value otherwise.

### GetLinesDeletedOk

`func (o *DeploymentRep) GetLinesDeletedOk() (*int64, bool)`

GetLinesDeletedOk returns a tuple with the LinesDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesDeleted

`func (o *DeploymentRep) SetLinesDeleted(v int64)`

SetLinesDeleted sets LinesDeleted field to given value.


### GetLeadTime

`func (o *DeploymentRep) GetLeadTime() int64`

GetLeadTime returns the LeadTime field if non-nil, zero value otherwise.

### GetLeadTimeOk

`func (o *DeploymentRep) GetLeadTimeOk() (*int64, bool)`

GetLeadTimeOk returns a tuple with the LeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTime

`func (o *DeploymentRep) SetLeadTime(v int64)`

SetLeadTime sets LeadTime field to given value.


### GetPullRequests

`func (o *DeploymentRep) GetPullRequests() PullRequestCollectionRep`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *DeploymentRep) GetPullRequestsOk() (*PullRequestCollectionRep, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *DeploymentRep) SetPullRequests(v PullRequestCollectionRep)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *DeploymentRep) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### GetFlagReferences

`func (o *DeploymentRep) GetFlagReferences() FlagReferenceCollectionRep`

GetFlagReferences returns the FlagReferences field if non-nil, zero value otherwise.

### GetFlagReferencesOk

`func (o *DeploymentRep) GetFlagReferencesOk() (*FlagReferenceCollectionRep, bool)`

GetFlagReferencesOk returns a tuple with the FlagReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagReferences

`func (o *DeploymentRep) SetFlagReferences(v FlagReferenceCollectionRep)`

SetFlagReferences sets FlagReferences field to given value.

### HasFlagReferences

`func (o *DeploymentRep) HasFlagReferences() bool`

HasFlagReferences returns a boolean if a field has been set.

### GetLeadTimeStages

`func (o *DeploymentRep) GetLeadTimeStages() LeadTimeStagesRep`

GetLeadTimeStages returns the LeadTimeStages field if non-nil, zero value otherwise.

### GetLeadTimeStagesOk

`func (o *DeploymentRep) GetLeadTimeStagesOk() (*LeadTimeStagesRep, bool)`

GetLeadTimeStagesOk returns a tuple with the LeadTimeStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTimeStages

`func (o *DeploymentRep) SetLeadTimeStages(v LeadTimeStagesRep)`

SetLeadTimeStages sets LeadTimeStages field to given value.

### HasLeadTimeStages

`func (o *DeploymentRep) HasLeadTimeStages() bool`

HasLeadTimeStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


