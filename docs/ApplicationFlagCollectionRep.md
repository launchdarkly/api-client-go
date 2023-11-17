# ApplicationFlagCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]FlagListingRep**](FlagListingRep.md) | A list of the flags that have been evaluated by the application | [optional] 
**TotalCount** | Pointer to **int32** | The number of flags that have been evaluated by the application | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewApplicationFlagCollectionRep

`func NewApplicationFlagCollectionRep() *ApplicationFlagCollectionRep`

NewApplicationFlagCollectionRep instantiates a new ApplicationFlagCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFlagCollectionRepWithDefaults

`func NewApplicationFlagCollectionRepWithDefaults() *ApplicationFlagCollectionRep`

NewApplicationFlagCollectionRepWithDefaults instantiates a new ApplicationFlagCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ApplicationFlagCollectionRep) GetItems() []FlagListingRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationFlagCollectionRep) GetItemsOk() (*[]FlagListingRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationFlagCollectionRep) SetItems(v []FlagListingRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationFlagCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *ApplicationFlagCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApplicationFlagCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApplicationFlagCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ApplicationFlagCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationFlagCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationFlagCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationFlagCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationFlagCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


