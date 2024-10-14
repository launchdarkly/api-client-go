# LayerRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the layer | 
**Name** | **string** | The name of the layer | 
**Description** | **string** | The description of the layer | 
**CreatedAt** | **int64** |  | 
**RandomizationUnit** | Pointer to **string** | The unit of randomization for the layer | [optional] 
**Environments** | Pointer to [**map[string]LayerConfigurationRep**](LayerConfigurationRep.md) | The layer configurations for each requested environment | [optional] 

## Methods

### NewLayerRep

`func NewLayerRep(key string, name string, description string, createdAt int64, ) *LayerRep`

NewLayerRep instantiates a new LayerRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerRepWithDefaults

`func NewLayerRepWithDefaults() *LayerRep`

NewLayerRepWithDefaults instantiates a new LayerRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LayerRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LayerRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LayerRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *LayerRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayerRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayerRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LayerRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LayerRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LayerRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *LayerRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LayerRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LayerRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetRandomizationUnit

`func (o *LayerRep) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *LayerRep) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *LayerRep) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *LayerRep) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetEnvironments

`func (o *LayerRep) GetEnvironments() map[string]LayerConfigurationRep`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *LayerRep) GetEnvironmentsOk() (*map[string]LayerConfigurationRep, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *LayerRep) SetEnvironments(v map[string]LayerConfigurationRep)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *LayerRep) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


