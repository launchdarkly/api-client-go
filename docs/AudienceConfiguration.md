# AudienceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseStrategy** | **string** |  | 
**RequireApproval** | **bool** | Whether or not the audience requires approval | 
**NotifyMemberIds** | Pointer to **[]string** | An array of member IDs. These members are notified to review the approval request. | [optional] 
**NotifyTeamKeys** | Pointer to **[]string** | An array of team keys. The members of these teams are notified to review the approval request. | [optional] 
**ReleaseGuardianConfiguration** | Pointer to [**ReleaseGuardianConfiguration**](ReleaseGuardianConfiguration.md) |  | [optional] 

## Methods

### NewAudienceConfiguration

`func NewAudienceConfiguration(releaseStrategy string, requireApproval bool, ) *AudienceConfiguration`

NewAudienceConfiguration instantiates a new AudienceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceConfigurationWithDefaults

`func NewAudienceConfigurationWithDefaults() *AudienceConfiguration`

NewAudienceConfigurationWithDefaults instantiates a new AudienceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseStrategy

`func (o *AudienceConfiguration) GetReleaseStrategy() string`

GetReleaseStrategy returns the ReleaseStrategy field if non-nil, zero value otherwise.

### GetReleaseStrategyOk

`func (o *AudienceConfiguration) GetReleaseStrategyOk() (*string, bool)`

GetReleaseStrategyOk returns a tuple with the ReleaseStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStrategy

`func (o *AudienceConfiguration) SetReleaseStrategy(v string)`

SetReleaseStrategy sets ReleaseStrategy field to given value.


### GetRequireApproval

`func (o *AudienceConfiguration) GetRequireApproval() bool`

GetRequireApproval returns the RequireApproval field if non-nil, zero value otherwise.

### GetRequireApprovalOk

`func (o *AudienceConfiguration) GetRequireApprovalOk() (*bool, bool)`

GetRequireApprovalOk returns a tuple with the RequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApproval

`func (o *AudienceConfiguration) SetRequireApproval(v bool)`

SetRequireApproval sets RequireApproval field to given value.


### GetNotifyMemberIds

`func (o *AudienceConfiguration) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *AudienceConfiguration) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *AudienceConfiguration) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *AudienceConfiguration) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetNotifyTeamKeys

`func (o *AudienceConfiguration) GetNotifyTeamKeys() []string`

GetNotifyTeamKeys returns the NotifyTeamKeys field if non-nil, zero value otherwise.

### GetNotifyTeamKeysOk

`func (o *AudienceConfiguration) GetNotifyTeamKeysOk() (*[]string, bool)`

GetNotifyTeamKeysOk returns a tuple with the NotifyTeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTeamKeys

`func (o *AudienceConfiguration) SetNotifyTeamKeys(v []string)`

SetNotifyTeamKeys sets NotifyTeamKeys field to given value.

### HasNotifyTeamKeys

`func (o *AudienceConfiguration) HasNotifyTeamKeys() bool`

HasNotifyTeamKeys returns a boolean if a field has been set.

### GetReleaseGuardianConfiguration

`func (o *AudienceConfiguration) GetReleaseGuardianConfiguration() ReleaseGuardianConfiguration`

GetReleaseGuardianConfiguration returns the ReleaseGuardianConfiguration field if non-nil, zero value otherwise.

### GetReleaseGuardianConfigurationOk

`func (o *AudienceConfiguration) GetReleaseGuardianConfigurationOk() (*ReleaseGuardianConfiguration, bool)`

GetReleaseGuardianConfigurationOk returns a tuple with the ReleaseGuardianConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGuardianConfiguration

`func (o *AudienceConfiguration) SetReleaseGuardianConfiguration(v ReleaseGuardianConfiguration)`

SetReleaseGuardianConfiguration sets ReleaseGuardianConfiguration field to given value.

### HasReleaseGuardianConfiguration

`func (o *AudienceConfiguration) HasReleaseGuardianConfiguration() bool`

HasReleaseGuardianConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


