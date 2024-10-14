# LayerCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]LayerRep**](LayerRep.md) | The layers in the project | 
**TotalCount** | **int32** | The total number of layers in the project | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewLayerCollectionRep

`func NewLayerCollectionRep(items []LayerRep, totalCount int32, links map[string]Link, ) *LayerCollectionRep`

NewLayerCollectionRep instantiates a new LayerCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerCollectionRepWithDefaults

`func NewLayerCollectionRepWithDefaults() *LayerCollectionRep`

NewLayerCollectionRepWithDefaults instantiates a new LayerCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *LayerCollectionRep) GetItems() []LayerRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LayerCollectionRep) GetItemsOk() (*[]LayerRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LayerCollectionRep) SetItems(v []LayerRep)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *LayerCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *LayerCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *LayerCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *LayerCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LayerCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LayerCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


