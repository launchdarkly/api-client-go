# EnvironmentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]CoreLink**](CoreLink.md) | Links to related resources. | 
**Id** | **string** |  | 
**Key** | **string** | A project-unique key for the new environment. | 
**Name** | **string** | A human-friendly name for the new environment. | 
**ApiKey** | **string** | API key to use with client-side SDKs. | 
**MobileKey** | **string** | API key to use with mobile SDKs. | 
**Color** | **string** | The color used to indicate this environment in the UI. | 
**DefaultTtl** | **int32** | The default time (in minutes) that the PHP SDK can cache feature flag rules locally. | 
**SecureMode** | **bool** | Secure mode ensures that a user of the client-side SDK cannot impersonate another user. | 
**DefaultTrackEvents** | **bool** | Enables tracking detailed information for new flags by default. | 
**Tags** | **[]string** |  | 
**ApprovalSettings** | Pointer to [**ApprovalSettingsRep**](ApprovalSettingsRep.md) |  | [optional] 

## Methods

### NewEnvironmentRep

`func NewEnvironmentRep(links map[string]CoreLink, id string, key string, name string, apiKey string, mobileKey string, color string, defaultTtl int32, secureMode bool, defaultTrackEvents bool, tags []string, ) *EnvironmentRep`

NewEnvironmentRep instantiates a new EnvironmentRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentRepWithDefaults

`func NewEnvironmentRepWithDefaults() *EnvironmentRep`

NewEnvironmentRepWithDefaults instantiates a new EnvironmentRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *EnvironmentRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentRep) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *EnvironmentRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *EnvironmentRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentRep) SetName(v string)`

SetName sets Name field to given value.


### GetApiKey

`func (o *EnvironmentRep) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *EnvironmentRep) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *EnvironmentRep) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetMobileKey

`func (o *EnvironmentRep) GetMobileKey() string`

GetMobileKey returns the MobileKey field if non-nil, zero value otherwise.

### GetMobileKeyOk

`func (o *EnvironmentRep) GetMobileKeyOk() (*string, bool)`

GetMobileKeyOk returns a tuple with the MobileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileKey

`func (o *EnvironmentRep) SetMobileKey(v string)`

SetMobileKey sets MobileKey field to given value.


### GetColor

`func (o *EnvironmentRep) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EnvironmentRep) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EnvironmentRep) SetColor(v string)`

SetColor sets Color field to given value.


### GetDefaultTtl

`func (o *EnvironmentRep) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *EnvironmentRep) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *EnvironmentRep) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.


### GetSecureMode

`func (o *EnvironmentRep) GetSecureMode() bool`

GetSecureMode returns the SecureMode field if non-nil, zero value otherwise.

### GetSecureModeOk

`func (o *EnvironmentRep) GetSecureModeOk() (*bool, bool)`

GetSecureModeOk returns a tuple with the SecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureMode

`func (o *EnvironmentRep) SetSecureMode(v bool)`

SetSecureMode sets SecureMode field to given value.


### GetDefaultTrackEvents

`func (o *EnvironmentRep) GetDefaultTrackEvents() bool`

GetDefaultTrackEvents returns the DefaultTrackEvents field if non-nil, zero value otherwise.

### GetDefaultTrackEventsOk

`func (o *EnvironmentRep) GetDefaultTrackEventsOk() (*bool, bool)`

GetDefaultTrackEventsOk returns a tuple with the DefaultTrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTrackEvents

`func (o *EnvironmentRep) SetDefaultTrackEvents(v bool)`

SetDefaultTrackEvents sets DefaultTrackEvents field to given value.


### GetTags

`func (o *EnvironmentRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EnvironmentRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EnvironmentRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetApprovalSettings

`func (o *EnvironmentRep) GetApprovalSettings() ApprovalSettingsRep`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *EnvironmentRep) GetApprovalSettingsOk() (*ApprovalSettingsRep, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *EnvironmentRep) SetApprovalSettings(v ApprovalSettingsRep)`

SetApprovalSettings sets ApprovalSettings field to given value.

### HasApprovalSettings

`func (o *EnvironmentRep) HasApprovalSettings() bool`

HasApprovalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


