# IntegrationConfigurationsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Id** | **string** | The unique identifier for this integration configuration | 
**Name** | **string** | A human-friendly name for the integration | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**IntegrationKey** | Pointer to **string** | The type of integration | [optional] 
**Tags** | Pointer to **[]string** | An array of tags for this integration | [optional] 
**Enabled** | Pointer to **bool** | Whether the integration is currently active | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**ConfigValues** | Pointer to **map[string]interface{}** | Details on configuration for an integration of this type. Refer to the &lt;code&gt;formVariables&lt;/code&gt; field in the corresponding &lt;code&gt;manifest.json&lt;/code&gt; for a full list of fields for each integration. | [optional] 
**CapabilityConfig** | Pointer to [**CapabilityConfigRep**](CapabilityConfigRep.md) |  | [optional] 

## Methods

### NewIntegrationConfigurationsRep

`func NewIntegrationConfigurationsRep(links map[string]Link, id string, name string, ) *IntegrationConfigurationsRep`

NewIntegrationConfigurationsRep instantiates a new IntegrationConfigurationsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConfigurationsRepWithDefaults

`func NewIntegrationConfigurationsRepWithDefaults() *IntegrationConfigurationsRep`

NewIntegrationConfigurationsRepWithDefaults instantiates a new IntegrationConfigurationsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntegrationConfigurationsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntegrationConfigurationsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntegrationConfigurationsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *IntegrationConfigurationsRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationConfigurationsRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationConfigurationsRep) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IntegrationConfigurationsRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationConfigurationsRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationConfigurationsRep) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *IntegrationConfigurationsRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationConfigurationsRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationConfigurationsRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IntegrationConfigurationsRep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *IntegrationConfigurationsRep) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *IntegrationConfigurationsRep) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *IntegrationConfigurationsRep) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *IntegrationConfigurationsRep) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetTags

`func (o *IntegrationConfigurationsRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationConfigurationsRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationConfigurationsRep) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IntegrationConfigurationsRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegrationConfigurationsRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationConfigurationsRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationConfigurationsRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationConfigurationsRep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAccess

`func (o *IntegrationConfigurationsRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *IntegrationConfigurationsRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *IntegrationConfigurationsRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *IntegrationConfigurationsRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetConfigValues

`func (o *IntegrationConfigurationsRep) GetConfigValues() map[string]interface{}`

GetConfigValues returns the ConfigValues field if non-nil, zero value otherwise.

### GetConfigValuesOk

`func (o *IntegrationConfigurationsRep) GetConfigValuesOk() (*map[string]interface{}, bool)`

GetConfigValuesOk returns a tuple with the ConfigValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigValues

`func (o *IntegrationConfigurationsRep) SetConfigValues(v map[string]interface{})`

SetConfigValues sets ConfigValues field to given value.

### HasConfigValues

`func (o *IntegrationConfigurationsRep) HasConfigValues() bool`

HasConfigValues returns a boolean if a field has been set.

### GetCapabilityConfig

`func (o *IntegrationConfigurationsRep) GetCapabilityConfig() CapabilityConfigRep`

GetCapabilityConfig returns the CapabilityConfig field if non-nil, zero value otherwise.

### GetCapabilityConfigOk

`func (o *IntegrationConfigurationsRep) GetCapabilityConfigOk() (*CapabilityConfigRep, bool)`

GetCapabilityConfigOk returns a tuple with the CapabilityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityConfig

`func (o *IntegrationConfigurationsRep) SetCapabilityConfig(v CapabilityConfigRep)`

SetCapabilityConfig sets CapabilityConfig field to given value.

### HasCapabilityConfig

`func (o *IntegrationConfigurationsRep) HasCapabilityConfig() bool`

HasCapabilityConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


