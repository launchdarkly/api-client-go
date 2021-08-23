# ScheduledChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreationDate** | **int64** |  | 
**MaintainerId** | **string** |  | 
**Version** | **int32** |  | 
**ExecutionDate** | **int64** |  | 
**Instructions** | **[]map[string]interface{}** |  | 
**Conflicts** | Pointer to **interface{}** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewScheduledChange

`func NewScheduledChange(id string, creationDate int64, maintainerId string, version int32, executionDate int64, instructions []map[string]interface{}, ) *ScheduledChange`

NewScheduledChange instantiates a new ScheduledChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledChangeWithDefaults

`func NewScheduledChangeWithDefaults() *ScheduledChange`

NewScheduledChangeWithDefaults instantiates a new ScheduledChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledChange) SetId(v string)`

SetId sets Id field to given value.


### GetCreationDate

`func (o *ScheduledChange) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ScheduledChange) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ScheduledChange) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetMaintainerId

`func (o *ScheduledChange) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ScheduledChange) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ScheduledChange) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetVersion

`func (o *ScheduledChange) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ScheduledChange) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ScheduledChange) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetExecutionDate

`func (o *ScheduledChange) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *ScheduledChange) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *ScheduledChange) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.


### GetInstructions

`func (o *ScheduledChange) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ScheduledChange) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ScheduledChange) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetConflicts

`func (o *ScheduledChange) GetConflicts() interface{}`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ScheduledChange) GetConflictsOk() (*interface{}, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ScheduledChange) SetConflicts(v interface{})`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *ScheduledChange) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### SetConflictsNil

`func (o *ScheduledChange) SetConflictsNil(b bool)`

 SetConflictsNil sets the value for Conflicts to be an explicit nil

### UnsetConflicts
`func (o *ScheduledChange) UnsetConflicts()`

UnsetConflicts ensures that no value is present for Conflicts, not even an explicit nil
### GetLinks

`func (o *ScheduledChange) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ScheduledChange) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ScheduledChange) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ScheduledChange) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


