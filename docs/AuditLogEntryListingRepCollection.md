# AuditLogEntryListingRepCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AuditLogEntryListingRep**](AuditLogEntryListingRep.md) | An array of audit log entries | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewAuditLogEntryListingRepCollection

`func NewAuditLogEntryListingRepCollection(items []AuditLogEntryListingRep, links map[string]Link, ) *AuditLogEntryListingRepCollection`

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

`func (o *AuditLogEntryListingRepCollection) GetItems() []AuditLogEntryListingRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditLogEntryListingRepCollection) GetItemsOk() (*[]AuditLogEntryListingRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditLogEntryListingRepCollection) SetItems(v []AuditLogEntryListingRep)`

SetItems sets Items field to given value.


### GetLinks

`func (o *AuditLogEntryListingRepCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuditLogEntryListingRepCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuditLogEntryListingRepCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


