# ExperimentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The experiment name | 
**Description** | Pointer to **string** | The experiment description | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this experiment | [optional] 
**Key** | **string** | The experiment key | 
**Iteration** | [**IterationInput**](IterationInput.md) |  | 

## Methods

### NewExperimentPost

`func NewExperimentPost(name string, key string, iteration IterationInput, ) *ExperimentPost`

NewExperimentPost instantiates a new ExperimentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentPostWithDefaults

`func NewExperimentPostWithDefaults() *ExperimentPost`

NewExperimentPostWithDefaults instantiates a new ExperimentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExperimentPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExperimentPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperimentPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainerId

`func (o *ExperimentPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ExperimentPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ExperimentPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ExperimentPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetKey

`func (o *ExperimentPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetIteration

`func (o *ExperimentPost) GetIteration() IterationInput`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *ExperimentPost) GetIterationOk() (*IterationInput, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *ExperimentPost) SetIteration(v IterationInput)`

SetIteration sets Iteration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


