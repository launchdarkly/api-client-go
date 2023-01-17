# BranchRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The branch name | 
**Head** | **string** | An ID representing the branch HEAD. For example, a commit SHA. | 
**UpdateSequenceId** | Pointer to **int64** | An optional ID used to prevent older data from overwriting newer data | [optional] 
**SyncTime** | **int64** |  | 
**References** | Pointer to [**[]ReferenceRep**](ReferenceRep.md) | An array of flag references found on the branch | [optional] 
**Links** | **map[string]interface{}** | The location and content type of related resources | 

## Methods

### NewBranchRep

`func NewBranchRep(name string, head string, syncTime int64, links map[string]interface{}, ) *BranchRep`

NewBranchRep instantiates a new BranchRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchRepWithDefaults

`func NewBranchRepWithDefaults() *BranchRep`

NewBranchRepWithDefaults instantiates a new BranchRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BranchRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchRep) SetName(v string)`

SetName sets Name field to given value.


### GetHead

`func (o *BranchRep) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *BranchRep) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *BranchRep) SetHead(v string)`

SetHead sets Head field to given value.


### GetUpdateSequenceId

`func (o *BranchRep) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *BranchRep) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *BranchRep) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.

### HasUpdateSequenceId

`func (o *BranchRep) HasUpdateSequenceId() bool`

HasUpdateSequenceId returns a boolean if a field has been set.

### GetSyncTime

`func (o *BranchRep) GetSyncTime() int64`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *BranchRep) GetSyncTimeOk() (*int64, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *BranchRep) SetSyncTime(v int64)`

SetSyncTime sets SyncTime field to given value.


### GetReferences

`func (o *BranchRep) GetReferences() []ReferenceRep`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BranchRep) GetReferencesOk() (*[]ReferenceRep, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BranchRep) SetReferences(v []ReferenceRep)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BranchRep) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetLinks

`func (o *BranchRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BranchRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BranchRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


