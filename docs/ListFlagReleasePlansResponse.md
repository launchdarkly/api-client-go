# ListFlagReleasePlansResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]FlagReleasePlan**](FlagReleasePlan.md) |  | 
**TotalCount** | **int32** | The total number of flag release plans in the project, before pagination. | 
**Links** | Pointer to [**FlagReleasePlansPaginatedLinks**](FlagReleasePlansPaginatedLinks.md) |  | [optional] 

## Methods

### NewListFlagReleasePlansResponse

`func NewListFlagReleasePlansResponse(items []FlagReleasePlan, totalCount int32, ) *ListFlagReleasePlansResponse`

NewListFlagReleasePlansResponse instantiates a new ListFlagReleasePlansResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlagReleasePlansResponseWithDefaults

`func NewListFlagReleasePlansResponseWithDefaults() *ListFlagReleasePlansResponse`

NewListFlagReleasePlansResponseWithDefaults instantiates a new ListFlagReleasePlansResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListFlagReleasePlansResponse) GetItems() []FlagReleasePlan`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFlagReleasePlansResponse) GetItemsOk() (*[]FlagReleasePlan, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFlagReleasePlansResponse) SetItems(v []FlagReleasePlan)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ListFlagReleasePlansResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListFlagReleasePlansResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListFlagReleasePlansResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *ListFlagReleasePlansResponse) GetLinks() FlagReleasePlansPaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFlagReleasePlansResponse) GetLinksOk() (*FlagReleasePlansPaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFlagReleasePlansResponse) SetLinks(v FlagReleasePlansPaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFlagReleasePlansResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


