# ApprovalSettingsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | **bool** |  | 
**BypassApprovalsForPendingChanges** | Pointer to **bool** |  | [optional] 
**MinNumApprovals** | Pointer to **int32** |  | [optional] 
**CanReviewOwnRequest** | Pointer to **bool** |  | [optional] 
**CanApplyDeclinedChanges** | Pointer to **bool** |  | [optional] 
**ServiceKind** | Pointer to **string** |  | [optional] 
**ServiceConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**RequiredApprovalTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApprovalSettingsRep

`func NewApprovalSettingsRep(required bool, ) *ApprovalSettingsRep`

NewApprovalSettingsRep instantiates a new ApprovalSettingsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalSettingsRepWithDefaults

`func NewApprovalSettingsRepWithDefaults() *ApprovalSettingsRep`

NewApprovalSettingsRepWithDefaults instantiates a new ApprovalSettingsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *ApprovalSettingsRep) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalSettingsRep) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalSettingsRep) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetBypassApprovalsForPendingChanges

`func (o *ApprovalSettingsRep) GetBypassApprovalsForPendingChanges() bool`

GetBypassApprovalsForPendingChanges returns the BypassApprovalsForPendingChanges field if non-nil, zero value otherwise.

### GetBypassApprovalsForPendingChangesOk

`func (o *ApprovalSettingsRep) GetBypassApprovalsForPendingChangesOk() (*bool, bool)`

GetBypassApprovalsForPendingChangesOk returns a tuple with the BypassApprovalsForPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApprovalsForPendingChanges

`func (o *ApprovalSettingsRep) SetBypassApprovalsForPendingChanges(v bool)`

SetBypassApprovalsForPendingChanges sets BypassApprovalsForPendingChanges field to given value.

### HasBypassApprovalsForPendingChanges

`func (o *ApprovalSettingsRep) HasBypassApprovalsForPendingChanges() bool`

HasBypassApprovalsForPendingChanges returns a boolean if a field has been set.

### GetMinNumApprovals

`func (o *ApprovalSettingsRep) GetMinNumApprovals() int32`

GetMinNumApprovals returns the MinNumApprovals field if non-nil, zero value otherwise.

### GetMinNumApprovalsOk

`func (o *ApprovalSettingsRep) GetMinNumApprovalsOk() (*int32, bool)`

GetMinNumApprovalsOk returns a tuple with the MinNumApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumApprovals

`func (o *ApprovalSettingsRep) SetMinNumApprovals(v int32)`

SetMinNumApprovals sets MinNumApprovals field to given value.

### HasMinNumApprovals

`func (o *ApprovalSettingsRep) HasMinNumApprovals() bool`

HasMinNumApprovals returns a boolean if a field has been set.

### GetCanReviewOwnRequest

`func (o *ApprovalSettingsRep) GetCanReviewOwnRequest() bool`

GetCanReviewOwnRequest returns the CanReviewOwnRequest field if non-nil, zero value otherwise.

### GetCanReviewOwnRequestOk

`func (o *ApprovalSettingsRep) GetCanReviewOwnRequestOk() (*bool, bool)`

GetCanReviewOwnRequestOk returns a tuple with the CanReviewOwnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReviewOwnRequest

`func (o *ApprovalSettingsRep) SetCanReviewOwnRequest(v bool)`

SetCanReviewOwnRequest sets CanReviewOwnRequest field to given value.

### HasCanReviewOwnRequest

`func (o *ApprovalSettingsRep) HasCanReviewOwnRequest() bool`

HasCanReviewOwnRequest returns a boolean if a field has been set.

### GetCanApplyDeclinedChanges

`func (o *ApprovalSettingsRep) GetCanApplyDeclinedChanges() bool`

GetCanApplyDeclinedChanges returns the CanApplyDeclinedChanges field if non-nil, zero value otherwise.

### GetCanApplyDeclinedChangesOk

`func (o *ApprovalSettingsRep) GetCanApplyDeclinedChangesOk() (*bool, bool)`

GetCanApplyDeclinedChangesOk returns a tuple with the CanApplyDeclinedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApplyDeclinedChanges

`func (o *ApprovalSettingsRep) SetCanApplyDeclinedChanges(v bool)`

SetCanApplyDeclinedChanges sets CanApplyDeclinedChanges field to given value.

### HasCanApplyDeclinedChanges

`func (o *ApprovalSettingsRep) HasCanApplyDeclinedChanges() bool`

HasCanApplyDeclinedChanges returns a boolean if a field has been set.

### GetServiceKind

`func (o *ApprovalSettingsRep) GetServiceKind() string`

GetServiceKind returns the ServiceKind field if non-nil, zero value otherwise.

### GetServiceKindOk

`func (o *ApprovalSettingsRep) GetServiceKindOk() (*string, bool)`

GetServiceKindOk returns a tuple with the ServiceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKind

`func (o *ApprovalSettingsRep) SetServiceKind(v string)`

SetServiceKind sets ServiceKind field to given value.

### HasServiceKind

`func (o *ApprovalSettingsRep) HasServiceKind() bool`

HasServiceKind returns a boolean if a field has been set.

### GetServiceConfig

`func (o *ApprovalSettingsRep) GetServiceConfig() map[string]interface{}`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *ApprovalSettingsRep) GetServiceConfigOk() (*map[string]interface{}, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *ApprovalSettingsRep) SetServiceConfig(v map[string]interface{})`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *ApprovalSettingsRep) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.

### GetRequiredApprovalTags

`func (o *ApprovalSettingsRep) GetRequiredApprovalTags() []string`

GetRequiredApprovalTags returns the RequiredApprovalTags field if non-nil, zero value otherwise.

### GetRequiredApprovalTagsOk

`func (o *ApprovalSettingsRep) GetRequiredApprovalTagsOk() (*[]string, bool)`

GetRequiredApprovalTagsOk returns a tuple with the RequiredApprovalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovalTags

`func (o *ApprovalSettingsRep) SetRequiredApprovalTags(v []string)`

SetRequiredApprovalTags sets RequiredApprovalTags field to given value.

### HasRequiredApprovalTags

`func (o *ApprovalSettingsRep) HasRequiredApprovalTags() bool`

HasRequiredApprovalTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


