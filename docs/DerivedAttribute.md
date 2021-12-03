# DerivedAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** |  | [optional] 
**LastDerived** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDerivedAttribute

`func NewDerivedAttribute() *DerivedAttribute`

NewDerivedAttribute instantiates a new DerivedAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDerivedAttributeWithDefaults

`func NewDerivedAttributeWithDefaults() *DerivedAttribute`

NewDerivedAttributeWithDefaults instantiates a new DerivedAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DerivedAttribute) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DerivedAttribute) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DerivedAttribute) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DerivedAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DerivedAttribute) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DerivedAttribute) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetLastDerived

`func (o *DerivedAttribute) GetLastDerived() time.Time`

GetLastDerived returns the LastDerived field if non-nil, zero value otherwise.

### GetLastDerivedOk

`func (o *DerivedAttribute) GetLastDerivedOk() (*time.Time, bool)`

GetLastDerivedOk returns a tuple with the LastDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDerived

`func (o *DerivedAttribute) SetLastDerived(v time.Time)`

SetLastDerived sets LastDerived field to given value.

### HasLastDerived

`func (o *DerivedAttribute) HasLastDerived() bool`

HasLastDerived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


