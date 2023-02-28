# ExpandableApprovalRequestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ExpandableApprovalRequestResponse**](ExpandableApprovalRequestResponse.md) | An array of approval requests | 
**TotalCount** | **int32** | Total number of approval requests | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewExpandableApprovalRequestsResponse

`func NewExpandableApprovalRequestsResponse(items []ExpandableApprovalRequestResponse, totalCount int32, links map[string]Link, ) *ExpandableApprovalRequestsResponse`

NewExpandableApprovalRequestsResponse instantiates a new ExpandableApprovalRequestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandableApprovalRequestsResponseWithDefaults

`func NewExpandableApprovalRequestsResponseWithDefaults() *ExpandableApprovalRequestsResponse`

NewExpandableApprovalRequestsResponseWithDefaults instantiates a new ExpandableApprovalRequestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExpandableApprovalRequestsResponse) GetItems() []ExpandableApprovalRequestResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExpandableApprovalRequestsResponse) GetItemsOk() (*[]ExpandableApprovalRequestResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExpandableApprovalRequestsResponse) SetItems(v []ExpandableApprovalRequestResponse)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ExpandableApprovalRequestsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExpandableApprovalRequestsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExpandableApprovalRequestsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *ExpandableApprovalRequestsResponse) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandableApprovalRequestsResponse) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandableApprovalRequestsResponse) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


