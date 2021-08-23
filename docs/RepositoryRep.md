# RepositoryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SourceLink** | Pointer to **string** |  | [optional] 
**CommitUrlTemplate** | Pointer to **string** |  | [optional] 
**HunkUrlTemplate** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**DefaultBranch** | **string** |  | 
**Enabled** | **bool** |  | 
**Version** | **int32** |  | 
**Branches** | Pointer to [**[]BranchRep**](BranchRep.md) |  | [optional] 
**Links** | **map[string]interface{}** |  | 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewRepositoryRep

`func NewRepositoryRep(name string, type_ string, defaultBranch string, enabled bool, version int32, links map[string]interface{}, ) *RepositoryRep`

NewRepositoryRep instantiates a new RepositoryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryRepWithDefaults

`func NewRepositoryRepWithDefaults() *RepositoryRep`

NewRepositoryRepWithDefaults instantiates a new RepositoryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryRep) SetName(v string)`

SetName sets Name field to given value.


### GetSourceLink

`func (o *RepositoryRep) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *RepositoryRep) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *RepositoryRep) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.

### HasSourceLink

`func (o *RepositoryRep) HasSourceLink() bool`

HasSourceLink returns a boolean if a field has been set.

### GetCommitUrlTemplate

`func (o *RepositoryRep) GetCommitUrlTemplate() string`

GetCommitUrlTemplate returns the CommitUrlTemplate field if non-nil, zero value otherwise.

### GetCommitUrlTemplateOk

`func (o *RepositoryRep) GetCommitUrlTemplateOk() (*string, bool)`

GetCommitUrlTemplateOk returns a tuple with the CommitUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrlTemplate

`func (o *RepositoryRep) SetCommitUrlTemplate(v string)`

SetCommitUrlTemplate sets CommitUrlTemplate field to given value.

### HasCommitUrlTemplate

`func (o *RepositoryRep) HasCommitUrlTemplate() bool`

HasCommitUrlTemplate returns a boolean if a field has been set.

### GetHunkUrlTemplate

`func (o *RepositoryRep) GetHunkUrlTemplate() string`

GetHunkUrlTemplate returns the HunkUrlTemplate field if non-nil, zero value otherwise.

### GetHunkUrlTemplateOk

`func (o *RepositoryRep) GetHunkUrlTemplateOk() (*string, bool)`

GetHunkUrlTemplateOk returns a tuple with the HunkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunkUrlTemplate

`func (o *RepositoryRep) SetHunkUrlTemplate(v string)`

SetHunkUrlTemplate sets HunkUrlTemplate field to given value.

### HasHunkUrlTemplate

`func (o *RepositoryRep) HasHunkUrlTemplate() bool`

HasHunkUrlTemplate returns a boolean if a field has been set.

### GetType

`func (o *RepositoryRep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryRep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryRep) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultBranch

`func (o *RepositoryRep) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *RepositoryRep) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *RepositoryRep) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetEnabled

`func (o *RepositoryRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RepositoryRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RepositoryRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVersion

`func (o *RepositoryRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RepositoryRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RepositoryRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetBranches

`func (o *RepositoryRep) GetBranches() []BranchRep`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *RepositoryRep) GetBranchesOk() (*[]BranchRep, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *RepositoryRep) SetBranches(v []BranchRep)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *RepositoryRep) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetLinks

`func (o *RepositoryRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RepositoryRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RepositoryRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.


### GetAccess

`func (o *RepositoryRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RepositoryRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RepositoryRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *RepositoryRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


