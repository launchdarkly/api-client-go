# JudgeAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JudgeConfigKey** | **string** | Key of the judge AI config | 
**SamplingRate** | **float32** | Sampling rate for this judge attachment (0.0 to 1.0) | 

## Methods

### NewJudgeAttachment

`func NewJudgeAttachment(judgeConfigKey string, samplingRate float32, ) *JudgeAttachment`

NewJudgeAttachment instantiates a new JudgeAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeAttachmentWithDefaults

`func NewJudgeAttachmentWithDefaults() *JudgeAttachment`

NewJudgeAttachmentWithDefaults instantiates a new JudgeAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJudgeConfigKey

`func (o *JudgeAttachment) GetJudgeConfigKey() string`

GetJudgeConfigKey returns the JudgeConfigKey field if non-nil, zero value otherwise.

### GetJudgeConfigKeyOk

`func (o *JudgeAttachment) GetJudgeConfigKeyOk() (*string, bool)`

GetJudgeConfigKeyOk returns a tuple with the JudgeConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeConfigKey

`func (o *JudgeAttachment) SetJudgeConfigKey(v string)`

SetJudgeConfigKey sets JudgeConfigKey field to given value.


### GetSamplingRate

`func (o *JudgeAttachment) GetSamplingRate() float32`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *JudgeAttachment) GetSamplingRateOk() (*float32, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *JudgeAttachment) SetSamplingRate(v float32)`

SetSamplingRate sets SamplingRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


