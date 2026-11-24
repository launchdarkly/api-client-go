# PutFlagReleasePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Maintainer** | Pointer to [**MaintainerIdentifier**](MaintainerIdentifier.md) |  | [optional] 
**ExecutionMode** | Pointer to [**ExecutionMode**](ExecutionMode.md) |  | [optional] [default to SEQUENTIAL]
**ApplicationKeys** | Pointer to **[]string** |  | [optional] 
**Default** | Pointer to **bool** | Whether this plan is the project&#39;s default. At most one active plan per project is the default. On update, omitting it demotes the current default. | [optional] 
**Environments** | [**[]FlagReleasePlanEnvironment**](FlagReleasePlanEnvironment.md) |  | 
**Version** | Pointer to **int64** | Optional. The plan version this request was based on, from a previous read. Plans start at version 1. When present it must match the plan&#39;s current version — on mismatch the request is rejected with 409 and no changes are applied, so a stale copy can&#39;t silently overwrite a concurrent edit. When omitted, the plan&#39;s version is not checked.  | [optional] 

## Methods

### NewPutFlagReleasePlanRequest

`func NewPutFlagReleasePlanRequest(name string, environments []FlagReleasePlanEnvironment, ) *PutFlagReleasePlanRequest`

NewPutFlagReleasePlanRequest instantiates a new PutFlagReleasePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFlagReleasePlanRequestWithDefaults

`func NewPutFlagReleasePlanRequestWithDefaults() *PutFlagReleasePlanRequest`

NewPutFlagReleasePlanRequestWithDefaults instantiates a new PutFlagReleasePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutFlagReleasePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutFlagReleasePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutFlagReleasePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PutFlagReleasePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutFlagReleasePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutFlagReleasePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutFlagReleasePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainer

`func (o *PutFlagReleasePlanRequest) GetMaintainer() MaintainerIdentifier`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *PutFlagReleasePlanRequest) GetMaintainerOk() (*MaintainerIdentifier, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *PutFlagReleasePlanRequest) SetMaintainer(v MaintainerIdentifier)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *PutFlagReleasePlanRequest) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetExecutionMode

`func (o *PutFlagReleasePlanRequest) GetExecutionMode() ExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *PutFlagReleasePlanRequest) GetExecutionModeOk() (*ExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *PutFlagReleasePlanRequest) SetExecutionMode(v ExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *PutFlagReleasePlanRequest) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetApplicationKeys

`func (o *PutFlagReleasePlanRequest) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *PutFlagReleasePlanRequest) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *PutFlagReleasePlanRequest) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *PutFlagReleasePlanRequest) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.

### GetDefault

`func (o *PutFlagReleasePlanRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PutFlagReleasePlanRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PutFlagReleasePlanRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PutFlagReleasePlanRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnvironments

`func (o *PutFlagReleasePlanRequest) GetEnvironments() []FlagReleasePlanEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PutFlagReleasePlanRequest) GetEnvironmentsOk() (*[]FlagReleasePlanEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PutFlagReleasePlanRequest) SetEnvironments(v []FlagReleasePlanEnvironment)`

SetEnvironments sets Environments field to given value.


### GetVersion

`func (o *PutFlagReleasePlanRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PutFlagReleasePlanRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PutFlagReleasePlanRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PutFlagReleasePlanRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


