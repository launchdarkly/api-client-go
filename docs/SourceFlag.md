# SourceFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Flag key to copy | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewSourceFlag

`func NewSourceFlag(key string, ) *SourceFlag`

NewSourceFlag instantiates a new SourceFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceFlagWithDefaults

`func NewSourceFlagWithDefaults() *SourceFlag`

NewSourceFlagWithDefaults instantiates a new SourceFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SourceFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SourceFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SourceFlag) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *SourceFlag) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SourceFlag) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SourceFlag) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SourceFlag) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


