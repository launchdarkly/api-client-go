# ScheduleConditionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleKind** | Pointer to **string** |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**WaitDuration** | Pointer to **int32** | For workflow stages whose scheduled execution is relative, how far in the future the stage should start. | [optional] 
**WaitDurationUnit** | Pointer to **string** |  | [optional] 
**ExecuteNow** | Pointer to **bool** | Whether the workflow stage should be executed immediately | [optional] 

## Methods

### NewScheduleConditionInput

`func NewScheduleConditionInput() *ScheduleConditionInput`

NewScheduleConditionInput instantiates a new ScheduleConditionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleConditionInputWithDefaults

`func NewScheduleConditionInputWithDefaults() *ScheduleConditionInput`

NewScheduleConditionInputWithDefaults instantiates a new ScheduleConditionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleKind

`func (o *ScheduleConditionInput) GetScheduleKind() string`

GetScheduleKind returns the ScheduleKind field if non-nil, zero value otherwise.

### GetScheduleKindOk

`func (o *ScheduleConditionInput) GetScheduleKindOk() (*string, bool)`

GetScheduleKindOk returns a tuple with the ScheduleKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleKind

`func (o *ScheduleConditionInput) SetScheduleKind(v string)`

SetScheduleKind sets ScheduleKind field to given value.

### HasScheduleKind

`func (o *ScheduleConditionInput) HasScheduleKind() bool`

HasScheduleKind returns a boolean if a field has been set.

### GetExecutionDate

`func (o *ScheduleConditionInput) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ScheduleConditionInput) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ScheduleConditionInput) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ScheduleConditionInput) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetWaitDuration

`func (o *ScheduleConditionInput) GetWaitDuration() int32`

GetWaitDuration returns the WaitDuration field if non-nil, zero value otherwise.

### GetWaitDurationOk

`func (o *ScheduleConditionInput) GetWaitDurationOk() (*int32, bool)`

GetWaitDurationOk returns a tuple with the WaitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDuration

`func (o *ScheduleConditionInput) SetWaitDuration(v int32)`

SetWaitDuration sets WaitDuration field to given value.

### HasWaitDuration

`func (o *ScheduleConditionInput) HasWaitDuration() bool`

HasWaitDuration returns a boolean if a field has been set.

### GetWaitDurationUnit

`func (o *ScheduleConditionInput) GetWaitDurationUnit() string`

GetWaitDurationUnit returns the WaitDurationUnit field if non-nil, zero value otherwise.

### GetWaitDurationUnitOk

`func (o *ScheduleConditionInput) GetWaitDurationUnitOk() (*string, bool)`

GetWaitDurationUnitOk returns a tuple with the WaitDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUnit

`func (o *ScheduleConditionInput) SetWaitDurationUnit(v string)`

SetWaitDurationUnit sets WaitDurationUnit field to given value.

### HasWaitDurationUnit

`func (o *ScheduleConditionInput) HasWaitDurationUnit() bool`

HasWaitDurationUnit returns a boolean if a field has been set.

### GetExecuteNow

`func (o *ScheduleConditionInput) GetExecuteNow() bool`

GetExecuteNow returns the ExecuteNow field if non-nil, zero value otherwise.

### GetExecuteNowOk

`func (o *ScheduleConditionInput) GetExecuteNowOk() (*bool, bool)`

GetExecuteNowOk returns a tuple with the ExecuteNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteNow

`func (o *ScheduleConditionInput) SetExecuteNow(v bool)`

SetExecuteNow sets ExecuteNow field to given value.

### HasExecuteNow

`func (o *ScheduleConditionInput) HasExecuteNow() bool`

HasExecuteNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


