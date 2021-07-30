# CoderefsBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**UpdateSequenceId** | Pointer to **int64** |  | [optional] 
**SyncTime** | Pointer to **int64** |  | [optional] 
**References** | Pointer to [**[]CoderefsReferenceFile**](CoderefsReferenceFile.md) |  | [optional] 

## Methods

### NewCoderefsBranch

`func NewCoderefsBranch() *CoderefsBranch`

NewCoderefsBranch instantiates a new CoderefsBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoderefsBranchWithDefaults

`func NewCoderefsBranchWithDefaults() *CoderefsBranch`

NewCoderefsBranchWithDefaults instantiates a new CoderefsBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CoderefsBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoderefsBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoderefsBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoderefsBranch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha

`func (o *CoderefsBranch) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CoderefsBranch) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CoderefsBranch) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CoderefsBranch) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetUpdateSequenceId

`func (o *CoderefsBranch) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *CoderefsBranch) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *CoderefsBranch) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.

### HasUpdateSequenceId

`func (o *CoderefsBranch) HasUpdateSequenceId() bool`

HasUpdateSequenceId returns a boolean if a field has been set.

### GetSyncTime

`func (o *CoderefsBranch) GetSyncTime() int64`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *CoderefsBranch) GetSyncTimeOk() (*int64, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *CoderefsBranch) SetSyncTime(v int64)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *CoderefsBranch) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetReferences

`func (o *CoderefsBranch) GetReferences() []CoderefsReferenceFile`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CoderefsBranch) GetReferencesOk() (*[]CoderefsReferenceFile, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CoderefsBranch) SetReferences(v []CoderefsReferenceFile)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CoderefsBranch) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


