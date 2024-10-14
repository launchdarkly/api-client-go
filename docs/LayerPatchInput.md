# LayerPatchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional comment describing the update | [optional] 
**EnvironmentKey** | Pointer to **string** | The environment key used for making environment specific updates. For example, updating the reservation of an experiment | [optional] 
**Instructions** | **[]map[string]interface{}** |  | 

## Methods

### NewLayerPatchInput

`func NewLayerPatchInput(instructions []map[string]interface{}, ) *LayerPatchInput`

NewLayerPatchInput instantiates a new LayerPatchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerPatchInputWithDefaults

`func NewLayerPatchInputWithDefaults() *LayerPatchInput`

NewLayerPatchInputWithDefaults instantiates a new LayerPatchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *LayerPatchInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LayerPatchInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LayerPatchInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LayerPatchInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *LayerPatchInput) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *LayerPatchInput) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *LayerPatchInput) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *LayerPatchInput) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetInstructions

`func (o *LayerPatchInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *LayerPatchInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *LayerPatchInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


