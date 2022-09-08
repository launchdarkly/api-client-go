# ApprovalConditionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**NotifyMemberIds** | **[]string** |  | 
**AllReviews** | [**[]ReviewOutput**](ReviewOutput.md) |  | 
**ReviewStatus** | **string** |  | 
**AppliedDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewApprovalConditionOutput

`func NewApprovalConditionOutput(description string, notifyMemberIds []string, allReviews []ReviewOutput, reviewStatus string, ) *ApprovalConditionOutput`

NewApprovalConditionOutput instantiates a new ApprovalConditionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConditionOutputWithDefaults

`func NewApprovalConditionOutputWithDefaults() *ApprovalConditionOutput`

NewApprovalConditionOutputWithDefaults instantiates a new ApprovalConditionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApprovalConditionOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalConditionOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalConditionOutput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifyMemberIds

`func (o *ApprovalConditionOutput) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ApprovalConditionOutput) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ApprovalConditionOutput) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAllReviews

`func (o *ApprovalConditionOutput) GetAllReviews() []ReviewOutput`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *ApprovalConditionOutput) GetAllReviewsOk() (*[]ReviewOutput, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *ApprovalConditionOutput) SetAllReviews(v []ReviewOutput)`

SetAllReviews sets AllReviews field to given value.


### GetReviewStatus

`func (o *ApprovalConditionOutput) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ApprovalConditionOutput) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ApprovalConditionOutput) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAppliedDate

`func (o *ApprovalConditionOutput) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *ApprovalConditionOutput) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *ApprovalConditionOutput) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *ApprovalConditionOutput) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


