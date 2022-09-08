# ApprovalConditionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**NotifyMemberIds** | Pointer to **[]string** |  | [optional] 
**NotifyTeamKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApprovalConditionInput

`func NewApprovalConditionInput() *ApprovalConditionInput`

NewApprovalConditionInput instantiates a new ApprovalConditionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConditionInputWithDefaults

`func NewApprovalConditionInputWithDefaults() *ApprovalConditionInput`

NewApprovalConditionInputWithDefaults instantiates a new ApprovalConditionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApprovalConditionInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalConditionInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalConditionInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalConditionInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifyMemberIds

`func (o *ApprovalConditionInput) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ApprovalConditionInput) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ApprovalConditionInput) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *ApprovalConditionInput) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetNotifyTeamKeys

`func (o *ApprovalConditionInput) GetNotifyTeamKeys() []string`

GetNotifyTeamKeys returns the NotifyTeamKeys field if non-nil, zero value otherwise.

### GetNotifyTeamKeysOk

`func (o *ApprovalConditionInput) GetNotifyTeamKeysOk() (*[]string, bool)`

GetNotifyTeamKeysOk returns a tuple with the NotifyTeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTeamKeys

`func (o *ApprovalConditionInput) SetNotifyTeamKeys(v []string)`

SetNotifyTeamKeys sets NotifyTeamKeys field to given value.

### HasNotifyTeamKeys

`func (o *ApprovalConditionInput) HasNotifyTeamKeys() bool`

HasNotifyTeamKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


