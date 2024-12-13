# ExpandedResourceRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of resource | 
**Flag** | Pointer to [**ExpandedFlagRep**](ExpandedFlagRep.md) |  | [optional] 
**Segment** | Pointer to [**UserSegment**](UserSegment.md) |  | [optional] 

## Methods

### NewExpandedResourceRep

`func NewExpandedResourceRep(kind string, ) *ExpandedResourceRep`

NewExpandedResourceRep instantiates a new ExpandedResourceRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedResourceRepWithDefaults

`func NewExpandedResourceRepWithDefaults() *ExpandedResourceRep`

NewExpandedResourceRepWithDefaults instantiates a new ExpandedResourceRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ExpandedResourceRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExpandedResourceRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExpandedResourceRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetFlag

`func (o *ExpandedResourceRep) GetFlag() ExpandedFlagRep`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ExpandedResourceRep) GetFlagOk() (*ExpandedFlagRep, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ExpandedResourceRep) SetFlag(v ExpandedFlagRep)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ExpandedResourceRep) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetSegment

`func (o *ExpandedResourceRep) GetSegment() UserSegment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *ExpandedResourceRep) GetSegmentOk() (*UserSegment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *ExpandedResourceRep) SetSegment(v UserSegment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *ExpandedResourceRep) HasSegment() bool`

HasSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


