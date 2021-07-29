# GlobalFlagRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**IncludeInSnippet** | Pointer to **bool** |  | [optional] 
**ClientSideAvailability** | Pointer to [**AccountsClientSideAvailability**](AccountsClientSideAvailability.md) |  | [optional] 
**Variations** | Pointer to [**[]VariateRep**](VariateRep.md) |  | [optional] 
**VariationJsonSchema** | Pointer to **interface{}** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**MaintainerId** | Pointer to **string** |  | [optional] 
**Maintainer** | Pointer to [**MemberSummaryRep**](MemberSummaryRep.md) |  | [optional] 
**GoalIds** | Pointer to **[]string** |  | [optional] 
**Experiments** | Pointer to [**ExperimentInfoRep**](ExperimentInfoRep.md) |  | [optional] 
**CustomProperties** | Pointer to [**map[string]CustomProperty**](CustomProperty.md) |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**ArchivedDate** | Pointer to **int64** |  | [optional] 
**Defaults** | Pointer to [**FlagDefaultsRep**](FlagDefaultsRep.md) |  | [optional] 
**Environments** | [**map[string]GlobalFlagRepEnvironments**](GlobalFlagRepEnvironments.md) |  | 

## Methods

### NewGlobalFlagRep

`func NewGlobalFlagRep(environments map[string]GlobalFlagRepEnvironments, ) *GlobalFlagRep`

NewGlobalFlagRep instantiates a new GlobalFlagRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalFlagRepWithDefaults

`func NewGlobalFlagRepWithDefaults() *GlobalFlagRep`

NewGlobalFlagRepWithDefaults instantiates a new GlobalFlagRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlobalFlagRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalFlagRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalFlagRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalFlagRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *GlobalFlagRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GlobalFlagRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GlobalFlagRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GlobalFlagRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalFlagRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalFlagRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalFlagRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalFlagRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *GlobalFlagRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GlobalFlagRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GlobalFlagRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GlobalFlagRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVersion

`func (o *GlobalFlagRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GlobalFlagRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GlobalFlagRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GlobalFlagRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationDate

`func (o *GlobalFlagRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *GlobalFlagRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *GlobalFlagRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *GlobalFlagRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetIncludeInSnippet

`func (o *GlobalFlagRep) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *GlobalFlagRep) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *GlobalFlagRep) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *GlobalFlagRep) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *GlobalFlagRep) GetClientSideAvailability() AccountsClientSideAvailability`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *GlobalFlagRep) GetClientSideAvailabilityOk() (*AccountsClientSideAvailability, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *GlobalFlagRep) SetClientSideAvailability(v AccountsClientSideAvailability)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *GlobalFlagRep) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariations

`func (o *GlobalFlagRep) GetVariations() []VariateRep`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *GlobalFlagRep) GetVariationsOk() (*[]VariateRep, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *GlobalFlagRep) SetVariations(v []VariateRep)`

SetVariations sets Variations field to given value.

### HasVariations

`func (o *GlobalFlagRep) HasVariations() bool`

HasVariations returns a boolean if a field has been set.

### GetVariationJsonSchema

`func (o *GlobalFlagRep) GetVariationJsonSchema() interface{}`

GetVariationJsonSchema returns the VariationJsonSchema field if non-nil, zero value otherwise.

### GetVariationJsonSchemaOk

`func (o *GlobalFlagRep) GetVariationJsonSchemaOk() (*interface{}, bool)`

GetVariationJsonSchemaOk returns a tuple with the VariationJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationJsonSchema

`func (o *GlobalFlagRep) SetVariationJsonSchema(v interface{})`

SetVariationJsonSchema sets VariationJsonSchema field to given value.

### HasVariationJsonSchema

`func (o *GlobalFlagRep) HasVariationJsonSchema() bool`

HasVariationJsonSchema returns a boolean if a field has been set.

### SetVariationJsonSchemaNil

`func (o *GlobalFlagRep) SetVariationJsonSchemaNil(b bool)`

 SetVariationJsonSchemaNil sets the value for VariationJsonSchema to be an explicit nil

### UnsetVariationJsonSchema
`func (o *GlobalFlagRep) UnsetVariationJsonSchema()`

UnsetVariationJsonSchema ensures that no value is present for VariationJsonSchema, not even an explicit nil
### GetTemporary

`func (o *GlobalFlagRep) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *GlobalFlagRep) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *GlobalFlagRep) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *GlobalFlagRep) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetTags

`func (o *GlobalFlagRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GlobalFlagRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GlobalFlagRep) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GlobalFlagRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLinks

`func (o *GlobalFlagRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GlobalFlagRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GlobalFlagRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GlobalFlagRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaintainerId

`func (o *GlobalFlagRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *GlobalFlagRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *GlobalFlagRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *GlobalFlagRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *GlobalFlagRep) GetMaintainer() MemberSummaryRep`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *GlobalFlagRep) GetMaintainerOk() (*MemberSummaryRep, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *GlobalFlagRep) SetMaintainer(v MemberSummaryRep)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *GlobalFlagRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetGoalIds

`func (o *GlobalFlagRep) GetGoalIds() []string`

GetGoalIds returns the GoalIds field if non-nil, zero value otherwise.

### GetGoalIdsOk

`func (o *GlobalFlagRep) GetGoalIdsOk() (*[]string, bool)`

GetGoalIdsOk returns a tuple with the GoalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalIds

`func (o *GlobalFlagRep) SetGoalIds(v []string)`

SetGoalIds sets GoalIds field to given value.

### HasGoalIds

`func (o *GlobalFlagRep) HasGoalIds() bool`

HasGoalIds returns a boolean if a field has been set.

### GetExperiments

`func (o *GlobalFlagRep) GetExperiments() ExperimentInfoRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *GlobalFlagRep) GetExperimentsOk() (*ExperimentInfoRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *GlobalFlagRep) SetExperiments(v ExperimentInfoRep)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *GlobalFlagRep) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetCustomProperties

`func (o *GlobalFlagRep) GetCustomProperties() map[string]CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *GlobalFlagRep) GetCustomPropertiesOk() (*map[string]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *GlobalFlagRep) SetCustomProperties(v map[string]CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *GlobalFlagRep) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetArchived

`func (o *GlobalFlagRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GlobalFlagRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GlobalFlagRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GlobalFlagRep) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedDate

`func (o *GlobalFlagRep) GetArchivedDate() int64`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *GlobalFlagRep) GetArchivedDateOk() (*int64, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *GlobalFlagRep) SetArchivedDate(v int64)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *GlobalFlagRep) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetDefaults

`func (o *GlobalFlagRep) GetDefaults() FlagDefaultsRep`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *GlobalFlagRep) GetDefaultsOk() (*FlagDefaultsRep, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *GlobalFlagRep) SetDefaults(v FlagDefaultsRep)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *GlobalFlagRep) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetEnvironments

`func (o *GlobalFlagRep) GetEnvironments() map[string]GlobalFlagRepEnvironments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *GlobalFlagRep) GetEnvironmentsOk() (*map[string]GlobalFlagRepEnvironments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *GlobalFlagRep) SetEnvironments(v map[string]GlobalFlagRepEnvironments)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


