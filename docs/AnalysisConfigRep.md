# AnalysisConfigRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BayesianThreshold** | Pointer to **string** | The threshold for the Probability to Beat Baseline (PBBL) and Probability to Be Best (PBB) comparisons for the Bayesian results analysis approach.  Value should be between 0-100 inclusive. | [optional] 
**SignificanceThreshold** | Pointer to **string** | The significance threshold for the frequentist results analysis approach. Value should be between 0.0-1.0 inclusive. | [optional] 
**TestDirection** | Pointer to **string** | The test sided direction for the frequentist results analysis approach. | [optional] 
**MultipleComparisonCorrectionMethod** | Pointer to **string** | The method for multiple comparison correction. | [optional] 
**MultipleComparisonCorrectionScope** | Pointer to **string** | The scope for multiple comparison correction. | [optional] 
**SequentialTestingEnabled** | Pointer to **bool** | Whether sequential testing is enabled for Frequentist analysis | [optional] 

## Methods

### NewAnalysisConfigRep

`func NewAnalysisConfigRep() *AnalysisConfigRep`

NewAnalysisConfigRep instantiates a new AnalysisConfigRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisConfigRepWithDefaults

`func NewAnalysisConfigRepWithDefaults() *AnalysisConfigRep`

NewAnalysisConfigRepWithDefaults instantiates a new AnalysisConfigRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBayesianThreshold

`func (o *AnalysisConfigRep) GetBayesianThreshold() string`

GetBayesianThreshold returns the BayesianThreshold field if non-nil, zero value otherwise.

### GetBayesianThresholdOk

`func (o *AnalysisConfigRep) GetBayesianThresholdOk() (*string, bool)`

GetBayesianThresholdOk returns a tuple with the BayesianThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBayesianThreshold

`func (o *AnalysisConfigRep) SetBayesianThreshold(v string)`

SetBayesianThreshold sets BayesianThreshold field to given value.

### HasBayesianThreshold

`func (o *AnalysisConfigRep) HasBayesianThreshold() bool`

HasBayesianThreshold returns a boolean if a field has been set.

### GetSignificanceThreshold

`func (o *AnalysisConfigRep) GetSignificanceThreshold() string`

GetSignificanceThreshold returns the SignificanceThreshold field if non-nil, zero value otherwise.

### GetSignificanceThresholdOk

`func (o *AnalysisConfigRep) GetSignificanceThresholdOk() (*string, bool)`

GetSignificanceThresholdOk returns a tuple with the SignificanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceThreshold

`func (o *AnalysisConfigRep) SetSignificanceThreshold(v string)`

SetSignificanceThreshold sets SignificanceThreshold field to given value.

### HasSignificanceThreshold

`func (o *AnalysisConfigRep) HasSignificanceThreshold() bool`

HasSignificanceThreshold returns a boolean if a field has been set.

### GetTestDirection

`func (o *AnalysisConfigRep) GetTestDirection() string`

GetTestDirection returns the TestDirection field if non-nil, zero value otherwise.

### GetTestDirectionOk

`func (o *AnalysisConfigRep) GetTestDirectionOk() (*string, bool)`

GetTestDirectionOk returns a tuple with the TestDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDirection

`func (o *AnalysisConfigRep) SetTestDirection(v string)`

SetTestDirection sets TestDirection field to given value.

### HasTestDirection

`func (o *AnalysisConfigRep) HasTestDirection() bool`

HasTestDirection returns a boolean if a field has been set.

### GetMultipleComparisonCorrectionMethod

`func (o *AnalysisConfigRep) GetMultipleComparisonCorrectionMethod() string`

GetMultipleComparisonCorrectionMethod returns the MultipleComparisonCorrectionMethod field if non-nil, zero value otherwise.

### GetMultipleComparisonCorrectionMethodOk

`func (o *AnalysisConfigRep) GetMultipleComparisonCorrectionMethodOk() (*string, bool)`

GetMultipleComparisonCorrectionMethodOk returns a tuple with the MultipleComparisonCorrectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleComparisonCorrectionMethod

`func (o *AnalysisConfigRep) SetMultipleComparisonCorrectionMethod(v string)`

SetMultipleComparisonCorrectionMethod sets MultipleComparisonCorrectionMethod field to given value.

### HasMultipleComparisonCorrectionMethod

`func (o *AnalysisConfigRep) HasMultipleComparisonCorrectionMethod() bool`

HasMultipleComparisonCorrectionMethod returns a boolean if a field has been set.

### GetMultipleComparisonCorrectionScope

`func (o *AnalysisConfigRep) GetMultipleComparisonCorrectionScope() string`

GetMultipleComparisonCorrectionScope returns the MultipleComparisonCorrectionScope field if non-nil, zero value otherwise.

### GetMultipleComparisonCorrectionScopeOk

`func (o *AnalysisConfigRep) GetMultipleComparisonCorrectionScopeOk() (*string, bool)`

GetMultipleComparisonCorrectionScopeOk returns a tuple with the MultipleComparisonCorrectionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleComparisonCorrectionScope

`func (o *AnalysisConfigRep) SetMultipleComparisonCorrectionScope(v string)`

SetMultipleComparisonCorrectionScope sets MultipleComparisonCorrectionScope field to given value.

### HasMultipleComparisonCorrectionScope

`func (o *AnalysisConfigRep) HasMultipleComparisonCorrectionScope() bool`

HasMultipleComparisonCorrectionScope returns a boolean if a field has been set.

### GetSequentialTestingEnabled

`func (o *AnalysisConfigRep) GetSequentialTestingEnabled() bool`

GetSequentialTestingEnabled returns the SequentialTestingEnabled field if non-nil, zero value otherwise.

### GetSequentialTestingEnabledOk

`func (o *AnalysisConfigRep) GetSequentialTestingEnabledOk() (*bool, bool)`

GetSequentialTestingEnabledOk returns a tuple with the SequentialTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialTestingEnabled

`func (o *AnalysisConfigRep) SetSequentialTestingEnabled(v bool)`

SetSequentialTestingEnabled sets SequentialTestingEnabled field to given value.

### HasSequentialTestingEnabled

`func (o *AnalysisConfigRep) HasSequentialTestingEnabled() bool`

HasSequentialTestingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


