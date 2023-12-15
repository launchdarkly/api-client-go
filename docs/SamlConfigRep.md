# SamlConfigRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestRedirectUrl** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RequireSso** | Pointer to **bool** |  | [optional] 
**SsoUrl** | Pointer to **string** |  | [optional] 
**X509Certificate** | Pointer to **string** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**DefaultRole** | Pointer to **string** |  | [optional] 
**IsManagingTeams** | Pointer to **bool** |  | [optional] 
**SyncedTeamCount** | Pointer to **int32** |  | [optional] 
**UnSyncedTeamCount** | Pointer to **int32** |  | [optional] 
**SsoAcsUrl** | Pointer to **string** |  | [optional] 
**SsoEntityId** | Pointer to **string** |  | [optional] 
**SsoStartUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewSamlConfigRep

`func NewSamlConfigRep() *SamlConfigRep`

NewSamlConfigRep instantiates a new SamlConfigRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlConfigRepWithDefaults

`func NewSamlConfigRepWithDefaults() *SamlConfigRep`

NewSamlConfigRepWithDefaults instantiates a new SamlConfigRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRedirectUrl

`func (o *SamlConfigRep) GetTestRedirectUrl() string`

GetTestRedirectUrl returns the TestRedirectUrl field if non-nil, zero value otherwise.

### GetTestRedirectUrlOk

`func (o *SamlConfigRep) GetTestRedirectUrlOk() (*string, bool)`

GetTestRedirectUrlOk returns a tuple with the TestRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRedirectUrl

`func (o *SamlConfigRep) SetTestRedirectUrl(v string)`

SetTestRedirectUrl sets TestRedirectUrl field to given value.

### HasTestRedirectUrl

`func (o *SamlConfigRep) HasTestRedirectUrl() bool`

HasTestRedirectUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *SamlConfigRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlConfigRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlConfigRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SamlConfigRep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequireSso

`func (o *SamlConfigRep) GetRequireSso() bool`

GetRequireSso returns the RequireSso field if non-nil, zero value otherwise.

### GetRequireSsoOk

`func (o *SamlConfigRep) GetRequireSsoOk() (*bool, bool)`

GetRequireSsoOk returns a tuple with the RequireSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSso

`func (o *SamlConfigRep) SetRequireSso(v bool)`

SetRequireSso sets RequireSso field to given value.

### HasRequireSso

`func (o *SamlConfigRep) HasRequireSso() bool`

HasRequireSso returns a boolean if a field has been set.

### GetSsoUrl

`func (o *SamlConfigRep) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *SamlConfigRep) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *SamlConfigRep) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *SamlConfigRep) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetX509Certificate

`func (o *SamlConfigRep) GetX509Certificate() string`

GetX509Certificate returns the X509Certificate field if non-nil, zero value otherwise.

### GetX509CertificateOk

`func (o *SamlConfigRep) GetX509CertificateOk() (*string, bool)`

GetX509CertificateOk returns a tuple with the X509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Certificate

`func (o *SamlConfigRep) SetX509Certificate(v string)`

SetX509Certificate sets X509Certificate field to given value.

### HasX509Certificate

`func (o *SamlConfigRep) HasX509Certificate() bool`

HasX509Certificate returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SamlConfigRep) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SamlConfigRep) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SamlConfigRep) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SamlConfigRep) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetDefaultRole

`func (o *SamlConfigRep) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *SamlConfigRep) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *SamlConfigRep) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *SamlConfigRep) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetIsManagingTeams

`func (o *SamlConfigRep) GetIsManagingTeams() bool`

GetIsManagingTeams returns the IsManagingTeams field if non-nil, zero value otherwise.

### GetIsManagingTeamsOk

`func (o *SamlConfigRep) GetIsManagingTeamsOk() (*bool, bool)`

GetIsManagingTeamsOk returns a tuple with the IsManagingTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagingTeams

`func (o *SamlConfigRep) SetIsManagingTeams(v bool)`

SetIsManagingTeams sets IsManagingTeams field to given value.

### HasIsManagingTeams

`func (o *SamlConfigRep) HasIsManagingTeams() bool`

HasIsManagingTeams returns a boolean if a field has been set.

### GetSyncedTeamCount

`func (o *SamlConfigRep) GetSyncedTeamCount() int32`

GetSyncedTeamCount returns the SyncedTeamCount field if non-nil, zero value otherwise.

### GetSyncedTeamCountOk

`func (o *SamlConfigRep) GetSyncedTeamCountOk() (*int32, bool)`

GetSyncedTeamCountOk returns a tuple with the SyncedTeamCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedTeamCount

`func (o *SamlConfigRep) SetSyncedTeamCount(v int32)`

SetSyncedTeamCount sets SyncedTeamCount field to given value.

### HasSyncedTeamCount

`func (o *SamlConfigRep) HasSyncedTeamCount() bool`

HasSyncedTeamCount returns a boolean if a field has been set.

### GetUnSyncedTeamCount

`func (o *SamlConfigRep) GetUnSyncedTeamCount() int32`

GetUnSyncedTeamCount returns the UnSyncedTeamCount field if non-nil, zero value otherwise.

### GetUnSyncedTeamCountOk

`func (o *SamlConfigRep) GetUnSyncedTeamCountOk() (*int32, bool)`

GetUnSyncedTeamCountOk returns a tuple with the UnSyncedTeamCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSyncedTeamCount

`func (o *SamlConfigRep) SetUnSyncedTeamCount(v int32)`

SetUnSyncedTeamCount sets UnSyncedTeamCount field to given value.

### HasUnSyncedTeamCount

`func (o *SamlConfigRep) HasUnSyncedTeamCount() bool`

HasUnSyncedTeamCount returns a boolean if a field has been set.

### GetSsoAcsUrl

`func (o *SamlConfigRep) GetSsoAcsUrl() string`

GetSsoAcsUrl returns the SsoAcsUrl field if non-nil, zero value otherwise.

### GetSsoAcsUrlOk

`func (o *SamlConfigRep) GetSsoAcsUrlOk() (*string, bool)`

GetSsoAcsUrlOk returns a tuple with the SsoAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAcsUrl

`func (o *SamlConfigRep) SetSsoAcsUrl(v string)`

SetSsoAcsUrl sets SsoAcsUrl field to given value.

### HasSsoAcsUrl

`func (o *SamlConfigRep) HasSsoAcsUrl() bool`

HasSsoAcsUrl returns a boolean if a field has been set.

### GetSsoEntityId

`func (o *SamlConfigRep) GetSsoEntityId() string`

GetSsoEntityId returns the SsoEntityId field if non-nil, zero value otherwise.

### GetSsoEntityIdOk

`func (o *SamlConfigRep) GetSsoEntityIdOk() (*string, bool)`

GetSsoEntityIdOk returns a tuple with the SsoEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEntityId

`func (o *SamlConfigRep) SetSsoEntityId(v string)`

SetSsoEntityId sets SsoEntityId field to given value.

### HasSsoEntityId

`func (o *SamlConfigRep) HasSsoEntityId() bool`

HasSsoEntityId returns a boolean if a field has been set.

### GetSsoStartUrl

`func (o *SamlConfigRep) GetSsoStartUrl() string`

GetSsoStartUrl returns the SsoStartUrl field if non-nil, zero value otherwise.

### GetSsoStartUrlOk

`func (o *SamlConfigRep) GetSsoStartUrlOk() (*string, bool)`

GetSsoStartUrlOk returns a tuple with the SsoStartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoStartUrl

`func (o *SamlConfigRep) SetSsoStartUrl(v string)`

SetSsoStartUrl sets SsoStartUrl field to given value.

### HasSsoStartUrl

`func (o *SamlConfigRep) HasSsoStartUrl() bool`

HasSsoStartUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


