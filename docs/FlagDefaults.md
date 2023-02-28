# FlagDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | A list of default tags for each flag | [optional] 
**Temporary** | Pointer to **bool** | Whether the flag should be temporary by default | [optional] 
**DefaultClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**BooleanDefaults** | Pointer to [**BooleanDefaults**](BooleanDefaults.md) |  | [optional] 

## Methods

### NewFlagDefaults

`func NewFlagDefaults() *FlagDefaults`

NewFlagDefaults instantiates a new FlagDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagDefaultsWithDefaults

`func NewFlagDefaultsWithDefaults() *FlagDefaults`

NewFlagDefaultsWithDefaults instantiates a new FlagDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *FlagDefaults) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagDefaults) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagDefaults) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlagDefaults) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemporary

`func (o *FlagDefaults) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FlagDefaults) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FlagDefaults) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *FlagDefaults) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetDefaultClientSideAvailability

`func (o *FlagDefaults) GetDefaultClientSideAvailability() ClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *FlagDefaults) GetDefaultClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *FlagDefaults) SetDefaultClientSideAvailability(v ClientSideAvailability)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.

### HasDefaultClientSideAvailability

`func (o *FlagDefaults) HasDefaultClientSideAvailability() bool`

HasDefaultClientSideAvailability returns a boolean if a field has been set.

### GetBooleanDefaults

`func (o *FlagDefaults) GetBooleanDefaults() BooleanDefaults`

GetBooleanDefaults returns the BooleanDefaults field if non-nil, zero value otherwise.

### GetBooleanDefaultsOk

`func (o *FlagDefaults) GetBooleanDefaultsOk() (*BooleanDefaults, bool)`

GetBooleanDefaultsOk returns a tuple with the BooleanDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanDefaults

`func (o *FlagDefaults) SetBooleanDefaults(v BooleanDefaults)`

SetBooleanDefaults sets BooleanDefaults field to given value.

### HasBooleanDefaults

`func (o *FlagDefaults) HasBooleanDefaults() bool`

HasBooleanDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


