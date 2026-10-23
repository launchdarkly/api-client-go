# CreateAdaptiveTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ResourceKind** | [**AdaptiveTriggerResourceKind**](AdaptiveTriggerResourceKind.md) |  | 
**ResourceKey** | **string** |  | 
**TriggerSource** | [**ObservabilityAlertSource**](ObservabilityAlertSource.md) |  | 
**Task** | [**AdaptiveTriggerTask**](AdaptiveTriggerTask.md) |  | 

## Methods

### NewCreateAdaptiveTriggerRequest

`func NewCreateAdaptiveTriggerRequest(name string, resourceKind AdaptiveTriggerResourceKind, resourceKey string, triggerSource ObservabilityAlertSource, task AdaptiveTriggerTask, ) *CreateAdaptiveTriggerRequest`

NewCreateAdaptiveTriggerRequest instantiates a new CreateAdaptiveTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdaptiveTriggerRequestWithDefaults

`func NewCreateAdaptiveTriggerRequestWithDefaults() *CreateAdaptiveTriggerRequest`

NewCreateAdaptiveTriggerRequestWithDefaults instantiates a new CreateAdaptiveTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAdaptiveTriggerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAdaptiveTriggerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAdaptiveTriggerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResourceKind

`func (o *CreateAdaptiveTriggerRequest) GetResourceKind() AdaptiveTriggerResourceKind`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *CreateAdaptiveTriggerRequest) GetResourceKindOk() (*AdaptiveTriggerResourceKind, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *CreateAdaptiveTriggerRequest) SetResourceKind(v AdaptiveTriggerResourceKind)`

SetResourceKind sets ResourceKind field to given value.


### GetResourceKey

`func (o *CreateAdaptiveTriggerRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *CreateAdaptiveTriggerRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *CreateAdaptiveTriggerRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetTriggerSource

`func (o *CreateAdaptiveTriggerRequest) GetTriggerSource() ObservabilityAlertSource`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *CreateAdaptiveTriggerRequest) GetTriggerSourceOk() (*ObservabilityAlertSource, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *CreateAdaptiveTriggerRequest) SetTriggerSource(v ObservabilityAlertSource)`

SetTriggerSource sets TriggerSource field to given value.


### GetTask

`func (o *CreateAdaptiveTriggerRequest) GetTask() AdaptiveTriggerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *CreateAdaptiveTriggerRequest) GetTaskOk() (*AdaptiveTriggerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *CreateAdaptiveTriggerRequest) SetTask(v AdaptiveTriggerTask)`

SetTask sets Task field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


