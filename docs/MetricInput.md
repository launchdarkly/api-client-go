# MetricInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The metric key | 
**Primary** | **bool** | Whether this is a primary metric (true) or a secondary metric (false) | 

## Methods

### NewMetricInput

`func NewMetricInput(key string, primary bool, ) *MetricInput`

NewMetricInput instantiates a new MetricInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricInputWithDefaults

`func NewMetricInputWithDefaults() *MetricInput`

NewMetricInputWithDefaults instantiates a new MetricInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetPrimary

`func (o *MetricInput) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MetricInput) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MetricInput) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


