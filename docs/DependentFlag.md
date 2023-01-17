# DependentFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The flag name | [optional] 
**Key** | **string** | The flag key | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Site** | [**Link**](Link.md) |  | 

## Methods

### NewDependentFlag

`func NewDependentFlag(key string, links map[string]Link, site Link, ) *DependentFlag`

NewDependentFlag instantiates a new DependentFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentFlagWithDefaults

`func NewDependentFlagWithDefaults() *DependentFlag`

NewDependentFlagWithDefaults instantiates a new DependentFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DependentFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentFlag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DependentFlag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *DependentFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentFlag) SetKey(v string)`

SetKey sets Key field to given value.


### GetLinks

`func (o *DependentFlag) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentFlag) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentFlag) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetSite

`func (o *DependentFlag) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DependentFlag) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DependentFlag) SetSite(v Link)`

SetSite sets Site field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


