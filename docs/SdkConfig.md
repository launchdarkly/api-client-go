# SdkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamplingRatio** | Pointer to **int32** |  | [optional] 
**ExcludeFromSummaries** | Pointer to **bool** |  | [optional] 

## Methods

### NewSdkConfig

`func NewSdkConfig() *SdkConfig`

NewSdkConfig instantiates a new SdkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkConfigWithDefaults

`func NewSdkConfigWithDefaults() *SdkConfig`

NewSdkConfigWithDefaults instantiates a new SdkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamplingRatio

`func (o *SdkConfig) GetSamplingRatio() int32`

GetSamplingRatio returns the SamplingRatio field if non-nil, zero value otherwise.

### GetSamplingRatioOk

`func (o *SdkConfig) GetSamplingRatioOk() (*int32, bool)`

GetSamplingRatioOk returns a tuple with the SamplingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRatio

`func (o *SdkConfig) SetSamplingRatio(v int32)`

SetSamplingRatio sets SamplingRatio field to given value.

### HasSamplingRatio

`func (o *SdkConfig) HasSamplingRatio() bool`

HasSamplingRatio returns a boolean if a field has been set.

### GetExcludeFromSummaries

`func (o *SdkConfig) GetExcludeFromSummaries() bool`

GetExcludeFromSummaries returns the ExcludeFromSummaries field if non-nil, zero value otherwise.

### GetExcludeFromSummariesOk

`func (o *SdkConfig) GetExcludeFromSummariesOk() (*bool, bool)`

GetExcludeFromSummariesOk returns a tuple with the ExcludeFromSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSummaries

`func (o *SdkConfig) SetExcludeFromSummaries(v bool)`

SetExcludeFromSummaries sets ExcludeFromSummaries field to given value.

### HasExcludeFromSummaries

`func (o *SdkConfig) HasExcludeFromSummaries() bool`

HasExcludeFromSummaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


