# AIConfigTargetingPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**EnvironmentKey** | **string** |  | 
**Instructions** | **[]map[string]interface{}** |  | 

## Methods

### NewAIConfigTargetingPatch

`func NewAIConfigTargetingPatch(environmentKey string, instructions []map[string]interface{}, ) *AIConfigTargetingPatch`

NewAIConfigTargetingPatch instantiates a new AIConfigTargetingPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigTargetingPatchWithDefaults

`func NewAIConfigTargetingPatchWithDefaults() *AIConfigTargetingPatch`

NewAIConfigTargetingPatchWithDefaults instantiates a new AIConfigTargetingPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *AIConfigTargetingPatch) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AIConfigTargetingPatch) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AIConfigTargetingPatch) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AIConfigTargetingPatch) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *AIConfigTargetingPatch) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *AIConfigTargetingPatch) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *AIConfigTargetingPatch) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetInstructions

`func (o *AIConfigTargetingPatch) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AIConfigTargetingPatch) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AIConfigTargetingPatch) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


