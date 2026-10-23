# EnvironmentReleaseClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextKind** | Pointer to **string** | Context kind the clause applies to. Omit for &#x60;segmentMatch&#x60;. | [optional] 
**Attribute** | **string** |  | 
**Op** | **string** | One of &#x60;in&#x60;, &#x60;endsWith&#x60;, &#x60;startsWith&#x60;, &#x60;matches&#x60;, &#x60;contains&#x60;, &#x60;lessThan&#x60;, &#x60;lessThanOrEqual&#x60;, &#x60;greaterThan&#x60;, &#x60;greaterThanOrEqual&#x60;, &#x60;before&#x60;, &#x60;after&#x60;, &#x60;segmentMatch&#x60;, &#x60;semVerEqual&#x60;, &#x60;semVerLessThan&#x60;, &#x60;semVerGreaterThan&#x60;.  | 
**Values** | **[]interface{}** | JSON primitives only (strings, numbers, booleans); objects, arrays, and null are rejected. | 
**Negate** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEnvironmentReleaseClause

`func NewEnvironmentReleaseClause(attribute string, op string, values []interface{}, ) *EnvironmentReleaseClause`

NewEnvironmentReleaseClause instantiates a new EnvironmentReleaseClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleaseClauseWithDefaults

`func NewEnvironmentReleaseClauseWithDefaults() *EnvironmentReleaseClause`

NewEnvironmentReleaseClauseWithDefaults instantiates a new EnvironmentReleaseClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextKind

`func (o *EnvironmentReleaseClause) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *EnvironmentReleaseClause) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *EnvironmentReleaseClause) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *EnvironmentReleaseClause) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.

### GetAttribute

`func (o *EnvironmentReleaseClause) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *EnvironmentReleaseClause) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *EnvironmentReleaseClause) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetOp

`func (o *EnvironmentReleaseClause) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EnvironmentReleaseClause) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EnvironmentReleaseClause) SetOp(v string)`

SetOp sets Op field to given value.


### GetValues

`func (o *EnvironmentReleaseClause) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EnvironmentReleaseClause) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EnvironmentReleaseClause) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetNegate

`func (o *EnvironmentReleaseClause) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *EnvironmentReleaseClause) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *EnvironmentReleaseClause) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *EnvironmentReleaseClause) HasNegate() bool`

HasNegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


