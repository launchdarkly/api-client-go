# FlagStatusCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Items** | Pointer to [**[]FlagStatusRep**](FlagStatusRep.md) |  | [optional] 

## Methods

### NewFlagStatusCollectionRep

`func NewFlagStatusCollectionRep() *FlagStatusCollectionRep`

NewFlagStatusCollectionRep instantiates a new FlagStatusCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagStatusCollectionRepWithDefaults

`func NewFlagStatusCollectionRepWithDefaults() *FlagStatusCollectionRep`

NewFlagStatusCollectionRepWithDefaults instantiates a new FlagStatusCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagStatusCollectionRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagStatusCollectionRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagStatusCollectionRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagStatusCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *FlagStatusCollectionRep) GetItems() []FlagStatusRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlagStatusCollectionRep) GetItemsOk() (*[]FlagStatusRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlagStatusCollectionRep) SetItems(v []FlagStatusRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *FlagStatusCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


