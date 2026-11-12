# PostSegmentScheduledChangeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionDate** | **int64** |  | 
**Comment** | Pointer to **string** | Optional comment describing the scheduled change | [optional] 
**Instructions** | **[]map[string]interface{}** | The segment semantic patch instructions to apply on the execution date | 

## Methods

### NewPostSegmentScheduledChangeInput

`func NewPostSegmentScheduledChangeInput(executionDate int64, instructions []map[string]interface{}, ) *PostSegmentScheduledChangeInput`

NewPostSegmentScheduledChangeInput instantiates a new PostSegmentScheduledChangeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSegmentScheduledChangeInputWithDefaults

`func NewPostSegmentScheduledChangeInputWithDefaults() *PostSegmentScheduledChangeInput`

NewPostSegmentScheduledChangeInputWithDefaults instantiates a new PostSegmentScheduledChangeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionDate

`func (o *PostSegmentScheduledChangeInput) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *PostSegmentScheduledChangeInput) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *PostSegmentScheduledChangeInput) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.


### GetComment

`func (o *PostSegmentScheduledChangeInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostSegmentScheduledChangeInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostSegmentScheduledChangeInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostSegmentScheduledChangeInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *PostSegmentScheduledChangeInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PostSegmentScheduledChangeInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PostSegmentScheduledChangeInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


