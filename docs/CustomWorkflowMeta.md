# CustomWorkflowMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to [**CustomWorkflowStageMeta**](CustomWorkflowStageMeta.md) |  | [optional] 

## Methods

### NewCustomWorkflowMeta

`func NewCustomWorkflowMeta() *CustomWorkflowMeta`

NewCustomWorkflowMeta instantiates a new CustomWorkflowMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWorkflowMetaWithDefaults

`func NewCustomWorkflowMetaWithDefaults() *CustomWorkflowMeta`

NewCustomWorkflowMetaWithDefaults instantiates a new CustomWorkflowMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomWorkflowMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomWorkflowMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomWorkflowMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomWorkflowMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStage

`func (o *CustomWorkflowMeta) GetStage() CustomWorkflowStageMeta`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CustomWorkflowMeta) GetStageOk() (*CustomWorkflowStageMeta, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CustomWorkflowMeta) SetStage(v CustomWorkflowStageMeta)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CustomWorkflowMeta) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


