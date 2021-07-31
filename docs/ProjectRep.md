# ProjectRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**IncludeInSnippetByDefault** | Pointer to **bool** |  | [optional] 
**DefaultClientSideAvailability** | Pointer to [**AccountsClientSideAvailability**](AccountsClientSideAvailability.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Environments** | Pointer to [**[]EnvironmentRep**](EnvironmentRep.md) |  | [optional] 

## Methods

### NewProjectRep

`func NewProjectRep() *ProjectRep`

NewProjectRep instantiates a new ProjectRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRepWithDefaults

`func NewProjectRepWithDefaults() *ProjectRep`

NewProjectRepWithDefaults instantiates a new ProjectRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProjectRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ProjectRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIncludeInSnippetByDefault

`func (o *ProjectRep) GetIncludeInSnippetByDefault() bool`

GetIncludeInSnippetByDefault returns the IncludeInSnippetByDefault field if non-nil, zero value otherwise.

### GetIncludeInSnippetByDefaultOk

`func (o *ProjectRep) GetIncludeInSnippetByDefaultOk() (*bool, bool)`

GetIncludeInSnippetByDefaultOk returns a tuple with the IncludeInSnippetByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippetByDefault

`func (o *ProjectRep) SetIncludeInSnippetByDefault(v bool)`

SetIncludeInSnippetByDefault sets IncludeInSnippetByDefault field to given value.

### HasIncludeInSnippetByDefault

`func (o *ProjectRep) HasIncludeInSnippetByDefault() bool`

HasIncludeInSnippetByDefault returns a boolean if a field has been set.

### GetDefaultClientSideAvailability

`func (o *ProjectRep) GetDefaultClientSideAvailability() AccountsClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *ProjectRep) GetDefaultClientSideAvailabilityOk() (*AccountsClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *ProjectRep) SetDefaultClientSideAvailability(v AccountsClientSideAvailability)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.

### HasDefaultClientSideAvailability

`func (o *ProjectRep) HasDefaultClientSideAvailability() bool`

HasDefaultClientSideAvailability returns a boolean if a field has been set.

### GetName

`func (o *ProjectRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *ProjectRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectRep) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvironments

`func (o *ProjectRep) GetEnvironments() []EnvironmentRep`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectRep) GetEnvironmentsOk() (*[]EnvironmentRep, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectRep) SetEnvironments(v []EnvironmentRep)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectRep) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


