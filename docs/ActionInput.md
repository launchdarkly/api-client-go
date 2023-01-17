# ActionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to **interface{}** | An array of instructions for the stage. Each object in the array uses the semantic patch format for updating a feature flag. | [optional] 

## Methods

### NewActionInput

`func NewActionInput() *ActionInput`

NewActionInput instantiates a new ActionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionInputWithDefaults

`func NewActionInputWithDefaults() *ActionInput`

NewActionInputWithDefaults instantiates a new ActionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *ActionInput) GetInstructions() interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ActionInput) GetInstructionsOk() (*interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ActionInput) SetInstructions(v interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ActionInput) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ActionInput) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ActionInput) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


