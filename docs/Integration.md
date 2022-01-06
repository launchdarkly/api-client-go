# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
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

### NewIntegration

`func NewIntegration() *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Integration) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Integration) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Integration) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Integration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Integration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Integration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Integration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Integration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Integration) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Integration) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Integration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *Integration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Integration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Integration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatements

`func (o *Integration) GetStatements() []StatementRep`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *Integration) GetStatementsOk() (*[]StatementRep, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *Integration) SetStatements(v []StatementRep)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *Integration) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetOn

`func (o *Integration) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *Integration) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *Integration) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *Integration) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetTags

`func (o *Integration) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Integration) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Integration) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Integration) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccess

`func (o *Integration) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Integration) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Integration) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Integration) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetStatus

`func (o *Integration) GetStatus() IntegrationSubscriptionStatusRep`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Integration) GetStatusOk() (*IntegrationSubscriptionStatusRep, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Integration) SetStatus(v IntegrationSubscriptionStatusRep)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Integration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *Integration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Integration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Integration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Integration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *Integration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Integration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Integration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Integration) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


