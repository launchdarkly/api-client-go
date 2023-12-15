# CreateApplicationsAndVersionsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationWithVersions**](ApplicationWithVersions.md) |  | [optional] 

## Methods

### NewCreateApplicationsAndVersionsInput

`func NewCreateApplicationsAndVersionsInput() *CreateApplicationsAndVersionsInput`

NewCreateApplicationsAndVersionsInput instantiates a new CreateApplicationsAndVersionsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationsAndVersionsInputWithDefaults

`func NewCreateApplicationsAndVersionsInputWithDefaults() *CreateApplicationsAndVersionsInput`

NewCreateApplicationsAndVersionsInputWithDefaults instantiates a new CreateApplicationsAndVersionsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *CreateApplicationsAndVersionsInput) GetApplications() []ApplicationWithVersions`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *CreateApplicationsAndVersionsInput) GetApplicationsOk() (*[]ApplicationWithVersions, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *CreateApplicationsAndVersionsInput) SetApplications(v []ApplicationWithVersions)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *CreateApplicationsAndVersionsInput) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


