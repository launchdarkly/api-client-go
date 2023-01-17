# IntegrationDeliveryConfigurationCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**IntegrationDeliveryConfigurationCollectionLinks**](IntegrationDeliveryConfigurationCollectionLinks.md) |  | 
**Items** | [**[]IntegrationDeliveryConfiguration**](IntegrationDeliveryConfiguration.md) | An array of integration delivery configurations | 

## Methods

### NewIntegrationDeliveryConfigurationCollection

`func NewIntegrationDeliveryConfigurationCollection(links IntegrationDeliveryConfigurationCollectionLinks, items []IntegrationDeliveryConfiguration, ) *IntegrationDeliveryConfigurationCollection`

NewIntegrationDeliveryConfigurationCollection instantiates a new IntegrationDeliveryConfigurationCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDeliveryConfigurationCollectionWithDefaults

`func NewIntegrationDeliveryConfigurationCollectionWithDefaults() *IntegrationDeliveryConfigurationCollection`

NewIntegrationDeliveryConfigurationCollectionWithDefaults instantiates a new IntegrationDeliveryConfigurationCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntegrationDeliveryConfigurationCollection) GetLinks() IntegrationDeliveryConfigurationCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntegrationDeliveryConfigurationCollection) GetLinksOk() (*IntegrationDeliveryConfigurationCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntegrationDeliveryConfigurationCollection) SetLinks(v IntegrationDeliveryConfigurationCollectionLinks)`

SetLinks sets Links field to given value.


### GetItems

`func (o *IntegrationDeliveryConfigurationCollection) GetItems() []IntegrationDeliveryConfiguration`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntegrationDeliveryConfigurationCollection) GetItemsOk() (*[]IntegrationDeliveryConfiguration, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntegrationDeliveryConfigurationCollection) SetItems(v []IntegrationDeliveryConfiguration)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


