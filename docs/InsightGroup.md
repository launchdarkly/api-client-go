# InsightGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**Scores** | Pointer to [**InsightGroupScores**](InsightGroupScores.md) |  | [optional] 
**ScoreMetadata** | Pointer to [**InsightGroupCollectionScoreMetadata**](InsightGroupCollectionScoreMetadata.md) |  | [optional] 
**Key** | **string** | The insight group key | 
**Name** | **string** | The insight group name | 
**ProjectKey** | **string** | The project key | 
**EnvironmentKey** | **string** | The environment key | 
**ApplicationKeys** | Pointer to **[]string** | The application keys | [optional] 
**CreatedAt** | **int64** |  | 

## Methods

### NewInsightGroup

`func NewInsightGroup(key string, name string, projectKey string, environmentKey string, createdAt int64, ) *InsightGroup`

NewInsightGroup instantiates a new InsightGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightGroupWithDefaults

`func NewInsightGroupWithDefaults() *InsightGroup`

NewInsightGroupWithDefaults instantiates a new InsightGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *InsightGroup) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InsightGroup) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InsightGroup) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InsightGroup) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetScores

`func (o *InsightGroup) GetScores() InsightGroupScores`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *InsightGroup) GetScoresOk() (*InsightGroupScores, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *InsightGroup) SetScores(v InsightGroupScores)`

SetScores sets Scores field to given value.

### HasScores

`func (o *InsightGroup) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetScoreMetadata

`func (o *InsightGroup) GetScoreMetadata() InsightGroupCollectionScoreMetadata`

GetScoreMetadata returns the ScoreMetadata field if non-nil, zero value otherwise.

### GetScoreMetadataOk

`func (o *InsightGroup) GetScoreMetadataOk() (*InsightGroupCollectionScoreMetadata, bool)`

GetScoreMetadataOk returns a tuple with the ScoreMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMetadata

`func (o *InsightGroup) SetScoreMetadata(v InsightGroupCollectionScoreMetadata)`

SetScoreMetadata sets ScoreMetadata field to given value.

### HasScoreMetadata

`func (o *InsightGroup) HasScoreMetadata() bool`

HasScoreMetadata returns a boolean if a field has been set.

### GetKey

`func (o *InsightGroup) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InsightGroup) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InsightGroup) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *InsightGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProjectKey

`func (o *InsightGroup) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *InsightGroup) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *InsightGroup) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetEnvironmentKey

`func (o *InsightGroup) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *InsightGroup) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *InsightGroup) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetApplicationKeys

`func (o *InsightGroup) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *InsightGroup) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *InsightGroup) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *InsightGroup) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InsightGroup) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InsightGroup) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InsightGroup) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


