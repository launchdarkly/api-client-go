# VariationTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the tool to use. | 
**Version** | **int32** | The version of the tool. | 
**CustomParameters** | Pointer to **map[string]interface{}** | Custom metadata and configuration for application-level use | [optional] 

## Methods

### NewVariationTool

`func NewVariationTool(key string, version int32, ) *VariationTool`

NewVariationTool instantiates a new VariationTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationToolWithDefaults

`func NewVariationToolWithDefaults() *VariationTool`

NewVariationToolWithDefaults instantiates a new VariationTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariationTool) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariationTool) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariationTool) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *VariationTool) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VariationTool) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VariationTool) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCustomParameters

`func (o *VariationTool) GetCustomParameters() map[string]interface{}`

GetCustomParameters returns the CustomParameters field if non-nil, zero value otherwise.

### GetCustomParametersOk

`func (o *VariationTool) GetCustomParametersOk() (*map[string]interface{}, bool)`

GetCustomParametersOk returns a tuple with the CustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameters

`func (o *VariationTool) SetCustomParameters(v map[string]interface{})`

SetCustomParameters sets CustomParameters field to given value.

### HasCustomParameters

`func (o *VariationTool) HasCustomParameters() bool`

HasCustomParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


