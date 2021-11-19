# ConditionOutputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Execution** | [**ExecutionOutputRep**](ExecutionOutputRep.md) |  | 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**Description** | **string** |  | 
**NotifyMemberIds** | **[]string** |  | 
**AllReviews** | [**[]ReviewOutputRep**](ReviewOutputRep.md) |  | 
**ReviewStatus** | **string** |  | 
**AppliedDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewConditionOutputRep

`func NewConditionOutputRep(id string, execution ExecutionOutputRep, description string, notifyMemberIds []string, allReviews []ReviewOutputRep, reviewStatus string, ) *ConditionOutputRep`

NewConditionOutputRep instantiates a new ConditionOutputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionOutputRepWithDefaults

`func NewConditionOutputRepWithDefaults() *ConditionOutputRep`

NewConditionOutputRepWithDefaults instantiates a new ConditionOutputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConditionOutputRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConditionOutputRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConditionOutputRep) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ConditionOutputRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConditionOutputRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConditionOutputRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConditionOutputRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetExecution

`func (o *ConditionOutputRep) GetExecution() ExecutionOutputRep`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ConditionOutputRep) GetExecutionOk() (*ExecutionOutputRep, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ConditionOutputRep) SetExecution(v ExecutionOutputRep)`

SetExecution sets Execution field to given value.


### GetExecutionDate

`func (o *ConditionOutputRep) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ConditionOutputRep) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ConditionOutputRep) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ConditionOutputRep) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetDescription

`func (o *ConditionOutputRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConditionOutputRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConditionOutputRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifyMemberIds

`func (o *ConditionOutputRep) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ConditionOutputRep) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ConditionOutputRep) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAllReviews

`func (o *ConditionOutputRep) GetAllReviews() []ReviewOutputRep`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *ConditionOutputRep) GetAllReviewsOk() (*[]ReviewOutputRep, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *ConditionOutputRep) SetAllReviews(v []ReviewOutputRep)`

SetAllReviews sets AllReviews field to given value.


### GetReviewStatus

`func (o *ConditionOutputRep) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ConditionOutputRep) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ConditionOutputRep) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAppliedDate

`func (o *ConditionOutputRep) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *ConditionOutputRep) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *ConditionOutputRep) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *ConditionOutputRep) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


