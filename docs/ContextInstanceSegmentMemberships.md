# ContextInstanceSegmentMemberships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ContextInstanceSegmentMembership**](ContextInstanceSegmentMembership.md) |  | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewContextInstanceSegmentMemberships

`func NewContextInstanceSegmentMemberships(items []ContextInstanceSegmentMembership, links map[string]Link, ) *ContextInstanceSegmentMemberships`

NewContextInstanceSegmentMemberships instantiates a new ContextInstanceSegmentMemberships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextInstanceSegmentMembershipsWithDefaults

`func NewContextInstanceSegmentMembershipsWithDefaults() *ContextInstanceSegmentMemberships`

NewContextInstanceSegmentMembershipsWithDefaults instantiates a new ContextInstanceSegmentMemberships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContextInstanceSegmentMemberships) GetItems() []ContextInstanceSegmentMembership`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContextInstanceSegmentMemberships) GetItemsOk() (*[]ContextInstanceSegmentMembership, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContextInstanceSegmentMemberships) SetItems(v []ContextInstanceSegmentMembership)`

SetItems sets Items field to given value.


### GetLinks

`func (o *ContextInstanceSegmentMemberships) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextInstanceSegmentMemberships) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextInstanceSegmentMemberships) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


