# ApiExtinctionCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 
**Items** | Pointer to [**map[string][]ApiExtinctionRep**](array.md) |  | [optional] 

## Methods

### NewApiExtinctionCollectionRep

`func NewApiExtinctionCollectionRep() *ApiExtinctionCollectionRep`

NewApiExtinctionCollectionRep instantiates a new ApiExtinctionCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExtinctionCollectionRepWithDefaults

`func NewApiExtinctionCollectionRepWithDefaults() *ApiExtinctionCollectionRep`

NewApiExtinctionCollectionRepWithDefaults instantiates a new ApiExtinctionCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApiExtinctionCollectionRep) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiExtinctionCollectionRep) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiExtinctionCollectionRep) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiExtinctionCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *ApiExtinctionCollectionRep) GetItems() map[string][]ApiExtinctionRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApiExtinctionCollectionRep) GetItemsOk() (*map[string][]ApiExtinctionRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApiExtinctionCollectionRep) SetItems(v map[string][]ApiExtinctionRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApiExtinctionCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


