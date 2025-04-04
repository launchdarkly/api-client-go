# MetricByVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariationKey** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**Metrics**](Metrics.md) |  | [optional] 

## Methods

### NewMetricByVariation

`func NewMetricByVariation() *MetricByVariation`

NewMetricByVariation instantiates a new MetricByVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricByVariationWithDefaults

`func NewMetricByVariationWithDefaults() *MetricByVariation`

NewMetricByVariationWithDefaults instantiates a new MetricByVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariationKey

`func (o *MetricByVariation) GetVariationKey() string`

GetVariationKey returns the VariationKey field if non-nil, zero value otherwise.

### GetVariationKeyOk

`func (o *MetricByVariation) GetVariationKeyOk() (*string, bool)`

GetVariationKeyOk returns a tuple with the VariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationKey

`func (o *MetricByVariation) SetVariationKey(v string)`

SetVariationKey sets VariationKey field to given value.

### HasVariationKey

`func (o *MetricByVariation) HasVariationKey() bool`

HasVariationKey returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricByVariation) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricByVariation) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricByVariation) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricByVariation) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


