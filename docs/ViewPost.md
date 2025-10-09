# ViewPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Unique key for the view within the account/project | 
**Name** | **string** | Human-readable name for the view | 
**Description** | Pointer to **string** | Optional detailed description of the view | [optional] [default to ""]
**GenerateSdkKeys** | Pointer to **bool** | Whether to generate SDK keys for this view | [optional] [default to false]
**MaintainerId** | Pointer to **string** | Member ID of the maintainer for this view. Only one of &#x60;maintainerId&#x60; or &#x60;maintainerTeamKey&#x60; can be specified. | [optional] 
**MaintainerTeamKey** | Pointer to **string** | Key of the maintainer team for this view. Only one of &#x60;maintainerId&#x60; or &#x60;maintainerTeamKey&#x60; can be specified. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with this view | [optional] 

## Methods

### NewViewPost

`func NewViewPost(key string, name string, ) *ViewPost`

NewViewPost instantiates a new ViewPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPostWithDefaults

`func NewViewPostWithDefaults() *ViewPost`

NewViewPostWithDefaults instantiates a new ViewPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ViewPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ViewPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ViewPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ViewPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ViewPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateSdkKeys

`func (o *ViewPost) GetGenerateSdkKeys() bool`

GetGenerateSdkKeys returns the GenerateSdkKeys field if non-nil, zero value otherwise.

### GetGenerateSdkKeysOk

`func (o *ViewPost) GetGenerateSdkKeysOk() (*bool, bool)`

GetGenerateSdkKeysOk returns a tuple with the GenerateSdkKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSdkKeys

`func (o *ViewPost) SetGenerateSdkKeys(v bool)`

SetGenerateSdkKeys sets GenerateSdkKeys field to given value.

### HasGenerateSdkKeys

`func (o *ViewPost) HasGenerateSdkKeys() bool`

HasGenerateSdkKeys returns a boolean if a field has been set.

### GetMaintainerId

`func (o *ViewPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ViewPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ViewPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ViewPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *ViewPost) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *ViewPost) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *ViewPost) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *ViewPost) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetTags

`func (o *ViewPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


