# ReferenceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**Hunks** | Pointer to [**[]HunkRep**](HunkRep.md) |  | [optional] 

## Methods

### NewReferenceRep

`func NewReferenceRep() *ReferenceRep`

NewReferenceRep instantiates a new ReferenceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceRepWithDefaults

`func NewReferenceRepWithDefaults() *ReferenceRep`

NewReferenceRepWithDefaults instantiates a new ReferenceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ReferenceRep) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReferenceRep) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReferenceRep) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReferenceRep) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHint

`func (o *ReferenceRep) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ReferenceRep) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ReferenceRep) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ReferenceRep) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetHunks

`func (o *ReferenceRep) GetHunks() []HunkRep`

GetHunks returns the Hunks field if non-nil, zero value otherwise.

### GetHunksOk

`func (o *ReferenceRep) GetHunksOk() (*[]HunkRep, bool)`

GetHunksOk returns a tuple with the Hunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunks

`func (o *ReferenceRep) SetHunks(v []HunkRep)`

SetHunks sets Hunks field to given value.

### HasHunks

`func (o *ReferenceRep) HasHunks() bool`

HasHunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


