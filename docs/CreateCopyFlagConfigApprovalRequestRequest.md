# CreateCopyFlagConfigApprovalRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional comment describing the approval request | [optional] 
**Description** | **string** | A brief description of your changes | 
**NotifyMemberIds** | Pointer to **[]string** | An array of member IDs. These members are notified to review the approval request. | [optional] 
**NotifyTeamKeys** | Pointer to **[]string** | An array of team keys. The members of these teams are notified to review the approval request. | [optional] 
**Source** | [**SourceFlag**](SourceFlag.md) |  | 
**IncludedActions** | Pointer to **[]string** | Optional list of the flag changes to copy from the source environment to the target environment. You may include either &lt;code&gt;includedActions&lt;/code&gt; or &lt;code&gt;excludedActions&lt;/code&gt;, but not both. If neither are included, then all flag changes will be copied. | [optional] 
**ExcludedActions** | Pointer to **[]string** | Optional list of the flag changes NOT to copy from the source environment to the target environment. You may include either &lt;code&gt;includedActions&lt;/code&gt; or &lt;code&gt;excludedActions&lt;/code&gt;, but not both. If neither are included, then all flag changes will be copied. | [optional] 

## Methods

### NewCreateCopyFlagConfigApprovalRequestRequest

`func NewCreateCopyFlagConfigApprovalRequestRequest(description string, source SourceFlag, ) *CreateCopyFlagConfigApprovalRequestRequest`

NewCreateCopyFlagConfigApprovalRequestRequest instantiates a new CreateCopyFlagConfigApprovalRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCopyFlagConfigApprovalRequestRequestWithDefaults

`func NewCreateCopyFlagConfigApprovalRequestRequestWithDefaults() *CreateCopyFlagConfigApprovalRequestRequest`

NewCreateCopyFlagConfigApprovalRequestRequestWithDefaults instantiates a new CreateCopyFlagConfigApprovalRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateCopyFlagConfigApprovalRequestRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifyMemberIds

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *CreateCopyFlagConfigApprovalRequestRequest) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetNotifyTeamKeys

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetNotifyTeamKeys() []string`

GetNotifyTeamKeys returns the NotifyTeamKeys field if non-nil, zero value otherwise.

### GetNotifyTeamKeysOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetNotifyTeamKeysOk() (*[]string, bool)`

GetNotifyTeamKeysOk returns a tuple with the NotifyTeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTeamKeys

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetNotifyTeamKeys(v []string)`

SetNotifyTeamKeys sets NotifyTeamKeys field to given value.

### HasNotifyTeamKeys

`func (o *CreateCopyFlagConfigApprovalRequestRequest) HasNotifyTeamKeys() bool`

HasNotifyTeamKeys returns a boolean if a field has been set.

### GetSource

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetSource() SourceFlag`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetSourceOk() (*SourceFlag, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetSource(v SourceFlag)`

SetSource sets Source field to given value.


### GetIncludedActions

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetIncludedActions() []string`

GetIncludedActions returns the IncludedActions field if non-nil, zero value otherwise.

### GetIncludedActionsOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetIncludedActionsOk() (*[]string, bool)`

GetIncludedActionsOk returns a tuple with the IncludedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedActions

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetIncludedActions(v []string)`

SetIncludedActions sets IncludedActions field to given value.

### HasIncludedActions

`func (o *CreateCopyFlagConfigApprovalRequestRequest) HasIncludedActions() bool`

HasIncludedActions returns a boolean if a field has been set.

### GetExcludedActions

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetExcludedActions() []string`

GetExcludedActions returns the ExcludedActions field if non-nil, zero value otherwise.

### GetExcludedActionsOk

`func (o *CreateCopyFlagConfigApprovalRequestRequest) GetExcludedActionsOk() (*[]string, bool)`

GetExcludedActionsOk returns a tuple with the ExcludedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedActions

`func (o *CreateCopyFlagConfigApprovalRequestRequest) SetExcludedActions(v []string)`

SetExcludedActions sets ExcludedActions field to given value.

### HasExcludedActions

`func (o *CreateCopyFlagConfigApprovalRequestRequest) HasExcludedActions() bool`

HasExcludedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


