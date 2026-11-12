# IterationAssignmentDataSourceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the assignment data source used for this iteration | 
**Name** | **string** | The display name of the assignment data source | 
**ExperimentId** | **string** | The experiment identifier in the assignment data source | 
**StartTime** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewIterationAssignmentDataSourceRep

`func NewIterationAssignmentDataSourceRep(key string, name string, experimentId string, ) *IterationAssignmentDataSourceRep`

NewIterationAssignmentDataSourceRep instantiates a new IterationAssignmentDataSourceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationAssignmentDataSourceRepWithDefaults

`func NewIterationAssignmentDataSourceRepWithDefaults() *IterationAssignmentDataSourceRep`

NewIterationAssignmentDataSourceRepWithDefaults instantiates a new IterationAssignmentDataSourceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *IterationAssignmentDataSourceRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IterationAssignmentDataSourceRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IterationAssignmentDataSourceRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *IterationAssignmentDataSourceRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IterationAssignmentDataSourceRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IterationAssignmentDataSourceRep) SetName(v string)`

SetName sets Name field to given value.


### GetExperimentId

`func (o *IterationAssignmentDataSourceRep) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *IterationAssignmentDataSourceRep) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *IterationAssignmentDataSourceRep) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.


### GetStartTime

`func (o *IterationAssignmentDataSourceRep) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IterationAssignmentDataSourceRep) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IterationAssignmentDataSourceRep) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IterationAssignmentDataSourceRep) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *IterationAssignmentDataSourceRep) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *IterationAssignmentDataSourceRep) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *IterationAssignmentDataSourceRep) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *IterationAssignmentDataSourceRep) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


