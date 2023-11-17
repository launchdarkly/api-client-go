# MetricInGroupRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The metric key | 
**Name** | **string** | The metric name | 
**Kind** | **string** | The kind of event the metric tracks | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**NameInGroup** | **string** | Name of the metric when used within the associated metric group. Can be different from the original name of the metric | 
**RandomizationUnits** | Pointer to **[]string** | The randomization units for the metric | [optional] 

## Methods

### NewMetricInGroupRep

`func NewMetricInGroupRep(key string, name string, kind string, links map[string]Link, nameInGroup string, ) *MetricInGroupRep`

NewMetricInGroupRep instantiates a new MetricInGroupRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricInGroupRepWithDefaults

`func NewMetricInGroupRepWithDefaults() *MetricInGroupRep`

NewMetricInGroupRepWithDefaults instantiates a new MetricInGroupRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricInGroupRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricInGroupRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricInGroupRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MetricInGroupRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricInGroupRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricInGroupRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *MetricInGroupRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricInGroupRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricInGroupRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLinks

`func (o *MetricInGroupRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricInGroupRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricInGroupRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetNameInGroup

`func (o *MetricInGroupRep) GetNameInGroup() string`

GetNameInGroup returns the NameInGroup field if non-nil, zero value otherwise.

### GetNameInGroupOk

`func (o *MetricInGroupRep) GetNameInGroupOk() (*string, bool)`

GetNameInGroupOk returns a tuple with the NameInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameInGroup

`func (o *MetricInGroupRep) SetNameInGroup(v string)`

SetNameInGroup sets NameInGroup field to given value.


### GetRandomizationUnits

`func (o *MetricInGroupRep) GetRandomizationUnits() []string`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *MetricInGroupRep) GetRandomizationUnitsOk() (*[]string, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *MetricInGroupRep) SetRandomizationUnits(v []string)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *MetricInGroupRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


