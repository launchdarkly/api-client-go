# FilterCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**TotalCount** | Pointer to **int32** | The number of payload filters returned | [optional] 
**Items** | Pointer to [**[]FilterRep**](FilterRep.md) | An array of payload filters | [optional] 

## Methods

### NewFilterCollectionRep

`func NewFilterCollectionRep() *FilterCollectionRep`

NewFilterCollectionRep instantiates a new FilterCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterCollectionRepWithDefaults

`func NewFilterCollectionRepWithDefaults() *FilterCollectionRep`

NewFilterCollectionRepWithDefaults instantiates a new FilterCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FilterCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FilterCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FilterCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FilterCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *FilterCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FilterCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FilterCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FilterCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *FilterCollectionRep) GetItems() []FilterRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FilterCollectionRep) GetItemsOk() (*[]FilterRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FilterCollectionRep) SetItems(v []FilterRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *FilterCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


