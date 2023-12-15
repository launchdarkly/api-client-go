# CustomWorkflowOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the workflow | 
**Version** | **int32** | The version of the workflow | 
**Conflicts** | [**[]ConflictOutput**](ConflictOutput.md) | Any conflicts that are present in the workflow stages | 
**CreationDate** | **int64** |  | 
**MaintainerId** | **string** | The member ID of the maintainer of the workflow. Defaults to the workflow creator. | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Name** | **string** | The name of the workflow | 
**Description** | Pointer to **string** | A brief description of the workflow | [optional] 
**Kind** | Pointer to **string** | The kind of workflow | [optional] 
**Stages** | Pointer to [**[]StageOutput**](StageOutput.md) | The stages that make up the workflow. Each stage contains conditions and actions. | [optional] 
**Execution** | [**ExecutionOutput**](ExecutionOutput.md) |  | 
**Meta** | Pointer to [**WorkflowTemplateMetadata**](WorkflowTemplateMetadata.md) |  | [optional] 
**TemplateKey** | Pointer to **string** | For workflows being created from a workflow template, this value is the template&#39;s key | [optional] 
**Environment** | Pointer to [**EnvironmentSummary**](EnvironmentSummary.md) |  | [optional] 

## Methods

### NewCustomWorkflowOutput

`func NewCustomWorkflowOutput(id string, version int32, conflicts []ConflictOutput, creationDate int64, maintainerId string, links map[string]Link, name string, execution ExecutionOutput, ) *CustomWorkflowOutput`

NewCustomWorkflowOutput instantiates a new CustomWorkflowOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWorkflowOutputWithDefaults

`func NewCustomWorkflowOutputWithDefaults() *CustomWorkflowOutput`

NewCustomWorkflowOutputWithDefaults instantiates a new CustomWorkflowOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomWorkflowOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomWorkflowOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomWorkflowOutput) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *CustomWorkflowOutput) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CustomWorkflowOutput) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CustomWorkflowOutput) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetConflicts

`func (o *CustomWorkflowOutput) GetConflicts() []ConflictOutput`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *CustomWorkflowOutput) GetConflictsOk() (*[]ConflictOutput, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *CustomWorkflowOutput) SetConflicts(v []ConflictOutput)`

SetConflicts sets Conflicts field to given value.


### GetCreationDate

`func (o *CustomWorkflowOutput) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CustomWorkflowOutput) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CustomWorkflowOutput) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetMaintainerId

`func (o *CustomWorkflowOutput) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *CustomWorkflowOutput) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *CustomWorkflowOutput) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetLinks

`func (o *CustomWorkflowOutput) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomWorkflowOutput) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomWorkflowOutput) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetName

`func (o *CustomWorkflowOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomWorkflowOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomWorkflowOutput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomWorkflowOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomWorkflowOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomWorkflowOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomWorkflowOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *CustomWorkflowOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CustomWorkflowOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CustomWorkflowOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CustomWorkflowOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetStages

`func (o *CustomWorkflowOutput) GetStages() []StageOutput`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CustomWorkflowOutput) GetStagesOk() (*[]StageOutput, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CustomWorkflowOutput) SetStages(v []StageOutput)`

SetStages sets Stages field to given value.

### HasStages

`func (o *CustomWorkflowOutput) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetExecution

`func (o *CustomWorkflowOutput) GetExecution() ExecutionOutput`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *CustomWorkflowOutput) GetExecutionOk() (*ExecutionOutput, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *CustomWorkflowOutput) SetExecution(v ExecutionOutput)`

SetExecution sets Execution field to given value.


### GetMeta

`func (o *CustomWorkflowOutput) GetMeta() WorkflowTemplateMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomWorkflowOutput) GetMetaOk() (*WorkflowTemplateMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomWorkflowOutput) SetMeta(v WorkflowTemplateMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CustomWorkflowOutput) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTemplateKey

`func (o *CustomWorkflowOutput) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *CustomWorkflowOutput) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *CustomWorkflowOutput) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *CustomWorkflowOutput) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.

### GetEnvironment

`func (o *CustomWorkflowOutput) GetEnvironment() EnvironmentSummary`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CustomWorkflowOutput) GetEnvironmentOk() (*EnvironmentSummary, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CustomWorkflowOutput) SetEnvironment(v EnvironmentSummary)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CustomWorkflowOutput) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


