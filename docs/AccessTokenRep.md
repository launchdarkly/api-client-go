# AccessTokenRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**MemberId** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberSummaryRep**](MemberSummaryRep.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**LastModified** | Pointer to **int64** |  | [optional] 
**CustomRoleIds** | Pointer to **[]string** |  | [optional] 
**InlineRole** | Pointer to [**[]StatementRep**](StatementRep.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
**DefaultApiVersion** | Pointer to **int32** |  | [optional] 
**LastUsed** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccessTokenRep

`func NewAccessTokenRep() *AccessTokenRep`

NewAccessTokenRep instantiates a new AccessTokenRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenRepWithDefaults

`func NewAccessTokenRepWithDefaults() *AccessTokenRep`

NewAccessTokenRepWithDefaults instantiates a new AccessTokenRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessTokenRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessTokenRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerId

`func (o *AccessTokenRep) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AccessTokenRep) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AccessTokenRep) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AccessTokenRep) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetMemberId

`func (o *AccessTokenRep) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *AccessTokenRep) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *AccessTokenRep) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *AccessTokenRep) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMember

`func (o *AccessTokenRep) GetMember() MemberSummaryRep`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AccessTokenRep) GetMemberOk() (*MemberSummaryRep, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AccessTokenRep) SetMember(v MemberSummaryRep)`

SetMember sets Member field to given value.

### HasMember

`func (o *AccessTokenRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *AccessTokenRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessTokenRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AccessTokenRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreationDate

`func (o *AccessTokenRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AccessTokenRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AccessTokenRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AccessTokenRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModified

`func (o *AccessTokenRep) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AccessTokenRep) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AccessTokenRep) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AccessTokenRep) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetCustomRoleIds

`func (o *AccessTokenRep) GetCustomRoleIds() []string`

GetCustomRoleIds returns the CustomRoleIds field if non-nil, zero value otherwise.

### GetCustomRoleIdsOk

`func (o *AccessTokenRep) GetCustomRoleIdsOk() (*[]string, bool)`

GetCustomRoleIdsOk returns a tuple with the CustomRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleIds

`func (o *AccessTokenRep) SetCustomRoleIds(v []string)`

SetCustomRoleIds sets CustomRoleIds field to given value.

### HasCustomRoleIds

`func (o *AccessTokenRep) HasCustomRoleIds() bool`

HasCustomRoleIds returns a boolean if a field has been set.

### GetInlineRole

`func (o *AccessTokenRep) GetInlineRole() []StatementRep`

GetInlineRole returns the InlineRole field if non-nil, zero value otherwise.

### GetInlineRoleOk

`func (o *AccessTokenRep) GetInlineRoleOk() (*[]StatementRep, bool)`

GetInlineRoleOk returns a tuple with the InlineRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineRole

`func (o *AccessTokenRep) SetInlineRole(v []StatementRep)`

SetInlineRole sets InlineRole field to given value.

### HasInlineRole

`func (o *AccessTokenRep) HasInlineRole() bool`

HasInlineRole returns a boolean if a field has been set.

### GetRole

`func (o *AccessTokenRep) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccessTokenRep) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccessTokenRep) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AccessTokenRep) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToken

`func (o *AccessTokenRep) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessTokenRep) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessTokenRep) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccessTokenRep) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetServiceToken

`func (o *AccessTokenRep) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *AccessTokenRep) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *AccessTokenRep) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *AccessTokenRep) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetLinks

`func (o *AccessTokenRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccessTokenRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccessTokenRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccessTokenRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDefaultApiVersion

`func (o *AccessTokenRep) GetDefaultApiVersion() int32`

GetDefaultApiVersion returns the DefaultApiVersion field if non-nil, zero value otherwise.

### GetDefaultApiVersionOk

`func (o *AccessTokenRep) GetDefaultApiVersionOk() (*int32, bool)`

GetDefaultApiVersionOk returns a tuple with the DefaultApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiVersion

`func (o *AccessTokenRep) SetDefaultApiVersion(v int32)`

SetDefaultApiVersion sets DefaultApiVersion field to given value.

### HasDefaultApiVersion

`func (o *AccessTokenRep) HasDefaultApiVersion() bool`

HasDefaultApiVersion returns a boolean if a field has been set.

### GetLastUsed

`func (o *AccessTokenRep) GetLastUsed() int64`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *AccessTokenRep) GetLastUsedOk() (*int64, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *AccessTokenRep) SetLastUsed(v int64)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *AccessTokenRep) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


