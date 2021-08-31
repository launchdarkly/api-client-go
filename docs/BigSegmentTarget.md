# BigSegmentTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserKey** | **string** | The user key | 
**Included** | **bool** | Indicates whether the user is included.&lt;br /&gt;Included users are always segment members, regardless of segment rules. | 
**Excluded** | **bool** | Indicates whether the user is excluded.&lt;br /&gt;Segment rules bypass excluded users, so they will never be included based on rules. Excluded users may still be included explicitly. | 

## Methods

### NewBigSegmentTarget

`func NewBigSegmentTarget(userKey string, included bool, excluded bool, ) *BigSegmentTarget`

NewBigSegmentTarget instantiates a new BigSegmentTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigSegmentTargetWithDefaults

`func NewBigSegmentTargetWithDefaults() *BigSegmentTarget`

NewBigSegmentTargetWithDefaults instantiates a new BigSegmentTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserKey

`func (o *BigSegmentTarget) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *BigSegmentTarget) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *BigSegmentTarget) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.


### GetIncluded

`func (o *BigSegmentTarget) GetIncluded() bool`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BigSegmentTarget) GetIncludedOk() (*bool, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BigSegmentTarget) SetIncluded(v bool)`

SetIncluded sets Included field to given value.


### GetExcluded

`func (o *BigSegmentTarget) GetExcluded() bool`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *BigSegmentTarget) GetExcludedOk() (*bool, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *BigSegmentTarget) SetExcluded(v bool)`

SetExcluded sets Excluded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


