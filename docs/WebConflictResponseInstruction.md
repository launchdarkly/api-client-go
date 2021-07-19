# WebConflictResponseInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Conflicts** | Pointer to [**[]WebConflictResponseConflicts**](WebConflictResponseConflicts.md) |  | [optional] 

## Methods

### NewWebConflictResponseInstruction

`func NewWebConflictResponseInstruction() *WebConflictResponseInstruction`

NewWebConflictResponseInstruction instantiates a new WebConflictResponseInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebConflictResponseInstructionWithDefaults

`func NewWebConflictResponseInstructionWithDefaults() *WebConflictResponseInstruction`

NewWebConflictResponseInstructionWithDefaults instantiates a new WebConflictResponseInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *WebConflictResponseInstruction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebConflictResponseInstruction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebConflictResponseInstruction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebConflictResponseInstruction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetConflicts

`func (o *WebConflictResponseInstruction) GetConflicts() []WebConflictResponseConflicts`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *WebConflictResponseInstruction) GetConflictsOk() (*[]WebConflictResponseConflicts, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *WebConflictResponseInstruction) SetConflicts(v []WebConflictResponseConflicts)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *WebConflictResponseInstruction) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


