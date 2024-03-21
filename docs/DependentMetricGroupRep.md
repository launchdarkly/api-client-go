# DependentMetricGroupRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key to reference the metric group | 
**Name** | **string** | A human-friendly name for the metric group | 
**Kind** | **string** | The type of the metric group | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewDependentMetricGroupRep

`func NewDependentMetricGroupRep(key string, name string, kind string, links map[string]Link, ) *DependentMetricGroupRep`

NewDependentMetricGroupRep instantiates a new DependentMetricGroupRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentMetricGroupRepWithDefaults

`func NewDependentMetricGroupRepWithDefaults() *DependentMetricGroupRep`

NewDependentMetricGroupRepWithDefaults instantiates a new DependentMetricGroupRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DependentMetricGroupRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentMetricGroupRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentMetricGroupRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DependentMetricGroupRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentMetricGroupRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentMetricGroupRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *DependentMetricGroupRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DependentMetricGroupRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DependentMetricGroupRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLinks

`func (o *DependentMetricGroupRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentMetricGroupRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentMetricGroupRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


