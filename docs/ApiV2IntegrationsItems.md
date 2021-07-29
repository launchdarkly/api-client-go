# ApiV2IntegrationsItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Statements** | Pointer to [**[]StatementRep**](StatementRep.md) |  | [optional] 
**On** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 
**Status** | Pointer to [**IntegrationSubscriptionStatusRep**](IntegrationSubscriptionStatusRep.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV2IntegrationsItems

`func NewApiV2IntegrationsItems() *ApiV2IntegrationsItems`

NewApiV2IntegrationsItems instantiates a new ApiV2IntegrationsItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2IntegrationsItemsWithDefaults

`func NewApiV2IntegrationsItemsWithDefaults() *ApiV2IntegrationsItems`

NewApiV2IntegrationsItemsWithDefaults instantiates a new ApiV2IntegrationsItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApiV2IntegrationsItems) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiV2IntegrationsItems) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiV2IntegrationsItems) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiV2IntegrationsItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ApiV2IntegrationsItems) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV2IntegrationsItems) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV2IntegrationsItems) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiV2IntegrationsItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ApiV2IntegrationsItems) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV2IntegrationsItems) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV2IntegrationsItems) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiV2IntegrationsItems) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ApiV2IntegrationsItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2IntegrationsItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2IntegrationsItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2IntegrationsItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *ApiV2IntegrationsItems) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiV2IntegrationsItems) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiV2IntegrationsItems) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiV2IntegrationsItems) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatements

`func (o *ApiV2IntegrationsItems) GetStatements() []StatementRep`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *ApiV2IntegrationsItems) GetStatementsOk() (*[]StatementRep, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *ApiV2IntegrationsItems) SetStatements(v []StatementRep)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *ApiV2IntegrationsItems) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetOn

`func (o *ApiV2IntegrationsItems) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *ApiV2IntegrationsItems) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *ApiV2IntegrationsItems) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *ApiV2IntegrationsItems) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetTags

`func (o *ApiV2IntegrationsItems) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiV2IntegrationsItems) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiV2IntegrationsItems) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiV2IntegrationsItems) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccess

`func (o *ApiV2IntegrationsItems) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApiV2IntegrationsItems) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApiV2IntegrationsItems) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ApiV2IntegrationsItems) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetStatus

`func (o *ApiV2IntegrationsItems) GetStatus() IntegrationSubscriptionStatusRep`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2IntegrationsItems) GetStatusOk() (*IntegrationSubscriptionStatusRep, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2IntegrationsItems) SetStatus(v IntegrationSubscriptionStatusRep)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2IntegrationsItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *ApiV2IntegrationsItems) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiV2IntegrationsItems) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiV2IntegrationsItems) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiV2IntegrationsItems) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *ApiV2IntegrationsItems) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiV2IntegrationsItems) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiV2IntegrationsItems) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ApiV2IntegrationsItems) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


