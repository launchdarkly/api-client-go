# EnvironmentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-friendly name for the new environment. | [optional] 
**Key** | Pointer to **string** | A project-unique key for the new environment. | [optional] 
**Color** | Pointer to **string** | A color to indicate this environment in the UI. | [optional] 
**DefaultTtl** | Pointer to **int32** | The default time (in minutes) that the PHP SDK can cache feature flag rules locally. | [optional] 
**SecureMode** | Pointer to **bool** | Secure mode ensures that a user of the client-side SDK cannot impersonate another user. | [optional] 
**DefaultTrackEvents** | Pointer to **bool** | Enables tracking detailed information for new flags by default. | [optional] 
**ConfirmChanges** | Pointer to **bool** | Require confirmation for all flag and segment changes via the UI in this environment. | [optional] 
**RequireComments** | Pointer to **bool** | Require comments for all flag and segment changes via the UI in this environment. | [optional] 
**Tags** | Pointer to **[]string** | Tags to apply to the new environment. | [optional] 

## Methods

### NewEnvironmentPost

`func NewEnvironmentPost() *EnvironmentPost`

NewEnvironmentPost instantiates a new EnvironmentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentPostWithDefaults

`func NewEnvironmentPostWithDefaults() *EnvironmentPost`

NewEnvironmentPostWithDefaults instantiates a new EnvironmentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *EnvironmentPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentPost) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EnvironmentPost) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetColor

`func (o *EnvironmentPost) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EnvironmentPost) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EnvironmentPost) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *EnvironmentPost) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDefaultTtl

`func (o *EnvironmentPost) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *EnvironmentPost) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *EnvironmentPost) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *EnvironmentPost) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### GetSecureMode

`func (o *EnvironmentPost) GetSecureMode() bool`

GetSecureMode returns the SecureMode field if non-nil, zero value otherwise.

### GetSecureModeOk

`func (o *EnvironmentPost) GetSecureModeOk() (*bool, bool)`

GetSecureModeOk returns a tuple with the SecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureMode

`func (o *EnvironmentPost) SetSecureMode(v bool)`

SetSecureMode sets SecureMode field to given value.

### HasSecureMode

`func (o *EnvironmentPost) HasSecureMode() bool`

HasSecureMode returns a boolean if a field has been set.

### GetDefaultTrackEvents

`func (o *EnvironmentPost) GetDefaultTrackEvents() bool`

GetDefaultTrackEvents returns the DefaultTrackEvents field if non-nil, zero value otherwise.

### GetDefaultTrackEventsOk

`func (o *EnvironmentPost) GetDefaultTrackEventsOk() (*bool, bool)`

GetDefaultTrackEventsOk returns a tuple with the DefaultTrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTrackEvents

`func (o *EnvironmentPost) SetDefaultTrackEvents(v bool)`

SetDefaultTrackEvents sets DefaultTrackEvents field to given value.

### HasDefaultTrackEvents

`func (o *EnvironmentPost) HasDefaultTrackEvents() bool`

HasDefaultTrackEvents returns a boolean if a field has been set.

### GetConfirmChanges

`func (o *EnvironmentPost) GetConfirmChanges() bool`

GetConfirmChanges returns the ConfirmChanges field if non-nil, zero value otherwise.

### GetConfirmChangesOk

`func (o *EnvironmentPost) GetConfirmChangesOk() (*bool, bool)`

GetConfirmChangesOk returns a tuple with the ConfirmChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmChanges

`func (o *EnvironmentPost) SetConfirmChanges(v bool)`

SetConfirmChanges sets ConfirmChanges field to given value.

### HasConfirmChanges

`func (o *EnvironmentPost) HasConfirmChanges() bool`

HasConfirmChanges returns a boolean if a field has been set.

### GetRequireComments

`func (o *EnvironmentPost) GetRequireComments() bool`

GetRequireComments returns the RequireComments field if non-nil, zero value otherwise.

### GetRequireCommentsOk

`func (o *EnvironmentPost) GetRequireCommentsOk() (*bool, bool)`

GetRequireCommentsOk returns a tuple with the RequireComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireComments

`func (o *EnvironmentPost) SetRequireComments(v bool)`

SetRequireComments sets RequireComments field to given value.

### HasRequireComments

`func (o *EnvironmentPost) HasRequireComments() bool`

HasRequireComments returns a boolean if a field has been set.

### GetTags

`func (o *EnvironmentPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EnvironmentPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EnvironmentPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EnvironmentPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


