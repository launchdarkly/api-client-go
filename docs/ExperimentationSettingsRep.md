# ExperimentationSettingsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | The project ID | [optional] 
**ProjectKey** | Pointer to **string** | The project key | [optional] 
**RandomizationUnits** | Pointer to [**[]RandomizationUnitRep**](RandomizationUnitRep.md) | An array of the randomization units in this project | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewExperimentationSettingsRep

`func NewExperimentationSettingsRep() *ExperimentationSettingsRep`

NewExperimentationSettingsRep instantiates a new ExperimentationSettingsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentationSettingsRepWithDefaults

`func NewExperimentationSettingsRepWithDefaults() *ExperimentationSettingsRep`

NewExperimentationSettingsRepWithDefaults instantiates a new ExperimentationSettingsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ExperimentationSettingsRep) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExperimentationSettingsRep) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExperimentationSettingsRep) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ExperimentationSettingsRep) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectKey

`func (o *ExperimentationSettingsRep) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *ExperimentationSettingsRep) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *ExperimentationSettingsRep) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *ExperimentationSettingsRep) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetRandomizationUnits

`func (o *ExperimentationSettingsRep) GetRandomizationUnits() []RandomizationUnitRep`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *ExperimentationSettingsRep) GetRandomizationUnitsOk() (*[]RandomizationUnitRep, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *ExperimentationSettingsRep) SetRandomizationUnits(v []RandomizationUnitRep)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *ExperimentationSettingsRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetCreationDate

`func (o *ExperimentationSettingsRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ExperimentationSettingsRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ExperimentationSettingsRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ExperimentationSettingsRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLinks

`func (o *ExperimentationSettingsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentationSettingsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentationSettingsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentationSettingsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


