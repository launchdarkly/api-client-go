# StageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The stage name | [optional] 
**ExecuteConditionsInSequence** | Pointer to **bool** | Whether to execute the conditions in sequence for the given stage | [optional] 
**Conditions** | Pointer to [**[]ConditionInput**](ConditionInput.md) | An array of conditions for the stage | [optional] 
**Action** | Pointer to [**ActionInput**](ActionInput.md) |  | [optional] 

## Methods

### NewStageInput

`func NewStageInput() *StageInput`

NewStageInput instantiates a new StageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageInputWithDefaults

`func NewStageInputWithDefaults() *StageInput`

NewStageInputWithDefaults instantiates a new StageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StageInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecuteConditionsInSequence

`func (o *StageInput) GetExecuteConditionsInSequence() bool`

GetExecuteConditionsInSequence returns the ExecuteConditionsInSequence field if non-nil, zero value otherwise.

### GetExecuteConditionsInSequenceOk

`func (o *StageInput) GetExecuteConditionsInSequenceOk() (*bool, bool)`

GetExecuteConditionsInSequenceOk returns a tuple with the ExecuteConditionsInSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteConditionsInSequence

`func (o *StageInput) SetExecuteConditionsInSequence(v bool)`

SetExecuteConditionsInSequence sets ExecuteConditionsInSequence field to given value.

### HasExecuteConditionsInSequence

`func (o *StageInput) HasExecuteConditionsInSequence() bool`

HasExecuteConditionsInSequence returns a boolean if a field has been set.

### GetConditions

`func (o *StageInput) GetConditions() []ConditionInput`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *StageInput) GetConditionsOk() (*[]ConditionInput, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *StageInput) SetConditions(v []ConditionInput)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *StageInput) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAction

`func (o *StageInput) GetAction() ActionInput`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StageInput) GetActionOk() (*ActionInput, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StageInput) SetAction(v ActionInput)`

SetAction sets Action field to given value.

### HasAction

`func (o *StageInput) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


