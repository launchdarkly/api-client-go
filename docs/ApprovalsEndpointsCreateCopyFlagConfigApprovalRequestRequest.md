# ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NotifyMemberIds** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to [**ApprovalsEndpointsSourceFlag**](ApprovalsEndpointsSourceFlag.md) |  | [optional] 
**IncludedActions** | Pointer to **[]string** |  | [optional] 
**ExcludedActions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest

`func NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest() *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest`

NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest instantiates a new ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequestWithDefaults

`func NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequestWithDefaults() *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest`

NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequestWithDefaults instantiates a new ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifyMemberIds

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetSource

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetSource() ApprovalsEndpointsSourceFlag`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetSourceOk() (*ApprovalsEndpointsSourceFlag, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) SetSource(v ApprovalsEndpointsSourceFlag)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIncludedActions

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetIncludedActions() []string`

GetIncludedActions returns the IncludedActions field if non-nil, zero value otherwise.

### GetIncludedActionsOk

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetIncludedActionsOk() (*[]string, bool)`

GetIncludedActionsOk returns a tuple with the IncludedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedActions

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) SetIncludedActions(v []string)`

SetIncludedActions sets IncludedActions field to given value.

### HasIncludedActions

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) HasIncludedActions() bool`

HasIncludedActions returns a boolean if a field has been set.

### GetExcludedActions

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetExcludedActions() []string`

GetExcludedActions returns the ExcludedActions field if non-nil, zero value otherwise.

### GetExcludedActionsOk

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) GetExcludedActionsOk() (*[]string, bool)`

GetExcludedActionsOk returns a tuple with the ExcludedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedActions

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) SetExcludedActions(v []string)`

SetExcludedActions sets ExcludedActions field to given value.

### HasExcludedActions

`func (o *ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest) HasExcludedActions() bool`

HasExcludedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


