# StageOutputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Conditions** | [**[]ConditionOutputRep**](ConditionOutputRep.md) |  | 
**Action** | [**ActionOutputRep**](ActionOutputRep.md) |  | 
**Execution** | [**ExecutionOutputRep**](ExecutionOutputRep.md) |  | 

## Methods

### NewStageOutputRep

`func NewStageOutputRep(id string, conditions []ConditionOutputRep, action ActionOutputRep, execution ExecutionOutputRep, ) *StageOutputRep`

NewStageOutputRep instantiates a new StageOutputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageOutputRepWithDefaults

`func NewStageOutputRepWithDefaults() *StageOutputRep`

NewStageOutputRepWithDefaults instantiates a new StageOutputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StageOutputRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StageOutputRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StageOutputRep) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StageOutputRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageOutputRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageOutputRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageOutputRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConditions

`func (o *StageOutputRep) GetConditions() []ConditionOutputRep`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *StageOutputRep) GetConditionsOk() (*[]ConditionOutputRep, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *StageOutputRep) SetConditions(v []ConditionOutputRep)`

SetConditions sets Conditions field to given value.


### GetAction

`func (o *StageOutputRep) GetAction() ActionOutputRep`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StageOutputRep) GetActionOk() (*ActionOutputRep, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StageOutputRep) SetAction(v ActionOutputRep)`

SetAction sets Action field to given value.


### GetExecution

`func (o *StageOutputRep) GetExecution() ExecutionOutputRep`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *StageOutputRep) GetExecutionOk() (*ExecutionOutputRep, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *StageOutputRep) SetExecution(v ExecutionOutputRep)`

SetExecution sets Execution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


