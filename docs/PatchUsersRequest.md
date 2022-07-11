# PatchUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional comment describing the change | [optional] 
**Instructions** | [**[]InstructionUserRequest**](InstructionUserRequest.md) | The instructions to perform when updating | 

## Methods

### NewPatchUsersRequest

`func NewPatchUsersRequest(instructions []InstructionUserRequest, ) *PatchUsersRequest`

NewPatchUsersRequest instantiates a new PatchUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUsersRequestWithDefaults

`func NewPatchUsersRequestWithDefaults() *PatchUsersRequest`

NewPatchUsersRequestWithDefaults instantiates a new PatchUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PatchUsersRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchUsersRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchUsersRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchUsersRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *PatchUsersRequest) GetInstructions() []InstructionUserRequest`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PatchUsersRequest) GetInstructionsOk() (*[]InstructionUserRequest, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PatchUsersRequest) SetInstructions(v []InstructionUserRequest)`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


