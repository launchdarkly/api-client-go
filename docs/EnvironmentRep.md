# EnvironmentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**MobileKey** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**DefaultTtl** | Pointer to **int32** |  | [optional] 
**SecureMode** | Pointer to **bool** |  | [optional] 
**DefaultTrackEvents** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ApprovalSettings** | Pointer to [**ApprovalSettingsRep**](ApprovalSettingsRep.md) |  | [optional] 

## Methods

### NewEnvironmentRep

`func NewEnvironmentRep() *EnvironmentRep`

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

`func (o *EnvironmentRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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

### HasId

`func (o *EnvironmentRep) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasKey

`func (o *EnvironmentRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

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

### HasName

`func (o *EnvironmentRep) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasApiKey

`func (o *EnvironmentRep) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

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

### HasMobileKey

`func (o *EnvironmentRep) HasMobileKey() bool`

HasMobileKey returns a boolean if a field has been set.

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

### HasColor

`func (o *EnvironmentRep) HasColor() bool`

HasColor returns a boolean if a field has been set.

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

### HasDefaultTtl

`func (o *EnvironmentRep) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

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

### HasSecureMode

`func (o *EnvironmentRep) HasSecureMode() bool`

HasSecureMode returns a boolean if a field has been set.

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

### HasDefaultTrackEvents

`func (o *EnvironmentRep) HasDefaultTrackEvents() bool`

HasDefaultTrackEvents returns a boolean if a field has been set.

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

### HasTags

`func (o *EnvironmentRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

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


