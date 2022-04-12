# Conflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instruction** | Pointer to **interface{}** |  | [optional] 
**Reason** | Pointer to **string** | Reason why the conflict exists | [optional] 

## Methods

### NewConflict

`func NewConflict() *Conflict`

NewConflict instantiates a new Conflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictWithDefaults

`func NewConflictWithDefaults() *Conflict`

NewConflictWithDefaults instantiates a new Conflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstruction

`func (o *Conflict) GetInstruction() interface{}`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *Conflict) GetInstructionOk() (*interface{}, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *Conflict) SetInstruction(v interface{})`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *Conflict) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### SetInstructionNil

`func (o *Conflict) SetInstructionNil(b bool)`

 SetInstructionNil sets the value for Instruction to be an explicit nil

### UnsetInstruction
`func (o *Conflict) UnsetInstruction()`

UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
### GetReason

`func (o *Conflict) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Conflict) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Conflict) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Conflict) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


