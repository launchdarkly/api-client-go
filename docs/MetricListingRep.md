# MetricListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Kind** | **string** |  | 
**AttachedFlagCount** | Pointer to **int32** |  | [optional] 
**Links** | [**map[string]Link**](Link.md) |  | 
**Site** | Pointer to [**Link**](Link.md) |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**Tags** | **[]string** |  | 
**CreationDate** | **int64** |  | 
**LastModified** | Pointer to [**Modification**](Modification.md) |  | [optional] 
**MaintainerId** | Pointer to **string** |  | [optional] 
**Maintainer** | Pointer to [**MemberSummaryRep**](MemberSummaryRep.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsNumeric** | Pointer to **bool** |  | [optional] 
**SuccessCriteria** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**EventKey** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricListingRep

`func NewMetricListingRep(id string, key string, name string, kind string, links map[string]Link, tags []string, creationDate int64, ) *MetricListingRep`

NewMetricListingRep instantiates a new MetricListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricListingRepWithDefaults

`func NewMetricListingRepWithDefaults() *MetricListingRep`

NewMetricListingRepWithDefaults instantiates a new MetricListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricListingRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricListingRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricListingRep) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *MetricListingRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricListingRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricListingRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MetricListingRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricListingRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricListingRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *MetricListingRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricListingRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricListingRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAttachedFlagCount

`func (o *MetricListingRep) GetAttachedFlagCount() int32`

GetAttachedFlagCount returns the AttachedFlagCount field if non-nil, zero value otherwise.

### GetAttachedFlagCountOk

`func (o *MetricListingRep) GetAttachedFlagCountOk() (*int32, bool)`

GetAttachedFlagCountOk returns a tuple with the AttachedFlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFlagCount

`func (o *MetricListingRep) SetAttachedFlagCount(v int32)`

SetAttachedFlagCount sets AttachedFlagCount field to given value.

### HasAttachedFlagCount

`func (o *MetricListingRep) HasAttachedFlagCount() bool`

HasAttachedFlagCount returns a boolean if a field has been set.

### GetLinks

`func (o *MetricListingRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricListingRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricListingRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetSite

`func (o *MetricListingRep) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MetricListingRep) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MetricListingRep) SetSite(v Link)`

SetSite sets Site field to given value.

### HasSite

`func (o *MetricListingRep) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccess

`func (o *MetricListingRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MetricListingRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MetricListingRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MetricListingRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTags

`func (o *MetricListingRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricListingRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricListingRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreationDate

`func (o *MetricListingRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MetricListingRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MetricListingRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *MetricListingRep) GetLastModified() Modification`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *MetricListingRep) GetLastModifiedOk() (*Modification, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *MetricListingRep) SetLastModified(v Modification)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *MetricListingRep) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMaintainerId

`func (o *MetricListingRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *MetricListingRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *MetricListingRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *MetricListingRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *MetricListingRep) GetMaintainer() MemberSummaryRep`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *MetricListingRep) GetMaintainerOk() (*MemberSummaryRep, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *MetricListingRep) SetMaintainer(v MemberSummaryRep)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *MetricListingRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *MetricListingRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricListingRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricListingRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricListingRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MetricListingRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricListingRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricListingRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricListingRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *MetricListingRep) GetSuccessCriteria() int32`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricListingRep) GetSuccessCriteriaOk() (*int32, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricListingRep) SetSuccessCriteria(v int32)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *MetricListingRep) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetUnit

`func (o *MetricListingRep) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricListingRep) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricListingRep) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricListingRep) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *MetricListingRep) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *MetricListingRep) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *MetricListingRep) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *MetricListingRep) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


