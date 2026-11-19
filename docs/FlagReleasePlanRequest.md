# FlagReleasePlanRequest

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

## Methods

### NewFlagReleasePlanRequest

`func NewFlagReleasePlanRequest(name string, environments []FlagReleasePlanEnvironment, ) *FlagReleasePlanRequest`

NewFlagReleasePlanRequest instantiates a new FlagReleasePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagReleasePlanRequestWithDefaults

`func NewFlagReleasePlanRequestWithDefaults() *FlagReleasePlanRequest`

NewFlagReleasePlanRequestWithDefaults instantiates a new FlagReleasePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlagReleasePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagReleasePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagReleasePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FlagReleasePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagReleasePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagReleasePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagReleasePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainer

`func (o *FlagReleasePlanRequest) GetMaintainer() MaintainerIdentifier`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *FlagReleasePlanRequest) GetMaintainerOk() (*MaintainerIdentifier, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *FlagReleasePlanRequest) SetMaintainer(v MaintainerIdentifier)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *FlagReleasePlanRequest) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetExecutionMode

`func (o *FlagReleasePlanRequest) GetExecutionMode() ExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *FlagReleasePlanRequest) GetExecutionModeOk() (*ExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *FlagReleasePlanRequest) SetExecutionMode(v ExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *FlagReleasePlanRequest) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetApplicationKeys

`func (o *FlagReleasePlanRequest) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *FlagReleasePlanRequest) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *FlagReleasePlanRequest) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *FlagReleasePlanRequest) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.

### GetDefault

`func (o *FlagReleasePlanRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FlagReleasePlanRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FlagReleasePlanRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FlagReleasePlanRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnvironments

`func (o *FlagReleasePlanRequest) GetEnvironments() []FlagReleasePlanEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *FlagReleasePlanRequest) GetEnvironmentsOk() (*[]FlagReleasePlanEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *FlagReleasePlanRequest) SetEnvironments(v []FlagReleasePlanEnvironment)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


