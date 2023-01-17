# ExecutionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the execution of this workflow stage | 
**StopDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewExecutionOutput

`func NewExecutionOutput(status string, ) *ExecutionOutput`

NewExecutionOutput instantiates a new ExecutionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionOutputWithDefaults

`func NewExecutionOutputWithDefaults() *ExecutionOutput`

NewExecutionOutputWithDefaults instantiates a new ExecutionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ExecutionOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionOutput) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStopDate

`func (o *ExecutionOutput) GetStopDate() int64`

GetStopDate returns the StopDate field if non-nil, zero value otherwise.

### GetStopDateOk

`func (o *ExecutionOutput) GetStopDateOk() (*int64, bool)`

GetStopDateOk returns a tuple with the StopDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDate

`func (o *ExecutionOutput) SetStopDate(v int64)`

SetStopDate sets StopDate field to given value.

### HasStopDate

`func (o *ExecutionOutput) HasStopDate() bool`

HasStopDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


