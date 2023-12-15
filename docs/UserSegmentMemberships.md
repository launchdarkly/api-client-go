# UserSegmentMemberships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]UserSegmentMembership**](UserSegmentMembership.md) |  | 
**Links** | [**map[string]Link**](Link.md) |  | 

## Methods

### NewUserSegmentMemberships

`func NewUserSegmentMemberships(items []UserSegmentMembership, links map[string]Link, ) *UserSegmentMemberships`

NewUserSegmentMemberships instantiates a new UserSegmentMemberships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSegmentMembershipsWithDefaults

`func NewUserSegmentMembershipsWithDefaults() *UserSegmentMemberships`

NewUserSegmentMembershipsWithDefaults instantiates a new UserSegmentMemberships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserSegmentMemberships) GetItems() []UserSegmentMembership`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserSegmentMemberships) GetItemsOk() (*[]UserSegmentMembership, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserSegmentMemberships) SetItems(v []UserSegmentMembership)`

SetItems sets Items field to given value.


### GetLinks

`func (o *UserSegmentMemberships) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSegmentMemberships) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSegmentMemberships) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


