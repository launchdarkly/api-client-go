# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the feature flag | 
**Kind** | **string** | Kind of feature flag | 
**Description** | Pointer to **string** | Description of the feature flag | [optional] 
**Key** | **string** | A unique key used to reference the flag in your code | 
**Version** | **int32** | Version of the feature flag | 
**CreationDate** | **int64** |  | 
**IncludeInSnippet** | Pointer to **bool** | Deprecated, use &lt;code&gt;clientSideAvailability&lt;/code&gt;. Whether this flag should be made available to the client-side JavaScript SDK | [optional] 
**ClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**Variations** | [**[]Variation**](Variation.md) | An array of possible variations for the flag | 
**Temporary** | **bool** | Whether the flag is a temporary flag | 
**Tags** | **[]string** | Tags for the feature flag | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**MaintainerId** | Pointer to **string** | Associated maintainerId for the feature flag | [optional] 
**Maintainer** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**MaintainerTeamKey** | Pointer to **string** | The key of the associated team that maintains this feature flag | [optional] 
**MaintainerTeam** | Pointer to [**MaintainerTeam**](MaintainerTeam.md) |  | [optional] 
**GoalIds** | Pointer to **[]string** | Deprecated, use &lt;code&gt;experiments&lt;/code&gt; instead | [optional] 
**Experiments** | [**ExperimentInfoRep**](ExperimentInfoRep.md) |  | 
**CustomProperties** | [**map[string]CustomProperty**](CustomProperty.md) |  | 
**Archived** | **bool** | Boolean indicating if the feature flag is archived | 
**ArchivedDate** | Pointer to **int64** |  | [optional] 
**Defaults** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**MigrationSettings** | Pointer to [**FlagMigrationSettingsRep**](FlagMigrationSettingsRep.md) |  | [optional] 
**Environments** | [**map[string]FeatureFlagConfig**](FeatureFlagConfig.md) | Details on the environments for this flag | 

## Methods

### NewFeatureFlag

`func NewFeatureFlag(name string, kind string, key string, version int32, creationDate int64, variations []Variation, temporary bool, tags []string, links map[string]Link, experiments ExperimentInfoRep, customProperties map[string]CustomProperty, archived bool, environments map[string]FeatureFlagConfig, ) *FeatureFlag`

NewFeatureFlag instantiates a new FeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagWithDefaults

`func NewFeatureFlagWithDefaults() *FeatureFlag`

NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlag) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *FeatureFlag) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FeatureFlag) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FeatureFlag) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *FeatureFlag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureFlag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureFlag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureFlag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *FeatureFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureFlag) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *FeatureFlag) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureFlag) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureFlag) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *FeatureFlag) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *FeatureFlag) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *FeatureFlag) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetIncludeInSnippet

`func (o *FeatureFlag) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *FeatureFlag) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *FeatureFlag) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *FeatureFlag) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *FeatureFlag) GetClientSideAvailability() ClientSideAvailability`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *FeatureFlag) GetClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *FeatureFlag) SetClientSideAvailability(v ClientSideAvailability)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *FeatureFlag) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariations

`func (o *FeatureFlag) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *FeatureFlag) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *FeatureFlag) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.


### GetTemporary

`func (o *FeatureFlag) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FeatureFlag) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FeatureFlag) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.


### GetTags

`func (o *FeatureFlag) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureFlag) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureFlag) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *FeatureFlag) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlag) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlag) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetMaintainerId

`func (o *FeatureFlag) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *FeatureFlag) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *FeatureFlag) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *FeatureFlag) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *FeatureFlag) GetMaintainer() MemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *FeatureFlag) GetMaintainerOk() (*MemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *FeatureFlag) SetMaintainer(v MemberSummary)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *FeatureFlag) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *FeatureFlag) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *FeatureFlag) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *FeatureFlag) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *FeatureFlag) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetMaintainerTeam

`func (o *FeatureFlag) GetMaintainerTeam() MaintainerTeam`

GetMaintainerTeam returns the MaintainerTeam field if non-nil, zero value otherwise.

### GetMaintainerTeamOk

`func (o *FeatureFlag) GetMaintainerTeamOk() (*MaintainerTeam, bool)`

GetMaintainerTeamOk returns a tuple with the MaintainerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeam

`func (o *FeatureFlag) SetMaintainerTeam(v MaintainerTeam)`

SetMaintainerTeam sets MaintainerTeam field to given value.

### HasMaintainerTeam

`func (o *FeatureFlag) HasMaintainerTeam() bool`

HasMaintainerTeam returns a boolean if a field has been set.

### GetGoalIds

`func (o *FeatureFlag) GetGoalIds() []string`

GetGoalIds returns the GoalIds field if non-nil, zero value otherwise.

### GetGoalIdsOk

`func (o *FeatureFlag) GetGoalIdsOk() (*[]string, bool)`

GetGoalIdsOk returns a tuple with the GoalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalIds

`func (o *FeatureFlag) SetGoalIds(v []string)`

SetGoalIds sets GoalIds field to given value.

### HasGoalIds

`func (o *FeatureFlag) HasGoalIds() bool`

HasGoalIds returns a boolean if a field has been set.

### GetExperiments

`func (o *FeatureFlag) GetExperiments() ExperimentInfoRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *FeatureFlag) GetExperimentsOk() (*ExperimentInfoRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *FeatureFlag) SetExperiments(v ExperimentInfoRep)`

SetExperiments sets Experiments field to given value.


### GetCustomProperties

`func (o *FeatureFlag) GetCustomProperties() map[string]CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *FeatureFlag) GetCustomPropertiesOk() (*map[string]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *FeatureFlag) SetCustomProperties(v map[string]CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.


### GetArchived

`func (o *FeatureFlag) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FeatureFlag) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FeatureFlag) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedDate

`func (o *FeatureFlag) GetArchivedDate() int64`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *FeatureFlag) GetArchivedDateOk() (*int64, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *FeatureFlag) SetArchivedDate(v int64)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *FeatureFlag) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetDefaults

`func (o *FeatureFlag) GetDefaults() Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *FeatureFlag) GetDefaultsOk() (*Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *FeatureFlag) SetDefaults(v Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *FeatureFlag) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetPurpose

`func (o *FeatureFlag) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *FeatureFlag) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *FeatureFlag) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *FeatureFlag) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetMigrationSettings

`func (o *FeatureFlag) GetMigrationSettings() FlagMigrationSettingsRep`

GetMigrationSettings returns the MigrationSettings field if non-nil, zero value otherwise.

### GetMigrationSettingsOk

`func (o *FeatureFlag) GetMigrationSettingsOk() (*FlagMigrationSettingsRep, bool)`

GetMigrationSettingsOk returns a tuple with the MigrationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationSettings

`func (o *FeatureFlag) SetMigrationSettings(v FlagMigrationSettingsRep)`

SetMigrationSettings sets MigrationSettings field to given value.

### HasMigrationSettings

`func (o *FeatureFlag) HasMigrationSettings() bool`

HasMigrationSettings returns a boolean if a field has been set.

### GetEnvironments

`func (o *FeatureFlag) GetEnvironments() map[string]FeatureFlagConfig`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *FeatureFlag) GetEnvironmentsOk() (*map[string]FeatureFlagConfig, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *FeatureFlag) SetEnvironments(v map[string]FeatureFlagConfig)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


