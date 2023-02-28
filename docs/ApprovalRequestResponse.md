# ApprovalRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this approval request | 
**Version** | **int32** | Version of the approval request | 
**CreationDate** | **int64** |  | 
**ServiceKind** | **string** |  | 
**RequestorId** | Pointer to **string** | The ID of the member who requested the approval | [optional] 
**Description** | Pointer to **string** | A human-friendly name for the approval request | [optional] 
**ReviewStatus** | **string** | Current status of the review of this approval request | 
**AllReviews** | [**[]ReviewResponse**](ReviewResponse.md) | An array of individual reviews of this approval request | 
**NotifyMemberIds** | **[]string** | An array of member IDs. These members are notified to review the approval request. | 
**AppliedDate** | Pointer to **int64** |  | [optional] 
**AppliedByMemberId** | Pointer to **string** | The member ID of the member who applied the approval request | [optional] 
**Status** | **string** | Current status of the approval request | 
**Instructions** | **[]map[string]interface{}** |  | 
**Conflicts** | [**[]Conflict**](Conflict.md) | Details on any conflicting approval requests | 
**Links** | **map[string]interface{}** | The location and content type of related resources | 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**OperatingOnId** | Pointer to **string** | ID of scheduled change to edit or delete | [optional] 
**IntegrationMetadata** | Pointer to [**IntegrationMetadata**](IntegrationMetadata.md) |  | [optional] 
**Source** | Pointer to [**CopiedFromEnv**](CopiedFromEnv.md) |  | [optional] 
**CustomWorkflowMetadata** | Pointer to [**CustomWorkflowMeta**](CustomWorkflowMeta.md) |  | [optional] 
**ResourceId** | Pointer to **string** | String representation of a resource | [optional] 
**ApprovalSettings** | Pointer to [**ApprovalSettings**](ApprovalSettings.md) |  | [optional] 

## Methods

### NewApprovalRequestResponse

`func NewApprovalRequestResponse(id string, version int32, creationDate int64, serviceKind string, reviewStatus string, allReviews []ReviewResponse, notifyMemberIds []string, status string, instructions []map[string]interface{}, conflicts []Conflict, links map[string]interface{}, ) *ApprovalRequestResponse`

NewApprovalRequestResponse instantiates a new ApprovalRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseWithDefaults

`func NewApprovalRequestResponseWithDefaults() *ApprovalRequestResponse`

NewApprovalRequestResponseWithDefaults instantiates a new ApprovalRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestResponse) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ApprovalRequestResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApprovalRequestResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApprovalRequestResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *ApprovalRequestResponse) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApprovalRequestResponse) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApprovalRequestResponse) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetServiceKind

`func (o *ApprovalRequestResponse) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *ApprovalRequestResponse) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *ApprovalRequestResponse) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.


### GetRequestorId

`func (o *ApprovalRequestResponse) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *ApprovalRequestResponse) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *ApprovalRequestResponse) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *ApprovalRequestResponse) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### GetDescription

`func (o *ApprovalRequestResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalRequestResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalRequestResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalRequestResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReviewStatus

`func (o *ApprovalRequestResponse) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ApprovalRequestResponse) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ApprovalRequestResponse) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAllReviews

`func (o *ApprovalRequestResponse) GetAllReviews() []ReviewResponse`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *ApprovalRequestResponse) GetAllReviewsOk() (*[]ReviewResponse, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *ApprovalRequestResponse) SetAllReviews(v []ReviewResponse)`

SetAllReviews sets AllReviews field to given value.


### GetNotifyMemberIds

`func (o *ApprovalRequestResponse) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ApprovalRequestResponse) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ApprovalRequestResponse) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAppliedDate

`func (o *ApprovalRequestResponse) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *ApprovalRequestResponse) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *ApprovalRequestResponse) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *ApprovalRequestResponse) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.

### GetAppliedByMemberId

`func (o *ApprovalRequestResponse) GetAppliedByMemberId() string`

GetAppliedByMemberId returns the AppliedByMemberId field if non-nil, zero value otherwise.

### GetAppliedByMemberIdOk

`func (o *ApprovalRequestResponse) GetAppliedByMemberIdOk() (*string, bool)`

GetAppliedByMemberIdOk returns a tuple with the AppliedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByMemberId

`func (o *ApprovalRequestResponse) SetAppliedByMemberId(v string)`

SetAppliedByMemberId sets AppliedByMemberId field to given value.

### HasAppliedByMemberId

`func (o *ApprovalRequestResponse) HasAppliedByMemberId() bool`

HasAppliedByMemberId returns a boolean if a field has been set.

### GetStatus

`func (o *ApprovalRequestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstructions

`func (o *ApprovalRequestResponse) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ApprovalRequestResponse) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ApprovalRequestResponse) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetConflicts

`func (o *ApprovalRequestResponse) GetConflicts() []Conflict`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ApprovalRequestResponse) GetConflictsOk() (*[]Conflict, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ApprovalRequestResponse) SetConflicts(v []Conflict)`

SetConflicts sets Conflicts field to given value.


### GetLinks

`func (o *ApprovalRequestResponse) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApprovalRequestResponse) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApprovalRequestResponse) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetExecutionDate

`func (o *ApprovalRequestResponse) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ApprovalRequestResponse) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ApprovalRequestResponse) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ApprovalRequestResponse) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetOperatingOnId

`func (o *ApprovalRequestResponse) GetOperatingOnId() string`

GetOperatingOnId returns the OperatingOnId field if non-nil, zero value otherwise.

### GetOperatingOnIdOk

`func (o *ApprovalRequestResponse) GetOperatingOnIdOk() (*string, bool)`

GetOperatingOnIdOk returns a tuple with the OperatingOnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingOnId

`func (o *ApprovalRequestResponse) SetOperatingOnId(v string)`

SetOperatingOnId sets OperatingOnId field to given value.

### HasOperatingOnId

`func (o *ApprovalRequestResponse) HasOperatingOnId() bool`

HasOperatingOnId returns a boolean if a field has been set.

### GetIntegrationMetadata

`func (o *ApprovalRequestResponse) GetIntegrationMetadata() IntegrationMetadata`

GetIntegrationMetadata returns the IntegrationMetadata field if non-nil, zero value otherwise.

### GetIntegrationMetadataOk

`func (o *ApprovalRequestResponse) GetIntegrationMetadataOk() (*IntegrationMetadata, bool)`

GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationMetadata

`func (o *ApprovalRequestResponse) SetIntegrationMetadata(v IntegrationMetadata)`

SetIntegrationMetadata sets IntegrationMetadata field to given value.

### HasIntegrationMetadata

`func (o *ApprovalRequestResponse) HasIntegrationMetadata() bool`

HasIntegrationMetadata returns a boolean if a field has been set.

### GetSource

`func (o *ApprovalRequestResponse) GetSource() CopiedFromEnv`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApprovalRequestResponse) GetSourceOk() (*CopiedFromEnv, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApprovalRequestResponse) SetSource(v CopiedFromEnv)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApprovalRequestResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCustomWorkflowMetadata

`func (o *ApprovalRequestResponse) GetCustomWorkflowMetadata() CustomWorkflowMeta`

GetCustomWorkflowMetadata returns the CustomWorkflowMetadata field if non-nil, zero value otherwise.

### GetCustomWorkflowMetadataOk

`func (o *ApprovalRequestResponse) GetCustomWorkflowMetadataOk() (*CustomWorkflowMeta, bool)`

GetCustomWorkflowMetadataOk returns a tuple with the CustomWorkflowMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWorkflowMetadata

`func (o *ApprovalRequestResponse) SetCustomWorkflowMetadata(v CustomWorkflowMeta)`

SetCustomWorkflowMetadata sets CustomWorkflowMetadata field to given value.

### HasCustomWorkflowMetadata

`func (o *ApprovalRequestResponse) HasCustomWorkflowMetadata() bool`

HasCustomWorkflowMetadata returns a boolean if a field has been set.

### GetResourceId

`func (o *ApprovalRequestResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ApprovalRequestResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ApprovalRequestResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ApprovalRequestResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetApprovalSettings

`func (o *ApprovalRequestResponse) GetApprovalSettings() ApprovalSettings`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *ApprovalRequestResponse) GetApprovalSettingsOk() (*ApprovalSettings, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *ApprovalRequestResponse) SetApprovalSettings(v ApprovalSettings)`

SetApprovalSettings sets ApprovalSettings field to given value.

### HasApprovalSettings

`func (o *ApprovalRequestResponse) HasApprovalSettings() bool`

HasApprovalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


