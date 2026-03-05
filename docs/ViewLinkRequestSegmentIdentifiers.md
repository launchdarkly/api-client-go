# ViewLinkRequestSegmentIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentIdentifiers** | [**[]ViewLinkRequestSegmentIdentifier**](ViewLinkRequestSegmentIdentifier.md) | Identifiers of the segments to link/unlink (environmentId and segmentKey) | 
**Filter** | Pointer to **string** | Optional filter string to determine which resources should be linked. Resources only need to match either the filter or explicitly-listed keys to be linked (union). Uses the same queryfilter syntax as the segments list endpoint.  Supported filters for segments: query, tags, keys, excludedKeys, unbounded, external, view, type  | [optional] 
**EnvironmentId** | Pointer to **string** | Required when using filter for segment resources. Specifies which environment to query for segments matching the filter. Ignored when only using explicit segmentIdentifiers (since each identifier contains its own environmentId).  | [optional] 
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


### GetFilter

`func (o *ViewLinkRequestSegmentIdentifiers) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ViewLinkRequestSegmentIdentifiers) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ViewLinkRequestSegmentIdentifiers) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ViewLinkRequestSegmentIdentifiers) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ViewLinkRequestSegmentIdentifiers) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ViewLinkRequestSegmentIdentifiers) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ViewLinkRequestSegmentIdentifiers) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ViewLinkRequestSegmentIdentifiers) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

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


