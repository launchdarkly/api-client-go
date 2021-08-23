# Destinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Items** | Pointer to [**[]Destination**](Destination.md) |  | [optional] 

## Methods

### NewDestinations

`func NewDestinations() *Destinations`

NewDestinations instantiates a new Destinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationsWithDefaults

`func NewDestinationsWithDefaults() *Destinations`

NewDestinationsWithDefaults instantiates a new Destinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Destinations) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Destinations) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Destinations) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Destinations) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *Destinations) GetItems() []Destination`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Destinations) GetItemsOk() (*[]Destination, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Destinations) SetItems(v []Destination)`

SetItems sets Items field to given value.

### HasItems

`func (o *Destinations) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


