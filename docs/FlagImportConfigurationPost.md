# FlagImportConfigurationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** |  | 
**Tags** | Pointer to **[]string** | Tags to associate with the configuration | [optional] 
**Name** | Pointer to **string** | Name to identify the configuration | [optional] 

## Methods

### NewFlagImportConfigurationPost

`func NewFlagImportConfigurationPost(config map[string]interface{}, ) *FlagImportConfigurationPost`

NewFlagImportConfigurationPost instantiates a new FlagImportConfigurationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagImportConfigurationPostWithDefaults

`func NewFlagImportConfigurationPostWithDefaults() *FlagImportConfigurationPost`

NewFlagImportConfigurationPostWithDefaults instantiates a new FlagImportConfigurationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *FlagImportConfigurationPost) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FlagImportConfigurationPost) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FlagImportConfigurationPost) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetTags

`func (o *FlagImportConfigurationPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlagImportConfigurationPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlagImportConfigurationPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlagImportConfigurationPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *FlagImportConfigurationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagImportConfigurationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagImportConfigurationPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlagImportConfigurationPost) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


