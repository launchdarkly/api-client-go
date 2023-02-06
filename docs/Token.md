# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OwnerId** | **string** |  | 
**MemberId** | **string** |  | 
**Member** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**Name** | Pointer to **string** | A human-friendly name for the access token | [optional] 
**Description** | Pointer to **string** | A description for the access token | [optional] 
**CreationDate** | **int64** |  | 
**LastModified** | **int64** |  | 
**CustomRoleIds** | Pointer to **[]string** | A list of custom role IDs to use as access limits for the access token | [optional] 
**InlineRole** | Pointer to [**[]Statement**](Statement.md) | An array of policy statements, with three attributes: effect, resources, actions. May be used in place of a built-in or custom role. | [optional] 
**Role** | Pointer to **string** | Built-in role for the token | [optional] 
**Token** | Pointer to **string** | The token value. When creating or resetting, contains the entire token value. Otherwise, contains the last four characters. | [optional] 
**ServiceToken** | Pointer to **bool** | Whether this is a service token or a personal token | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**DefaultApiVersion** | Pointer to **int32** | The default API version for this token | [optional] 
**LastUsed** | Pointer to **int64** |  | [optional] 

## Methods

### NewToken

`func NewToken(id string, ownerId string, memberId string, creationDate int64, lastModified int64, links map[string]Link, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Token) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Token) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Token) SetId(v string)`

SetId sets Id field to given value.


### GetOwnerId

`func (o *Token) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Token) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Token) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetMemberId

`func (o *Token) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *Token) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *Token) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetMember

`func (o *Token) GetMember() MemberSummary`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Token) GetMemberOk() (*MemberSummary, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Token) SetMember(v MemberSummary)`

SetMember sets Member field to given value.

### HasMember

`func (o *Token) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *Token) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Token) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Token) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Token) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Token) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Token) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Token) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Token) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreationDate

`func (o *Token) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Token) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Token) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *Token) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Token) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Token) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetCustomRoleIds

`func (o *Token) GetCustomRoleIds() []string`

GetCustomRoleIds returns the CustomRoleIds field if non-nil, zero value otherwise.

### GetCustomRoleIdsOk

`func (o *Token) GetCustomRoleIdsOk() (*[]string, bool)`

GetCustomRoleIdsOk returns a tuple with the CustomRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleIds

`func (o *Token) SetCustomRoleIds(v []string)`

SetCustomRoleIds sets CustomRoleIds field to given value.

### HasCustomRoleIds

`func (o *Token) HasCustomRoleIds() bool`

HasCustomRoleIds returns a boolean if a field has been set.

### GetInlineRole

`func (o *Token) GetInlineRole() []Statement`

GetInlineRole returns the InlineRole field if non-nil, zero value otherwise.

### GetInlineRoleOk

`func (o *Token) GetInlineRoleOk() (*[]Statement, bool)`

GetInlineRoleOk returns a tuple with the InlineRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineRole

`func (o *Token) SetInlineRole(v []Statement)`

SetInlineRole sets InlineRole field to given value.

### HasInlineRole

`func (o *Token) HasInlineRole() bool`

HasInlineRole returns a boolean if a field has been set.

### GetRole

`func (o *Token) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Token) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Token) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Token) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToken

`func (o *Token) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Token) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Token) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Token) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetServiceToken

`func (o *Token) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *Token) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *Token) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *Token) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetLinks

`func (o *Token) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Token) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Token) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetDefaultApiVersion

`func (o *Token) GetDefaultApiVersion() int32`

GetDefaultApiVersion returns the DefaultApiVersion field if non-nil, zero value otherwise.

### GetDefaultApiVersionOk

`func (o *Token) GetDefaultApiVersionOk() (*int32, bool)`

GetDefaultApiVersionOk returns a tuple with the DefaultApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiVersion

`func (o *Token) SetDefaultApiVersion(v int32)`

SetDefaultApiVersion sets DefaultApiVersion field to given value.

### HasDefaultApiVersion

`func (o *Token) HasDefaultApiVersion() bool`

HasDefaultApiVersion returns a boolean if a field has been set.

### GetLastUsed

`func (o *Token) GetLastUsed() int64`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *Token) GetLastUsedOk() (*int64, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *Token) SetLastUsed(v int64)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *Token) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


