# AccessTokensAccessTokenPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**CustomRoleIds** | Pointer to **[]string** |  | [optional] 
**InlineRole** | Pointer to [**[]AccessTokensAccessTokenPostInlineRole**](AccessTokensAccessTokenPostInlineRole.md) |  | [optional] 
**ServiceToken** | Pointer to **bool** |  | [optional] 
**DefaultApiVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccessTokensAccessTokenPost

`func NewAccessTokensAccessTokenPost() *AccessTokensAccessTokenPost`

NewAccessTokensAccessTokenPost instantiates a new AccessTokensAccessTokenPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokensAccessTokenPostWithDefaults

`func NewAccessTokensAccessTokenPostWithDefaults() *AccessTokensAccessTokenPost`

NewAccessTokensAccessTokenPostWithDefaults instantiates a new AccessTokensAccessTokenPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessTokensAccessTokenPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokensAccessTokenPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokensAccessTokenPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessTokensAccessTokenPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AccessTokensAccessTokenPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokensAccessTokenPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokensAccessTokenPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokensAccessTokenPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *AccessTokensAccessTokenPost) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccessTokensAccessTokenPost) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccessTokensAccessTokenPost) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AccessTokensAccessTokenPost) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCustomRoleIds

`func (o *AccessTokensAccessTokenPost) GetCustomRoleIds() []string`

GetCustomRoleIds returns the CustomRoleIds field if non-nil, zero value otherwise.

### GetCustomRoleIdsOk

`func (o *AccessTokensAccessTokenPost) GetCustomRoleIdsOk() (*[]string, bool)`

GetCustomRoleIdsOk returns a tuple with the CustomRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleIds

`func (o *AccessTokensAccessTokenPost) SetCustomRoleIds(v []string)`

SetCustomRoleIds sets CustomRoleIds field to given value.

### HasCustomRoleIds

`func (o *AccessTokensAccessTokenPost) HasCustomRoleIds() bool`

HasCustomRoleIds returns a boolean if a field has been set.

### GetInlineRole

`func (o *AccessTokensAccessTokenPost) GetInlineRole() []AccessTokensAccessTokenPostInlineRole`

GetInlineRole returns the InlineRole field if non-nil, zero value otherwise.

### GetInlineRoleOk

`func (o *AccessTokensAccessTokenPost) GetInlineRoleOk() (*[]AccessTokensAccessTokenPostInlineRole, bool)`

GetInlineRoleOk returns a tuple with the InlineRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineRole

`func (o *AccessTokensAccessTokenPost) SetInlineRole(v []AccessTokensAccessTokenPostInlineRole)`

SetInlineRole sets InlineRole field to given value.

### HasInlineRole

`func (o *AccessTokensAccessTokenPost) HasInlineRole() bool`

HasInlineRole returns a boolean if a field has been set.

### GetServiceToken

`func (o *AccessTokensAccessTokenPost) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *AccessTokensAccessTokenPost) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *AccessTokensAccessTokenPost) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *AccessTokensAccessTokenPost) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetDefaultApiVersion

`func (o *AccessTokensAccessTokenPost) GetDefaultApiVersion() int32`

GetDefaultApiVersion returns the DefaultApiVersion field if non-nil, zero value otherwise.

### GetDefaultApiVersionOk

`func (o *AccessTokensAccessTokenPost) GetDefaultApiVersionOk() (*int32, bool)`

GetDefaultApiVersionOk returns a tuple with the DefaultApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiVersion

`func (o *AccessTokensAccessTokenPost) SetDefaultApiVersion(v int32)`

SetDefaultApiVersion sets DefaultApiVersion field to given value.

### HasDefaultApiVersion

`func (o *AccessTokensAccessTokenPost) HasDefaultApiVersion() bool`

HasDefaultApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


