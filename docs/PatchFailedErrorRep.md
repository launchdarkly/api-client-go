# PatchFailedErrorRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Specific error code encountered | [optional] 
**Message** | Pointer to **string** | Description of the error | [optional] 

## Methods

### NewPatchFailedErrorRep

`func NewPatchFailedErrorRep() *PatchFailedErrorRep`

NewPatchFailedErrorRep instantiates a new PatchFailedErrorRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFailedErrorRepWithDefaults

`func NewPatchFailedErrorRepWithDefaults() *PatchFailedErrorRep`

NewPatchFailedErrorRepWithDefaults instantiates a new PatchFailedErrorRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PatchFailedErrorRep) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PatchFailedErrorRep) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PatchFailedErrorRep) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PatchFailedErrorRep) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *PatchFailedErrorRep) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PatchFailedErrorRep) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PatchFailedErrorRep) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PatchFailedErrorRep) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


