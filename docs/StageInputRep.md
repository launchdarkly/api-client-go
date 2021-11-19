# StageInputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]ConditionInputRep**](ConditionInputRep.md) |  | [optional] 
**Action** | Pointer to [**ActionInputRep**](ActionInputRep.md) |  | [optional] 

## Methods

### NewStageInputRep

`func NewStageInputRep() *StageInputRep`

NewStageInputRep instantiates a new StageInputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageInputRepWithDefaults

`func NewStageInputRepWithDefaults() *StageInputRep`

NewStageInputRepWithDefaults instantiates a new StageInputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StageInputRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageInputRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageInputRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageInputRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConditions

`func (o *StageInputRep) GetConditions() []ConditionInputRep`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *StageInputRep) GetConditionsOk() (*[]ConditionInputRep, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *StageInputRep) SetConditions(v []ConditionInputRep)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *StageInputRep) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAction

`func (o *StageInputRep) GetAction() ActionInputRep`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StageInputRep) GetActionOk() (*ActionInputRep, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StageInputRep) SetAction(v ActionInputRep)`

SetAction sets Action field to given value.

### HasAction

`func (o *StageInputRep) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


