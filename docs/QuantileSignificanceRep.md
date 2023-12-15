# QuantileSignificanceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Significant** | Pointer to **bool** | Whether the treatment is significant against the control | [optional] 
**Regression** | Pointer to **bool** | Whether the treatment is a regression against the control | [optional] 
**FromTreatmentId** | Pointer to **string** | The treatment ID | [optional] 

## Methods

### NewQuantileSignificanceRep

`func NewQuantileSignificanceRep() *QuantileSignificanceRep`

NewQuantileSignificanceRep instantiates a new QuantileSignificanceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantileSignificanceRepWithDefaults

`func NewQuantileSignificanceRepWithDefaults() *QuantileSignificanceRep`

NewQuantileSignificanceRepWithDefaults instantiates a new QuantileSignificanceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignificant

`func (o *QuantileSignificanceRep) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *QuantileSignificanceRep) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *QuantileSignificanceRep) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.

### HasSignificant

`func (o *QuantileSignificanceRep) HasSignificant() bool`

HasSignificant returns a boolean if a field has been set.

### GetRegression

`func (o *QuantileSignificanceRep) GetRegression() bool`

GetRegression returns the Regression field if non-nil, zero value otherwise.

### GetRegressionOk

`func (o *QuantileSignificanceRep) GetRegressionOk() (*bool, bool)`

GetRegressionOk returns a tuple with the Regression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegression

`func (o *QuantileSignificanceRep) SetRegression(v bool)`

SetRegression sets Regression field to given value.

### HasRegression

`func (o *QuantileSignificanceRep) HasRegression() bool`

HasRegression returns a boolean if a field has been set.

### GetFromTreatmentId

`func (o *QuantileSignificanceRep) GetFromTreatmentId() string`

GetFromTreatmentId returns the FromTreatmentId field if non-nil, zero value otherwise.

### GetFromTreatmentIdOk

`func (o *QuantileSignificanceRep) GetFromTreatmentIdOk() (*string, bool)`

GetFromTreatmentIdOk returns a tuple with the FromTreatmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTreatmentId

`func (o *QuantileSignificanceRep) SetFromTreatmentId(v string)`

SetFromTreatmentId sets FromTreatmentId field to given value.

### HasFromTreatmentId

`func (o *QuantileSignificanceRep) HasFromTreatmentId() bool`

HasFromTreatmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


