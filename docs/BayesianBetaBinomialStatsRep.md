# BayesianBetaBinomialStatsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorAlpha** | Pointer to **float32** | Sum of converted pseudo-units for prior distribution | [optional] 
**PriorBeta** | Pointer to **float32** | Sum of non-converted pseudo-units for prior distribution | [optional] 
**PriorMean** | Pointer to **float32** | Mean of the prior distribution | [optional] 
**DataWeight** | Pointer to **float32** | The precision weight of the data mean | [optional] 

## Methods

### NewBayesianBetaBinomialStatsRep

`func NewBayesianBetaBinomialStatsRep() *BayesianBetaBinomialStatsRep`

NewBayesianBetaBinomialStatsRep instantiates a new BayesianBetaBinomialStatsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBayesianBetaBinomialStatsRepWithDefaults

`func NewBayesianBetaBinomialStatsRepWithDefaults() *BayesianBetaBinomialStatsRep`

NewBayesianBetaBinomialStatsRepWithDefaults instantiates a new BayesianBetaBinomialStatsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorAlpha

`func (o *BayesianBetaBinomialStatsRep) GetPriorAlpha() float32`

GetPriorAlpha returns the PriorAlpha field if non-nil, zero value otherwise.

### GetPriorAlphaOk

`func (o *BayesianBetaBinomialStatsRep) GetPriorAlphaOk() (*float32, bool)`

GetPriorAlphaOk returns a tuple with the PriorAlpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAlpha

`func (o *BayesianBetaBinomialStatsRep) SetPriorAlpha(v float32)`

SetPriorAlpha sets PriorAlpha field to given value.

### HasPriorAlpha

`func (o *BayesianBetaBinomialStatsRep) HasPriorAlpha() bool`

HasPriorAlpha returns a boolean if a field has been set.

### GetPriorBeta

`func (o *BayesianBetaBinomialStatsRep) GetPriorBeta() float32`

GetPriorBeta returns the PriorBeta field if non-nil, zero value otherwise.

### GetPriorBetaOk

`func (o *BayesianBetaBinomialStatsRep) GetPriorBetaOk() (*float32, bool)`

GetPriorBetaOk returns a tuple with the PriorBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorBeta

`func (o *BayesianBetaBinomialStatsRep) SetPriorBeta(v float32)`

SetPriorBeta sets PriorBeta field to given value.

### HasPriorBeta

`func (o *BayesianBetaBinomialStatsRep) HasPriorBeta() bool`

HasPriorBeta returns a boolean if a field has been set.

### GetPriorMean

`func (o *BayesianBetaBinomialStatsRep) GetPriorMean() float32`

GetPriorMean returns the PriorMean field if non-nil, zero value otherwise.

### GetPriorMeanOk

`func (o *BayesianBetaBinomialStatsRep) GetPriorMeanOk() (*float32, bool)`

GetPriorMeanOk returns a tuple with the PriorMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorMean

`func (o *BayesianBetaBinomialStatsRep) SetPriorMean(v float32)`

SetPriorMean sets PriorMean field to given value.

### HasPriorMean

`func (o *BayesianBetaBinomialStatsRep) HasPriorMean() bool`

HasPriorMean returns a boolean if a field has been set.

### GetDataWeight

`func (o *BayesianBetaBinomialStatsRep) GetDataWeight() float32`

GetDataWeight returns the DataWeight field if non-nil, zero value otherwise.

### GetDataWeightOk

`func (o *BayesianBetaBinomialStatsRep) GetDataWeightOk() (*float32, bool)`

GetDataWeightOk returns a tuple with the DataWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWeight

`func (o *BayesianBetaBinomialStatsRep) SetDataWeight(v float32)`

SetDataWeight sets DataWeight field to given value.

### HasDataWeight

`func (o *BayesianBetaBinomialStatsRep) HasDataWeight() bool`

HasDataWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


