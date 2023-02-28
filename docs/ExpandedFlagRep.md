# ExpandedFlagRep

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
**MaintainerId** | Pointer to **string** | The ID of the member who maintains the flag | [optional] 
**Maintainer** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**CustomProperties** | [**map[string]CustomProperty**](CustomProperty.md) |  | 
**Archived** | **bool** | Boolean indicating if the feature flag is archived | 
**ArchivedDate** | Pointer to **int64** |  | [optional] 
**Defaults** | Pointer to [**Defaults**](Defaults.md) |  | [optional] 

## Methods

### NewExpandedFlagRep

`func NewExpandedFlagRep(name string, kind string, key string, version int32, creationDate int64, variations []Variation, temporary bool, tags []string, links map[string]Link, customProperties map[string]CustomProperty, archived bool, ) *ExpandedFlagRep`

NewExpandedFlagRep instantiates a new ExpandedFlagRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedFlagRepWithDefaults

`func NewExpandedFlagRepWithDefaults() *ExpandedFlagRep`

NewExpandedFlagRepWithDefaults instantiates a new ExpandedFlagRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExpandedFlagRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedFlagRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedFlagRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *ExpandedFlagRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExpandedFlagRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExpandedFlagRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *ExpandedFlagRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandedFlagRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandedFlagRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandedFlagRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *ExpandedFlagRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedFlagRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedFlagRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *ExpandedFlagRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpandedFlagRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpandedFlagRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *ExpandedFlagRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ExpandedFlagRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ExpandedFlagRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetIncludeInSnippet

`func (o *ExpandedFlagRep) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *ExpandedFlagRep) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *ExpandedFlagRep) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *ExpandedFlagRep) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetClientSideAvailability

`func (o *ExpandedFlagRep) GetClientSideAvailability() ClientSideAvailability`

GetClientSideAvailability returns the ClientSideAvailability field if non-nil, zero value otherwise.

### GetClientSideAvailabilityOk

`func (o *ExpandedFlagRep) GetClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetClientSideAvailabilityOk returns a tuple with the ClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideAvailability

`func (o *ExpandedFlagRep) SetClientSideAvailability(v ClientSideAvailability)`

SetClientSideAvailability sets ClientSideAvailability field to given value.

### HasClientSideAvailability

`func (o *ExpandedFlagRep) HasClientSideAvailability() bool`

HasClientSideAvailability returns a boolean if a field has been set.

### GetVariations

`func (o *ExpandedFlagRep) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *ExpandedFlagRep) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *ExpandedFlagRep) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.


### GetTemporary

`func (o *ExpandedFlagRep) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *ExpandedFlagRep) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *ExpandedFlagRep) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.


### GetTags

`func (o *ExpandedFlagRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExpandedFlagRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExpandedFlagRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *ExpandedFlagRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandedFlagRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandedFlagRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetMaintainerId

`func (o *ExpandedFlagRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ExpandedFlagRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ExpandedFlagRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ExpandedFlagRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *ExpandedFlagRep) GetMaintainer() MemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *ExpandedFlagRep) GetMaintainerOk() (*MemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *ExpandedFlagRep) SetMaintainer(v MemberSummary)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *ExpandedFlagRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ExpandedFlagRep) GetCustomProperties() map[string]CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ExpandedFlagRep) GetCustomPropertiesOk() (*map[string]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ExpandedFlagRep) SetCustomProperties(v map[string]CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.


### GetArchived

`func (o *ExpandedFlagRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ExpandedFlagRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ExpandedFlagRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedDate

`func (o *ExpandedFlagRep) GetArchivedDate() int64`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *ExpandedFlagRep) GetArchivedDateOk() (*int64, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *ExpandedFlagRep) SetArchivedDate(v int64)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *ExpandedFlagRep) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetDefaults

`func (o *ExpandedFlagRep) GetDefaults() Defaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *ExpandedFlagRep) GetDefaultsOk() (*Defaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *ExpandedFlagRep) SetDefaults(v Defaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *ExpandedFlagRep) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


