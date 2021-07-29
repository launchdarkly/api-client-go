# GlobalFlagCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GlobalFlagRep**](GlobalFlagRep.md) |  | 
**Links** | [**map[string]InlineResponse200**](InlineResponse200.md) |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewGlobalFlagCollectionRep

`func NewGlobalFlagCollectionRep(items []GlobalFlagRep, links map[string]InlineResponse200, ) *GlobalFlagCollectionRep`

NewGlobalFlagCollectionRep instantiates a new GlobalFlagCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalFlagCollectionRepWithDefaults

`func NewGlobalFlagCollectionRepWithDefaults() *GlobalFlagCollectionRep`

NewGlobalFlagCollectionRepWithDefaults instantiates a new GlobalFlagCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GlobalFlagCollectionRep) GetItems() []GlobalFlagRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GlobalFlagCollectionRep) GetItemsOk() (*[]GlobalFlagRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GlobalFlagCollectionRep) SetItems(v []GlobalFlagRep)`

SetItems sets Items field to given value.


### GetLinks

`func (o *GlobalFlagCollectionRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GlobalFlagCollectionRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GlobalFlagCollectionRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.


### GetTotalCount

`func (o *GlobalFlagCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GlobalFlagCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GlobalFlagCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GlobalFlagCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


