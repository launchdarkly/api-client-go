# DependentMetricGroupRepWithMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key to reference the metric group | 
**Name** | **string** | A human-friendly name for the metric group | 
**Kind** | **string** | The type of the metric group | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Metrics** | Pointer to [**[]MetricInGroupRep**](MetricInGroupRep.md) | The metrics in the metric group | [optional] 

## Methods

### NewDependentMetricGroupRepWithMetrics

`func NewDependentMetricGroupRepWithMetrics(key string, name string, kind string, links map[string]Link, ) *DependentMetricGroupRepWithMetrics`

NewDependentMetricGroupRepWithMetrics instantiates a new DependentMetricGroupRepWithMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentMetricGroupRepWithMetricsWithDefaults

`func NewDependentMetricGroupRepWithMetricsWithDefaults() *DependentMetricGroupRepWithMetrics`

NewDependentMetricGroupRepWithMetricsWithDefaults instantiates a new DependentMetricGroupRepWithMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DependentMetricGroupRepWithMetrics) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentMetricGroupRepWithMetrics) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentMetricGroupRepWithMetrics) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DependentMetricGroupRepWithMetrics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentMetricGroupRepWithMetrics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentMetricGroupRepWithMetrics) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *DependentMetricGroupRepWithMetrics) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DependentMetricGroupRepWithMetrics) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DependentMetricGroupRepWithMetrics) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLinks

`func (o *DependentMetricGroupRepWithMetrics) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentMetricGroupRepWithMetrics) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentMetricGroupRepWithMetrics) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetMetrics

`func (o *DependentMetricGroupRepWithMetrics) GetMetrics() []MetricInGroupRep`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DependentMetricGroupRepWithMetrics) GetMetricsOk() (*[]MetricInGroupRep, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DependentMetricGroupRepWithMetrics) SetMetrics(v []MetricInGroupRep)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DependentMetricGroupRepWithMetrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


