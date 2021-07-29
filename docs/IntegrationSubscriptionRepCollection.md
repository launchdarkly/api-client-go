# IntegrationSubscriptionRepCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Items** | Pointer to [**[]ApiV2IntegrationsItems**](ApiV2IntegrationsItems.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationSubscriptionRepCollection

`func NewIntegrationSubscriptionRepCollection() *IntegrationSubscriptionRepCollection`

NewIntegrationSubscriptionRepCollection instantiates a new IntegrationSubscriptionRepCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSubscriptionRepCollectionWithDefaults

`func NewIntegrationSubscriptionRepCollectionWithDefaults() *IntegrationSubscriptionRepCollection`

NewIntegrationSubscriptionRepCollectionWithDefaults instantiates a new IntegrationSubscriptionRepCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntegrationSubscriptionRepCollection) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntegrationSubscriptionRepCollection) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntegrationSubscriptionRepCollection) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IntegrationSubscriptionRepCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *IntegrationSubscriptionRepCollection) GetItems() []ApiV2IntegrationsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntegrationSubscriptionRepCollection) GetItemsOk() (*[]ApiV2IntegrationsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntegrationSubscriptionRepCollection) SetItems(v []ApiV2IntegrationsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *IntegrationSubscriptionRepCollection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKey

`func (o *IntegrationSubscriptionRepCollection) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IntegrationSubscriptionRepCollection) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IntegrationSubscriptionRepCollection) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IntegrationSubscriptionRepCollection) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


