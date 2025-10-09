# ApprovalRequestSettingWithEnvs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**map[string]ApprovalRequestSetting**](ApprovalRequestSetting.md) | Environment-specific overrides. | [optional] 
**Default** | Pointer to [**ApprovalRequestSetting**](ApprovalRequestSetting.md) |  | [optional] 
**Strict** | Pointer to [**ApprovalRequestSetting**](ApprovalRequestSetting.md) |  | [optional] 

## Methods

### NewApprovalRequestSettingWithEnvs

`func NewApprovalRequestSettingWithEnvs() *ApprovalRequestSettingWithEnvs`

NewApprovalRequestSettingWithEnvs instantiates a new ApprovalRequestSettingWithEnvs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestSettingWithEnvsWithDefaults

`func NewApprovalRequestSettingWithEnvsWithDefaults() *ApprovalRequestSettingWithEnvs`

NewApprovalRequestSettingWithEnvsWithDefaults instantiates a new ApprovalRequestSettingWithEnvs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *ApprovalRequestSettingWithEnvs) GetEnvironments() map[string]ApprovalRequestSetting`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ApprovalRequestSettingWithEnvs) GetEnvironmentsOk() (*map[string]ApprovalRequestSetting, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ApprovalRequestSettingWithEnvs) SetEnvironments(v map[string]ApprovalRequestSetting)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ApprovalRequestSettingWithEnvs) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetDefault

`func (o *ApprovalRequestSettingWithEnvs) GetDefault() ApprovalRequestSetting`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApprovalRequestSettingWithEnvs) GetDefaultOk() (*ApprovalRequestSetting, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApprovalRequestSettingWithEnvs) SetDefault(v ApprovalRequestSetting)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApprovalRequestSettingWithEnvs) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetStrict

`func (o *ApprovalRequestSettingWithEnvs) GetStrict() ApprovalRequestSetting`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *ApprovalRequestSettingWithEnvs) GetStrictOk() (*ApprovalRequestSetting, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *ApprovalRequestSettingWithEnvs) SetStrict(v ApprovalRequestSetting)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *ApprovalRequestSettingWithEnvs) HasStrict() bool`

HasStrict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


