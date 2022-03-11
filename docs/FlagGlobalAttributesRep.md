# FlagGlobalAttributesRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the feature flag | 
**Kind** | **string** | Kind of feature flag | 
**Description** | Pointer to **string** | Description of the feature flag | [optional] 
**Key** | **string** | A unique key used to reference the flag in your code | 
**Version** | **int32** | Version of the feature flag | 
**CreationDate** | **int64** |  | 
**IncludeInSnippet** | Pointer to **bool** | Deprecated, use clientSideAvailability. Whether or not this flag should be made available to the client-side JavaScript SDK | [optional] 
**ClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**Variations** | [**[]Variation**](Variation.md) | An array of possible variations for the flag | 
**VariationJsonSchema** | Pointer to **interface{}** |  | [optional] 
**Temporary** | **bool** | Whether or not the flag is a temporary flag | 
**Tags** | **[]string** | Tags for the feature flag | 
**Links** | [**map[string]Link**](Link.md) |  | 
**MaintainerId** | Pointer to **string** | Associated maintainerId for the feature flag | [optional] 
**Maintainer** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**GoalIds** | Pointer to **[]string** |  | [optional] 
**Experiments** | [**ExperimentInfoRep**](ExperimentInfoRep.md) |  | 
**CustomProperties** | [**map[string]CustomProperty**](CustomProperty.md) |  | 
**Archived** | **bool** | Boolean indicating if the feature flag is archived | 
**ArchivedDate** | Pointer to **int64** |  | [optional] 
**Defaults** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 

## Methods

### NewFlagGlobalAttributesRep

`func NewFlagGlobalAttributesRep(name string, kind string, key string, version int32, creationDate int64, variations []Variation, temporary bool, tags []string, links map[string]Link, experiments ExperimentInfoRep, customProperties map[string]CustomProperty, archived bool, ) *FlagGlobalAttributesRep`

NewFlagGlobalAttributesRep instantiates a new FlagGlobalAttributesRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagGlobalAttributesRepWithDefaults

`func NewFlagGlobalAttributesRepWithDefaults() *FlagGlobalAttributesRep`

NewFlagGlobalAttributesRepWithDefaults instantiates a new FlagGlobalAttributesRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlagGlobalAttributesRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagGlobalAttributesRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagGlobalAttributesRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *FlagGlobalAttributesRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FlagGlobalAttributesRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FlagGlobalAttributesRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *FlagGlobalAttributesRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagGlobalAttributesRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagGlobalAttributesRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagGlobalAttributesRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *FlagGlobalAttributesRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagGlobalAttributesRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagGlobalAttributesRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *FlagGlobalAttributesRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlagGlobalAttributesRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlagGlobalAttributesRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *FlagGlobalAttributesRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *FlagGlobalAttributesRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *FlagGlobalAttributesRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetIncludeInSnippet

`func (o *FlagGlobalAttributesRep) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *FlagGlobalAttributesRep) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *FlagGlobalAttributesRep) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *FlagGlobalAttributesRep) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *FlagGlobalAttributesRep) GetClientSideAvailability() ClientSideAvailability`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *FlagGlobalAttributesRep) GetClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *FlagGlobalAttributesRep) SetClientSideAvailability(v ClientSideAvailability)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *FlagGlobalAttributesRep) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariations

`func (o *FlagGlobalAttributesRep) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *FlagGlobalAttributesRep) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *FlagGlobalAttributesRep) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.


### GetVariationJsonSchema

`func (o *FlagGlobalAttributesRep) GetVariationJsonSchema() interface{}`

GetVariationJsonSchema returns the VariationJsonSchema field if non-nil, zero value otherwise.

### GetVariationJsonSchemaOk

`func (o *FlagGlobalAttributesRep) GetVariationJsonSchemaOk() (*interface{}, bool)`

GetVariationJsonSchemaOk returns a tuple with the VariationJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationJsonSchema

`func (o *FlagGlobalAttributesRep) SetVariationJsonSchema(v interface{})`

SetVariationJsonSchema sets VariationJsonSchema field to given value.

### HasVariationJsonSchema

`func (o *FlagGlobalAttributesRep) HasVariationJsonSchema() bool`

HasVariationJsonSchema returns a boolean if a field has been set.

### SetVariationJsonSchemaNil

`func (o *FlagGlobalAttributesRep) SetVariationJsonSchemaNil(b bool)`

 SetVariationJsonSchemaNil sets the value for VariationJsonSchema to be an explicit nil

### UnsetVariationJsonSchema
`func (o *FlagGlobalAttributesRep) UnsetVariationJsonSchema()`

UnsetVariationJsonSchema ensures that no value is present for VariationJsonSchema, not even an explicit nil
### GetTemporary

`func (o *FlagGlobalAttributesRep) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FlagGlobalAttributesRep) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FlagGlobalAttributesRep) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.


### GetTags

`func (o *FlagGlobalAttributesRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagGlobalAttributesRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagGlobalAttributesRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *FlagGlobalAttributesRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagGlobalAttributesRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagGlobalAttributesRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetMaintainerId

`func (o *FlagGlobalAttributesRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *FlagGlobalAttributesRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *FlagGlobalAttributesRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *FlagGlobalAttributesRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *FlagGlobalAttributesRep) GetMaintainer() MemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *FlagGlobalAttributesRep) GetMaintainerOk() (*MemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *FlagGlobalAttributesRep) SetMaintainer(v MemberSummary)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *FlagGlobalAttributesRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetGoalIds

`func (o *FlagGlobalAttributesRep) GetGoalIds() []string`

GetGoalIds returns the GoalIds field if non-nil, zero value otherwise.

### GetGoalIdsOk

`func (o *FlagGlobalAttributesRep) GetGoalIdsOk() (*[]string, bool)`

GetGoalIdsOk returns a tuple with the GoalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalIds

`func (o *FlagGlobalAttributesRep) SetGoalIds(v []string)`

SetGoalIds sets GoalIds field to given value.

### HasGoalIds

`func (o *FlagGlobalAttributesRep) HasGoalIds() bool`

HasGoalIds returns a boolean if a field has been set.

### GetExperiments

`func (o *FlagGlobalAttributesRep) GetExperiments() ExperimentInfoRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *FlagGlobalAttributesRep) GetExperimentsOk() (*ExperimentInfoRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *FlagGlobalAttributesRep) SetExperiments(v ExperimentInfoRep)`

SetExperiments sets Experiments field to given value.


### GetCustomProperties

`func (o *FlagGlobalAttributesRep) GetCustomProperties() map[string]CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *FlagGlobalAttributesRep) GetCustomPropertiesOk() (*map[string]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *FlagGlobalAttributesRep) SetCustomProperties(v map[string]CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.


### GetArchived

`func (o *FlagGlobalAttributesRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FlagGlobalAttributesRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FlagGlobalAttributesRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedDate

`func (o *FlagGlobalAttributesRep) GetArchivedDate() int64`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *FlagGlobalAttributesRep) GetArchivedDateOk() (*int64, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *FlagGlobalAttributesRep) SetArchivedDate(v int64)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *FlagGlobalAttributesRep) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetDefaults

`func (o *FlagGlobalAttributesRep) GetDefaults() Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *FlagGlobalAttributesRep) GetDefaultsOk() (*Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *FlagGlobalAttributesRep) SetDefaults(v Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *FlagGlobalAttributesRep) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


