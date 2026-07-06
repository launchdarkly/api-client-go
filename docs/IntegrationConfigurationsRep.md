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
**SnowflakeSetupScript** | Pointer to **string** | Consolidated SQL script for Snowflake Warehouse Native Experimentation setup. Present only for setup endpoint responses. | [optional] 
**RedshiftSetupScripts** | Pointer to **[]string** | SQL setup scripts (4 parts) for Redshift Native Experimentation setup. Present only for setup endpoint responses. | [optional] 
**RedshiftIAMPermissionsPolicy** | Pointer to **string** | IAM permissions policy JSON for the customer&#39;s Redshift IAM role. Present only for setup endpoint responses. | [optional] 
**RedshiftIAMTrustPolicy** | Pointer to **string** | IAM trust policy JSON for the customer&#39;s Redshift IAM role. Present only for setup endpoint responses. | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this integration configuration. Defaults to the member who created it. | [optional] 

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

### GetSnowflakeSetupScript

`func (o *IntegrationConfigurationsRep) GetSnowflakeSetupScript() string`

GetSnowflakeSetupScript returns the SnowflakeSetupScript field if non-nil, zero value otherwise.

### GetSnowflakeSetupScriptOk

`func (o *IntegrationConfigurationsRep) GetSnowflakeSetupScriptOk() (*string, bool)`

GetSnowflakeSetupScriptOk returns a tuple with the SnowflakeSetupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeSetupScript

`func (o *IntegrationConfigurationsRep) SetSnowflakeSetupScript(v string)`

SetSnowflakeSetupScript sets SnowflakeSetupScript field to given value.

### HasSnowflakeSetupScript

`func (o *IntegrationConfigurationsRep) HasSnowflakeSetupScript() bool`

HasSnowflakeSetupScript returns a boolean if a field has been set.

### GetRedshiftSetupScripts

`func (o *IntegrationConfigurationsRep) GetRedshiftSetupScripts() []string`

GetRedshiftSetupScripts returns the RedshiftSetupScripts field if non-nil, zero value otherwise.

### GetRedshiftSetupScriptsOk

`func (o *IntegrationConfigurationsRep) GetRedshiftSetupScriptsOk() (*[]string, bool)`

GetRedshiftSetupScriptsOk returns a tuple with the RedshiftSetupScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftSetupScripts

`func (o *IntegrationConfigurationsRep) SetRedshiftSetupScripts(v []string)`

SetRedshiftSetupScripts sets RedshiftSetupScripts field to given value.

### HasRedshiftSetupScripts

`func (o *IntegrationConfigurationsRep) HasRedshiftSetupScripts() bool`

HasRedshiftSetupScripts returns a boolean if a field has been set.

### GetRedshiftIAMPermissionsPolicy

`func (o *IntegrationConfigurationsRep) GetRedshiftIAMPermissionsPolicy() string`

GetRedshiftIAMPermissionsPolicy returns the RedshiftIAMPermissionsPolicy field if non-nil, zero value otherwise.

### GetRedshiftIAMPermissionsPolicyOk

`func (o *IntegrationConfigurationsRep) GetRedshiftIAMPermissionsPolicyOk() (*string, bool)`

GetRedshiftIAMPermissionsPolicyOk returns a tuple with the RedshiftIAMPermissionsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftIAMPermissionsPolicy

`func (o *IntegrationConfigurationsRep) SetRedshiftIAMPermissionsPolicy(v string)`

SetRedshiftIAMPermissionsPolicy sets RedshiftIAMPermissionsPolicy field to given value.

### HasRedshiftIAMPermissionsPolicy

`func (o *IntegrationConfigurationsRep) HasRedshiftIAMPermissionsPolicy() bool`

HasRedshiftIAMPermissionsPolicy returns a boolean if a field has been set.

### GetRedshiftIAMTrustPolicy

`func (o *IntegrationConfigurationsRep) GetRedshiftIAMTrustPolicy() string`

GetRedshiftIAMTrustPolicy returns the RedshiftIAMTrustPolicy field if non-nil, zero value otherwise.

### GetRedshiftIAMTrustPolicyOk

`func (o *IntegrationConfigurationsRep) GetRedshiftIAMTrustPolicyOk() (*string, bool)`

GetRedshiftIAMTrustPolicyOk returns a tuple with the RedshiftIAMTrustPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftIAMTrustPolicy

`func (o *IntegrationConfigurationsRep) SetRedshiftIAMTrustPolicy(v string)`

SetRedshiftIAMTrustPolicy sets RedshiftIAMTrustPolicy field to given value.

### HasRedshiftIAMTrustPolicy

`func (o *IntegrationConfigurationsRep) HasRedshiftIAMTrustPolicy() bool`

HasRedshiftIAMTrustPolicy returns a boolean if a field has been set.

### GetMaintainerId

`func (o *IntegrationConfigurationsRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *IntegrationConfigurationsRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *IntegrationConfigurationsRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *IntegrationConfigurationsRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


