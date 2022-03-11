# IntegrationDeliveryConfigurationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | Pointer to **bool** | Default value is false | [optional] 
**Config** | **map[string]interface{}** |  | 
**Tags** | Pointer to **[]string** | Tags to associate with integration | [optional] 
**Name** | Pointer to **string** | Name to identify integration | [optional] 

## Methods

### NewIntegrationDeliveryConfigurationPost

`func NewIntegrationDeliveryConfigurationPost(config map[string]interface{}, ) *IntegrationDeliveryConfigurationPost`

NewIntegrationDeliveryConfigurationPost instantiates a new IntegrationDeliveryConfigurationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDeliveryConfigurationPostWithDefaults

`func NewIntegrationDeliveryConfigurationPostWithDefaults() *IntegrationDeliveryConfigurationPost`

NewIntegrationDeliveryConfigurationPostWithDefaults instantiates a new IntegrationDeliveryConfigurationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *IntegrationDeliveryConfigurationPost) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *IntegrationDeliveryConfigurationPost) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *IntegrationDeliveryConfigurationPost) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *IntegrationDeliveryConfigurationPost) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationDeliveryConfigurationPost) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationDeliveryConfigurationPost) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationDeliveryConfigurationPost) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetTags

`func (o *IntegrationDeliveryConfigurationPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationDeliveryConfigurationPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationDeliveryConfigurationPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IntegrationDeliveryConfigurationPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *IntegrationDeliveryConfigurationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationDeliveryConfigurationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationDeliveryConfigurationPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationDeliveryConfigurationPost) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


