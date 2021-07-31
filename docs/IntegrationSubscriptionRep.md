# IntegrationSubscriptionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
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

### NewIntegrationSubscriptionRep

`func NewIntegrationSubscriptionRep() *IntegrationSubscriptionRep`

NewIntegrationSubscriptionRep instantiates a new IntegrationSubscriptionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSubscriptionRepWithDefaults

`func NewIntegrationSubscriptionRepWithDefaults() *IntegrationSubscriptionRep`

NewIntegrationSubscriptionRepWithDefaults instantiates a new IntegrationSubscriptionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntegrationSubscriptionRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntegrationSubscriptionRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntegrationSubscriptionRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IntegrationSubscriptionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *IntegrationSubscriptionRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationSubscriptionRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationSubscriptionRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationSubscriptionRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *IntegrationSubscriptionRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IntegrationSubscriptionRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IntegrationSubscriptionRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IntegrationSubscriptionRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *IntegrationSubscriptionRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationSubscriptionRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationSubscriptionRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationSubscriptionRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationSubscriptionRep) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationSubscriptionRep) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationSubscriptionRep) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationSubscriptionRep) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatements

`func (o *IntegrationSubscriptionRep) GetStatements() []StatementRep`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *IntegrationSubscriptionRep) GetStatementsOk() (*[]StatementRep, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *IntegrationSubscriptionRep) SetStatements(v []StatementRep)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *IntegrationSubscriptionRep) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### GetOn

`func (o *IntegrationSubscriptionRep) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *IntegrationSubscriptionRep) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *IntegrationSubscriptionRep) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *IntegrationSubscriptionRep) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetTags

`func (o *IntegrationSubscriptionRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IntegrationSubscriptionRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IntegrationSubscriptionRep) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IntegrationSubscriptionRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAccess

`func (o *IntegrationSubscriptionRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *IntegrationSubscriptionRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *IntegrationSubscriptionRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *IntegrationSubscriptionRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationSubscriptionRep) GetStatus() IntegrationSubscriptionStatusRep`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationSubscriptionRep) GetStatusOk() (*IntegrationSubscriptionStatusRep, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationSubscriptionRep) SetStatus(v IntegrationSubscriptionStatusRep)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationSubscriptionRep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *IntegrationSubscriptionRep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IntegrationSubscriptionRep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IntegrationSubscriptionRep) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IntegrationSubscriptionRep) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *IntegrationSubscriptionRep) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *IntegrationSubscriptionRep) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *IntegrationSubscriptionRep) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *IntegrationSubscriptionRep) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


