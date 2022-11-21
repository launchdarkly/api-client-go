# ProjectSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Key** | Pointer to **string** | The project key | [optional] 
**Name** | Pointer to **string** | The project name | [optional] 

## Methods

### NewProjectSummary

`func NewProjectSummary() *ProjectSummary`

NewProjectSummary instantiates a new ProjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSummaryWithDefaults

`func NewProjectSummaryWithDefaults() *ProjectSummary`

NewProjectSummaryWithDefaults instantiates a new ProjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectSummary) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectSummary) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectSummary) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectSummary) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetKey

`func (o *ProjectSummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectSummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectSummary) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectSummary) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ProjectSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectSummary) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


