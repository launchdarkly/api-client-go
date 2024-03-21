# LeadTimeStagesRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodingDurationMs** | **int64** | The coding duration in milliseconds | 
**ReviewDurationMs** | Pointer to **int64** | The review duration in milliseconds | [optional] 
**WaitDurationMs** | Pointer to **int64** | The wait duration between merge time and deploy start time in milliseconds | [optional] 
**DeployDurationMs** | Pointer to **int64** | The deploy duration in milliseconds | [optional] 
**TotalLeadTimeMs** | Pointer to **int64** | The total lead time in milliseconds | [optional] 

## Methods

### NewLeadTimeStagesRep

`func NewLeadTimeStagesRep(codingDurationMs int64, ) *LeadTimeStagesRep`

NewLeadTimeStagesRep instantiates a new LeadTimeStagesRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeadTimeStagesRepWithDefaults

`func NewLeadTimeStagesRepWithDefaults() *LeadTimeStagesRep`

NewLeadTimeStagesRepWithDefaults instantiates a new LeadTimeStagesRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodingDurationMs

`func (o *LeadTimeStagesRep) GetCodingDurationMs() int64`

GetCodingDurationMs returns the CodingDurationMs field if non-nil, zero value otherwise.

### GetCodingDurationMsOk

`func (o *LeadTimeStagesRep) GetCodingDurationMsOk() (*int64, bool)`

GetCodingDurationMsOk returns a tuple with the CodingDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingDurationMs

`func (o *LeadTimeStagesRep) SetCodingDurationMs(v int64)`

SetCodingDurationMs sets CodingDurationMs field to given value.


### GetReviewDurationMs

`func (o *LeadTimeStagesRep) GetReviewDurationMs() int64`

GetReviewDurationMs returns the ReviewDurationMs field if non-nil, zero value otherwise.

### GetReviewDurationMsOk

`func (o *LeadTimeStagesRep) GetReviewDurationMsOk() (*int64, bool)`

GetReviewDurationMsOk returns a tuple with the ReviewDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewDurationMs

`func (o *LeadTimeStagesRep) SetReviewDurationMs(v int64)`

SetReviewDurationMs sets ReviewDurationMs field to given value.

### HasReviewDurationMs

`func (o *LeadTimeStagesRep) HasReviewDurationMs() bool`

HasReviewDurationMs returns a boolean if a field has been set.

### GetWaitDurationMs

`func (o *LeadTimeStagesRep) GetWaitDurationMs() int64`

GetWaitDurationMs returns the WaitDurationMs field if non-nil, zero value otherwise.

### GetWaitDurationMsOk

`func (o *LeadTimeStagesRep) GetWaitDurationMsOk() (*int64, bool)`

GetWaitDurationMsOk returns a tuple with the WaitDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitDurationMs

`func (o *LeadTimeStagesRep) SetWaitDurationMs(v int64)`

SetWaitDurationMs sets WaitDurationMs field to given value.

### HasWaitDurationMs

`func (o *LeadTimeStagesRep) HasWaitDurationMs() bool`

HasWaitDurationMs returns a boolean if a field has been set.

### GetDeployDurationMs

`func (o *LeadTimeStagesRep) GetDeployDurationMs() int64`

GetDeployDurationMs returns the DeployDurationMs field if non-nil, zero value otherwise.

### GetDeployDurationMsOk

`func (o *LeadTimeStagesRep) GetDeployDurationMsOk() (*int64, bool)`

GetDeployDurationMsOk returns a tuple with the DeployDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDurationMs

`func (o *LeadTimeStagesRep) SetDeployDurationMs(v int64)`

SetDeployDurationMs sets DeployDurationMs field to given value.

### HasDeployDurationMs

`func (o *LeadTimeStagesRep) HasDeployDurationMs() bool`

HasDeployDurationMs returns a boolean if a field has been set.

### GetTotalLeadTimeMs

`func (o *LeadTimeStagesRep) GetTotalLeadTimeMs() int64`

GetTotalLeadTimeMs returns the TotalLeadTimeMs field if non-nil, zero value otherwise.

### GetTotalLeadTimeMsOk

`func (o *LeadTimeStagesRep) GetTotalLeadTimeMsOk() (*int64, bool)`

GetTotalLeadTimeMsOk returns a tuple with the TotalLeadTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLeadTimeMs

`func (o *LeadTimeStagesRep) SetTotalLeadTimeMs(v int64)`

SetTotalLeadTimeMs sets TotalLeadTimeMs field to given value.

### HasTotalLeadTimeMs

`func (o *LeadTimeStagesRep) HasTotalLeadTimeMs() bool`

HasTotalLeadTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


