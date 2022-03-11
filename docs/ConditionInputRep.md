# ConditionInputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleKind** | Pointer to **string** |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**WaitDuration** | Pointer to **int32** |  | [optional] 
**WaitDurationUnit** | Pointer to **string** |  | [optional] 
**ExecuteNow** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NotifyMemberIds** | Pointer to **[]string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewConditionInputRep

`func NewConditionInputRep() *ConditionInputRep`

NewConditionInputRep instantiates a new ConditionInputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionInputRepWithDefaults

`func NewConditionInputRepWithDefaults() *ConditionInputRep`

NewConditionInputRepWithDefaults instantiates a new ConditionInputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleKind

`func (o *ConditionInputRep) GetScheduleKind() string`

GetScheduleKind returns the ScheduleKind field if non-nil, zero value otherwise.

### GetScheduleKindOk

`func (o *ConditionInputRep) GetScheduleKindOk() (*string, bool)`

GetScheduleKindOk returns a tuple with the ScheduleKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleKind

`func (o *ConditionInputRep) SetScheduleKind(v string)`

SetScheduleKind sets ScheduleKind field to given value.

### HasScheduleKind

`func (o *ConditionInputRep) HasScheduleKind() bool`

HasScheduleKind returns a boolean if a field has been set.

### GetExecutionDate

`func (o *ConditionInputRep) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ConditionInputRep) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ConditionInputRep) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ConditionInputRep) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetWaitDuration

`func (o *ConditionInputRep) GetWaitDuration() int32`

GetWaitDuration returns the WaitDuration field if non-nil, zero value otherwise.

### GetWaitDurationOk

`func (o *ConditionInputRep) GetWaitDurationOk() (*int32, bool)`

GetWaitDurationOk returns a tuple with the WaitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDuration

`func (o *ConditionInputRep) SetWaitDuration(v int32)`

SetWaitDuration sets WaitDuration field to given value.

### HasWaitDuration

`func (o *ConditionInputRep) HasWaitDuration() bool`

HasWaitDuration returns a boolean if a field has been set.

### GetWaitDurationUnit

`func (o *ConditionInputRep) GetWaitDurationUnit() string`

GetWaitDurationUnit returns the WaitDurationUnit field if non-nil, zero value otherwise.

### GetWaitDurationUnitOk

`func (o *ConditionInputRep) GetWaitDurationUnitOk() (*string, bool)`

GetWaitDurationUnitOk returns a tuple with the WaitDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUnit

`func (o *ConditionInputRep) SetWaitDurationUnit(v string)`

SetWaitDurationUnit sets WaitDurationUnit field to given value.

### HasWaitDurationUnit

`func (o *ConditionInputRep) HasWaitDurationUnit() bool`

HasWaitDurationUnit returns a boolean if a field has been set.

### GetExecuteNow

`func (o *ConditionInputRep) GetExecuteNow() bool`

GetExecuteNow returns the ExecuteNow field if non-nil, zero value otherwise.

### GetExecuteNowOk

`func (o *ConditionInputRep) GetExecuteNowOk() (*bool, bool)`

GetExecuteNowOk returns a tuple with the ExecuteNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteNow

`func (o *ConditionInputRep) SetExecuteNow(v bool)`

SetExecuteNow sets ExecuteNow field to given value.

### HasExecuteNow

`func (o *ConditionInputRep) HasExecuteNow() bool`

HasExecuteNow returns a boolean if a field has been set.

### GetDescription

`func (o *ConditionInputRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConditionInputRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConditionInputRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConditionInputRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotifyMemberIds

`func (o *ConditionInputRep) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ConditionInputRep) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ConditionInputRep) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.

### HasNotifyMemberIds

`func (o *ConditionInputRep) HasNotifyMemberIds() bool`

HasNotifyMemberIds returns a boolean if a field has been set.

### GetKind

`func (o *ConditionInputRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConditionInputRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConditionInputRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConditionInputRep) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


