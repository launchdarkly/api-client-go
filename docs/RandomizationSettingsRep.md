# RandomizationSettingsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | The project ID | [optional] 
**ProjectKey** | Pointer to **string** | The project key | [optional] 
**RandomizationUnits** | Pointer to [**[]RandomizationUnitRep**](RandomizationUnitRep.md) | An array of the randomization units in this project | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewRandomizationSettingsRep

`func NewRandomizationSettingsRep() *RandomizationSettingsRep`

NewRandomizationSettingsRep instantiates a new RandomizationSettingsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomizationSettingsRepWithDefaults

`func NewRandomizationSettingsRepWithDefaults() *RandomizationSettingsRep`

NewRandomizationSettingsRepWithDefaults instantiates a new RandomizationSettingsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *RandomizationSettingsRep) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RandomizationSettingsRep) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RandomizationSettingsRep) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RandomizationSettingsRep) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectKey

`func (o *RandomizationSettingsRep) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *RandomizationSettingsRep) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *RandomizationSettingsRep) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *RandomizationSettingsRep) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetRandomizationUnits

`func (o *RandomizationSettingsRep) GetRandomizationUnits() []RandomizationUnitRep`

GetRandomizationUnits returns the RandomizationUnits field if non-nil, zero value otherwise.

### GetRandomizationUnitsOk

`func (o *RandomizationSettingsRep) GetRandomizationUnitsOk() (*[]RandomizationUnitRep, bool)`

GetRandomizationUnitsOk returns a tuple with the RandomizationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnits

`func (o *RandomizationSettingsRep) SetRandomizationUnits(v []RandomizationUnitRep)`

SetRandomizationUnits sets RandomizationUnits field to given value.

### HasRandomizationUnits

`func (o *RandomizationSettingsRep) HasRandomizationUnits() bool`

HasRandomizationUnits returns a boolean if a field has been set.

### GetCreationDate

`func (o *RandomizationSettingsRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *RandomizationSettingsRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *RandomizationSettingsRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *RandomizationSettingsRep) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLinks

`func (o *RandomizationSettingsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RandomizationSettingsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RandomizationSettingsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RandomizationSettingsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


