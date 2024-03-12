# FlagDefaultsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Key** | Pointer to **string** | A unique key for the flag default | [optional] 
**Tags** | Pointer to **[]string** | A list of default tags for each flag | [optional] 
**Temporary** | Pointer to **bool** | Whether the flag should be temporary by default | [optional] 
**DefaultClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**BooleanDefaults** | Pointer to [**BooleanDefaults**](BooleanDefaults.md) |  | [optional] 

## Methods

### NewFlagDefaultsRep

`func NewFlagDefaultsRep() *FlagDefaultsRep`

NewFlagDefaultsRep instantiates a new FlagDefaultsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagDefaultsRepWithDefaults

`func NewFlagDefaultsRepWithDefaults() *FlagDefaultsRep`

NewFlagDefaultsRepWithDefaults instantiates a new FlagDefaultsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagDefaultsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagDefaultsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagDefaultsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagDefaultsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetKey

`func (o *FlagDefaultsRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagDefaultsRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagDefaultsRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlagDefaultsRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTags

`func (o *FlagDefaultsRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagDefaultsRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagDefaultsRep) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlagDefaultsRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemporary

`func (o *FlagDefaultsRep) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *FlagDefaultsRep) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *FlagDefaultsRep) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *FlagDefaultsRep) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetDefaultClientSideAvailability

`func (o *FlagDefaultsRep) GetDefaultClientSideAvailability() ClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *FlagDefaultsRep) GetDefaultClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *FlagDefaultsRep) SetDefaultClientSideAvailability(v ClientSideAvailability)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.

### HasDefaultClientSideAvailability

`func (o *FlagDefaultsRep) HasDefaultClientSideAvailability() bool`

HasDefaultClientSideAvailability returns a boolean if a field has been set.

### GetBooleanDefaults

`func (o *FlagDefaultsRep) GetBooleanDefaults() BooleanDefaults`

GetBooleanDefaults returns the BooleanDefaults field if non-nil, zero value otherwise.

### GetBooleanDefaultsOk

`func (o *FlagDefaultsRep) GetBooleanDefaultsOk() (*BooleanDefaults, bool)`

GetBooleanDefaultsOk returns a tuple with the BooleanDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanDefaults

`func (o *FlagDefaultsRep) SetBooleanDefaults(v BooleanDefaults)`

SetBooleanDefaults sets BooleanDefaults field to given value.

### HasBooleanDefaults

`func (o *FlagDefaultsRep) HasBooleanDefaults() bool`

HasBooleanDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


