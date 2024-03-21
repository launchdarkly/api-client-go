# ValidationFailedErrorRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specific error code encountered | 
**Message** | **string** | Description of the error | 
**Errors** | [**[]FailureReasonRep**](FailureReasonRep.md) | List of validation errors | 

## Methods

### NewValidationFailedErrorRep

`func NewValidationFailedErrorRep(code string, message string, errors []FailureReasonRep, ) *ValidationFailedErrorRep`

NewValidationFailedErrorRep instantiates a new ValidationFailedErrorRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationFailedErrorRepWithDefaults

`func NewValidationFailedErrorRepWithDefaults() *ValidationFailedErrorRep`

NewValidationFailedErrorRepWithDefaults instantiates a new ValidationFailedErrorRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ValidationFailedErrorRep) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationFailedErrorRep) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationFailedErrorRep) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ValidationFailedErrorRep) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationFailedErrorRep) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationFailedErrorRep) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ValidationFailedErrorRep) GetErrors() []FailureReasonRep`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationFailedErrorRep) GetErrorsOk() (*[]FailureReasonRep, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationFailedErrorRep) SetErrors(v []FailureReasonRep)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


