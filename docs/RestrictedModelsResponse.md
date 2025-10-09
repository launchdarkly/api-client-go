# RestrictedModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Successes** | **[]string** |  | 
**Errors** | [**[]RestrictedModelError**](RestrictedModelError.md) |  | 

## Methods

### NewRestrictedModelsResponse

`func NewRestrictedModelsResponse(successes []string, errors []RestrictedModelError, ) *RestrictedModelsResponse`

NewRestrictedModelsResponse instantiates a new RestrictedModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedModelsResponseWithDefaults

`func NewRestrictedModelsResponseWithDefaults() *RestrictedModelsResponse`

NewRestrictedModelsResponseWithDefaults instantiates a new RestrictedModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccesses

`func (o *RestrictedModelsResponse) GetSuccesses() []string`

GetSuccesses returns the Successes field if non-nil, zero value otherwise.

### GetSuccessesOk

`func (o *RestrictedModelsResponse) GetSuccessesOk() (*[]string, bool)`

GetSuccessesOk returns a tuple with the Successes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesses

`func (o *RestrictedModelsResponse) SetSuccesses(v []string)`

SetSuccesses sets Successes field to given value.


### GetErrors

`func (o *RestrictedModelsResponse) GetErrors() []RestrictedModelError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RestrictedModelsResponse) GetErrorsOk() (*[]RestrictedModelError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RestrictedModelsResponse) SetErrors(v []RestrictedModelError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


