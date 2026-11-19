# EnvironmentReleasePlanClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The clause&#39;s server-generated identifier | [optional] [readonly] 
**ContextKind** | Pointer to **string** | Context kind this clause applies to | [optional] [default to "user"]
**Attribute** | **string** | Context attribute to match on (may be empty for segmentMatch) | 
**Op** | **string** | Clause operator | 
**Values** | **[]interface{}** | Values to match against (strings, numbers, or booleans). Each string value can be at most 1,024 bytes when UTF-8 encoded, so a value with multibyte characters can be rejected even if it has fewer than 1,024 characters. | 
**Negate** | Pointer to **bool** | Whether to negate the clause | [optional] [default to false]

## Methods

### NewEnvironmentReleasePlanClause

`func NewEnvironmentReleasePlanClause(attribute string, op string, values []interface{}, ) *EnvironmentReleasePlanClause`

NewEnvironmentReleasePlanClause instantiates a new EnvironmentReleasePlanClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanClauseWithDefaults

`func NewEnvironmentReleasePlanClauseWithDefaults() *EnvironmentReleasePlanClause`

NewEnvironmentReleasePlanClauseWithDefaults instantiates a new EnvironmentReleasePlanClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReleasePlanClause) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReleasePlanClause) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReleasePlanClause) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentReleasePlanClause) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContextKind

`func (o *EnvironmentReleasePlanClause) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *EnvironmentReleasePlanClause) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *EnvironmentReleasePlanClause) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *EnvironmentReleasePlanClause) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.

### GetAttribute

`func (o *EnvironmentReleasePlanClause) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *EnvironmentReleasePlanClause) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *EnvironmentReleasePlanClause) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetOp

`func (o *EnvironmentReleasePlanClause) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EnvironmentReleasePlanClause) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EnvironmentReleasePlanClause) SetOp(v string)`

SetOp sets Op field to given value.


### GetValues

`func (o *EnvironmentReleasePlanClause) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EnvironmentReleasePlanClause) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EnvironmentReleasePlanClause) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetNegate

`func (o *EnvironmentReleasePlanClause) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *EnvironmentReleasePlanClause) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *EnvironmentReleasePlanClause) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *EnvironmentReleasePlanClause) HasNegate() bool`

HasNegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


