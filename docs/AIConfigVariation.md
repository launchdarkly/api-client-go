# AIConfigVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ParentLink**](ParentLink.md) |  | [optional] 
**Key** | **string** |  | 
**Messages** | [**[]Message**](Message.md) |  | 
**Model** | **map[string]interface{}** |  | 
**ModelConfigKey** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**CreatedAt** | **int64** |  | 
**Version** | **int32** |  | 
**State** | Pointer to **string** |  | [optional] 
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**PublishedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAIConfigVariation

`func NewAIConfigVariation(key string, messages []Message, model map[string]interface{}, name string, createdAt int64, version int32, ) *AIConfigVariation`

NewAIConfigVariation instantiates a new AIConfigVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigVariationWithDefaults

`func NewAIConfigVariationWithDefaults() *AIConfigVariation`

NewAIConfigVariationWithDefaults instantiates a new AIConfigVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AIConfigVariation) GetLinks() ParentLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AIConfigVariation) GetLinksOk() (*ParentLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AIConfigVariation) SetLinks(v ParentLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AIConfigVariation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetKey

`func (o *AIConfigVariation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIConfigVariation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIConfigVariation) SetKey(v string)`

SetKey sets Key field to given value.


### GetMessages

`func (o *AIConfigVariation) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AIConfigVariation) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AIConfigVariation) SetMessages(v []Message)`

SetMessages sets Messages field to given value.


### GetModel

`func (o *AIConfigVariation) GetModel() map[string]interface{}`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AIConfigVariation) GetModelOk() (*map[string]interface{}, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AIConfigVariation) SetModel(v map[string]interface{})`

SetModel sets Model field to given value.


### GetModelConfigKey

`func (o *AIConfigVariation) GetModelConfigKey() string`

GetModelConfigKey returns the ModelConfigKey field if non-nil, zero value otherwise.

### GetModelConfigKeyOk

`func (o *AIConfigVariation) GetModelConfigKeyOk() (*string, bool)`

GetModelConfigKeyOk returns a tuple with the ModelConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConfigKey

`func (o *AIConfigVariation) SetModelConfigKey(v string)`

SetModelConfigKey sets ModelConfigKey field to given value.

### HasModelConfigKey

`func (o *AIConfigVariation) HasModelConfigKey() bool`

HasModelConfigKey returns a boolean if a field has been set.

### GetName

`func (o *AIConfigVariation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIConfigVariation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIConfigVariation) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *AIConfigVariation) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIConfigVariation) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIConfigVariation) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetVersion

`func (o *AIConfigVariation) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AIConfigVariation) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AIConfigVariation) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetState

`func (o *AIConfigVariation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AIConfigVariation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AIConfigVariation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AIConfigVariation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetArchivedAt

`func (o *AIConfigVariation) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *AIConfigVariation) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *AIConfigVariation) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *AIConfigVariation) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *AIConfigVariation) GetPublishedAt() int64`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *AIConfigVariation) GetPublishedAtOk() (*int64, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *AIConfigVariation) SetPublishedAt(v int64)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *AIConfigVariation) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


