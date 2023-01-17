# SegmentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the segment | 
**Key** | **string** | A unique key used to reference the segment | 
**Description** | Pointer to **string** | A description of the segment&#39;s purpose | [optional] 
**Tags** | Pointer to **[]string** | Tags for the segment | [optional] 
**Unbounded** | Pointer to **bool** | Whether to create a standard segment (false) or a Big Segment (true). Only use a Big Segment if you need to add more than 15,000 users. | [optional] 
**UnboundedContextKind** | Pointer to **string** | If unbounded is true, you can use this field to set the Big Segment&#39;s context kind | [optional] 

## Methods

### NewSegmentBody

`func NewSegmentBody(name string, key string, ) *SegmentBody`

NewSegmentBody instantiates a new SegmentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentBodyWithDefaults

`func NewSegmentBodyWithDefaults() *SegmentBody`

NewSegmentBodyWithDefaults instantiates a new SegmentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentBody) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *SegmentBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SegmentBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SegmentBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *SegmentBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *SegmentBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SegmentBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SegmentBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SegmentBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnbounded

`func (o *SegmentBody) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *SegmentBody) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *SegmentBody) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.

### HasUnbounded

`func (o *SegmentBody) HasUnbounded() bool`

HasUnbounded returns a boolean if a field has been set.

### GetUnboundedContextKind

`func (o *SegmentBody) GetUnboundedContextKind() string`

GetUnboundedContextKind returns the UnboundedContextKind field if non-nil, zero value otherwise.

### GetUnboundedContextKindOk

`func (o *SegmentBody) GetUnboundedContextKindOk() (*string, bool)`

GetUnboundedContextKindOk returns a tuple with the UnboundedContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnboundedContextKind

`func (o *SegmentBody) SetUnboundedContextKind(v string)`

SetUnboundedContextKind sets UnboundedContextKind field to given value.

### HasUnboundedContextKind

`func (o *SegmentBody) HasUnboundedContextKind() bool`

HasUnboundedContextKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


