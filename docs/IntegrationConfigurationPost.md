# IntegrationConfigurationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the integration configuration | 
**Enabled** | Pointer to **bool** | Whether the integration configuration is enabled. If omitted, defaults to true | [optional] 
**Tags** | Pointer to **[]string** | Tags for the integration | [optional] 
**ConfigValues** | **map[string]interface{}** | The unique set of fields required to configure the integration. Refer to the &lt;code&gt;formVariables&lt;/code&gt; field in the corresponding &lt;code&gt;manifest.json&lt;/code&gt; at https://github.com/launchdarkly/integration-framework/tree/main/integrations for a full list of fields for the integration you wish to configure. | 
**CapabilityConfig** | Pointer to [**CapabilityConfigPost**](CapabilityConfigPost.md) |  | [optional] 

## Methods

### NewIntegrationConfigurationPost

`func NewIntegrationConfigurationPost(name string, configValues map[string]interface{}, ) *IntegrationConfigurationPost`

NewIntegrationConfigurationPost instantiates a new IntegrationConfigurationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConfigurationPostWithDefaults

`func NewIntegrationConfigurationPostWithDefaults() *IntegrationConfigurationPost`

NewIntegrationConfigurationPostWithDefaults instantiates a new IntegrationConfigurationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegrationConfigurationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationConfigurationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationConfigurationPost) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *IntegrationConfigurationPost) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationConfigurationPost) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationConfigurationPost) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationConfigurationPost) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *IntegrationConfigurationPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationConfigurationPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationConfigurationPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IntegrationConfigurationPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfigValues

`func (o *IntegrationConfigurationPost) GetConfigValues() map[string]interface{}`

GetConfigValues returns the ConfigValues field if non-nil, zero value otherwise.

### GetConfigValuesOk

`func (o *IntegrationConfigurationPost) GetConfigValuesOk() (*map[string]interface{}, bool)`

GetConfigValuesOk returns a tuple with the ConfigValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigValues

`func (o *IntegrationConfigurationPost) SetConfigValues(v map[string]interface{})`

SetConfigValues sets ConfigValues field to given value.


### GetCapabilityConfig

`func (o *IntegrationConfigurationPost) GetCapabilityConfig() CapabilityConfigPost`

GetCapabilityConfig returns the CapabilityConfig field if non-nil, zero value otherwise.

### GetCapabilityConfigOk

`func (o *IntegrationConfigurationPost) GetCapabilityConfigOk() (*CapabilityConfigPost, bool)`

GetCapabilityConfigOk returns a tuple with the CapabilityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityConfig

`func (o *IntegrationConfigurationPost) SetCapabilityConfig(v CapabilityConfigPost)`

SetCapabilityConfig sets CapabilityConfig field to given value.

### HasCapabilityConfig

`func (o *IntegrationConfigurationPost) HasCapabilityConfig() bool`

HasCapabilityConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


