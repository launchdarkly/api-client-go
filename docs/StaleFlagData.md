# StaleFlagData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadyForCodeRemoval** | Pointer to **bool** | Whether the flag is ready for code removal | [optional] 
**ReadyToArchive** | Pointer to **bool** | Whether the flag is ready to be archived | [optional] 
**CleanupId** | Pointer to **string** | If a third-party system helps clean up the flag, the ID from that system | [optional] 

## Methods

### NewStaleFlagData

`func NewStaleFlagData() *StaleFlagData`

NewStaleFlagData instantiates a new StaleFlagData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaleFlagDataWithDefaults

`func NewStaleFlagDataWithDefaults() *StaleFlagData`

NewStaleFlagDataWithDefaults instantiates a new StaleFlagData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadyForCodeRemoval

`func (o *StaleFlagData) GetReadyForCodeRemoval() bool`

GetReadyForCodeRemoval returns the ReadyForCodeRemoval field if non-nil, zero value otherwise.

### GetReadyForCodeRemovalOk

`func (o *StaleFlagData) GetReadyForCodeRemovalOk() (*bool, bool)`

GetReadyForCodeRemovalOk returns a tuple with the ReadyForCodeRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyForCodeRemoval

`func (o *StaleFlagData) SetReadyForCodeRemoval(v bool)`

SetReadyForCodeRemoval sets ReadyForCodeRemoval field to given value.

### HasReadyForCodeRemoval

`func (o *StaleFlagData) HasReadyForCodeRemoval() bool`

HasReadyForCodeRemoval returns a boolean if a field has been set.

### GetReadyToArchive

`func (o *StaleFlagData) GetReadyToArchive() bool`

GetReadyToArchive returns the ReadyToArchive field if non-nil, zero value otherwise.

### GetReadyToArchiveOk

`func (o *StaleFlagData) GetReadyToArchiveOk() (*bool, bool)`

GetReadyToArchiveOk returns a tuple with the ReadyToArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToArchive

`func (o *StaleFlagData) SetReadyToArchive(v bool)`

SetReadyToArchive sets ReadyToArchive field to given value.

### HasReadyToArchive

`func (o *StaleFlagData) HasReadyToArchive() bool`

HasReadyToArchive returns a boolean if a field has been set.

### GetCleanupId

`func (o *StaleFlagData) GetCleanupId() string`

GetCleanupId returns the CleanupId field if non-nil, zero value otherwise.

### GetCleanupIdOk

`func (o *StaleFlagData) GetCleanupIdOk() (*string, bool)`

GetCleanupIdOk returns a tuple with the CleanupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupId

`func (o *StaleFlagData) SetCleanupId(v string)`

SetCleanupId sets CleanupId field to given value.

### HasCleanupId

`func (o *StaleFlagData) HasCleanupId() bool`

HasCleanupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


