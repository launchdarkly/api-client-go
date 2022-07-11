# ScheduleConditionInputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleKind** | Pointer to **string** |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**WaitDuration** | Pointer to **int32** | For workflow stages whose scheduled execution is relative, how far in the future the stage should start. | [optional] 
**WaitDurationUnit** | Pointer to **string** |  | [optional] 
**ExecuteNow** | Pointer to **bool** | Whether the workflow stage should be executed immediately | [optional] 

## Methods

### NewScheduleConditionInputRep

`func NewScheduleConditionInputRep() *ScheduleConditionInputRep`

NewScheduleConditionInputRep instantiates a new ScheduleConditionInputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleConditionInputRepWithDefaults

`func NewScheduleConditionInputRepWithDefaults() *ScheduleConditionInputRep`

NewScheduleConditionInputRepWithDefaults instantiates a new ScheduleConditionInputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleKind

`func (o *ScheduleConditionInputRep) GetScheduleKind() string`

GetScheduleKind returns the ScheduleKind field if non-nil, zero value otherwise.

### GetScheduleKindOk

`func (o *ScheduleConditionInputRep) GetScheduleKindOk() (*string, bool)`

GetScheduleKindOk returns a tuple with the ScheduleKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleKind

`func (o *ScheduleConditionInputRep) SetScheduleKind(v string)`

SetScheduleKind sets ScheduleKind field to given value.

### HasScheduleKind

`func (o *ScheduleConditionInputRep) HasScheduleKind() bool`

HasScheduleKind returns a boolean if a field has been set.

### GetExecutionDate

`func (o *ScheduleConditionInputRep) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ScheduleConditionInputRep) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ScheduleConditionInputRep) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ScheduleConditionInputRep) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetWaitDuration

`func (o *ScheduleConditionInputRep) GetWaitDuration() int32`

GetWaitDuration returns the WaitDuration field if non-nil, zero value otherwise.

### GetWaitDurationOk

`func (o *ScheduleConditionInputRep) GetWaitDurationOk() (*int32, bool)`

GetWaitDurationOk returns a tuple with the WaitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDuration

`func (o *ScheduleConditionInputRep) SetWaitDuration(v int32)`

SetWaitDuration sets WaitDuration field to given value.

### HasWaitDuration

`func (o *ScheduleConditionInputRep) HasWaitDuration() bool`

HasWaitDuration returns a boolean if a field has been set.

### GetWaitDurationUnit

`func (o *ScheduleConditionInputRep) GetWaitDurationUnit() string`

GetWaitDurationUnit returns the WaitDurationUnit field if non-nil, zero value otherwise.

### GetWaitDurationUnitOk

`func (o *ScheduleConditionInputRep) GetWaitDurationUnitOk() (*string, bool)`

GetWaitDurationUnitOk returns a tuple with the WaitDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUnit

`func (o *ScheduleConditionInputRep) SetWaitDurationUnit(v string)`

SetWaitDurationUnit sets WaitDurationUnit field to given value.

### HasWaitDurationUnit

`func (o *ScheduleConditionInputRep) HasWaitDurationUnit() bool`

HasWaitDurationUnit returns a boolean if a field has been set.

### GetExecuteNow

`func (o *ScheduleConditionInputRep) GetExecuteNow() bool`

GetExecuteNow returns the ExecuteNow field if non-nil, zero value otherwise.

### GetExecuteNowOk

`func (o *ScheduleConditionInputRep) GetExecuteNowOk() (*bool, bool)`

GetExecuteNowOk returns a tuple with the ExecuteNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteNow

`func (o *ScheduleConditionInputRep) SetExecuteNow(v bool)`

SetExecuteNow sets ExecuteNow field to given value.

### HasExecuteNow

`func (o *ScheduleConditionInputRep) HasExecuteNow() bool`

HasExecuteNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


