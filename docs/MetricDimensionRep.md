# MetricDimensionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Dimension attribute key | 
**Kind** | **string** |  | 
**ContextKindKey** | Pointer to **string** | Key of the context kind that the dimension is for | [optional] 
**ContextKindName** | Pointer to **string** | Name of the context kind that the dimension is for | [optional] 

## Methods

### NewMetricDimensionRep

`func NewMetricDimensionRep(key string, kind string, ) *MetricDimensionRep`

NewMetricDimensionRep instantiates a new MetricDimensionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDimensionRepWithDefaults

`func NewMetricDimensionRepWithDefaults() *MetricDimensionRep`

NewMetricDimensionRepWithDefaults instantiates a new MetricDimensionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricDimensionRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricDimensionRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricDimensionRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *MetricDimensionRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricDimensionRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricDimensionRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetContextKindKey

`func (o *MetricDimensionRep) GetContextKindKey() string`

GetContextKindKey returns the ContextKindKey field if non-nil, zero value otherwise.

### GetContextKindKeyOk

`func (o *MetricDimensionRep) GetContextKindKeyOk() (*string, bool)`

GetContextKindKeyOk returns a tuple with the ContextKindKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKindKey

`func (o *MetricDimensionRep) SetContextKindKey(v string)`

SetContextKindKey sets ContextKindKey field to given value.

### HasContextKindKey

`func (o *MetricDimensionRep) HasContextKindKey() bool`

HasContextKindKey returns a boolean if a field has been set.

### GetContextKindName

`func (o *MetricDimensionRep) GetContextKindName() string`

GetContextKindName returns the ContextKindName field if non-nil, zero value otherwise.

### GetContextKindNameOk

`func (o *MetricDimensionRep) GetContextKindNameOk() (*string, bool)`

GetContextKindNameOk returns a tuple with the ContextKindName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKindName

`func (o *MetricDimensionRep) SetContextKindName(v string)`

SetContextKindName sets ContextKindName field to given value.

### HasContextKindName

`func (o *MetricDimensionRep) HasContextKindName() bool`

HasContextKindName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


