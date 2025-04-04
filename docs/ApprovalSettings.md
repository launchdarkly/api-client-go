# ApprovalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | **bool** | If approvals are required for this environment | 
**BypassApprovalsForPendingChanges** | **bool** | Whether to skip approvals for pending changes | 
**MinNumApprovals** | **int32** | Sets the amount of approvals required before a member can apply a change. The minimum is one and the maximum is five. | 
**CanReviewOwnRequest** | **bool** | Allow someone who makes an approval request to apply their own change | 
**CanApplyDeclinedChanges** | **bool** | Allow applying the change as long as at least one person has approved | 
**AutoApplyApprovedChanges** | Pointer to **bool** | Automatically apply changes that have been approved by all reviewers. This field is only applicable for approval services other than LaunchDarkly. | [optional] 
**ServiceKind** | **string** | Which service to use for managing approvals | 
**ServiceConfig** | **map[string]interface{}** |  | 
**RequiredApprovalTags** | **[]string** | Require approval only on flags with the provided tags. Otherwise all flags will require approval. | 
**ServiceKindConfigurationId** | Pointer to **string** | Optional field for integration configuration ID of a custom approval integration. This is an Enterprise-only feature. | [optional] 
**ResourceKind** | Pointer to **string** | The kind of resource for which the approval settings apply, for example, flag or segment | [optional] 

## Methods

### NewApprovalSettings

`func NewApprovalSettings(required bool, bypassApprovalsForPendingChanges bool, minNumApprovals int32, canReviewOwnRequest bool, canApplyDeclinedChanges bool, serviceKind string, serviceConfig map[string]interface{}, requiredApprovalTags []string, ) *ApprovalSettings`

NewApprovalSettings instantiates a new ApprovalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalSettingsWithDefaults

`func NewApprovalSettingsWithDefaults() *ApprovalSettings`

NewApprovalSettingsWithDefaults instantiates a new ApprovalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *ApprovalSettings) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalSettings) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalSettings) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetBypassApprovalsForPendingChanges

`func (o *ApprovalSettings) GetBypassApprovalsForPendingChanges() bool`

GetBypassApprovalsForPendingChanges returns the BypassApprovalsForPendingChanges field if non-nil, zero value otherwise.

### GetBypassApprovalsForPendingChangesOk

`func (o *ApprovalSettings) GetBypassApprovalsForPendingChangesOk() (*bool, bool)`

GetBypassApprovalsForPendingChangesOk returns a tuple with the BypassApprovalsForPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApprovalsForPendingChanges

`func (o *ApprovalSettings) SetBypassApprovalsForPendingChanges(v bool)`

SetBypassApprovalsForPendingChanges sets BypassApprovalsForPendingChanges field to given value.


### GetMinNumApprovals

`func (o *ApprovalSettings) GetMinNumApprovals() int32`

GetMinNumApprovals returns the MinNumApprovals field if non-nil, zero value otherwise.

### GetMinNumApprovalsOk

`func (o *ApprovalSettings) GetMinNumApprovalsOk() (*int32, bool)`

GetMinNumApprovalsOk returns a tuple with the MinNumApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumApprovals

`func (o *ApprovalSettings) SetMinNumApprovals(v int32)`

SetMinNumApprovals sets MinNumApprovals field to given value.


### GetCanReviewOwnRequest

`func (o *ApprovalSettings) GetCanReviewOwnRequest() bool`

GetCanReviewOwnRequest returns the CanReviewOwnRequest field if non-nil, zero value otherwise.

### GetCanReviewOwnRequestOk

`func (o *ApprovalSettings) GetCanReviewOwnRequestOk() (*bool, bool)`

GetCanReviewOwnRequestOk returns a tuple with the CanReviewOwnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReviewOwnRequest

`func (o *ApprovalSettings) SetCanReviewOwnRequest(v bool)`

SetCanReviewOwnRequest sets CanReviewOwnRequest field to given value.


### GetCanApplyDeclinedChanges

`func (o *ApprovalSettings) GetCanApplyDeclinedChanges() bool`

GetCanApplyDeclinedChanges returns the CanApplyDeclinedChanges field if non-nil, zero value otherwise.

### GetCanApplyDeclinedChangesOk

`func (o *ApprovalSettings) GetCanApplyDeclinedChangesOk() (*bool, bool)`

GetCanApplyDeclinedChangesOk returns a tuple with the CanApplyDeclinedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyDeclinedChanges

`func (o *ApprovalSettings) SetCanApplyDeclinedChanges(v bool)`

SetCanApplyDeclinedChanges sets CanApplyDeclinedChanges field to given value.


### GetAutoApplyApprovedChanges

`func (o *ApprovalSettings) GetAutoApplyApprovedChanges() bool`

GetAutoApplyApprovedChanges returns the AutoApplyApprovedChanges field if non-nil, zero value otherwise.

### GetAutoApplyApprovedChangesOk

`func (o *ApprovalSettings) GetAutoApplyApprovedChangesOk() (*bool, bool)`

GetAutoApplyApprovedChangesOk returns a tuple with the AutoApplyApprovedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApplyApprovedChanges

`func (o *ApprovalSettings) SetAutoApplyApprovedChanges(v bool)`

SetAutoApplyApprovedChanges sets AutoApplyApprovedChanges field to given value.

### HasAutoApplyApprovedChanges

`func (o *ApprovalSettings) HasAutoApplyApprovedChanges() bool`

HasAutoApplyApprovedChanges returns a boolean if a field has been set.

### GetServiceKind

`func (o *ApprovalSettings) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *ApprovalSettings) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *ApprovalSettings) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.


### GetServiceConfig

`func (o *ApprovalSettings) GetServiceConfig() map[string]interface{}`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *ApprovalSettings) GetServiceConfigOk() (*map[string]interface{}, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *ApprovalSettings) SetServiceConfig(v map[string]interface{})`

SetServiceConfig sets ServiceConfig field to given value.


### GetRequiredApprovalTags

`func (o *ApprovalSettings) GetRequiredApprovalTags() []string`

GetRequiredApprovalTags returns the RequiredApprovalTags field if non-nil, zero value otherwise.

### GetRequiredApprovalTagsOk

`func (o *ApprovalSettings) GetRequiredApprovalTagsOk() (*[]string, bool)`

GetRequiredApprovalTagsOk returns a tuple with the RequiredApprovalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovalTags

`func (o *ApprovalSettings) SetRequiredApprovalTags(v []string)`

SetRequiredApprovalTags sets RequiredApprovalTags field to given value.


### GetServiceKindConfigurationId

`func (o *ApprovalSettings) GetServiceKindConfigurationId() string`

GetServiceKindConfigurationId returns the ServiceKindConfigurationId field if non-nil, zero value otherwise.

### GetServiceKindConfigurationIdOk

`func (o *ApprovalSettings) GetServiceKindConfigurationIdOk() (*string, bool)`

GetServiceKindConfigurationIdOk returns a tuple with the ServiceKindConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKindConfigurationId

`func (o *ApprovalSettings) SetServiceKindConfigurationId(v string)`

SetServiceKindConfigurationId sets ServiceKindConfigurationId field to given value.

### HasServiceKindConfigurationId

`func (o *ApprovalSettings) HasServiceKindConfigurationId() bool`

HasServiceKindConfigurationId returns a boolean if a field has been set.

### GetResourceKind

`func (o *ApprovalSettings) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *ApprovalSettings) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *ApprovalSettings) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *ApprovalSettings) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


