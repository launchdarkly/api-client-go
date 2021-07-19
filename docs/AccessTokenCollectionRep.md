# AccessTokenCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AccessTokenRep**](AccessTokenRep.md) |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 

## Methods

### NewAccessTokenCollectionRep

`func NewAccessTokenCollectionRep() *AccessTokenCollectionRep`

NewAccessTokenCollectionRep instantiates a new AccessTokenCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCollectionRepWithDefaults

`func NewAccessTokenCollectionRepWithDefaults() *AccessTokenCollectionRep`

NewAccessTokenCollectionRepWithDefaults instantiates a new AccessTokenCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AccessTokenCollectionRep) GetItems() []AccessTokenRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccessTokenCollectionRep) GetItemsOk() (*[]AccessTokenRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccessTokenCollectionRep) SetItems(v []AccessTokenRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *AccessTokenCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *AccessTokenCollectionRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccessTokenCollectionRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccessTokenCollectionRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccessTokenCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


