# AssignmentDataSourceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the assignment data source to use for this iteration | 
**ExperimentId** | Pointer to **string** | The experiment identifier in the assignment data source. You can omit this while the experiment is a draft, but it must be set before you start the experiment. | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAssignmentDataSourceInput

`func NewAssignmentDataSourceInput(key string, ) *AssignmentDataSourceInput`

NewAssignmentDataSourceInput instantiates a new AssignmentDataSourceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentDataSourceInputWithDefaults

`func NewAssignmentDataSourceInputWithDefaults() *AssignmentDataSourceInput`

NewAssignmentDataSourceInputWithDefaults instantiates a new AssignmentDataSourceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AssignmentDataSourceInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AssignmentDataSourceInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AssignmentDataSourceInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetExperimentId

`func (o *AssignmentDataSourceInput) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *AssignmentDataSourceInput) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *AssignmentDataSourceInput) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *AssignmentDataSourceInput) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetStartTime

`func (o *AssignmentDataSourceInput) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AssignmentDataSourceInput) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AssignmentDataSourceInput) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AssignmentDataSourceInput) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *AssignmentDataSourceInput) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AssignmentDataSourceInput) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AssignmentDataSourceInput) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AssignmentDataSourceInput) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


