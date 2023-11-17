# StatusServiceUnavailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specific error code encountered | 
**Message** | **string** | Description of the error | 

## Methods

### NewStatusServiceUnavailable

`func NewStatusServiceUnavailable(code string, message string, ) *StatusServiceUnavailable`

NewStatusServiceUnavailable instantiates a new StatusServiceUnavailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusServiceUnavailableWithDefaults

`func NewStatusServiceUnavailableWithDefaults() *StatusServiceUnavailable`

NewStatusServiceUnavailableWithDefaults instantiates a new StatusServiceUnavailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StatusServiceUnavailable) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatusServiceUnavailable) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatusServiceUnavailable) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *StatusServiceUnavailable) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusServiceUnavailable) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusServiceUnavailable) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


