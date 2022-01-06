# Integrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Items** | Pointer to [**[]Integration**](Integration.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrations

`func NewIntegrations() *Integrations`

NewIntegrations instantiates a new Integrations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsWithDefaults

`func NewIntegrationsWithDefaults() *Integrations`

NewIntegrationsWithDefaults instantiates a new Integrations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Integrations) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Integrations) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Integrations) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Integrations) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *Integrations) GetItems() []Integration`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Integrations) GetItemsOk() (*[]Integration, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Integrations) SetItems(v []Integration)`

SetItems sets Items field to given value.

### HasItems

`func (o *Integrations) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKey

`func (o *Integrations) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Integrations) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Integrations) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Integrations) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


