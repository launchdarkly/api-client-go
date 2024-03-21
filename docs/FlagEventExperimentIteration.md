# FlagEventExperimentIteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The experiment iteration ID | 
**Status** | **string** |  | 
**StartedAt** | **int64** |  | 
**EndedAt** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewFlagEventExperimentIteration

`func NewFlagEventExperimentIteration(id string, status string, startedAt int64, ) *FlagEventExperimentIteration`

NewFlagEventExperimentIteration instantiates a new FlagEventExperimentIteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagEventExperimentIterationWithDefaults

`func NewFlagEventExperimentIterationWithDefaults() *FlagEventExperimentIteration`

NewFlagEventExperimentIterationWithDefaults instantiates a new FlagEventExperimentIteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlagEventExperimentIteration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagEventExperimentIteration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagEventExperimentIteration) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *FlagEventExperimentIteration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlagEventExperimentIteration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlagEventExperimentIteration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *FlagEventExperimentIteration) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlagEventExperimentIteration) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlagEventExperimentIteration) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *FlagEventExperimentIteration) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *FlagEventExperimentIteration) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *FlagEventExperimentIteration) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *FlagEventExperimentIteration) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetLinks

`func (o *FlagEventExperimentIteration) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagEventExperimentIteration) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagEventExperimentIteration) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagEventExperimentIteration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


