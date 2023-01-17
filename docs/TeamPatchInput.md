# TeamPatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional comment describing the update | [optional] 
**Instructions** | **[]map[string]interface{}** |  | 

## Methods

### NewTeamPatchInput

`func NewTeamPatchInput(instructions []map[string]interface{}, ) *TeamPatchInput`

NewTeamPatchInput instantiates a new TeamPatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPatchInputWithDefaults

`func NewTeamPatchInputWithDefaults() *TeamPatchInput`

NewTeamPatchInputWithDefaults instantiates a new TeamPatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *TeamPatchInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TeamPatchInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TeamPatchInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TeamPatchInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *TeamPatchInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TeamPatchInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TeamPatchInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


