# SegmentUserState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**SegmentUserList**](SegmentUserList.md) |  | [optional] 
**Excluded** | Pointer to [**SegmentUserList**](SegmentUserList.md) |  | [optional] 

## Methods

### NewSegmentUserState

`func NewSegmentUserState() *SegmentUserState`

NewSegmentUserState instantiates a new SegmentUserState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentUserStateWithDefaults

`func NewSegmentUserStateWithDefaults() *SegmentUserState`

NewSegmentUserStateWithDefaults instantiates a new SegmentUserState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *SegmentUserState) GetIncluded() SegmentUserList`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SegmentUserState) GetIncludedOk() (*SegmentUserList, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SegmentUserState) SetIncluded(v SegmentUserList)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *SegmentUserState) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetExcluded

`func (o *SegmentUserState) GetExcluded() SegmentUserList`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *SegmentUserState) GetExcludedOk() (*SegmentUserList, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *SegmentUserState) SetExcluded(v SegmentUserList)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *SegmentUserState) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


