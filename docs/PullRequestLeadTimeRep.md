# PullRequestLeadTimeRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodingDurationMs** | **int64** | The coding duration in milliseconds | 
**ReviewDurationMs** | Pointer to **int64** | The review duration in milliseconds | [optional] 
**MaxWaitDurationMs** | Pointer to **int64** | The max wait duration between merge time and deploy start time in milliseconds | [optional] 
**AvgWaitDurationMs** | Pointer to **int64** | The average wait duration between merge time and deploy start time in milliseconds | [optional] 
**MaxDeployDurationMs** | Pointer to **int64** | The max deploy duration in milliseconds | [optional] 
**AvgDeployDurationMs** | Pointer to **int64** | The average deploy duration in milliseconds | [optional] 
**MaxTotalLeadTimeMs** | Pointer to **int64** | The max total lead time in milliseconds | [optional] 
**AvgTotalLeadTimeMs** | Pointer to **int64** | The average total lead time in milliseconds | [optional] 

## Methods

### NewPullRequestLeadTimeRep

`func NewPullRequestLeadTimeRep(codingDurationMs int64, ) *PullRequestLeadTimeRep`

NewPullRequestLeadTimeRep instantiates a new PullRequestLeadTimeRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestLeadTimeRepWithDefaults

`func NewPullRequestLeadTimeRepWithDefaults() *PullRequestLeadTimeRep`

NewPullRequestLeadTimeRepWithDefaults instantiates a new PullRequestLeadTimeRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodingDurationMs

`func (o *PullRequestLeadTimeRep) GetCodingDurationMs() int64`

GetCodingDurationMs returns the CodingDurationMs field if non-nil, zero value otherwise.

### GetCodingDurationMsOk

`func (o *PullRequestLeadTimeRep) GetCodingDurationMsOk() (*int64, bool)`

GetCodingDurationMsOk returns a tuple with the CodingDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingDurationMs

`func (o *PullRequestLeadTimeRep) SetCodingDurationMs(v int64)`

SetCodingDurationMs sets CodingDurationMs field to given value.


### GetReviewDurationMs

`func (o *PullRequestLeadTimeRep) GetReviewDurationMs() int64`

GetReviewDurationMs returns the ReviewDurationMs field if non-nil, zero value otherwise.

### GetReviewDurationMsOk

`func (o *PullRequestLeadTimeRep) GetReviewDurationMsOk() (*int64, bool)`

GetReviewDurationMsOk returns a tuple with the ReviewDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewDurationMs

`func (o *PullRequestLeadTimeRep) SetReviewDurationMs(v int64)`

SetReviewDurationMs sets ReviewDurationMs field to given value.

### HasReviewDurationMs

`func (o *PullRequestLeadTimeRep) HasReviewDurationMs() bool`

HasReviewDurationMs returns a boolean if a field has been set.

### GetMaxWaitDurationMs

`func (o *PullRequestLeadTimeRep) GetMaxWaitDurationMs() int64`

GetMaxWaitDurationMs returns the MaxWaitDurationMs field if non-nil, zero value otherwise.

### GetMaxWaitDurationMsOk

`func (o *PullRequestLeadTimeRep) GetMaxWaitDurationMsOk() (*int64, bool)`

GetMaxWaitDurationMsOk returns a tuple with the MaxWaitDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitDurationMs

`func (o *PullRequestLeadTimeRep) SetMaxWaitDurationMs(v int64)`

SetMaxWaitDurationMs sets MaxWaitDurationMs field to given value.

### HasMaxWaitDurationMs

`func (o *PullRequestLeadTimeRep) HasMaxWaitDurationMs() bool`

HasMaxWaitDurationMs returns a boolean if a field has been set.

### GetAvgWaitDurationMs

`func (o *PullRequestLeadTimeRep) GetAvgWaitDurationMs() int64`

GetAvgWaitDurationMs returns the AvgWaitDurationMs field if non-nil, zero value otherwise.

### GetAvgWaitDurationMsOk

`func (o *PullRequestLeadTimeRep) GetAvgWaitDurationMsOk() (*int64, bool)`

GetAvgWaitDurationMsOk returns a tuple with the AvgWaitDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgWaitDurationMs

`func (o *PullRequestLeadTimeRep) SetAvgWaitDurationMs(v int64)`

SetAvgWaitDurationMs sets AvgWaitDurationMs field to given value.

### HasAvgWaitDurationMs

`func (o *PullRequestLeadTimeRep) HasAvgWaitDurationMs() bool`

HasAvgWaitDurationMs returns a boolean if a field has been set.

### GetMaxDeployDurationMs

`func (o *PullRequestLeadTimeRep) GetMaxDeployDurationMs() int64`

GetMaxDeployDurationMs returns the MaxDeployDurationMs field if non-nil, zero value otherwise.

### GetMaxDeployDurationMsOk

`func (o *PullRequestLeadTimeRep) GetMaxDeployDurationMsOk() (*int64, bool)`

GetMaxDeployDurationMsOk returns a tuple with the MaxDeployDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeployDurationMs

`func (o *PullRequestLeadTimeRep) SetMaxDeployDurationMs(v int64)`

SetMaxDeployDurationMs sets MaxDeployDurationMs field to given value.

### HasMaxDeployDurationMs

`func (o *PullRequestLeadTimeRep) HasMaxDeployDurationMs() bool`

HasMaxDeployDurationMs returns a boolean if a field has been set.

### GetAvgDeployDurationMs

`func (o *PullRequestLeadTimeRep) GetAvgDeployDurationMs() int64`

GetAvgDeployDurationMs returns the AvgDeployDurationMs field if non-nil, zero value otherwise.

### GetAvgDeployDurationMsOk

`func (o *PullRequestLeadTimeRep) GetAvgDeployDurationMsOk() (*int64, bool)`

GetAvgDeployDurationMsOk returns a tuple with the AvgDeployDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDeployDurationMs

`func (o *PullRequestLeadTimeRep) SetAvgDeployDurationMs(v int64)`

SetAvgDeployDurationMs sets AvgDeployDurationMs field to given value.

### HasAvgDeployDurationMs

`func (o *PullRequestLeadTimeRep) HasAvgDeployDurationMs() bool`

HasAvgDeployDurationMs returns a boolean if a field has been set.

### GetMaxTotalLeadTimeMs

`func (o *PullRequestLeadTimeRep) GetMaxTotalLeadTimeMs() int64`

GetMaxTotalLeadTimeMs returns the MaxTotalLeadTimeMs field if non-nil, zero value otherwise.

### GetMaxTotalLeadTimeMsOk

`func (o *PullRequestLeadTimeRep) GetMaxTotalLeadTimeMsOk() (*int64, bool)`

GetMaxTotalLeadTimeMsOk returns a tuple with the MaxTotalLeadTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalLeadTimeMs

`func (o *PullRequestLeadTimeRep) SetMaxTotalLeadTimeMs(v int64)`

SetMaxTotalLeadTimeMs sets MaxTotalLeadTimeMs field to given value.

### HasMaxTotalLeadTimeMs

`func (o *PullRequestLeadTimeRep) HasMaxTotalLeadTimeMs() bool`

HasMaxTotalLeadTimeMs returns a boolean if a field has been set.

### GetAvgTotalLeadTimeMs

`func (o *PullRequestLeadTimeRep) GetAvgTotalLeadTimeMs() int64`

GetAvgTotalLeadTimeMs returns the AvgTotalLeadTimeMs field if non-nil, zero value otherwise.

### GetAvgTotalLeadTimeMsOk

`func (o *PullRequestLeadTimeRep) GetAvgTotalLeadTimeMsOk() (*int64, bool)`

GetAvgTotalLeadTimeMsOk returns a tuple with the AvgTotalLeadTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTotalLeadTimeMs

`func (o *PullRequestLeadTimeRep) SetAvgTotalLeadTimeMs(v int64)`

SetAvgTotalLeadTimeMs sets AvgTotalLeadTimeMs field to given value.

### HasAvgTotalLeadTimeMs

`func (o *PullRequestLeadTimeRep) HasAvgTotalLeadTimeMs() bool`

HasAvgTotalLeadTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


