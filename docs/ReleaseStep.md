# ReleaseStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepIndex** | **int32** | Zero-based index of this step. | [readonly] 
**Allocation** | **int32** | Thousandths of a percent. &#x60;100000&#x60; is 100%, and &#x60;10000&#x60; is 10%. Guarded releases cannot exceed &#x60;50000&#x60;, preserving a control group.  | 
**DurationMillis** | **int64** |  | 
**StartedAt** | Pointer to **int64** | Start of the step in Unix milliseconds. | [optional] [readonly] 
**EndedAt** | Pointer to **int64** | End of the step in Unix milliseconds. | [optional] [readonly] 

## Methods

### NewReleaseStep

`func NewReleaseStep(stepIndex int32, allocation int32, durationMillis int64, ) *ReleaseStep`

NewReleaseStep instantiates a new ReleaseStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseStepWithDefaults

`func NewReleaseStepWithDefaults() *ReleaseStep`

NewReleaseStepWithDefaults instantiates a new ReleaseStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepIndex

`func (o *ReleaseStep) GetStepIndex() int32`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *ReleaseStep) GetStepIndexOk() (*int32, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *ReleaseStep) SetStepIndex(v int32)`

SetStepIndex sets StepIndex field to given value.


### GetAllocation

`func (o *ReleaseStep) GetAllocation() int32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *ReleaseStep) GetAllocationOk() (*int32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *ReleaseStep) SetAllocation(v int32)`

SetAllocation sets Allocation field to given value.


### GetDurationMillis

`func (o *ReleaseStep) GetDurationMillis() int64`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *ReleaseStep) GetDurationMillisOk() (*int64, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *ReleaseStep) SetDurationMillis(v int64)`

SetDurationMillis sets DurationMillis field to given value.


### GetStartedAt

`func (o *ReleaseStep) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ReleaseStep) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ReleaseStep) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ReleaseStep) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *ReleaseStep) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ReleaseStep) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ReleaseStep) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ReleaseStep) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


