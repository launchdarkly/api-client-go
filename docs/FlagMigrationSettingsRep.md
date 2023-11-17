# FlagMigrationSettingsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextKind** | Pointer to **string** | The context kind targeted by this migration flag. Only applicable for six-stage migrations. | [optional] 
**StageCount** | Pointer to **int32** | The number of stages for this migration flag | [optional] 

## Methods

### NewFlagMigrationSettingsRep

`func NewFlagMigrationSettingsRep() *FlagMigrationSettingsRep`

NewFlagMigrationSettingsRep instantiates a new FlagMigrationSettingsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagMigrationSettingsRepWithDefaults

`func NewFlagMigrationSettingsRepWithDefaults() *FlagMigrationSettingsRep`

NewFlagMigrationSettingsRepWithDefaults instantiates a new FlagMigrationSettingsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextKind

`func (o *FlagMigrationSettingsRep) GetContextKind() string`

GetContextKind returns the ContextKind field if non-nil, zero value otherwise.

### GetContextKindOk

`func (o *FlagMigrationSettingsRep) GetContextKindOk() (*string, bool)`

GetContextKindOk returns a tuple with the ContextKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKind

`func (o *FlagMigrationSettingsRep) SetContextKind(v string)`

SetContextKind sets ContextKind field to given value.

### HasContextKind

`func (o *FlagMigrationSettingsRep) HasContextKind() bool`

HasContextKind returns a boolean if a field has been set.

### GetStageCount

`func (o *FlagMigrationSettingsRep) GetStageCount() int32`

GetStageCount returns the StageCount field if non-nil, zero value otherwise.

### GetStageCountOk

`func (o *FlagMigrationSettingsRep) GetStageCountOk() (*int32, bool)`

GetStageCountOk returns a tuple with the StageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageCount

`func (o *FlagMigrationSettingsRep) SetStageCount(v int32)`

SetStageCount sets StageCount field to given value.

### HasStageCount

`func (o *FlagMigrationSettingsRep) HasStageCount() bool`

HasStageCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


