# UnlinkResourceSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessCount** | **int32** | The number of resources successfully unlinked. | 
**FailureCount** | **int32** | The number of resources that failed to unlink. | 
**FailedResources** | Pointer to [**[]FailedResourceLink**](FailedResourceLink.md) | Details of resources that failed to unlink. | [optional] 

## Methods

### NewUnlinkResourceSuccessResponse

`func NewUnlinkResourceSuccessResponse(successCount int32, failureCount int32, ) *UnlinkResourceSuccessResponse`

NewUnlinkResourceSuccessResponse instantiates a new UnlinkResourceSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkResourceSuccessResponseWithDefaults

`func NewUnlinkResourceSuccessResponseWithDefaults() *UnlinkResourceSuccessResponse`

NewUnlinkResourceSuccessResponseWithDefaults instantiates a new UnlinkResourceSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessCount

`func (o *UnlinkResourceSuccessResponse) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnlinkResourceSuccessResponse) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnlinkResourceSuccessResponse) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.


### GetFailureCount

`func (o *UnlinkResourceSuccessResponse) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *UnlinkResourceSuccessResponse) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *UnlinkResourceSuccessResponse) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.


### GetFailedResources

`func (o *UnlinkResourceSuccessResponse) GetFailedResources() []FailedResourceLink`

GetFailedResources returns the FailedResources field if non-nil, zero value otherwise.

### GetFailedResourcesOk

`func (o *UnlinkResourceSuccessResponse) GetFailedResourcesOk() (*[]FailedResourceLink, bool)`

GetFailedResourcesOk returns a tuple with the FailedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedResources

`func (o *UnlinkResourceSuccessResponse) SetFailedResources(v []FailedResourceLink)`

SetFailedResources sets FailedResources field to given value.

### HasFailedResources

`func (o *UnlinkResourceSuccessResponse) HasFailedResources() bool`

HasFailedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


