# InsightsRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The repository ID | 
**Version** | **int32** | The repository version | 
**Key** | **string** | The repository key | 
**Type** | **string** | The repository type | 
**Url** | **string** | The repository URL | 
**MainBranch** | **string** | The repository main branch | 
**Projects** | Pointer to [**ProjectSummaryCollection**](ProjectSummaryCollection.md) |  | [optional] 

## Methods

### NewInsightsRepository

`func NewInsightsRepository(id string, version int32, key string, type_ string, url string, mainBranch string, ) *InsightsRepository`

NewInsightsRepository instantiates a new InsightsRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsRepositoryWithDefaults

`func NewInsightsRepositoryWithDefaults() *InsightsRepository`

NewInsightsRepositoryWithDefaults instantiates a new InsightsRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InsightsRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightsRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightsRepository) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *InsightsRepository) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InsightsRepository) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InsightsRepository) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetKey

`func (o *InsightsRepository) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InsightsRepository) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InsightsRepository) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *InsightsRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InsightsRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InsightsRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *InsightsRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InsightsRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InsightsRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMainBranch

`func (o *InsightsRepository) GetMainBranch() string`

GetMainBranch returns the MainBranch field if non-nil, zero value otherwise.

### GetMainBranchOk

`func (o *InsightsRepository) GetMainBranchOk() (*string, bool)`

GetMainBranchOk returns a tuple with the MainBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainBranch

`func (o *InsightsRepository) SetMainBranch(v string)`

SetMainBranch sets MainBranch field to given value.


### GetProjects

`func (o *InsightsRepository) GetProjects() ProjectSummaryCollection`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *InsightsRepository) GetProjectsOk() (*ProjectSummaryCollection, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *InsightsRepository) SetProjects(v ProjectSummaryCollection)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *InsightsRepository) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


