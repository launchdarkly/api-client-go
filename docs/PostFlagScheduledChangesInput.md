# PostFlagScheduledChangesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ExecutionDate** | **int64** |  | 
**Instructions** | **[]map[string]interface{}** |  | 

## Methods

### NewPostFlagScheduledChangesInput

`func NewPostFlagScheduledChangesInput(executionDate int64, instructions []map[string]interface{}, ) *PostFlagScheduledChangesInput`

NewPostFlagScheduledChangesInput instantiates a new PostFlagScheduledChangesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFlagScheduledChangesInputWithDefaults

`func NewPostFlagScheduledChangesInputWithDefaults() *PostFlagScheduledChangesInput`

NewPostFlagScheduledChangesInputWithDefaults instantiates a new PostFlagScheduledChangesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PostFlagScheduledChangesInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostFlagScheduledChangesInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostFlagScheduledChangesInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostFlagScheduledChangesInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExecutionDate

`func (o *PostFlagScheduledChangesInput) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *PostFlagScheduledChangesInput) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *PostFlagScheduledChangesInput) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.


### GetInstructions

`func (o *PostFlagScheduledChangesInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PostFlagScheduledChangesInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PostFlagScheduledChangesInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


