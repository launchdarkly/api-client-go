# TriggerWorkflowCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TriggerWorkflowRep**](TriggerWorkflowRep.md) | An array of flag triggers | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewTriggerWorkflowCollectionRep

`func NewTriggerWorkflowCollectionRep() *TriggerWorkflowCollectionRep`

NewTriggerWorkflowCollectionRep instantiates a new TriggerWorkflowCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWorkflowCollectionRepWithDefaults

`func NewTriggerWorkflowCollectionRepWithDefaults() *TriggerWorkflowCollectionRep`

NewTriggerWorkflowCollectionRepWithDefaults instantiates a new TriggerWorkflowCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TriggerWorkflowCollectionRep) GetItems() []TriggerWorkflowRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TriggerWorkflowCollectionRep) GetItemsOk() (*[]TriggerWorkflowRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TriggerWorkflowCollectionRep) SetItems(v []TriggerWorkflowRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *TriggerWorkflowCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *TriggerWorkflowCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TriggerWorkflowCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TriggerWorkflowCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TriggerWorkflowCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


