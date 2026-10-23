# StopEnvironmentReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | [**EnvironmentReleaseResolution**](EnvironmentReleaseResolution.md) |  | 
**Comment** | Pointer to **string** | Optional comment written to the flag&#39;s audit-log entry. | [optional] 

## Methods

### NewStopEnvironmentReleaseRequest

`func NewStopEnvironmentReleaseRequest(resolution EnvironmentReleaseResolution, ) *StopEnvironmentReleaseRequest`

NewStopEnvironmentReleaseRequest instantiates a new StopEnvironmentReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopEnvironmentReleaseRequestWithDefaults

`func NewStopEnvironmentReleaseRequestWithDefaults() *StopEnvironmentReleaseRequest`

NewStopEnvironmentReleaseRequestWithDefaults instantiates a new StopEnvironmentReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolution

`func (o *StopEnvironmentReleaseRequest) GetResolution() EnvironmentReleaseResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *StopEnvironmentReleaseRequest) GetResolutionOk() (*EnvironmentReleaseResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *StopEnvironmentReleaseRequest) SetResolution(v EnvironmentReleaseResolution)`

SetResolution sets Resolution field to given value.


### GetComment

`func (o *StopEnvironmentReleaseRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StopEnvironmentReleaseRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StopEnvironmentReleaseRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StopEnvironmentReleaseRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


