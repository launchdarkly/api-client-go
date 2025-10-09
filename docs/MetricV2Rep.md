# MetricV2Rep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The metric key | 
**VersionId** | Pointer to **string** | The version ID of the metric | [optional] 
**Name** | **string** | The metric name | 
**Kind** | **string** | The kind of event the metric tracks | 
**IsNumeric** | Pointer to **bool** | For custom metrics, whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). | [optional] 
**UnitAggregationType** | Pointer to **string** | The type of unit aggregation to use for the metric | [optional] 
**EventKey** | Pointer to **string** | The event key sent with the metric. Only relevant for custom metrics. | [optional] 
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


### GetVersionId

`func (o *MetricV2Rep) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *MetricV2Rep) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *MetricV2Rep) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *MetricV2Rep) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

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


### GetIsNumeric

`func (o *MetricV2Rep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MetricV2Rep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MetricV2Rep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *MetricV2Rep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetUnitAggregationType

`func (o *MetricV2Rep) GetUnitAggregationType() string`

GetUnitAggregationType returns the UnitAggregationType field if non-nil, zero value otherwise.

### GetUnitAggregationTypeOk

`func (o *MetricV2Rep) GetUnitAggregationTypeOk() (*string, bool)`

GetUnitAggregationTypeOk returns a tuple with the UnitAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAggregationType

`func (o *MetricV2Rep) SetUnitAggregationType(v string)`

SetUnitAggregationType sets UnitAggregationType field to given value.

### HasUnitAggregationType

`func (o *MetricV2Rep) HasUnitAggregationType() bool`

HasUnitAggregationType returns a boolean if a field has been set.

### GetEventKey

`func (o *MetricV2Rep) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *MetricV2Rep) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *MetricV2Rep) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *MetricV2Rep) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

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


