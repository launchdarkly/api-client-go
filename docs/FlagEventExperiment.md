# FlagEventExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The experiment key | 
**Name** | **string** | The experiment name | 
**Iteration** | [**FlagEventExperimentIteration**](FlagEventExperimentIteration.md) |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewFlagEventExperiment

`func NewFlagEventExperiment(key string, name string, iteration FlagEventExperimentIteration, ) *FlagEventExperiment`

NewFlagEventExperiment instantiates a new FlagEventExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagEventExperimentWithDefaults

`func NewFlagEventExperimentWithDefaults() *FlagEventExperiment`

NewFlagEventExperimentWithDefaults instantiates a new FlagEventExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlagEventExperiment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagEventExperiment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagEventExperiment) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *FlagEventExperiment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagEventExperiment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagEventExperiment) SetName(v string)`

SetName sets Name field to given value.


### GetIteration

`func (o *FlagEventExperiment) GetIteration() FlagEventExperimentIteration`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *FlagEventExperiment) GetIterationOk() (*FlagEventExperimentIteration, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *FlagEventExperiment) SetIteration(v FlagEventExperimentIteration)`

SetIteration sets Iteration field to given value.


### GetLinks

`func (o *FlagEventExperiment) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagEventExperiment) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagEventExperiment) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagEventExperiment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


