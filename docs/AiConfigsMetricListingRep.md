# AiConfigsMetricListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentCount** | Pointer to **int32** | The number of experiments using this metric | [optional] 
**MetricGroupCount** | Pointer to **int32** | The number of metric groups using this metric | [optional] 
**GuardedRolloutCount** | Pointer to **int32** | The number of guarded rollouts using this metric | [optional] 
**ActiveExperimentCount** | Pointer to **int32** | The number of active experiments using this metric | [optional] 
**ActiveGuardedRolloutCount** | Pointer to **int32** | The number of active guarded rollouts using this metric | [optional] 
**Id** | **string** | The ID of this metric | 
**VersionId** | **string** | The version ID of the metric | 
**Version** | Pointer to **int32** | Version of the metric | [optional] 
**Key** | **string** | A unique key to reference the metric | 
**Name** | **string** | A human-friendly name for the metric | 
**Kind** | **string** | The kind of event the metric tracks | 
**AttachedFlagCount** | Pointer to **int32** | The number of feature flags currently attached to this metric | [optional] 
**Links** | [**map[string]AiConfigsLink**](AiConfigsLink.md) | The location and content type of related resources | 
**Site** | Pointer to [**AiConfigsLink**](AiConfigsLink.md) |  | [optional] 
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Tags** | **[]string** | Tags for the metric | 
**CreationDate** | **int64** |  | 
**LastModified** | Pointer to [**AiConfigsModification**](AiConfigsModification.md) |  | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this metric | [optional] 
**Maintainer** | Pointer to [**AiConfigsMemberSummary**](AiConfigsMemberSummary.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the metric | [optional] 
**Category** | Pointer to **string** | The category of the metric | [optional] 
**IsNumeric** | Pointer to **bool** | For custom metrics, whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). | [optional] 
**SuccessCriteria** | Pointer to **string** | For custom metrics, the success criteria | [optional] 
**Unit** | Pointer to **string** | For numeric custom metrics, the unit of measure | [optional] 
**EventKey** | Pointer to **string** | For custom metrics, the event key to use in your code | [optional] 
**RandomizationUnits** | Pointer to **[]string** | An array of randomization units allowed for this metric | [optional] 
**Filters** | Pointer to [**AiConfigsFilter**](AiConfigsFilter.md) |  | [optional] 
**UnitAggregationType** | Pointer to **string** | The method by which multiple unit event values are aggregated | [optional] 
**AnalysisType** | Pointer to **string** | The method for analyzing metric events | [optional] 
**PercentileValue** | Pointer to **int32** | The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when &lt;code&gt;analysisType&lt;/code&gt; is &lt;code&gt;percentile&lt;/code&gt;. | [optional] 
**EventDefault** | Pointer to [**AiConfigsMetricEventDefaultRep**](AiConfigsMetricEventDefaultRep.md) |  | [optional] 
**DataSource** | Pointer to [**AiConfigsMetricDataSourceRefRep**](AiConfigsMetricDataSourceRefRep.md) |  | [optional] 
**Archived** | Pointer to **bool** | Whether the metric version is archived | [optional] 
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**Selector** | Pointer to **string** | For click metrics, the CSS selectors | [optional] 
**Urls** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewAiConfigsMetricListingRep

`func NewAiConfigsMetricListingRep(id string, versionId string, key string, name string, kind string, links map[string]AiConfigsLink, tags []string, creationDate int64, ) *AiConfigsMetricListingRep`

NewAiConfigsMetricListingRep instantiates a new AiConfigsMetricListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsMetricListingRepWithDefaults

`func NewAiConfigsMetricListingRepWithDefaults() *AiConfigsMetricListingRep`

NewAiConfigsMetricListingRepWithDefaults instantiates a new AiConfigsMetricListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentCount

`func (o *AiConfigsMetricListingRep) GetExperimentCount() int32`

GetExperimentCount returns the ExperimentCount field if non-nil, zero value otherwise.

### GetExperimentCountOk

`func (o *AiConfigsMetricListingRep) GetExperimentCountOk() (*int32, bool)`

GetExperimentCountOk returns a tuple with the ExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentCount

`func (o *AiConfigsMetricListingRep) SetExperimentCount(v int32)`

SetExperimentCount sets ExperimentCount field to given value.

### HasExperimentCount

`func (o *AiConfigsMetricListingRep) HasExperimentCount() bool`

HasExperimentCount returns a boolean if a field has been set.

### GetMetricGroupCount

`func (o *AiConfigsMetricListingRep) GetMetricGroupCount() int32`

GetMetricGroupCount returns the MetricGroupCount field if non-nil, zero value otherwise.

### GetMetricGroupCountOk

`func (o *AiConfigsMetricListingRep) GetMetricGroupCountOk() (*int32, bool)`

GetMetricGroupCountOk returns a tuple with the MetricGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroupCount

`func (o *AiConfigsMetricListingRep) SetMetricGroupCount(v int32)`

SetMetricGroupCount sets MetricGroupCount field to given value.

### HasMetricGroupCount

`func (o *AiConfigsMetricListingRep) HasMetricGroupCount() bool`

HasMetricGroupCount returns a boolean if a field has been set.

### GetGuardedRolloutCount

`func (o *AiConfigsMetricListingRep) GetGuardedRolloutCount() int32`

GetGuardedRolloutCount returns the GuardedRolloutCount field if non-nil, zero value otherwise.

### GetGuardedRolloutCountOk

`func (o *AiConfigsMetricListingRep) GetGuardedRolloutCountOk() (*int32, bool)`

GetGuardedRolloutCountOk returns a tuple with the GuardedRolloutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardedRolloutCount

`func (o *AiConfigsMetricListingRep) SetGuardedRolloutCount(v int32)`

SetGuardedRolloutCount sets GuardedRolloutCount field to given value.

### HasGuardedRolloutCount

`func (o *AiConfigsMetricListingRep) HasGuardedRolloutCount() bool`

HasGuardedRolloutCount returns a boolean if a field has been set.

### GetActiveExperimentCount

`func (o *AiConfigsMetricListingRep) GetActiveExperimentCount() int32`

GetActiveExperimentCount returns the ActiveExperimentCount field if non-nil, zero value otherwise.

### GetActiveExperimentCountOk

`func (o *AiConfigsMetricListingRep) GetActiveExperimentCountOk() (*int32, bool)`

GetActiveExperimentCountOk returns a tuple with the ActiveExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExperimentCount

`func (o *AiConfigsMetricListingRep) SetActiveExperimentCount(v int32)`

SetActiveExperimentCount sets ActiveExperimentCount field to given value.

### HasActiveExperimentCount

`func (o *AiConfigsMetricListingRep) HasActiveExperimentCount() bool`

HasActiveExperimentCount returns a boolean if a field has been set.

### GetActiveGuardedRolloutCount

`func (o *AiConfigsMetricListingRep) GetActiveGuardedRolloutCount() int32`

GetActiveGuardedRolloutCount returns the ActiveGuardedRolloutCount field if non-nil, zero value otherwise.

### GetActiveGuardedRolloutCountOk

`func (o *AiConfigsMetricListingRep) GetActiveGuardedRolloutCountOk() (*int32, bool)`

GetActiveGuardedRolloutCountOk returns a tuple with the ActiveGuardedRolloutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGuardedRolloutCount

`func (o *AiConfigsMetricListingRep) SetActiveGuardedRolloutCount(v int32)`

SetActiveGuardedRolloutCount sets ActiveGuardedRolloutCount field to given value.

### HasActiveGuardedRolloutCount

`func (o *AiConfigsMetricListingRep) HasActiveGuardedRolloutCount() bool`

HasActiveGuardedRolloutCount returns a boolean if a field has been set.

### GetId

`func (o *AiConfigsMetricListingRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiConfigsMetricListingRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiConfigsMetricListingRep) SetId(v string)`

SetId sets Id field to given value.


### GetVersionId

`func (o *AiConfigsMetricListingRep) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AiConfigsMetricListingRep) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AiConfigsMetricListingRep) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersion

`func (o *AiConfigsMetricListingRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AiConfigsMetricListingRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AiConfigsMetricListingRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AiConfigsMetricListingRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetKey

`func (o *AiConfigsMetricListingRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiConfigsMetricListingRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiConfigsMetricListingRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AiConfigsMetricListingRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiConfigsMetricListingRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiConfigsMetricListingRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *AiConfigsMetricListingRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiConfigsMetricListingRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiConfigsMetricListingRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAttachedFlagCount

`func (o *AiConfigsMetricListingRep) GetAttachedFlagCount() int32`

GetAttachedFlagCount returns the AttachedFlagCount field if non-nil, zero value otherwise.

### GetAttachedFlagCountOk

`func (o *AiConfigsMetricListingRep) GetAttachedFlagCountOk() (*int32, bool)`

GetAttachedFlagCountOk returns a tuple with the AttachedFlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFlagCount

`func (o *AiConfigsMetricListingRep) SetAttachedFlagCount(v int32)`

SetAttachedFlagCount sets AttachedFlagCount field to given value.

### HasAttachedFlagCount

`func (o *AiConfigsMetricListingRep) HasAttachedFlagCount() bool`

HasAttachedFlagCount returns a boolean if a field has been set.

### GetLinks

`func (o *AiConfigsMetricListingRep) GetLinks() map[string]AiConfigsLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AiConfigsMetricListingRep) GetLinksOk() (*map[string]AiConfigsLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AiConfigsMetricListingRep) SetLinks(v map[string]AiConfigsLink)`

SetLinks sets Links field to given value.


### GetSite

`func (o *AiConfigsMetricListingRep) GetSite() AiConfigsLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AiConfigsMetricListingRep) GetSiteOk() (*AiConfigsLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AiConfigsMetricListingRep) SetSite(v AiConfigsLink)`

SetSite sets Site field to given value.

### HasSite

`func (o *AiConfigsMetricListingRep) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccess

`func (o *AiConfigsMetricListingRep) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AiConfigsMetricListingRep) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AiConfigsMetricListingRep) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AiConfigsMetricListingRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTags

`func (o *AiConfigsMetricListingRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AiConfigsMetricListingRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AiConfigsMetricListingRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreationDate

`func (o *AiConfigsMetricListingRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AiConfigsMetricListingRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AiConfigsMetricListingRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *AiConfigsMetricListingRep) GetLastModified() AiConfigsModification`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AiConfigsMetricListingRep) GetLastModifiedOk() (*AiConfigsModification, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AiConfigsMetricListingRep) SetLastModified(v AiConfigsModification)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AiConfigsMetricListingRep) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMaintainerId

`func (o *AiConfigsMetricListingRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AiConfigsMetricListingRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AiConfigsMetricListingRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AiConfigsMetricListingRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *AiConfigsMetricListingRep) GetMaintainer() AiConfigsMemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *AiConfigsMetricListingRep) GetMaintainerOk() (*AiConfigsMemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *AiConfigsMetricListingRep) SetMaintainer(v AiConfigsMemberSummary)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *AiConfigsMetricListingRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *AiConfigsMetricListingRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiConfigsMetricListingRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiConfigsMetricListingRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AiConfigsMetricListingRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *AiConfigsMetricListingRep) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AiConfigsMetricListingRep) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AiConfigsMetricListingRep) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AiConfigsMetricListingRep) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIsNumeric

`func (o *AiConfigsMetricListingRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *AiConfigsMetricListingRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *AiConfigsMetricListingRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *AiConfigsMetricListingRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *AiConfigsMetricListingRep) GetSuccessCriteria() string`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *AiConfigsMetricListingRep) GetSuccessCriteriaOk() (*string, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *AiConfigsMetricListingRep) SetSuccessCriteria(v string)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *AiConfigsMetricListingRep) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetUnit

`func (o *AiConfigsMetricListingRep) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AiConfigsMetricListingRep) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AiConfigsMetricListingRep) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AiConfigsMetricListingRep) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *AiConfigsMetricListingRep) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *AiConfigsMetricListingRep) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *AiConfigsMetricListingRep) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *AiConfigsMetricListingRep) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetRandomizationUnits

`func (o *AiConfigsMetricListingRep) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *AiConfigsMetricListingRep) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *AiConfigsMetricListingRep) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *AiConfigsMetricListingRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetFilters

`func (o *AiConfigsMetricListingRep) GetFilters() AiConfigsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AiConfigsMetricListingRep) GetFiltersOk() (*AiConfigsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AiConfigsMetricListingRep) SetFilters(v AiConfigsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AiConfigsMetricListingRep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *AiConfigsMetricListingRep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *AiConfigsMetricListingRep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *AiConfigsMetricListingRep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *AiConfigsMetricListingRep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetAnalysisType

`func (o *AiConfigsMetricListingRep) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *AiConfigsMetricListingRep) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *AiConfigsMetricListingRep) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.

### HasAnalysisType

`func (o *AiConfigsMetricListingRep) HasAnalysisType() bool`

HasAnalysisType returns a boolean if a field has been set.

### GetPercentileValue

`func (o *AiConfigsMetricListingRep) GetPercentileValue() int32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *AiConfigsMetricListingRep) GetPercentileValueOk() (*int32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *AiConfigsMetricListingRep) SetPercentileValue(v int32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *AiConfigsMetricListingRep) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetEventDefault

`func (o *AiConfigsMetricListingRep) GetEventDefault() AiConfigsMetricEventDefaultRep`

GetEventDefault returns the EventDefault field if non-nil, zero value otherwise.

### GetEventDefaultOk

`func (o *AiConfigsMetricListingRep) GetEventDefaultOk() (*AiConfigsMetricEventDefaultRep, bool)`

GetEventDefaultOk returns a tuple with the EventDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDefault

`func (o *AiConfigsMetricListingRep) SetEventDefault(v AiConfigsMetricEventDefaultRep)`

SetEventDefault sets EventDefault field to given value.

### HasEventDefault

`func (o *AiConfigsMetricListingRep) HasEventDefault() bool`

HasEventDefault returns a boolean if a field has been set.

### GetDataSource

`func (o *AiConfigsMetricListingRep) GetDataSource() AiConfigsMetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *AiConfigsMetricListingRep) GetDataSourceOk() (*AiConfigsMetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *AiConfigsMetricListingRep) SetDataSource(v AiConfigsMetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *AiConfigsMetricListingRep) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetArchived

`func (o *AiConfigsMetricListingRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AiConfigsMetricListingRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AiConfigsMetricListingRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AiConfigsMetricListingRep) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *AiConfigsMetricListingRep) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *AiConfigsMetricListingRep) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *AiConfigsMetricListingRep) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *AiConfigsMetricListingRep) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetSelector

`func (o *AiConfigsMetricListingRep) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *AiConfigsMetricListingRep) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *AiConfigsMetricListingRep) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *AiConfigsMetricListingRep) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *AiConfigsMetricListingRep) GetUrls() []map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *AiConfigsMetricListingRep) GetUrlsOk() (*[]map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *AiConfigsMetricListingRep) SetUrls(v []map[string]interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *AiConfigsMetricListingRep) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


