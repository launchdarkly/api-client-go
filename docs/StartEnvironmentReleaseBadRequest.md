# StartEnvironmentReleaseBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**FailingChecks** | Pointer to [**[]ReleaseHealthCheckFailure**](ReleaseHealthCheckFailure.md) | Present when &#x60;code&#x60; is &#x60;health_checks_failed&#x60;. | [optional] 

## Methods

### NewStartEnvironmentReleaseBadRequest

`func NewStartEnvironmentReleaseBadRequest(code string, message string, ) *StartEnvironmentReleaseBadRequest`

NewStartEnvironmentReleaseBadRequest instantiates a new StartEnvironmentReleaseBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartEnvironmentReleaseBadRequestWithDefaults

`func NewStartEnvironmentReleaseBadRequestWithDefaults() *StartEnvironmentReleaseBadRequest`

NewStartEnvironmentReleaseBadRequestWithDefaults instantiates a new StartEnvironmentReleaseBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StartEnvironmentReleaseBadRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StartEnvironmentReleaseBadRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StartEnvironmentReleaseBadRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *StartEnvironmentReleaseBadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StartEnvironmentReleaseBadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StartEnvironmentReleaseBadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFailingChecks

`func (o *StartEnvironmentReleaseBadRequest) GetFailingChecks() []ReleaseHealthCheckFailure`

GetFailingChecks returns the FailingChecks field if non-nil, zero value otherwise.

### GetFailingChecksOk

`func (o *StartEnvironmentReleaseBadRequest) GetFailingChecksOk() (*[]ReleaseHealthCheckFailure, bool)`

GetFailingChecksOk returns a tuple with the FailingChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingChecks

`func (o *StartEnvironmentReleaseBadRequest) SetFailingChecks(v []ReleaseHealthCheckFailure)`

SetFailingChecks sets FailingChecks field to given value.

### HasFailingChecks

`func (o *StartEnvironmentReleaseBadRequest) HasFailingChecks() bool`

HasFailingChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


