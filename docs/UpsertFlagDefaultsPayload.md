# UpsertFlagDefaultsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** |  | 
**Temporary** | **bool** |  | 
**DefaultClientSideAvailability** | [**DefaultClientSideAvailability**](DefaultClientSideAvailability.md) |  | 
**BooleanDefaults** | [**BooleanFlagDefaults**](BooleanFlagDefaults.md) |  | 

## Methods

### NewUpsertFlagDefaultsPayload

`func NewUpsertFlagDefaultsPayload(tags []string, temporary bool, defaultClientSideAvailability DefaultClientSideAvailability, booleanDefaults BooleanFlagDefaults, ) *UpsertFlagDefaultsPayload`

NewUpsertFlagDefaultsPayload instantiates a new UpsertFlagDefaultsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertFlagDefaultsPayloadWithDefaults

`func NewUpsertFlagDefaultsPayloadWithDefaults() *UpsertFlagDefaultsPayload`

NewUpsertFlagDefaultsPayloadWithDefaults instantiates a new UpsertFlagDefaultsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpsertFlagDefaultsPayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpsertFlagDefaultsPayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpsertFlagDefaultsPayload) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTemporary

`func (o *UpsertFlagDefaultsPayload) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *UpsertFlagDefaultsPayload) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *UpsertFlagDefaultsPayload) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.


### GetDefaultClientSideAvailability

`func (o *UpsertFlagDefaultsPayload) GetDefaultClientSideAvailability() DefaultClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *UpsertFlagDefaultsPayload) GetDefaultClientSideAvailabilityOk() (*DefaultClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *UpsertFlagDefaultsPayload) SetDefaultClientSideAvailability(v DefaultClientSideAvailability)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.


### GetBooleanDefaults

`func (o *UpsertFlagDefaultsPayload) GetBooleanDefaults() BooleanFlagDefaults`

GetBooleanDefaults returns the BooleanDefaults field if non-nil, zero value otherwise.

### GetBooleanDefaultsOk

`func (o *UpsertFlagDefaultsPayload) GetBooleanDefaultsOk() (*BooleanFlagDefaults, bool)`

GetBooleanDefaultsOk returns a tuple with the BooleanDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanDefaults

`func (o *UpsertFlagDefaultsPayload) SetBooleanDefaults(v BooleanFlagDefaults)`

SetBooleanDefaults sets BooleanDefaults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


