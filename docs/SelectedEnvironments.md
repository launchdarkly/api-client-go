# SelectedEnvironments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**SelectedEnvironmentKey** | **string** | The key for the selected environment in the project | 
**Environments** | [**[]InternalEnvironment**](InternalEnvironment.md) | An ordered list of selected environments in this project | 

## Methods

### NewSelectedEnvironments

`func NewSelectedEnvironments(selectedEnvironmentKey string, environments []InternalEnvironment, ) *SelectedEnvironments`

NewSelectedEnvironments instantiates a new SelectedEnvironments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedEnvironmentsWithDefaults

`func NewSelectedEnvironmentsWithDefaults() *SelectedEnvironments`

NewSelectedEnvironmentsWithDefaults instantiates a new SelectedEnvironments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SelectedEnvironments) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SelectedEnvironments) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SelectedEnvironments) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SelectedEnvironments) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSelectedEnvironmentKey

`func (o *SelectedEnvironments) GetSelectedEnvironmentKey() string`

GetSelectedEnvironmentKey returns the SelectedEnvironmentKey field if non-nil, zero value otherwise.

### GetSelectedEnvironmentKeyOk

`func (o *SelectedEnvironments) GetSelectedEnvironmentKeyOk() (*string, bool)`

GetSelectedEnvironmentKeyOk returns a tuple with the SelectedEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEnvironmentKey

`func (o *SelectedEnvironments) SetSelectedEnvironmentKey(v string)`

SetSelectedEnvironmentKey sets SelectedEnvironmentKey field to given value.


### GetEnvironments

`func (o *SelectedEnvironments) GetEnvironments() []InternalEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *SelectedEnvironments) GetEnvironmentsOk() (*[]InternalEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *SelectedEnvironments) SetEnvironments(v []InternalEnvironment)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


