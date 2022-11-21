# StageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this stage | 
**Name** | Pointer to **string** | The stage name | [optional] 
**Conditions** | [**[]ConditionOutput**](ConditionOutput.md) | An array of conditions for the stage | 
**Action** | [**ActionOutput**](ActionOutput.md) |  | 
**Execution** | [**ExecutionOutput**](ExecutionOutput.md) |  | 

## Methods

### NewStageOutput

`func NewStageOutput(id string, conditions []ConditionOutput, action ActionOutput, execution ExecutionOutput, ) *StageOutput`

NewStageOutput instantiates a new StageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageOutputWithDefaults

`func NewStageOutputWithDefaults() *StageOutput`

NewStageOutputWithDefaults instantiates a new StageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StageOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StageOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StageOutput) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StageOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConditions

`func (o *StageOutput) GetConditions() []ConditionOutput`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *StageOutput) GetConditionsOk() (*[]ConditionOutput, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *StageOutput) SetConditions(v []ConditionOutput)`

SetConditions sets Conditions field to given value.


### GetAction

`func (o *StageOutput) GetAction() ActionOutput`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StageOutput) GetActionOk() (*ActionOutput, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StageOutput) SetAction(v ActionOutput)`

SetAction sets Action field to given value.


### GetExecution

`func (o *StageOutput) GetExecution() ExecutionOutput`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *StageOutput) GetExecutionOk() (*ExecutionOutput, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *StageOutput) SetExecution(v ExecutionOutput)`

SetExecution sets Execution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


