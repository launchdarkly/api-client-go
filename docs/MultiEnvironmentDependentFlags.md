# MultiEnvironmentDependentFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]MultiEnvironmentDependentFlag**](MultiEnvironmentDependentFlag.md) | An array of dependent flags with their environment information | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Site** | [**Link**](Link.md) |  | 

## Methods

### NewMultiEnvironmentDependentFlags

`func NewMultiEnvironmentDependentFlags(items []MultiEnvironmentDependentFlag, links map[string]Link, site Link, ) *MultiEnvironmentDependentFlags`

NewMultiEnvironmentDependentFlags instantiates a new MultiEnvironmentDependentFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiEnvironmentDependentFlagsWithDefaults

`func NewMultiEnvironmentDependentFlagsWithDefaults() *MultiEnvironmentDependentFlags`

NewMultiEnvironmentDependentFlagsWithDefaults instantiates a new MultiEnvironmentDependentFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MultiEnvironmentDependentFlags) GetItems() []MultiEnvironmentDependentFlag`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MultiEnvironmentDependentFlags) GetItemsOk() (*[]MultiEnvironmentDependentFlag, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MultiEnvironmentDependentFlags) SetItems(v []MultiEnvironmentDependentFlag)`

SetItems sets Items field to given value.


### GetLinks

`func (o *MultiEnvironmentDependentFlags) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MultiEnvironmentDependentFlags) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MultiEnvironmentDependentFlags) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetSite

`func (o *MultiEnvironmentDependentFlags) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MultiEnvironmentDependentFlags) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MultiEnvironmentDependentFlags) SetSite(v Link)`

SetSite sets Site field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


