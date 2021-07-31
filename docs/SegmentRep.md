# SegmentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** |  | 
**CreationDate** | **int64** |  | 
**Key** | **string** |  | 
**Included** | **[]string** |  | 
**Excluded** | **[]string** |  | 
**Links** | [**map[string]CoreLink**](CoreLink.md) |  | 
**Rules** | [**[]SegmentRuleRep**](SegmentRuleRep.md) |  | 
**Version** | **int32** |  | 
**Deleted** | **bool** |  | 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**Flags** | Pointer to [**[]FlagListingRep**](FlagListingRep.md) |  | [optional] 
**Unbounded** | Pointer to **bool** |  | [optional] 
**Generation** | **int32** |  | 
**UnboundedMetadata** | Pointer to [**UnboundedSegmentMetadata**](UnboundedSegmentMetadata.md) |  | [optional] 
**External** | Pointer to **string** |  | [optional] 
**ExternalLink** | Pointer to **string** |  | [optional] 

## Methods

### NewSegmentRep

`func NewSegmentRep(name string, tags []string, creationDate int64, key string, included []string, excluded []string, links map[string]CoreLink, rules []SegmentRuleRep, version int32, deleted bool, generation int32, ) *SegmentRep`

NewSegmentRep instantiates a new SegmentRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentRepWithDefaults

`func NewSegmentRepWithDefaults() *SegmentRep`

NewSegmentRepWithDefaults instantiates a new SegmentRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SegmentRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *SegmentRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SegmentRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SegmentRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreationDate

`func (o *SegmentRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SegmentRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SegmentRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetKey

`func (o *SegmentRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SegmentRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SegmentRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetIncluded

`func (o *SegmentRep) GetIncluded() []string`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SegmentRep) GetIncludedOk() (*[]string, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SegmentRep) SetIncluded(v []string)`

SetIncluded sets Included field to given value.


### GetExcluded

`func (o *SegmentRep) GetExcluded() []string`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *SegmentRep) GetExcludedOk() (*[]string, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *SegmentRep) SetExcluded(v []string)`

SetExcluded sets Excluded field to given value.


### GetLinks

`func (o *SegmentRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SegmentRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SegmentRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.


### GetRules

`func (o *SegmentRep) GetRules() []SegmentRuleRep`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SegmentRep) GetRulesOk() (*[]SegmentRuleRep, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SegmentRep) SetRules(v []SegmentRuleRep)`

SetRules sets Rules field to given value.


### GetVersion

`func (o *SegmentRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SegmentRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SegmentRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDeleted

`func (o *SegmentRep) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SegmentRep) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SegmentRep) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetAccess

`func (o *SegmentRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SegmentRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SegmentRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SegmentRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetFlags

`func (o *SegmentRep) GetFlags() []FlagListingRep`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SegmentRep) GetFlagsOk() (*[]FlagListingRep, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SegmentRep) SetFlags(v []FlagListingRep)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SegmentRep) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetUnbounded

`func (o *SegmentRep) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *SegmentRep) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *SegmentRep) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.

### HasUnbounded

`func (o *SegmentRep) HasUnbounded() bool`

HasUnbounded returns a boolean if a field has been set.

### GetGeneration

`func (o *SegmentRep) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *SegmentRep) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *SegmentRep) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.


### GetUnboundedMetadata

`func (o *SegmentRep) GetUnboundedMetadata() UnboundedSegmentMetadata`

GetUnboundedMetadata returns the UnboundedMetadata field if non-nil, zero value otherwise.

### GetUnboundedMetadataOk

`func (o *SegmentRep) GetUnboundedMetadataOk() (*UnboundedSegmentMetadata, bool)`

GetUnboundedMetadataOk returns a tuple with the UnboundedMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundedMetadata

`func (o *SegmentRep) SetUnboundedMetadata(v UnboundedSegmentMetadata)`

SetUnboundedMetadata sets UnboundedMetadata field to given value.

### HasUnboundedMetadata

`func (o *SegmentRep) HasUnboundedMetadata() bool`

HasUnboundedMetadata returns a boolean if a field has been set.

### GetExternal

`func (o *SegmentRep) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *SegmentRep) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *SegmentRep) SetExternal(v string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *SegmentRep) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalLink

`func (o *SegmentRep) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *SegmentRep) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *SegmentRep) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *SegmentRep) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


