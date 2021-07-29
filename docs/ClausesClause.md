# ClausesClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Op** | **string** |  | 
**Values** | **[]interface{}** |  | 
**Negate** | **bool** |  | 

## Methods

### NewClausesClause

`func NewClausesClause(op string, values []interface{}, negate bool, ) *ClausesClause`

NewClausesClause instantiates a new ClausesClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClausesClauseWithDefaults

`func NewClausesClauseWithDefaults() *ClausesClause`

NewClausesClauseWithDefaults instantiates a new ClausesClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClausesClause) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClausesClause) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClausesClause) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClausesClause) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttribute

`func (o *ClausesClause) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ClausesClause) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ClausesClause) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *ClausesClause) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOp

`func (o *ClausesClause) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ClausesClause) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ClausesClause) SetOp(v string)`

SetOp sets Op field to given value.


### GetValues

`func (o *ClausesClause) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ClausesClause) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ClausesClause) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetNegate

`func (o *ClausesClause) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ClausesClause) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ClausesClause) SetNegate(v bool)`

SetNegate sets Negate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


