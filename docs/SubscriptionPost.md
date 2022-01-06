# SubscriptionPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for your audit log subscription. | 
**Statements** | Pointer to [**[]StatementPost**](StatementPost.md) |  | [optional] 
**On** | Pointer to **bool** | Whether or not you want your subscription to actively send events. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Config** | **map[string]interface{}** | The unique set of fields required to configure an audit log subscription integration of this type. Refer to the \&quot;formVariables\&quot; field in the corresponding manifest.json  at https://github.com/launchdarkly/integration-framework/tree/master/integrations for a full list of fields for the integration you wish to configure. | 
**Url** | Pointer to **string** | Slack webhook receiver URL. Only necessary for legacy Slack webhook integrations. | [optional] 
**ApiKey** | Pointer to **string** | Datadog API key. Only necessary for legacy Datadog webhook subscriptions. | [optional] 

## Methods

### NewSubscriptionPost

`func NewSubscriptionPost(name string, config map[string]interface{}, ) *SubscriptionPost`

NewSubscriptionPost instantiates a new SubscriptionPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPostWithDefaults

`func NewSubscriptionPostWithDefaults() *SubscriptionPost`

NewSubscriptionPostWithDefaults instantiates a new SubscriptionPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubscriptionPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionPost) SetName(v string)`

SetName sets Name field to given value.


### GetStatements

`func (o *SubscriptionPost) GetStatements() []StatementPost`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *SubscriptionPost) GetStatementsOk() (*[]StatementPost, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *SubscriptionPost) SetStatements(v []StatementPost)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *SubscriptionPost) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetOn

`func (o *SubscriptionPost) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *SubscriptionPost) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *SubscriptionPost) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *SubscriptionPost) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetTags

`func (o *SubscriptionPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SubscriptionPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SubscriptionPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SubscriptionPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfig

`func (o *SubscriptionPost) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SubscriptionPost) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SubscriptionPost) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetUrl

`func (o *SubscriptionPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubscriptionPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubscriptionPost) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SubscriptionPost) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *SubscriptionPost) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SubscriptionPost) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SubscriptionPost) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *SubscriptionPost) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


