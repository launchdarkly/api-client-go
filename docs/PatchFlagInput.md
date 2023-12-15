# PatchFlagInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentKey** | Pointer to **string** |  | [optional] 
**Instructions** | **[]map[string]interface{}** |  | 

## Methods

### NewPatchFlagInput

`func NewPatchFlagInput(instructions []map[string]interface{}, ) *PatchFlagInput`

NewPatchFlagInput instantiates a new PatchFlagInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFlagInputWithDefaults

`func NewPatchFlagInputWithDefaults() *PatchFlagInput`

NewPatchFlagInputWithDefaults instantiates a new PatchFlagInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentKey

`func (o *PatchFlagInput) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *PatchFlagInput) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *PatchFlagInput) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *PatchFlagInput) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetInstructions

`func (o *PatchFlagInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PatchFlagInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PatchFlagInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


