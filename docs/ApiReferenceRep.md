# ApiReferenceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**Hunks** | Pointer to [**[]ApiHunkRep**](ApiHunkRep.md) |  | [optional] 

## Methods

### NewApiReferenceRep

`func NewApiReferenceRep() *ApiReferenceRep`

NewApiReferenceRep instantiates a new ApiReferenceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReferenceRepWithDefaults

`func NewApiReferenceRepWithDefaults() *ApiReferenceRep`

NewApiReferenceRepWithDefaults instantiates a new ApiReferenceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ApiReferenceRep) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiReferenceRep) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiReferenceRep) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiReferenceRep) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHint

`func (o *ApiReferenceRep) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ApiReferenceRep) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ApiReferenceRep) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ApiReferenceRep) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetHunks

`func (o *ApiReferenceRep) GetHunks() []ApiHunkRep`

GetHunks returns the Hunks field if non-nil, zero value otherwise.

### GetHunksOk

`func (o *ApiReferenceRep) GetHunksOk() (*[]ApiHunkRep, bool)`

GetHunksOk returns a tuple with the Hunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunks

`func (o *ApiReferenceRep) SetHunks(v []ApiHunkRep)`

SetHunks sets Hunks field to given value.

### HasHunks

`func (o *ApiReferenceRep) HasHunks() bool`

HasHunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


