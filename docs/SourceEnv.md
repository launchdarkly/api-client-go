# SourceEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the source environment to clone from | [optional] 
**Version** | Pointer to **int32** | (Optional) The version number of the source environment to clone from. Used for optimistic locking | [optional] 

## Methods

### NewSourceEnv

`func NewSourceEnv() *SourceEnv`

NewSourceEnv instantiates a new SourceEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceEnvWithDefaults

`func NewSourceEnvWithDefaults() *SourceEnv`

NewSourceEnvWithDefaults instantiates a new SourceEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SourceEnv) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SourceEnv) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SourceEnv) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SourceEnv) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVersion

`func (o *SourceEnv) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SourceEnv) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SourceEnv) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SourceEnv) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


