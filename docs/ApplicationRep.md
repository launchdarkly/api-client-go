# ApplicationRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to [**ApplicationFlagCollectionRep**](ApplicationFlagCollectionRep.md) |  | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Version** | Pointer to **int32** | Version of the application | [optional] 
**AutoAdded** | **bool** | Whether the application was automatically created because it was included in a context when a LaunchDarkly SDK evaluated a feature flag, or was created through the LaunchDarkly UI or REST API. | 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** | The application description | [optional] 
**Key** | **string** | The unique identifier of this application | 
**Kind** | **string** | To distinguish the kind of application | 
**Maintainer** | Pointer to [**MaintainerRep**](MaintainerRep.md) |  | [optional] 
**Name** | **string** | The name of the application | 

## Methods

### NewApplicationRep

`func NewApplicationRep(autoAdded bool, key string, kind string, name string, ) *ApplicationRep`

NewApplicationRep instantiates a new ApplicationRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRepWithDefaults

`func NewApplicationRepWithDefaults() *ApplicationRep`

NewApplicationRepWithDefaults instantiates a new ApplicationRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *ApplicationRep) GetFlags() ApplicationFlagCollectionRep`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApplicationRep) GetFlagsOk() (*ApplicationFlagCollectionRep, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApplicationRep) SetFlags(v ApplicationFlagCollectionRep)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ApplicationRep) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetAccess

`func (o *ApplicationRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApplicationRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApplicationRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ApplicationRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetVersion

`func (o *ApplicationRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplicationRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAutoAdded

`func (o *ApplicationRep) GetAutoAdded() bool`

GetAutoAdded returns the AutoAdded field if non-nil, zero value otherwise.

### GetAutoAddedOk

`func (o *ApplicationRep) GetAutoAddedOk() (*bool, bool)`

GetAutoAddedOk returns a tuple with the AutoAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdded

`func (o *ApplicationRep) SetAutoAdded(v bool)`

SetAutoAdded sets AutoAdded field to given value.


### GetCreationDate

`func (o *ApplicationRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApplicationRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApplicationRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ApplicationRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *ApplicationRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *ApplicationRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApplicationRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApplicationRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMaintainer

`func (o *ApplicationRep) GetMaintainer() MaintainerRep`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *ApplicationRep) GetMaintainerOk() (*MaintainerRep, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *ApplicationRep) SetMaintainer(v MaintainerRep)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *ApplicationRep) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRep) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


