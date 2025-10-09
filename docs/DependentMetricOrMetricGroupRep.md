# DependentMetricOrMetricGroupRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key to reference the metric or metric group | 
**VersionId** | **string** | The version ID of the metric or metric group | 
**Name** | **string** | A human-friendly name for the metric or metric group | 
**Kind** | **string** | If this is a metric, then it represents the kind of event the metric tracks. If this is a metric group, then it represents the group type | 
**IsNumeric** | Pointer to **bool** | For custom metrics, whether to track numeric changes in value against a baseline (&lt;code&gt;true&lt;/code&gt;) or to track a conversion when an end user takes an action (&lt;code&gt;false&lt;/code&gt;). | [optional] 
**EventKey** | Pointer to **string** | The event key sent with the metric. Only relevant for custom metrics. | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**IsGroup** | **bool** | Whether this is a metric group or a metric | 
**Metrics** | Pointer to [**[]MetricInGroupRep**](MetricInGroupRep.md) | An ordered list of the metrics in this metric group | [optional] 

## Methods

### NewDependentMetricOrMetricGroupRep

`func NewDependentMetricOrMetricGroupRep(key string, versionId string, name string, kind string, links map[string]Link, isGroup bool, ) *DependentMetricOrMetricGroupRep`

NewDependentMetricOrMetricGroupRep instantiates a new DependentMetricOrMetricGroupRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentMetricOrMetricGroupRepWithDefaults

`func NewDependentMetricOrMetricGroupRepWithDefaults() *DependentMetricOrMetricGroupRep`

NewDependentMetricOrMetricGroupRepWithDefaults instantiates a new DependentMetricOrMetricGroupRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DependentMetricOrMetricGroupRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentMetricOrMetricGroupRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentMetricOrMetricGroupRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersionId

`func (o *DependentMetricOrMetricGroupRep) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DependentMetricOrMetricGroupRep) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DependentMetricOrMetricGroupRep) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetName

`func (o *DependentMetricOrMetricGroupRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentMetricOrMetricGroupRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentMetricOrMetricGroupRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *DependentMetricOrMetricGroupRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DependentMetricOrMetricGroupRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DependentMetricOrMetricGroupRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIsNumeric

`func (o *DependentMetricOrMetricGroupRep) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *DependentMetricOrMetricGroupRep) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *DependentMetricOrMetricGroupRep) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.

### HasIsNumeric

`func (o *DependentMetricOrMetricGroupRep) HasIsNumeric() bool`

HasIsNumeric returns a boolean if a field has been set.

### GetEventKey

`func (o *DependentMetricOrMetricGroupRep) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *DependentMetricOrMetricGroupRep) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *DependentMetricOrMetricGroupRep) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *DependentMetricOrMetricGroupRep) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetLinks

`func (o *DependentMetricOrMetricGroupRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentMetricOrMetricGroupRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentMetricOrMetricGroupRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetIsGroup

`func (o *DependentMetricOrMetricGroupRep) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *DependentMetricOrMetricGroupRep) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *DependentMetricOrMetricGroupRep) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.


### GetMetrics

`func (o *DependentMetricOrMetricGroupRep) GetMetrics() []MetricInGroupRep`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DependentMetricOrMetricGroupRep) GetMetricsOk() (*[]MetricInGroupRep, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DependentMetricOrMetricGroupRep) SetMetrics(v []MetricInGroupRep)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DependentMetricOrMetricGroupRep) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


