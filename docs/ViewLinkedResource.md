# ViewLinkedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | 
**ResourceKey** | **string** | Key of the resource (flag, segment, AI config or metric) | 
**EnvironmentId** | Pointer to **string** | Environment ID of the resource (only present for segments) | [optional] 
**EnvironmentKey** | Pointer to **string** | Environment Key of the resource (only present for segments) | [optional] 
**ResourceType** | **string** |  | 
**LinkedAt** | **int64** |  | 
**ResourceDetails** | Pointer to [**ViewLinkedResourceDetails**](ViewLinkedResourceDetails.md) |  | [optional] 

## Methods

### NewViewLinkedResource

`func NewViewLinkedResource(links ParentAndSelfLinks, resourceKey string, resourceType string, linkedAt int64, ) *ViewLinkedResource`

NewViewLinkedResource instantiates a new ViewLinkedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLinkedResourceWithDefaults

`func NewViewLinkedResourceWithDefaults() *ViewLinkedResource`

NewViewLinkedResourceWithDefaults instantiates a new ViewLinkedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ViewLinkedResource) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ViewLinkedResource) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ViewLinkedResource) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.


### GetResourceKey

`func (o *ViewLinkedResource) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ViewLinkedResource) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ViewLinkedResource) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetEnvironmentId

`func (o *ViewLinkedResource) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ViewLinkedResource) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ViewLinkedResource) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ViewLinkedResource) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *ViewLinkedResource) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *ViewLinkedResource) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *ViewLinkedResource) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *ViewLinkedResource) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetResourceType

`func (o *ViewLinkedResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ViewLinkedResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ViewLinkedResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetLinkedAt

`func (o *ViewLinkedResource) GetLinkedAt() int64`

GetLinkedAt returns the LinkedAt field if non-nil, zero value otherwise.

### GetLinkedAtOk

`func (o *ViewLinkedResource) GetLinkedAtOk() (*int64, bool)`

GetLinkedAtOk returns a tuple with the LinkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAt

`func (o *ViewLinkedResource) SetLinkedAt(v int64)`

SetLinkedAt sets LinkedAt field to given value.


### GetResourceDetails

`func (o *ViewLinkedResource) GetResourceDetails() ViewLinkedResourceDetails`

GetResourceDetails returns the ResourceDetails field if non-nil, zero value otherwise.

### GetResourceDetailsOk

`func (o *ViewLinkedResource) GetResourceDetailsOk() (*ViewLinkedResourceDetails, bool)`

GetResourceDetailsOk returns a tuple with the ResourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDetails

`func (o *ViewLinkedResource) SetResourceDetails(v ViewLinkedResourceDetails)`

SetResourceDetails sets ResourceDetails field to given value.

### HasResourceDetails

`func (o *ViewLinkedResource) HasResourceDetails() bool`

HasResourceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


