# FlagCopyConfigEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The environment key | 
**CurrentVersion** | Pointer to **int32** | Optional flag version. If you include this, the operation only succeeds if the current flag version in the environment matches this version. | [optional] 

## Methods

### NewFlagCopyConfigEnvironment

`func NewFlagCopyConfigEnvironment(key string, ) *FlagCopyConfigEnvironment`

NewFlagCopyConfigEnvironment instantiates a new FlagCopyConfigEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagCopyConfigEnvironmentWithDefaults

`func NewFlagCopyConfigEnvironmentWithDefaults() *FlagCopyConfigEnvironment`

NewFlagCopyConfigEnvironmentWithDefaults instantiates a new FlagCopyConfigEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlagCopyConfigEnvironment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagCopyConfigEnvironment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagCopyConfigEnvironment) SetKey(v string)`

SetKey sets Key field to given value.


### GetCurrentVersion

`func (o *FlagCopyConfigEnvironment) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *FlagCopyConfigEnvironment) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *FlagCopyConfigEnvironment) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *FlagCopyConfigEnvironment) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


