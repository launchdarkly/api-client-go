# VariationEvalSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** | The variation value | [optional] 
**Before** | Pointer to **int64** | The number of evaluations in the ten minutes before the flag event | [optional] 
**After** | Pointer to **int64** | The number of evaluations in the ten minutes after the flag event | [optional] 

## Methods

### NewVariationEvalSummary

`func NewVariationEvalSummary() *VariationEvalSummary`

NewVariationEvalSummary instantiates a new VariationEvalSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationEvalSummaryWithDefaults

`func NewVariationEvalSummaryWithDefaults() *VariationEvalSummary`

NewVariationEvalSummaryWithDefaults instantiates a new VariationEvalSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VariationEvalSummary) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariationEvalSummary) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariationEvalSummary) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *VariationEvalSummary) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariationEvalSummary) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariationEvalSummary) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetBefore

`func (o *VariationEvalSummary) GetBefore() int64`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *VariationEvalSummary) GetBeforeOk() (*int64, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *VariationEvalSummary) SetBefore(v int64)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *VariationEvalSummary) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetAfter

`func (o *VariationEvalSummary) GetAfter() int64`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *VariationEvalSummary) GetAfterOk() (*int64, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *VariationEvalSummary) SetAfter(v int64)`

SetAfter sets After field to given value.

### HasAfter

`func (o *VariationEvalSummary) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


