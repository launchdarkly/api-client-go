# ReleasePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The phase ID | 
**Name** | **string** | The release phase name | 
**Complete** | **bool** | Whether this phase is complete | 
**CreationDate** | **int64** |  | 
**CompletionDate** | Pointer to **int64** |  | [optional] 
**CompletedBy** | Pointer to [**CompletedBy**](CompletedBy.md) |  | [optional] 
**Audiences** | [**[]ReleaseAudience**](ReleaseAudience.md) | A logical grouping of one or more environments that share attributes for rolling out changes | 
**Status** | Pointer to **string** |  | [optional] 
**Started** | Pointer to **bool** | Whether or not this phase has started | [optional] 
**StartedDate** | Pointer to **int64** |  | [optional] 
**Configuration** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewReleasePhase

`func NewReleasePhase(id string, name string, complete bool, creationDate int64, audiences []ReleaseAudience, ) *ReleasePhase`

NewReleasePhase instantiates a new ReleasePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePhaseWithDefaults

`func NewReleasePhaseWithDefaults() *ReleasePhase`

NewReleasePhaseWithDefaults instantiates a new ReleasePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleasePhase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleasePhase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleasePhase) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ReleasePhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleasePhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleasePhase) SetName(v string)`

SetName sets Name field to given value.


### GetComplete

`func (o *ReleasePhase) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ReleasePhase) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ReleasePhase) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetCreationDate

`func (o *ReleasePhase) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ReleasePhase) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ReleasePhase) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetCompletionDate

`func (o *ReleasePhase) GetCompletionDate() int64`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *ReleasePhase) GetCompletionDateOk() (*int64, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *ReleasePhase) SetCompletionDate(v int64)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *ReleasePhase) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetCompletedBy

`func (o *ReleasePhase) GetCompletedBy() CompletedBy`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *ReleasePhase) GetCompletedByOk() (*CompletedBy, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *ReleasePhase) SetCompletedBy(v CompletedBy)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *ReleasePhase) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### GetAudiences

`func (o *ReleasePhase) GetAudiences() []ReleaseAudience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *ReleasePhase) GetAudiencesOk() (*[]ReleaseAudience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *ReleasePhase) SetAudiences(v []ReleaseAudience)`

SetAudiences sets Audiences field to given value.


### GetStatus

`func (o *ReleasePhase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReleasePhase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReleasePhase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReleasePhase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStarted

`func (o *ReleasePhase) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ReleasePhase) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ReleasePhase) SetStarted(v bool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *ReleasePhase) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStartedDate

`func (o *ReleasePhase) GetStartedDate() int64`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *ReleasePhase) GetStartedDateOk() (*int64, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *ReleasePhase) SetStartedDate(v int64)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *ReleasePhase) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### GetConfiguration

`func (o *ReleasePhase) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ReleasePhase) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ReleasePhase) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ReleasePhase) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


