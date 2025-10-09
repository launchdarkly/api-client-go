# SegmentsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**LinkedSegments** | Pointer to [**ExpandedDirectlyLinkedSegments**](ExpandedDirectlyLinkedSegments.md) |  | [optional] 

## Methods

### NewSegmentsSummary

`func NewSegmentsSummary(count int32, ) *SegmentsSummary`

NewSegmentsSummary instantiates a new SegmentsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentsSummaryWithDefaults

`func NewSegmentsSummaryWithDefaults() *SegmentsSummary`

NewSegmentsSummaryWithDefaults instantiates a new SegmentsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SegmentsSummary) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SegmentsSummary) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SegmentsSummary) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLinkedSegments

`func (o *SegmentsSummary) GetLinkedSegments() ExpandedDirectlyLinkedSegments`

GetLinkedSegments returns the LinkedSegments field if non-nil, zero value otherwise.

### GetLinkedSegmentsOk

`func (o *SegmentsSummary) GetLinkedSegmentsOk() (*ExpandedDirectlyLinkedSegments, bool)`

GetLinkedSegmentsOk returns a tuple with the LinkedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedSegments

`func (o *SegmentsSummary) SetLinkedSegments(v ExpandedDirectlyLinkedSegments)`

SetLinkedSegments sets LinkedSegments field to given value.

### HasLinkedSegments

`func (o *SegmentsSummary) HasLinkedSegments() bool`

HasLinkedSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


