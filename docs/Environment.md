# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | Links to other resources within the API. Includes the URL and content type of those resources | 
**Id** | **string** |  | 
**Key** | **string** | A project-unique key for the new environment. | 
**Name** | **string** | A human-friendly name for the new environment. | 
**ApiKey** | **string** | API key to use with client-side SDKs. | 
**MobileKey** | **string** | API key to use with mobile SDKs. | 
**Color** | **string** | The color used to indicate this environment in the UI. | 
**DefaultTtl** | **int32** | The default time (in minutes) that the PHP SDK can cache feature flag rules locally. | 
**SecureMode** | **bool** | Ensures that one end user of the client-side SDK cannot inspect the variations for another end user. | 
**DefaultTrackEvents** | **bool** | Enables tracking detailed information for new flags by default. | 
**RequireComments** | **bool** | Whether members who modify flags and segments through the LaunchDarkly user interface are required to add a comment | 
**ConfirmChanges** | **bool** | Whether members who modify flags and segments through the LaunchDarkly user interface are required to confirm those changes | 
**Tags** | **[]string** | A list of tags for this environment | 
**ApprovalSettings** | Pointer to [**ApprovalSettings**](ApprovalSettings.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(links map[string]Link, id string, key string, name string, apiKey string, mobileKey string, color string, defaultTtl int32, secureMode bool, defaultTrackEvents bool, requireComments bool, confirmChanges bool, tags []string, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Environment) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Environment) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Environment) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *Environment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *Environment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Environment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Environment) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.


### GetApiKey

`func (o *Environment) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Environment) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Environment) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetMobileKey

`func (o *Environment) GetMobileKey() string`

GetMobileKey returns the MobileKey field if non-nil, zero value otherwise.

### GetMobileKeyOk

`func (o *Environment) GetMobileKeyOk() (*string, bool)`

GetMobileKeyOk returns a tuple with the MobileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileKey

`func (o *Environment) SetMobileKey(v string)`

SetMobileKey sets MobileKey field to given value.


### GetColor

`func (o *Environment) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Environment) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Environment) SetColor(v string)`

SetColor sets Color field to given value.


### GetDefaultTtl

`func (o *Environment) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *Environment) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *Environment) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.


### GetSecureMode

`func (o *Environment) GetSecureMode() bool`

GetSecureMode returns the SecureMode field if non-nil, zero value otherwise.

### GetSecureModeOk

`func (o *Environment) GetSecureModeOk() (*bool, bool)`

GetSecureModeOk returns a tuple with the SecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureMode

`func (o *Environment) SetSecureMode(v bool)`

SetSecureMode sets SecureMode field to given value.


### GetDefaultTrackEvents

`func (o *Environment) GetDefaultTrackEvents() bool`

GetDefaultTrackEvents returns the DefaultTrackEvents field if non-nil, zero value otherwise.

### GetDefaultTrackEventsOk

`func (o *Environment) GetDefaultTrackEventsOk() (*bool, bool)`

GetDefaultTrackEventsOk returns a tuple with the DefaultTrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTrackEvents

`func (o *Environment) SetDefaultTrackEvents(v bool)`

SetDefaultTrackEvents sets DefaultTrackEvents field to given value.


### GetRequireComments

`func (o *Environment) GetRequireComments() bool`

GetRequireComments returns the RequireComments field if non-nil, zero value otherwise.

### GetRequireCommentsOk

`func (o *Environment) GetRequireCommentsOk() (*bool, bool)`

GetRequireCommentsOk returns a tuple with the RequireComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireComments

`func (o *Environment) SetRequireComments(v bool)`

SetRequireComments sets RequireComments field to given value.


### GetConfirmChanges

`func (o *Environment) GetConfirmChanges() bool`

GetConfirmChanges returns the ConfirmChanges field if non-nil, zero value otherwise.

### GetConfirmChangesOk

`func (o *Environment) GetConfirmChangesOk() (*bool, bool)`

GetConfirmChangesOk returns a tuple with the ConfirmChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmChanges

`func (o *Environment) SetConfirmChanges(v bool)`

SetConfirmChanges sets ConfirmChanges field to given value.


### GetTags

`func (o *Environment) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Environment) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Environment) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetApprovalSettings

`func (o *Environment) GetApprovalSettings() ApprovalSettings`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *Environment) GetApprovalSettingsOk() (*ApprovalSettings, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *Environment) SetApprovalSettings(v ApprovalSettings)`

SetApprovalSettings sets ApprovalSettings field to given value.

### HasApprovalSettings

`func (o *Environment) HasApprovalSettings() bool`

HasApprovalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


