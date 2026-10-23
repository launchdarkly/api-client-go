# PipelineReleaseAudienceConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceId** | Pointer to **string** | UUID of the audience. | [optional] 
**ReleaseGuardianConfiguration** | Pointer to [**PipelineReleaseGuardianConfigurationInput**](PipelineReleaseGuardianConfigurationInput.md) |  | [optional] 
**NotifyMemberIds** | Pointer to **[]string** | An array of member IDs. These members are notified to review the approval request. | [optional] 
**NotifyTeamKeys** | Pointer to **[]string** | An array of team keys. The members of these teams are notified to review the approval request. | [optional] 

## Methods

### NewPipelineReleaseAudienceConfigInput

`func NewPipelineReleaseAudienceConfigInput() *PipelineReleaseAudienceConfigInput`

NewPipelineReleaseAudienceConfigInput instantiates a new PipelineReleaseAudienceConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReleaseAudienceConfigInputWithDefaults

`func NewPipelineReleaseAudienceConfigInputWithDefaults() *PipelineReleaseAudienceConfigInput`

NewPipelineReleaseAudienceConfigInputWithDefaults instantiates a new PipelineReleaseAudienceConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceId

`func (o *PipelineReleaseAudienceConfigInput) GetAudienceId() string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *PipelineReleaseAudienceConfigInput) GetAudienceIdOk() (*string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *PipelineReleaseAudienceConfigInput) SetAudienceId(v string)`

SetAudienceId sets AudienceId field to given value.

### HasAudienceId

`func (o *PipelineReleaseAudienceConfigInput) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### GetReleaseGuardianConfiguration

`func (o *PipelineReleaseAudienceConfigInput) GetReleaseGuardianConfiguration() PipelineReleaseGuardianConfigurationInput`

GetReleaseGuardianConfiguration returns the ReleaseGuardianConfiguration field if non-nil, zero value otherwise.

### GetReleaseGuardianConfigurationOk

`func (o *PipelineReleaseAudienceConfigInput) GetReleaseGuardianConfigurationOk() (*PipelineReleaseGuardianConfigurationInput, bool)`

GetReleaseGuardianConfigurationOk returns a tuple with the ReleaseGuardianConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGuardianConfiguration

`func (o *PipelineReleaseAudienceConfigInput) SetReleaseGuardianConfiguration(v PipelineReleaseGuardianConfigurationInput)`

SetReleaseGuardianConfiguration sets ReleaseGuardianConfiguration field to given value.

### HasReleaseGuardianConfiguration

`func (o *PipelineReleaseAudienceConfigInput) HasReleaseGuardianConfiguration() bool`

HasReleaseGuardianConfiguration returns a boolean if a field has been set.

### GetNotifyMemberIds

`func (o *PipelineReleaseAudienceConfigInput) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *PipelineReleaseAudienceConfigInput) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *PipelineReleaseAudienceConfigInput) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *PipelineReleaseAudienceConfigInput) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetNotifyTeamKeys

`func (o *PipelineReleaseAudienceConfigInput) GetNotifyTeamKeys() []string`

GetNotifyTeamKeys returns the NotifyTeamKeys field if non-nil, zero value otherwise.

### GetNotifyTeamKeysOk

`func (o *PipelineReleaseAudienceConfigInput) GetNotifyTeamKeysOk() (*[]string, bool)`

GetNotifyTeamKeysOk returns a tuple with the NotifyTeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTeamKeys

`func (o *PipelineReleaseAudienceConfigInput) SetNotifyTeamKeys(v []string)`

SetNotifyTeamKeys sets NotifyTeamKeys field to given value.

### HasNotifyTeamKeys

`func (o *PipelineReleaseAudienceConfigInput) HasNotifyTeamKeys() bool`

HasNotifyTeamKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


