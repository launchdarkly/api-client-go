# ToolReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiConfigKey** | **string** | The key of the AgentControl config with a variation that references this tool. | 
**AiConfigName** | **string** | The name of the AgentControl config with a variation that references this tool. | 
**VariationId** | **string** | The ID of the variation that references this tool. | 
**VariationKey** | **string** | The key of the AgentControl config variation that references this tool. | 
**VariationName** | **string** | The name of the variation that references this tool. | 
**ToolVersion** | **int32** | The version of the tool being referenced. | 

## Methods

### NewToolReference

`func NewToolReference(aiConfigKey string, aiConfigName string, variationId string, variationKey string, variationName string, toolVersion int32, ) *ToolReference`

NewToolReference instantiates a new ToolReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolReferenceWithDefaults

`func NewToolReferenceWithDefaults() *ToolReference`

NewToolReferenceWithDefaults instantiates a new ToolReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiConfigKey

`func (o *ToolReference) GetAiConfigKey() string`

GetAiConfigKey returns the AiConfigKey field if non-nil, zero value otherwise.

### GetAiConfigKeyOk

`func (o *ToolReference) GetAiConfigKeyOk() (*string, bool)`

GetAiConfigKeyOk returns a tuple with the AiConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigKey

`func (o *ToolReference) SetAiConfigKey(v string)`

SetAiConfigKey sets AiConfigKey field to given value.


### GetAiConfigName

`func (o *ToolReference) GetAiConfigName() string`

GetAiConfigName returns the AiConfigName field if non-nil, zero value otherwise.

### GetAiConfigNameOk

`func (o *ToolReference) GetAiConfigNameOk() (*string, bool)`

GetAiConfigNameOk returns a tuple with the AiConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigName

`func (o *ToolReference) SetAiConfigName(v string)`

SetAiConfigName sets AiConfigName field to given value.


### GetVariationId

`func (o *ToolReference) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *ToolReference) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *ToolReference) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.


### GetVariationKey

`func (o *ToolReference) GetVariationKey() string`

GetVariationKey returns the VariationKey field if non-nil, zero value otherwise.

### GetVariationKeyOk

`func (o *ToolReference) GetVariationKeyOk() (*string, bool)`

GetVariationKeyOk returns a tuple with the VariationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationKey

`func (o *ToolReference) SetVariationKey(v string)`

SetVariationKey sets VariationKey field to given value.


### GetVariationName

`func (o *ToolReference) GetVariationName() string`

GetVariationName returns the VariationName field if non-nil, zero value otherwise.

### GetVariationNameOk

`func (o *ToolReference) GetVariationNameOk() (*string, bool)`

GetVariationNameOk returns a tuple with the VariationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationName

`func (o *ToolReference) SetVariationName(v string)`

SetVariationName sets VariationName field to given value.


### GetToolVersion

`func (o *ToolReference) GetToolVersion() int32`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *ToolReference) GetToolVersionOk() (*int32, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *ToolReference) SetToolVersion(v int32)`

SetToolVersion sets ToolVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


