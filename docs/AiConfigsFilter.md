# AiConfigsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Filter type. One of [contextAttribute, eventProperty, group] | 
**Attribute** | Pointer to **string** | If not a group node, the context attribute name or event property name to filter on | [optional] 
**Op** | **string** |  | 
**Values** | **[]interface{}** | The context attribute / event property values or group member nodes | 
**ContextKind** | Pointer to **string** | For context attribute filters, the context kind. | [optional] 
**Negate** | **bool** | If set, then take the inverse of the operator. &#39;in&#39; becomes &#39;not in&#39;. | 

## Methods

### NewAiConfigsFilter

`func NewAiConfigsFilter(type_ string, op string, values []interface{}, negate bool, ) *AiConfigsFilter`

NewAiConfigsFilter instantiates a new AiConfigsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsFilterWithDefaults

`func NewAiConfigsFilterWithDefaults() *AiConfigsFilter`

NewAiConfigsFilterWithDefaults instantiates a new AiConfigsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiConfigsFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiConfigsFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiConfigsFilter) SetType(v string)`

SetType sets Type field to given value.


### GetAttribute

`func (o *AiConfigsFilter) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AiConfigsFilter) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AiConfigsFilter) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *AiConfigsFilter) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOp

`func (o *AiConfigsFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AiConfigsFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AiConfigsFilter) SetOp(v string)`

SetOp sets Op field to given value.


### GetValues

`func (o *AiConfigsFilter) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AiConfigsFilter) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AiConfigsFilter) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetContextKind

`func (o *AiConfigsFilter) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *AiConfigsFilter) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *AiConfigsFilter) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *AiConfigsFilter) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.

### GetNegate

`func (o *AiConfigsFilter) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AiConfigsFilter) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AiConfigsFilter) SetNegate(v bool)`

SetNegate sets Negate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


