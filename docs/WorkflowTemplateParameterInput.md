# WorkflowTemplateParameterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Default** | Pointer to [**ParameterDefaultInput**](ParameterDefaultInput.md) |  | [optional] 

## Methods

### NewWorkflowTemplateParameterInput

`func NewWorkflowTemplateParameterInput() *WorkflowTemplateParameterInput`

NewWorkflowTemplateParameterInput instantiates a new WorkflowTemplateParameterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateParameterInputWithDefaults

`func NewWorkflowTemplateParameterInputWithDefaults() *WorkflowTemplateParameterInput`

NewWorkflowTemplateParameterInputWithDefaults instantiates a new WorkflowTemplateParameterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTemplateParameterInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTemplateParameterInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTemplateParameterInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTemplateParameterInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *WorkflowTemplateParameterInput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkflowTemplateParameterInput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkflowTemplateParameterInput) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WorkflowTemplateParameterInput) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDefault

`func (o *WorkflowTemplateParameterInput) GetDefault() ParameterDefaultInput`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WorkflowTemplateParameterInput) GetDefaultOk() (*ParameterDefaultInput, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WorkflowTemplateParameterInput) SetDefault(v ParameterDefaultInput)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WorkflowTemplateParameterInput) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


