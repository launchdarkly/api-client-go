# ConflictResponseInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Conflicts** | [**[]ConflictResponseConflict**](ConflictResponseConflict.md) |  | 

## Methods

### NewConflictResponseInstruction

`func NewConflictResponseInstruction(kind string, conflicts []ConflictResponseConflict, ) *ConflictResponseInstruction`

NewConflictResponseInstruction instantiates a new ConflictResponseInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictResponseInstructionWithDefaults

`func NewConflictResponseInstructionWithDefaults() *ConflictResponseInstruction`

NewConflictResponseInstructionWithDefaults instantiates a new ConflictResponseInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConflictResponseInstruction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConflictResponseInstruction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConflictResponseInstruction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetConflicts

`func (o *ConflictResponseInstruction) GetConflicts() []ConflictResponseConflict`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ConflictResponseInstruction) GetConflictsOk() (*[]ConflictResponseConflict, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ConflictResponseInstruction) SetConflicts(v []ConflictResponseConflict)`

SetConflicts sets Conflicts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


