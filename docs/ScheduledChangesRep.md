# ScheduledChangesRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**MaintainerId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**ExecutionDate** | Pointer to **int64** |  | [optional] 
**Instructions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Conflicts** | Pointer to **interface{}** |  | [optional] 
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 

## Methods

### NewScheduledChangesRep

`func NewScheduledChangesRep() *ScheduledChangesRep`

NewScheduledChangesRep instantiates a new ScheduledChangesRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledChangesRepWithDefaults

`func NewScheduledChangesRepWithDefaults() *ScheduledChangesRep`

NewScheduledChangesRepWithDefaults instantiates a new ScheduledChangesRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledChangesRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledChangesRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledChangesRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduledChangesRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationDate

`func (o *ScheduledChangesRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ScheduledChangesRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ScheduledChangesRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ScheduledChangesRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetMaintainerId

`func (o *ScheduledChangesRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ScheduledChangesRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ScheduledChangesRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ScheduledChangesRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetVersion

`func (o *ScheduledChangesRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScheduledChangesRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScheduledChangesRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ScheduledChangesRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetExecutionDate

`func (o *ScheduledChangesRep) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ScheduledChangesRep) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ScheduledChangesRep) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *ScheduledChangesRep) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetInstructions

`func (o *ScheduledChangesRep) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ScheduledChangesRep) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ScheduledChangesRep) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ScheduledChangesRep) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetConflicts

`func (o *ScheduledChangesRep) GetConflicts() interface{}`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ScheduledChangesRep) GetConflictsOk() (*interface{}, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ScheduledChangesRep) SetConflicts(v interface{})`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *ScheduledChangesRep) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### SetConflictsNil

`func (o *ScheduledChangesRep) SetConflictsNil(b bool)`

 SetConflictsNil sets the value for Conflicts to be an explicit nil

### UnsetConflicts
`func (o *ScheduledChangesRep) UnsetConflicts()`

UnsetConflicts ensures that no value is present for Conflicts, not even an explicit nil
### GetLinks

`func (o *ScheduledChangesRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ScheduledChangesRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ScheduledChangesRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ScheduledChangesRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


