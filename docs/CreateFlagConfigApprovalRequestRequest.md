# CreateFlagConfigApprovalRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment describing the approval request | [optional] 
**Description** | **string** | A human-friendly name for the approval request | 
**Instructions** | **[]interface{}** |  | 
**NotifyMemberIds** | **[]string** | An array of member IDs. These members are notified to review the approval request | 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**OperatingOnId** | Pointer to **string** | ID of scheduled change to edit or delete | [optional] 
**IntegrationConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateFlagConfigApprovalRequestRequest

`func NewCreateFlagConfigApprovalRequestRequest(description string, instructions []interface{}, notifyMemberIds []string, ) *CreateFlagConfigApprovalRequestRequest`

NewCreateFlagConfigApprovalRequestRequest instantiates a new CreateFlagConfigApprovalRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlagConfigApprovalRequestRequestWithDefaults

`func NewCreateFlagConfigApprovalRequestRequestWithDefaults() *CreateFlagConfigApprovalRequestRequest`

NewCreateFlagConfigApprovalRequestRequestWithDefaults instantiates a new CreateFlagConfigApprovalRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateFlagConfigApprovalRequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateFlagConfigApprovalRequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateFlagConfigApprovalRequestRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *CreateFlagConfigApprovalRequestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFlagConfigApprovalRequestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstructions

`func (o *CreateFlagConfigApprovalRequestRequest) GetInstructions() []interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetInstructionsOk() (*[]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateFlagConfigApprovalRequestRequest) SetInstructions(v []interface{})`

SetInstructions sets Instructions field to given value.


### GetNotifyMemberIds

`func (o *CreateFlagConfigApprovalRequestRequest) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *CreateFlagConfigApprovalRequestRequest) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetExecutionDate

`func (o *CreateFlagConfigApprovalRequestRequest) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *CreateFlagConfigApprovalRequestRequest) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *CreateFlagConfigApprovalRequestRequest) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetOperatingOnId

`func (o *CreateFlagConfigApprovalRequestRequest) GetOperatingOnId() string`

GetOperatingOnId returns the OperatingOnId field if non-nil, zero value otherwise.

### GetOperatingOnIdOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetOperatingOnIdOk() (*string, bool)`

GetOperatingOnIdOk returns a tuple with the OperatingOnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingOnId

`func (o *CreateFlagConfigApprovalRequestRequest) SetOperatingOnId(v string)`

SetOperatingOnId sets OperatingOnId field to given value.

### HasOperatingOnId

`func (o *CreateFlagConfigApprovalRequestRequest) HasOperatingOnId() bool`

HasOperatingOnId returns a boolean if a field has been set.

### GetIntegrationConfig

`func (o *CreateFlagConfigApprovalRequestRequest) GetIntegrationConfig() map[string]interface{}`

GetIntegrationConfig returns the IntegrationConfig field if non-nil, zero value otherwise.

### GetIntegrationConfigOk

`func (o *CreateFlagConfigApprovalRequestRequest) GetIntegrationConfigOk() (*map[string]interface{}, bool)`

GetIntegrationConfigOk returns a tuple with the IntegrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfig

`func (o *CreateFlagConfigApprovalRequestRequest) SetIntegrationConfig(v map[string]interface{})`

SetIntegrationConfig sets IntegrationConfig field to given value.

### HasIntegrationConfig

`func (o *CreateFlagConfigApprovalRequestRequest) HasIntegrationConfig() bool`

HasIntegrationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


