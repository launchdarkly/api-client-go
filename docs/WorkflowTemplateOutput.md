# WorkflowTemplateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**CreationDate** | **int64** |  | 
**OwnerId** | **string** |  | 
**MaintainerId** | **string** |  | 
**Links** | [**map[string]Link**](Link.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Stages** | Pointer to [**[]StageOutput**](StageOutput.md) |  | [optional] 

## Methods

### NewWorkflowTemplateOutput

`func NewWorkflowTemplateOutput(id string, key string, creationDate int64, ownerId string, maintainerId string, links map[string]Link, ) *WorkflowTemplateOutput`

NewWorkflowTemplateOutput instantiates a new WorkflowTemplateOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateOutputWithDefaults

`func NewWorkflowTemplateOutputWithDefaults() *WorkflowTemplateOutput`

NewWorkflowTemplateOutputWithDefaults instantiates a new WorkflowTemplateOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTemplateOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTemplateOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTemplateOutput) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *WorkflowTemplateOutput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowTemplateOutput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowTemplateOutput) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *WorkflowTemplateOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTemplateOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTemplateOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTemplateOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationDate

`func (o *WorkflowTemplateOutput) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *WorkflowTemplateOutput) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *WorkflowTemplateOutput) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetOwnerId

`func (o *WorkflowTemplateOutput) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *WorkflowTemplateOutput) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *WorkflowTemplateOutput) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetMaintainerId

`func (o *WorkflowTemplateOutput) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *WorkflowTemplateOutput) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *WorkflowTemplateOutput) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetLinks

`func (o *WorkflowTemplateOutput) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkflowTemplateOutput) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkflowTemplateOutput) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetDescription

`func (o *WorkflowTemplateOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTemplateOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTemplateOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTemplateOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStages

`func (o *WorkflowTemplateOutput) GetStages() []StageOutput`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *WorkflowTemplateOutput) GetStagesOk() (*[]StageOutput, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *WorkflowTemplateOutput) SetStages(v []StageOutput)`

SetStages sets Stages field to given value.

### HasStages

`func (o *WorkflowTemplateOutput) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


