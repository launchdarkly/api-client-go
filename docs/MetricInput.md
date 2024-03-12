# MetricInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The metric key | 
**IsGroup** | Pointer to **bool** | Whether this is a metric group (true) or a metric (false). Defaults to false | [optional] 
**Primary** | Pointer to **bool** | Deprecated, use &lt;code&gt;primarySingleMetricKey&lt;/code&gt; and &lt;code&gt;primaryFunnelKey&lt;/code&gt;. Whether this is a primary metric (true) or a secondary metric (false) | [optional] 

## Methods

### NewMetricInput

`func NewMetricInput(key string, ) *MetricInput`

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


### GetIsGroup

`func (o *MetricInput) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *MetricInput) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *MetricInput) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *MetricInput) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

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

### HasPrimary

`func (o *MetricInput) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


