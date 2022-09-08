# ExpiringUserTargetPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExpiringUserTargetItem**](ExpiringUserTargetItem.md) | An array of expiring user targets | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TotalInstructions** | Pointer to **int32** | The total count of instructions sent in the PATCH request | [optional] 
**SuccessfulInstructions** | Pointer to **int32** | The total count of successful instructions sent in the PATCH request | [optional] 
**FailedInstructions** | Pointer to **int32** | The total count of the failed instructions sent in the PATCH request | [optional] 
**Errors** | Pointer to [**[]ExpiringTargetError**](ExpiringTargetError.md) | An array of error messages for the failed instructions | [optional] 

## Methods

### NewExpiringUserTargetPatchResponse

`func NewExpiringUserTargetPatchResponse(items []ExpiringUserTargetItem, ) *ExpiringUserTargetPatchResponse`

NewExpiringUserTargetPatchResponse instantiates a new ExpiringUserTargetPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringUserTargetPatchResponseWithDefaults

`func NewExpiringUserTargetPatchResponseWithDefaults() *ExpiringUserTargetPatchResponse`

NewExpiringUserTargetPatchResponseWithDefaults instantiates a new ExpiringUserTargetPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExpiringUserTargetPatchResponse) GetItems() []ExpiringUserTargetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExpiringUserTargetPatchResponse) GetItemsOk() (*[]ExpiringUserTargetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExpiringUserTargetPatchResponse) SetItems(v []ExpiringUserTargetItem)`

SetItems sets Items field to given value.


### GetLinks

`func (o *ExpiringUserTargetPatchResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpiringUserTargetPatchResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpiringUserTargetPatchResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpiringUserTargetPatchResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalInstructions

`func (o *ExpiringUserTargetPatchResponse) GetTotalInstructions() int32`

GetTotalInstructions returns the TotalInstructions field if non-nil, zero value otherwise.

### GetTotalInstructionsOk

`func (o *ExpiringUserTargetPatchResponse) GetTotalInstructionsOk() (*int32, bool)`

GetTotalInstructionsOk returns a tuple with the TotalInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstructions

`func (o *ExpiringUserTargetPatchResponse) SetTotalInstructions(v int32)`

SetTotalInstructions sets TotalInstructions field to given value.

### HasTotalInstructions

`func (o *ExpiringUserTargetPatchResponse) HasTotalInstructions() bool`

HasTotalInstructions returns a boolean if a field has been set.

### GetSuccessfulInstructions

`func (o *ExpiringUserTargetPatchResponse) GetSuccessfulInstructions() int32`

GetSuccessfulInstructions returns the SuccessfulInstructions field if non-nil, zero value otherwise.

### GetSuccessfulInstructionsOk

`func (o *ExpiringUserTargetPatchResponse) GetSuccessfulInstructionsOk() (*int32, bool)`

GetSuccessfulInstructionsOk returns a tuple with the SuccessfulInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulInstructions

`func (o *ExpiringUserTargetPatchResponse) SetSuccessfulInstructions(v int32)`

SetSuccessfulInstructions sets SuccessfulInstructions field to given value.

### HasSuccessfulInstructions

`func (o *ExpiringUserTargetPatchResponse) HasSuccessfulInstructions() bool`

HasSuccessfulInstructions returns a boolean if a field has been set.

### GetFailedInstructions

`func (o *ExpiringUserTargetPatchResponse) GetFailedInstructions() int32`

GetFailedInstructions returns the FailedInstructions field if non-nil, zero value otherwise.

### GetFailedInstructionsOk

`func (o *ExpiringUserTargetPatchResponse) GetFailedInstructionsOk() (*int32, bool)`

GetFailedInstructionsOk returns a tuple with the FailedInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInstructions

`func (o *ExpiringUserTargetPatchResponse) SetFailedInstructions(v int32)`

SetFailedInstructions sets FailedInstructions field to given value.

### HasFailedInstructions

`func (o *ExpiringUserTargetPatchResponse) HasFailedInstructions() bool`

HasFailedInstructions returns a boolean if a field has been set.

### GetErrors

`func (o *ExpiringUserTargetPatchResponse) GetErrors() []ExpiringTargetError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExpiringUserTargetPatchResponse) GetErrorsOk() (*[]ExpiringTargetError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExpiringUserTargetPatchResponse) SetErrors(v []ExpiringTargetError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExpiringUserTargetPatchResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


