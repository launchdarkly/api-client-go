# ExpandableApprovalRequestResponse

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
**AppliedByServiceTokenId** | Pointer to **string** | The service token ID of the service token which applied the approval request | [optional] 
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
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Environments** | Pointer to [**[]Environment**](Environment.md) | List of environments the approval impacts | [optional] 
**Flag** | Pointer to [**ExpandedFlagRep**](ExpandedFlagRep.md) |  | [optional] 
**Resource** | Pointer to [**ExpandedResourceRep**](ExpandedResourceRep.md) |  | [optional] 

## Methods

### NewExpandableApprovalRequestResponse

`func NewExpandableApprovalRequestResponse(id string, version int32, creationDate int64, serviceKind string, reviewStatus string, allReviews []ReviewResponse, notifyMemberIds []string, status string, instructions []map[string]interface{}, conflicts []Conflict, links map[string]interface{}, ) *ExpandableApprovalRequestResponse`

NewExpandableApprovalRequestResponse instantiates a new ExpandableApprovalRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandableApprovalRequestResponseWithDefaults

`func NewExpandableApprovalRequestResponseWithDefaults() *ExpandableApprovalRequestResponse`

NewExpandableApprovalRequestResponseWithDefaults instantiates a new ExpandableApprovalRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpandableApprovalRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpandableApprovalRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpandableApprovalRequestResponse) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ExpandableApprovalRequestResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpandableApprovalRequestResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpandableApprovalRequestResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *ExpandableApprovalRequestResponse) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ExpandableApprovalRequestResponse) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ExpandableApprovalRequestResponse) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetServiceKind

`func (o *ExpandableApprovalRequestResponse) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *ExpandableApprovalRequestResponse) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *ExpandableApprovalRequestResponse) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.


### GetRequestorId

`func (o *ExpandableApprovalRequestResponse) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *ExpandableApprovalRequestResponse) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *ExpandableApprovalRequestResponse) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *ExpandableApprovalRequestResponse) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### GetDescription

`func (o *ExpandableApprovalRequestResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandableApprovalRequestResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandableApprovalRequestResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandableApprovalRequestResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReviewStatus

`func (o *ExpandableApprovalRequestResponse) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ExpandableApprovalRequestResponse) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ExpandableApprovalRequestResponse) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAllReviews

`func (o *ExpandableApprovalRequestResponse) GetAllReviews() []ReviewResponse`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *ExpandableApprovalRequestResponse) GetAllReviewsOk() (*[]ReviewResponse, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *ExpandableApprovalRequestResponse) SetAllReviews(v []ReviewResponse)`

SetAllReviews sets AllReviews field to given value.


### GetNotifyMemberIds

`func (o *ExpandableApprovalRequestResponse) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ExpandableApprovalRequestResponse) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ExpandableApprovalRequestResponse) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAppliedDate

`func (o *ExpandableApprovalRequestResponse) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *ExpandableApprovalRequestResponse) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *ExpandableApprovalRequestResponse) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *ExpandableApprovalRequestResponse) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.

### GetAppliedByMemberId

`func (o *ExpandableApprovalRequestResponse) GetAppliedByMemberId() string`

GetAppliedByMemberId returns the AppliedByMemberId field if non-nil, zero value otherwise.

### GetAppliedByMemberIdOk

`func (o *ExpandableApprovalRequestResponse) GetAppliedByMemberIdOk() (*string, bool)`

GetAppliedByMemberIdOk returns a tuple with the AppliedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByMemberId

`func (o *ExpandableApprovalRequestResponse) SetAppliedByMemberId(v string)`

SetAppliedByMemberId sets AppliedByMemberId field to given value.

### HasAppliedByMemberId

`func (o *ExpandableApprovalRequestResponse) HasAppliedByMemberId() bool`

HasAppliedByMemberId returns a boolean if a field has been set.

### GetAppliedByServiceTokenId

`func (o *ExpandableApprovalRequestResponse) GetAppliedByServiceTokenId() string`

GetAppliedByServiceTokenId returns the AppliedByServiceTokenId field if non-nil, zero value otherwise.

### GetAppliedByServiceTokenIdOk

`func (o *ExpandableApprovalRequestResponse) GetAppliedByServiceTokenIdOk() (*string, bool)`

GetAppliedByServiceTokenIdOk returns a tuple with the AppliedByServiceTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByServiceTokenId

`func (o *ExpandableApprovalRequestResponse) SetAppliedByServiceTokenId(v string)`

SetAppliedByServiceTokenId sets AppliedByServiceTokenId field to given value.

### HasAppliedByServiceTokenId

`func (o *ExpandableApprovalRequestResponse) HasAppliedByServiceTokenId() bool`

HasAppliedByServiceTokenId returns a boolean if a field has been set.

### GetStatus

`func (o *ExpandableApprovalRequestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpandableApprovalRequestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpandableApprovalRequestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstructions

`func (o *ExpandableApprovalRequestResponse) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ExpandableApprovalRequestResponse) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ExpandableApprovalRequestResponse) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetConflicts

`func (o *ExpandableApprovalRequestResponse) GetConflicts() []Conflict`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ExpandableApprovalRequestResponse) GetConflictsOk() (*[]Conflict, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ExpandableApprovalRequestResponse) SetConflicts(v []Conflict)`

SetConflicts sets Conflicts field to given value.


### GetLinks

`func (o *ExpandableApprovalRequestResponse) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandableApprovalRequestResponse) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandableApprovalRequestResponse) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetExecutionDate

`func (o *ExpandableApprovalRequestResponse) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ExpandableApprovalRequestResponse) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ExpandableApprovalRequestResponse) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ExpandableApprovalRequestResponse) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetOperatingOnId

`func (o *ExpandableApprovalRequestResponse) GetOperatingOnId() string`

GetOperatingOnId returns the OperatingOnId field if non-nil, zero value otherwise.

### GetOperatingOnIdOk

`func (o *ExpandableApprovalRequestResponse) GetOperatingOnIdOk() (*string, bool)`

GetOperatingOnIdOk returns a tuple with the OperatingOnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingOnId

`func (o *ExpandableApprovalRequestResponse) SetOperatingOnId(v string)`

SetOperatingOnId sets OperatingOnId field to given value.

### HasOperatingOnId

`func (o *ExpandableApprovalRequestResponse) HasOperatingOnId() bool`

HasOperatingOnId returns a boolean if a field has been set.

### GetIntegrationMetadata

`func (o *ExpandableApprovalRequestResponse) GetIntegrationMetadata() IntegrationMetadata`

GetIntegrationMetadata returns the IntegrationMetadata field if non-nil, zero value otherwise.

### GetIntegrationMetadataOk

`func (o *ExpandableApprovalRequestResponse) GetIntegrationMetadataOk() (*IntegrationMetadata, bool)`

GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationMetadata

`func (o *ExpandableApprovalRequestResponse) SetIntegrationMetadata(v IntegrationMetadata)`

SetIntegrationMetadata sets IntegrationMetadata field to given value.

### HasIntegrationMetadata

`func (o *ExpandableApprovalRequestResponse) HasIntegrationMetadata() bool`

HasIntegrationMetadata returns a boolean if a field has been set.

### GetSource

`func (o *ExpandableApprovalRequestResponse) GetSource() CopiedFromEnv`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExpandableApprovalRequestResponse) GetSourceOk() (*CopiedFromEnv, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExpandableApprovalRequestResponse) SetSource(v CopiedFromEnv)`

SetSource sets Source field to given value.

### HasSource

`func (o *ExpandableApprovalRequestResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCustomWorkflowMetadata

`func (o *ExpandableApprovalRequestResponse) GetCustomWorkflowMetadata() CustomWorkflowMeta`

GetCustomWorkflowMetadata returns the CustomWorkflowMetadata field if non-nil, zero value otherwise.

### GetCustomWorkflowMetadataOk

`func (o *ExpandableApprovalRequestResponse) GetCustomWorkflowMetadataOk() (*CustomWorkflowMeta, bool)`

GetCustomWorkflowMetadataOk returns a tuple with the CustomWorkflowMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWorkflowMetadata

`func (o *ExpandableApprovalRequestResponse) SetCustomWorkflowMetadata(v CustomWorkflowMeta)`

SetCustomWorkflowMetadata sets CustomWorkflowMetadata field to given value.

### HasCustomWorkflowMetadata

`func (o *ExpandableApprovalRequestResponse) HasCustomWorkflowMetadata() bool`

HasCustomWorkflowMetadata returns a boolean if a field has been set.

### GetResourceId

`func (o *ExpandableApprovalRequestResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExpandableApprovalRequestResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExpandableApprovalRequestResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ExpandableApprovalRequestResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetApprovalSettings

`func (o *ExpandableApprovalRequestResponse) GetApprovalSettings() ApprovalSettings`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *ExpandableApprovalRequestResponse) GetApprovalSettingsOk() (*ApprovalSettings, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *ExpandableApprovalRequestResponse) SetApprovalSettings(v ApprovalSettings)`

SetApprovalSettings sets ApprovalSettings field to given value.

### HasApprovalSettings

`func (o *ExpandableApprovalRequestResponse) HasApprovalSettings() bool`

HasApprovalSettings returns a boolean if a field has been set.

### GetProject

`func (o *ExpandableApprovalRequestResponse) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ExpandableApprovalRequestResponse) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ExpandableApprovalRequestResponse) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ExpandableApprovalRequestResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironments

`func (o *ExpandableApprovalRequestResponse) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ExpandableApprovalRequestResponse) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ExpandableApprovalRequestResponse) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ExpandableApprovalRequestResponse) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetFlag

`func (o *ExpandableApprovalRequestResponse) GetFlag() ExpandedFlagRep`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ExpandableApprovalRequestResponse) GetFlagOk() (*ExpandedFlagRep, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ExpandableApprovalRequestResponse) SetFlag(v ExpandedFlagRep)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ExpandableApprovalRequestResponse) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetResource

`func (o *ExpandableApprovalRequestResponse) GetResource() ExpandedResourceRep`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ExpandableApprovalRequestResponse) GetResourceOk() (*ExpandedResourceRep, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ExpandableApprovalRequestResponse) SetResource(v ExpandedResourceRep)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ExpandableApprovalRequestResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


