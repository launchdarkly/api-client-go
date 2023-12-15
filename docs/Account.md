# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to **[]string** |  | [optional] 
**RequireMfa** | Pointer to **bool** |  | [optional] 
**SamlConfig** | Pointer to [**SamlConfigRep**](SamlConfigRep.md) |  | [optional] 
**BillingContact** | Pointer to [**BillingContact**](BillingContact.md) |  | [optional] 
**SessionConfig** | Pointer to [**SessionCfgRep**](SessionCfgRep.md) |  | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**EnableAccountImpersonationUntil** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *Account) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Account) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Account) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Account) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrganization

`func (o *Account) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Account) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Account) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Account) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTokens

`func (o *Account) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Account) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Account) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *Account) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetRequireMfa

`func (o *Account) GetRequireMfa() bool`

GetRequireMfa returns the RequireMfa field if non-nil, zero value otherwise.

### GetRequireMfaOk

`func (o *Account) GetRequireMfaOk() (*bool, bool)`

GetRequireMfaOk returns a tuple with the RequireMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfa

`func (o *Account) SetRequireMfa(v bool)`

SetRequireMfa sets RequireMfa field to given value.

### HasRequireMfa

`func (o *Account) HasRequireMfa() bool`

HasRequireMfa returns a boolean if a field has been set.

### GetSamlConfig

`func (o *Account) GetSamlConfig() SamlConfigRep`

GetSamlConfig returns the SamlConfig field if non-nil, zero value otherwise.

### GetSamlConfigOk

`func (o *Account) GetSamlConfigOk() (*SamlConfigRep, bool)`

GetSamlConfigOk returns a tuple with the SamlConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConfig

`func (o *Account) SetSamlConfig(v SamlConfigRep)`

SetSamlConfig sets SamlConfig field to given value.

### HasSamlConfig

`func (o *Account) HasSamlConfig() bool`

HasSamlConfig returns a boolean if a field has been set.

### GetBillingContact

`func (o *Account) GetBillingContact() BillingContact`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *Account) GetBillingContactOk() (*BillingContact, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *Account) SetBillingContact(v BillingContact)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *Account) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetSessionConfig

`func (o *Account) GetSessionConfig() SessionCfgRep`

GetSessionConfig returns the SessionConfig field if non-nil, zero value otherwise.

### GetSessionConfigOk

`func (o *Account) GetSessionConfigOk() (*SessionCfgRep, bool)`

GetSessionConfigOk returns a tuple with the SessionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionConfig

`func (o *Account) SetSessionConfig(v SessionCfgRep)`

SetSessionConfig sets SessionConfig field to given value.

### HasSessionConfig

`func (o *Account) HasSessionConfig() bool`

HasSessionConfig returns a boolean if a field has been set.

### GetAccess

`func (o *Account) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Account) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Account) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Account) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Account) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Account) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Account) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *Account) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetEnableAccountImpersonationUntil

`func (o *Account) GetEnableAccountImpersonationUntil() int64`

GetEnableAccountImpersonationUntil returns the EnableAccountImpersonationUntil field if non-nil, zero value otherwise.

### GetEnableAccountImpersonationUntilOk

`func (o *Account) GetEnableAccountImpersonationUntilOk() (*int64, bool)`

GetEnableAccountImpersonationUntilOk returns a tuple with the EnableAccountImpersonationUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountImpersonationUntil

`func (o *Account) SetEnableAccountImpersonationUntil(v int64)`

SetEnableAccountImpersonationUntil sets EnableAccountImpersonationUntil field to given value.

### HasEnableAccountImpersonationUntil

`func (o *Account) HasEnableAccountImpersonationUntil() bool`

HasEnableAccountImpersonationUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


