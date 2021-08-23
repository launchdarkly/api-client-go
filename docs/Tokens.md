# Tokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewTokens

`func NewTokens() *Tokens`

NewTokens instantiates a new Tokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensWithDefaults

`func NewTokensWithDefaults() *Tokens`

NewTokensWithDefaults instantiates a new Tokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Tokens) GetItems() []Token`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Tokens) GetItemsOk() (*[]Token, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Tokens) SetItems(v []Token)`

SetItems sets Items field to given value.

### HasItems

`func (o *Tokens) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *Tokens) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Tokens) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Tokens) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Tokens) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


