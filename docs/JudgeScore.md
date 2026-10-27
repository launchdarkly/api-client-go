# JudgeScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** | The judge&#39;s evaluation metric key | 
**Score** | **float32** | Average judge score across the recorded samples | 
**Count** | **int32** | Number of judge scores recorded | 

## Methods

### NewJudgeScore

`func NewJudgeScore(metricKey string, score float32, count int32, ) *JudgeScore`

NewJudgeScore instantiates a new JudgeScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeScoreWithDefaults

`func NewJudgeScoreWithDefaults() *JudgeScore`

NewJudgeScoreWithDefaults instantiates a new JudgeScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *JudgeScore) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *JudgeScore) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *JudgeScore) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetScore

`func (o *JudgeScore) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *JudgeScore) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *JudgeScore) SetScore(v float32)`

SetScore sets Score field to given value.


### GetCount

`func (o *JudgeScore) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *JudgeScore) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *JudgeScore) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


