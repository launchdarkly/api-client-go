# AITool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Access** | Pointer to [**AiConfigsAccess**](AiConfigsAccess.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Maintainer** | Pointer to [**AIConfigMaintainer**](AIConfigMaintainer.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Schema** | **map[string]interface{}** |  | 
**Version** | **int32** |  | 
**CreatedAt** | **int64** |  | 

## Methods

### NewAITool

`func NewAITool(key string, schema map[string]interface{}, version int32, createdAt int64, ) *AITool`

NewAITool instantiates a new AITool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIToolWithDefaults

`func NewAIToolWithDefaults() *AITool`

NewAIToolWithDefaults instantiates a new AITool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AITool) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AITool) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AITool) SetKey(v string)`

SetKey sets Key field to given value.


### GetAccess

`func (o *AITool) GetAccess() AiConfigsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AITool) GetAccessOk() (*AiConfigsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AITool) SetAccess(v AiConfigsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AITool) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *AITool) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AITool) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AITool) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AITool) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMaintainer

`func (o *AITool) GetMaintainer() AIConfigMaintainer`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *AITool) GetMaintainerOk() (*AIConfigMaintainer, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *AITool) SetMaintainer(v AIConfigMaintainer)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *AITool) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetDescription

`func (o *AITool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AITool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AITool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AITool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchema

`func (o *AITool) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AITool) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AITool) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetVersion

`func (o *AITool) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AITool) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AITool) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *AITool) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AITool) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AITool) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


