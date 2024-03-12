# FlagEventCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of flag events | 
**Items** | [**[]FlagEventRep**](FlagEventRep.md) | A list of flag events | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewFlagEventCollectionRep

`func NewFlagEventCollectionRep(totalCount int32, items []FlagEventRep, ) *FlagEventCollectionRep`

NewFlagEventCollectionRep instantiates a new FlagEventCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagEventCollectionRepWithDefaults

`func NewFlagEventCollectionRepWithDefaults() *FlagEventCollectionRep`

NewFlagEventCollectionRepWithDefaults instantiates a new FlagEventCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *FlagEventCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FlagEventCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FlagEventCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *FlagEventCollectionRep) GetItems() []FlagEventRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlagEventCollectionRep) GetItemsOk() (*[]FlagEventRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlagEventCollectionRep) SetItems(v []FlagEventRep)`

SetItems sets Items field to given value.


### GetLinks

`func (o *FlagEventCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagEventCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagEventCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagEventCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


