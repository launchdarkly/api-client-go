# PutBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The branch name | 
**Head** | **string** | An ID representing the branch HEAD. For example, a commit SHA. | 
**UpdateSequenceId** | Pointer to **int64** | An optional ID used to prevent older data from overwriting newer data. If no sequence ID is included, the newly submitted data will always be saved. | [optional] 
**SyncTime** | **int64** |  | 
**References** | Pointer to [**[]ReferenceRep**](ReferenceRep.md) | An array of flag references found on the branch | [optional] 

## Methods

### NewPutBranch

`func NewPutBranch(name string, head string, syncTime int64, ) *PutBranch`

NewPutBranch instantiates a new PutBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBranchWithDefaults

`func NewPutBranchWithDefaults() *PutBranch`

NewPutBranchWithDefaults instantiates a new PutBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutBranch) SetName(v string)`

SetName sets Name field to given value.


### GetHead

`func (o *PutBranch) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PutBranch) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PutBranch) SetHead(v string)`

SetHead sets Head field to given value.


### GetUpdateSequenceId

`func (o *PutBranch) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *PutBranch) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *PutBranch) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.

### HasUpdateSequenceId

`func (o *PutBranch) HasUpdateSequenceId() bool`

HasUpdateSequenceId returns a boolean if a field has been set.

### GetSyncTime

`func (o *PutBranch) GetSyncTime() int64`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *PutBranch) GetSyncTimeOk() (*int64, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *PutBranch) SetSyncTime(v int64)`

SetSyncTime sets SyncTime field to given value.


### GetReferences

`func (o *PutBranch) GetReferences() []ReferenceRep`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PutBranch) GetReferencesOk() (*[]ReferenceRep, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PutBranch) SetReferences(v []ReferenceRep)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PutBranch) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


