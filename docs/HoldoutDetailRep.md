# HoldoutDetailRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**HoldoutAmount** | **string** | The percentage of traffic allocated to this holdout. | 
**IsDirty** | Pointer to **bool** | Indicates if the holdout experiment is running and if any related experiments are running. | [optional] 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**BaseExperiment** | [**Experiment**](Experiment.md) |  | 
**RelatedExperiments** | Pointer to [**[]Experiment**](Experiment.md) |  | [optional] 

## Methods

### NewHoldoutDetailRep

`func NewHoldoutDetailRep(id string, status string, holdoutAmount string, createdAt int64, updatedAt int64, baseExperiment Experiment, ) *HoldoutDetailRep`

NewHoldoutDetailRep instantiates a new HoldoutDetailRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldoutDetailRepWithDefaults

`func NewHoldoutDetailRepWithDefaults() *HoldoutDetailRep`

NewHoldoutDetailRepWithDefaults instantiates a new HoldoutDetailRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HoldoutDetailRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HoldoutDetailRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HoldoutDetailRep) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *HoldoutDetailRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HoldoutDetailRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HoldoutDetailRep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *HoldoutDetailRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldoutDetailRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldoutDetailRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldoutDetailRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHoldoutAmount

`func (o *HoldoutDetailRep) GetHoldoutAmount() string`

GetHoldoutAmount returns the HoldoutAmount field if non-nil, zero value otherwise.

### GetHoldoutAmountOk

`func (o *HoldoutDetailRep) GetHoldoutAmountOk() (*string, bool)`

GetHoldoutAmountOk returns a tuple with the HoldoutAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutAmount

`func (o *HoldoutDetailRep) SetHoldoutAmount(v string)`

SetHoldoutAmount sets HoldoutAmount field to given value.


### GetIsDirty

`func (o *HoldoutDetailRep) GetIsDirty() bool`

GetIsDirty returns the IsDirty field if non-nil, zero value otherwise.

### GetIsDirtyOk

`func (o *HoldoutDetailRep) GetIsDirtyOk() (*bool, bool)`

GetIsDirtyOk returns a tuple with the IsDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirty

`func (o *HoldoutDetailRep) SetIsDirty(v bool)`

SetIsDirty sets IsDirty field to given value.

### HasIsDirty

`func (o *HoldoutDetailRep) HasIsDirty() bool`

HasIsDirty returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HoldoutDetailRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HoldoutDetailRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HoldoutDetailRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HoldoutDetailRep) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HoldoutDetailRep) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HoldoutDetailRep) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBaseExperiment

`func (o *HoldoutDetailRep) GetBaseExperiment() Experiment`

GetBaseExperiment returns the BaseExperiment field if non-nil, zero value otherwise.

### GetBaseExperimentOk

`func (o *HoldoutDetailRep) GetBaseExperimentOk() (*Experiment, bool)`

GetBaseExperimentOk returns a tuple with the BaseExperiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseExperiment

`func (o *HoldoutDetailRep) SetBaseExperiment(v Experiment)`

SetBaseExperiment sets BaseExperiment field to given value.


### GetRelatedExperiments

`func (o *HoldoutDetailRep) GetRelatedExperiments() []Experiment`

GetRelatedExperiments returns the RelatedExperiments field if non-nil, zero value otherwise.

### GetRelatedExperimentsOk

`func (o *HoldoutDetailRep) GetRelatedExperimentsOk() (*[]Experiment, bool)`

GetRelatedExperimentsOk returns a tuple with the RelatedExperiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedExperiments

`func (o *HoldoutDetailRep) SetRelatedExperiments(v []Experiment)`

SetRelatedExperiments sets RelatedExperiments field to given value.

### HasRelatedExperiments

`func (o *HoldoutDetailRep) HasRelatedExperiments() bool`

HasRelatedExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


