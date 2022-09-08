# ExpiringTargetError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstructionIndex** | **int32** | The index of the PATCH instruction where the error occurred | 
**Message** | **string** | The error message related to a failed PATCH instruction | 

## Methods

### NewExpiringTargetError

`func NewExpiringTargetError(instructionIndex int32, message string, ) *ExpiringTargetError`

NewExpiringTargetError instantiates a new ExpiringTargetError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringTargetErrorWithDefaults

`func NewExpiringTargetErrorWithDefaults() *ExpiringTargetError`

NewExpiringTargetErrorWithDefaults instantiates a new ExpiringTargetError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructionIndex

`func (o *ExpiringTargetError) GetInstructionIndex() int32`

GetInstructionIndex returns the InstructionIndex field if non-nil, zero value otherwise.

### GetInstructionIndexOk

`func (o *ExpiringTargetError) GetInstructionIndexOk() (*int32, bool)`

GetInstructionIndexOk returns a tuple with the InstructionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionIndex

`func (o *ExpiringTargetError) SetInstructionIndex(v int32)`

SetInstructionIndex sets InstructionIndex field to given value.


### GetMessage

`func (o *ExpiringTargetError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExpiringTargetError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExpiringTargetError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


