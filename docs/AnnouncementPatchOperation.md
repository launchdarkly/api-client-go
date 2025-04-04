# AnnouncementPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | The type of operation to perform | 
**Path** | **string** | A JSON Pointer string specifying the part of the document to operate on | 
**Value** | Pointer to **interface{}** | A JSON value used in \&quot;add\&quot;, \&quot;replace\&quot;, and \&quot;test\&quot; operations | [optional] 

## Methods

### NewAnnouncementPatchOperation

`func NewAnnouncementPatchOperation(op string, path string, ) *AnnouncementPatchOperation`

NewAnnouncementPatchOperation instantiates a new AnnouncementPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementPatchOperationWithDefaults

`func NewAnnouncementPatchOperationWithDefaults() *AnnouncementPatchOperation`

NewAnnouncementPatchOperationWithDefaults instantiates a new AnnouncementPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AnnouncementPatchOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AnnouncementPatchOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AnnouncementPatchOperation) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *AnnouncementPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AnnouncementPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AnnouncementPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *AnnouncementPatchOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AnnouncementPatchOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AnnouncementPatchOperation) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *AnnouncementPatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AnnouncementPatchOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AnnouncementPatchOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


