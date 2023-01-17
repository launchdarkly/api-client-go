# ContextAttributeValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind associated with this collection of context attribute values. | 
**Values** | [**[]ContextAttributeValue**](ContextAttributeValue.md) | A collection of context attribute values. | 

## Methods

### NewContextAttributeValues

`func NewContextAttributeValues(kind string, values []ContextAttributeValue, ) *ContextAttributeValues`

NewContextAttributeValues instantiates a new ContextAttributeValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextAttributeValuesWithDefaults

`func NewContextAttributeValuesWithDefaults() *ContextAttributeValues`

NewContextAttributeValuesWithDefaults instantiates a new ContextAttributeValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ContextAttributeValues) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContextAttributeValues) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContextAttributeValues) SetKind(v string)`

SetKind sets Kind field to given value.


### GetValues

`func (o *ContextAttributeValues) GetValues() []ContextAttributeValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ContextAttributeValues) GetValuesOk() (*[]ContextAttributeValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ContextAttributeValues) SetValues(v []ContextAttributeValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


