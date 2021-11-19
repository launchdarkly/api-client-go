# RepositoryCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Items** | [**[]RepositoryRep**](RepositoryRep.md) | An array of repositories | 

## Methods

### NewRepositoryCollectionRep

`func NewRepositoryCollectionRep(links map[string]Link, items []RepositoryRep, ) *RepositoryCollectionRep`

NewRepositoryCollectionRep instantiates a new RepositoryCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryCollectionRepWithDefaults

`func NewRepositoryCollectionRepWithDefaults() *RepositoryCollectionRep`

NewRepositoryCollectionRepWithDefaults instantiates a new RepositoryCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RepositoryCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RepositoryCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RepositoryCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *RepositoryCollectionRep) GetItems() []RepositoryRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RepositoryCollectionRep) GetItemsOk() (*[]RepositoryRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RepositoryCollectionRep) SetItems(v []RepositoryRep)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


