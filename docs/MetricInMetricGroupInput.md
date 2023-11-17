# MetricInMetricGroupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The metric key | 
**NameInGroup** | **string** | Name of the metric when used within the associated metric group. Can be different from the original name of the metric | 

## Methods

### NewMetricInMetricGroupInput

`func NewMetricInMetricGroupInput(key string, nameInGroup string, ) *MetricInMetricGroupInput`

NewMetricInMetricGroupInput instantiates a new MetricInMetricGroupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricInMetricGroupInputWithDefaults

`func NewMetricInMetricGroupInputWithDefaults() *MetricInMetricGroupInput`

NewMetricInMetricGroupInputWithDefaults instantiates a new MetricInMetricGroupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricInMetricGroupInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricInMetricGroupInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricInMetricGroupInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetNameInGroup

`func (o *MetricInMetricGroupInput) GetNameInGroup() string`

GetNameInGroup returns the NameInGroup field if non-nil, zero value otherwise.

### GetNameInGroupOk

`func (o *MetricInMetricGroupInput) GetNameInGroupOk() (*string, bool)`

GetNameInGroupOk returns a tuple with the NameInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameInGroup

`func (o *MetricInMetricGroupInput) SetNameInGroup(v string)`

SetNameInGroup sets NameInGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


