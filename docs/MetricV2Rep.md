# MetricV2Rep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The metric key | 
**Name** | **string** | The metric name | 
**Kind** | **string** | The kind of event the metric tracks | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewMetricV2Rep

`func NewMetricV2Rep(key string, name string, kind string, links map[string]Link, ) *MetricV2Rep`

NewMetricV2Rep instantiates a new MetricV2Rep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricV2RepWithDefaults

`func NewMetricV2RepWithDefaults() *MetricV2Rep`

NewMetricV2RepWithDefaults instantiates a new MetricV2Rep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricV2Rep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricV2Rep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricV2Rep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MetricV2Rep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricV2Rep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricV2Rep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *MetricV2Rep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricV2Rep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricV2Rep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLinks

`func (o *MetricV2Rep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricV2Rep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricV2Rep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


