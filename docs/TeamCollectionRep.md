# TeamCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TeamRep**](TeamRep.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewTeamCollectionRep

`func NewTeamCollectionRep() *TeamCollectionRep`

NewTeamCollectionRep instantiates a new TeamCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCollectionRepWithDefaults

`func NewTeamCollectionRepWithDefaults() *TeamCollectionRep`

NewTeamCollectionRepWithDefaults instantiates a new TeamCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TeamCollectionRep) GetItems() []TeamRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamCollectionRep) GetItemsOk() (*[]TeamRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamCollectionRep) SetItems(v []TeamRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *TeamCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *TeamCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *TeamCollectionRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamCollectionRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamCollectionRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamCollectionRep) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


