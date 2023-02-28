# ContextInstanceEvaluations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ContextInstanceEvaluation**](ContextInstanceEvaluation.md) | Details on the flag evaluations for this context instance | 
**TotalCount** | Pointer to **int32** | The number of flags | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewContextInstanceEvaluations

`func NewContextInstanceEvaluations(items []ContextInstanceEvaluation, links map[string]Link, ) *ContextInstanceEvaluations`

NewContextInstanceEvaluations instantiates a new ContextInstanceEvaluations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextInstanceEvaluationsWithDefaults

`func NewContextInstanceEvaluationsWithDefaults() *ContextInstanceEvaluations`

NewContextInstanceEvaluationsWithDefaults instantiates a new ContextInstanceEvaluations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContextInstanceEvaluations) GetItems() []ContextInstanceEvaluation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContextInstanceEvaluations) GetItemsOk() (*[]ContextInstanceEvaluation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContextInstanceEvaluations) SetItems(v []ContextInstanceEvaluation)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *ContextInstanceEvaluations) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ContextInstanceEvaluations) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ContextInstanceEvaluations) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ContextInstanceEvaluations) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetLinks

`func (o *ContextInstanceEvaluations) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextInstanceEvaluations) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextInstanceEvaluations) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


