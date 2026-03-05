# CountBucketsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]CountBucket**](CountBucket.md) |  | 
**TotalCount** | **int64** |  | 
**BucketIntervalMs** | **int64** |  | 

## Methods

### NewCountBucketsResult

`func NewCountBucketsResult(buckets []CountBucket, totalCount int64, bucketIntervalMs int64, ) *CountBucketsResult`

NewCountBucketsResult instantiates a new CountBucketsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountBucketsResultWithDefaults

`func NewCountBucketsResultWithDefaults() *CountBucketsResult`

NewCountBucketsResultWithDefaults instantiates a new CountBucketsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *CountBucketsResult) GetBuckets() []CountBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *CountBucketsResult) GetBucketsOk() (*[]CountBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *CountBucketsResult) SetBuckets(v []CountBucket)`

SetBuckets sets Buckets field to given value.


### GetTotalCount

`func (o *CountBucketsResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CountBucketsResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CountBucketsResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetBucketIntervalMs

`func (o *CountBucketsResult) GetBucketIntervalMs() int64`

GetBucketIntervalMs returns the BucketIntervalMs field if non-nil, zero value otherwise.

### GetBucketIntervalMsOk

`func (o *CountBucketsResult) GetBucketIntervalMsOk() (*int64, bool)`

GetBucketIntervalMsOk returns a tuple with the BucketIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketIntervalMs

`func (o *CountBucketsResult) SetBucketIntervalMs(v int64)`

SetBucketIntervalMs sets BucketIntervalMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


