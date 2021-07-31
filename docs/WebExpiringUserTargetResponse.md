# WebExpiringUserTargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WebExpiringUserTargetItem**](WebExpiringUserTargetItem.md) |  | [optional] 
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
**TotalInstructions** | Pointer to **int32** |  | [optional] 
**SuccessfulInstructions** | Pointer to **int32** |  | [optional] 
**FailedInstructions** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]WebExpiringUserTargetError**](WebExpiringUserTargetError.md) |  | [optional] 

## Methods

### NewWebExpiringUserTargetResponse

`func NewWebExpiringUserTargetResponse() *WebExpiringUserTargetResponse`

NewWebExpiringUserTargetResponse instantiates a new WebExpiringUserTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebExpiringUserTargetResponseWithDefaults

`func NewWebExpiringUserTargetResponseWithDefaults() *WebExpiringUserTargetResponse`

NewWebExpiringUserTargetResponseWithDefaults instantiates a new WebExpiringUserTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WebExpiringUserTargetResponse) GetItems() []WebExpiringUserTargetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WebExpiringUserTargetResponse) GetItemsOk() (*[]WebExpiringUserTargetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WebExpiringUserTargetResponse) SetItems(v []WebExpiringUserTargetItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *WebExpiringUserTargetResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *WebExpiringUserTargetResponse) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebExpiringUserTargetResponse) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebExpiringUserTargetResponse) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebExpiringUserTargetResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalInstructions

`func (o *WebExpiringUserTargetResponse) GetTotalInstructions() int32`

GetTotalInstructions returns the TotalInstructions field if non-nil, zero value otherwise.

### GetTotalInstructionsOk

`func (o *WebExpiringUserTargetResponse) GetTotalInstructionsOk() (*int32, bool)`

GetTotalInstructionsOk returns a tuple with the TotalInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstructions

`func (o *WebExpiringUserTargetResponse) SetTotalInstructions(v int32)`

SetTotalInstructions sets TotalInstructions field to given value.

### HasTotalInstructions

`func (o *WebExpiringUserTargetResponse) HasTotalInstructions() bool`

HasTotalInstructions returns a boolean if a field has been set.

### GetSuccessfulInstructions

`func (o *WebExpiringUserTargetResponse) GetSuccessfulInstructions() int32`

GetSuccessfulInstructions returns the SuccessfulInstructions field if non-nil, zero value otherwise.

### GetSuccessfulInstructionsOk

`func (o *WebExpiringUserTargetResponse) GetSuccessfulInstructionsOk() (*int32, bool)`

GetSuccessfulInstructionsOk returns a tuple with the SuccessfulInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulInstructions

`func (o *WebExpiringUserTargetResponse) SetSuccessfulInstructions(v int32)`

SetSuccessfulInstructions sets SuccessfulInstructions field to given value.

### HasSuccessfulInstructions

`func (o *WebExpiringUserTargetResponse) HasSuccessfulInstructions() bool`

HasSuccessfulInstructions returns a boolean if a field has been set.

### GetFailedInstructions

`func (o *WebExpiringUserTargetResponse) GetFailedInstructions() int32`

GetFailedInstructions returns the FailedInstructions field if non-nil, zero value otherwise.

### GetFailedInstructionsOk

`func (o *WebExpiringUserTargetResponse) GetFailedInstructionsOk() (*int32, bool)`

GetFailedInstructionsOk returns a tuple with the FailedInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInstructions

`func (o *WebExpiringUserTargetResponse) SetFailedInstructions(v int32)`

SetFailedInstructions sets FailedInstructions field to given value.

### HasFailedInstructions

`func (o *WebExpiringUserTargetResponse) HasFailedInstructions() bool`

HasFailedInstructions returns a boolean if a field has been set.

### GetErrors

`func (o *WebExpiringUserTargetResponse) GetErrors() []WebExpiringUserTargetError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WebExpiringUserTargetResponse) GetErrorsOk() (*[]WebExpiringUserTargetError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WebExpiringUserTargetResponse) SetErrors(v []WebExpiringUserTargetError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WebExpiringUserTargetResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


