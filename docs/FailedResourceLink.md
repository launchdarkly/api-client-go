# FailedResourceLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceKey** | **string** | The key of the resource that failed to link. | 
**EnvironmentId** | Pointer to **string** | Environment ID of the resource (only present for segments) | [optional] 
**ResourceType** | **string** | The type of the resource that failed to link. | 
**ErrorMessage** | **string** | The reason why linking this resource failed. | 

## Methods

### NewFailedResourceLink

`func NewFailedResourceLink(resourceKey string, resourceType string, errorMessage string, ) *FailedResourceLink`

NewFailedResourceLink instantiates a new FailedResourceLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedResourceLinkWithDefaults

`func NewFailedResourceLinkWithDefaults() *FailedResourceLink`

NewFailedResourceLinkWithDefaults instantiates a new FailedResourceLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceKey

`func (o *FailedResourceLink) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *FailedResourceLink) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *FailedResourceLink) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetEnvironmentId

`func (o *FailedResourceLink) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FailedResourceLink) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FailedResourceLink) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *FailedResourceLink) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetResourceType

`func (o *FailedResourceLink) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *FailedResourceLink) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *FailedResourceLink) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetErrorMessage

`func (o *FailedResourceLink) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FailedResourceLink) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FailedResourceLink) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


