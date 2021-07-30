# DependentFlagWithEnvs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Environments** | [**[]DependentFlagEnvironment**](DependentFlagEnvironment.md) |  | 

## Methods

### NewDependentFlagWithEnvs

`func NewDependentFlagWithEnvs(key string, environments []DependentFlagEnvironment, ) *DependentFlagWithEnvs`

NewDependentFlagWithEnvs instantiates a new DependentFlagWithEnvs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentFlagWithEnvsWithDefaults

`func NewDependentFlagWithEnvsWithDefaults() *DependentFlagWithEnvs`

NewDependentFlagWithEnvsWithDefaults instantiates a new DependentFlagWithEnvs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DependentFlagWithEnvs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentFlagWithEnvs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentFlagWithEnvs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DependentFlagWithEnvs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *DependentFlagWithEnvs) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentFlagWithEnvs) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentFlagWithEnvs) SetKey(v string)`

SetKey sets Key field to given value.


### GetEnvironments

`func (o *DependentFlagWithEnvs) GetEnvironments() []DependentFlagEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *DependentFlagWithEnvs) GetEnvironmentsOk() (*[]DependentFlagEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *DependentFlagWithEnvs) SetEnvironments(v []DependentFlagEnvironment)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


