# UserSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the segment. | 
**Description** | Pointer to **string** | A description of the segment&#39;s purpose. Defaults to &lt;code&gt;null&lt;/code&gt; and is omitted in the response if not provided. | [optional] 
**Tags** | **[]string** | Tags for the segment. Defaults to an empty array. | 
**CreationDate** | **int64** |  | 
**LastModifiedDate** | **int64** |  | 
**Key** | **string** | A unique key used to reference the segment | 
**Included** | Pointer to **[]string** | An array of keys for included targets. Included individual targets are always segment members, regardless of segment rules. For list-based segments over 15,000 entries, also called big segments, this array is either empty or omitted. | [optional] 
**Excluded** | Pointer to **[]string** | An array of keys for excluded targets. Segment rules bypass individual excluded targets, so they will never be included based on rules. Excluded targets may still be included explicitly. This value is omitted for list-based segments over 15,000 entries, also called big segments. | [optional] 
**IncludedContexts** | Pointer to [**[]SegmentTarget**](SegmentTarget.md) |  | [optional] 
**ExcludedContexts** | Pointer to [**[]SegmentTarget**](SegmentTarget.md) |  | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Rules** | [**[]UserSegmentRule**](UserSegmentRule.md) | An array of the targeting rules for this segment. | 
**Version** | **int32** | Version of the segment | 
**Deleted** | **bool** | Whether the segment has been deleted | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Flags** | Pointer to [**[]FlagListingRep**](FlagListingRep.md) | A list of flags targeting this segment. Only included when getting a single segment, using the &lt;code&gt;getSegment&lt;/code&gt; endpoint. | [optional] 
**Unbounded** | Pointer to **bool** | Whether this is a standard segment (&lt;code&gt;false&lt;/code&gt;) or a big segment (&lt;code&gt;true&lt;/code&gt;). Standard segments include rule-based segments and smaller list-based segments. Big segments include larger list-based segments and synced segments. If omitted, the segment is a standard segment. | [optional] 
**UnboundedContextKind** | Pointer to **string** | For big segments, the targeted context kind. | [optional] 
**Generation** | **int32** | For big segments, how many times this segment has been created. | 
**UnboundedMetadata** | Pointer to [**SegmentMetadata**](SegmentMetadata.md) |  | [optional] 
**External** | Pointer to **string** | The external data store backing this segment. Only applies to synced segments. | [optional] 
**ExternalLink** | Pointer to **string** | The URL for the external data store backing this segment. Only applies to synced segments. | [optional] 
**ImportInProgress** | Pointer to **bool** | Whether an import is currently in progress for the specified segment. Only applies to big segments. | [optional] 

## Methods

### NewUserSegment

`func NewUserSegment(name string, tags []string, creationDate int64, lastModifiedDate int64, key string, links map[string]Link, rules []UserSegmentRule, version int32, deleted bool, generation int32, ) *UserSegment`

NewUserSegment instantiates a new UserSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSegmentWithDefaults

`func NewUserSegmentWithDefaults() *UserSegment`

NewUserSegmentWithDefaults instantiates a new UserSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSegment) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UserSegment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSegment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSegment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserSegment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *UserSegment) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserSegment) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserSegment) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreationDate

`func (o *UserSegment) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UserSegment) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UserSegment) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModifiedDate

`func (o *UserSegment) GetLastModifiedDate() int64`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *UserSegment) GetLastModifiedDateOk() (*int64, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *UserSegment) SetLastModifiedDate(v int64)`

SetLastModifiedDate sets LastModifiedDate field to given value.


### GetKey

`func (o *UserSegment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserSegment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserSegment) SetKey(v string)`

SetKey sets Key field to given value.


### GetIncluded

`func (o *UserSegment) GetIncluded() []string`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UserSegment) GetIncludedOk() (*[]string, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UserSegment) SetIncluded(v []string)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UserSegment) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetExcluded

`func (o *UserSegment) GetExcluded() []string`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *UserSegment) GetExcludedOk() (*[]string, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *UserSegment) SetExcluded(v []string)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *UserSegment) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetIncludedContexts

`func (o *UserSegment) GetIncludedContexts() []SegmentTarget`

GetIncludedContexts returns the IncludedContexts field if non-nil, zero value otherwise.

### GetIncludedContextsOk

`func (o *UserSegment) GetIncludedContextsOk() (*[]SegmentTarget, bool)`

GetIncludedContextsOk returns a tuple with the IncludedContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedContexts

`func (o *UserSegment) SetIncludedContexts(v []SegmentTarget)`

SetIncludedContexts sets IncludedContexts field to given value.

### HasIncludedContexts

`func (o *UserSegment) HasIncludedContexts() bool`

HasIncludedContexts returns a boolean if a field has been set.

### GetExcludedContexts

`func (o *UserSegment) GetExcludedContexts() []SegmentTarget`

GetExcludedContexts returns the ExcludedContexts field if non-nil, zero value otherwise.

### GetExcludedContextsOk

`func (o *UserSegment) GetExcludedContextsOk() (*[]SegmentTarget, bool)`

GetExcludedContextsOk returns a tuple with the ExcludedContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedContexts

`func (o *UserSegment) SetExcludedContexts(v []SegmentTarget)`

SetExcludedContexts sets ExcludedContexts field to given value.

### HasExcludedContexts

`func (o *UserSegment) HasExcludedContexts() bool`

HasExcludedContexts returns a boolean if a field has been set.

### GetLinks

`func (o *UserSegment) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSegment) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSegment) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetRules

`func (o *UserSegment) GetRules() []UserSegmentRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UserSegment) GetRulesOk() (*[]UserSegmentRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UserSegment) SetRules(v []UserSegmentRule)`

SetRules sets Rules field to given value.


### GetVersion

`func (o *UserSegment) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserSegment) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserSegment) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDeleted

`func (o *UserSegment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UserSegment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UserSegment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetAccess

`func (o *UserSegment) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UserSegment) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UserSegment) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UserSegment) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetFlags

`func (o *UserSegment) GetFlags() []FlagListingRep`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UserSegment) GetFlagsOk() (*[]FlagListingRep, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UserSegment) SetFlags(v []FlagListingRep)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UserSegment) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetUnbounded

`func (o *UserSegment) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *UserSegment) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *UserSegment) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.

### HasUnbounded

`func (o *UserSegment) HasUnbounded() bool`

HasUnbounded returns a boolean if a field has been set.

### GetUnboundedContextKind

`func (o *UserSegment) GetUnboundedContextKind() string`

GetUnboundedContextKind returns the UnboundedContextKind field if non-nil, zero value otherwise.

### GetUnboundedContextKindOk

`func (o *UserSegment) GetUnboundedContextKindOk() (*string, bool)`

GetUnboundedContextKindOk returns a tuple with the UnboundedContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundedContextKind

`func (o *UserSegment) SetUnboundedContextKind(v string)`

SetUnboundedContextKind sets UnboundedContextKind field to given value.

### HasUnboundedContextKind

`func (o *UserSegment) HasUnboundedContextKind() bool`

HasUnboundedContextKind returns a boolean if a field has been set.

### GetGeneration

`func (o *UserSegment) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *UserSegment) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *UserSegment) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.


### GetUnboundedMetadata

`func (o *UserSegment) GetUnboundedMetadata() SegmentMetadata`

GetUnboundedMetadata returns the UnboundedMetadata field if non-nil, zero value otherwise.

### GetUnboundedMetadataOk

`func (o *UserSegment) GetUnboundedMetadataOk() (*SegmentMetadata, bool)`

GetUnboundedMetadataOk returns a tuple with the UnboundedMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundedMetadata

`func (o *UserSegment) SetUnboundedMetadata(v SegmentMetadata)`

SetUnboundedMetadata sets UnboundedMetadata field to given value.

### HasUnboundedMetadata

`func (o *UserSegment) HasUnboundedMetadata() bool`

HasUnboundedMetadata returns a boolean if a field has been set.

### GetExternal

`func (o *UserSegment) GetExternal() string`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *UserSegment) GetExternalOk() (*string, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *UserSegment) SetExternal(v string)`

SetExternal sets External field to given value.

### HasExternal

`func (o *UserSegment) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetExternalLink

`func (o *UserSegment) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *UserSegment) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *UserSegment) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *UserSegment) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetImportInProgress

`func (o *UserSegment) GetImportInProgress() bool`

GetImportInProgress returns the ImportInProgress field if non-nil, zero value otherwise.

### GetImportInProgressOk

`func (o *UserSegment) GetImportInProgressOk() (*bool, bool)`

GetImportInProgressOk returns a tuple with the ImportInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportInProgress

`func (o *UserSegment) SetImportInProgress(v bool)`

SetImportInProgress sets ImportInProgress field to given value.

### HasImportInProgress

`func (o *UserSegment) HasImportInProgress() bool`

HasImportInProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


