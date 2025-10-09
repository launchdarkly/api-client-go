# ViewPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable name for the view | [optional] 
**Description** | Pointer to **string** | Optional detailed description of the view | [optional] 
**GenerateSdkKeys** | Pointer to **bool** | Whether to generate SDK keys for this view | [optional] 
**MaintainerId** | Pointer to **string** | Member ID of the maintainer for this view. Only one of &#x60;maintainerId&#x60; or &#x60;maintainerTeamKey&#x60; can be specified. | [optional] 
**MaintainerTeamKey** | Pointer to **string** | Key of the maintainer team for this view. Only one of &#x60;maintainerId&#x60; or &#x60;maintainerTeamKey&#x60; can be specified. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with this view | [optional] 
**Archived** | Pointer to **bool** | Whether or not the view is archived | [optional] 

## Methods

### NewViewPatch

`func NewViewPatch() *ViewPatch`

NewViewPatch instantiates a new ViewPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPatchWithDefaults

`func NewViewPatchWithDefaults() *ViewPatch`

NewViewPatchWithDefaults instantiates a new ViewPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ViewPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ViewPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateSdkKeys

`func (o *ViewPatch) GetGenerateSdkKeys() bool`

GetGenerateSdkKeys returns the GenerateSdkKeys field if non-nil, zero value otherwise.

### GetGenerateSdkKeysOk

`func (o *ViewPatch) GetGenerateSdkKeysOk() (*bool, bool)`

GetGenerateSdkKeysOk returns a tuple with the GenerateSdkKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSdkKeys

`func (o *ViewPatch) SetGenerateSdkKeys(v bool)`

SetGenerateSdkKeys sets GenerateSdkKeys field to given value.

### HasGenerateSdkKeys

`func (o *ViewPatch) HasGenerateSdkKeys() bool`

HasGenerateSdkKeys returns a boolean if a field has been set.

### GetMaintainerId

`func (o *ViewPatch) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ViewPatch) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ViewPatch) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ViewPatch) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *ViewPatch) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *ViewPatch) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *ViewPatch) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *ViewPatch) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetTags

`func (o *ViewPatch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ViewPatch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ViewPatch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ViewPatch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetArchived

`func (o *ViewPatch) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ViewPatch) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ViewPatch) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ViewPatch) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


