# ExperimentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The experiment name | 
**Description** | Pointer to **string** | The experiment description | [optional] 
**MaintainerId** | Pointer to **string** | The ID of the member who maintains this experiment | [optional] 
**Key** | **string** | The experiment key | 
**Iteration** | [**IterationInput**](IterationInput.md) |  | 
**HoldoutId** | Pointer to **string** | The ID of the holdout | [optional] 
**Tags** | Pointer to **[]string** | Tags for the experiment | [optional] 
**Methodology** | Pointer to **string** | The results analysis approach. | [optional] 
**AnalysisConfig** | Pointer to [**AnalysisConfigInput**](AnalysisConfigInput.md) |  | [optional] 
**DataSource** | Pointer to **string** | The source of metric data in order to analyze results. Defaults to \&quot;launchdarkly\&quot; when not provided. | [optional] 
**Type** | Pointer to **string** | The type of experiment. | [optional] 

## Methods

### NewExperimentPost

`func NewExperimentPost(name string, key string, iteration IterationInput, ) *ExperimentPost`

NewExperimentPost instantiates a new ExperimentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentPostWithDefaults

`func NewExperimentPostWithDefaults() *ExperimentPost`

NewExperimentPostWithDefaults instantiates a new ExperimentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExperimentPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentPost) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExperimentPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperimentPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainerId

`func (o *ExperimentPost) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *ExperimentPost) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *ExperimentPost) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *ExperimentPost) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetKey

`func (o *ExperimentPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentPost) SetKey(v string)`

SetKey sets Key field to given value.


### GetIteration

`func (o *ExperimentPost) GetIteration() IterationInput`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *ExperimentPost) GetIterationOk() (*IterationInput, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *ExperimentPost) SetIteration(v IterationInput)`

SetIteration sets Iteration field to given value.


### GetHoldoutId

`func (o *ExperimentPost) GetHoldoutId() string`

GetHoldoutId returns the HoldoutId field if non-nil, zero value otherwise.

### GetHoldoutIdOk

`func (o *ExperimentPost) GetHoldoutIdOk() (*string, bool)`

GetHoldoutIdOk returns a tuple with the HoldoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutId

`func (o *ExperimentPost) SetHoldoutId(v string)`

SetHoldoutId sets HoldoutId field to given value.

### HasHoldoutId

`func (o *ExperimentPost) HasHoldoutId() bool`

HasHoldoutId returns a boolean if a field has been set.

### GetTags

`func (o *ExperimentPost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentPost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentPost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentPost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMethodology

`func (o *ExperimentPost) GetMethodology() string`

GetMethodology returns the Methodology field if non-nil, zero value otherwise.

### GetMethodologyOk

`func (o *ExperimentPost) GetMethodologyOk() (*string, bool)`

GetMethodologyOk returns a tuple with the Methodology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodology

`func (o *ExperimentPost) SetMethodology(v string)`

SetMethodology sets Methodology field to given value.

### HasMethodology

`func (o *ExperimentPost) HasMethodology() bool`

HasMethodology returns a boolean if a field has been set.

### GetAnalysisConfig

`func (o *ExperimentPost) GetAnalysisConfig() AnalysisConfigInput`

GetAnalysisConfig returns the AnalysisConfig field if non-nil, zero value otherwise.

### GetAnalysisConfigOk

`func (o *ExperimentPost) GetAnalysisConfigOk() (*AnalysisConfigInput, bool)`

GetAnalysisConfigOk returns a tuple with the AnalysisConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisConfig

`func (o *ExperimentPost) SetAnalysisConfig(v AnalysisConfigInput)`

SetAnalysisConfig sets AnalysisConfig field to given value.

### HasAnalysisConfig

`func (o *ExperimentPost) HasAnalysisConfig() bool`

HasAnalysisConfig returns a boolean if a field has been set.

### GetDataSource

`func (o *ExperimentPost) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ExperimentPost) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ExperimentPost) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ExperimentPost) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetType

`func (o *ExperimentPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExperimentPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExperimentPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExperimentPost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


