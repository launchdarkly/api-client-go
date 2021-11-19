# CustomWorkflowOutputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | **int32** |  | 
**Conflicts** | [**[]ConflictOutputRep**](ConflictOutputRep.md) |  | 
**CreationDate** | **int64** |  | 
**MaintainerId** | **string** |  | 
**Links** | [**map[string]Link**](Link.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Stages** | Pointer to [**[]StageOutputRep**](StageOutputRep.md) |  | [optional] 
**Execution** | [**ExecutionOutputRep**](ExecutionOutputRep.md) |  | 

## Methods

### NewCustomWorkflowOutputRep

`func NewCustomWorkflowOutputRep(id string, version int32, conflicts []ConflictOutputRep, creationDate int64, maintainerId string, links map[string]Link, name string, execution ExecutionOutputRep, ) *CustomWorkflowOutputRep`

NewCustomWorkflowOutputRep instantiates a new CustomWorkflowOutputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWorkflowOutputRepWithDefaults

`func NewCustomWorkflowOutputRepWithDefaults() *CustomWorkflowOutputRep`

NewCustomWorkflowOutputRepWithDefaults instantiates a new CustomWorkflowOutputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomWorkflowOutputRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomWorkflowOutputRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomWorkflowOutputRep) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *CustomWorkflowOutputRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CustomWorkflowOutputRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CustomWorkflowOutputRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetConflicts

`func (o *CustomWorkflowOutputRep) GetConflicts() []ConflictOutputRep`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *CustomWorkflowOutputRep) GetConflictsOk() (*[]ConflictOutputRep, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *CustomWorkflowOutputRep) SetConflicts(v []ConflictOutputRep)`

SetConflicts sets Conflicts field to given value.


### GetCreationDate

`func (o *CustomWorkflowOutputRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CustomWorkflowOutputRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CustomWorkflowOutputRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetMaintainerId

`func (o *CustomWorkflowOutputRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *CustomWorkflowOutputRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *CustomWorkflowOutputRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetLinks

`func (o *CustomWorkflowOutputRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomWorkflowOutputRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomWorkflowOutputRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetName

`func (o *CustomWorkflowOutputRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomWorkflowOutputRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomWorkflowOutputRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomWorkflowOutputRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomWorkflowOutputRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomWorkflowOutputRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomWorkflowOutputRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *CustomWorkflowOutputRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CustomWorkflowOutputRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CustomWorkflowOutputRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CustomWorkflowOutputRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetStages

`func (o *CustomWorkflowOutputRep) GetStages() []StageOutputRep`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CustomWorkflowOutputRep) GetStagesOk() (*[]StageOutputRep, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CustomWorkflowOutputRep) SetStages(v []StageOutputRep)`

SetStages sets Stages field to given value.

### HasStages

`func (o *CustomWorkflowOutputRep) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetExecution

`func (o *CustomWorkflowOutputRep) GetExecution() ExecutionOutputRep`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *CustomWorkflowOutputRep) GetExecutionOk() (*ExecutionOutputRep, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *CustomWorkflowOutputRep) SetExecution(v ExecutionOutputRep)`

SetExecution sets Execution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


