# Experiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**MaintainerId** | **string** |  | 
**CreationDate** | **int64** |  | 
**Links** | [**map[string]Link**](Link.md) |  | 
**CurrentIteration** | Pointer to [**IterationRep**](IterationRep.md) |  | [optional] 
**DraftIteration** | Pointer to [**IterationRep**](IterationRep.md) |  | [optional] 
**PreviousIterations** | Pointer to [**[]IterationRep**](IterationRep.md) |  | [optional] 

## Methods

### NewExperiment

`func NewExperiment(key string, name string, maintainerId string, creationDate int64, links map[string]Link, ) *Experiment`

NewExperiment instantiates a new Experiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentWithDefaults

`func NewExperimentWithDefaults() *Experiment`

NewExperimentWithDefaults instantiates a new Experiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Experiment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Experiment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Experiment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Experiment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Experiment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Experiment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Experiment) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Experiment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Experiment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Experiment) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Experiment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Experiment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Experiment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Experiment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainerId

`func (o *Experiment) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *Experiment) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *Experiment) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.


### GetCreationDate

`func (o *Experiment) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Experiment) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Experiment) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLinks

`func (o *Experiment) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Experiment) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Experiment) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetCurrentIteration

`func (o *Experiment) GetCurrentIteration() IterationRep`

GetCurrentIteration returns the CurrentIteration field if non-nil, zero value otherwise.

### GetCurrentIterationOk

`func (o *Experiment) GetCurrentIterationOk() (*IterationRep, bool)`

GetCurrentIterationOk returns a tuple with the CurrentIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIteration

`func (o *Experiment) SetCurrentIteration(v IterationRep)`

SetCurrentIteration sets CurrentIteration field to given value.

### HasCurrentIteration

`func (o *Experiment) HasCurrentIteration() bool`

HasCurrentIteration returns a boolean if a field has been set.

### GetDraftIteration

`func (o *Experiment) GetDraftIteration() IterationRep`

GetDraftIteration returns the DraftIteration field if non-nil, zero value otherwise.

### GetDraftIterationOk

`func (o *Experiment) GetDraftIterationOk() (*IterationRep, bool)`

GetDraftIterationOk returns a tuple with the DraftIteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftIteration

`func (o *Experiment) SetDraftIteration(v IterationRep)`

SetDraftIteration sets DraftIteration field to given value.

### HasDraftIteration

`func (o *Experiment) HasDraftIteration() bool`

HasDraftIteration returns a boolean if a field has been set.

### GetPreviousIterations

`func (o *Experiment) GetPreviousIterations() []IterationRep`

GetPreviousIterations returns the PreviousIterations field if non-nil, zero value otherwise.

### GetPreviousIterationsOk

`func (o *Experiment) GetPreviousIterationsOk() (*[]IterationRep, bool)`

GetPreviousIterationsOk returns a tuple with the PreviousIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIterations

`func (o *Experiment) SetPreviousIterations(v []IterationRep)`

SetPreviousIterations sets PreviousIterations field to given value.

### HasPreviousIterations

`func (o *Experiment) HasPreviousIterations() bool`

HasPreviousIterations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


