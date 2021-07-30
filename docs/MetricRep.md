# MetricRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**AttachedFlagCount** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 
**Site** | Pointer to [**CoreLink**](CoreLink.md) |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**LastModified** | Pointer to [**GoalsModification**](GoalsModification.md) |  | [optional] 
**MaintainerId** | Pointer to **string** |  | [optional] 
**Maintainer** | Pointer to [**MemberSummaryRep**](MemberSummaryRep.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsNumeric** | Pointer to **bool** |  | [optional] 
**SuccessCriteria** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**EventKey** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**AttachedFeatures** | Pointer to [**[]FlagListingRep**](FlagListingRep.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Selector** | Pointer to **string** |  | [optional] 
**Urls** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewMetricRep

`func NewMetricRep() *MetricRep`

NewMetricRep instantiates a new MetricRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRepWithDefaults

`func NewMetricRepWithDefaults() *MetricRep`

NewMetricRepWithDefaults instantiates a new MetricRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *MetricRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *MetricRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *MetricRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetricRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAttachedFlagCount

`func (o *MetricRep) GetAttachedFlagCount() int32`

GetAttachedFlagCount returns the AttachedFlagCount field if non-nil, zero value otherwise.

### GetAttachedFlagCountOk

`func (o *MetricRep) GetAttachedFlagCountOk() (*int32, bool)`

GetAttachedFlagCountOk returns a tuple with the AttachedFlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFlagCount

`func (o *MetricRep) SetAttachedFlagCount(v int32)`

SetAttachedFlagCount sets AttachedFlagCount field to given value.

### HasAttachedFlagCount

`func (o *MetricRep) HasAttachedFlagCount() bool`

HasAttachedFlagCount returns a boolean if a field has been set.

### GetLinks

`func (o *MetricRep) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricRep) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricRep) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MetricRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSite

`func (o *MetricRep) GetSite() CoreLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MetricRep) GetSiteOk() (*CoreLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MetricRep) SetSite(v CoreLink)`

SetSite sets Site field to given value.

### HasSite

`func (o *MetricRep) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccess

`func (o *MetricRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MetricRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MetricRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MetricRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTags

`func (o *MetricRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricRep) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreationDate

`func (o *MetricRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MetricRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MetricRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *MetricRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModified

`func (o *MetricRep) GetLastModified() GoalsModification`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *MetricRep) GetLastModifiedOk() (*GoalsModification, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *MetricRep) SetLastModified(v GoalsModification)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *MetricRep) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMaintainerId

`func (o *MetricRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *MetricRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *MetricRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *MetricRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *MetricRep) GetMaintainer() MemberSummaryRep`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *MetricRep) GetMaintainerOk() (*MemberSummaryRep, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *MetricRep) SetMaintainer(v MemberSummaryRep)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *MetricRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *MetricRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MetricRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *MetricRep) GetSuccessCriteria() int32`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricRep) GetSuccessCriteriaOk() (*int32, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricRep) SetSuccessCriteria(v int32)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *MetricRep) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetUnit

`func (o *MetricRep) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricRep) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricRep) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricRep) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *MetricRep) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *MetricRep) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *MetricRep) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *MetricRep) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetIsActive

`func (o *MetricRep) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MetricRep) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MetricRep) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MetricRep) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetAttachedFeatures

`func (o *MetricRep) GetAttachedFeatures() []FlagListingRep`

GetAttachedFeatures returns the AttachedFeatures field if non-nil, zero value otherwise.

### GetAttachedFeaturesOk

`func (o *MetricRep) GetAttachedFeaturesOk() (*[]FlagListingRep, bool)`

GetAttachedFeaturesOk returns a tuple with the AttachedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFeatures

`func (o *MetricRep) SetAttachedFeatures(v []FlagListingRep)`

SetAttachedFeatures sets AttachedFeatures field to given value.

### HasAttachedFeatures

`func (o *MetricRep) HasAttachedFeatures() bool`

HasAttachedFeatures returns a boolean if a field has been set.

### GetVersion

`func (o *MetricRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetricRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetricRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetricRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSelector

`func (o *MetricRep) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MetricRep) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MetricRep) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MetricRep) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *MetricRep) GetUrls() []interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MetricRep) GetUrlsOk() (*[]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MetricRep) SetUrls(v []interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MetricRep) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


