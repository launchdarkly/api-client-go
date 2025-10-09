# ConditionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Execution** | [**ExecutionOutput**](ExecutionOutput.md) |  | 
**ScheduleKind** | Pointer to **string** |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**WaitDuration** | Pointer to **int32** |  | [optional] 
**WaitDurationUnit** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**NotifyMemberIds** | **[]string** |  | 
**AllReviews** | [**[]ReviewOutput**](ReviewOutput.md) |  | 
**ReviewStatus** | **string** |  | 
**AppliedDate** | Pointer to **int64** |  | [optional] 
**CreationConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConditionOutput

`func NewConditionOutput(id string, execution ExecutionOutput, description string, notifyMemberIds []string, allReviews []ReviewOutput, reviewStatus string, ) *ConditionOutput`

NewConditionOutput instantiates a new ConditionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionOutputWithDefaults

`func NewConditionOutputWithDefaults() *ConditionOutput`

NewConditionOutputWithDefaults instantiates a new ConditionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConditionOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConditionOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConditionOutput) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ConditionOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConditionOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConditionOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConditionOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetExecution

`func (o *ConditionOutput) GetExecution() ExecutionOutput`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ConditionOutput) GetExecutionOk() (*ExecutionOutput, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ConditionOutput) SetExecution(v ExecutionOutput)`

SetExecution sets Execution field to given value.


### GetScheduleKind

`func (o *ConditionOutput) GetScheduleKind() string`

GetScheduleKind returns the ScheduleKind field if non-nil, zero value otherwise.

### GetScheduleKindOk

`func (o *ConditionOutput) GetScheduleKindOk() (*string, bool)`

GetScheduleKindOk returns a tuple with the ScheduleKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleKind

`func (o *ConditionOutput) SetScheduleKind(v string)`

SetScheduleKind sets ScheduleKind field to given value.

### HasScheduleKind

`func (o *ConditionOutput) HasScheduleKind() bool`

HasScheduleKind returns a boolean if a field has been set.

### GetExecutionDate

`func (o *ConditionOutput) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ConditionOutput) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ConditionOutput) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ConditionOutput) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetWaitDuration

`func (o *ConditionOutput) GetWaitDuration() int32`

GetWaitDuration returns the WaitDuration field if non-nil, zero value otherwise.

### GetWaitDurationOk

`func (o *ConditionOutput) GetWaitDurationOk() (*int32, bool)`

GetWaitDurationOk returns a tuple with the WaitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDuration

`func (o *ConditionOutput) SetWaitDuration(v int32)`

SetWaitDuration sets WaitDuration field to given value.

### HasWaitDuration

`func (o *ConditionOutput) HasWaitDuration() bool`

HasWaitDuration returns a boolean if a field has been set.

### GetWaitDurationUnit

`func (o *ConditionOutput) GetWaitDurationUnit() string`

GetWaitDurationUnit returns the WaitDurationUnit field if non-nil, zero value otherwise.

### GetWaitDurationUnitOk

`func (o *ConditionOutput) GetWaitDurationUnitOk() (*string, bool)`

GetWaitDurationUnitOk returns a tuple with the WaitDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationUnit

`func (o *ConditionOutput) SetWaitDurationUnit(v string)`

SetWaitDurationUnit sets WaitDurationUnit field to given value.

### HasWaitDurationUnit

`func (o *ConditionOutput) HasWaitDurationUnit() bool`

HasWaitDurationUnit returns a boolean if a field has been set.

### GetDescription

`func (o *ConditionOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConditionOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConditionOutput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifyMemberIds

`func (o *ConditionOutput) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ConditionOutput) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ConditionOutput) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAllReviews

`func (o *ConditionOutput) GetAllReviews() []ReviewOutput`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *ConditionOutput) GetAllReviewsOk() (*[]ReviewOutput, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *ConditionOutput) SetAllReviews(v []ReviewOutput)`

SetAllReviews sets AllReviews field to given value.


### GetReviewStatus

`func (o *ConditionOutput) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ConditionOutput) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ConditionOutput) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAppliedDate

`func (o *ConditionOutput) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *ConditionOutput) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *ConditionOutput) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *ConditionOutput) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.

### GetCreationConfig

`func (o *ConditionOutput) GetCreationConfig() map[string]interface{}`

GetCreationConfig returns the CreationConfig field if non-nil, zero value otherwise.

### GetCreationConfigOk

`func (o *ConditionOutput) GetCreationConfigOk() (*map[string]interface{}, bool)`

GetCreationConfigOk returns a tuple with the CreationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationConfig

`func (o *ConditionOutput) SetCreationConfig(v map[string]interface{})`

SetCreationConfig sets CreationConfig field to given value.

### HasCreationConfig

`func (o *ConditionOutput) HasCreationConfig() bool`

HasCreationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


