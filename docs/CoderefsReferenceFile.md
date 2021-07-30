# CoderefsReferenceFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**Hunks** | Pointer to [**[]CoderefsHunk**](CoderefsHunk.md) |  | [optional] 

## Methods

### NewCoderefsReferenceFile

`func NewCoderefsReferenceFile() *CoderefsReferenceFile`

NewCoderefsReferenceFile instantiates a new CoderefsReferenceFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoderefsReferenceFileWithDefaults

`func NewCoderefsReferenceFileWithDefaults() *CoderefsReferenceFile`

NewCoderefsReferenceFileWithDefaults instantiates a new CoderefsReferenceFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *CoderefsReferenceFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CoderefsReferenceFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CoderefsReferenceFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CoderefsReferenceFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHint

`func (o *CoderefsReferenceFile) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CoderefsReferenceFile) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CoderefsReferenceFile) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *CoderefsReferenceFile) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetHunks

`func (o *CoderefsReferenceFile) GetHunks() []CoderefsHunk`

GetHunks returns the Hunks field if non-nil, zero value otherwise.

### GetHunksOk

`func (o *CoderefsReferenceFile) GetHunksOk() (*[]CoderefsHunk, bool)`

GetHunksOk returns a tuple with the Hunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunks

`func (o *CoderefsReferenceFile) SetHunks(v []CoderefsHunk)`

SetHunks sets Hunks field to given value.

### HasHunks

`func (o *CoderefsReferenceFile) HasHunks() bool`

HasHunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


