# BigSegmentTargetChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserKey** | **string** | The user key | 
**Included** | **bool** | Whether or not the user is included in the segment | 
**Excluded** | **bool** | Whether or not the user is excluded from the segment | 

## Methods

### NewBigSegmentTargetChanges

`func NewBigSegmentTargetChanges(userKey string, included bool, excluded bool, ) *BigSegmentTargetChanges`

NewBigSegmentTargetChanges instantiates a new BigSegmentTargetChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigSegmentTargetChangesWithDefaults

`func NewBigSegmentTargetChangesWithDefaults() *BigSegmentTargetChanges`

NewBigSegmentTargetChangesWithDefaults instantiates a new BigSegmentTargetChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserKey

`func (o *BigSegmentTargetChanges) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *BigSegmentTargetChanges) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *BigSegmentTargetChanges) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.


### GetIncluded

`func (o *BigSegmentTargetChanges) GetIncluded() bool`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BigSegmentTargetChanges) GetIncludedOk() (*bool, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BigSegmentTargetChanges) SetIncluded(v bool)`

SetIncluded sets Included field to given value.


### GetExcluded

`func (o *BigSegmentTargetChanges) GetExcluded() bool`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *BigSegmentTargetChanges) GetExcludedOk() (*bool, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *BigSegmentTargetChanges) SetExcluded(v bool)`

SetExcluded sets Excluded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


