# BayesianNormalStatsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataWeight** | Pointer to **float32** | The precision weight of the data mean | [optional] 
**PriorMean** | Pointer to **float32** | Mean of the prior distribution | [optional] 

## Methods

### NewBayesianNormalStatsRep

`func NewBayesianNormalStatsRep() *BayesianNormalStatsRep`

NewBayesianNormalStatsRep instantiates a new BayesianNormalStatsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBayesianNormalStatsRepWithDefaults

`func NewBayesianNormalStatsRepWithDefaults() *BayesianNormalStatsRep`

NewBayesianNormalStatsRepWithDefaults instantiates a new BayesianNormalStatsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataWeight

`func (o *BayesianNormalStatsRep) GetDataWeight() float32`

GetDataWeight returns the DataWeight field if non-nil, zero value otherwise.

### GetDataWeightOk

`func (o *BayesianNormalStatsRep) GetDataWeightOk() (*float32, bool)`

GetDataWeightOk returns a tuple with the DataWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWeight

`func (o *BayesianNormalStatsRep) SetDataWeight(v float32)`

SetDataWeight sets DataWeight field to given value.

### HasDataWeight

`func (o *BayesianNormalStatsRep) HasDataWeight() bool`

HasDataWeight returns a boolean if a field has been set.

### GetPriorMean

`func (o *BayesianNormalStatsRep) GetPriorMean() float32`

GetPriorMean returns the PriorMean field if non-nil, zero value otherwise.

### GetPriorMeanOk

`func (o *BayesianNormalStatsRep) GetPriorMeanOk() (*float32, bool)`

GetPriorMeanOk returns a tuple with the PriorMean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorMean

`func (o *BayesianNormalStatsRep) SetPriorMean(v float32)`

SetPriorMean sets PriorMean field to given value.

### HasPriorMean

`func (o *BayesianNormalStatsRep) HasPriorMean() bool`

HasPriorMean returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


