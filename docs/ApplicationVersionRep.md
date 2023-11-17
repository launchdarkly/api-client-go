# ApplicationVersionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Version** | Pointer to **int32** | Version of the application version | [optional] 
**AutoAdded** | **bool** | Whether the application version was automatically created, because it was included in a context when a LaunchDarkly SDK evaluated a feature flag, or if the application version was created through the LaunchDarkly UI or REST API.  | 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Key** | **string** | The unique identifier of this application version | 
**Name** | **string** | The name of this version | 
**Supported** | Pointer to **bool** | Whether this version is supported. Only applicable if the application &lt;code&gt;kind&lt;/code&gt; is &lt;code&gt;mobile&lt;/code&gt;. | [optional] 

## Methods

### NewApplicationVersionRep

`func NewApplicationVersionRep(autoAdded bool, key string, name string, ) *ApplicationVersionRep`

NewApplicationVersionRep instantiates a new ApplicationVersionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVersionRepWithDefaults

`func NewApplicationVersionRepWithDefaults() *ApplicationVersionRep`

NewApplicationVersionRepWithDefaults instantiates a new ApplicationVersionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ApplicationVersionRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApplicationVersionRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApplicationVersionRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ApplicationVersionRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationVersionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationVersionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationVersionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationVersionRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetVersion

`func (o *ApplicationVersionRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationVersionRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationVersionRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplicationVersionRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAutoAdded

`func (o *ApplicationVersionRep) GetAutoAdded() bool`

GetAutoAdded returns the AutoAdded field if non-nil, zero value otherwise.

### GetAutoAddedOk

`func (o *ApplicationVersionRep) GetAutoAddedOk() (*bool, bool)`

GetAutoAddedOk returns a tuple with the AutoAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdded

`func (o *ApplicationVersionRep) SetAutoAdded(v bool)`

SetAutoAdded sets AutoAdded field to given value.


### GetCreationDate

`func (o *ApplicationVersionRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApplicationVersionRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApplicationVersionRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ApplicationVersionRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetKey

`func (o *ApplicationVersionRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationVersionRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationVersionRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ApplicationVersionRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationVersionRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationVersionRep) SetName(v string)`

SetName sets Name field to given value.


### GetSupported

`func (o *ApplicationVersionRep) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ApplicationVersionRep) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *ApplicationVersionRep) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *ApplicationVersionRep) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


