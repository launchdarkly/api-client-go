# MetricQueryFilterRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionKey** | **string** |  | 
**DimensionKind** | **string** |  | 
**DimensionKeyContextKind** | Pointer to **string** |  | [optional] 
**Op** | **string** |  | 
**Values** | **[]string** |  | 

## Methods

### NewMetricQueryFilterRep

`func NewMetricQueryFilterRep(dimensionKey string, dimensionKind string, op string, values []string, ) *MetricQueryFilterRep`

NewMetricQueryFilterRep instantiates a new MetricQueryFilterRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricQueryFilterRepWithDefaults

`func NewMetricQueryFilterRepWithDefaults() *MetricQueryFilterRep`

NewMetricQueryFilterRepWithDefaults instantiates a new MetricQueryFilterRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionKey

`func (o *MetricQueryFilterRep) GetDimensionKey() string`

GetDimensionKey returns the DimensionKey field if non-nil, zero value otherwise.

### GetDimensionKeyOk

`func (o *MetricQueryFilterRep) GetDimensionKeyOk() (*string, bool)`

GetDimensionKeyOk returns a tuple with the DimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKey

`func (o *MetricQueryFilterRep) SetDimensionKey(v string)`

SetDimensionKey sets DimensionKey field to given value.


### GetDimensionKind

`func (o *MetricQueryFilterRep) GetDimensionKind() string`

GetDimensionKind returns the DimensionKind field if non-nil, zero value otherwise.

### GetDimensionKindOk

`func (o *MetricQueryFilterRep) GetDimensionKindOk() (*string, bool)`

GetDimensionKindOk returns a tuple with the DimensionKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKind

`func (o *MetricQueryFilterRep) SetDimensionKind(v string)`

SetDimensionKind sets DimensionKind field to given value.


### GetDimensionKeyContextKind

`func (o *MetricQueryFilterRep) GetDimensionKeyContextKind() string`

GetDimensionKeyContextKind returns the DimensionKeyContextKind field if non-nil, zero value otherwise.

### GetDimensionKeyContextKindOk

`func (o *MetricQueryFilterRep) GetDimensionKeyContextKindOk() (*string, bool)`

GetDimensionKeyContextKindOk returns a tuple with the DimensionKeyContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKeyContextKind

`func (o *MetricQueryFilterRep) SetDimensionKeyContextKind(v string)`

SetDimensionKeyContextKind sets DimensionKeyContextKind field to given value.

### HasDimensionKeyContextKind

`func (o *MetricQueryFilterRep) HasDimensionKeyContextKind() bool`

HasDimensionKeyContextKind returns a boolean if a field has been set.

### GetOp

`func (o *MetricQueryFilterRep) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MetricQueryFilterRep) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MetricQueryFilterRep) SetOp(v string)`

SetOp sets Op field to given value.


### GetValues

`func (o *MetricQueryFilterRep) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricQueryFilterRep) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricQueryFilterRep) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


