# DenominatorPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** | The warehouse event column for the denominator. Required. | [optional] 
**IsNumeric** | Pointer to **bool** | Whether the denominator aggregates a numeric value | [optional] 
**DataSource** | Pointer to [**MetricDataSourceRefRep**](MetricDataSourceRefRep.md) |  | [optional] 
**UnitAggregationType** | Pointer to **string** | How individual unit values are aggregated. One of: average, sum, count_distinct | [optional] 
**UnitAggregationField** | Pointer to **string** | The warehouse column to use for counting distinct values. Required when the unitAggregationType is count_distinct. | [optional] 
**ValueColumn** | Pointer to **string** | For a numeric denominator, the column holding the numeric value. Overrides the value column mapped on the denominator data source. | [optional] 
**Filters** | Pointer to [**EventFilter**](EventFilter.md) |  | [optional] 
**WindowStartOffset** | Pointer to **int64** | Start of the measurement window in milliseconds | [optional] 
**WindowEndOffset** | Pointer to **int64** | End of the measurement window in milliseconds | [optional] 
**WinsorLowerPercentile** | Pointer to **float32** | Lower winsorization percentile in the open interval (0, 100) | [optional] 
**WinsorUpperPercentile** | Pointer to **float32** | Upper winsorization percentile in the open interval (0, 100) | [optional] 
**WinsorIncludeImputed** | Pointer to **bool** | When true, includes imputed zeros in the percentile bound calculation | [optional] 

## Methods

### NewDenominatorPost

`func NewDenominatorPost() *DenominatorPost`

NewDenominatorPost instantiates a new DenominatorPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDenominatorPostWithDefaults

`func NewDenominatorPostWithDefaults() *DenominatorPost`

NewDenominatorPostWithDefaults instantiates a new DenominatorPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *DenominatorPost) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DenominatorPost) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DenominatorPost) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DenominatorPost) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetIsNumeric

`func (o *DenominatorPost) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *DenominatorPost) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *DenominatorPost) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *DenominatorPost) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetDataSource

`func (o *DenominatorPost) GetDataSource() MetricDataSourceRefRep`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DenominatorPost) GetDataSourceOk() (*MetricDataSourceRefRep, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DenominatorPost) SetDataSource(v MetricDataSourceRefRep)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DenominatorPost) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *DenominatorPost) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *DenominatorPost) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *DenominatorPost) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *DenominatorPost) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetUnitAggregationField

`func (o *DenominatorPost) GetUnitAggregationField() string`

GetUnitAggregationField returns the UnitAggregationField field if non-nil, zero value otherwise.

### GetUnitAggregationFieldOk

`func (o *DenominatorPost) GetUnitAggregationFieldOk() (*string, bool)`

GetUnitAggregationFieldOk returns a tuple with the UnitAggregationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationField

`func (o *DenominatorPost) SetUnitAggregationField(v string)`

SetUnitAggregationField sets UnitAggregationField field to given value.

### HasUnitAggregationField

`func (o *DenominatorPost) HasUnitAggregationField() bool`

HasUnitAggregationField returns a boolean if a field has been set.

### GetValueColumn

`func (o *DenominatorPost) GetValueColumn() string`

GetValueColumn returns the ValueColumn field if non-nil, zero value otherwise.

### GetValueColumnOk

`func (o *DenominatorPost) GetValueColumnOk() (*string, bool)`

GetValueColumnOk returns a tuple with the ValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueColumn

`func (o *DenominatorPost) SetValueColumn(v string)`

SetValueColumn sets ValueColumn field to given value.

### HasValueColumn

`func (o *DenominatorPost) HasValueColumn() bool`

HasValueColumn returns a boolean if a field has been set.

### GetFilters

`func (o *DenominatorPost) GetFilters() EventFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DenominatorPost) GetFiltersOk() (*EventFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DenominatorPost) SetFilters(v EventFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DenominatorPost) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetWindowStartOffset

`func (o *DenominatorPost) GetWindowStartOffset() int64`

GetWindowStartOffset returns the WindowStartOffset field if non-nil, zero value otherwise.

### GetWindowStartOffsetOk

`func (o *DenominatorPost) GetWindowStartOffsetOk() (*int64, bool)`

GetWindowStartOffsetOk returns a tuple with the WindowStartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartOffset

`func (o *DenominatorPost) SetWindowStartOffset(v int64)`

SetWindowStartOffset sets WindowStartOffset field to given value.

### HasWindowStartOffset

`func (o *DenominatorPost) HasWindowStartOffset() bool`

HasWindowStartOffset returns a boolean if a field has been set.

### GetWindowEndOffset

`func (o *DenominatorPost) GetWindowEndOffset() int64`

GetWindowEndOffset returns the WindowEndOffset field if non-nil, zero value otherwise.

### GetWindowEndOffsetOk

`func (o *DenominatorPost) GetWindowEndOffsetOk() (*int64, bool)`

GetWindowEndOffsetOk returns a tuple with the WindowEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEndOffset

`func (o *DenominatorPost) SetWindowEndOffset(v int64)`

SetWindowEndOffset sets WindowEndOffset field to given value.

### HasWindowEndOffset

`func (o *DenominatorPost) HasWindowEndOffset() bool`

HasWindowEndOffset returns a boolean if a field has been set.

### GetWinsorLowerPercentile

`func (o *DenominatorPost) GetWinsorLowerPercentile() float32`

GetWinsorLowerPercentile returns the WinsorLowerPercentile field if non-nil, zero value otherwise.

### GetWinsorLowerPercentileOk

`func (o *DenominatorPost) GetWinsorLowerPercentileOk() (*float32, bool)`

GetWinsorLowerPercentileOk returns a tuple with the WinsorLowerPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorLowerPercentile

`func (o *DenominatorPost) SetWinsorLowerPercentile(v float32)`

SetWinsorLowerPercentile sets WinsorLowerPercentile field to given value.

### HasWinsorLowerPercentile

`func (o *DenominatorPost) HasWinsorLowerPercentile() bool`

HasWinsorLowerPercentile returns a boolean if a field has been set.

### GetWinsorUpperPercentile

`func (o *DenominatorPost) GetWinsorUpperPercentile() float32`

GetWinsorUpperPercentile returns the WinsorUpperPercentile field if non-nil, zero value otherwise.

### GetWinsorUpperPercentileOk

`func (o *DenominatorPost) GetWinsorUpperPercentileOk() (*float32, bool)`

GetWinsorUpperPercentileOk returns a tuple with the WinsorUpperPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorUpperPercentile

`func (o *DenominatorPost) SetWinsorUpperPercentile(v float32)`

SetWinsorUpperPercentile sets WinsorUpperPercentile field to given value.

### HasWinsorUpperPercentile

`func (o *DenominatorPost) HasWinsorUpperPercentile() bool`

HasWinsorUpperPercentile returns a boolean if a field has been set.

### GetWinsorIncludeImputed

`func (o *DenominatorPost) GetWinsorIncludeImputed() bool`

GetWinsorIncludeImputed returns the WinsorIncludeImputed field if non-nil, zero value otherwise.

### GetWinsorIncludeImputedOk

`func (o *DenominatorPost) GetWinsorIncludeImputedOk() (*bool, bool)`

GetWinsorIncludeImputedOk returns a tuple with the WinsorIncludeImputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsorIncludeImputed

`func (o *DenominatorPost) SetWinsorIncludeImputed(v bool)`

SetWinsorIncludeImputed sets WinsorIncludeImputed field to given value.

### HasWinsorIncludeImputed

`func (o *DenominatorPost) HasWinsorIncludeImputed() bool`

HasWinsorIncludeImputed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


