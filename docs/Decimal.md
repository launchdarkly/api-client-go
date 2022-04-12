# Decimal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** |  | [optional] 
**Exp** | Pointer to **int32** |  | [optional] 

## Methods

### NewDecimal

`func NewDecimal() *Decimal`

NewDecimal instantiates a new Decimal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecimalWithDefaults

`func NewDecimalWithDefaults() *Decimal`

NewDecimalWithDefaults instantiates a new Decimal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Decimal) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Decimal) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Decimal) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Decimal) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Decimal) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Decimal) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetExp

`func (o *Decimal) GetExp() int32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *Decimal) GetExpOk() (*int32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *Decimal) SetExp(v int32)`

SetExp sets Exp field to given value.

### HasExp

`func (o *Decimal) HasExp() bool`

HasExp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


