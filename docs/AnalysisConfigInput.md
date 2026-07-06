# AnalysisConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BayesianThreshold** | Pointer to **string** | The threshold for the Probability to Beat Baseline (PBBL) and Probability to Be Best (PBB) comparisons for the Bayesian results analysis approach. | [optional] 
**SignificanceThreshold** | Pointer to **string** | The significance threshold for the frequentist results analysis approach. | [optional] 
**TestDirection** | Pointer to **string** | The test sided direction for the frequentist results analysis approach. | [optional] 
**MultipleComparisonCorrectionMethod** | Pointer to **string** | The method to use for multiple comparison correction. | [optional] 
**MultipleComparisonCorrectionScope** | Pointer to **string** | The scope of the multiple comparison correction. | [optional] 
**SequentialTestingEnabled** | Pointer to **bool** | Whether sequential testing is enabled for Frequentist analysis | [optional] 

## Methods

### NewAnalysisConfigInput

`func NewAnalysisConfigInput() *AnalysisConfigInput`

NewAnalysisConfigInput instantiates a new AnalysisConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisConfigInputWithDefaults

`func NewAnalysisConfigInputWithDefaults() *AnalysisConfigInput`

NewAnalysisConfigInputWithDefaults instantiates a new AnalysisConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBayesianThreshold

`func (o *AnalysisConfigInput) GetBayesianThreshold() string`

GetBayesianThreshold returns the BayesianThreshold field if non-nil, zero value otherwise.

### GetBayesianThresholdOk

`func (o *AnalysisConfigInput) GetBayesianThresholdOk() (*string, bool)`

GetBayesianThresholdOk returns a tuple with the BayesianThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBayesianThreshold

`func (o *AnalysisConfigInput) SetBayesianThreshold(v string)`

SetBayesianThreshold sets BayesianThreshold field to given value.

### HasBayesianThreshold

`func (o *AnalysisConfigInput) HasBayesianThreshold() bool`

HasBayesianThreshold returns a boolean if a field has been set.

### GetSignificanceThreshold

`func (o *AnalysisConfigInput) GetSignificanceThreshold() string`

GetSignificanceThreshold returns the SignificanceThreshold field if non-nil, zero value otherwise.

### GetSignificanceThresholdOk

`func (o *AnalysisConfigInput) GetSignificanceThresholdOk() (*string, bool)`

GetSignificanceThresholdOk returns a tuple with the SignificanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceThreshold

`func (o *AnalysisConfigInput) SetSignificanceThreshold(v string)`

SetSignificanceThreshold sets SignificanceThreshold field to given value.

### HasSignificanceThreshold

`func (o *AnalysisConfigInput) HasSignificanceThreshold() bool`

HasSignificanceThreshold returns a boolean if a field has been set.

### GetTestDirection

`func (o *AnalysisConfigInput) GetTestDirection() string`

GetTestDirection returns the TestDirection field if non-nil, zero value otherwise.

### GetTestDirectionOk

`func (o *AnalysisConfigInput) GetTestDirectionOk() (*string, bool)`

GetTestDirectionOk returns a tuple with the TestDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDirection

`func (o *AnalysisConfigInput) SetTestDirection(v string)`

SetTestDirection sets TestDirection field to given value.

### HasTestDirection

`func (o *AnalysisConfigInput) HasTestDirection() bool`

HasTestDirection returns a boolean if a field has been set.

### GetMultipleComparisonCorrectionMethod

`func (o *AnalysisConfigInput) GetMultipleComparisonCorrectionMethod() string`

GetMultipleComparisonCorrectionMethod returns the MultipleComparisonCorrectionMethod field if non-nil, zero value otherwise.

### GetMultipleComparisonCorrectionMethodOk

`func (o *AnalysisConfigInput) GetMultipleComparisonCorrectionMethodOk() (*string, bool)`

GetMultipleComparisonCorrectionMethodOk returns a tuple with the MultipleComparisonCorrectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleComparisonCorrectionMethod

`func (o *AnalysisConfigInput) SetMultipleComparisonCorrectionMethod(v string)`

SetMultipleComparisonCorrectionMethod sets MultipleComparisonCorrectionMethod field to given value.

### HasMultipleComparisonCorrectionMethod

`func (o *AnalysisConfigInput) HasMultipleComparisonCorrectionMethod() bool`

HasMultipleComparisonCorrectionMethod returns a boolean if a field has been set.

### GetMultipleComparisonCorrectionScope

`func (o *AnalysisConfigInput) GetMultipleComparisonCorrectionScope() string`

GetMultipleComparisonCorrectionScope returns the MultipleComparisonCorrectionScope field if non-nil, zero value otherwise.

### GetMultipleComparisonCorrectionScopeOk

`func (o *AnalysisConfigInput) GetMultipleComparisonCorrectionScopeOk() (*string, bool)`

GetMultipleComparisonCorrectionScopeOk returns a tuple with the MultipleComparisonCorrectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleComparisonCorrectionScope

`func (o *AnalysisConfigInput) SetMultipleComparisonCorrectionScope(v string)`

SetMultipleComparisonCorrectionScope sets MultipleComparisonCorrectionScope field to given value.

### HasMultipleComparisonCorrectionScope

`func (o *AnalysisConfigInput) HasMultipleComparisonCorrectionScope() bool`

HasMultipleComparisonCorrectionScope returns a boolean if a field has been set.

### GetSequentialTestingEnabled

`func (o *AnalysisConfigInput) GetSequentialTestingEnabled() bool`

GetSequentialTestingEnabled returns the SequentialTestingEnabled field if non-nil, zero value otherwise.

### GetSequentialTestingEnabledOk

`func (o *AnalysisConfigInput) GetSequentialTestingEnabledOk() (*bool, bool)`

GetSequentialTestingEnabledOk returns a tuple with the SequentialTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialTestingEnabled

`func (o *AnalysisConfigInput) SetSequentialTestingEnabled(v bool)`

SetSequentialTestingEnabled sets SequentialTestingEnabled field to given value.

### HasSequentialTestingEnabled

`func (o *AnalysisConfigInput) HasSequentialTestingEnabled() bool`

HasSequentialTestingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


