# ExtinctionCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Items** | [**map[string][]Extinction**](array.md) | An array of extinction events | 

## Methods

### NewExtinctionCollectionRep

`func NewExtinctionCollectionRep(links map[string]Link, items map[string][]Extinction, ) *ExtinctionCollectionRep`

NewExtinctionCollectionRep instantiates a new ExtinctionCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtinctionCollectionRepWithDefaults

`func NewExtinctionCollectionRepWithDefaults() *ExtinctionCollectionRep`

NewExtinctionCollectionRepWithDefaults instantiates a new ExtinctionCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ExtinctionCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtinctionCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtinctionCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *ExtinctionCollectionRep) GetItems() map[string][]Extinction`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExtinctionCollectionRep) GetItemsOk() (*map[string][]Extinction, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExtinctionCollectionRep) SetItems(v map[string][]Extinction)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


