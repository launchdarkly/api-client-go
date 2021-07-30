# ProjectListingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**IncludeInSnippetByDefault** | Pointer to **bool** |  | [optional] 
**DefaultClientSideAvailability** | Pointer to [**AccountsClientSideAvailability**](AccountsClientSideAvailability.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Environments** | Pointer to [**[]EnvironmentRep**](EnvironmentRep.md) |  | [optional] 

## Methods

### NewProjectListingRep

`func NewProjectListingRep() *ProjectListingRep`

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

`func (o *ProjectListingRep) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectListingRep) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectListingRep) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectListingRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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

### HasId

`func (o *ProjectListingRep) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasKey

`func (o *ProjectListingRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

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

### HasIncludeInSnippetByDefault

`func (o *ProjectListingRep) HasIncludeInSnippetByDefault() bool`

HasIncludeInSnippetByDefault returns a boolean if a field has been set.

### GetDefaultClientSideAvailability

`func (o *ProjectListingRep) GetDefaultClientSideAvailability() AccountsClientSideAvailability`

GetDefaultClientSideAvailability returns the DefaultClientSideAvailability field if non-nil, zero value otherwise.

### GetDefaultClientSideAvailabilityOk

`func (o *ProjectListingRep) GetDefaultClientSideAvailabilityOk() (*AccountsClientSideAvailability, bool)`

GetDefaultClientSideAvailabilityOk returns a tuple with the DefaultClientSideAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientSideAvailability

`func (o *ProjectListingRep) SetDefaultClientSideAvailability(v AccountsClientSideAvailability)`

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

### HasName

`func (o *ProjectListingRep) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasTags

`func (o *ProjectListingRep) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnvironments

`func (o *ProjectListingRep) GetEnvironments() []EnvironmentRep`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ProjectListingRep) GetEnvironmentsOk() (*[]EnvironmentRep, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ProjectListingRep) SetEnvironments(v []EnvironmentRep)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ProjectListingRep) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


