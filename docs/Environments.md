# Environments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 

## Methods

### NewEnvironments

`func NewEnvironments() *Environments`

NewEnvironments instantiates a new Environments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsWithDefaults

`func NewEnvironmentsWithDefaults() *Environments`

NewEnvironmentsWithDefaults instantiates a new Environments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Environments) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Environments) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Environments) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Environments) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalCount

`func (o *Environments) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Environments) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Environments) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Environments) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *Environments) GetItems() []Environment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Environments) GetItemsOk() (*[]Environment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Environments) SetItems(v []Environment)`

SetItems sets Items field to given value.

### HasItems

`func (o *Environments) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

