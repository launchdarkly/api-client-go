# ExpandedFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key used to reference the flag | 
**Name** | **string** | A human-friendly name for the flag | 
**Description** | Pointer to **string** | Description of the flag | [optional] 
**CreationDate** | Pointer to **int64** | Creation date in milliseconds | [optional] 
**Version** | Pointer to **int32** | Version of the flag | [optional] 
**Archived** | Pointer to **bool** | Whether the flag is archived | [optional] 
**Tags** | Pointer to **[]string** | Tags for the flag | [optional] 
**Temporary** | Pointer to **bool** | Whether the flag is temporary | [optional] 
**IncludeInSnippet** | Pointer to **bool** | Whether to include in snippet | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 

## Methods

### NewExpandedFlag

`func NewExpandedFlag(key string, name string, ) *ExpandedFlag`

NewExpandedFlag instantiates a new ExpandedFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedFlagWithDefaults

`func NewExpandedFlagWithDefaults() *ExpandedFlag`

NewExpandedFlagWithDefaults instantiates a new ExpandedFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExpandedFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedFlag) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ExpandedFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedFlag) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExpandedFlag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandedFlag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandedFlag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandedFlag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreationDate

`func (o *ExpandedFlag) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ExpandedFlag) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ExpandedFlag) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ExpandedFlag) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetVersion

`func (o *ExpandedFlag) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpandedFlag) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpandedFlag) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExpandedFlag) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArchived

`func (o *ExpandedFlag) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ExpandedFlag) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ExpandedFlag) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ExpandedFlag) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTags

`func (o *ExpandedFlag) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExpandedFlag) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExpandedFlag) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExpandedFlag) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemporary

`func (o *ExpandedFlag) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *ExpandedFlag) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *ExpandedFlag) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *ExpandedFlag) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetIncludeInSnippet

`func (o *ExpandedFlag) GetIncludeInSnippet() bool`

GetIncludeInSnippet returns the IncludeInSnippet field if non-nil, zero value otherwise.

### GetIncludeInSnippetOk

`func (o *ExpandedFlag) GetIncludeInSnippetOk() (*bool, bool)`

GetIncludeInSnippetOk returns a tuple with the IncludeInSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippet

`func (o *ExpandedFlag) SetIncludeInSnippet(v bool)`

SetIncludeInSnippet sets IncludeInSnippet field to given value.

### HasIncludeInSnippet

`func (o *ExpandedFlag) HasIncludeInSnippet() bool`

HasIncludeInSnippet returns a boolean if a field has been set.

### GetLinks

`func (o *ExpandedFlag) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandedFlag) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandedFlag) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpandedFlag) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


