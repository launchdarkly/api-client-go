# Capabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to [**Approval**](Approval.md) |  | [optional] 
**AuditLogEventsHook** | Pointer to [**AuditLogEventsHook**](AuditLogEventsHook.md) |  | [optional] 
**BigSegmentStore** | Pointer to [**BigSegmentStore**](BigSegmentStore.md) |  | [optional] 
**ExternalConfigurationPages** | Pointer to [**ExternalConfigurationPages**](ExternalConfigurationPages.md) |  | [optional] 
**ExternalConfigurationURL** | Pointer to **string** |  | [optional] 
**FeatureStore** | Pointer to [**FeatureStore**](FeatureStore.md) |  | [optional] 
**FlagLink** | Pointer to [**FlagLink**](FlagLink.md) |  | [optional] 
**HideConfiguration** | Pointer to **bool** |  | [optional] 
**ReservedCustomProperties** | Pointer to [**[]ReservedCustomPropertiesItems**](ReservedCustomPropertiesItems.md) |  | [optional] 
**SyncedSegment** | Pointer to [**SyncedSegment**](SyncedSegment.md) |  | [optional] 
**Trigger** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 

## Methods

### NewCapabilities

`func NewCapabilities() *Capabilities`

NewCapabilities instantiates a new Capabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesWithDefaults

`func NewCapabilitiesWithDefaults() *Capabilities`

NewCapabilitiesWithDefaults instantiates a new Capabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *Capabilities) GetApproval() Approval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *Capabilities) GetApprovalOk() (*Approval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *Capabilities) SetApproval(v Approval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *Capabilities) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetAuditLogEventsHook

`func (o *Capabilities) GetAuditLogEventsHook() AuditLogEventsHook`

GetAuditLogEventsHook returns the AuditLogEventsHook field if non-nil, zero value otherwise.

### GetAuditLogEventsHookOk

`func (o *Capabilities) GetAuditLogEventsHookOk() (*AuditLogEventsHook, bool)`

GetAuditLogEventsHookOk returns a tuple with the AuditLogEventsHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogEventsHook

`func (o *Capabilities) SetAuditLogEventsHook(v AuditLogEventsHook)`

SetAuditLogEventsHook sets AuditLogEventsHook field to given value.

### HasAuditLogEventsHook

`func (o *Capabilities) HasAuditLogEventsHook() bool`

HasAuditLogEventsHook returns a boolean if a field has been set.

### GetBigSegmentStore

`func (o *Capabilities) GetBigSegmentStore() BigSegmentStore`

GetBigSegmentStore returns the BigSegmentStore field if non-nil, zero value otherwise.

### GetBigSegmentStoreOk

`func (o *Capabilities) GetBigSegmentStoreOk() (*BigSegmentStore, bool)`

GetBigSegmentStoreOk returns a tuple with the BigSegmentStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigSegmentStore

`func (o *Capabilities) SetBigSegmentStore(v BigSegmentStore)`

SetBigSegmentStore sets BigSegmentStore field to given value.

### HasBigSegmentStore

`func (o *Capabilities) HasBigSegmentStore() bool`

HasBigSegmentStore returns a boolean if a field has been set.

### GetExternalConfigurationPages

`func (o *Capabilities) GetExternalConfigurationPages() ExternalConfigurationPages`

GetExternalConfigurationPages returns the ExternalConfigurationPages field if non-nil, zero value otherwise.

### GetExternalConfigurationPagesOk

`func (o *Capabilities) GetExternalConfigurationPagesOk() (*ExternalConfigurationPages, bool)`

GetExternalConfigurationPagesOk returns a tuple with the ExternalConfigurationPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfigurationPages

`func (o *Capabilities) SetExternalConfigurationPages(v ExternalConfigurationPages)`

SetExternalConfigurationPages sets ExternalConfigurationPages field to given value.

### HasExternalConfigurationPages

`func (o *Capabilities) HasExternalConfigurationPages() bool`

HasExternalConfigurationPages returns a boolean if a field has been set.

### GetExternalConfigurationURL

`func (o *Capabilities) GetExternalConfigurationURL() string`

GetExternalConfigurationURL returns the ExternalConfigurationURL field if non-nil, zero value otherwise.

### GetExternalConfigurationURLOk

`func (o *Capabilities) GetExternalConfigurationURLOk() (*string, bool)`

GetExternalConfigurationURLOk returns a tuple with the ExternalConfigurationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfigurationURL

`func (o *Capabilities) SetExternalConfigurationURL(v string)`

SetExternalConfigurationURL sets ExternalConfigurationURL field to given value.

### HasExternalConfigurationURL

`func (o *Capabilities) HasExternalConfigurationURL() bool`

HasExternalConfigurationURL returns a boolean if a field has been set.

### GetFeatureStore

`func (o *Capabilities) GetFeatureStore() FeatureStore`

GetFeatureStore returns the FeatureStore field if non-nil, zero value otherwise.

### GetFeatureStoreOk

`func (o *Capabilities) GetFeatureStoreOk() (*FeatureStore, bool)`

GetFeatureStoreOk returns a tuple with the FeatureStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStore

`func (o *Capabilities) SetFeatureStore(v FeatureStore)`

SetFeatureStore sets FeatureStore field to given value.

### HasFeatureStore

`func (o *Capabilities) HasFeatureStore() bool`

HasFeatureStore returns a boolean if a field has been set.

### GetFlagLink

`func (o *Capabilities) GetFlagLink() FlagLink`

GetFlagLink returns the FlagLink field if non-nil, zero value otherwise.

### GetFlagLinkOk

`func (o *Capabilities) GetFlagLinkOk() (*FlagLink, bool)`

GetFlagLinkOk returns a tuple with the FlagLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagLink

`func (o *Capabilities) SetFlagLink(v FlagLink)`

SetFlagLink sets FlagLink field to given value.

### HasFlagLink

`func (o *Capabilities) HasFlagLink() bool`

HasFlagLink returns a boolean if a field has been set.

### GetHideConfiguration

`func (o *Capabilities) GetHideConfiguration() bool`

GetHideConfiguration returns the HideConfiguration field if non-nil, zero value otherwise.

### GetHideConfigurationOk

`func (o *Capabilities) GetHideConfigurationOk() (*bool, bool)`

GetHideConfigurationOk returns a tuple with the HideConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideConfiguration

`func (o *Capabilities) SetHideConfiguration(v bool)`

SetHideConfiguration sets HideConfiguration field to given value.

### HasHideConfiguration

`func (o *Capabilities) HasHideConfiguration() bool`

HasHideConfiguration returns a boolean if a field has been set.

### GetReservedCustomProperties

`func (o *Capabilities) GetReservedCustomProperties() []ReservedCustomPropertiesItems`

GetReservedCustomProperties returns the ReservedCustomProperties field if non-nil, zero value otherwise.

### GetReservedCustomPropertiesOk

`func (o *Capabilities) GetReservedCustomPropertiesOk() (*[]ReservedCustomPropertiesItems, bool)`

GetReservedCustomPropertiesOk returns a tuple with the ReservedCustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCustomProperties

`func (o *Capabilities) SetReservedCustomProperties(v []ReservedCustomPropertiesItems)`

SetReservedCustomProperties sets ReservedCustomProperties field to given value.

### HasReservedCustomProperties

`func (o *Capabilities) HasReservedCustomProperties() bool`

HasReservedCustomProperties returns a boolean if a field has been set.

### GetSyncedSegment

`func (o *Capabilities) GetSyncedSegment() SyncedSegment`

GetSyncedSegment returns the SyncedSegment field if non-nil, zero value otherwise.

### GetSyncedSegmentOk

`func (o *Capabilities) GetSyncedSegmentOk() (*SyncedSegment, bool)`

GetSyncedSegmentOk returns a tuple with the SyncedSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedSegment

`func (o *Capabilities) SetSyncedSegment(v SyncedSegment)`

SetSyncedSegment sets SyncedSegment field to given value.

### HasSyncedSegment

`func (o *Capabilities) HasSyncedSegment() bool`

HasSyncedSegment returns a boolean if a field has been set.

### GetTrigger

`func (o *Capabilities) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Capabilities) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Capabilities) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Capabilities) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


