# ProjectRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Id** | **string** |  | 
**Key** | **string** |  | 
**IncludeInSnippetByDefault** | **bool** |  | 
**DefaultClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**Name** | **string** |  | 
**Tags** | **[]string** |  | 
**Environments** | [**[]Environment**](Environment.md) |  | 

## Methods

### NewProjectRep

`func NewProjectRep(links map[string]Link, id string, key string, includeInSnippetByDefault bool, name string, tags []string, environments []Environment, ) *ProjectRep`

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

`func (o *ProjectRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


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


### GetDefaultClientSideAvailability

`func (o *ProjectRep) GetDefaultClientSideAvailability() ClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *ProjectRep) GetDefaultClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *ProjectRep) SetDefaultClientSideAvailability(v ClientSideAvailability)`

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


### GetEnvironments

`func (o *ProjectRep) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectRep) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectRep) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


