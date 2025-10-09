# ViewLinkRequestSegmentIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentIdentifiers** | [**[]ViewLinkRequestSegmentIdentifier**](ViewLinkRequestSegmentIdentifier.md) | Identifiers of the segments to link/unlink (environmentId and segmentKey) | 
**Comment** | Pointer to **string** | Optional comment for the link/unlink operation | [optional] [default to ""]

## Methods

### NewViewLinkRequestSegmentIdentifiers

`func NewViewLinkRequestSegmentIdentifiers(segmentIdentifiers []ViewLinkRequestSegmentIdentifier, ) *ViewLinkRequestSegmentIdentifiers`

NewViewLinkRequestSegmentIdentifiers instantiates a new ViewLinkRequestSegmentIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLinkRequestSegmentIdentifiersWithDefaults

`func NewViewLinkRequestSegmentIdentifiersWithDefaults() *ViewLinkRequestSegmentIdentifiers`

NewViewLinkRequestSegmentIdentifiersWithDefaults instantiates a new ViewLinkRequestSegmentIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentIdentifiers

`func (o *ViewLinkRequestSegmentIdentifiers) GetSegmentIdentifiers() []ViewLinkRequestSegmentIdentifier`

GetSegmentIdentifiers returns the SegmentIdentifiers field if non-nil, zero value otherwise.

### GetSegmentIdentifiersOk

`func (o *ViewLinkRequestSegmentIdentifiers) GetSegmentIdentifiersOk() (*[]ViewLinkRequestSegmentIdentifier, bool)`

GetSegmentIdentifiersOk returns a tuple with the SegmentIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIdentifiers

`func (o *ViewLinkRequestSegmentIdentifiers) SetSegmentIdentifiers(v []ViewLinkRequestSegmentIdentifier)`

SetSegmentIdentifiers sets SegmentIdentifiers field to given value.


### GetComment

`func (o *ViewLinkRequestSegmentIdentifiers) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ViewLinkRequestSegmentIdentifiers) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ViewLinkRequestSegmentIdentifiers) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ViewLinkRequestSegmentIdentifiers) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


