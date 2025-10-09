# ReleasePoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ReleasePolicy**](ReleasePolicy.md) | List of release policies | 
**TotalCount** | **int32** | Total number of release policies | 

## Methods

### NewReleasePoliciesResponse

`func NewReleasePoliciesResponse(items []ReleasePolicy, totalCount int32, ) *ReleasePoliciesResponse`

NewReleasePoliciesResponse instantiates a new ReleasePoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePoliciesResponseWithDefaults

`func NewReleasePoliciesResponseWithDefaults() *ReleasePoliciesResponse`

NewReleasePoliciesResponseWithDefaults instantiates a new ReleasePoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ReleasePoliciesResponse) GetItems() []ReleasePolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ReleasePoliciesResponse) GetItemsOk() (*[]ReleasePolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ReleasePoliciesResponse) SetItems(v []ReleasePolicy)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ReleasePoliciesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReleasePoliciesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReleasePoliciesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


