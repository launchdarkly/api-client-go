# SegmentScheduledChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]SegmentScheduledChangesLink**](SegmentScheduledChangesLink.md) | The location and content type of related resources | 
**Items** | [**[]SegmentScheduledChange**](SegmentScheduledChange.md) | The segment&#39;s scheduled changes that have not been applied yet, plus any that failed to apply, soonest first | 

## Methods

### NewSegmentScheduledChanges

`func NewSegmentScheduledChanges(links map[string]SegmentScheduledChangesLink, items []SegmentScheduledChange, ) *SegmentScheduledChanges`

NewSegmentScheduledChanges instantiates a new SegmentScheduledChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentScheduledChangesWithDefaults

`func NewSegmentScheduledChangesWithDefaults() *SegmentScheduledChanges`

NewSegmentScheduledChangesWithDefaults instantiates a new SegmentScheduledChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SegmentScheduledChanges) GetLinks() map[string]SegmentScheduledChangesLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SegmentScheduledChanges) GetLinksOk() (*map[string]SegmentScheduledChangesLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SegmentScheduledChanges) SetLinks(v map[string]SegmentScheduledChangesLink)`

SetLinks sets Links field to given value.


### GetItems

`func (o *SegmentScheduledChanges) GetItems() []SegmentScheduledChange`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SegmentScheduledChanges) GetItemsOk() (*[]SegmentScheduledChange, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SegmentScheduledChanges) SetItems(v []SegmentScheduledChange)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


