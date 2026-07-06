# AiConfigsMetricDenominatorRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** | The warehouse event column for the denominator | [optional] 
**IsNumeric** | Pointer to **bool** | Whether the denominator aggregates a numeric value | [optional] 
**UnitAggregationType** | Pointer to **string** | How individual unit values are aggregated for the denominator | [optional] 
**UnitAggregationField** | Pointer to **string** | The column to count distinct values of; required when unitAggregationType is count_distinct | [optional] 
**DataSource** | Pointer to [**AiConfigsMetricDataSourceRefRep**](AiConfigsMetricDataSourceRefRep.md) |  | [optional] 
**Filters** | Pointer to [**AiConfigsFilter**](AiConfigsFilter.md) |  | [optional] 
**WindowStartOffset** | Pointer to **int64** | Start of the measurement window in milliseconds | [optional] 
**WindowEndOffset** | Pointer to **int64** | End of the measurement window in milliseconds | [optional] 
**WinsorLowerPercentile** | Pointer to **float32** | Lower winsorization percentile in the open interval (0, 100) | [optional] 
**WinsorUpperPercentile** | Pointer to **float32** | Upper winsorization percentile in the open interval (0, 100) | [optional] 
**WinsorExcludeImputed** | Pointer to **bool** | Deprecated and ignored. Use winsorIncludeImputed instead. | [optional] 
**WinsorIncludeImputed** | Pointer to **bool** | When true, the percentile bound calculation includes imputed zeros | [optional] 

## Methods

### NewAiConfigsMetricDenominatorRep

`func NewAiConfigsMetricDenominatorRep() *AiConfigsMetricDenominatorRep`

NewAiConfigsMetricDenominatorRep instantiates a new AiConfigsMetricDenominatorRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsMetricDenominatorRepWithDefaults

`func NewAiConfigsMetricDenominatorRepWithDefaults() *AiConfigsMetricDenominatorRep`

NewAiConfigsMetricDenominatorRepWithDefaults instantiates a new AiConfigsMetricDenominatorRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *AiConfigsMetricDenominatorRep) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *AiConfigsMetricDenominatorRep) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *AiConfigsMetricDenominatorRep) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *AiConfigsMetricDenominatorRep) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetIsNumeric

`func (o *AiConfigsMetricDenominatorRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *AiConfigsMetricDenominatorRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *AiConfigsMetricDenominatorRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *AiConfigsMetricDenominatorRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *AiConfigsMetricDenominatorRep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *AiConfigsMetricDenominatorRep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *AiConfigsMetricDenominatorRep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *AiConfigsMetricDenominatorRep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetUnitAggregationField

`func (o *AiConfigsMetricDenominatorRep) GetUnitAggregationField() string`

GetUnitAggregationField returns the UnitAggregationField field if non-nil, zero value otherwise.

### GetUnitAggregationFieldOk

`func (o *AiConfigsMetricDenominatorRep) GetUnitAggregationFieldOk() (*string, bool)`

GetUnitAggregationFieldOk returns a tuple with the UnitAggregationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationField

`func (o *AiConfigsMetricDenominatorRep) SetUnitAggregationField(v string)`

SetUnitAggregationField sets UnitAggregationField field to given value.

### HasUnitAggregationField

`func (o *AiConfigsMetricDenominatorRep) HasUnitAggregationField() bool`

HasUnitAggregationField returns a boolean if a field has been set.

### GetDataSource

`func (o *AiConfigsMetricDenominatorRep) GetDataSource() AiConfigsMetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *AiConfigsMetricDenominatorRep) GetDataSourceOk() (*AiConfigsMetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *AiConfigsMetricDenominatorRep) SetDataSource(v AiConfigsMetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *AiConfigsMetricDenominatorRep) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetFilters

`func (o *AiConfigsMetricDenominatorRep) GetFilters() AiConfigsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AiConfigsMetricDenominatorRep) GetFiltersOk() (*AiConfigsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AiConfigsMetricDenominatorRep) SetFilters(v AiConfigsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AiConfigsMetricDenominatorRep) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetWindowStartOffset

`func (o *AiConfigsMetricDenominatorRep) GetWindowStartOffset() int64`

GetWindowStartOffset returns the WindowStartOffset field if non-nil, zero value otherwise.

### GetWindowStartOffsetOk

`func (o *AiConfigsMetricDenominatorRep) GetWindowStartOffsetOk() (*int64, bool)`

GetWindowStartOffsetOk returns a tuple with the WindowStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartOffset

`func (o *AiConfigsMetricDenominatorRep) SetWindowStartOffset(v int64)`

SetWindowStartOffset sets WindowStartOffset field to given value.

### HasWindowStartOffset

`func (o *AiConfigsMetricDenominatorRep) HasWindowStartOffset() bool`

HasWindowStartOffset returns a boolean if a field has been set.

### GetWindowEndOffset

`func (o *AiConfigsMetricDenominatorRep) GetWindowEndOffset() int64`

GetWindowEndOffset returns the WindowEndOffset field if non-nil, zero value otherwise.

### GetWindowEndOffsetOk

`func (o *AiConfigsMetricDenominatorRep) GetWindowEndOffsetOk() (*int64, bool)`

GetWindowEndOffsetOk returns a tuple with the WindowEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEndOffset

`func (o *AiConfigsMetricDenominatorRep) SetWindowEndOffset(v int64)`

SetWindowEndOffset sets WindowEndOffset field to given value.

### HasWindowEndOffset

`func (o *AiConfigsMetricDenominatorRep) HasWindowEndOffset() bool`

HasWindowEndOffset returns a boolean if a field has been set.

### GetWinsorLowerPercentile

`func (o *AiConfigsMetricDenominatorRep) GetWinsorLowerPercentile() float32`

GetWinsorLowerPercentile returns the WinsorLowerPercentile field if non-nil, zero value otherwise.

### GetWinsorLowerPercentileOk

`func (o *AiConfigsMetricDenominatorRep) GetWinsorLowerPercentileOk() (*float32, bool)`

GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorLowerPercentile

`func (o *AiConfigsMetricDenominatorRep) SetWinsorLowerPercentile(v float32)`

SetWinsorLowerPercentile sets WinsorLowerPercentile field to given value.

### HasWinsorLowerPercentile

`func (o *AiConfigsMetricDenominatorRep) HasWinsorLowerPercentile() bool`

HasWinsorLowerPercentile returns a boolean if a field has been set.

### GetWinsorUpperPercentile

`func (o *AiConfigsMetricDenominatorRep) GetWinsorUpperPercentile() float32`

GetWinsorUpperPercentile returns the WinsorUpperPercentile field if non-nil, zero value otherwise.

### GetWinsorUpperPercentileOk

`func (o *AiConfigsMetricDenominatorRep) GetWinsorUpperPercentileOk() (*float32, bool)`

GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorUpperPercentile

`func (o *AiConfigsMetricDenominatorRep) SetWinsorUpperPercentile(v float32)`

SetWinsorUpperPercentile sets WinsorUpperPercentile field to given value.

### HasWinsorUpperPercentile

`func (o *AiConfigsMetricDenominatorRep) HasWinsorUpperPercentile() bool`

HasWinsorUpperPercentile returns a boolean if a field has been set.

### GetWinsorExcludeImputed

`func (o *AiConfigsMetricDenominatorRep) GetWinsorExcludeImputed() bool`

GetWinsorExcludeImputed returns the WinsorExcludeImputed field if non-nil, zero value otherwise.

### GetWinsorExcludeImputedOk

`func (o *AiConfigsMetricDenominatorRep) GetWinsorExcludeImputedOk() (*bool, bool)`

GetWinsorExcludeImputedOk returns a tuple with the WinsorExcludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorExcludeImputed

`func (o *AiConfigsMetricDenominatorRep) SetWinsorExcludeImputed(v bool)`

SetWinsorExcludeImputed sets WinsorExcludeImputed field to given value.

### HasWinsorExcludeImputed

`func (o *AiConfigsMetricDenominatorRep) HasWinsorExcludeImputed() bool`

HasWinsorExcludeImputed returns a boolean if a field has been set.

### GetWinsorIncludeImputed

`func (o *AiConfigsMetricDenominatorRep) GetWinsorIncludeImputed() bool`

GetWinsorIncludeImputed returns the WinsorIncludeImputed field if non-nil, zero value otherwise.

### GetWinsorIncludeImputedOk

`func (o *AiConfigsMetricDenominatorRep) GetWinsorIncludeImputedOk() (*bool, bool)`

GetWinsorIncludeImputedOk returns a tuple with the WinsorIncludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorIncludeImputed

`func (o *AiConfigsMetricDenominatorRep) SetWinsorIncludeImputed(v bool)`

SetWinsorIncludeImputed sets WinsorIncludeImputed field to given value.

### HasWinsorIncludeImputed

`func (o *AiConfigsMetricDenominatorRep) HasWinsorIncludeImputed() bool`

HasWinsorIncludeImputed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


