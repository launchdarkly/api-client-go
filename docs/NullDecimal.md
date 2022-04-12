# NullDecimal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimal** | Pointer to [**Decimal**](Decimal.md) |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewNullDecimal

`func NewNullDecimal() *NullDecimal`

NewNullDecimal instantiates a new NullDecimal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullDecimalWithDefaults

`func NewNullDecimalWithDefaults() *NullDecimal`

NewNullDecimalWithDefaults instantiates a new NullDecimal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimal

`func (o *NullDecimal) GetDecimal() Decimal`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *NullDecimal) GetDecimalOk() (*Decimal, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *NullDecimal) SetDecimal(v Decimal)`

SetDecimal sets Decimal field to given value.

### HasDecimal

`func (o *NullDecimal) HasDecimal() bool`

HasDecimal returns a boolean if a field has been set.

### GetValid

`func (o *NullDecimal) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *NullDecimal) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *NullDecimal) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *NullDecimal) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


