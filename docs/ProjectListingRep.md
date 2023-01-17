# ProjectListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Id** | **string** | The ID of this project | 
**Key** | **string** | The key of this project | 
**IncludeInSnippetByDefault** | **bool** | Whether or not flags created in this project are made available to the client-side JavaScript SDK by default | 
**DefaultClientSideAvailability** | Pointer to [**ClientSideAvailability**](ClientSideAvailability.md) |  | [optional] 
**Name** | **string** | A human-friendly name for the project | 
**Tags** | **[]string** | A list of tags for the project | 

## Methods

### NewProjectListingRep

`func NewProjectListingRep(links map[string]Link, id string, key string, includeInSnippetByDefault bool, name string, tags []string, ) *ProjectListingRep`

NewProjectListingRep instantiates a new ProjectListingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListingRepWithDefaults

`func NewProjectListingRepWithDefaults() *ProjectListingRep`

NewProjectListingRepWithDefaults instantiates a new ProjectListingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectListingRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectListingRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectListingRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *ProjectListingRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectListingRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectListingRep) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *ProjectListingRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectListingRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectListingRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetIncludeInSnippetByDefault

`func (o *ProjectListingRep) GetIncludeInSnippetByDefault() bool`

GetIncludeInSnippetByDefault returns the IncludeInSnippetByDefault field if non-nil, zero value otherwise.

### GetIncludeInSnippetByDefaultOk

`func (o *ProjectListingRep) GetIncludeInSnippetByDefaultOk() (*bool, bool)`

GetIncludeInSnippetByDefaultOk returns a tuple with the IncludeInSnippetByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSnippetByDefault

`func (o *ProjectListingRep) SetIncludeInSnippetByDefault(v bool)`

SetIncludeInSnippetByDefault sets IncludeInSnippetByDefault field to given value.


### GetDefaultClientSideAvailability

`func (o *ProjectListingRep) GetDefaultClientSideAvailability() ClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *ProjectListingRep) GetDefaultClientSideAvailabilityOk() (*ClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *ProjectListingRep) SetDefaultClientSideAvailability(v ClientSideAvailability)`

SetDefaultClientSideAvailability sets DefaultClientSideAvailability field to given value.

### HasDefaultClientSideAvailability

`func (o *ProjectListingRep) HasDefaultClientSideAvailability() bool`

HasDefaultClientSideAvailability returns a boolean if a field has been set.

### GetName

`func (o *ProjectListingRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectListingRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectListingRep) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *ProjectListingRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectListingRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectListingRep) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


