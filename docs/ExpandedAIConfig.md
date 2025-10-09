# ExpandedAIConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A unique key used to reference the AI config | [optional] 
**Name** | Pointer to **string** | A human-friendly name for the AI config | [optional] 
**Tags** | Pointer to **[]string** | Tags for the AI config | [optional] 
**Description** | Pointer to **string** | Description of the AI config | [optional] 
**Version** | Pointer to **int32** | Version of the AI config | [optional] 
**CreatedAt** | Pointer to **int64** | Creation date in milliseconds | [optional] 
**UpdatedAt** | Pointer to **int64** | Last modification date in milliseconds | [optional] 
**FlagKey** | Pointer to **string** | Key of the flag that this AI config is attached to | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 

## Methods

### NewExpandedAIConfig

`func NewExpandedAIConfig() *ExpandedAIConfig`

NewExpandedAIConfig instantiates a new ExpandedAIConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedAIConfigWithDefaults

`func NewExpandedAIConfigWithDefaults() *ExpandedAIConfig`

NewExpandedAIConfigWithDefaults instantiates a new ExpandedAIConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExpandedAIConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedAIConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedAIConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ExpandedAIConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ExpandedAIConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedAIConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedAIConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpandedAIConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *ExpandedAIConfig) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExpandedAIConfig) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExpandedAIConfig) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExpandedAIConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDescription

`func (o *ExpandedAIConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandedAIConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandedAIConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandedAIConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *ExpandedAIConfig) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpandedAIConfig) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpandedAIConfig) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExpandedAIConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExpandedAIConfig) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExpandedAIConfig) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExpandedAIConfig) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExpandedAIConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExpandedAIConfig) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExpandedAIConfig) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExpandedAIConfig) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExpandedAIConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFlagKey

`func (o *ExpandedAIConfig) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *ExpandedAIConfig) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *ExpandedAIConfig) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *ExpandedAIConfig) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetLinks

`func (o *ExpandedAIConfig) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandedAIConfig) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandedAIConfig) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpandedAIConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


