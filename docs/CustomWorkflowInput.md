# CustomWorkflowInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintainerId** | Pointer to **string** |  | [optional] 
**Name** | **string** | The workflow name | 
**Description** | Pointer to **string** | The workflow description | [optional] 
**Stages** | Pointer to [**[]StageInput**](StageInput.md) | A list of the workflow stages | [optional] 
**TemplateKey** | Pointer to **string** | The template key | [optional] 

## Methods

### NewCustomWorkflowInput

`func NewCustomWorkflowInput(name string, ) *CustomWorkflowInput`

NewCustomWorkflowInput instantiates a new CustomWorkflowInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWorkflowInputWithDefaults

`func NewCustomWorkflowInputWithDefaults() *CustomWorkflowInput`

NewCustomWorkflowInputWithDefaults instantiates a new CustomWorkflowInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintainerId

`func (o *CustomWorkflowInput) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *CustomWorkflowInput) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *CustomWorkflowInput) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *CustomWorkflowInput) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetName

`func (o *CustomWorkflowInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomWorkflowInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomWorkflowInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomWorkflowInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomWorkflowInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomWorkflowInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomWorkflowInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStages

`func (o *CustomWorkflowInput) GetStages() []StageInput`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CustomWorkflowInput) GetStagesOk() (*[]StageInput, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CustomWorkflowInput) SetStages(v []StageInput)`

SetStages sets Stages field to given value.

### HasStages

`func (o *CustomWorkflowInput) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetTemplateKey

`func (o *CustomWorkflowInput) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *CustomWorkflowInput) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *CustomWorkflowInput) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *CustomWorkflowInput) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


