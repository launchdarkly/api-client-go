# PatchWithComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patch** | Pointer to [**[]JSONPatchElt**](JSONPatchElt.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchWithComment

`func NewPatchWithComment() *PatchWithComment`

NewPatchWithComment instantiates a new PatchWithComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchWithCommentWithDefaults

`func NewPatchWithCommentWithDefaults() *PatchWithComment`

NewPatchWithCommentWithDefaults instantiates a new PatchWithComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatch

`func (o *PatchWithComment) GetPatch() []JSONPatchElt`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *PatchWithComment) GetPatchOk() (*[]JSONPatchElt, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *PatchWithComment) SetPatch(v []JSONPatchElt)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *PatchWithComment) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetComment

`func (o *PatchWithComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatchWithComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatchWithComment) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatchWithComment) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


