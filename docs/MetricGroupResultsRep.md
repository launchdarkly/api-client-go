# MetricGroupResultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Metrics** | [**[]MetricInGroupResultsRep**](MetricInGroupResultsRep.md) | An ordered list of the metrics in this metric group, and each of their results | 

## Methods

### NewMetricGroupResultsRep

`func NewMetricGroupResultsRep(links map[string]Link, metrics []MetricInGroupResultsRep, ) *MetricGroupResultsRep`

NewMetricGroupResultsRep instantiates a new MetricGroupResultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGroupResultsRepWithDefaults

`func NewMetricGroupResultsRepWithDefaults() *MetricGroupResultsRep`

NewMetricGroupResultsRepWithDefaults instantiates a new MetricGroupResultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *MetricGroupResultsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricGroupResultsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricGroupResultsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetMetrics

`func (o *MetricGroupResultsRep) GetMetrics() []MetricInGroupResultsRep`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricGroupResultsRep) GetMetricsOk() (*[]MetricInGroupResultsRep, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricGroupResultsRep) SetMetrics(v []MetricInGroupResultsRep)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


