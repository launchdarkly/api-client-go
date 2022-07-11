# ConditionBaseOutputRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Execution** | [**ExecutionOutputRep**](ExecutionOutputRep.md) |  | 

## Methods

### NewConditionBaseOutputRep

`func NewConditionBaseOutputRep(id string, execution ExecutionOutputRep, ) *ConditionBaseOutputRep`

NewConditionBaseOutputRep instantiates a new ConditionBaseOutputRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionBaseOutputRepWithDefaults

`func NewConditionBaseOutputRepWithDefaults() *ConditionBaseOutputRep`

NewConditionBaseOutputRepWithDefaults instantiates a new ConditionBaseOutputRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConditionBaseOutputRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConditionBaseOutputRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConditionBaseOutputRep) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ConditionBaseOutputRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConditionBaseOutputRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConditionBaseOutputRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConditionBaseOutputRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetExecution

`func (o *ConditionBaseOutputRep) GetExecution() ExecutionOutputRep`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ConditionBaseOutputRep) GetExecutionOk() (*ExecutionOutputRep, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ConditionBaseOutputRep) SetExecution(v ExecutionOutputRep)`

SetExecution sets Execution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

