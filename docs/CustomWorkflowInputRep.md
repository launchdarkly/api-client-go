# CustomWorkflowInputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintainerId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The workflow name | [optional] 
**Description** | **string** | The workflow description | 
**Stages** | Pointer to [**[]StageInputRep**](StageInputRep.md) |  | [optional] 
**TemplateKey** | Pointer to **string** | The template key | [optional] 

## Methods

### NewCustomWorkflowInputRep

`func NewCustomWorkflowInputRep(description string, ) *CustomWorkflowInputRep`

NewCustomWorkflowInputRep instantiates a new CustomWorkflowInputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWorkflowInputRepWithDefaults

`func NewCustomWorkflowInputRepWithDefaults() *CustomWorkflowInputRep`

NewCustomWorkflowInputRepWithDefaults instantiates a new CustomWorkflowInputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintainerId

`func (o *CustomWorkflowInputRep) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *CustomWorkflowInputRep) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *CustomWorkflowInputRep) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *CustomWorkflowInputRep) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetName

`func (o *CustomWorkflowInputRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomWorkflowInputRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomWorkflowInputRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomWorkflowInputRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CustomWorkflowInputRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomWorkflowInputRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomWorkflowInputRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStages

`func (o *CustomWorkflowInputRep) GetStages() []StageInputRep`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CustomWorkflowInputRep) GetStagesOk() (*[]StageInputRep, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CustomWorkflowInputRep) SetStages(v []StageInputRep)`

SetStages sets Stages field to given value.

### HasStages

`func (o *CustomWorkflowInputRep) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetTemplateKey

`func (o *CustomWorkflowInputRep) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *CustomWorkflowInputRep) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *CustomWorkflowInputRep) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *CustomWorkflowInputRep) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


