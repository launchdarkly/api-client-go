# EnvironmentReleasePlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**EnvironmentReleasePlansPaginatedLinks**](EnvironmentReleasePlansPaginatedLinks.md) |  | [optional] 
**Items** | [**[]EnvironmentReleasePlan**](EnvironmentReleasePlan.md) | One page of the listed release plans, ordered by creation time | 
**TotalCount** | **int32** | The total number of release plans in the listed scope, before pagination | 

## Methods

### NewEnvironmentReleasePlans

`func NewEnvironmentReleasePlans(items []EnvironmentReleasePlan, totalCount int32, ) *EnvironmentReleasePlans`

NewEnvironmentReleasePlans instantiates a new EnvironmentReleasePlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlansWithDefaults

`func NewEnvironmentReleasePlansWithDefaults() *EnvironmentReleasePlans`

NewEnvironmentReleasePlansWithDefaults instantiates a new EnvironmentReleasePlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentReleasePlans) GetLinks() EnvironmentReleasePlansPaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentReleasePlans) GetLinksOk() (*EnvironmentReleasePlansPaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentReleasePlans) SetLinks(v EnvironmentReleasePlansPaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentReleasePlans) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *EnvironmentReleasePlans) GetItems() []EnvironmentReleasePlan`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentReleasePlans) GetItemsOk() (*[]EnvironmentReleasePlan, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentReleasePlans) SetItems(v []EnvironmentReleasePlan)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *EnvironmentReleasePlans) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EnvironmentReleasePlans) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EnvironmentReleasePlans) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


