# MigrationSettingsPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextKind** | Pointer to **string** | Context kind for a migration with 6 stages, where data is being moved | [optional] 
**StageCount** | **int32** |  | 

## Methods

### NewMigrationSettingsPost

`func NewMigrationSettingsPost(stageCount int32, ) *MigrationSettingsPost`

NewMigrationSettingsPost instantiates a new MigrationSettingsPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationSettingsPostWithDefaults

`func NewMigrationSettingsPostWithDefaults() *MigrationSettingsPost`

NewMigrationSettingsPostWithDefaults instantiates a new MigrationSettingsPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextKind

`func (o *MigrationSettingsPost) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *MigrationSettingsPost) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *MigrationSettingsPost) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *MigrationSettingsPost) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.

### GetStageCount

`func (o *MigrationSettingsPost) GetStageCount() int32`

GetStageCount returns the StageCount field if non-nil, zero value otherwise.

### GetStageCountOk

`func (o *MigrationSettingsPost) GetStageCountOk() (*int32, bool)`

GetStageCountOk returns a tuple with the StageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageCount

`func (o *MigrationSettingsPost) SetStageCount(v int32)`

SetStageCount sets StageCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


