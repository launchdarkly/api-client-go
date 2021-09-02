# FlagListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | **string** |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Site** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewFlagListingRep

`func NewFlagListingRep(name string, key string, ) *FlagListingRep`

NewFlagListingRep instantiates a new FlagListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagListingRepWithDefaults

`func NewFlagListingRepWithDefaults() *FlagListingRep`

NewFlagListingRepWithDefaults instantiates a new FlagListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlagListingRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagListingRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagListingRep) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *FlagListingRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagListingRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagListingRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetLinks

`func (o *FlagListingRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagListingRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagListingRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagListingRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSite

`func (o *FlagListingRep) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *FlagListingRep) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *FlagListingRep) SetSite(v Link)`

SetSite sets Site field to given value.

### HasSite

`func (o *FlagListingRep) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


