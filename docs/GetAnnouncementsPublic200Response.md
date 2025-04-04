# GetAnnouncementsPublic200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AnnouncementResponse**](AnnouncementResponse.md) |  | 
**Links** | [**AnnouncementPaginatedLinks**](AnnouncementPaginatedLinks.md) |  | 

## Methods

### NewGetAnnouncementsPublic200Response

`func NewGetAnnouncementsPublic200Response(items []AnnouncementResponse, links AnnouncementPaginatedLinks, ) *GetAnnouncementsPublic200Response`

NewGetAnnouncementsPublic200Response instantiates a new GetAnnouncementsPublic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnnouncementsPublic200ResponseWithDefaults

`func NewGetAnnouncementsPublic200ResponseWithDefaults() *GetAnnouncementsPublic200Response`

NewGetAnnouncementsPublic200ResponseWithDefaults instantiates a new GetAnnouncementsPublic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetAnnouncementsPublic200Response) GetItems() []AnnouncementResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetAnnouncementsPublic200Response) GetItemsOk() (*[]AnnouncementResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetAnnouncementsPublic200Response) SetItems(v []AnnouncementResponse)`

SetItems sets Items field to given value.


### GetLinks

`func (o *GetAnnouncementsPublic200Response) GetLinks() AnnouncementPaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAnnouncementsPublic200Response) GetLinksOk() (*AnnouncementPaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAnnouncementsPublic200Response) SetLinks(v AnnouncementPaginatedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


