# ApplicationCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Items** | Pointer to [**[]ApplicationRep**](ApplicationRep.md) | A list of applications | [optional] 
**TotalCount** | Pointer to **int32** | The number of applications | [optional] 

## Methods

### NewApplicationCollectionRep

`func NewApplicationCollectionRep() *ApplicationCollectionRep`

NewApplicationCollectionRep instantiates a new ApplicationCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCollectionRepWithDefaults

`func NewApplicationCollectionRepWithDefaults() *ApplicationCollectionRep`

NewApplicationCollectionRepWithDefaults instantiates a new ApplicationCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *ApplicationCollectionRep) GetItems() []ApplicationRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApplicationCollectionRep) GetItemsOk() (*[]ApplicationRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApplicationCollectionRep) SetItems(v []ApplicationRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApplicationCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *ApplicationCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApplicationCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApplicationCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ApplicationCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


