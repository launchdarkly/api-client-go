# DependentFlagsByEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DependentFlag**](DependentFlag.md) | A list of dependent flags, which are flags that use the requested flag as a prerequisite | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Site** | [**Link**](Link.md) |  | 

## Methods

### NewDependentFlagsByEnvironment

`func NewDependentFlagsByEnvironment(items []DependentFlag, links map[string]Link, site Link, ) *DependentFlagsByEnvironment`

NewDependentFlagsByEnvironment instantiates a new DependentFlagsByEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentFlagsByEnvironmentWithDefaults

`func NewDependentFlagsByEnvironmentWithDefaults() *DependentFlagsByEnvironment`

NewDependentFlagsByEnvironmentWithDefaults instantiates a new DependentFlagsByEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DependentFlagsByEnvironment) GetItems() []DependentFlag`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DependentFlagsByEnvironment) GetItemsOk() (*[]DependentFlag, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DependentFlagsByEnvironment) SetItems(v []DependentFlag)`

SetItems sets Items field to given value.


### GetLinks

`func (o *DependentFlagsByEnvironment) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentFlagsByEnvironment) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentFlagsByEnvironment) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetSite

`func (o *DependentFlagsByEnvironment) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DependentFlagsByEnvironment) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DependentFlagsByEnvironment) SetSite(v Link)`

SetSite sets Site field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


