# PatchIpAllowlistConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionAllowlistEnabled** | Pointer to **bool** | Enable or disable session allowlist | [optional] 
**ApiTokenAllowlistEnabled** | Pointer to **bool** | Enable or disable API token allowlist | [optional] 

## Methods

### NewPatchIpAllowlistConfigRequest

`func NewPatchIpAllowlistConfigRequest() *PatchIpAllowlistConfigRequest`

NewPatchIpAllowlistConfigRequest instantiates a new PatchIpAllowlistConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchIpAllowlistConfigRequestWithDefaults

`func NewPatchIpAllowlistConfigRequestWithDefaults() *PatchIpAllowlistConfigRequest`

NewPatchIpAllowlistConfigRequestWithDefaults instantiates a new PatchIpAllowlistConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionAllowlistEnabled

`func (o *PatchIpAllowlistConfigRequest) GetSessionAllowlistEnabled() bool`

GetSessionAllowlistEnabled returns the SessionAllowlistEnabled field if non-nil, zero value otherwise.

### GetSessionAllowlistEnabledOk

`func (o *PatchIpAllowlistConfigRequest) GetSessionAllowlistEnabledOk() (*bool, bool)`

GetSessionAllowlistEnabledOk returns a tuple with the SessionAllowlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAllowlistEnabled

`func (o *PatchIpAllowlistConfigRequest) SetSessionAllowlistEnabled(v bool)`

SetSessionAllowlistEnabled sets SessionAllowlistEnabled field to given value.

### HasSessionAllowlistEnabled

`func (o *PatchIpAllowlistConfigRequest) HasSessionAllowlistEnabled() bool`

HasSessionAllowlistEnabled returns a boolean if a field has been set.

### GetApiTokenAllowlistEnabled

`func (o *PatchIpAllowlistConfigRequest) GetApiTokenAllowlistEnabled() bool`

GetApiTokenAllowlistEnabled returns the ApiTokenAllowlistEnabled field if non-nil, zero value otherwise.

### GetApiTokenAllowlistEnabledOk

`func (o *PatchIpAllowlistConfigRequest) GetApiTokenAllowlistEnabledOk() (*bool, bool)`

GetApiTokenAllowlistEnabledOk returns a tuple with the ApiTokenAllowlistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenAllowlistEnabled

`func (o *PatchIpAllowlistConfigRequest) SetApiTokenAllowlistEnabled(v bool)`

SetApiTokenAllowlistEnabled sets ApiTokenAllowlistEnabled field to given value.

### HasApiTokenAllowlistEnabled

`func (o *PatchIpAllowlistConfigRequest) HasApiTokenAllowlistEnabled() bool`

HasApiTokenAllowlistEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


