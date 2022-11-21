# UpsertPayloadRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Tags** | **[]string** |  | 
**Temporary** | **bool** |  | 
**DefaultClientSideAvailability** | [**DefaultClientSideAvailability**](DefaultClientSideAvailability.md) |  | 
**BooleanDefaults** | [**BooleanFlagDefaults**](BooleanFlagDefaults.md) |  | 

## Methods

### NewUpsertPayloadRep

`func NewUpsertPayloadRep(tags []string, temporary bool, defaultClientSideAvailability DefaultClientSideAvailability, booleanDefaults BooleanFlagDefaults, ) *UpsertPayloadRep`

NewUpsertPayloadRep instantiates a new UpsertPayloadRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertPayloadRepWithDefaults

`func NewUpsertPayloadRepWithDefaults() *UpsertPayloadRep`

NewUpsertPayloadRepWithDefaults instantiates a new UpsertPayloadRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UpsertPayloadRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpsertPayloadRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpsertPayloadRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpsertPayloadRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTags

`func (o *UpsertPayloadRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpsertPayloadRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpsertPayloadRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTemporary

`func (o *UpsertPayloadRep) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *UpsertPayloadRep) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *UpsertPayloadRep) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.


### GetDefaultClientSideAvailability

`func (o *UpsertPayloadRep) GetDefaultClientSideAvailability() DefaultClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *UpsertPayloadRep) GetDefaultClientSideAvailabilityOk() (*DefaultClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *UpsertPayloadRep) SetDefaultClientSideAvailability(v DefaultClientSideAvailability)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.


### GetBooleanDefaults

`func (o *UpsertPayloadRep) GetBooleanDefaults() BooleanFlagDefaults`

GetBooleanDefaults returns the BooleanDefaults field if non-nil, zero value otherwise.

### GetBooleanDefaultsOk

`func (o *UpsertPayloadRep) GetBooleanDefaultsOk() (*BooleanFlagDefaults, bool)`

GetBooleanDefaultsOk returns a tuple with the BooleanDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanDefaults

`func (o *UpsertPayloadRep) SetBooleanDefaults(v BooleanFlagDefaults)`

SetBooleanDefaults sets BooleanDefaults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


