# ApiBranchCollectionRepItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Head** | Pointer to **string** |  | [optional] 
**UpdateSequenceId** | Pointer to **int64** |  | [optional] 
**SyncTime** | Pointer to **int64** |  | [optional] 
**References** | Pointer to [**[]ApiBranchCollectionRepReferences**](ApiBranchCollectionRepReferences.md) |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiBranchCollectionRepItems

`func NewApiBranchCollectionRepItems() *ApiBranchCollectionRepItems`

NewApiBranchCollectionRepItems instantiates a new ApiBranchCollectionRepItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBranchCollectionRepItemsWithDefaults

`func NewApiBranchCollectionRepItemsWithDefaults() *ApiBranchCollectionRepItems`

NewApiBranchCollectionRepItemsWithDefaults instantiates a new ApiBranchCollectionRepItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiBranchCollectionRepItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBranchCollectionRepItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBranchCollectionRepItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiBranchCollectionRepItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHead

`func (o *ApiBranchCollectionRepItems) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *ApiBranchCollectionRepItems) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *ApiBranchCollectionRepItems) SetHead(v string)`

SetHead sets Head field to given value.

### HasHead

`func (o *ApiBranchCollectionRepItems) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetUpdateSequenceId

`func (o *ApiBranchCollectionRepItems) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *ApiBranchCollectionRepItems) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *ApiBranchCollectionRepItems) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.

### HasUpdateSequenceId

`func (o *ApiBranchCollectionRepItems) HasUpdateSequenceId() bool`

HasUpdateSequenceId returns a boolean if a field has been set.

### GetSyncTime

`func (o *ApiBranchCollectionRepItems) GetSyncTime() int64`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *ApiBranchCollectionRepItems) GetSyncTimeOk() (*int64, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *ApiBranchCollectionRepItems) SetSyncTime(v int64)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *ApiBranchCollectionRepItems) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetReferences

`func (o *ApiBranchCollectionRepItems) GetReferences() []ApiBranchCollectionRepReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiBranchCollectionRepItems) GetReferencesOk() (*[]ApiBranchCollectionRepReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiBranchCollectionRepItems) SetReferences(v []ApiBranchCollectionRepReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiBranchCollectionRepItems) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetLinks

`func (o *ApiBranchCollectionRepItems) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiBranchCollectionRepItems) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiBranchCollectionRepItems) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiBranchCollectionRepItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


