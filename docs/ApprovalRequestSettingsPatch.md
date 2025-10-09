# ApprovalRequestSettingsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApplyApprovedChanges** | Pointer to **NullableBool** | Automatically apply changes that have been approved by all reviewers. This field is only applicable for approval services other than LaunchDarkly.  | [optional] 
**BypassApprovalsForPendingChanges** | Pointer to **bool** | Whether to skip approvals for pending changes | [optional] 
**CanApplyDeclinedChanges** | Pointer to **bool** | Allow applying the change as long as at least one person has approved | [optional] 
**CanReviewOwnRequest** | Pointer to **bool** | Allow someone who makes an approval request to apply their own change | [optional] 
**EnvironmentKey** | **string** |  | 
**MinNumApprovals** | Pointer to **int32** | Sets the amount of approvals required before a member can apply a change. The minimum is one and the maximum is five.  | [optional] 
**Required** | Pointer to **bool** | If approvals are required for this environment | [optional] 
**RequiredApprovalTags** | Pointer to **[]string** | Require approval only on flags with the provided tags. Otherwise all flags will require approval.  | [optional] 
**ResourceKind** | **string** |  | 
**ServiceConfig** | Pointer to **map[string]interface{}** | Arbitrary service-specific configuration | [optional] 
**ServiceKind** | Pointer to **string** | Which service to use for managing approvals | [optional] 
**ServiceKindConfigurationId** | Pointer to **NullableString** | Optional integration configuration ID of a custom approval integration. This is an Enterprise-only feature.  | [optional] 

## Methods

### NewApprovalRequestSettingsPatch

`func NewApprovalRequestSettingsPatch(environmentKey string, resourceKind string, ) *ApprovalRequestSettingsPatch`

NewApprovalRequestSettingsPatch instantiates a new ApprovalRequestSettingsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestSettingsPatchWithDefaults

`func NewApprovalRequestSettingsPatchWithDefaults() *ApprovalRequestSettingsPatch`

NewApprovalRequestSettingsPatchWithDefaults instantiates a new ApprovalRequestSettingsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApplyApprovedChanges

`func (o *ApprovalRequestSettingsPatch) GetAutoApplyApprovedChanges() bool`

GetAutoApplyApprovedChanges returns the AutoApplyApprovedChanges field if non-nil, zero value otherwise.

### GetAutoApplyApprovedChangesOk

`func (o *ApprovalRequestSettingsPatch) GetAutoApplyApprovedChangesOk() (*bool, bool)`

GetAutoApplyApprovedChangesOk returns a tuple with the AutoApplyApprovedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApplyApprovedChanges

`func (o *ApprovalRequestSettingsPatch) SetAutoApplyApprovedChanges(v bool)`

SetAutoApplyApprovedChanges sets AutoApplyApprovedChanges field to given value.

### HasAutoApplyApprovedChanges

`func (o *ApprovalRequestSettingsPatch) HasAutoApplyApprovedChanges() bool`

HasAutoApplyApprovedChanges returns a boolean if a field has been set.

### SetAutoApplyApprovedChangesNil

`func (o *ApprovalRequestSettingsPatch) SetAutoApplyApprovedChangesNil(b bool)`

 SetAutoApplyApprovedChangesNil sets the value for AutoApplyApprovedChanges to be an explicit nil

### UnsetAutoApplyApprovedChanges
`func (o *ApprovalRequestSettingsPatch) UnsetAutoApplyApprovedChanges()`

UnsetAutoApplyApprovedChanges ensures that no value is present for AutoApplyApprovedChanges, not even an explicit nil
### GetBypassApprovalsForPendingChanges

`func (o *ApprovalRequestSettingsPatch) GetBypassApprovalsForPendingChanges() bool`

GetBypassApprovalsForPendingChanges returns the BypassApprovalsForPendingChanges field if non-nil, zero value otherwise.

### GetBypassApprovalsForPendingChangesOk

`func (o *ApprovalRequestSettingsPatch) GetBypassApprovalsForPendingChangesOk() (*bool, bool)`

GetBypassApprovalsForPendingChangesOk returns a tuple with the BypassApprovalsForPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApprovalsForPendingChanges

`func (o *ApprovalRequestSettingsPatch) SetBypassApprovalsForPendingChanges(v bool)`

SetBypassApprovalsForPendingChanges sets BypassApprovalsForPendingChanges field to given value.

### HasBypassApprovalsForPendingChanges

`func (o *ApprovalRequestSettingsPatch) HasBypassApprovalsForPendingChanges() bool`

HasBypassApprovalsForPendingChanges returns a boolean if a field has been set.

### GetCanApplyDeclinedChanges

`func (o *ApprovalRequestSettingsPatch) GetCanApplyDeclinedChanges() bool`

GetCanApplyDeclinedChanges returns the CanApplyDeclinedChanges field if non-nil, zero value otherwise.

### GetCanApplyDeclinedChangesOk

`func (o *ApprovalRequestSettingsPatch) GetCanApplyDeclinedChangesOk() (*bool, bool)`

GetCanApplyDeclinedChangesOk returns a tuple with the CanApplyDeclinedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyDeclinedChanges

`func (o *ApprovalRequestSettingsPatch) SetCanApplyDeclinedChanges(v bool)`

SetCanApplyDeclinedChanges sets CanApplyDeclinedChanges field to given value.

### HasCanApplyDeclinedChanges

`func (o *ApprovalRequestSettingsPatch) HasCanApplyDeclinedChanges() bool`

HasCanApplyDeclinedChanges returns a boolean if a field has been set.

### GetCanReviewOwnRequest

`func (o *ApprovalRequestSettingsPatch) GetCanReviewOwnRequest() bool`

GetCanReviewOwnRequest returns the CanReviewOwnRequest field if non-nil, zero value otherwise.

### GetCanReviewOwnRequestOk

`func (o *ApprovalRequestSettingsPatch) GetCanReviewOwnRequestOk() (*bool, bool)`

GetCanReviewOwnRequestOk returns a tuple with the CanReviewOwnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReviewOwnRequest

`func (o *ApprovalRequestSettingsPatch) SetCanReviewOwnRequest(v bool)`

SetCanReviewOwnRequest sets CanReviewOwnRequest field to given value.

### HasCanReviewOwnRequest

`func (o *ApprovalRequestSettingsPatch) HasCanReviewOwnRequest() bool`

HasCanReviewOwnRequest returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *ApprovalRequestSettingsPatch) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *ApprovalRequestSettingsPatch) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *ApprovalRequestSettingsPatch) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetMinNumApprovals

`func (o *ApprovalRequestSettingsPatch) GetMinNumApprovals() int32`

GetMinNumApprovals returns the MinNumApprovals field if non-nil, zero value otherwise.

### GetMinNumApprovalsOk

`func (o *ApprovalRequestSettingsPatch) GetMinNumApprovalsOk() (*int32, bool)`

GetMinNumApprovalsOk returns a tuple with the MinNumApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumApprovals

`func (o *ApprovalRequestSettingsPatch) SetMinNumApprovals(v int32)`

SetMinNumApprovals sets MinNumApprovals field to given value.

### HasMinNumApprovals

`func (o *ApprovalRequestSettingsPatch) HasMinNumApprovals() bool`

HasMinNumApprovals returns a boolean if a field has been set.

### GetRequired

`func (o *ApprovalRequestSettingsPatch) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalRequestSettingsPatch) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalRequestSettingsPatch) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApprovalRequestSettingsPatch) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetRequiredApprovalTags

`func (o *ApprovalRequestSettingsPatch) GetRequiredApprovalTags() []string`

GetRequiredApprovalTags returns the RequiredApprovalTags field if non-nil, zero value otherwise.

### GetRequiredApprovalTagsOk

`func (o *ApprovalRequestSettingsPatch) GetRequiredApprovalTagsOk() (*[]string, bool)`

GetRequiredApprovalTagsOk returns a tuple with the RequiredApprovalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovalTags

`func (o *ApprovalRequestSettingsPatch) SetRequiredApprovalTags(v []string)`

SetRequiredApprovalTags sets RequiredApprovalTags field to given value.

### HasRequiredApprovalTags

`func (o *ApprovalRequestSettingsPatch) HasRequiredApprovalTags() bool`

HasRequiredApprovalTags returns a boolean if a field has been set.

### GetResourceKind

`func (o *ApprovalRequestSettingsPatch) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *ApprovalRequestSettingsPatch) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *ApprovalRequestSettingsPatch) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.


### GetServiceConfig

`func (o *ApprovalRequestSettingsPatch) GetServiceConfig() map[string]interface{}`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *ApprovalRequestSettingsPatch) GetServiceConfigOk() (*map[string]interface{}, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *ApprovalRequestSettingsPatch) SetServiceConfig(v map[string]interface{})`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *ApprovalRequestSettingsPatch) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.

### GetServiceKind

`func (o *ApprovalRequestSettingsPatch) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *ApprovalRequestSettingsPatch) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *ApprovalRequestSettingsPatch) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.

### HasServiceKind

`func (o *ApprovalRequestSettingsPatch) HasServiceKind() bool`

HasServiceKind returns a boolean if a field has been set.

### GetServiceKindConfigurationId

`func (o *ApprovalRequestSettingsPatch) GetServiceKindConfigurationId() string`

GetServiceKindConfigurationId returns the ServiceKindConfigurationId field if non-nil, zero value otherwise.

### GetServiceKindConfigurationIdOk

`func (o *ApprovalRequestSettingsPatch) GetServiceKindConfigurationIdOk() (*string, bool)`

GetServiceKindConfigurationIdOk returns a tuple with the ServiceKindConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKindConfigurationId

`func (o *ApprovalRequestSettingsPatch) SetServiceKindConfigurationId(v string)`

SetServiceKindConfigurationId sets ServiceKindConfigurationId field to given value.

### HasServiceKindConfigurationId

`func (o *ApprovalRequestSettingsPatch) HasServiceKindConfigurationId() bool`

HasServiceKindConfigurationId returns a boolean if a field has been set.

### SetServiceKindConfigurationIdNil

`func (o *ApprovalRequestSettingsPatch) SetServiceKindConfigurationIdNil(b bool)`

 SetServiceKindConfigurationIdNil sets the value for ServiceKindConfigurationId to be an explicit nil

### UnsetServiceKindConfigurationId
`func (o *ApprovalRequestSettingsPatch) UnsetServiceKindConfigurationId()`

UnsetServiceKindConfigurationId ensures that no value is present for ServiceKindConfigurationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


