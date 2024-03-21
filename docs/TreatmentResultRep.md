# TreatmentResultRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TreatmentId** | Pointer to **string** | The ID of the treatment | [optional] 
**TreatmentName** | Pointer to **string** | The name of the treatment | [optional] 
**Mean** | Pointer to **float32** | The average value of the variation in this sample. It doesn’t capture the uncertainty in the measurement, so it should not be the only measurement you use to make decisions. | [optional] 
**CredibleInterval** | Pointer to [**CredibleIntervalRep**](CredibleIntervalRep.md) |  | [optional] 
**PBest** | Pointer to **float32** | The likelihood that this variation has the biggest effect on the primary metric. The variation with the highest probability is likely the best of the variations you&#39;re testing | [optional] 
**RelativeDifferences** | Pointer to [**[]RelativeDifferenceRep**](RelativeDifferenceRep.md) | Estimates of the relative difference between this treatment&#39;s mean and the mean of each other treatment | [optional] 
**Units** | Pointer to **int64** | The number of units exposed to this treatment that have event values, including those that are configured to default to 0 | [optional] 
**Traffic** | Pointer to **int64** | The number of units exposed to this treatment. | [optional] 
**Distribution** | Pointer to [**Distribution**](Distribution.md) |  | [optional] 

## Methods

### NewTreatmentResultRep

`func NewTreatmentResultRep() *TreatmentResultRep`

NewTreatmentResultRep instantiates a new TreatmentResultRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreatmentResultRepWithDefaults

`func NewTreatmentResultRepWithDefaults() *TreatmentResultRep`

NewTreatmentResultRepWithDefaults instantiates a new TreatmentResultRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTreatmentId

`func (o *TreatmentResultRep) GetTreatmentId() string`

GetTreatmentId returns the TreatmentId field if non-nil, zero value otherwise.

### GetTreatmentIdOk

`func (o *TreatmentResultRep) GetTreatmentIdOk() (*string, bool)`

GetTreatmentIdOk returns a tuple with the TreatmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentId

`func (o *TreatmentResultRep) SetTreatmentId(v string)`

SetTreatmentId sets TreatmentId field to given value.

### HasTreatmentId

`func (o *TreatmentResultRep) HasTreatmentId() bool`

HasTreatmentId returns a boolean if a field has been set.

### GetTreatmentName

`func (o *TreatmentResultRep) GetTreatmentName() string`

GetTreatmentName returns the TreatmentName field if non-nil, zero value otherwise.

### GetTreatmentNameOk

`func (o *TreatmentResultRep) GetTreatmentNameOk() (*string, bool)`

GetTreatmentNameOk returns a tuple with the TreatmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentName

`func (o *TreatmentResultRep) SetTreatmentName(v string)`

SetTreatmentName sets TreatmentName field to given value.

### HasTreatmentName

`func (o *TreatmentResultRep) HasTreatmentName() bool`

HasTreatmentName returns a boolean if a field has been set.

### GetMean

`func (o *TreatmentResultRep) GetMean() float32`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *TreatmentResultRep) GetMeanOk() (*float32, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *TreatmentResultRep) SetMean(v float32)`

SetMean sets Mean field to given value.

### HasMean

`func (o *TreatmentResultRep) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetCredibleInterval

`func (o *TreatmentResultRep) GetCredibleInterval() CredibleIntervalRep`

GetCredibleInterval returns the CredibleInterval field if non-nil, zero value otherwise.

### GetCredibleIntervalOk

`func (o *TreatmentResultRep) GetCredibleIntervalOk() (*CredibleIntervalRep, bool)`

GetCredibleIntervalOk returns a tuple with the CredibleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleInterval

`func (o *TreatmentResultRep) SetCredibleInterval(v CredibleIntervalRep)`

SetCredibleInterval sets CredibleInterval field to given value.

### HasCredibleInterval

`func (o *TreatmentResultRep) HasCredibleInterval() bool`

HasCredibleInterval returns a boolean if a field has been set.

### GetPBest

`func (o *TreatmentResultRep) GetPBest() float32`

GetPBest returns the PBest field if non-nil, zero value otherwise.

### GetPBestOk

`func (o *TreatmentResultRep) GetPBestOk() (*float32, bool)`

GetPBestOk returns a tuple with the PBest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPBest

`func (o *TreatmentResultRep) SetPBest(v float32)`

SetPBest sets PBest field to given value.

### HasPBest

`func (o *TreatmentResultRep) HasPBest() bool`

HasPBest returns a boolean if a field has been set.

### GetRelativeDifferences

`func (o *TreatmentResultRep) GetRelativeDifferences() []RelativeDifferenceRep`

GetRelativeDifferences returns the RelativeDifferences field if non-nil, zero value otherwise.

### GetRelativeDifferencesOk

`func (o *TreatmentResultRep) GetRelativeDifferencesOk() (*[]RelativeDifferenceRep, bool)`

GetRelativeDifferencesOk returns a tuple with the RelativeDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDifferences

`func (o *TreatmentResultRep) SetRelativeDifferences(v []RelativeDifferenceRep)`

SetRelativeDifferences sets RelativeDifferences field to given value.

### HasRelativeDifferences

`func (o *TreatmentResultRep) HasRelativeDifferences() bool`

HasRelativeDifferences returns a boolean if a field has been set.

### GetUnits

`func (o *TreatmentResultRep) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TreatmentResultRep) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TreatmentResultRep) SetUnits(v int64)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *TreatmentResultRep) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetTraffic

`func (o *TreatmentResultRep) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *TreatmentResultRep) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *TreatmentResultRep) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *TreatmentResultRep) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetDistribution

`func (o *TreatmentResultRep) GetDistribution() Distribution`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *TreatmentResultRep) GetDistributionOk() (*Distribution, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *TreatmentResultRep) SetDistribution(v Distribution)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *TreatmentResultRep) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


