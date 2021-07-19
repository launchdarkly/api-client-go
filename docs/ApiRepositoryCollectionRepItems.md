# ApiRepositoryCollectionRepItems

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

### NewApiRepositoryCollectionRepItems

`func NewApiRepositoryCollectionRepItems() *ApiRepositoryCollectionRepItems`

NewApiRepositoryCollectionRepItems instantiates a new ApiRepositoryCollectionRepItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryCollectionRepItemsWithDefaults

`func NewApiRepositoryCollectionRepItemsWithDefaults() *ApiRepositoryCollectionRepItems`

NewApiRepositoryCollectionRepItemsWithDefaults instantiates a new ApiRepositoryCollectionRepItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiRepositoryCollectionRepItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiRepositoryCollectionRepItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiRepositoryCollectionRepItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiRepositoryCollectionRepItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceLink

`func (o *ApiRepositoryCollectionRepItems) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *ApiRepositoryCollectionRepItems) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *ApiRepositoryCollectionRepItems) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.

### HasSourceLink

`func (o *ApiRepositoryCollectionRepItems) HasSourceLink() bool`

HasSourceLink returns a boolean if a field has been set.

### GetCommitUrlTemplate

`func (o *ApiRepositoryCollectionRepItems) GetCommitUrlTemplate() string`

GetCommitUrlTemplate returns the CommitUrlTemplate field if non-nil, zero value otherwise.

### GetCommitUrlTemplateOk

`func (o *ApiRepositoryCollectionRepItems) GetCommitUrlTemplateOk() (*string, bool)`

GetCommitUrlTemplateOk returns a tuple with the CommitUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrlTemplate

`func (o *ApiRepositoryCollectionRepItems) SetCommitUrlTemplate(v string)`

SetCommitUrlTemplate sets CommitUrlTemplate field to given value.

### HasCommitUrlTemplate

`func (o *ApiRepositoryCollectionRepItems) HasCommitUrlTemplate() bool`

HasCommitUrlTemplate returns a boolean if a field has been set.

### GetHunkUrlTemplate

`func (o *ApiRepositoryCollectionRepItems) GetHunkUrlTemplate() string`

GetHunkUrlTemplate returns the HunkUrlTemplate field if non-nil, zero value otherwise.

### GetHunkUrlTemplateOk

`func (o *ApiRepositoryCollectionRepItems) GetHunkUrlTemplateOk() (*string, bool)`

GetHunkUrlTemplateOk returns a tuple with the HunkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunkUrlTemplate

`func (o *ApiRepositoryCollectionRepItems) SetHunkUrlTemplate(v string)`

SetHunkUrlTemplate sets HunkUrlTemplate field to given value.

### HasHunkUrlTemplate

`func (o *ApiRepositoryCollectionRepItems) HasHunkUrlTemplate() bool`

HasHunkUrlTemplate returns a boolean if a field has been set.

### GetType

`func (o *ApiRepositoryCollectionRepItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRepositoryCollectionRepItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRepositoryCollectionRepItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiRepositoryCollectionRepItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ApiRepositoryCollectionRepItems) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiRepositoryCollectionRepItems) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiRepositoryCollectionRepItems) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ApiRepositoryCollectionRepItems) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiRepositoryCollectionRepItems) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiRepositoryCollectionRepItems) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiRepositoryCollectionRepItems) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiRepositoryCollectionRepItems) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *ApiRepositoryCollectionRepItems) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiRepositoryCollectionRepItems) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiRepositoryCollectionRepItems) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiRepositoryCollectionRepItems) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBranches

`func (o *ApiRepositoryCollectionRepItems) GetBranches() []ApiBranchRep`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *ApiRepositoryCollectionRepItems) GetBranchesOk() (*[]ApiBranchRep, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *ApiRepositoryCollectionRepItems) SetBranches(v []ApiBranchRep)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *ApiRepositoryCollectionRepItems) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetLinks

`func (o *ApiRepositoryCollectionRepItems) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiRepositoryCollectionRepItems) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiRepositoryCollectionRepItems) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiRepositoryCollectionRepItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccess

`func (o *ApiRepositoryCollectionRepItems) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApiRepositoryCollectionRepItems) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApiRepositoryCollectionRepItems) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ApiRepositoryCollectionRepItems) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


