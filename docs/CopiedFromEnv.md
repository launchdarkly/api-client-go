# CopiedFromEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of feature flag copied | 
**Version** | Pointer to **int32** | Version of feature flag copied | [optional] 

## Methods

### NewCopiedFromEnv

`func NewCopiedFromEnv(key string, ) *CopiedFromEnv`

NewCopiedFromEnv instantiates a new CopiedFromEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopiedFromEnvWithDefaults

`func NewCopiedFromEnvWithDefaults() *CopiedFromEnv`

NewCopiedFromEnvWithDefaults instantiates a new CopiedFromEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CopiedFromEnv) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CopiedFromEnv) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CopiedFromEnv) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *CopiedFromEnv) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CopiedFromEnv) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CopiedFromEnv) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CopiedFromEnv) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


