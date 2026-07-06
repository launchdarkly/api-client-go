# ExpandedExperimentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the experiment | 
**Name** | **string** | The name of the experiment | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 

## Methods

### NewExpandedExperimentRep

`func NewExpandedExperimentRep(key string, name string, ) *ExpandedExperimentRep`

NewExpandedExperimentRep instantiates a new ExpandedExperimentRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedExperimentRepWithDefaults

`func NewExpandedExperimentRepWithDefaults() *ExpandedExperimentRep`

NewExpandedExperimentRepWithDefaults instantiates a new ExpandedExperimentRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExpandedExperimentRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedExperimentRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedExperimentRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ExpandedExperimentRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedExperimentRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedExperimentRep) SetName(v string)`

SetName sets Name field to given value.


### GetAccess

`func (o *ExpandedExperimentRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ExpandedExperimentRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ExpandedExperimentRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ExpandedExperimentRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


