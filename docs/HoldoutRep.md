# HoldoutRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**HoldoutAmount** | **string** | The percentage of traffic allocated to this holdout. | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**BaseExperiment** | [**Experiment**](Experiment.md) |  | 
**Experiments** | Pointer to [**[]RelatedExperimentRep**](RelatedExperimentRep.md) |  | [optional] 

## Methods

### NewHoldoutRep

`func NewHoldoutRep(id string, status string, holdoutAmount string, createdAt int64, updatedAt int64, baseExperiment Experiment, ) *HoldoutRep`

NewHoldoutRep instantiates a new HoldoutRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutRepWithDefaults

`func NewHoldoutRepWithDefaults() *HoldoutRep`

NewHoldoutRepWithDefaults instantiates a new HoldoutRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HoldoutRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HoldoutRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HoldoutRep) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *HoldoutRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HoldoutRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HoldoutRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *HoldoutRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldoutRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHoldoutAmount

`func (o *HoldoutRep) GetHoldoutAmount() string`

GetHoldoutAmount returns the HoldoutAmount field if non-nil, zero value otherwise.

### GetHoldoutAmountOk

`func (o *HoldoutRep) GetHoldoutAmountOk() (*string, bool)`

GetHoldoutAmountOk returns a tuple with the HoldoutAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutAmount

`func (o *HoldoutRep) SetHoldoutAmount(v string)`

SetHoldoutAmount sets HoldoutAmount field to given value.


### GetCreatedAt

`func (o *HoldoutRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HoldoutRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HoldoutRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HoldoutRep) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HoldoutRep) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HoldoutRep) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBaseExperiment

`func (o *HoldoutRep) GetBaseExperiment() Experiment`

GetBaseExperiment returns the BaseExperiment field if non-nil, zero value otherwise.

### GetBaseExperimentOk

`func (o *HoldoutRep) GetBaseExperimentOk() (*Experiment, bool)`

GetBaseExperimentOk returns a tuple with the BaseExperiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseExperiment

`func (o *HoldoutRep) SetBaseExperiment(v Experiment)`

SetBaseExperiment sets BaseExperiment field to given value.


### GetExperiments

`func (o *HoldoutRep) GetExperiments() []RelatedExperimentRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *HoldoutRep) GetExperimentsOk() (*[]RelatedExperimentRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *HoldoutRep) SetExperiments(v []RelatedExperimentRep)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *HoldoutRep) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


