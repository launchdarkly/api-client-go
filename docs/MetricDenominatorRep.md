# MetricDenominatorRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** | The warehouse event column for the denominator | [optional] 
**IsNumeric** | Pointer to **bool** | Whether the denominator aggregates a numeric value | [optional] 
**UnitAggregationType** | Pointer to **string** | How individual unit values are aggregated for the denominator | [optional] 
**UnitAggregationField** | Pointer to **string** | The column to count distinct values of; required when unitAggregationType is count_distinct | [optional] 
**ValueColumn** | Pointer to **string** | For a numeric denominator, the column holding the numeric value | [optional] 
**DataSource** | Pointer to [**MetricDataSourceRefRep**](MetricDataSourceRefRep.md) |  | [optional] 
**Filters** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**WindowStartOffset** | Pointer to **int64** | Start of the measurement window in milliseconds | [optional] 
**WindowEndOffset** | Pointer to **int64** | End of the measurement window in milliseconds | [optional] 
**WinsorLowerPercentile** | Pointer to **float32** | Lower winsorization percentile in the open interval (0, 100) | [optional] 
**WinsorUpperPercentile** | Pointer to **float32** | Upper winsorization percentile in the open interval (0, 100) | [optional] 
**WinsorIncludeImputed** | Pointer to **bool** | When true, the percentile bound calculation includes imputed zeros | [optional] 

## Methods

### NewMetricDenominatorRep

`func NewMetricDenominatorRep() *MetricDenominatorRep`

NewMetricDenominatorRep instantiates a new MetricDenominatorRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDenominatorRepWithDefaults

`func NewMetricDenominatorRepWithDefaults() *MetricDenominatorRep`

NewMetricDenominatorRepWithDefaults instantiates a new MetricDenominatorRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *MetricDenominatorRep) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *MetricDenominatorRep) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *MetricDenominatorRep) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *MetricDenominatorRep) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MetricDenominatorRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricDenominatorRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricDenominatorRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricDenominatorRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *MetricDenominatorRep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *MetricDenominatorRep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *MetricDenominatorRep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *MetricDenominatorRep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetUnitAggregationField

`func (o *MetricDenominatorRep) GetUnitAggregationField() string`

GetUnitAggregationField returns the UnitAggregationField field if non-nil, zero value otherwise.

### GetUnitAggregationFieldOk

`func (o *MetricDenominatorRep) GetUnitAggregationFieldOk() (*string, bool)`

GetUnitAggregationFieldOk returns a tuple with the UnitAggregationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationField

`func (o *MetricDenominatorRep) SetUnitAggregationField(v string)`

SetUnitAggregationField sets UnitAggregationField field to given value.

### HasUnitAggregationField

`func (o *MetricDenominatorRep) HasUnitAggregationField() bool`

HasUnitAggregationField returns a boolean if a field has been set.

### GetValueColumn

`func (o *MetricDenominatorRep) GetValueColumn() string`

GetValueColumn returns the ValueColumn field if non-nil, zero value otherwise.

### GetValueColumnOk

`func (o *MetricDenominatorRep) GetValueColumnOk() (*string, bool)`

GetValueColumnOk returns a tuple with the ValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueColumn

`func (o *MetricDenominatorRep) SetValueColumn(v string)`

SetValueColumn sets ValueColumn field to given value.

### HasValueColumn

`func (o *MetricDenominatorRep) HasValueColumn() bool`

HasValueColumn returns a boolean if a field has been set.

### GetDataSource

`func (o *MetricDenominatorRep) GetDataSource() MetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *MetricDenominatorRep) GetDataSourceOk() (*MetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *MetricDenominatorRep) SetDataSource(v MetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *MetricDenominatorRep) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetFilters

`func (o *MetricDenominatorRep) GetFilters() Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricDenominatorRep) GetFiltersOk() (*Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricDenominatorRep) SetFilters(v Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricDenominatorRep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetWindowStartOffset

`func (o *MetricDenominatorRep) GetWindowStartOffset() int64`

GetWindowStartOffset returns the WindowStartOffset field if non-nil, zero value otherwise.

### GetWindowStartOffsetOk

`func (o *MetricDenominatorRep) GetWindowStartOffsetOk() (*int64, bool)`

GetWindowStartOffsetOk returns a tuple with the WindowStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartOffset

`func (o *MetricDenominatorRep) SetWindowStartOffset(v int64)`

SetWindowStartOffset sets WindowStartOffset field to given value.

### HasWindowStartOffset

`func (o *MetricDenominatorRep) HasWindowStartOffset() bool`

HasWindowStartOffset returns a boolean if a field has been set.

### GetWindowEndOffset

`func (o *MetricDenominatorRep) GetWindowEndOffset() int64`

GetWindowEndOffset returns the WindowEndOffset field if non-nil, zero value otherwise.

### GetWindowEndOffsetOk

`func (o *MetricDenominatorRep) GetWindowEndOffsetOk() (*int64, bool)`

GetWindowEndOffsetOk returns a tuple with the WindowEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEndOffset

`func (o *MetricDenominatorRep) SetWindowEndOffset(v int64)`

SetWindowEndOffset sets WindowEndOffset field to given value.

### HasWindowEndOffset

`func (o *MetricDenominatorRep) HasWindowEndOffset() bool`

HasWindowEndOffset returns a boolean if a field has been set.

### GetWinsorLowerPercentile

`func (o *MetricDenominatorRep) GetWinsorLowerPercentile() float32`

GetWinsorLowerPercentile returns the WinsorLowerPercentile field if non-nil, zero value otherwise.

### GetWinsorLowerPercentileOk

`func (o *MetricDenominatorRep) GetWinsorLowerPercentileOk() (*float32, bool)`

GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorLowerPercentile

`func (o *MetricDenominatorRep) SetWinsorLowerPercentile(v float32)`

SetWinsorLowerPercentile sets WinsorLowerPercentile field to given value.

### HasWinsorLowerPercentile

`func (o *MetricDenominatorRep) HasWinsorLowerPercentile() bool`

HasWinsorLowerPercentile returns a boolean if a field has been set.

### GetWinsorUpperPercentile

`func (o *MetricDenominatorRep) GetWinsorUpperPercentile() float32`

GetWinsorUpperPercentile returns the WinsorUpperPercentile field if non-nil, zero value otherwise.

### GetWinsorUpperPercentileOk

`func (o *MetricDenominatorRep) GetWinsorUpperPercentileOk() (*float32, bool)`

GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorUpperPercentile

`func (o *MetricDenominatorRep) SetWinsorUpperPercentile(v float32)`

SetWinsorUpperPercentile sets WinsorUpperPercentile field to given value.

### HasWinsorUpperPercentile

`func (o *MetricDenominatorRep) HasWinsorUpperPercentile() bool`

HasWinsorUpperPercentile returns a boolean if a field has been set.

### GetWinsorIncludeImputed

`func (o *MetricDenominatorRep) GetWinsorIncludeImputed() bool`

GetWinsorIncludeImputed returns the WinsorIncludeImputed field if non-nil, zero value otherwise.

### GetWinsorIncludeImputedOk

`func (o *MetricDenominatorRep) GetWinsorIncludeImputedOk() (*bool, bool)`

GetWinsorIncludeImputedOk returns a tuple with the WinsorIncludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorIncludeImputed

`func (o *MetricDenominatorRep) SetWinsorIncludeImputed(v bool)`

SetWinsorIncludeImputed sets WinsorIncludeImputed field to given value.

### HasWinsorIncludeImputed

`func (o *MetricDenominatorRep) HasWinsorIncludeImputed() bool`

HasWinsorIncludeImputed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


