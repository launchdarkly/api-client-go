# EnvironmentReleasePlanStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The step&#39;s server-generated identifier | [optional] [readonly] 
**Allocation** | **int64** | Traffic allocation in per-100k units (100000 &#x3D; 100%) | 
**DurationMillis** | **int64** | How long this step holds before the next one starts | 

## Methods

### NewEnvironmentReleasePlanStep

`func NewEnvironmentReleasePlanStep(allocation int64, durationMillis int64, ) *EnvironmentReleasePlanStep`

NewEnvironmentReleasePlanStep instantiates a new EnvironmentReleasePlanStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanStepWithDefaults

`func NewEnvironmentReleasePlanStepWithDefaults() *EnvironmentReleasePlanStep`

NewEnvironmentReleasePlanStepWithDefaults instantiates a new EnvironmentReleasePlanStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReleasePlanStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReleasePlanStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReleasePlanStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentReleasePlanStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllocation

`func (o *EnvironmentReleasePlanStep) GetAllocation() int64`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *EnvironmentReleasePlanStep) GetAllocationOk() (*int64, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *EnvironmentReleasePlanStep) SetAllocation(v int64)`

SetAllocation sets Allocation field to given value.


### GetDurationMillis

`func (o *EnvironmentReleasePlanStep) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *EnvironmentReleasePlanStep) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *EnvironmentReleasePlanStep) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


