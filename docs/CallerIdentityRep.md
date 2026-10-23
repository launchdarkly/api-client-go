# CallerIdentityRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**EnvironmentName** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**AuthKind** | Pointer to **string** |  | [optional] 
**TokenKind** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientName** | Pointer to **string** | The registered name of the OAuth client the access token was issued to. Only present for OAuth access tokens, and omitted when the client&#39;s name can&#39;t be resolved. | [optional] 
**Audience** | Pointer to **[]string** | The resources (RFC 8707) the OAuth access token was issued for. Only present for OAuth access tokens issued for a specific resource. | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**TokenName** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**MemberId** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **bool** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCallerIdentityRep

`func NewCallerIdentityRep() *CallerIdentityRep`

NewCallerIdentityRep instantiates a new CallerIdentityRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallerIdentityRepWithDefaults

`func NewCallerIdentityRepWithDefaults() *CallerIdentityRep`

NewCallerIdentityRepWithDefaults instantiates a new CallerIdentityRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CallerIdentityRep) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CallerIdentityRep) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CallerIdentityRep) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CallerIdentityRep) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *CallerIdentityRep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CallerIdentityRep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CallerIdentityRep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CallerIdentityRep) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetProjectId

`func (o *CallerIdentityRep) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CallerIdentityRep) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CallerIdentityRep) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CallerIdentityRep) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *CallerIdentityRep) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *CallerIdentityRep) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *CallerIdentityRep) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *CallerIdentityRep) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetProjectName

`func (o *CallerIdentityRep) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CallerIdentityRep) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CallerIdentityRep) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *CallerIdentityRep) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetAuthKind

`func (o *CallerIdentityRep) GetAuthKind() string`

GetAuthKind returns the AuthKind field if non-nil, zero value otherwise.

### GetAuthKindOk

`func (o *CallerIdentityRep) GetAuthKindOk() (*string, bool)`

GetAuthKindOk returns a tuple with the AuthKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKind

`func (o *CallerIdentityRep) SetAuthKind(v string)`

SetAuthKind sets AuthKind field to given value.

### HasAuthKind

`func (o *CallerIdentityRep) HasAuthKind() bool`

HasAuthKind returns a boolean if a field has been set.

### GetTokenKind

`func (o *CallerIdentityRep) GetTokenKind() string`

GetTokenKind returns the TokenKind field if non-nil, zero value otherwise.

### GetTokenKindOk

`func (o *CallerIdentityRep) GetTokenKindOk() (*string, bool)`

GetTokenKindOk returns a tuple with the TokenKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenKind

`func (o *CallerIdentityRep) SetTokenKind(v string)`

SetTokenKind sets TokenKind field to given value.

### HasTokenKind

`func (o *CallerIdentityRep) HasTokenKind() bool`

HasTokenKind returns a boolean if a field has been set.

### GetClientId

`func (o *CallerIdentityRep) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CallerIdentityRep) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CallerIdentityRep) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CallerIdentityRep) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *CallerIdentityRep) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CallerIdentityRep) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CallerIdentityRep) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *CallerIdentityRep) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetAudience

`func (o *CallerIdentityRep) GetAudience() []string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CallerIdentityRep) GetAudienceOk() (*[]string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CallerIdentityRep) SetAudience(v []string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CallerIdentityRep) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CallerIdentityRep) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CallerIdentityRep) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CallerIdentityRep) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CallerIdentityRep) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetTokenName

`func (o *CallerIdentityRep) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *CallerIdentityRep) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *CallerIdentityRep) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *CallerIdentityRep) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTokenId

`func (o *CallerIdentityRep) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CallerIdentityRep) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CallerIdentityRep) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CallerIdentityRep) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetMemberId

`func (o *CallerIdentityRep) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *CallerIdentityRep) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *CallerIdentityRep) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *CallerIdentityRep) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetServiceToken

`func (o *CallerIdentityRep) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *CallerIdentityRep) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *CallerIdentityRep) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *CallerIdentityRep) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetScopes

`func (o *CallerIdentityRep) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CallerIdentityRep) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CallerIdentityRep) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CallerIdentityRep) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


