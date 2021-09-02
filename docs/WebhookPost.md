# WebhookPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name for your webhook | [optional] 
**Url** | **string** | The URL of the remote webhook | 
**Secret** | Pointer to **string** | If sign is true, and the secret attribute is omitted, LaunchDarkly automatically generates a secret for you. | [optional] 
**Statements** | Pointer to [**[]StatementPost**](StatementPost.md) |  | [optional] 
**Sign** | **bool** | If sign is false, the webhook does not include a signature header, and the secret can be omitted. | 
**On** | **bool** | Whether or not this webhook is enabled. | 
**Tags** | Pointer to **[]string** | List of tags for this webhook | [optional] 

## Methods

### NewWebhookPost

`func NewWebhookPost(url string, sign bool, on bool, ) *WebhookPost`

NewWebhookPost instantiates a new WebhookPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPostWithDefaults

`func NewWebhookPostWithDefaults() *WebhookPost`

NewWebhookPostWithDefaults instantiates a new WebhookPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookPost) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSecret

`func (o *WebhookPost) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookPost) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookPost) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookPost) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatements

`func (o *WebhookPost) GetStatements() []StatementPost`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *WebhookPost) GetStatementsOk() (*[]StatementPost, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *WebhookPost) SetStatements(v []StatementPost)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *WebhookPost) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetSign

`func (o *WebhookPost) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *WebhookPost) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *WebhookPost) SetSign(v bool)`

SetSign sets Sign field to given value.


### GetOn

`func (o *WebhookPost) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *WebhookPost) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *WebhookPost) SetOn(v bool)`

SetOn sets On field to given value.


### GetTags

`func (o *WebhookPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebhookPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebhookPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WebhookPost) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


