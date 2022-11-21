# ConditionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleKind** | Pointer to **string** |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**WaitDuration** | Pointer to **int32** | For workflow stages whose scheduled execution is relative, how far in the future the stage should start. | [optional] 
**WaitDurationUnit** | Pointer to **string** |  | [optional] 
**ExecuteNow** | Pointer to **bool** | Whether the workflow stage should be executed immediately | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NotifyMemberIds** | Pointer to **[]string** |  | [optional] 
**NotifyTeamKeys** | Pointer to **[]string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewConditionInput

`func NewConditionInput() *ConditionInput`

NewConditionInput instantiates a new ConditionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionInputWithDefaults

`func NewConditionInputWithDefaults() *ConditionInput`

NewConditionInputWithDefaults instantiates a new ConditionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleKind

`func (o *ConditionInput) GetScheduleKind() string`

GetScheduleKind returns the ScheduleKind field if non-nil, zero value otherwise.

### GetScheduleKindOk

`func (o *ConditionInput) GetScheduleKindOk() (*string, bool)`

GetScheduleKindOk returns a tuple with the ScheduleKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleKind

`func (o *ConditionInput) SetScheduleKind(v string)`

SetScheduleKind sets ScheduleKind field to given value.

### HasScheduleKind

`func (o *ConditionInput) HasScheduleKind() bool`

HasScheduleKind returns a boolean if a field has been set.

### GetExecutionDate

`func (o *ConditionInput) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ConditionInput) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ConditionInput) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ConditionInput) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetWaitDuration

`func (o *ConditionInput) GetWaitDuration() int32`

GetWaitDuration returns the WaitDuration field if non-nil, zero value otherwise.

### GetWaitDurationOk

`func (o *ConditionInput) GetWaitDurationOk() (*int32, bool)`

GetWaitDurationOk returns a tuple with the WaitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDuration

`func (o *ConditionInput) SetWaitDuration(v int32)`

SetWaitDuration sets WaitDuration field to given value.

### HasWaitDuration

`func (o *ConditionInput) HasWaitDuration() bool`

HasWaitDuration returns a boolean if a field has been set.

### GetWaitDurationUnit

`func (o *ConditionInput) GetWaitDurationUnit() string`

GetWaitDurationUnit returns the WaitDurationUnit field if non-nil, zero value otherwise.

### GetWaitDurationUnitOk

`func (o *ConditionInput) GetWaitDurationUnitOk() (*string, bool)`

GetWaitDurationUnitOk returns a tuple with the WaitDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUnit

`func (o *ConditionInput) SetWaitDurationUnit(v string)`

SetWaitDurationUnit sets WaitDurationUnit field to given value.

### HasWaitDurationUnit

`func (o *ConditionInput) HasWaitDurationUnit() bool`

HasWaitDurationUnit returns a boolean if a field has been set.

### GetExecuteNow

`func (o *ConditionInput) GetExecuteNow() bool`

GetExecuteNow returns the ExecuteNow field if non-nil, zero value otherwise.

### GetExecuteNowOk

`func (o *ConditionInput) GetExecuteNowOk() (*bool, bool)`

GetExecuteNowOk returns a tuple with the ExecuteNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteNow

`func (o *ConditionInput) SetExecuteNow(v bool)`

SetExecuteNow sets ExecuteNow field to given value.

### HasExecuteNow

`func (o *ConditionInput) HasExecuteNow() bool`

HasExecuteNow returns a boolean if a field has been set.

### GetDescription

`func (o *ConditionInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConditionInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConditionInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConditionInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifyMemberIds

`func (o *ConditionInput) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ConditionInput) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ConditionInput) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *ConditionInput) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetNotifyTeamKeys

`func (o *ConditionInput) GetNotifyTeamKeys() []string`

GetNotifyTeamKeys returns the NotifyTeamKeys field if non-nil, zero value otherwise.

### GetNotifyTeamKeysOk

`func (o *ConditionInput) GetNotifyTeamKeysOk() (*[]string, bool)`

GetNotifyTeamKeysOk returns a tuple with the NotifyTeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTeamKeys

`func (o *ConditionInput) SetNotifyTeamKeys(v []string)`

SetNotifyTeamKeys sets NotifyTeamKeys field to given value.

### HasNotifyTeamKeys

`func (o *ConditionInput) HasNotifyTeamKeys() bool`

HasNotifyTeamKeys returns a boolean if a field has been set.

### GetKind

`func (o *ConditionInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConditionInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConditionInput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConditionInput) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


