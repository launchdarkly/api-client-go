# SimpleHoldoutRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Experiments** | Pointer to [**[]RelatedExperimentRep**](RelatedExperimentRep.md) |  | [optional] 

## Methods

### NewSimpleHoldoutRep

`func NewSimpleHoldoutRep() *SimpleHoldoutRep`

NewSimpleHoldoutRep instantiates a new SimpleHoldoutRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleHoldoutRepWithDefaults

`func NewSimpleHoldoutRepWithDefaults() *SimpleHoldoutRep`

NewSimpleHoldoutRepWithDefaults instantiates a new SimpleHoldoutRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleHoldoutRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleHoldoutRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleHoldoutRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleHoldoutRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *SimpleHoldoutRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SimpleHoldoutRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SimpleHoldoutRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SimpleHoldoutRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *SimpleHoldoutRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleHoldoutRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleHoldoutRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleHoldoutRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SimpleHoldoutRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleHoldoutRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleHoldoutRep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimpleHoldoutRep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimpleHoldoutRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimpleHoldoutRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimpleHoldoutRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SimpleHoldoutRep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SimpleHoldoutRep) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimpleHoldoutRep) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimpleHoldoutRep) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SimpleHoldoutRep) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExperiments

`func (o *SimpleHoldoutRep) GetExperiments() []RelatedExperimentRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *SimpleHoldoutRep) GetExperimentsOk() (*[]RelatedExperimentRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *SimpleHoldoutRep) SetExperiments(v []RelatedExperimentRep)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *SimpleHoldoutRep) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


