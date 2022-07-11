# ParameterDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** | The default value for the given parameter | [optional] 
**BooleanVariationValue** | Pointer to **bool** | Variation value for boolean flags. Not applicable for non-boolean flags. | [optional] 
**RuleClause** | Pointer to [**RuleClause**](RuleClause.md) |  | [optional] 

## Methods

### NewParameterDefault

`func NewParameterDefault() *ParameterDefault`

NewParameterDefault instantiates a new ParameterDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterDefaultWithDefaults

`func NewParameterDefaultWithDefaults() *ParameterDefault`

NewParameterDefaultWithDefaults instantiates a new ParameterDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ParameterDefault) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParameterDefault) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParameterDefault) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ParameterDefault) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ParameterDefault) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ParameterDefault) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetBooleanVariationValue

`func (o *ParameterDefault) GetBooleanVariationValue() bool`

GetBooleanVariationValue returns the BooleanVariationValue field if non-nil, zero value otherwise.

### GetBooleanVariationValueOk

`func (o *ParameterDefault) GetBooleanVariationValueOk() (*bool, bool)`

GetBooleanVariationValueOk returns a tuple with the BooleanVariationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanVariationValue

`func (o *ParameterDefault) SetBooleanVariationValue(v bool)`

SetBooleanVariationValue sets BooleanVariationValue field to given value.

### HasBooleanVariationValue

`func (o *ParameterDefault) HasBooleanVariationValue() bool`

HasBooleanVariationValue returns a boolean if a field has been set.

### GetRuleClause

`func (o *ParameterDefault) GetRuleClause() RuleClause`

GetRuleClause returns the RuleClause field if non-nil, zero value otherwise.

### GetRuleClauseOk

`func (o *ParameterDefault) GetRuleClauseOk() (*RuleClause, bool)`

GetRuleClauseOk returns a tuple with the RuleClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleClause

`func (o *ParameterDefault) SetRuleClause(v RuleClause)`

SetRuleClause sets RuleClause field to given value.

### HasRuleClause

`func (o *ParameterDefault) HasRuleClause() bool`

HasRuleClause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


