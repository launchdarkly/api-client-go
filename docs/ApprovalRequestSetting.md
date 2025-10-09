# ApprovalRequestSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | **bool** | If approvals are required for this environment | 
**BypassApprovalsForPendingChanges** | **bool** | Whether to skip approvals for pending changes | 
**MinNumApprovals** | **int32** | Sets the amount of approvals required before a member can apply a change. The minimum is one and the maximum is five.  | 
**CanReviewOwnRequest** | **bool** | Allow someone who makes an approval request to apply their own change | 
**CanApplyDeclinedChanges** | **bool** | Allow applying the change as long as at least one person has approved | 
**AutoApplyApprovedChanges** | Pointer to **NullableBool** | Automatically apply changes that have been approved by all reviewers. This field is only applicable for approval services other than LaunchDarkly.  | [optional] 
**ServiceKind** | **string** | Which service to use for managing approvals | 
**ServiceConfig** | **map[string]interface{}** | Arbitrary service-specific configuration | 
**RequiredApprovalTags** | **[]string** | Require approval only on flags with the provided tags. Otherwise all flags will require approval.  | 
**ServiceKindConfigurationId** | Pointer to **NullableString** | Optional integration configuration ID of a custom approval integration. This is an Enterprise-only feature.  | [optional] 

## Methods

### NewApprovalRequestSetting

`func NewApprovalRequestSetting(required bool, bypassApprovalsForPendingChanges bool, minNumApprovals int32, canReviewOwnRequest bool, canApplyDeclinedChanges bool, serviceKind string, serviceConfig map[string]interface{}, requiredApprovalTags []string, ) *ApprovalRequestSetting`

NewApprovalRequestSetting instantiates a new ApprovalRequestSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestSettingWithDefaults

`func NewApprovalRequestSettingWithDefaults() *ApprovalRequestSetting`

NewApprovalRequestSettingWithDefaults instantiates a new ApprovalRequestSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *ApprovalRequestSetting) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalRequestSetting) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalRequestSetting) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetBypassApprovalsForPendingChanges

`func (o *ApprovalRequestSetting) GetBypassApprovalsForPendingChanges() bool`

GetBypassApprovalsForPendingChanges returns the BypassApprovalsForPendingChanges field if non-nil, zero value otherwise.

### GetBypassApprovalsForPendingChangesOk

`func (o *ApprovalRequestSetting) GetBypassApprovalsForPendingChangesOk() (*bool, bool)`

GetBypassApprovalsForPendingChangesOk returns a tuple with the BypassApprovalsForPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApprovalsForPendingChanges

`func (o *ApprovalRequestSetting) SetBypassApprovalsForPendingChanges(v bool)`

SetBypassApprovalsForPendingChanges sets BypassApprovalsForPendingChanges field to given value.


### GetMinNumApprovals

`func (o *ApprovalRequestSetting) GetMinNumApprovals() int32`

GetMinNumApprovals returns the MinNumApprovals field if non-nil, zero value otherwise.

### GetMinNumApprovalsOk

`func (o *ApprovalRequestSetting) GetMinNumApprovalsOk() (*int32, bool)`

GetMinNumApprovalsOk returns a tuple with the MinNumApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumApprovals

`func (o *ApprovalRequestSetting) SetMinNumApprovals(v int32)`

SetMinNumApprovals sets MinNumApprovals field to given value.


### GetCanReviewOwnRequest

`func (o *ApprovalRequestSetting) GetCanReviewOwnRequest() bool`

GetCanReviewOwnRequest returns the CanReviewOwnRequest field if non-nil, zero value otherwise.

### GetCanReviewOwnRequestOk

`func (o *ApprovalRequestSetting) GetCanReviewOwnRequestOk() (*bool, bool)`

GetCanReviewOwnRequestOk returns a tuple with the CanReviewOwnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReviewOwnRequest

`func (o *ApprovalRequestSetting) SetCanReviewOwnRequest(v bool)`

SetCanReviewOwnRequest sets CanReviewOwnRequest field to given value.


### GetCanApplyDeclinedChanges

`func (o *ApprovalRequestSetting) GetCanApplyDeclinedChanges() bool`

GetCanApplyDeclinedChanges returns the CanApplyDeclinedChanges field if non-nil, zero value otherwise.

### GetCanApplyDeclinedChangesOk

`func (o *ApprovalRequestSetting) GetCanApplyDeclinedChangesOk() (*bool, bool)`

GetCanApplyDeclinedChangesOk returns a tuple with the CanApplyDeclinedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyDeclinedChanges

`func (o *ApprovalRequestSetting) SetCanApplyDeclinedChanges(v bool)`

SetCanApplyDeclinedChanges sets CanApplyDeclinedChanges field to given value.


### GetAutoApplyApprovedChanges

`func (o *ApprovalRequestSetting) GetAutoApplyApprovedChanges() bool`

GetAutoApplyApprovedChanges returns the AutoApplyApprovedChanges field if non-nil, zero value otherwise.

### GetAutoApplyApprovedChangesOk

`func (o *ApprovalRequestSetting) GetAutoApplyApprovedChangesOk() (*bool, bool)`

GetAutoApplyApprovedChangesOk returns a tuple with the AutoApplyApprovedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApplyApprovedChanges

`func (o *ApprovalRequestSetting) SetAutoApplyApprovedChanges(v bool)`

SetAutoApplyApprovedChanges sets AutoApplyApprovedChanges field to given value.

### HasAutoApplyApprovedChanges

`func (o *ApprovalRequestSetting) HasAutoApplyApprovedChanges() bool`

HasAutoApplyApprovedChanges returns a boolean if a field has been set.

### SetAutoApplyApprovedChangesNil

`func (o *ApprovalRequestSetting) SetAutoApplyApprovedChangesNil(b bool)`

 SetAutoApplyApprovedChangesNil sets the value for AutoApplyApprovedChanges to be an explicit nil

### UnsetAutoApplyApprovedChanges
`func (o *ApprovalRequestSetting) UnsetAutoApplyApprovedChanges()`

UnsetAutoApplyApprovedChanges ensures that no value is present for AutoApplyApprovedChanges, not even an explicit nil
### GetServiceKind

`func (o *ApprovalRequestSetting) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *ApprovalRequestSetting) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *ApprovalRequestSetting) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.


### GetServiceConfig

`func (o *ApprovalRequestSetting) GetServiceConfig() map[string]interface{}`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *ApprovalRequestSetting) GetServiceConfigOk() (*map[string]interface{}, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *ApprovalRequestSetting) SetServiceConfig(v map[string]interface{})`

SetServiceConfig sets ServiceConfig field to given value.


### GetRequiredApprovalTags

`func (o *ApprovalRequestSetting) GetRequiredApprovalTags() []string`

GetRequiredApprovalTags returns the RequiredApprovalTags field if non-nil, zero value otherwise.

### GetRequiredApprovalTagsOk

`func (o *ApprovalRequestSetting) GetRequiredApprovalTagsOk() (*[]string, bool)`

GetRequiredApprovalTagsOk returns a tuple with the RequiredApprovalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovalTags

`func (o *ApprovalRequestSetting) SetRequiredApprovalTags(v []string)`

SetRequiredApprovalTags sets RequiredApprovalTags field to given value.


### GetServiceKindConfigurationId

`func (o *ApprovalRequestSetting) GetServiceKindConfigurationId() string`

GetServiceKindConfigurationId returns the ServiceKindConfigurationId field if non-nil, zero value otherwise.

### GetServiceKindConfigurationIdOk

`func (o *ApprovalRequestSetting) GetServiceKindConfigurationIdOk() (*string, bool)`

GetServiceKindConfigurationIdOk returns a tuple with the ServiceKindConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKindConfigurationId

`func (o *ApprovalRequestSetting) SetServiceKindConfigurationId(v string)`

SetServiceKindConfigurationId sets ServiceKindConfigurationId field to given value.

### HasServiceKindConfigurationId

`func (o *ApprovalRequestSetting) HasServiceKindConfigurationId() bool`

HasServiceKindConfigurationId returns a boolean if a field has been set.

### SetServiceKindConfigurationIdNil

`func (o *ApprovalRequestSetting) SetServiceKindConfigurationIdNil(b bool)`

 SetServiceKindConfigurationIdNil sets the value for ServiceKindConfigurationId to be an explicit nil

### UnsetServiceKindConfigurationId
`func (o *ApprovalRequestSetting) UnsetServiceKindConfigurationId()`

UnsetServiceKindConfigurationId ensures that no value is present for ServiceKindConfigurationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


