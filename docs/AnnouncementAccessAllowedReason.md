# AnnouncementAccessAllowedReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]string** | Resource specifier strings | [optional] 
**NotResources** | Pointer to **[]string** | Targeted resources are the resources NOT in this list. The &lt;code&gt;resources&lt;/code&gt; and &lt;code&gt;notActions&lt;/code&gt; fields must be empty to use this field. | [optional] 
**Actions** | Pointer to **[]string** | Actions to perform on a resource | [optional] 
**NotActions** | Pointer to **[]string** | Targeted actions are the actions NOT in this list. The &lt;code&gt;actions&lt;/code&gt; and &lt;code&gt;notResources&lt;/code&gt; fields must be empty to use this field. | [optional] 
**Effect** | **string** | Whether this statement should allow or deny actions on the resources. | 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewAnnouncementAccessAllowedReason

`func NewAnnouncementAccessAllowedReason(effect string, ) *AnnouncementAccessAllowedReason`

NewAnnouncementAccessAllowedReason instantiates a new AnnouncementAccessAllowedReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementAccessAllowedReasonWithDefaults

`func NewAnnouncementAccessAllowedReasonWithDefaults() *AnnouncementAccessAllowedReason`

NewAnnouncementAccessAllowedReasonWithDefaults instantiates a new AnnouncementAccessAllowedReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *AnnouncementAccessAllowedReason) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AnnouncementAccessAllowedReason) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AnnouncementAccessAllowedReason) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AnnouncementAccessAllowedReason) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetNotResources

`func (o *AnnouncementAccessAllowedReason) GetNotResources() []string`

GetNotResources returns the NotResources field if non-nil, zero value otherwise.

### GetNotResourcesOk

`func (o *AnnouncementAccessAllowedReason) GetNotResourcesOk() (*[]string, bool)`

GetNotResourcesOk returns a tuple with the NotResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotResources

`func (o *AnnouncementAccessAllowedReason) SetNotResources(v []string)`

SetNotResources sets NotResources field to given value.

### HasNotResources

`func (o *AnnouncementAccessAllowedReason) HasNotResources() bool`

HasNotResources returns a boolean if a field has been set.

### GetActions

`func (o *AnnouncementAccessAllowedReason) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AnnouncementAccessAllowedReason) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AnnouncementAccessAllowedReason) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AnnouncementAccessAllowedReason) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetNotActions

`func (o *AnnouncementAccessAllowedReason) GetNotActions() []string`

GetNotActions returns the NotActions field if non-nil, zero value otherwise.

### GetNotActionsOk

`func (o *AnnouncementAccessAllowedReason) GetNotActionsOk() (*[]string, bool)`

GetNotActionsOk returns a tuple with the NotActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotActions

`func (o *AnnouncementAccessAllowedReason) SetNotActions(v []string)`

SetNotActions sets NotActions field to given value.

### HasNotActions

`func (o *AnnouncementAccessAllowedReason) HasNotActions() bool`

HasNotActions returns a boolean if a field has been set.

### GetEffect

`func (o *AnnouncementAccessAllowedReason) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AnnouncementAccessAllowedReason) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AnnouncementAccessAllowedReason) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetRoleName

`func (o *AnnouncementAccessAllowedReason) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AnnouncementAccessAllowedReason) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AnnouncementAccessAllowedReason) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AnnouncementAccessAllowedReason) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


