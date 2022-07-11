# RepositoryPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The repository name | 
**SourceLink** | Pointer to **string** | A URL to access the repository | [optional] 
**CommitUrlTemplate** | Pointer to **string** | A template for constructing a valid URL to view the commit | [optional] 
**HunkUrlTemplate** | Pointer to **string** | A template for constructing a valid URL to view the hunk | [optional] 
**Type** | Pointer to **string** | The type of repository. If not specified, the default value is &lt;code&gt;custom&lt;/code&gt;. | [optional] 
**DefaultBranch** | Pointer to **string** | The repository&#39;s default branch. If not specified, the default value is &lt;code&gt;master&lt;/code&gt;. | [optional] 

## Methods

### NewRepositoryPost

`func NewRepositoryPost(name string, ) *RepositoryPost`

NewRepositoryPost instantiates a new RepositoryPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryPostWithDefaults

`func NewRepositoryPostWithDefaults() *RepositoryPost`

NewRepositoryPostWithDefaults instantiates a new RepositoryPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryPost) SetName(v string)`

SetName sets Name field to given value.


### GetSourceLink

`func (o *RepositoryPost) GetSourceLink() string`

GetSourceLink returns the SourceLink field if non-nil, zero value otherwise.

### GetSourceLinkOk

`func (o *RepositoryPost) GetSourceLinkOk() (*string, bool)`

GetSourceLinkOk returns a tuple with the SourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLink

`func (o *RepositoryPost) SetSourceLink(v string)`

SetSourceLink sets SourceLink field to given value.

### HasSourceLink

`func (o *RepositoryPost) HasSourceLink() bool`

HasSourceLink returns a boolean if a field has been set.

### GetCommitUrlTemplate

`func (o *RepositoryPost) GetCommitUrlTemplate() string`

GetCommitUrlTemplate returns the CommitUrlTemplate field if non-nil, zero value otherwise.

### GetCommitUrlTemplateOk

`func (o *RepositoryPost) GetCommitUrlTemplateOk() (*string, bool)`

GetCommitUrlTemplateOk returns a tuple with the CommitUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrlTemplate

`func (o *RepositoryPost) SetCommitUrlTemplate(v string)`

SetCommitUrlTemplate sets CommitUrlTemplate field to given value.

### HasCommitUrlTemplate

`func (o *RepositoryPost) HasCommitUrlTemplate() bool`

HasCommitUrlTemplate returns a boolean if a field has been set.

### GetHunkUrlTemplate

`func (o *RepositoryPost) GetHunkUrlTemplate() string`

GetHunkUrlTemplate returns the HunkUrlTemplate field if non-nil, zero value otherwise.

### GetHunkUrlTemplateOk

`func (o *RepositoryPost) GetHunkUrlTemplateOk() (*string, bool)`

GetHunkUrlTemplateOk returns a tuple with the HunkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunkUrlTemplate

`func (o *RepositoryPost) SetHunkUrlTemplate(v string)`

SetHunkUrlTemplate sets HunkUrlTemplate field to given value.

### HasHunkUrlTemplate

`func (o *RepositoryPost) HasHunkUrlTemplate() bool`

HasHunkUrlTemplate returns a boolean if a field has been set.

### GetType

`func (o *RepositoryPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RepositoryPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *RepositoryPost) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *RepositoryPost) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *RepositoryPost) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *RepositoryPost) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


