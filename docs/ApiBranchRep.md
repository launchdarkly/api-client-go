# ApiBranchRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Head** | Pointer to **string** |  | [optional] 
**UpdateSequenceId** | Pointer to **int64** |  | [optional] 
**SyncTime** | Pointer to **int64** |  | [optional] 
**References** | Pointer to [**[]ApiReferenceRep**](ApiReferenceRep.md) |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiBranchRep

`func NewApiBranchRep() *ApiBranchRep`

NewApiBranchRep instantiates a new ApiBranchRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBranchRepWithDefaults

`func NewApiBranchRepWithDefaults() *ApiBranchRep`

NewApiBranchRepWithDefaults instantiates a new ApiBranchRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiBranchRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBranchRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBranchRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiBranchRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHead

`func (o *ApiBranchRep) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *ApiBranchRep) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *ApiBranchRep) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *ApiBranchRep) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetUpdateSequenceId

`func (o *ApiBranchRep) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *ApiBranchRep) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *ApiBranchRep) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.

### HasUpdateSequenceId

`func (o *ApiBranchRep) HasUpdateSequenceId() bool`

HasUpdateSequenceId returns a boolean if a field has been set.

### GetSyncTime

`func (o *ApiBranchRep) GetSyncTime() int64`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *ApiBranchRep) GetSyncTimeOk() (*int64, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *ApiBranchRep) SetSyncTime(v int64)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *ApiBranchRep) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetReferences

`func (o *ApiBranchRep) GetReferences() []ApiReferenceRep`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiBranchRep) GetReferencesOk() (*[]ApiReferenceRep, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiBranchRep) SetReferences(v []ApiReferenceRep)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiBranchRep) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetLinks

`func (o *ApiBranchRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiBranchRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiBranchRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiBranchRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


