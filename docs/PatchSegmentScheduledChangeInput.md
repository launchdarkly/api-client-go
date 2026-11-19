# PatchSegmentScheduledChangeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**Comment** | Pointer to **string** | The new comment describing the scheduled change. Leave out to keep the current comment, or send an empty string to clear it. | [optional] 

## Methods

### NewPatchSegmentScheduledChangeInput

`func NewPatchSegmentScheduledChangeInput() *PatchSegmentScheduledChangeInput`

NewPatchSegmentScheduledChangeInput instantiates a new PatchSegmentScheduledChangeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSegmentScheduledChangeInputWithDefaults

`func NewPatchSegmentScheduledChangeInputWithDefaults() *PatchSegmentScheduledChangeInput`

NewPatchSegmentScheduledChangeInputWithDefaults instantiates a new PatchSegmentScheduledChangeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionDate

`func (o *PatchSegmentScheduledChangeInput) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *PatchSegmentScheduledChangeInput) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *PatchSegmentScheduledChangeInput) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *PatchSegmentScheduledChangeInput) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetComment

`func (o *PatchSegmentScheduledChangeInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchSegmentScheduledChangeInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchSegmentScheduledChangeInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchSegmentScheduledChangeInput) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


