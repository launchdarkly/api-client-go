# JSONPatchElt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | The type of operation to perform | 
**Path** | **string** | A JSON Pointer string specifying the part of the document to operate on | 
**Value** | **interface{}** | A JSON value used in \&quot;add\&quot;, \&quot;replace\&quot;, and \&quot;test\&quot; operations | 

## Methods

### NewJSONPatchElt

`func NewJSONPatchElt(op string, path string, value interface{}, ) *JSONPatchElt`

NewJSONPatchElt instantiates a new JSONPatchElt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchEltWithDefaults

`func NewJSONPatchEltWithDefaults() *JSONPatchElt`

NewJSONPatchEltWithDefaults instantiates a new JSONPatchElt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *JSONPatchElt) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JSONPatchElt) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JSONPatchElt) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *JSONPatchElt) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JSONPatchElt) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JSONPatchElt) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JSONPatchElt) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPatchElt) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPatchElt) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JSONPatchElt) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JSONPatchElt) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


