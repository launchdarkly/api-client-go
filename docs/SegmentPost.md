# SegmentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the segment | 
**Key** | **string** | A unique key used to reference the segment | 
**Description** | Pointer to **string** | A description of the segment&#39;s purpose | [optional] 
**Tags** | Pointer to **[]string** | Tags for the segment | [optional] 
**Unbounded** | Pointer to **bool** |  | [optional] 

## Methods

### NewSegmentPost

`func NewSegmentPost(name string, key string, ) *SegmentPost`

NewSegmentPost instantiates a new SegmentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentPostWithDefaults

`func NewSegmentPostWithDefaults() *SegmentPost`

NewSegmentPostWithDefaults instantiates a new SegmentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentPost) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *SegmentPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SegmentPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SegmentPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *SegmentPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *SegmentPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SegmentPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SegmentPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SegmentPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnbounded

`func (o *SegmentPost) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *SegmentPost) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *SegmentPost) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.

### HasUnbounded

`func (o *SegmentPost) HasUnbounded() bool`

HasUnbounded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


