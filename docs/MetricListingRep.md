# MetricListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentCount** | Pointer to **int32** | The number of experiments using this metric | [optional] 
**MetricGroupCount** | Pointer to **int32** | The number of metric groups using this metric | [optional] 
**ActiveExperimentCount** | Pointer to **int32** | The number of active experiments using this metric | [optional] 
**ActiveGuardedRolloutCount** | Pointer to **int32** | The number of active guarded rollouts using this metric | [optional] 
**Id** | **string** | The ID of this metric | 
**VersionId** | **string** | The version ID of the metric | 
**Version** | Pointer to **int32** | Version of the metric | [optional] 
**Key** | **string** | A unique key to reference the metric | 
**Name** | **string** | A human-friendly name for the metric | 
**Kind** | **string** | The kind of event the metric tracks | 
**AttachedFlagCount** | Pointer to **int32** | The number of feature flags currently attached to this metric | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Site** | Pointer to [**Link**](Link.md) |  | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Tags** | **[]string** | Tags for the metric | 
**CreationDate** | **int64** |  | 
**LastModified** | Pointer to [**Modification**](Modification.md) |  | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this metric | [optional] 
**Maintainer** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the metric | [optional] 
**Category** | Pointer to **string** | The category of the metric | [optional] 
**IsNumeric** | Pointer to **bool** | For custom metrics, whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). | [optional] 
**SuccessCriteria** | Pointer to **string** | For custom metrics, the success criteria | [optional] 
**Unit** | Pointer to **string** | For numeric custom metrics, the unit of measure | [optional] 
**EventKey** | Pointer to **string** | For custom metrics, the event key to use in your code | [optional] 
**RandomizationUnits** | Pointer to **[]string** | An array of randomization units allowed for this metric | [optional] 
**Filters** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**UnitAggregationType** | Pointer to **string** | The method by which multiple unit event values are aggregated | [optional] 
**AnalysisType** | Pointer to **string** | The method for analyzing metric events | [optional] 
**PercentileValue** | Pointer to **int32** | The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when &lt;code&gt;analysisType&lt;/code&gt; is &lt;code&gt;percentile&lt;/code&gt;. | [optional] 
**EventDefault** | Pointer to [**MetricEventDefaultRep**](MetricEventDefaultRep.md) |  | [optional] 
**DataSource** | Pointer to [**MetricDataSourceRefRep**](MetricDataSourceRefRep.md) |  | [optional] 
**Archived** | Pointer to **bool** | Whether the metric version is archived | [optional] 
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**Selector** | Pointer to **string** | For click metrics, the CSS selectors | [optional] 
**Urls** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewMetricListingRep

`func NewMetricListingRep(id string, versionId string, key string, name string, kind string, links map[string]Link, tags []string, creationDate int64, ) *MetricListingRep`

NewMetricListingRep instantiates a new MetricListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricListingRepWithDefaults

`func NewMetricListingRepWithDefaults() *MetricListingRep`

NewMetricListingRepWithDefaults instantiates a new MetricListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentCount

`func (o *MetricListingRep) GetExperimentCount() int32`

GetExperimentCount returns the ExperimentCount field if non-nil, zero value otherwise.

### GetExperimentCountOk

`func (o *MetricListingRep) GetExperimentCountOk() (*int32, bool)`

GetExperimentCountOk returns a tuple with the ExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentCount

`func (o *MetricListingRep) SetExperimentCount(v int32)`

SetExperimentCount sets ExperimentCount field to given value.

### HasExperimentCount

`func (o *MetricListingRep) HasExperimentCount() bool`

HasExperimentCount returns a boolean if a field has been set.

### GetMetricGroupCount

`func (o *MetricListingRep) GetMetricGroupCount() int32`

GetMetricGroupCount returns the MetricGroupCount field if non-nil, zero value otherwise.

### GetMetricGroupCountOk

`func (o *MetricListingRep) GetMetricGroupCountOk() (*int32, bool)`

GetMetricGroupCountOk returns a tuple with the MetricGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroupCount

`func (o *MetricListingRep) SetMetricGroupCount(v int32)`

SetMetricGroupCount sets MetricGroupCount field to given value.

### HasMetricGroupCount

`func (o *MetricListingRep) HasMetricGroupCount() bool`

HasMetricGroupCount returns a boolean if a field has been set.

### GetActiveExperimentCount

`func (o *MetricListingRep) GetActiveExperimentCount() int32`

GetActiveExperimentCount returns the ActiveExperimentCount field if non-nil, zero value otherwise.

### GetActiveExperimentCountOk

`func (o *MetricListingRep) GetActiveExperimentCountOk() (*int32, bool)`

GetActiveExperimentCountOk returns a tuple with the ActiveExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExperimentCount

`func (o *MetricListingRep) SetActiveExperimentCount(v int32)`

SetActiveExperimentCount sets ActiveExperimentCount field to given value.

### HasActiveExperimentCount

`func (o *MetricListingRep) HasActiveExperimentCount() bool`

HasActiveExperimentCount returns a boolean if a field has been set.

### GetActiveGuardedRolloutCount

`func (o *MetricListingRep) GetActiveGuardedRolloutCount() int32`

GetActiveGuardedRolloutCount returns the ActiveGuardedRolloutCount field if non-nil, zero value otherwise.

### GetActiveGuardedRolloutCountOk

`func (o *MetricListingRep) GetActiveGuardedRolloutCountOk() (*int32, bool)`

GetActiveGuardedRolloutCountOk returns a tuple with the ActiveGuardedRolloutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGuardedRolloutCount

`func (o *MetricListingRep) SetActiveGuardedRolloutCount(v int32)`

SetActiveGuardedRolloutCount sets ActiveGuardedRolloutCount field to given value.

### HasActiveGuardedRolloutCount

`func (o *MetricListingRep) HasActiveGuardedRolloutCount() bool`

HasActiveGuardedRolloutCount returns a boolean if a field has been set.

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


### GetVersionId

`func (o *MetricListingRep) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *MetricListingRep) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *MetricListingRep) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersion

`func (o *MetricListingRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetricListingRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetricListingRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetricListingRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

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

`func (o *MetricListingRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MetricListingRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MetricListingRep) SetAccess(v Access)`

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

`func (o *MetricListingRep) GetMaintainer() MemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *MetricListingRep) GetMaintainerOk() (*MemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *MetricListingRep) SetMaintainer(v MemberSummary)`

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

### GetCategory

`func (o *MetricListingRep) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MetricListingRep) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MetricListingRep) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MetricListingRep) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

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

`func (o *MetricListingRep) GetSuccessCriteria() string`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricListingRep) GetSuccessCriteriaOk() (*string, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricListingRep) SetSuccessCriteria(v string)`

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

### GetRandomizationUnits

`func (o *MetricListingRep) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *MetricListingRep) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *MetricListingRep) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *MetricListingRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetFilters

`func (o *MetricListingRep) GetFilters() Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricListingRep) GetFiltersOk() (*Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricListingRep) SetFilters(v Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricListingRep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *MetricListingRep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *MetricListingRep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *MetricListingRep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *MetricListingRep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetAnalysisType

`func (o *MetricListingRep) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *MetricListingRep) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *MetricListingRep) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.

### HasAnalysisType

`func (o *MetricListingRep) HasAnalysisType() bool`

HasAnalysisType returns a boolean if a field has been set.

### GetPercentileValue

`func (o *MetricListingRep) GetPercentileValue() int32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *MetricListingRep) GetPercentileValueOk() (*int32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *MetricListingRep) SetPercentileValue(v int32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *MetricListingRep) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetEventDefault

`func (o *MetricListingRep) GetEventDefault() MetricEventDefaultRep`

GetEventDefault returns the EventDefault field if non-nil, zero value otherwise.

### GetEventDefaultOk

`func (o *MetricListingRep) GetEventDefaultOk() (*MetricEventDefaultRep, bool)`

GetEventDefaultOk returns a tuple with the EventDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDefault

`func (o *MetricListingRep) SetEventDefault(v MetricEventDefaultRep)`

SetEventDefault sets EventDefault field to given value.

### HasEventDefault

`func (o *MetricListingRep) HasEventDefault() bool`

HasEventDefault returns a boolean if a field has been set.

### GetDataSource

`func (o *MetricListingRep) GetDataSource() MetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *MetricListingRep) GetDataSourceOk() (*MetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *MetricListingRep) SetDataSource(v MetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *MetricListingRep) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetArchived

`func (o *MetricListingRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *MetricListingRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *MetricListingRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *MetricListingRep) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *MetricListingRep) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *MetricListingRep) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *MetricListingRep) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *MetricListingRep) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetSelector

`func (o *MetricListingRep) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MetricListingRep) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MetricListingRep) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MetricListingRep) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *MetricListingRep) GetUrls() []map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MetricListingRep) GetUrlsOk() (*[]map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MetricListingRep) SetUrls(v []map[string]interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MetricListingRep) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


