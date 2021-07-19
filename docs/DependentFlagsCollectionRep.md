# DependentFlagsCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DependentFlagWithEnvsEnvironments**](DependentFlagWithEnvsEnvironments.md) |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Site** | Pointer to [**CoreLink**](CoreLink.md) |  | [optional] 

## Methods

### NewDependentFlagsCollectionRep

`func NewDependentFlagsCollectionRep() *DependentFlagsCollectionRep`

NewDependentFlagsCollectionRep instantiates a new DependentFlagsCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentFlagsCollectionRepWithDefaults

`func NewDependentFlagsCollectionRepWithDefaults() *DependentFlagsCollectionRep`

NewDependentFlagsCollectionRepWithDefaults instantiates a new DependentFlagsCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DependentFlagsCollectionRep) GetItems() []DependentFlagWithEnvsEnvironments`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DependentFlagsCollectionRep) GetItemsOk() (*[]DependentFlagWithEnvsEnvironments, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DependentFlagsCollectionRep) SetItems(v []DependentFlagWithEnvsEnvironments)`

SetItems sets Items field to given value.

### HasItems

`func (o *DependentFlagsCollectionRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *DependentFlagsCollectionRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentFlagsCollectionRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentFlagsCollectionRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DependentFlagsCollectionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSite

`func (o *DependentFlagsCollectionRep) GetSite() CoreLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DependentFlagsCollectionRep) GetSiteOk() (*CoreLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DependentFlagsCollectionRep) SetSite(v CoreLink)`

SetSite sets Site field to given value.

### HasSite

`func (o *DependentFlagsCollectionRep) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


