# EnvironmentReleaseStepInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocation** | **int32** | Thousandths of a percent. &#x60;100000&#x60; is 100%, and &#x60;10000&#x60; is 10%. Guarded releases cannot exceed &#x60;50000&#x60;, preserving a control group. Each step&#39;s allocation must be greater than the previous step&#39;s.  In a progressive release, an allocation of N% serves the target variation to N% of the phase&#39;s contexts and the original variation to the rest. In a guarded release, it serves the target variation to N%, the original variation to another N% as the measured control group, and the original variation to the rest without measuring them.  | 
**DurationMillis** | **int64** |  | 

## Methods

### NewEnvironmentReleaseStepInput

`func NewEnvironmentReleaseStepInput(allocation int32, durationMillis int64, ) *EnvironmentReleaseStepInput`

NewEnvironmentReleaseStepInput instantiates a new EnvironmentReleaseStepInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleaseStepInputWithDefaults

`func NewEnvironmentReleaseStepInputWithDefaults() *EnvironmentReleaseStepInput`

NewEnvironmentReleaseStepInputWithDefaults instantiates a new EnvironmentReleaseStepInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocation

`func (o *EnvironmentReleaseStepInput) GetAllocation() int32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *EnvironmentReleaseStepInput) GetAllocationOk() (*int32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *EnvironmentReleaseStepInput) SetAllocation(v int32)`

SetAllocation sets Allocation field to given value.


### GetDurationMillis

`func (o *EnvironmentReleaseStepInput) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *EnvironmentReleaseStepInput) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *EnvironmentReleaseStepInput) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


