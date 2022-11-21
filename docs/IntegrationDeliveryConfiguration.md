# IntegrationDeliveryConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**IntegrationDeliveryConfigurationLinks**](IntegrationDeliveryConfigurationLinks.md) |  | 
**Id** | **string** | The integration ID | 
**IntegrationKey** | **string** | The integration key | 
**ProjectKey** | **string** | The project key | 
**EnvironmentKey** | **string** | The environment key | 
**Config** | **map[string]interface{}** |  | 
**On** | **bool** | Whether the configuration is turned on | 
**Tags** | **[]string** | List of tags for this configuration | 
**Name** | **string** | Name of the configuration | 
**Version** | **int32** | Version of the current configuration | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 

## Methods

### NewIntegrationDeliveryConfiguration

`func NewIntegrationDeliveryConfiguration(links IntegrationDeliveryConfigurationLinks, id string, integrationKey string, projectKey string, environmentKey string, config map[string]interface{}, on bool, tags []string, name string, version int32, ) *IntegrationDeliveryConfiguration`

NewIntegrationDeliveryConfiguration instantiates a new IntegrationDeliveryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDeliveryConfigurationWithDefaults

`func NewIntegrationDeliveryConfigurationWithDefaults() *IntegrationDeliveryConfiguration`

NewIntegrationDeliveryConfigurationWithDefaults instantiates a new IntegrationDeliveryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntegrationDeliveryConfiguration) GetLinks() IntegrationDeliveryConfigurationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntegrationDeliveryConfiguration) GetLinksOk() (*IntegrationDeliveryConfigurationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntegrationDeliveryConfiguration) SetLinks(v IntegrationDeliveryConfigurationLinks)`

SetLinks sets Links field to given value.


### GetId

`func (o *IntegrationDeliveryConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationDeliveryConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationDeliveryConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationKey

`func (o *IntegrationDeliveryConfiguration) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *IntegrationDeliveryConfiguration) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *IntegrationDeliveryConfiguration) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.


### GetProjectKey

`func (o *IntegrationDeliveryConfiguration) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *IntegrationDeliveryConfiguration) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *IntegrationDeliveryConfiguration) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetEnvironmentKey

`func (o *IntegrationDeliveryConfiguration) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *IntegrationDeliveryConfiguration) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *IntegrationDeliveryConfiguration) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetConfig

`func (o *IntegrationDeliveryConfiguration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationDeliveryConfiguration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationDeliveryConfiguration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetOn

`func (o *IntegrationDeliveryConfiguration) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *IntegrationDeliveryConfiguration) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *IntegrationDeliveryConfiguration) SetOn(v bool)`

SetOn sets On field to given value.


### GetTags

`func (o *IntegrationDeliveryConfiguration) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationDeliveryConfiguration) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationDeliveryConfiguration) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetName

`func (o *IntegrationDeliveryConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationDeliveryConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationDeliveryConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *IntegrationDeliveryConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IntegrationDeliveryConfiguration) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IntegrationDeliveryConfiguration) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetAccess

`func (o *IntegrationDeliveryConfiguration) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *IntegrationDeliveryConfiguration) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *IntegrationDeliveryConfiguration) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *IntegrationDeliveryConfiguration) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


