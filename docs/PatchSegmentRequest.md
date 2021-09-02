# PatchSegmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional description of changes | [optional] 
**Instructions** | [**[]PatchSegmentInstruction**](PatchSegmentInstruction.md) | Semantic patch instructions for the desired changes to the resource | 

## Methods

### NewPatchSegmentRequest

`func NewPatchSegmentRequest(instructions []PatchSegmentInstruction, ) *PatchSegmentRequest`

NewPatchSegmentRequest instantiates a new PatchSegmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSegmentRequestWithDefaults

`func NewPatchSegmentRequestWithDefaults() *PatchSegmentRequest`

NewPatchSegmentRequestWithDefaults instantiates a new PatchSegmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PatchSegmentRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchSegmentRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchSegmentRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchSegmentRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *PatchSegmentRequest) GetInstructions() []PatchSegmentInstruction`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PatchSegmentRequest) GetInstructionsOk() (*[]PatchSegmentInstruction, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PatchSegmentRequest) SetInstructions(v []PatchSegmentInstruction)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


