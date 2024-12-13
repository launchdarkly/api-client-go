# CapabilityConfigPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**ApprovalsCapabilityConfig**](ApprovalsCapabilityConfig.md) |  | [optional] 
**AuditLogEventsHook** | Pointer to [**AuditLogEventsHookCapabilityConfigPost**](AuditLogEventsHookCapabilityConfigPost.md) |  | [optional] 

## Methods

### NewCapabilityConfigPost

`func NewCapabilityConfigPost() *CapabilityConfigPost`

NewCapabilityConfigPost instantiates a new CapabilityConfigPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityConfigPostWithDefaults

`func NewCapabilityConfigPostWithDefaults() *CapabilityConfigPost`

NewCapabilityConfigPostWithDefaults instantiates a new CapabilityConfigPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *CapabilityConfigPost) GetApprovals() ApprovalsCapabilityConfig`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *CapabilityConfigPost) GetApprovalsOk() (*ApprovalsCapabilityConfig, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *CapabilityConfigPost) SetApprovals(v ApprovalsCapabilityConfig)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *CapabilityConfigPost) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetAuditLogEventsHook

`func (o *CapabilityConfigPost) GetAuditLogEventsHook() AuditLogEventsHookCapabilityConfigPost`

GetAuditLogEventsHook returns the AuditLogEventsHook field if non-nil, zero value otherwise.

### GetAuditLogEventsHookOk

`func (o *CapabilityConfigPost) GetAuditLogEventsHookOk() (*AuditLogEventsHookCapabilityConfigPost, bool)`

GetAuditLogEventsHookOk returns a tuple with the AuditLogEventsHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogEventsHook

`func (o *CapabilityConfigPost) SetAuditLogEventsHook(v AuditLogEventsHookCapabilityConfigPost)`

SetAuditLogEventsHook sets AuditLogEventsHook field to given value.

### HasAuditLogEventsHook

`func (o *CapabilityConfigPost) HasAuditLogEventsHook() bool`

HasAuditLogEventsHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


