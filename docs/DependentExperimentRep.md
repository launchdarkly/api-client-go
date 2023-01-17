# DependentExperimentRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The experiment key | 
**Name** | **string** | The experiment name | 
**EnvironmentId** | **string** | The environment ID | 
**CreationDate** | **int64** |  | 
**ArchivedDate** | Pointer to **int64** |  | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewDependentExperimentRep

`func NewDependentExperimentRep(key string, name string, environmentId string, creationDate int64, links map[string]Link, ) *DependentExperimentRep`

NewDependentExperimentRep instantiates a new DependentExperimentRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentExperimentRepWithDefaults

`func NewDependentExperimentRepWithDefaults() *DependentExperimentRep`

NewDependentExperimentRepWithDefaults instantiates a new DependentExperimentRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DependentExperimentRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DependentExperimentRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DependentExperimentRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DependentExperimentRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentExperimentRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentExperimentRep) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironmentId

`func (o *DependentExperimentRep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DependentExperimentRep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DependentExperimentRep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetCreationDate

`func (o *DependentExperimentRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *DependentExperimentRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *DependentExperimentRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetArchivedDate

`func (o *DependentExperimentRep) GetArchivedDate() int64`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *DependentExperimentRep) GetArchivedDateOk() (*int64, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *DependentExperimentRep) SetArchivedDate(v int64)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *DependentExperimentRep) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetLinks

`func (o *DependentExperimentRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DependentExperimentRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DependentExperimentRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


