# UserSegments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]UserSegment**](UserSegment.md) | An array of segments | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**TotalCount** | Pointer to **int32** | The total number of segments | [optional] 

## Methods

### NewUserSegments

`func NewUserSegments(items []UserSegment, links map[string]Link, ) *UserSegments`

NewUserSegments instantiates a new UserSegments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSegmentsWithDefaults

`func NewUserSegmentsWithDefaults() *UserSegments`

NewUserSegmentsWithDefaults instantiates a new UserSegments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserSegments) GetItems() []UserSegment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserSegments) GetItemsOk() (*[]UserSegment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserSegments) SetItems(v []UserSegment)`

SetItems sets Items field to given value.


### GetLinks

`func (o *UserSegments) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSegments) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSegments) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetTotalCount

`func (o *UserSegments) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UserSegments) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UserSegments) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *UserSegments) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


