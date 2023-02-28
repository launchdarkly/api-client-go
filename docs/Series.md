# Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int64** | The timestamp | 
**Value** | **float32** | The value for the given timestamp | 

## Methods

### NewSeries

`func NewSeries(time int64, value float32, ) *Series`

NewSeries instantiates a new Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesWithDefaults

`func NewSeriesWithDefaults() *Series`

NewSeriesWithDefaults instantiates a new Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *Series) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Series) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Series) SetTime(v int64)`

SetTime sets Time field to given value.


### GetValue

`func (o *Series) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Series) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Series) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


