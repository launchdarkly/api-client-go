# CreateApprovalRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | String representation of the resource specifier | 
**Comment** | Pointer to **string** | Optional comment describing the approval request | [optional] 
**Description** | **string** | A brief description of the changes you&#39;re requesting | 
**Instructions** | **[]map[string]interface{}** |  | 
**NotifyMemberIds** | Pointer to **[]string** | An array of member IDs. These members are notified to review the approval request. | [optional] 
**NotifyTeamKeys** | Pointer to **[]string** | An array of team keys. The members of these teams are notified to review the approval request. | [optional] 
**IntegrationConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateApprovalRequestRequest

`func NewCreateApprovalRequestRequest(resourceId string, description string, instructions []map[string]interface{}, ) *CreateApprovalRequestRequest`

NewCreateApprovalRequestRequest instantiates a new CreateApprovalRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApprovalRequestRequestWithDefaults

`func NewCreateApprovalRequestRequestWithDefaults() *CreateApprovalRequestRequest`

NewCreateApprovalRequestRequestWithDefaults instantiates a new CreateApprovalRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *CreateApprovalRequestRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CreateApprovalRequestRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CreateApprovalRequestRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetComment

`func (o *CreateApprovalRequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateApprovalRequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateApprovalRequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateApprovalRequestRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateApprovalRequestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApprovalRequestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApprovalRequestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstructions

`func (o *CreateApprovalRequestRequest) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateApprovalRequestRequest) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateApprovalRequestRequest) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetNotifyMemberIds

`func (o *CreateApprovalRequestRequest) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *CreateApprovalRequestRequest) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *CreateApprovalRequestRequest) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *CreateApprovalRequestRequest) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetNotifyTeamKeys

`func (o *CreateApprovalRequestRequest) GetNotifyTeamKeys() []string`

GetNotifyTeamKeys returns the NotifyTeamKeys field if non-nil, zero value otherwise.

### GetNotifyTeamKeysOk

`func (o *CreateApprovalRequestRequest) GetNotifyTeamKeysOk() (*[]string, bool)`

GetNotifyTeamKeysOk returns a tuple with the NotifyTeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTeamKeys

`func (o *CreateApprovalRequestRequest) SetNotifyTeamKeys(v []string)`

SetNotifyTeamKeys sets NotifyTeamKeys field to given value.

### HasNotifyTeamKeys

`func (o *CreateApprovalRequestRequest) HasNotifyTeamKeys() bool`

HasNotifyTeamKeys returns a boolean if a field has been set.

### GetIntegrationConfig

`func (o *CreateApprovalRequestRequest) GetIntegrationConfig() map[string]interface{}`

GetIntegrationConfig returns the IntegrationConfig field if non-nil, zero value otherwise.

### GetIntegrationConfigOk

`func (o *CreateApprovalRequestRequest) GetIntegrationConfigOk() (*map[string]interface{}, bool)`

GetIntegrationConfigOk returns a tuple with the IntegrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfig

`func (o *CreateApprovalRequestRequest) SetIntegrationConfig(v map[string]interface{})`

SetIntegrationConfig sets IntegrationConfig field to given value.

### HasIntegrationConfig

`func (o *CreateApprovalRequestRequest) HasIntegrationConfig() bool`

HasIntegrationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


