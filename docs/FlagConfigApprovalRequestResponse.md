# FlagConfigApprovalRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | **int32** |  | 
**CreationDate** | **int64** |  | 
**ServiceKind** | **string** |  | 
**RequestorId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | A human-friendly name for the approval request | [optional] 
**ReviewStatus** | **string** |  | 
**AllReviews** | [**[]ReviewResponse**](ReviewResponse.md) |  | 
**NotifyMemberIds** | **[]string** | An array of member IDs. These members are notified to review the approval request. | 
**AppliedDate** | Pointer to **int64** |  | [optional] 
**AppliedByMemberId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Instructions** | **[]map[string]interface{}** |  | 
**Conflicts** | [**[]Conflict**](Conflict.md) |  | 
**Links** | [**map[string]Link**](Link.md) |  | 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**OperatingOnId** | Pointer to **string** | ID of scheduled change to edit or delete | [optional] 
**IntegrationMetadata** | Pointer to [**IntegrationMetadata**](IntegrationMetadata.md) |  | [optional] 
**Source** | Pointer to [**CopiedFromEnv**](CopiedFromEnv.md) |  | [optional] 

## Methods

### NewFlagConfigApprovalRequestResponse

`func NewFlagConfigApprovalRequestResponse(id string, version int32, creationDate int64, serviceKind string, reviewStatus string, allReviews []ReviewResponse, notifyMemberIds []string, status string, instructions []map[string]interface{}, conflicts []Conflict, links map[string]Link, ) *FlagConfigApprovalRequestResponse`

NewFlagConfigApprovalRequestResponse instantiates a new FlagConfigApprovalRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagConfigApprovalRequestResponseWithDefaults

`func NewFlagConfigApprovalRequestResponseWithDefaults() *FlagConfigApprovalRequestResponse`

NewFlagConfigApprovalRequestResponseWithDefaults instantiates a new FlagConfigApprovalRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlagConfigApprovalRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagConfigApprovalRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagConfigApprovalRequestResponse) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *FlagConfigApprovalRequestResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlagConfigApprovalRequestResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlagConfigApprovalRequestResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *FlagConfigApprovalRequestResponse) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *FlagConfigApprovalRequestResponse) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *FlagConfigApprovalRequestResponse) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetServiceKind

`func (o *FlagConfigApprovalRequestResponse) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *FlagConfigApprovalRequestResponse) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *FlagConfigApprovalRequestResponse) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.


### GetRequestorId

`func (o *FlagConfigApprovalRequestResponse) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *FlagConfigApprovalRequestResponse) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *FlagConfigApprovalRequestResponse) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *FlagConfigApprovalRequestResponse) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### GetDescription

`func (o *FlagConfigApprovalRequestResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagConfigApprovalRequestResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagConfigApprovalRequestResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagConfigApprovalRequestResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReviewStatus

`func (o *FlagConfigApprovalRequestResponse) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *FlagConfigApprovalRequestResponse) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *FlagConfigApprovalRequestResponse) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAllReviews

`func (o *FlagConfigApprovalRequestResponse) GetAllReviews() []ReviewResponse`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *FlagConfigApprovalRequestResponse) GetAllReviewsOk() (*[]ReviewResponse, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *FlagConfigApprovalRequestResponse) SetAllReviews(v []ReviewResponse)`

SetAllReviews sets AllReviews field to given value.


### GetNotifyMemberIds

`func (o *FlagConfigApprovalRequestResponse) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *FlagConfigApprovalRequestResponse) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *FlagConfigApprovalRequestResponse) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAppliedDate

`func (o *FlagConfigApprovalRequestResponse) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *FlagConfigApprovalRequestResponse) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *FlagConfigApprovalRequestResponse) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *FlagConfigApprovalRequestResponse) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.

### GetAppliedByMemberId

`func (o *FlagConfigApprovalRequestResponse) GetAppliedByMemberId() string`

GetAppliedByMemberId returns the AppliedByMemberId field if non-nil, zero value otherwise.

### GetAppliedByMemberIdOk

`func (o *FlagConfigApprovalRequestResponse) GetAppliedByMemberIdOk() (*string, bool)`

GetAppliedByMemberIdOk returns a tuple with the AppliedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByMemberId

`func (o *FlagConfigApprovalRequestResponse) SetAppliedByMemberId(v string)`

SetAppliedByMemberId sets AppliedByMemberId field to given value.

### HasAppliedByMemberId

`func (o *FlagConfigApprovalRequestResponse) HasAppliedByMemberId() bool`

HasAppliedByMemberId returns a boolean if a field has been set.

### GetStatus

`func (o *FlagConfigApprovalRequestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlagConfigApprovalRequestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlagConfigApprovalRequestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstructions

`func (o *FlagConfigApprovalRequestResponse) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *FlagConfigApprovalRequestResponse) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *FlagConfigApprovalRequestResponse) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetConflicts

`func (o *FlagConfigApprovalRequestResponse) GetConflicts() []Conflict`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *FlagConfigApprovalRequestResponse) GetConflictsOk() (*[]Conflict, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *FlagConfigApprovalRequestResponse) SetConflicts(v []Conflict)`

SetConflicts sets Conflicts field to given value.


### GetLinks

`func (o *FlagConfigApprovalRequestResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagConfigApprovalRequestResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagConfigApprovalRequestResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetExecutionDate

`func (o *FlagConfigApprovalRequestResponse) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *FlagConfigApprovalRequestResponse) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *FlagConfigApprovalRequestResponse) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *FlagConfigApprovalRequestResponse) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetOperatingOnId

`func (o *FlagConfigApprovalRequestResponse) GetOperatingOnId() string`

GetOperatingOnId returns the OperatingOnId field if non-nil, zero value otherwise.

### GetOperatingOnIdOk

`func (o *FlagConfigApprovalRequestResponse) GetOperatingOnIdOk() (*string, bool)`

GetOperatingOnIdOk returns a tuple with the OperatingOnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingOnId

`func (o *FlagConfigApprovalRequestResponse) SetOperatingOnId(v string)`

SetOperatingOnId sets OperatingOnId field to given value.

### HasOperatingOnId

`func (o *FlagConfigApprovalRequestResponse) HasOperatingOnId() bool`

HasOperatingOnId returns a boolean if a field has been set.

### GetIntegrationMetadata

`func (o *FlagConfigApprovalRequestResponse) GetIntegrationMetadata() IntegrationMetadata`

GetIntegrationMetadata returns the IntegrationMetadata field if non-nil, zero value otherwise.

### GetIntegrationMetadataOk

`func (o *FlagConfigApprovalRequestResponse) GetIntegrationMetadataOk() (*IntegrationMetadata, bool)`

GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationMetadata

`func (o *FlagConfigApprovalRequestResponse) SetIntegrationMetadata(v IntegrationMetadata)`

SetIntegrationMetadata sets IntegrationMetadata field to given value.

### HasIntegrationMetadata

`func (o *FlagConfigApprovalRequestResponse) HasIntegrationMetadata() bool`

HasIntegrationMetadata returns a boolean if a field has been set.

### GetSource

`func (o *FlagConfigApprovalRequestResponse) GetSource() CopiedFromEnv`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FlagConfigApprovalRequestResponse) GetSourceOk() (*CopiedFromEnv, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FlagConfigApprovalRequestResponse) SetSource(v CopiedFromEnv)`

SetSource sets Source field to given value.

### HasSource

`func (o *FlagConfigApprovalRequestResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


