# MetricRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentCount** | Pointer to **int32** | The number of experiments using this metric | [optional] 
**Id** | **string** | The ID of this metric | 
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
**IsNumeric** | Pointer to **bool** | For custom metrics, whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). | [optional] 
**SuccessCriteria** | Pointer to **string** | For custom metrics, the success criteria | [optional] 
**Unit** | Pointer to **string** | For numeric custom metrics, the unit of measure | [optional] 
**EventKey** | Pointer to **string** | For custom metrics, the event key to use in your code | [optional] 
**RandomizationUnits** | Pointer to **[]string** | An array of randomization units allowed for this metric | [optional] 
**UnitAggregationType** | Pointer to **string** | The method in which multiple unit event values are aggregated | [optional] 
**AnalysisType** | Pointer to **string** | The strategy for analyzing metric events | [optional] 
**PercentileValue** | Pointer to **int32** | The percentile, an integer denoting the target percentile between 0 and 100. Only present when &lt;code&gt;analysisType&lt;/code&gt; is &lt;code&gt;percentile&lt;/code&gt;. | [optional] 
**EventDefault** | Pointer to [**MetricEventDefaultRep**](MetricEventDefaultRep.md) |  | [optional] 
**Experiments** | Pointer to [**[]DependentExperimentRep**](DependentExperimentRep.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Whether the metric is active | [optional] 
**AttachedFeatures** | Pointer to [**[]FlagListingRep**](FlagListingRep.md) | Details on the flags attached to this metric | [optional] 
**Version** | Pointer to **int32** | Version of the metric | [optional] 
**Selector** | Pointer to **string** | For click metrics, the CSS selectors | [optional] 
**Urls** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewMetricRep

`func NewMetricRep(id string, key string, name string, kind string, links map[string]Link, tags []string, creationDate int64, ) *MetricRep`

NewMetricRep instantiates a new MetricRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRepWithDefaults

`func NewMetricRepWithDefaults() *MetricRep`

NewMetricRepWithDefaults instantiates a new MetricRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentCount

`func (o *MetricRep) GetExperimentCount() int32`

GetExperimentCount returns the ExperimentCount field if non-nil, zero value otherwise.

### GetExperimentCountOk

`func (o *MetricRep) GetExperimentCountOk() (*int32, bool)`

GetExperimentCountOk returns a tuple with the ExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentCount

`func (o *MetricRep) SetExperimentCount(v int32)`

SetExperimentCount sets ExperimentCount field to given value.

### HasExperimentCount

`func (o *MetricRep) HasExperimentCount() bool`

HasExperimentCount returns a boolean if a field has been set.

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

`func (o *MetricRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetSite

`func (o *MetricRep) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MetricRep) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MetricRep) SetSite(v Link)`

SetSite sets Site field to given value.

### HasSite

`func (o *MetricRep) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAccess

`func (o *MetricRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MetricRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MetricRep) SetAccess(v Access)`

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


### GetLastModified

`func (o *MetricRep) GetLastModified() Modification`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *MetricRep) GetLastModifiedOk() (*Modification, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *MetricRep) SetLastModified(v Modification)`

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

`func (o *MetricRep) GetMaintainer() MemberSummary`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *MetricRep) GetMaintainerOk() (*MemberSummary, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *MetricRep) SetMaintainer(v MemberSummary)`

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

`func (o *MetricRep) GetSuccessCriteria() string`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricRep) GetSuccessCriteriaOk() (*string, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricRep) SetSuccessCriteria(v string)`

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

### GetRandomizationUnits

`func (o *MetricRep) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *MetricRep) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *MetricRep) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *MetricRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *MetricRep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *MetricRep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *MetricRep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *MetricRep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetAnalysisType

`func (o *MetricRep) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *MetricRep) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *MetricRep) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.

### HasAnalysisType

`func (o *MetricRep) HasAnalysisType() bool`

HasAnalysisType returns a boolean if a field has been set.

### GetPercentileValue

`func (o *MetricRep) GetPercentileValue() int32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *MetricRep) GetPercentileValueOk() (*int32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *MetricRep) SetPercentileValue(v int32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *MetricRep) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetEventDefault

`func (o *MetricRep) GetEventDefault() MetricEventDefaultRep`

GetEventDefault returns the EventDefault field if non-nil, zero value otherwise.

### GetEventDefaultOk

`func (o *MetricRep) GetEventDefaultOk() (*MetricEventDefaultRep, bool)`

GetEventDefaultOk returns a tuple with the EventDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDefault

`func (o *MetricRep) SetEventDefault(v MetricEventDefaultRep)`

SetEventDefault sets EventDefault field to given value.

### HasEventDefault

`func (o *MetricRep) HasEventDefault() bool`

HasEventDefault returns a boolean if a field has been set.

### GetExperiments

`func (o *MetricRep) GetExperiments() []DependentExperimentRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *MetricRep) GetExperimentsOk() (*[]DependentExperimentRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *MetricRep) SetExperiments(v []DependentExperimentRep)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *MetricRep) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

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

`func (o *MetricRep) GetUrls() []map[string]interface{}`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MetricRep) GetUrlsOk() (*[]map[string]interface{}, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MetricRep) SetUrls(v []map[string]interface{})`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MetricRep) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


