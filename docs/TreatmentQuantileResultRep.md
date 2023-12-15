# TreatmentQuantileResultRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TreatmentId** | Pointer to **string** | The ID of the treatment | [optional] 
**TreatmentName** | Pointer to **string** | The name of the treatment | [optional] 
**PercentileValue** | Pointer to **float32** | The value at the specified percentile. | [optional] 
**PercentileInterval** | Pointer to [**CredibleIntervalRep**](CredibleIntervalRep.md) |  | [optional] 
**Significance** | Pointer to [**QuantileSignificanceRep**](QuantileSignificanceRep.md) |  | [optional] 
**RelativeDifferences** | Pointer to [**[]RelativeDifferenceRep**](RelativeDifferenceRep.md) | Estimates of the relative difference between this treatment&#39;s percentile and the percentile of each other treatment | [optional] 
**Units** | Pointer to **int64** | The number of experiment users for this variation | [optional] 

## Methods

### NewTreatmentQuantileResultRep

`func NewTreatmentQuantileResultRep() *TreatmentQuantileResultRep`

NewTreatmentQuantileResultRep instantiates a new TreatmentQuantileResultRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreatmentQuantileResultRepWithDefaults

`func NewTreatmentQuantileResultRepWithDefaults() *TreatmentQuantileResultRep`

NewTreatmentQuantileResultRepWithDefaults instantiates a new TreatmentQuantileResultRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTreatmentId

`func (o *TreatmentQuantileResultRep) GetTreatmentId() string`

GetTreatmentId returns the TreatmentId field if non-nil, zero value otherwise.

### GetTreatmentIdOk

`func (o *TreatmentQuantileResultRep) GetTreatmentIdOk() (*string, bool)`

GetTreatmentIdOk returns a tuple with the TreatmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentId

`func (o *TreatmentQuantileResultRep) SetTreatmentId(v string)`

SetTreatmentId sets TreatmentId field to given value.

### HasTreatmentId

`func (o *TreatmentQuantileResultRep) HasTreatmentId() bool`

HasTreatmentId returns a boolean if a field has been set.

### GetTreatmentName

`func (o *TreatmentQuantileResultRep) GetTreatmentName() string`

GetTreatmentName returns the TreatmentName field if non-nil, zero value otherwise.

### GetTreatmentNameOk

`func (o *TreatmentQuantileResultRep) GetTreatmentNameOk() (*string, bool)`

GetTreatmentNameOk returns a tuple with the TreatmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatmentName

`func (o *TreatmentQuantileResultRep) SetTreatmentName(v string)`

SetTreatmentName sets TreatmentName field to given value.

### HasTreatmentName

`func (o *TreatmentQuantileResultRep) HasTreatmentName() bool`

HasTreatmentName returns a boolean if a field has been set.

### GetPercentileValue

`func (o *TreatmentQuantileResultRep) GetPercentileValue() float32`

GetPercentileValue returns the PercentileValue field if non-nil, zero value otherwise.

### GetPercentileValueOk

`func (o *TreatmentQuantileResultRep) GetPercentileValueOk() (*float32, bool)`

GetPercentileValueOk returns a tuple with the PercentileValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileValue

`func (o *TreatmentQuantileResultRep) SetPercentileValue(v float32)`

SetPercentileValue sets PercentileValue field to given value.

### HasPercentileValue

`func (o *TreatmentQuantileResultRep) HasPercentileValue() bool`

HasPercentileValue returns a boolean if a field has been set.

### GetPercentileInterval

`func (o *TreatmentQuantileResultRep) GetPercentileInterval() CredibleIntervalRep`

GetPercentileInterval returns the PercentileInterval field if non-nil, zero value otherwise.

### GetPercentileIntervalOk

`func (o *TreatmentQuantileResultRep) GetPercentileIntervalOk() (*CredibleIntervalRep, bool)`

GetPercentileIntervalOk returns a tuple with the PercentileInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileInterval

`func (o *TreatmentQuantileResultRep) SetPercentileInterval(v CredibleIntervalRep)`

SetPercentileInterval sets PercentileInterval field to given value.

### HasPercentileInterval

`func (o *TreatmentQuantileResultRep) HasPercentileInterval() bool`

HasPercentileInterval returns a boolean if a field has been set.

### GetSignificance

`func (o *TreatmentQuantileResultRep) GetSignificance() QuantileSignificanceRep`

GetSignificance returns the Significance field if non-nil, zero value otherwise.

### GetSignificanceOk

`func (o *TreatmentQuantileResultRep) GetSignificanceOk() (*QuantileSignificanceRep, bool)`

GetSignificanceOk returns a tuple with the Significance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificance

`func (o *TreatmentQuantileResultRep) SetSignificance(v QuantileSignificanceRep)`

SetSignificance sets Significance field to given value.

### HasSignificance

`func (o *TreatmentQuantileResultRep) HasSignificance() bool`

HasSignificance returns a boolean if a field has been set.

### GetRelativeDifferences

`func (o *TreatmentQuantileResultRep) GetRelativeDifferences() []RelativeDifferenceRep`

GetRelativeDifferences returns the RelativeDifferences field if non-nil, zero value otherwise.

### GetRelativeDifferencesOk

`func (o *TreatmentQuantileResultRep) GetRelativeDifferencesOk() (*[]RelativeDifferenceRep, bool)`

GetRelativeDifferencesOk returns a tuple with the RelativeDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDifferences

`func (o *TreatmentQuantileResultRep) SetRelativeDifferences(v []RelativeDifferenceRep)`

SetRelativeDifferences sets RelativeDifferences field to given value.

### HasRelativeDifferences

`func (o *TreatmentQuantileResultRep) HasRelativeDifferences() bool`

HasRelativeDifferences returns a boolean if a field has been set.

### GetUnits

`func (o *TreatmentQuantileResultRep) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TreatmentQuantileResultRep) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TreatmentQuantileResultRep) SetUnits(v int64)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *TreatmentQuantileResultRep) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


