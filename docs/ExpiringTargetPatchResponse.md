# ExpiringTargetPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExpiringTarget**](ExpiringTarget.md) | A list of the results from each instruction | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TotalInstructions** | Pointer to **int32** |  | [optional] 
**SuccessfulInstructions** | Pointer to **int32** |  | [optional] 
**FailedInstructions** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]ExpiringTargetError**](ExpiringTargetError.md) |  | [optional] 

## Methods

### NewExpiringTargetPatchResponse

`func NewExpiringTargetPatchResponse(items []ExpiringTarget, ) *ExpiringTargetPatchResponse`

NewExpiringTargetPatchResponse instantiates a new ExpiringTargetPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringTargetPatchResponseWithDefaults

`func NewExpiringTargetPatchResponseWithDefaults() *ExpiringTargetPatchResponse`

NewExpiringTargetPatchResponseWithDefaults instantiates a new ExpiringTargetPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExpiringTargetPatchResponse) GetItems() []ExpiringTarget`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExpiringTargetPatchResponse) GetItemsOk() (*[]ExpiringTarget, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExpiringTargetPatchResponse) SetItems(v []ExpiringTarget)`

SetItems sets Items field to given value.


### GetLinks

`func (o *ExpiringTargetPatchResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpiringTargetPatchResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpiringTargetPatchResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpiringTargetPatchResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalInstructions

`func (o *ExpiringTargetPatchResponse) GetTotalInstructions() int32`

GetTotalInstructions returns the TotalInstructions field if non-nil, zero value otherwise.

### GetTotalInstructionsOk

`func (o *ExpiringTargetPatchResponse) GetTotalInstructionsOk() (*int32, bool)`

GetTotalInstructionsOk returns a tuple with the TotalInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstructions

`func (o *ExpiringTargetPatchResponse) SetTotalInstructions(v int32)`

SetTotalInstructions sets TotalInstructions field to given value.

### HasTotalInstructions

`func (o *ExpiringTargetPatchResponse) HasTotalInstructions() bool`

HasTotalInstructions returns a boolean if a field has been set.

### GetSuccessfulInstructions

`func (o *ExpiringTargetPatchResponse) GetSuccessfulInstructions() int32`

GetSuccessfulInstructions returns the SuccessfulInstructions field if non-nil, zero value otherwise.

### GetSuccessfulInstructionsOk

`func (o *ExpiringTargetPatchResponse) GetSuccessfulInstructionsOk() (*int32, bool)`

GetSuccessfulInstructionsOk returns a tuple with the SuccessfulInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulInstructions

`func (o *ExpiringTargetPatchResponse) SetSuccessfulInstructions(v int32)`

SetSuccessfulInstructions sets SuccessfulInstructions field to given value.

### HasSuccessfulInstructions

`func (o *ExpiringTargetPatchResponse) HasSuccessfulInstructions() bool`

HasSuccessfulInstructions returns a boolean if a field has been set.

### GetFailedInstructions

`func (o *ExpiringTargetPatchResponse) GetFailedInstructions() int32`

GetFailedInstructions returns the FailedInstructions field if non-nil, zero value otherwise.

### GetFailedInstructionsOk

`func (o *ExpiringTargetPatchResponse) GetFailedInstructionsOk() (*int32, bool)`

GetFailedInstructionsOk returns a tuple with the FailedInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInstructions

`func (o *ExpiringTargetPatchResponse) SetFailedInstructions(v int32)`

SetFailedInstructions sets FailedInstructions field to given value.

### HasFailedInstructions

`func (o *ExpiringTargetPatchResponse) HasFailedInstructions() bool`

HasFailedInstructions returns a boolean if a field has been set.

### GetErrors

`func (o *ExpiringTargetPatchResponse) GetErrors() []ExpiringTargetError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExpiringTargetPatchResponse) GetErrorsOk() (*[]ExpiringTargetError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExpiringTargetPatchResponse) SetErrors(v []ExpiringTargetError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExpiringTargetPatchResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


