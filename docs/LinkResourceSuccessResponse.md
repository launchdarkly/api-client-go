# LinkResourceSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessCount** | **int32** | The number of resources successfully linked. | 
**FailureCount** | **int32** | The number of resources that failed to link. | 
**LinkedResources** | Pointer to [**ViewLinkedResources**](ViewLinkedResources.md) |  | [optional] 
**FailedResources** | Pointer to [**[]FailedResourceLink**](FailedResourceLink.md) | Details of resources that failed to link. | [optional] 

## Methods

### NewLinkResourceSuccessResponse

`func NewLinkResourceSuccessResponse(successCount int32, failureCount int32, ) *LinkResourceSuccessResponse`

NewLinkResourceSuccessResponse instantiates a new LinkResourceSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkResourceSuccessResponseWithDefaults

`func NewLinkResourceSuccessResponseWithDefaults() *LinkResourceSuccessResponse`

NewLinkResourceSuccessResponseWithDefaults instantiates a new LinkResourceSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessCount

`func (o *LinkResourceSuccessResponse) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *LinkResourceSuccessResponse) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *LinkResourceSuccessResponse) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.


### GetFailureCount

`func (o *LinkResourceSuccessResponse) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *LinkResourceSuccessResponse) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *LinkResourceSuccessResponse) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.


### GetLinkedResources

`func (o *LinkResourceSuccessResponse) GetLinkedResources() ViewLinkedResources`

GetLinkedResources returns the LinkedResources field if non-nil, zero value otherwise.

### GetLinkedResourcesOk

`func (o *LinkResourceSuccessResponse) GetLinkedResourcesOk() (*ViewLinkedResources, bool)`

GetLinkedResourcesOk returns a tuple with the LinkedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResources

`func (o *LinkResourceSuccessResponse) SetLinkedResources(v ViewLinkedResources)`

SetLinkedResources sets LinkedResources field to given value.

### HasLinkedResources

`func (o *LinkResourceSuccessResponse) HasLinkedResources() bool`

HasLinkedResources returns a boolean if a field has been set.

### GetFailedResources

`func (o *LinkResourceSuccessResponse) GetFailedResources() []FailedResourceLink`

GetFailedResources returns the FailedResources field if non-nil, zero value otherwise.

### GetFailedResourcesOk

`func (o *LinkResourceSuccessResponse) GetFailedResourcesOk() (*[]FailedResourceLink, bool)`

GetFailedResourcesOk returns a tuple with the FailedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedResources

`func (o *LinkResourceSuccessResponse) SetFailedResources(v []FailedResourceLink)`

SetFailedResources sets FailedResources field to given value.

### HasFailedResources

`func (o *LinkResourceSuccessResponse) HasFailedResources() bool`

HasFailedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


