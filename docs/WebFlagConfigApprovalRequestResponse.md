# WebFlagConfigApprovalRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**ServiceKind** | Pointer to **string** |  | [optional] 
**RequestorId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ReviewStatus** | Pointer to **string** |  | [optional] 
**AllReviews** | Pointer to [**[]WebReviewResponse**](WebReviewResponse.md) |  | [optional] 
**NotifyMemberIds** | Pointer to **[]string** |  | [optional] 
**AppliedDate** | Pointer to **int64** |  | [optional] 
**AppliedByMemberId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Conflicts** | Pointer to [**[]WebConflict**](WebConflict.md) |  | [optional] 
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**OperatingOnId** | Pointer to **string** |  | [optional] 
**IntegrationMetadata** | Pointer to [**WebIntegrationMetadata**](WebIntegrationMetadata.md) |  | [optional] 
**Source** | Pointer to [**WebCopiedFromEnv**](WebCopiedFromEnv.md) |  | [optional] 

## Methods

### NewWebFlagConfigApprovalRequestResponse

`func NewWebFlagConfigApprovalRequestResponse() *WebFlagConfigApprovalRequestResponse`

NewWebFlagConfigApprovalRequestResponse instantiates a new WebFlagConfigApprovalRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebFlagConfigApprovalRequestResponseWithDefaults

`func NewWebFlagConfigApprovalRequestResponseWithDefaults() *WebFlagConfigApprovalRequestResponse`

NewWebFlagConfigApprovalRequestResponseWithDefaults instantiates a new WebFlagConfigApprovalRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebFlagConfigApprovalRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebFlagConfigApprovalRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebFlagConfigApprovalRequestResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebFlagConfigApprovalRequestResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *WebFlagConfigApprovalRequestResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebFlagConfigApprovalRequestResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebFlagConfigApprovalRequestResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebFlagConfigApprovalRequestResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationDate

`func (o *WebFlagConfigApprovalRequestResponse) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *WebFlagConfigApprovalRequestResponse) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *WebFlagConfigApprovalRequestResponse) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *WebFlagConfigApprovalRequestResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetServiceKind

`func (o *WebFlagConfigApprovalRequestResponse) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *WebFlagConfigApprovalRequestResponse) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *WebFlagConfigApprovalRequestResponse) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.

### HasServiceKind

`func (o *WebFlagConfigApprovalRequestResponse) HasServiceKind() bool`

HasServiceKind returns a boolean if a field has been set.

### GetRequestorId

`func (o *WebFlagConfigApprovalRequestResponse) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *WebFlagConfigApprovalRequestResponse) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *WebFlagConfigApprovalRequestResponse) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *WebFlagConfigApprovalRequestResponse) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### GetDescription

`func (o *WebFlagConfigApprovalRequestResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebFlagConfigApprovalRequestResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebFlagConfigApprovalRequestResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebFlagConfigApprovalRequestResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReviewStatus

`func (o *WebFlagConfigApprovalRequestResponse) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *WebFlagConfigApprovalRequestResponse) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *WebFlagConfigApprovalRequestResponse) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.

### HasReviewStatus

`func (o *WebFlagConfigApprovalRequestResponse) HasReviewStatus() bool`

HasReviewStatus returns a boolean if a field has been set.

### GetAllReviews

`func (o *WebFlagConfigApprovalRequestResponse) GetAllReviews() []WebReviewResponse`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *WebFlagConfigApprovalRequestResponse) GetAllReviewsOk() (*[]WebReviewResponse, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *WebFlagConfigApprovalRequestResponse) SetAllReviews(v []WebReviewResponse)`

SetAllReviews sets AllReviews field to given value.

### HasAllReviews

`func (o *WebFlagConfigApprovalRequestResponse) HasAllReviews() bool`

HasAllReviews returns a boolean if a field has been set.

### GetNotifyMemberIds

`func (o *WebFlagConfigApprovalRequestResponse) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *WebFlagConfigApprovalRequestResponse) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *WebFlagConfigApprovalRequestResponse) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *WebFlagConfigApprovalRequestResponse) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetAppliedDate

`func (o *WebFlagConfigApprovalRequestResponse) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *WebFlagConfigApprovalRequestResponse) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *WebFlagConfigApprovalRequestResponse) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *WebFlagConfigApprovalRequestResponse) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.

### GetAppliedByMemberId

`func (o *WebFlagConfigApprovalRequestResponse) GetAppliedByMemberId() string`

GetAppliedByMemberId returns the AppliedByMemberId field if non-nil, zero value otherwise.

### GetAppliedByMemberIdOk

`func (o *WebFlagConfigApprovalRequestResponse) GetAppliedByMemberIdOk() (*string, bool)`

GetAppliedByMemberIdOk returns a tuple with the AppliedByMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByMemberId

`func (o *WebFlagConfigApprovalRequestResponse) SetAppliedByMemberId(v string)`

SetAppliedByMemberId sets AppliedByMemberId field to given value.

### HasAppliedByMemberId

`func (o *WebFlagConfigApprovalRequestResponse) HasAppliedByMemberId() bool`

HasAppliedByMemberId returns a boolean if a field has been set.

### GetStatus

`func (o *WebFlagConfigApprovalRequestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebFlagConfigApprovalRequestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebFlagConfigApprovalRequestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebFlagConfigApprovalRequestResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInstructions

`func (o *WebFlagConfigApprovalRequestResponse) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *WebFlagConfigApprovalRequestResponse) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *WebFlagConfigApprovalRequestResponse) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *WebFlagConfigApprovalRequestResponse) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetConflicts

`func (o *WebFlagConfigApprovalRequestResponse) GetConflicts() []WebConflict`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *WebFlagConfigApprovalRequestResponse) GetConflictsOk() (*[]WebConflict, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *WebFlagConfigApprovalRequestResponse) SetConflicts(v []WebConflict)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *WebFlagConfigApprovalRequestResponse) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetLinks

`func (o *WebFlagConfigApprovalRequestResponse) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebFlagConfigApprovalRequestResponse) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebFlagConfigApprovalRequestResponse) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebFlagConfigApprovalRequestResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetExecutionDate

`func (o *WebFlagConfigApprovalRequestResponse) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *WebFlagConfigApprovalRequestResponse) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *WebFlagConfigApprovalRequestResponse) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *WebFlagConfigApprovalRequestResponse) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetOperatingOnId

`func (o *WebFlagConfigApprovalRequestResponse) GetOperatingOnId() string`

GetOperatingOnId returns the OperatingOnId field if non-nil, zero value otherwise.

### GetOperatingOnIdOk

`func (o *WebFlagConfigApprovalRequestResponse) GetOperatingOnIdOk() (*string, bool)`

GetOperatingOnIdOk returns a tuple with the OperatingOnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingOnId

`func (o *WebFlagConfigApprovalRequestResponse) SetOperatingOnId(v string)`

SetOperatingOnId sets OperatingOnId field to given value.

### HasOperatingOnId

`func (o *WebFlagConfigApprovalRequestResponse) HasOperatingOnId() bool`

HasOperatingOnId returns a boolean if a field has been set.

### GetIntegrationMetadata

`func (o *WebFlagConfigApprovalRequestResponse) GetIntegrationMetadata() WebIntegrationMetadata`

GetIntegrationMetadata returns the IntegrationMetadata field if non-nil, zero value otherwise.

### GetIntegrationMetadataOk

`func (o *WebFlagConfigApprovalRequestResponse) GetIntegrationMetadataOk() (*WebIntegrationMetadata, bool)`

GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationMetadata

`func (o *WebFlagConfigApprovalRequestResponse) SetIntegrationMetadata(v WebIntegrationMetadata)`

SetIntegrationMetadata sets IntegrationMetadata field to given value.

### HasIntegrationMetadata

`func (o *WebFlagConfigApprovalRequestResponse) HasIntegrationMetadata() bool`

HasIntegrationMetadata returns a boolean if a field has been set.

### GetSource

`func (o *WebFlagConfigApprovalRequestResponse) GetSource() WebCopiedFromEnv`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebFlagConfigApprovalRequestResponse) GetSourceOk() (*WebCopiedFromEnv, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebFlagConfigApprovalRequestResponse) SetSource(v WebCopiedFromEnv)`

SetSource sets Source field to given value.

### HasSource

`func (o *WebFlagConfigApprovalRequestResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


