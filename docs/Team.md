# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**LastModified** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewTeam

`func NewTeam() *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *Team) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Team) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Team) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Team) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Team) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccess

`func (o *Team) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Team) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Team) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Team) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCreationDate

`func (o *Team) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Team) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Team) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Team) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLinks

`func (o *Team) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Team) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Team) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Team) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetLastModified

`func (o *Team) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Team) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Team) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *Team) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetVersion

`func (o *Team) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Team) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Team) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Team) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


