# UpdateAdaptiveTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Task** | Pointer to [**AdaptiveTriggerTask**](AdaptiveTriggerTask.md) |  | [optional] 

## Methods

### NewUpdateAdaptiveTriggerRequest

`func NewUpdateAdaptiveTriggerRequest() *UpdateAdaptiveTriggerRequest`

NewUpdateAdaptiveTriggerRequest instantiates a new UpdateAdaptiveTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdaptiveTriggerRequestWithDefaults

`func NewUpdateAdaptiveTriggerRequestWithDefaults() *UpdateAdaptiveTriggerRequest`

NewUpdateAdaptiveTriggerRequestWithDefaults instantiates a new UpdateAdaptiveTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAdaptiveTriggerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAdaptiveTriggerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAdaptiveTriggerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAdaptiveTriggerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTask

`func (o *UpdateAdaptiveTriggerRequest) GetTask() AdaptiveTriggerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UpdateAdaptiveTriggerRequest) GetTaskOk() (*AdaptiveTriggerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UpdateAdaptiveTriggerRequest) SetTask(v AdaptiveTriggerTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *UpdateAdaptiveTriggerRequest) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


