# CustomWorkflowsListingOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CustomWorkflowOutput**](CustomWorkflowOutput.md) | An array of workflows | 
**TotalCount** | **int32** | Total number of workflows | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewCustomWorkflowsListingOutput

`func NewCustomWorkflowsListingOutput(items []CustomWorkflowOutput, totalCount int32, links map[string]Link, ) *CustomWorkflowsListingOutput`

NewCustomWorkflowsListingOutput instantiates a new CustomWorkflowsListingOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWorkflowsListingOutputWithDefaults

`func NewCustomWorkflowsListingOutputWithDefaults() *CustomWorkflowsListingOutput`

NewCustomWorkflowsListingOutputWithDefaults instantiates a new CustomWorkflowsListingOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CustomWorkflowsListingOutput) GetItems() []CustomWorkflowOutput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomWorkflowsListingOutput) GetItemsOk() (*[]CustomWorkflowOutput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomWorkflowsListingOutput) SetItems(v []CustomWorkflowOutput)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *CustomWorkflowsListingOutput) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CustomWorkflowsListingOutput) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CustomWorkflowsListingOutput) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *CustomWorkflowsListingOutput) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomWorkflowsListingOutput) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomWorkflowsListingOutput) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


