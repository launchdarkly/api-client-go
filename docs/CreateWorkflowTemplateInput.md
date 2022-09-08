# CreateWorkflowTemplateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**Stages** | Pointer to [**[]StageInput**](StageInput.md) |  | [optional] 
**ProjectKey** | Pointer to **string** |  | [optional] 
**EnvironmentKey** | Pointer to **string** |  | [optional] 
**FlagKey** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateWorkflowTemplateInput

`func NewCreateWorkflowTemplateInput(key string, ) *CreateWorkflowTemplateInput`

NewCreateWorkflowTemplateInput instantiates a new CreateWorkflowTemplateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowTemplateInputWithDefaults

`func NewCreateWorkflowTemplateInputWithDefaults() *CreateWorkflowTemplateInput`

NewCreateWorkflowTemplateInputWithDefaults instantiates a new CreateWorkflowTemplateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateWorkflowTemplateInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateWorkflowTemplateInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateWorkflowTemplateInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateWorkflowTemplateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflowTemplateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflowTemplateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWorkflowTemplateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateWorkflowTemplateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkflowTemplateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkflowTemplateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorkflowTemplateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkflowId

`func (o *CreateWorkflowTemplateInput) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CreateWorkflowTemplateInput) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CreateWorkflowTemplateInput) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *CreateWorkflowTemplateInput) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetStages

`func (o *CreateWorkflowTemplateInput) GetStages() []StageInput`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CreateWorkflowTemplateInput) GetStagesOk() (*[]StageInput, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CreateWorkflowTemplateInput) SetStages(v []StageInput)`

SetStages sets Stages field to given value.

### HasStages

`func (o *CreateWorkflowTemplateInput) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetProjectKey

`func (o *CreateWorkflowTemplateInput) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *CreateWorkflowTemplateInput) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *CreateWorkflowTemplateInput) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *CreateWorkflowTemplateInput) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *CreateWorkflowTemplateInput) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *CreateWorkflowTemplateInput) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *CreateWorkflowTemplateInput) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *CreateWorkflowTemplateInput) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetFlagKey

`func (o *CreateWorkflowTemplateInput) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *CreateWorkflowTemplateInput) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *CreateWorkflowTemplateInput) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *CreateWorkflowTemplateInput) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


