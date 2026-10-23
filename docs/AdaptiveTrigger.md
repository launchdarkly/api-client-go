# AdaptiveTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ResourceKind** | [**AdaptiveTriggerResourceKind**](AdaptiveTriggerResourceKind.md) |  | 
**ResourceKey** | **string** |  | 
**Enabled** | **bool** |  | 
**TriggerSource** | [**ObservabilityAlertSource**](ObservabilityAlertSource.md) |  | 
**Task** | [**AdaptiveTriggerTask**](AdaptiveTriggerTask.md) |  | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewAdaptiveTrigger

`func NewAdaptiveTrigger(id string, name string, resourceKind AdaptiveTriggerResourceKind, resourceKey string, enabled bool, triggerSource ObservabilityAlertSource, task AdaptiveTriggerTask, createdAt int64, updatedAt int64, ) *AdaptiveTrigger`

NewAdaptiveTrigger instantiates a new AdaptiveTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdaptiveTriggerWithDefaults

`func NewAdaptiveTriggerWithDefaults() *AdaptiveTrigger`

NewAdaptiveTriggerWithDefaults instantiates a new AdaptiveTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdaptiveTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdaptiveTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdaptiveTrigger) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AdaptiveTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdaptiveTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdaptiveTrigger) SetName(v string)`

SetName sets Name field to given value.


### GetResourceKind

`func (o *AdaptiveTrigger) GetResourceKind() AdaptiveTriggerResourceKind`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *AdaptiveTrigger) GetResourceKindOk() (*AdaptiveTriggerResourceKind, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *AdaptiveTrigger) SetResourceKind(v AdaptiveTriggerResourceKind)`

SetResourceKind sets ResourceKind field to given value.


### GetResourceKey

`func (o *AdaptiveTrigger) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *AdaptiveTrigger) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *AdaptiveTrigger) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetEnabled

`func (o *AdaptiveTrigger) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdaptiveTrigger) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdaptiveTrigger) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTriggerSource

`func (o *AdaptiveTrigger) GetTriggerSource() ObservabilityAlertSource`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *AdaptiveTrigger) GetTriggerSourceOk() (*ObservabilityAlertSource, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *AdaptiveTrigger) SetTriggerSource(v ObservabilityAlertSource)`

SetTriggerSource sets TriggerSource field to given value.


### GetTask

`func (o *AdaptiveTrigger) GetTask() AdaptiveTriggerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *AdaptiveTrigger) GetTaskOk() (*AdaptiveTriggerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *AdaptiveTrigger) SetTask(v AdaptiveTriggerTask)`

SetTask sets Task field to given value.


### GetCreatedAt

`func (o *AdaptiveTrigger) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdaptiveTrigger) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdaptiveTrigger) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AdaptiveTrigger) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdaptiveTrigger) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdaptiveTrigger) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


