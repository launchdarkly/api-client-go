# RuleRepClauses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Op** | **string** |  | 
**Values** | **[]interface{}** |  | 
**Negate** | **bool** |  | 

## Methods

### NewRuleRepClauses

`func NewRuleRepClauses(op string, values []interface{}, negate bool, ) *RuleRepClauses`

NewRuleRepClauses instantiates a new RuleRepClauses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleRepClausesWithDefaults

`func NewRuleRepClausesWithDefaults() *RuleRepClauses`

NewRuleRepClausesWithDefaults instantiates a new RuleRepClauses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleRepClauses) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleRepClauses) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleRepClauses) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleRepClauses) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttribute

`func (o *RuleRepClauses) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *RuleRepClauses) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *RuleRepClauses) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *RuleRepClauses) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOp

`func (o *RuleRepClauses) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RuleRepClauses) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RuleRepClauses) SetOp(v string)`

SetOp sets Op field to given value.


### GetValues

`func (o *RuleRepClauses) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RuleRepClauses) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RuleRepClauses) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetNegate

`func (o *RuleRepClauses) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *RuleRepClauses) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *RuleRepClauses) SetNegate(v bool)`

SetNegate sets Negate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


