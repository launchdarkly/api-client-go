# DependentFlagEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Site** | Pointer to [**CoreLink**](CoreLink.md) |  | [optional] 

## Methods

### NewDependentFlagEnvironment

`func NewDependentFlagEnvironment() *DependentFlagEnvironment`

NewDependentFlagEnvironment instantiates a new DependentFlagEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentFlagEnvironmentWithDefaults

`func NewDependentFlagEnvironmentWithDefaults() *DependentFlagEnvironment`

NewDependentFlagEnvironmentWithDefaults instantiates a new DependentFlagEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DependentFlagEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentFlagEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentFlagEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DependentFlagEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *DependentFlagEnvironment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentFlagEnvironment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentFlagEnvironment) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DependentFlagEnvironment) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLinks

`func (o *DependentFlagEnvironment) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentFlagEnvironment) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentFlagEnvironment) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DependentFlagEnvironment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSite

`func (o *DependentFlagEnvironment) GetSite() CoreLink`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DependentFlagEnvironment) GetSiteOk() (*CoreLink, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DependentFlagEnvironment) SetSite(v CoreLink)`

SetSite sets Site field to given value.

### HasSite

`func (o *DependentFlagEnvironment) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


