# AuditLogEntryListingRepCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AuditLogEntryListingRepCollectionItems**](AuditLogEntryListingRepCollectionItems.md) |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 

## Methods

### NewAuditLogEntryListingRepCollection

`func NewAuditLogEntryListingRepCollection() *AuditLogEntryListingRepCollection`

NewAuditLogEntryListingRepCollection instantiates a new AuditLogEntryListingRepCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryListingRepCollectionWithDefaults

`func NewAuditLogEntryListingRepCollectionWithDefaults() *AuditLogEntryListingRepCollection`

NewAuditLogEntryListingRepCollectionWithDefaults instantiates a new AuditLogEntryListingRepCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuditLogEntryListingRepCollection) GetItems() []AuditLogEntryListingRepCollectionItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditLogEntryListingRepCollection) GetItemsOk() (*[]AuditLogEntryListingRepCollectionItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditLogEntryListingRepCollection) SetItems(v []AuditLogEntryListingRepCollectionItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuditLogEntryListingRepCollection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *AuditLogEntryListingRepCollection) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditLogEntryListingRepCollection) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditLogEntryListingRepCollection) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuditLogEntryListingRepCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


