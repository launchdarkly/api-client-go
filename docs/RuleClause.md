# RuleClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | The attribute the rule applies to, for example, last name or email address | [optional] 
**Op** | Pointer to **string** | The operator to apply to the given attribute | [optional] 
**Negate** | Pointer to **bool** | Whether the operator should be negated | [optional] 

## Methods

### NewRuleClause

`func NewRuleClause() *RuleClause`

NewRuleClause instantiates a new RuleClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleClauseWithDefaults

`func NewRuleClauseWithDefaults() *RuleClause`

NewRuleClauseWithDefaults instantiates a new RuleClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *RuleClause) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *RuleClause) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *RuleClause) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *RuleClause) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOp

`func (o *RuleClause) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RuleClause) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RuleClause) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *RuleClause) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetNegate

`func (o *RuleClause) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *RuleClause) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *RuleClause) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *RuleClause) HasNegate() bool`

HasNegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


