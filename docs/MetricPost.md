# MetricPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key to reference the metric | 
**Name** | Pointer to **string** | A human-friendly name for the metric | [optional] 
**Description** | Pointer to **string** | Description of the metric | [optional] 
**Kind** | **string** | The kind of event your metric will track | 
**Selector** | Pointer to **string** | One or more CSS selectors. Required for click metrics only. | [optional] 
**Urls** | Pointer to [**[]UrlPost**](UrlPost.md) | One or more target URLs. Required for click and pageview metrics only. | [optional] 
**IsNumeric** | Pointer to **bool** | Whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). Required for custom and trace metrics only. | [optional] 
**Unit** | Pointer to **string** | The unit of measure. Applicable for numeric custom and trace metrics only. | [optional] 
**EventKey** | Pointer to **string** | The event key to use in your code. Required for custom conversion/binary and custom numeric metrics only. | [optional] 
**SuccessCriteria** | Pointer to **string** | Success criteria. Required for custom and trace numeric metrics, optional for custom and trace conversion metrics. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the metric | [optional] 
**RandomizationUnits** | Pointer to **[]string** | Deprecated, use &lt;code&gt;analysisUnits&lt;/code&gt; instead. | [optional] 
**AnalysisUnits** | Pointer to **[]string** | An array of analysis units allowed for this metric. | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this metric | [optional] 
**UnitAggregationType** | Pointer to **string** | The method by which multiple unit event values are aggregated | [optional] 
**AnalysisType** | Pointer to **string** | The method for analyzing metric events | [optional] 
**PercentileValue** | Pointer to **int32** | The percentile for the analysis method. An integer denoting the target percentile between 0 and 100. Required when &lt;code&gt;analysisType&lt;/code&gt; is &lt;code&gt;percentile&lt;/code&gt;. | [optional] 
**EventDefault** | Pointer to [**MetricEventDefaultRep**](MetricEventDefaultRep.md) |  | [optional] 
**DataSource** | Pointer to [**MetricDataSourceRefRep**](MetricDataSourceRefRep.md) |  | [optional] 
**Filters** | Pointer to [**EventFilter**](EventFilter.md) |  | [optional] 
**WindowStartOffset** | Pointer to **int64** | Not yet implemented - The start of the measurement window, in milliseconds relative to the unit&#39;s first exposure to a flag variation | [optional] 
**WindowEndOffset** | Pointer to **int64** | Not yet implemented - The end of the measurement window, in milliseconds relative to the unit&#39;s first exposure to a flag variation | [optional] 
**WinsorLowerPercentile** | Pointer to **float32** | Lower winsorization percentile, expressed as a percent in the open interval (0, 100). When both bounds are set, defines a two-sided clamp range. Otherwise lower-only winsorization. | [optional] 
**WinsorUpperPercentile** | Pointer to **float32** | Upper winsorization percentile, expressed as a percent in the open interval (0, 100). When both bounds are set, must be greater than winsorLowerPercentile. | [optional] 
**WinsorExcludeImputed** | Pointer to **bool** | Deprecated and ignored. Use winsorIncludeImputed instead. | [optional] 
**WinsorIncludeImputed** | Pointer to **bool** | When true, the percentile bound calculation includes imputed zeros. Only meaningful when at least one bound is set and the metric includes units that didn&#39;t send events. | [optional] 
**TraceQuery** | Pointer to **string** | The trace query to use for the metric. Required for trace metrics. | [optional] 
**TraceValueLocation** | Pointer to **string** | The location in the trace to use for numeric values. Required for numeric trace metrics. | [optional] 
**UnitAggregationField** | Pointer to **string** | The warehouse column to use for counting distinct values. Required when the unitAggregationType is count_distinct. | [optional] 
**Denominator** | Pointer to [**DenominatorPost**](DenominatorPost.md) |  | [optional] 

## Methods

### NewMetricPost

`func NewMetricPost(key string, kind string, ) *MetricPost`

NewMetricPost instantiates a new MetricPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricPostWithDefaults

`func NewMetricPostWithDefaults() *MetricPost`

NewMetricPostWithDefaults instantiates a new MetricPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MetricPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetricPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *MetricPost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricPost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricPost) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSelector

`func (o *MetricPost) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MetricPost) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MetricPost) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MetricPost) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetUrls

`func (o *MetricPost) GetUrls() []UrlPost`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MetricPost) GetUrlsOk() (*[]UrlPost, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MetricPost) SetUrls(v []UrlPost)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MetricPost) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MetricPost) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricPost) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricPost) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricPost) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetUnit

`func (o *MetricPost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricPost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricPost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricPost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetEventKey

`func (o *MetricPost) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *MetricPost) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *MetricPost) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *MetricPost) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetSuccessCriteria

`func (o *MetricPost) GetSuccessCriteria() string`

GetSuccessCriteria returns the SuccessCriteria field if non-nil, zero value otherwise.

### GetSuccessCriteriaOk

`func (o *MetricPost) GetSuccessCriteriaOk() (*string, bool)`

GetSuccessCriteriaOk returns a tuple with the SuccessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCriteria

`func (o *MetricPost) SetSuccessCriteria(v string)`

SetSuccessCriteria sets SuccessCriteria field to given value.

### HasSuccessCriteria

`func (o *MetricPost) HasSuccessCriteria() bool`

HasSuccessCriteria returns a boolean if a field has been set.

### GetTags

`func (o *MetricPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRandomizationUnits

`func (o *MetricPost) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *MetricPost) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *MetricPost) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *MetricPost) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetAnalysisUnits

`func (o *MetricPost) GetAnalysisUnits() []string`

GetAnalysisUnits returns the AnalysisUnits field if non-nil, zero value otherwise.

### GetAnalysisUnitsOk

`func (o *MetricPost) GetAnalysisUnitsOk() (*[]string, bool)`

GetAnalysisUnitsOk returns a tuple with the AnalysisUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisUnits

`func (o *MetricPost) SetAnalysisUnits(v []string)`

SetAnalysisUnits sets AnalysisUnits field to given value.

### HasAnalysisUnits

`func (o *MetricPost) HasAnalysisUnits() bool`

HasAnalysisUnits returns a boolean if a field has been set.

### GetMaintainerId

`func (o *MetricPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *MetricPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *MetricPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *MetricPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *MetricPost) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *MetricPost) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *MetricPost) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *MetricPost) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetAnalysisType

`func (o *MetricPost) GetAnalysisType() string`

GetAnalysisType returns the AnalysisType field if non-nil, zero value otherwise.

### GetAnalysisTypeOk

`func (o *MetricPost) GetAnalysisTypeOk() (*string, bool)`

GetAnalysisTypeOk returns a tuple with the AnalysisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisType

`func (o *MetricPost) SetAnalysisType(v string)`

SetAnalysisType sets AnalysisType field to given value.

### HasAnalysisType

`func (o *MetricPost) HasAnalysisType() bool`

HasAnalysisType returns a boolean if a field has been set.

### GetPercentileValue

`func (o *MetricPost) GetPercentileValue() int32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *MetricPost) GetPercentileValueOk() (*int32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *MetricPost) SetPercentileValue(v int32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *MetricPost) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetEventDefault

`func (o *MetricPost) GetEventDefault() MetricEventDefaultRep`

GetEventDefault returns the EventDefault field if non-nil, zero value otherwise.

### GetEventDefaultOk

`func (o *MetricPost) GetEventDefaultOk() (*MetricEventDefaultRep, bool)`

GetEventDefaultOk returns a tuple with the EventDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDefault

`func (o *MetricPost) SetEventDefault(v MetricEventDefaultRep)`

SetEventDefault sets EventDefault field to given value.

### HasEventDefault

`func (o *MetricPost) HasEventDefault() bool`

HasEventDefault returns a boolean if a field has been set.

### GetDataSource

`func (o *MetricPost) GetDataSource() MetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *MetricPost) GetDataSourceOk() (*MetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *MetricPost) SetDataSource(v MetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *MetricPost) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetFilters

`func (o *MetricPost) GetFilters() EventFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricPost) GetFiltersOk() (*EventFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricPost) SetFilters(v EventFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricPost) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetWindowStartOffset

`func (o *MetricPost) GetWindowStartOffset() int64`

GetWindowStartOffset returns the WindowStartOffset field if non-nil, zero value otherwise.

### GetWindowStartOffsetOk

`func (o *MetricPost) GetWindowStartOffsetOk() (*int64, bool)`

GetWindowStartOffsetOk returns a tuple with the WindowStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartOffset

`func (o *MetricPost) SetWindowStartOffset(v int64)`

SetWindowStartOffset sets WindowStartOffset field to given value.

### HasWindowStartOffset

`func (o *MetricPost) HasWindowStartOffset() bool`

HasWindowStartOffset returns a boolean if a field has been set.

### GetWindowEndOffset

`func (o *MetricPost) GetWindowEndOffset() int64`

GetWindowEndOffset returns the WindowEndOffset field if non-nil, zero value otherwise.

### GetWindowEndOffsetOk

`func (o *MetricPost) GetWindowEndOffsetOk() (*int64, bool)`

GetWindowEndOffsetOk returns a tuple with the WindowEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEndOffset

`func (o *MetricPost) SetWindowEndOffset(v int64)`

SetWindowEndOffset sets WindowEndOffset field to given value.

### HasWindowEndOffset

`func (o *MetricPost) HasWindowEndOffset() bool`

HasWindowEndOffset returns a boolean if a field has been set.

### GetWinsorLowerPercentile

`func (o *MetricPost) GetWinsorLowerPercentile() float32`

GetWinsorLowerPercentile returns the WinsorLowerPercentile field if non-nil, zero value otherwise.

### GetWinsorLowerPercentileOk

`func (o *MetricPost) GetWinsorLowerPercentileOk() (*float32, bool)`

GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorLowerPercentile

`func (o *MetricPost) SetWinsorLowerPercentile(v float32)`

SetWinsorLowerPercentile sets WinsorLowerPercentile field to given value.

### HasWinsorLowerPercentile

`func (o *MetricPost) HasWinsorLowerPercentile() bool`

HasWinsorLowerPercentile returns a boolean if a field has been set.

### GetWinsorUpperPercentile

`func (o *MetricPost) GetWinsorUpperPercentile() float32`

GetWinsorUpperPercentile returns the WinsorUpperPercentile field if non-nil, zero value otherwise.

### GetWinsorUpperPercentileOk

`func (o *MetricPost) GetWinsorUpperPercentileOk() (*float32, bool)`

GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorUpperPercentile

`func (o *MetricPost) SetWinsorUpperPercentile(v float32)`

SetWinsorUpperPercentile sets WinsorUpperPercentile field to given value.

### HasWinsorUpperPercentile

`func (o *MetricPost) HasWinsorUpperPercentile() bool`

HasWinsorUpperPercentile returns a boolean if a field has been set.

### GetWinsorExcludeImputed

`func (o *MetricPost) GetWinsorExcludeImputed() bool`

GetWinsorExcludeImputed returns the WinsorExcludeImputed field if non-nil, zero value otherwise.

### GetWinsorExcludeImputedOk

`func (o *MetricPost) GetWinsorExcludeImputedOk() (*bool, bool)`

GetWinsorExcludeImputedOk returns a tuple with the WinsorExcludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorExcludeImputed

`func (o *MetricPost) SetWinsorExcludeImputed(v bool)`

SetWinsorExcludeImputed sets WinsorExcludeImputed field to given value.

### HasWinsorExcludeImputed

`func (o *MetricPost) HasWinsorExcludeImputed() bool`

HasWinsorExcludeImputed returns a boolean if a field has been set.

### GetWinsorIncludeImputed

`func (o *MetricPost) GetWinsorIncludeImputed() bool`

GetWinsorIncludeImputed returns the WinsorIncludeImputed field if non-nil, zero value otherwise.

### GetWinsorIncludeImputedOk

`func (o *MetricPost) GetWinsorIncludeImputedOk() (*bool, bool)`

GetWinsorIncludeImputedOk returns a tuple with the WinsorIncludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorIncludeImputed

`func (o *MetricPost) SetWinsorIncludeImputed(v bool)`

SetWinsorIncludeImputed sets WinsorIncludeImputed field to given value.

### HasWinsorIncludeImputed

`func (o *MetricPost) HasWinsorIncludeImputed() bool`

HasWinsorIncludeImputed returns a boolean if a field has been set.

### GetTraceQuery

`func (o *MetricPost) GetTraceQuery() string`

GetTraceQuery returns the TraceQuery field if non-nil, zero value otherwise.

### GetTraceQueryOk

`func (o *MetricPost) GetTraceQueryOk() (*string, bool)`

GetTraceQueryOk returns a tuple with the TraceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceQuery

`func (o *MetricPost) SetTraceQuery(v string)`

SetTraceQuery sets TraceQuery field to given value.

### HasTraceQuery

`func (o *MetricPost) HasTraceQuery() bool`

HasTraceQuery returns a boolean if a field has been set.

### GetTraceValueLocation

`func (o *MetricPost) GetTraceValueLocation() string`

GetTraceValueLocation returns the TraceValueLocation field if non-nil, zero value otherwise.

### GetTraceValueLocationOk

`func (o *MetricPost) GetTraceValueLocationOk() (*string, bool)`

GetTraceValueLocationOk returns a tuple with the TraceValueLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceValueLocation

`func (o *MetricPost) SetTraceValueLocation(v string)`

SetTraceValueLocation sets TraceValueLocation field to given value.

### HasTraceValueLocation

`func (o *MetricPost) HasTraceValueLocation() bool`

HasTraceValueLocation returns a boolean if a field has been set.

### GetUnitAggregationField

`func (o *MetricPost) GetUnitAggregationField() string`

GetUnitAggregationField returns the UnitAggregationField field if non-nil, zero value otherwise.

### GetUnitAggregationFieldOk

`func (o *MetricPost) GetUnitAggregationFieldOk() (*string, bool)`

GetUnitAggregationFieldOk returns a tuple with the UnitAggregationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationField

`func (o *MetricPost) SetUnitAggregationField(v string)`

SetUnitAggregationField sets UnitAggregationField field to given value.

### HasUnitAggregationField

`func (o *MetricPost) HasUnitAggregationField() bool`

HasUnitAggregationField returns a boolean if a field has been set.

### GetDenominator

`func (o *MetricPost) GetDenominator() DenominatorPost`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *MetricPost) GetDenominatorOk() (*DenominatorPost, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *MetricPost) SetDenominator(v DenominatorPost)`

SetDenominator sets Denominator field to given value.

### HasDenominator

`func (o *MetricPost) HasDenominator() bool`

HasDenominator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


