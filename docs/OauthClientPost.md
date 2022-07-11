# OauthClientPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of your new LaunchDarkly OAuth 2.0 client. | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI for your new OAuth 2.0 application. This should be an absolute URL conforming with the standard HTTPS protocol. | [optional] 
**Description** | Pointer to **string** | Description of your OAuth 2.0 client. | [optional] 

## Methods

### NewOauthClientPost

`func NewOauthClientPost() *OauthClientPost`

NewOauthClientPost instantiates a new OauthClientPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthClientPostWithDefaults

`func NewOauthClientPostWithDefaults() *OauthClientPost`

NewOauthClientPostWithDefaults instantiates a new OauthClientPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OauthClientPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OauthClientPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OauthClientPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OauthClientPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OauthClientPost) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OauthClientPost) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OauthClientPost) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OauthClientPost) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetDescription

`func (o *OauthClientPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OauthClientPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OauthClientPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OauthClientPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


