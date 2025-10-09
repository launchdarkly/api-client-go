# AIConfigPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Key** | **string** |  | 
**MaintainerId** | Pointer to **string** |  | [optional] 
**MaintainerTeamKey** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "completion"]
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAIConfigPost

`func NewAIConfigPost(key string, name string, ) *AIConfigPost`

NewAIConfigPost instantiates a new AIConfigPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigPostWithDefaults

`func NewAIConfigPostWithDefaults() *AIConfigPost`

NewAIConfigPostWithDefaults instantiates a new AIConfigPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AIConfigPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIConfigPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIConfigPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIConfigPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *AIConfigPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIConfigPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIConfigPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaintainerId

`func (o *AIConfigPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AIConfigPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AIConfigPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AIConfigPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *AIConfigPost) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *AIConfigPost) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *AIConfigPost) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *AIConfigPost) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetMode

`func (o *AIConfigPost) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AIConfigPost) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AIConfigPost) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AIConfigPost) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *AIConfigPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfigPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfigPost) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *AIConfigPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AIConfigPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AIConfigPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AIConfigPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


