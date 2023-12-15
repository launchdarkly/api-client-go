# InternalEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Id** | **string** | The ID for the environment. Use this as the client-side ID for authorization in some client-side SDKs, and to associate LaunchDarkly environments with CDN integrations in edge SDKs. | 
**Key** | **string** | A project-unique key for the new environment | 
**Name** | **string** | A human-friendly name for the new environment | 
**ApiKey** | **string** | The SDK key for the environment. Use this for authorization in server-side SDKs. | 
**MobileKey** | **string** | The mobile key for the environment. Use this for authorization in mobile SDKs. | 
**Color** | **string** | The color used to indicate this environment in the UI | 
**DefaultTtl** | **int32** | The default time (in minutes) that the PHP SDK can cache feature flag rules locally | 
**SecureMode** | **bool** | Ensures that one end user of the client-side SDK cannot inspect the variations for another end user | 
**DefaultTrackEvents** | **bool** | Enables tracking detailed information for new flags by default | 
**RequireComments** | **bool** | Whether members who modify flags and segments through the LaunchDarkly user interface are required to add a comment | 
**ConfirmChanges** | **bool** | Whether members who modify flags and segments through the LaunchDarkly user interface are required to confirm those changes | 
**Tags** | **[]string** | A list of tags for this environment | 
**ApprovalSettings** | Pointer to [**ApprovalSettings**](ApprovalSettings.md) |  | [optional] 
**Kind** | Pointer to **string** | The kind of environment | [optional] 
**Pubnub** | [**InternalPubnubRep**](InternalPubnubRep.md) |  | 

## Methods

### NewInternalEnvironment

`func NewInternalEnvironment(links map[string]Link, id string, key string, name string, apiKey string, mobileKey string, color string, defaultTtl int32, secureMode bool, defaultTrackEvents bool, requireComments bool, confirmChanges bool, tags []string, pubnub InternalPubnubRep, ) *InternalEnvironment`

NewInternalEnvironment instantiates a new InternalEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalEnvironmentWithDefaults

`func NewInternalEnvironmentWithDefaults() *InternalEnvironment`

NewInternalEnvironmentWithDefaults instantiates a new InternalEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *InternalEnvironment) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InternalEnvironment) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InternalEnvironment) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *InternalEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalEnvironment) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *InternalEnvironment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InternalEnvironment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InternalEnvironment) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *InternalEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetApiKey

`func (o *InternalEnvironment) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *InternalEnvironment) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *InternalEnvironment) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetMobileKey

`func (o *InternalEnvironment) GetMobileKey() string`

GetMobileKey returns the MobileKey field if non-nil, zero value otherwise.

### GetMobileKeyOk

`func (o *InternalEnvironment) GetMobileKeyOk() (*string, bool)`

GetMobileKeyOk returns a tuple with the MobileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileKey

`func (o *InternalEnvironment) SetMobileKey(v string)`

SetMobileKey sets MobileKey field to given value.


### GetColor

`func (o *InternalEnvironment) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InternalEnvironment) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InternalEnvironment) SetColor(v string)`

SetColor sets Color field to given value.


### GetDefaultTtl

`func (o *InternalEnvironment) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *InternalEnvironment) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *InternalEnvironment) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.


### GetSecureMode

`func (o *InternalEnvironment) GetSecureMode() bool`

GetSecureMode returns the SecureMode field if non-nil, zero value otherwise.

### GetSecureModeOk

`func (o *InternalEnvironment) GetSecureModeOk() (*bool, bool)`

GetSecureModeOk returns a tuple with the SecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureMode

`func (o *InternalEnvironment) SetSecureMode(v bool)`

SetSecureMode sets SecureMode field to given value.


### GetDefaultTrackEvents

`func (o *InternalEnvironment) GetDefaultTrackEvents() bool`

GetDefaultTrackEvents returns the DefaultTrackEvents field if non-nil, zero value otherwise.

### GetDefaultTrackEventsOk

`func (o *InternalEnvironment) GetDefaultTrackEventsOk() (*bool, bool)`

GetDefaultTrackEventsOk returns a tuple with the DefaultTrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTrackEvents

`func (o *InternalEnvironment) SetDefaultTrackEvents(v bool)`

SetDefaultTrackEvents sets DefaultTrackEvents field to given value.


### GetRequireComments

`func (o *InternalEnvironment) GetRequireComments() bool`

GetRequireComments returns the RequireComments field if non-nil, zero value otherwise.

### GetRequireCommentsOk

`func (o *InternalEnvironment) GetRequireCommentsOk() (*bool, bool)`

GetRequireCommentsOk returns a tuple with the RequireComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireComments

`func (o *InternalEnvironment) SetRequireComments(v bool)`

SetRequireComments sets RequireComments field to given value.


### GetConfirmChanges

`func (o *InternalEnvironment) GetConfirmChanges() bool`

GetConfirmChanges returns the ConfirmChanges field if non-nil, zero value otherwise.

### GetConfirmChangesOk

`func (o *InternalEnvironment) GetConfirmChangesOk() (*bool, bool)`

GetConfirmChangesOk returns a tuple with the ConfirmChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmChanges

`func (o *InternalEnvironment) SetConfirmChanges(v bool)`

SetConfirmChanges sets ConfirmChanges field to given value.


### GetTags

`func (o *InternalEnvironment) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InternalEnvironment) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InternalEnvironment) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetApprovalSettings

`func (o *InternalEnvironment) GetApprovalSettings() ApprovalSettings`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *InternalEnvironment) GetApprovalSettingsOk() (*ApprovalSettings, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *InternalEnvironment) SetApprovalSettings(v ApprovalSettings)`

SetApprovalSettings sets ApprovalSettings field to given value.

### HasApprovalSettings

`func (o *InternalEnvironment) HasApprovalSettings() bool`

HasApprovalSettings returns a boolean if a field has been set.

### GetKind

`func (o *InternalEnvironment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InternalEnvironment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InternalEnvironment) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InternalEnvironment) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPubnub

`func (o *InternalEnvironment) GetPubnub() InternalPubnubRep`

GetPubnub returns the Pubnub field if non-nil, zero value otherwise.

### GetPubnubOk

`func (o *InternalEnvironment) GetPubnubOk() (*InternalPubnubRep, bool)`

GetPubnubOk returns a tuple with the Pubnub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubnub

`func (o *InternalEnvironment) SetPubnub(v InternalPubnubRep)`

SetPubnub sets Pubnub field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


