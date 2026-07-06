# SdkKeyPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | (Optional)The human-readable name of the SDK key. | [optional] 
**Description** | Pointer to **string** | (Optional) The description of the SDK key. | [optional] 
**Expiry** | Pointer to **int64** |  | [optional] 

## Methods

### NewSdkKeyPatch

`func NewSdkKeyPatch() *SdkKeyPatch`

NewSdkKeyPatch instantiates a new SdkKeyPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeyPatchWithDefaults

`func NewSdkKeyPatchWithDefaults() *SdkKeyPatch`

NewSdkKeyPatchWithDefaults instantiates a new SdkKeyPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SdkKeyPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdkKeyPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdkKeyPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdkKeyPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SdkKeyPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdkKeyPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdkKeyPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdkKeyPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiry

`func (o *SdkKeyPatch) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SdkKeyPatch) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SdkKeyPatch) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SdkKeyPatch) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


