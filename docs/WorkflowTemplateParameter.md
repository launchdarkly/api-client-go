# WorkflowTemplateParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** | The path of the property to parameterize, relative to its parent condition or instruction | [optional] 
**Default** | Pointer to [**ParameterDefault**](ParameterDefault.md) |  | [optional] 
**Valid** | Pointer to **bool** | Whether the default value is valid for the target flag and environment | [optional] 

## Methods

### NewWorkflowTemplateParameter

`func NewWorkflowTemplateParameter() *WorkflowTemplateParameter`

NewWorkflowTemplateParameter instantiates a new WorkflowTemplateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateParameterWithDefaults

`func NewWorkflowTemplateParameterWithDefaults() *WorkflowTemplateParameter`

NewWorkflowTemplateParameterWithDefaults instantiates a new WorkflowTemplateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTemplateParameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTemplateParameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTemplateParameter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTemplateParameter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *WorkflowTemplateParameter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkflowTemplateParameter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkflowTemplateParameter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WorkflowTemplateParameter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDefault

`func (o *WorkflowTemplateParameter) GetDefault() ParameterDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WorkflowTemplateParameter) GetDefaultOk() (*ParameterDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WorkflowTemplateParameter) SetDefault(v ParameterDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WorkflowTemplateParameter) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetValid

`func (o *WorkflowTemplateParameter) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *WorkflowTemplateParameter) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *WorkflowTemplateParameter) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *WorkflowTemplateParameter) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


