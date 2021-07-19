# ExperimentTotalsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CumulativeValue** | Pointer to **float32** |  | [optional] 
**CumulativeCount** | Pointer to **int64** |  | [optional] 
**CumulativeImpressionCount** | Pointer to **int64** |  | [optional] 
**CumulativeConversionRate** | Pointer to **float32** |  | [optional] 
**CumulativeConfidenceInterval** | Pointer to [**ConfidenceIntervalRep**](ConfidenceIntervalRep.md) |  | [optional] 
**PValue** | Pointer to **float32** |  | [optional] 
**Improvement** | Pointer to **float32** |  | [optional] 
**MinSampleSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewExperimentTotalsRep

`func NewExperimentTotalsRep() *ExperimentTotalsRep`

NewExperimentTotalsRep instantiates a new ExperimentTotalsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentTotalsRepWithDefaults

`func NewExperimentTotalsRepWithDefaults() *ExperimentTotalsRep`

NewExperimentTotalsRepWithDefaults instantiates a new ExperimentTotalsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCumulativeValue

`func (o *ExperimentTotalsRep) GetCumulativeValue() float32`

GetCumulativeValue returns the CumulativeValue field if non-nil, zero value otherwise.

### GetCumulativeValueOk

`func (o *ExperimentTotalsRep) GetCumulativeValueOk() (*float32, bool)`

GetCumulativeValueOk returns a tuple with the CumulativeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeValue

`func (o *ExperimentTotalsRep) SetCumulativeValue(v float32)`

SetCumulativeValue sets CumulativeValue field to given value.

### HasCumulativeValue

`func (o *ExperimentTotalsRep) HasCumulativeValue() bool`

HasCumulativeValue returns a boolean if a field has been set.

### GetCumulativeCount

`func (o *ExperimentTotalsRep) GetCumulativeCount() int64`

GetCumulativeCount returns the CumulativeCount field if non-nil, zero value otherwise.

### GetCumulativeCountOk

`func (o *ExperimentTotalsRep) GetCumulativeCountOk() (*int64, bool)`

GetCumulativeCountOk returns a tuple with the CumulativeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCount

`func (o *ExperimentTotalsRep) SetCumulativeCount(v int64)`

SetCumulativeCount sets CumulativeCount field to given value.

### HasCumulativeCount

`func (o *ExperimentTotalsRep) HasCumulativeCount() bool`

HasCumulativeCount returns a boolean if a field has been set.

### GetCumulativeImpressionCount

`func (o *ExperimentTotalsRep) GetCumulativeImpressionCount() int64`

GetCumulativeImpressionCount returns the CumulativeImpressionCount field if non-nil, zero value otherwise.

### GetCumulativeImpressionCountOk

`func (o *ExperimentTotalsRep) GetCumulativeImpressionCountOk() (*int64, bool)`

GetCumulativeImpressionCountOk returns a tuple with the CumulativeImpressionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeImpressionCount

`func (o *ExperimentTotalsRep) SetCumulativeImpressionCount(v int64)`

SetCumulativeImpressionCount sets CumulativeImpressionCount field to given value.

### HasCumulativeImpressionCount

`func (o *ExperimentTotalsRep) HasCumulativeImpressionCount() bool`

HasCumulativeImpressionCount returns a boolean if a field has been set.

### GetCumulativeConversionRate

`func (o *ExperimentTotalsRep) GetCumulativeConversionRate() float32`

GetCumulativeConversionRate returns the CumulativeConversionRate field if non-nil, zero value otherwise.

### GetCumulativeConversionRateOk

`func (o *ExperimentTotalsRep) GetCumulativeConversionRateOk() (*float32, bool)`

GetCumulativeConversionRateOk returns a tuple with the CumulativeConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeConversionRate

`func (o *ExperimentTotalsRep) SetCumulativeConversionRate(v float32)`

SetCumulativeConversionRate sets CumulativeConversionRate field to given value.

### HasCumulativeConversionRate

`func (o *ExperimentTotalsRep) HasCumulativeConversionRate() bool`

HasCumulativeConversionRate returns a boolean if a field has been set.

### GetCumulativeConfidenceInterval

`func (o *ExperimentTotalsRep) GetCumulativeConfidenceInterval() ConfidenceIntervalRep`

GetCumulativeConfidenceInterval returns the CumulativeConfidenceInterval field if non-nil, zero value otherwise.

### GetCumulativeConfidenceIntervalOk

`func (o *ExperimentTotalsRep) GetCumulativeConfidenceIntervalOk() (*ConfidenceIntervalRep, bool)`

GetCumulativeConfidenceIntervalOk returns a tuple with the CumulativeConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeConfidenceInterval

`func (o *ExperimentTotalsRep) SetCumulativeConfidenceInterval(v ConfidenceIntervalRep)`

SetCumulativeConfidenceInterval sets CumulativeConfidenceInterval field to given value.

### HasCumulativeConfidenceInterval

`func (o *ExperimentTotalsRep) HasCumulativeConfidenceInterval() bool`

HasCumulativeConfidenceInterval returns a boolean if a field has been set.

### GetPValue

`func (o *ExperimentTotalsRep) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ExperimentTotalsRep) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ExperimentTotalsRep) SetPValue(v float32)`

SetPValue sets PValue field to given value.

### HasPValue

`func (o *ExperimentTotalsRep) HasPValue() bool`

HasPValue returns a boolean if a field has been set.

### GetImprovement

`func (o *ExperimentTotalsRep) GetImprovement() float32`

GetImprovement returns the Improvement field if non-nil, zero value otherwise.

### GetImprovementOk

`func (o *ExperimentTotalsRep) GetImprovementOk() (*float32, bool)`

GetImprovementOk returns a tuple with the Improvement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImprovement

`func (o *ExperimentTotalsRep) SetImprovement(v float32)`

SetImprovement sets Improvement field to given value.

### HasImprovement

`func (o *ExperimentTotalsRep) HasImprovement() bool`

HasImprovement returns a boolean if a field has been set.

### GetMinSampleSize

`func (o *ExperimentTotalsRep) GetMinSampleSize() int64`

GetMinSampleSize returns the MinSampleSize field if non-nil, zero value otherwise.

### GetMinSampleSizeOk

`func (o *ExperimentTotalsRep) GetMinSampleSizeOk() (*int64, bool)`

GetMinSampleSizeOk returns a tuple with the MinSampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSampleSize

`func (o *ExperimentTotalsRep) SetMinSampleSize(v int64)`

SetMinSampleSize sets MinSampleSize field to given value.

### HasMinSampleSize

`func (o *ExperimentTotalsRep) HasMinSampleSize() bool`

HasMinSampleSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


