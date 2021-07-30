# ProjectPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the project. | 
**Key** | **string** | A unique key used to reference the project in your code. | 
**IncludeInSnippetByDefault** | Pointer to **bool** | Whether or not flags created in this project are made available to the client-side JavaScript SDK by default. | [optional] 
**DefaultClientSideAvailability** | Pointer to [**DefaultClientSideAvailabilityPost**](DefaultClientSideAvailabilityPost.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Environments** | Pointer to **[]map[string]interface{}** | Creates the provided environments for this project. If omitted default environments will be created instead. | [optional] 

## Methods

### NewProjectPost

`func NewProjectPost(name string, key string, ) *ProjectPost`

NewProjectPost instantiates a new ProjectPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPostWithDefaults

`func NewProjectPostWithDefaults() *ProjectPost`

NewProjectPostWithDefaults instantiates a new ProjectPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectPost) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ProjectPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetIncludeInSnippetByDefault

`func (o *ProjectPost) GetIncludeInSnippetByDefault() bool`

GetIncludeInSnippetByDefault returns the IncludeInSnippetByDefault field if non-nil, zero value otherwise.

### GetIncludeInSnippetByDefaultOk

`func (o *ProjectPost) GetIncludeInSnippetByDefaultOk() (*bool, bool)`

GetIncludeInSnippetByDefaultOk returns a tuple with the IncludeInSnippetByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippetByDefault

`func (o *ProjectPost) SetIncludeInSnippetByDefault(v bool)`

SetIncludeInSnippetByDefault sets IncludeInSnippetByDefault field to given value.

### HasIncludeInSnippetByDefault

`func (o *ProjectPost) HasIncludeInSnippetByDefault() bool`

HasIncludeInSnippetByDefault returns a boolean if a field has been set.

### GetDefaultClientSideAvailability

`func (o *ProjectPost) GetDefaultClientSideAvailability() DefaultClientSideAvailabilityPost`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *ProjectPost) GetDefaultClientSideAvailabilityOk() (*DefaultClientSideAvailabilityPost, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *ProjectPost) SetDefaultClientSideAvailability(v DefaultClientSideAvailabilityPost)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.

### HasDefaultClientSideAvailability

`func (o *ProjectPost) HasDefaultClientSideAvailability() bool`

HasDefaultClientSideAvailability returns a boolean if a field has been set.

### GetTags

`func (o *ProjectPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvironments

`func (o *ProjectPost) GetEnvironments() []map[string]interface{}`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectPost) GetEnvironmentsOk() (*[]map[string]interface{}, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectPost) SetEnvironments(v []map[string]interface{})`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectPost) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


