# ApprovalConditionOutputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**NotifyMemberIds** | **[]string** |  | 
**AllReviews** | [**[]ReviewOutputRep**](ReviewOutputRep.md) |  | 
**ReviewStatus** | **string** |  | 
**AppliedDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewApprovalConditionOutputRep

`func NewApprovalConditionOutputRep(description string, notifyMemberIds []string, allReviews []ReviewOutputRep, reviewStatus string, ) *ApprovalConditionOutputRep`

NewApprovalConditionOutputRep instantiates a new ApprovalConditionOutputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConditionOutputRepWithDefaults

`func NewApprovalConditionOutputRepWithDefaults() *ApprovalConditionOutputRep`

NewApprovalConditionOutputRepWithDefaults instantiates a new ApprovalConditionOutputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApprovalConditionOutputRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalConditionOutputRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalConditionOutputRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifyMemberIds

`func (o *ApprovalConditionOutputRep) GetNotifyMemberIds() []string`

GetNotifyMemberIds returns the NotifyMemberIds field if non-nil, zero value otherwise.

### GetNotifyMemberIdsOk

`func (o *ApprovalConditionOutputRep) GetNotifyMemberIdsOk() (*[]string, bool)`

GetNotifyMemberIdsOk returns a tuple with the NotifyMemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyMemberIds

`func (o *ApprovalConditionOutputRep) SetNotifyMemberIds(v []string)`

SetNotifyMemberIds sets NotifyMemberIds field to given value.


### GetAllReviews

`func (o *ApprovalConditionOutputRep) GetAllReviews() []ReviewOutputRep`

GetAllReviews returns the AllReviews field if non-nil, zero value otherwise.

### GetAllReviewsOk

`func (o *ApprovalConditionOutputRep) GetAllReviewsOk() (*[]ReviewOutputRep, bool)`

GetAllReviewsOk returns a tuple with the AllReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllReviews

`func (o *ApprovalConditionOutputRep) SetAllReviews(v []ReviewOutputRep)`

SetAllReviews sets AllReviews field to given value.


### GetReviewStatus

`func (o *ApprovalConditionOutputRep) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *ApprovalConditionOutputRep) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *ApprovalConditionOutputRep) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.


### GetAppliedDate

`func (o *ApprovalConditionOutputRep) GetAppliedDate() int64`

GetAppliedDate returns the AppliedDate field if non-nil, zero value otherwise.

### GetAppliedDateOk

`func (o *ApprovalConditionOutputRep) GetAppliedDateOk() (*int64, bool)`

GetAppliedDateOk returns a tuple with the AppliedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDate

`func (o *ApprovalConditionOutputRep) SetAppliedDate(v int64)`

SetAppliedDate sets AppliedDate field to given value.

### HasAppliedDate

`func (o *ApprovalConditionOutputRep) HasAppliedDate() bool`

HasAppliedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


