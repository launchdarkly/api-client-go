# FeatureFlagBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the feature flag | 
**Key** | **string** | A unique key used to reference the flag in your code | 
**Description** | Pointer to **string** | Description of the feature flag. Defaults to an empty string. | [optional] 
**IncludeInSnippet** | Pointer to **bool** | Deprecated, use &lt;code&gt;clientSideAvailability&lt;/code&gt;. Whether this flag should be made available to the client-side JavaScript SDK. Defaults to &lt;code&gt;false&lt;/code&gt;. | [optional] 
**ClientSideAvailability** | Pointer to [**ClientSideAvailabilityPost**](ClientSideAvailabilityPost.md) |  | [optional] 
**Variations** | Pointer to [**[]Variation**](Variation.md) | An array of possible variations for the flag. The variation values must be unique. If omitted, two boolean variations of &lt;code&gt;true&lt;/code&gt; and &lt;code&gt;false&lt;/code&gt; will be used. | [optional] 
**Temporary** | Pointer to **bool** | Whether the flag is a temporary flag. Defaults to &lt;code&gt;true&lt;/code&gt;. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the feature flag. Defaults to an empty array. | [optional] 
**CustomProperties** | Pointer to [**map[string]CustomProperty**](CustomProperty.md) |  | [optional] 
**Defaults** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 
**Purpose** | Pointer to **string** | Purpose of the flag | [optional] 
**MigrationSettings** | Pointer to [**MigrationSettingsPost**](MigrationSettingsPost.md) |  | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this feature flag | [optional] 
**MaintainerTeamKey** | Pointer to **string** | The key of the team that maintains this feature flag | [optional] 
**InitialPrerequisites** | Pointer to [**[]FlagPrerequisitePost**](FlagPrerequisitePost.md) | Initial set of prerequisite flags for all environments | [optional] 
**IsFlagOn** | Pointer to **bool** | Whether to automatically turn the flag on across all environments at creation. Defaults to &lt;code&gt;false&lt;/code&gt;. | [optional] 

## Methods

### NewFeatureFlagBody

`func NewFeatureFlagBody(name string, key string, ) *FeatureFlagBody`

NewFeatureFlagBody instantiates a new FeatureFlagBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagBodyWithDefaults

`func NewFeatureFlagBodyWithDefaults() *FeatureFlagBody`

NewFeatureFlagBodyWithDefaults instantiates a new FeatureFlagBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureFlagBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlagBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlagBody) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *FeatureFlagBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureFlagBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureFlagBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *FeatureFlagBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureFlagBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureFlagBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureFlagBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeInSnippet

`func (o *FeatureFlagBody) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *FeatureFlagBody) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *FeatureFlagBody) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *FeatureFlagBody) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *FeatureFlagBody) GetClientSideAvailability() ClientSideAvailabilityPost`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *FeatureFlagBody) GetClientSideAvailabilityOk() (*ClientSideAvailabilityPost, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *FeatureFlagBody) SetClientSideAvailability(v ClientSideAvailabilityPost)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *FeatureFlagBody) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariations

`func (o *FeatureFlagBody) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *FeatureFlagBody) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *FeatureFlagBody) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.

### HasVariations

`func (o *FeatureFlagBody) HasVariations() bool`

HasVariations returns a boolean if a field has been set.

### GetTemporary

`func (o *FeatureFlagBody) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FeatureFlagBody) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FeatureFlagBody) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *FeatureFlagBody) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetTags

`func (o *FeatureFlagBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureFlagBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureFlagBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeatureFlagBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomProperties

`func (o *FeatureFlagBody) GetCustomProperties() map[string]CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *FeatureFlagBody) GetCustomPropertiesOk() (*map[string]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *FeatureFlagBody) SetCustomProperties(v map[string]CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *FeatureFlagBody) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDefaults

`func (o *FeatureFlagBody) GetDefaults() Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *FeatureFlagBody) GetDefaultsOk() (*Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *FeatureFlagBody) SetDefaults(v Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *FeatureFlagBody) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetPurpose

`func (o *FeatureFlagBody) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *FeatureFlagBody) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *FeatureFlagBody) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *FeatureFlagBody) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMigrationSettings

`func (o *FeatureFlagBody) GetMigrationSettings() MigrationSettingsPost`

GetMigrationSettings returns the MigrationSettings field if non-nil, zero value otherwise.

### GetMigrationSettingsOk

`func (o *FeatureFlagBody) GetMigrationSettingsOk() (*MigrationSettingsPost, bool)`

GetMigrationSettingsOk returns a tuple with the MigrationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationSettings

`func (o *FeatureFlagBody) SetMigrationSettings(v MigrationSettingsPost)`

SetMigrationSettings sets MigrationSettings field to given value.

### HasMigrationSettings

`func (o *FeatureFlagBody) HasMigrationSettings() bool`

HasMigrationSettings returns a boolean if a field has been set.

### GetMaintainerId

`func (o *FeatureFlagBody) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *FeatureFlagBody) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *FeatureFlagBody) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *FeatureFlagBody) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *FeatureFlagBody) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *FeatureFlagBody) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *FeatureFlagBody) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *FeatureFlagBody) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetInitialPrerequisites

`func (o *FeatureFlagBody) GetInitialPrerequisites() []FlagPrerequisitePost`

GetInitialPrerequisites returns the InitialPrerequisites field if non-nil, zero value otherwise.

### GetInitialPrerequisitesOk

`func (o *FeatureFlagBody) GetInitialPrerequisitesOk() (*[]FlagPrerequisitePost, bool)`

GetInitialPrerequisitesOk returns a tuple with the InitialPrerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPrerequisites

`func (o *FeatureFlagBody) SetInitialPrerequisites(v []FlagPrerequisitePost)`

SetInitialPrerequisites sets InitialPrerequisites field to given value.

### HasInitialPrerequisites

`func (o *FeatureFlagBody) HasInitialPrerequisites() bool`

HasInitialPrerequisites returns a boolean if a field has been set.

### GetIsFlagOn

`func (o *FeatureFlagBody) GetIsFlagOn() bool`

GetIsFlagOn returns the IsFlagOn field if non-nil, zero value otherwise.

### GetIsFlagOnOk

`func (o *FeatureFlagBody) GetIsFlagOnOk() (*bool, bool)`

GetIsFlagOnOk returns a tuple with the IsFlagOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlagOn

`func (o *FeatureFlagBody) SetIsFlagOn(v bool)`

SetIsFlagOn sets IsFlagOn field to given value.

### HasIsFlagOn

`func (o *FeatureFlagBody) HasIsFlagOn() bool`

HasIsFlagOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


