# AIToolPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**MaintainerId** | Pointer to **string** |  | [optional] 
**MaintainerTeamKey** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Schema** | **map[string]interface{}** |  | 

## Methods

### NewAIToolPost

`func NewAIToolPost(key string, schema map[string]interface{}, ) *AIToolPost`

NewAIToolPost instantiates a new AIToolPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIToolPostWithDefaults

`func NewAIToolPostWithDefaults() *AIToolPost`

NewAIToolPostWithDefaults instantiates a new AIToolPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AIToolPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIToolPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIToolPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaintainerId

`func (o *AIToolPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *AIToolPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *AIToolPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *AIToolPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerTeamKey

`func (o *AIToolPost) GetMaintainerTeamKey() string`

GetMaintainerTeamKey returns the MaintainerTeamKey field if non-nil, zero value otherwise.

### GetMaintainerTeamKeyOk

`func (o *AIToolPost) GetMaintainerTeamKeyOk() (*string, bool)`

GetMaintainerTeamKeyOk returns a tuple with the MaintainerTeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeamKey

`func (o *AIToolPost) SetMaintainerTeamKey(v string)`

SetMaintainerTeamKey sets MaintainerTeamKey field to given value.

### HasMaintainerTeamKey

`func (o *AIToolPost) HasMaintainerTeamKey() bool`

HasMaintainerTeamKey returns a boolean if a field has been set.

### GetDescription

`func (o *AIToolPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIToolPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIToolPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIToolPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchema

`func (o *AIToolPost) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AIToolPost) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AIToolPost) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


