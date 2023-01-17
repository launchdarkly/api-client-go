# FlagScheduledChangesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional comment describing the update to the scheduled changes | [optional] 
**Instructions** | **[]map[string]interface{}** |  | 

## Methods

### NewFlagScheduledChangesInput

`func NewFlagScheduledChangesInput(instructions []map[string]interface{}, ) *FlagScheduledChangesInput`

NewFlagScheduledChangesInput instantiates a new FlagScheduledChangesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagScheduledChangesInputWithDefaults

`func NewFlagScheduledChangesInputWithDefaults() *FlagScheduledChangesInput`

NewFlagScheduledChangesInputWithDefaults instantiates a new FlagScheduledChangesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *FlagScheduledChangesInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FlagScheduledChangesInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FlagScheduledChangesInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FlagScheduledChangesInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *FlagScheduledChangesInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *FlagScheduledChangesInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *FlagScheduledChangesInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


