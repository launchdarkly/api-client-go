# EnvironmentReleasesMetricListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentCount** | Pointer to **int32** | The number of experiments using this metric | [optional] 
**MetricGroupCount** | Pointer to **int32** | The number of metric groups using this metric | [optional] 
**GuardedRolloutCount** | Pointer to **int32** | The number of guarded rollouts using this metric | [optional] 
**ReleasePolicyCount** | Pointer to **int32** | The number of release policies using this metric | [optional] 
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
**Access** | Pointer to [**EnvironmentReleasesAccess**](EnvironmentReleasesAccess.md) |  | [optional] 
**Tags** | **[]string** | Tags for the metric | 
**CreationDate** | **int64** |  | 
**LastModified** | Pointer to [**EnvironmentReleasesModification**](EnvironmentReleasesModification.md) |  | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this metric | [optional] 
**Maintainer** | Pointer to [**EnvironmentReleasesMemberSummary**](EnvironmentReleasesMemberSummary.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the metric | [optional] 
**Category** | Pointer to **string** | The category of the metric | [optional] 
**IsNumeric** | Pointer to **bool** | For custom and trace metrics, whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). | [optional] 
**SuccessCriteria** | Pointer to **string** | For custom and trace metrics, the success criteria | [optional] 
**Unit** | Pointer to **string** | For numeric custom and trace metrics, the unit of measure | [optional] 
**EventKey** | Pointer to **string** | For custom metrics, the event key to use in your code | [optional] 
**RandomizationUnits** | Pointer to **[]string** | Deprecated, use &lt;code&gt;analysisUnits&lt;/code&gt; instead. | [optional] 
**AnalysisUnits** | Pointer to **[]string** | An array of analysis units allowed for this metric. | [optional] 
**Filters** | Pointer to [**EnvironmentReleasesFilter**](EnvironmentReleasesFilter.md) |  | [optional] 
**UnitAggregationType** | Pointer to **string** | The method by which multiple unit event values are aggregated | [optional] 
**AnalysisType** | Pointer to **string** | The method for analyzing metric events | [optional] 
**PercentileValue** | Pointer to **int32** | The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when &lt;code&gt;analysisType&lt;/code&gt; is &lt;code&gt;percentile&lt;/code&gt;. | [optional] 
**EventDefault** | Pointer to [**MetricEventDefaultRep**](MetricEventDefaultRep.md) |  | [optional] 
**DataSource** | [**MetricDataSourceRefRep**](MetricDataSourceRefRep.md) |  | 
**LastSeen** | Pointer to **int64** |  | [optional] 
**Archived** | Pointer to **bool** | Whether the metric version is archived | [optional] 
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**Selector** | Pointer to **string** | For click metrics, the CSS selectors | [optional] 
**Urls** | Pointer to **[]map[string]interface{}** |  | [optional] 
**WindowStartOffset** | Pointer to **int64** | Not yet implemented - The start of the measurement window, in milliseconds relative to the unit&#39;s first exposure to a flag variation | [optional] 
**WindowEndOffset** | Pointer to **int64** | Not yet implemented - The end of the measurement window, in milliseconds relative to the unit&#39;s first exposure to a flag variation | [optional] 
**WinsorLowerPercentile** | Pointer to **float32** | Lower winsorization percentile, expressed as a percent in the open interval (0, 100). When both bounds are set, defines a two-sided clamp range. Otherwise lower-only winsorization. | [optional] 
**WinsorUpperPercentile** | Pointer to **float32** | Upper winsorization percentile, expressed as a percent in the open interval (0, 100). When both bounds are set, must be greater than winsorLowerPercentile. | [optional] 
**WinsorIncludeImputed** | Pointer to **bool** | When true, the percentile bound calculation includes imputed zeros. Only meaningful when at least one bound is set and the metric includes units that didn&#39;t send events. | [optional] 
**TraceQuery** | Pointer to **string** | For trace metrics, the trace query to use for the metric. | [optional] 
**TraceValueLocation** | Pointer to **string** | For trace metrics, the location in the trace to use for numeric values. | [optional] 
**UnitAggregationField** | Pointer to **string** | For count_distinct metrics, the column to count distinct values of | [optional] 
**Denominator** | Pointer to [**EnvironmentReleasesMetricDenominatorRep**](EnvironmentReleasesMetricDenominatorRep.md) |  | [optional] 

## Methods

### NewEnvironmentReleasesMetricListingRep

`func NewEnvironmentReleasesMetricListingRep(id string, versionId string, key string, name string, kind string, links map[string]Link, tags []string, creationDate int64, dataSource MetricDataSourceRefRep, ) *EnvironmentReleasesMetricListingRep`

NewEnvironmentReleasesMetricListingRep instantiates a new EnvironmentReleasesMetricListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasesMetricListingRepWithDefaults

`func NewEnvironmentReleasesMetricListingRepWithDefaults() *EnvironmentReleasesMetricListingRep`

NewEnvironmentReleasesMetricListingRepWithDefaults instantiates a new EnvironmentReleasesMetricListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentCount

`func (o *EnvironmentReleasesMetricListingRep) GetExperimentCount() int32`

GetExperimentCount returns the ExperimentCount field if non-nil, zero value otherwise.

### GetExperimentCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetExperimentCountOk() (*int32, bool)`

GetExperimentCountOk returns a tuple with the ExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentCount

`func (o *EnvironmentReleasesMetricListingRep) SetExperimentCount(v int32)`

SetExperimentCount sets ExperimentCount field to given value.

### HasExperimentCount

`func (o *EnvironmentReleasesMetricListingRep) HasExperimentCount() bool`

HasExperimentCount returns a boolean if a field has been set.

### GetMetricGroupCount

`func (o *EnvironmentReleasesMetricListingRep) GetMetricGroupCount() int32`

GetMetricGroupCount returns the MetricGroupCount field if non-nil, zero value otherwise.

### GetMetricGroupCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetMetricGroupCountOk() (*int32, bool)`

GetMetricGroupCountOk returns a tuple with the MetricGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricGroupCount

`func (o *EnvironmentReleasesMetricListingRep) SetMetricGroupCount(v int32)`

SetMetricGroupCount sets MetricGroupCount field to given value.

### HasMetricGroupCount

`func (o *EnvironmentReleasesMetricListingRep) HasMetricGroupCount() bool`

HasMetricGroupCount returns a boolean if a field has been set.

### GetGuardedRolloutCount

`func (o *EnvironmentReleasesMetricListingRep) GetGuardedRolloutCount() int32`

GetGuardedRolloutCount returns the GuardedRolloutCount field if non-nil, zero value otherwise.

### GetGuardedRolloutCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetGuardedRolloutCountOk() (*int32, bool)`

GetGuardedRolloutCountOk returns a tuple with the GuardedRolloutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardedRolloutCount

`func (o *EnvironmentReleasesMetricListingRep) SetGuardedRolloutCount(v int32)`

SetGuardedRolloutCount sets GuardedRolloutCount field to given value.

### HasGuardedRolloutCount

`func (o *EnvironmentReleasesMetricListingRep) HasGuardedRolloutCount() bool`

HasGuardedRolloutCount returns a boolean if a field has been set.

### GetReleasePolicyCount

`func (o *EnvironmentReleasesMetricListingRep) GetReleasePolicyCount() int32`

GetReleasePolicyCount returns the ReleasePolicyCount field if non-nil, zero value otherwise.

### GetReleasePolicyCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetReleasePolicyCountOk() (*int32, bool)`

GetReleasePolicyCountOk returns a tuple with the ReleasePolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePolicyCount

`func (o *EnvironmentReleasesMetricListingRep) SetReleasePolicyCount(v int32)`

SetReleasePolicyCount sets ReleasePolicyCount field to given value.

### HasReleasePolicyCount

`func (o *EnvironmentReleasesMetricListingRep) HasReleasePolicyCount() bool`

HasReleasePolicyCount returns a boolean if a field has been set.

### GetActiveExperimentCount

`func (o *EnvironmentReleasesMetricListingRep) GetActiveExperimentCount() int32`

GetActiveExperimentCount returns the ActiveExperimentCount field if non-nil, zero value otherwise.

### GetActiveExperimentCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetActiveExperimentCountOk() (*int32, bool)`

GetActiveExperimentCountOk returns a tuple with the ActiveExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExperimentCount

`func (o *EnvironmentReleasesMetricListingRep) SetActiveExperimentCount(v int32)`

SetActiveExperimentCount sets ActiveExperimentCount field to given value.

### HasActiveExperimentCount

`func (o *EnvironmentReleasesMetricListingRep) HasActiveExperimentCount() bool`

HasActiveExperimentCount returns a boolean if a field has been set.

### GetActiveGuardedRolloutCount

`func (o *EnvironmentReleasesMetricListingRep) GetActiveGuardedRolloutCount() int32`

GetActiveGuardedRolloutCount returns the ActiveGuardedRolloutCount field if non-nil, zero value otherwise.

### GetActiveGuardedRolloutCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetActiveGuardedRolloutCountOk() (*int32, bool)`

GetActiveGuardedRolloutCountOk returns a tuple with the ActiveGuardedRolloutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGuardedRolloutCount

`func (o *EnvironmentReleasesMetricListingRep) SetActiveGuardedRolloutCount(v int32)`

SetActiveGuardedRolloutCount sets ActiveGuardedRolloutCount field to given value.

### HasActiveGuardedRolloutCount

`func (o *EnvironmentReleasesMetricListingRep) HasActiveGuardedRolloutCount() bool`

HasActiveGuardedRolloutCount returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentReleasesMetricListingRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReleasesMetricListingRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReleasesMetricListingRep) SetId(v string)`

SetId sets Id field to given value.


### GetVersionId

`func (o *EnvironmentReleasesMetricListingRep) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EnvironmentReleasesMetricListingRep) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EnvironmentReleasesMetricListingRep) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersion

`func (o *EnvironmentReleasesMetricListingRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentReleasesMetricListingRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentReleasesMetricListingRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnvironmentReleasesMetricListingRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetKey

`func (o *EnvironmentReleasesMetricListingRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentReleasesMetricListingRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentReleasesMetricListingRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *EnvironmentReleasesMetricListingRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentReleasesMetricListingRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentReleasesMetricListingRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *EnvironmentReleasesMetricListingRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnvironmentReleasesMetricListingRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnvironmentReleasesMetricListingRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAttachedFlagCount

`func (o *EnvironmentReleasesMetricListingRep) GetAttachedFlagCount() int32`

GetAttachedFlagCount returns the AttachedFlagCount field if non-nil, zero value otherwise.

### GetAttachedFlagCountOk

`func (o *EnvironmentReleasesMetricListingRep) GetAttachedFlagCountOk() (*int32, bool)`

GetAttachedFlagCountOk returns a tuple with the AttachedFlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedFlagCount

`func (o *EnvironmentReleasesMetricListingRep) SetAttachedFlagCount(v int32)`

SetAttachedFlagCount sets AttachedFlagCount field to given value.

### HasAttachedFlagCount

`func (o *EnvironmentReleasesMetricListingRep) HasAttachedFlagCount() bool`

HasAttachedFlagCount returns a boolean if a field has been set.

### GetLinks

`func (o *EnvironmentReleasesMetricListingRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentReleasesMetricListingRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentReleasesMetricListingRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetSite

`func (o *EnvironmentReleasesMetricListingRep) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *EnvironmentReleasesMetricListingRep) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *EnvironmentReleasesMetricListingRep) SetSite(v Link)`

SetSite sets Site field to given value.

### HasSite

`func (o *EnvironmentReleasesMetricListingRep) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccess

`func (o *EnvironmentReleasesMetricListingRep) GetAccess() EnvironmentReleasesAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *EnvironmentReleasesMetricListingRep) GetAccessOk() (*EnvironmentReleasesAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *EnvironmentReleasesMetricListingRep) SetAccess(v EnvironmentReleasesAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *EnvironmentReleasesMetricListingRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTags

`func (o *EnvironmentReleasesMetricListingRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EnvironmentReleasesMetricListingRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EnvironmentReleasesMetricListingRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreationDate

`func (o *EnvironmentReleasesMetricListingRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *EnvironmentReleasesMetricListingRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *EnvironmentReleasesMetricListingRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *EnvironmentReleasesMetricListingRep) GetLastModified() EnvironmentReleasesModification`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EnvironmentReleasesMetricListingRep) GetLastModifiedOk() (*EnvironmentReleasesModification, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EnvironmentReleasesMetricListingRep) SetLastModified(v EnvironmentReleasesModification)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *EnvironmentReleasesMetricListingRep) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMaintainerId

`func (o *EnvironmentReleasesMetricListingRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *EnvironmentReleasesMetricListingRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *EnvironmentReleasesMetricListingRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *EnvironmentReleasesMetricListingRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainer

`func (o *EnvironmentReleasesMetricListingRep) GetMaintainer() EnvironmentReleasesMemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *EnvironmentReleasesMetricListingRep) GetMaintainerOk() (*EnvironmentReleasesMemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *EnvironmentReleasesMetricListingRep) SetMaintainer(v EnvironmentReleasesMemberSummary)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *EnvironmentReleasesMetricListingRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentReleasesMetricListingRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReleasesMetricListingRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReleasesMetricListingRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReleasesMetricListingRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *EnvironmentReleasesMetricListingRep) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EnvironmentReleasesMetricListingRep) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EnvironmentReleasesMetricListingRep) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EnvironmentReleasesMetricListingRep) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIsNumeric

`func (o *EnvironmentReleasesMetricListingRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *EnvironmentReleasesMetricListingRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *EnvironmentReleasesMetricListingRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *EnvironmentReleasesMetricListingRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *EnvironmentReleasesMetricListingRep) GetSuccessCriteria() string`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *EnvironmentReleasesMetricListingRep) GetSuccessCriteriaOk() (*string, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *EnvironmentReleasesMetricListingRep) SetSuccessCriteria(v string)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *EnvironmentReleasesMetricListingRep) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetUnit

`func (o *EnvironmentReleasesMetricListingRep) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *EnvironmentReleasesMetricListingRep) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *EnvironmentReleasesMetricListingRep) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *EnvironmentReleasesMetricListingRep) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *EnvironmentReleasesMetricListingRep) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *EnvironmentReleasesMetricListingRep) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *EnvironmentReleasesMetricListingRep) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *EnvironmentReleasesMetricListingRep) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetRandomizationUnits

`func (o *EnvironmentReleasesMetricListingRep) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *EnvironmentReleasesMetricListingRep) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *EnvironmentReleasesMetricListingRep) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *EnvironmentReleasesMetricListingRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetAnalysisUnits

`func (o *EnvironmentReleasesMetricListingRep) GetAnalysisUnits() []string`

GetAnalysisUnits returns the AnalysisUnits field if non-nil, zero value otherwise.

### GetAnalysisUnitsOk

`func (o *EnvironmentReleasesMetricListingRep) GetAnalysisUnitsOk() (*[]string, bool)`

GetAnalysisUnitsOk returns a tuple with the AnalysisUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisUnits

`func (o *EnvironmentReleasesMetricListingRep) SetAnalysisUnits(v []string)`

SetAnalysisUnits sets AnalysisUnits field to given value.

### HasAnalysisUnits

`func (o *EnvironmentReleasesMetricListingRep) HasAnalysisUnits() bool`

HasAnalysisUnits returns a boolean if a field has been set.

### GetFilters

`func (o *EnvironmentReleasesMetricListingRep) GetFilters() EnvironmentReleasesFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EnvironmentReleasesMetricListingRep) GetFiltersOk() (*EnvironmentReleasesFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EnvironmentReleasesMetricListingRep) SetFilters(v EnvironmentReleasesFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EnvironmentReleasesMetricListingRep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *EnvironmentReleasesMetricListingRep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *EnvironmentReleasesMetricListingRep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *EnvironmentReleasesMetricListingRep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *EnvironmentReleasesMetricListingRep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetAnalysisType

`func (o *EnvironmentReleasesMetricListingRep) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *EnvironmentReleasesMetricListingRep) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *EnvironmentReleasesMetricListingRep) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.

### HasAnalysisType

`func (o *EnvironmentReleasesMetricListingRep) HasAnalysisType() bool`

HasAnalysisType returns a boolean if a field has been set.

### GetPercentileValue

`func (o *EnvironmentReleasesMetricListingRep) GetPercentileValue() int32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *EnvironmentReleasesMetricListingRep) GetPercentileValueOk() (*int32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *EnvironmentReleasesMetricListingRep) SetPercentileValue(v int32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *EnvironmentReleasesMetricListingRep) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetEventDefault

`func (o *EnvironmentReleasesMetricListingRep) GetEventDefault() MetricEventDefaultRep`

GetEventDefault returns the EventDefault field if non-nil, zero value otherwise.

### GetEventDefaultOk

`func (o *EnvironmentReleasesMetricListingRep) GetEventDefaultOk() (*MetricEventDefaultRep, bool)`

GetEventDefaultOk returns a tuple with the EventDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDefault

`func (o *EnvironmentReleasesMetricListingRep) SetEventDefault(v MetricEventDefaultRep)`

SetEventDefault sets EventDefault field to given value.

### HasEventDefault

`func (o *EnvironmentReleasesMetricListingRep) HasEventDefault() bool`

HasEventDefault returns a boolean if a field has been set.

### GetDataSource

`func (o *EnvironmentReleasesMetricListingRep) GetDataSource() MetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *EnvironmentReleasesMetricListingRep) GetDataSourceOk() (*MetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *EnvironmentReleasesMetricListingRep) SetDataSource(v MetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.


### GetLastSeen

`func (o *EnvironmentReleasesMetricListingRep) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *EnvironmentReleasesMetricListingRep) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *EnvironmentReleasesMetricListingRep) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *EnvironmentReleasesMetricListingRep) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetArchived

`func (o *EnvironmentReleasesMetricListingRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *EnvironmentReleasesMetricListingRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *EnvironmentReleasesMetricListingRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *EnvironmentReleasesMetricListingRep) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedAt

`func (o *EnvironmentReleasesMetricListingRep) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *EnvironmentReleasesMetricListingRep) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *EnvironmentReleasesMetricListingRep) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *EnvironmentReleasesMetricListingRep) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetSelector

`func (o *EnvironmentReleasesMetricListingRep) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *EnvironmentReleasesMetricListingRep) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *EnvironmentReleasesMetricListingRep) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *EnvironmentReleasesMetricListingRep) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *EnvironmentReleasesMetricListingRep) GetUrls() []map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *EnvironmentReleasesMetricListingRep) GetUrlsOk() (*[]map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *EnvironmentReleasesMetricListingRep) SetUrls(v []map[string]interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *EnvironmentReleasesMetricListingRep) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetWindowStartOffset

`func (o *EnvironmentReleasesMetricListingRep) GetWindowStartOffset() int64`

GetWindowStartOffset returns the WindowStartOffset field if non-nil, zero value otherwise.

### GetWindowStartOffsetOk

`func (o *EnvironmentReleasesMetricListingRep) GetWindowStartOffsetOk() (*int64, bool)`

GetWindowStartOffsetOk returns a tuple with the WindowStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartOffset

`func (o *EnvironmentReleasesMetricListingRep) SetWindowStartOffset(v int64)`

SetWindowStartOffset sets WindowStartOffset field to given value.

### HasWindowStartOffset

`func (o *EnvironmentReleasesMetricListingRep) HasWindowStartOffset() bool`

HasWindowStartOffset returns a boolean if a field has been set.

### GetWindowEndOffset

`func (o *EnvironmentReleasesMetricListingRep) GetWindowEndOffset() int64`

GetWindowEndOffset returns the WindowEndOffset field if non-nil, zero value otherwise.

### GetWindowEndOffsetOk

`func (o *EnvironmentReleasesMetricListingRep) GetWindowEndOffsetOk() (*int64, bool)`

GetWindowEndOffsetOk returns a tuple with the WindowEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEndOffset

`func (o *EnvironmentReleasesMetricListingRep) SetWindowEndOffset(v int64)`

SetWindowEndOffset sets WindowEndOffset field to given value.

### HasWindowEndOffset

`func (o *EnvironmentReleasesMetricListingRep) HasWindowEndOffset() bool`

HasWindowEndOffset returns a boolean if a field has been set.

### GetWinsorLowerPercentile

`func (o *EnvironmentReleasesMetricListingRep) GetWinsorLowerPercentile() float32`

GetWinsorLowerPercentile returns the WinsorLowerPercentile field if non-nil, zero value otherwise.

### GetWinsorLowerPercentileOk

`func (o *EnvironmentReleasesMetricListingRep) GetWinsorLowerPercentileOk() (*float32, bool)`

GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorLowerPercentile

`func (o *EnvironmentReleasesMetricListingRep) SetWinsorLowerPercentile(v float32)`

SetWinsorLowerPercentile sets WinsorLowerPercentile field to given value.

### HasWinsorLowerPercentile

`func (o *EnvironmentReleasesMetricListingRep) HasWinsorLowerPercentile() bool`

HasWinsorLowerPercentile returns a boolean if a field has been set.

### GetWinsorUpperPercentile

`func (o *EnvironmentReleasesMetricListingRep) GetWinsorUpperPercentile() float32`

GetWinsorUpperPercentile returns the WinsorUpperPercentile field if non-nil, zero value otherwise.

### GetWinsorUpperPercentileOk

`func (o *EnvironmentReleasesMetricListingRep) GetWinsorUpperPercentileOk() (*float32, bool)`

GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorUpperPercentile

`func (o *EnvironmentReleasesMetricListingRep) SetWinsorUpperPercentile(v float32)`

SetWinsorUpperPercentile sets WinsorUpperPercentile field to given value.

### HasWinsorUpperPercentile

`func (o *EnvironmentReleasesMetricListingRep) HasWinsorUpperPercentile() bool`

HasWinsorUpperPercentile returns a boolean if a field has been set.

### GetWinsorIncludeImputed

`func (o *EnvironmentReleasesMetricListingRep) GetWinsorIncludeImputed() bool`

GetWinsorIncludeImputed returns the WinsorIncludeImputed field if non-nil, zero value otherwise.

### GetWinsorIncludeImputedOk

`func (o *EnvironmentReleasesMetricListingRep) GetWinsorIncludeImputedOk() (*bool, bool)`

GetWinsorIncludeImputedOk returns a tuple with the WinsorIncludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorIncludeImputed

`func (o *EnvironmentReleasesMetricListingRep) SetWinsorIncludeImputed(v bool)`

SetWinsorIncludeImputed sets WinsorIncludeImputed field to given value.

### HasWinsorIncludeImputed

`func (o *EnvironmentReleasesMetricListingRep) HasWinsorIncludeImputed() bool`

HasWinsorIncludeImputed returns a boolean if a field has been set.

### GetTraceQuery

`func (o *EnvironmentReleasesMetricListingRep) GetTraceQuery() string`

GetTraceQuery returns the TraceQuery field if non-nil, zero value otherwise.

### GetTraceQueryOk

`func (o *EnvironmentReleasesMetricListingRep) GetTraceQueryOk() (*string, bool)`

GetTraceQueryOk returns a tuple with the TraceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceQuery

`func (o *EnvironmentReleasesMetricListingRep) SetTraceQuery(v string)`

SetTraceQuery sets TraceQuery field to given value.

### HasTraceQuery

`func (o *EnvironmentReleasesMetricListingRep) HasTraceQuery() bool`

HasTraceQuery returns a boolean if a field has been set.

### GetTraceValueLocation

`func (o *EnvironmentReleasesMetricListingRep) GetTraceValueLocation() string`

GetTraceValueLocation returns the TraceValueLocation field if non-nil, zero value otherwise.

### GetTraceValueLocationOk

`func (o *EnvironmentReleasesMetricListingRep) GetTraceValueLocationOk() (*string, bool)`

GetTraceValueLocationOk returns a tuple with the TraceValueLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceValueLocation

`func (o *EnvironmentReleasesMetricListingRep) SetTraceValueLocation(v string)`

SetTraceValueLocation sets TraceValueLocation field to given value.

### HasTraceValueLocation

`func (o *EnvironmentReleasesMetricListingRep) HasTraceValueLocation() bool`

HasTraceValueLocation returns a boolean if a field has been set.

### GetUnitAggregationField

`func (o *EnvironmentReleasesMetricListingRep) GetUnitAggregationField() string`

GetUnitAggregationField returns the UnitAggregationField field if non-nil, zero value otherwise.

### GetUnitAggregationFieldOk

`func (o *EnvironmentReleasesMetricListingRep) GetUnitAggregationFieldOk() (*string, bool)`

GetUnitAggregationFieldOk returns a tuple with the UnitAggregationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationField

`func (o *EnvironmentReleasesMetricListingRep) SetUnitAggregationField(v string)`

SetUnitAggregationField sets UnitAggregationField field to given value.

### HasUnitAggregationField

`func (o *EnvironmentReleasesMetricListingRep) HasUnitAggregationField() bool`

HasUnitAggregationField returns a boolean if a field has been set.

### GetDenominator

`func (o *EnvironmentReleasesMetricListingRep) GetDenominator() EnvironmentReleasesMetricDenominatorRep`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *EnvironmentReleasesMetricListingRep) GetDenominatorOk() (*EnvironmentReleasesMetricDenominatorRep, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *EnvironmentReleasesMetricListingRep) SetDenominator(v EnvironmentReleasesMetricDenominatorRep)`

SetDenominator sets Denominator field to given value.

### HasDenominator

`func (o *EnvironmentReleasesMetricListingRep) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


