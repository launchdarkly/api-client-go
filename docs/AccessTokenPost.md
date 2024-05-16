# AccessTokenPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-friendly name for the access token | [optional] 
**Description** | Pointer to **string** | A description for the access token | [optional] 
**Role** | Pointer to **string** | Built-in role for the token | [optional] 
**CustomRoleIds** | Pointer to **[]string** | A list of custom role IDs to use as access limits for the access token | [optional] 
**InlineRole** | Pointer to [**[]StatementPost**](StatementPost.md) | A JSON array of statements represented as JSON objects with three attributes: effect, resources, actions. May be used in place of a built-in or custom role. | [optional] 
**ServiceToken** | Pointer to **bool** | Whether the token is a service token https://docs.launchdarkly.com/home/account/api#service-tokens | [optional] 
**DefaultApiVersion** | Pointer to **int32** | The default API version for this token | [optional] 

## Methods

### NewAccessTokenPost

`func NewAccessTokenPost() *AccessTokenPost`

NewAccessTokenPost instantiates a new AccessTokenPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenPostWithDefaults

`func NewAccessTokenPostWithDefaults() *AccessTokenPost`

NewAccessTokenPostWithDefaults instantiates a new AccessTokenPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessTokenPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessTokenPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AccessTokenPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *AccessTokenPost) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccessTokenPost) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccessTokenPost) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AccessTokenPost) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCustomRoleIds

`func (o *AccessTokenPost) GetCustomRoleIds() []string`

GetCustomRoleIds returns the CustomRoleIds field if non-nil, zero value otherwise.

### GetCustomRoleIdsOk

`func (o *AccessTokenPost) GetCustomRoleIdsOk() (*[]string, bool)`

GetCustomRoleIdsOk returns a tuple with the CustomRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleIds

`func (o *AccessTokenPost) SetCustomRoleIds(v []string)`

SetCustomRoleIds sets CustomRoleIds field to given value.

### HasCustomRoleIds

`func (o *AccessTokenPost) HasCustomRoleIds() bool`

HasCustomRoleIds returns a boolean if a field has been set.

### GetInlineRole

`func (o *AccessTokenPost) GetInlineRole() []StatementPost`

GetInlineRole returns the InlineRole field if non-nil, zero value otherwise.

### GetInlineRoleOk

`func (o *AccessTokenPost) GetInlineRoleOk() (*[]StatementPost, bool)`

GetInlineRoleOk returns a tuple with the InlineRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineRole

`func (o *AccessTokenPost) SetInlineRole(v []StatementPost)`

SetInlineRole sets InlineRole field to given value.

### HasInlineRole

`func (o *AccessTokenPost) HasInlineRole() bool`

HasInlineRole returns a boolean if a field has been set.

### GetServiceToken

`func (o *AccessTokenPost) GetServiceToken() bool`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *AccessTokenPost) GetServiceTokenOk() (*bool, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *AccessTokenPost) SetServiceToken(v bool)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *AccessTokenPost) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetDefaultApiVersion

`func (o *AccessTokenPost) GetDefaultApiVersion() int32`

GetDefaultApiVersion returns the DefaultApiVersion field if non-nil, zero value otherwise.

### GetDefaultApiVersionOk

`func (o *AccessTokenPost) GetDefaultApiVersionOk() (*int32, bool)`

GetDefaultApiVersionOk returns a tuple with the DefaultApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiVersion

`func (o *AccessTokenPost) SetDefaultApiVersion(v int32)`

SetDefaultApiVersion sets DefaultApiVersion field to given value.

### HasDefaultApiVersion

`func (o *AccessTokenPost) HasDefaultApiVersion() bool`

HasDefaultApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


