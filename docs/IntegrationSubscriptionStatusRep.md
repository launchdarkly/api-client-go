# IntegrationSubscriptionStatusRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessCount** | Pointer to **int32** |  | [optional] 
**LastSuccess** | Pointer to **int64** |  | [optional] 
**LastError** | Pointer to **int64** |  | [optional] 
**ErrorCount** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]IntegrationSubscriptionStatusRepErrors**](IntegrationSubscriptionStatusRepErrors.md) |  | [optional] 

## Methods

### NewIntegrationSubscriptionStatusRep

`func NewIntegrationSubscriptionStatusRep() *IntegrationSubscriptionStatusRep`

NewIntegrationSubscriptionStatusRep instantiates a new IntegrationSubscriptionStatusRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSubscriptionStatusRepWithDefaults

`func NewIntegrationSubscriptionStatusRepWithDefaults() *IntegrationSubscriptionStatusRep`

NewIntegrationSubscriptionStatusRepWithDefaults instantiates a new IntegrationSubscriptionStatusRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessCount

`func (o *IntegrationSubscriptionStatusRep) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *IntegrationSubscriptionStatusRep) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *IntegrationSubscriptionStatusRep) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *IntegrationSubscriptionStatusRep) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetLastSuccess

`func (o *IntegrationSubscriptionStatusRep) GetLastSuccess() int64`

GetLastSuccess returns the LastSuccess field if non-nil, zero value otherwise.

### GetLastSuccessOk

`func (o *IntegrationSubscriptionStatusRep) GetLastSuccessOk() (*int64, bool)`

GetLastSuccessOk returns a tuple with the LastSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccess

`func (o *IntegrationSubscriptionStatusRep) SetLastSuccess(v int64)`

SetLastSuccess sets LastSuccess field to given value.

### HasLastSuccess

`func (o *IntegrationSubscriptionStatusRep) HasLastSuccess() bool`

HasLastSuccess returns a boolean if a field has been set.

### GetLastError

`func (o *IntegrationSubscriptionStatusRep) GetLastError() int64`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *IntegrationSubscriptionStatusRep) GetLastErrorOk() (*int64, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *IntegrationSubscriptionStatusRep) SetLastError(v int64)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *IntegrationSubscriptionStatusRep) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetErrorCount

`func (o *IntegrationSubscriptionStatusRep) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *IntegrationSubscriptionStatusRep) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *IntegrationSubscriptionStatusRep) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *IntegrationSubscriptionStatusRep) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetErrors

`func (o *IntegrationSubscriptionStatusRep) GetErrors() []IntegrationSubscriptionStatusRepErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IntegrationSubscriptionStatusRep) GetErrorsOk() (*[]IntegrationSubscriptionStatusRepErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IntegrationSubscriptionStatusRep) SetErrors(v []IntegrationSubscriptionStatusRepErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *IntegrationSubscriptionStatusRep) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


