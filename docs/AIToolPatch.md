# AIToolPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintainerId** | Pointer to **string** |  | [optional] 
**MaintainerTeamKey** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **map[string]interface{}** | JSON Schema defining the tool&#39;s parameters for LLM consumption | [optional] 
**CustomParameters** | Pointer to **map[string]interface{}** | Custom metadata and configuration for application-level use (not sent to LLM) | [optional] 

## Methods

### NewAIToolPatch

`func NewAIToolPatch() *AIToolPatch`

NewAIToolPatch instantiates a new AIToolPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIToolPatchWithDefaults

`func NewAIToolPatchWithDefaults() *AIToolPatch`

NewAIToolPatchWithDefaults instantiates a new AIToolPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintainerId

`func (o *AIToolPatch) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AIToolPatch) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AIToolPatch) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AIToolPatch) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *AIToolPatch) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *AIToolPatch) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *AIToolPatch) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *AIToolPatch) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetDescription

`func (o *AIToolPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIToolPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIToolPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIToolPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchema

`func (o *AIToolPatch) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AIToolPatch) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AIToolPatch) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AIToolPatch) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCustomParameters

`func (o *AIToolPatch) GetCustomParameters() map[string]interface{}`

GetCustomParameters returns the CustomParameters field if non-nil, zero value otherwise.

### GetCustomParametersOk

`func (o *AIToolPatch) GetCustomParametersOk() (*map[string]interface{}, bool)`

GetCustomParametersOk returns a tuple with the CustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameters

`func (o *AIToolPatch) SetCustomParameters(v map[string]interface{})`

SetCustomParameters sets CustomParameters field to given value.

### HasCustomParameters

`func (o *AIToolPatch) HasCustomParameters() bool`

HasCustomParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


