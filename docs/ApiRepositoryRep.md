# ApiRepositoryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SourceLink** | Pointer to **string** |  | [optional] 
**CommitUrlTemplate** | Pointer to **string** |  | [optional] 
**HunkUrlTemplate** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DefaultBranch** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Branches** | Pointer to [**[]ApiBranchRep**](ApiBranchRep.md) |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewApiRepositoryRep

`func NewApiRepositoryRep() *ApiRepositoryRep`

NewApiRepositoryRep instantiates a new ApiRepositoryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryRepWithDefaults

`func NewApiRepositoryRepWithDefaults() *ApiRepositoryRep`

NewApiRepositoryRepWithDefaults instantiates a new ApiRepositoryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiRepositoryRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiRepositoryRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiRepositoryRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiRepositoryRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceLink

`func (o *ApiRepositoryRep) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *ApiRepositoryRep) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *ApiRepositoryRep) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.

### HasSourceLink

`func (o *ApiRepositoryRep) HasSourceLink() bool`

HasSourceLink returns a boolean if a field has been set.

### GetCommitUrlTemplate

`func (o *ApiRepositoryRep) GetCommitUrlTemplate() string`

GetCommitUrlTemplate returns the CommitUrlTemplate field if non-nil, zero value otherwise.

### GetCommitUrlTemplateOk

`func (o *ApiRepositoryRep) GetCommitUrlTemplateOk() (*string, bool)`

GetCommitUrlTemplateOk returns a tuple with the CommitUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrlTemplate

`func (o *ApiRepositoryRep) SetCommitUrlTemplate(v string)`

SetCommitUrlTemplate sets CommitUrlTemplate field to given value.

### HasCommitUrlTemplate

`func (o *ApiRepositoryRep) HasCommitUrlTemplate() bool`

HasCommitUrlTemplate returns a boolean if a field has been set.

### GetHunkUrlTemplate

`func (o *ApiRepositoryRep) GetHunkUrlTemplate() string`

GetHunkUrlTemplate returns the HunkUrlTemplate field if non-nil, zero value otherwise.

### GetHunkUrlTemplateOk

`func (o *ApiRepositoryRep) GetHunkUrlTemplateOk() (*string, bool)`

GetHunkUrlTemplateOk returns a tuple with the HunkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunkUrlTemplate

`func (o *ApiRepositoryRep) SetHunkUrlTemplate(v string)`

SetHunkUrlTemplate sets HunkUrlTemplate field to given value.

### HasHunkUrlTemplate

`func (o *ApiRepositoryRep) HasHunkUrlTemplate() bool`

HasHunkUrlTemplate returns a boolean if a field has been set.

### GetType

`func (o *ApiRepositoryRep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRepositoryRep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRepositoryRep) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiRepositoryRep) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ApiRepositoryRep) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiRepositoryRep) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiRepositoryRep) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ApiRepositoryRep) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiRepositoryRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiRepositoryRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiRepositoryRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiRepositoryRep) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *ApiRepositoryRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiRepositoryRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiRepositoryRep) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiRepositoryRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBranches

`func (o *ApiRepositoryRep) GetBranches() []ApiBranchRep`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *ApiRepositoryRep) GetBranchesOk() (*[]ApiBranchRep, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *ApiRepositoryRep) SetBranches(v []ApiBranchRep)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *ApiRepositoryRep) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetLinks

`func (o *ApiRepositoryRep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiRepositoryRep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiRepositoryRep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiRepositoryRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccess

`func (o *ApiRepositoryRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApiRepositoryRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApiRepositoryRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ApiRepositoryRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


