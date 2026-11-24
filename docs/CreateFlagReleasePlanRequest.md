# CreateFlagReleasePlanRequest

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
**Key** | **string** |  | 

## Methods

### NewCreateFlagReleasePlanRequest

`func NewCreateFlagReleasePlanRequest(name string, environments []FlagReleasePlanEnvironment, key string, ) *CreateFlagReleasePlanRequest`

NewCreateFlagReleasePlanRequest instantiates a new CreateFlagReleasePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlagReleasePlanRequestWithDefaults

`func NewCreateFlagReleasePlanRequestWithDefaults() *CreateFlagReleasePlanRequest`

NewCreateFlagReleasePlanRequestWithDefaults instantiates a new CreateFlagReleasePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFlagReleasePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFlagReleasePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFlagReleasePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateFlagReleasePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFlagReleasePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFlagReleasePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFlagReleasePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainer

`func (o *CreateFlagReleasePlanRequest) GetMaintainer() MaintainerIdentifier`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *CreateFlagReleasePlanRequest) GetMaintainerOk() (*MaintainerIdentifier, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *CreateFlagReleasePlanRequest) SetMaintainer(v MaintainerIdentifier)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *CreateFlagReleasePlanRequest) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CreateFlagReleasePlanRequest) GetExecutionMode() ExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CreateFlagReleasePlanRequest) GetExecutionModeOk() (*ExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CreateFlagReleasePlanRequest) SetExecutionMode(v ExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CreateFlagReleasePlanRequest) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetApplicationKeys

`func (o *CreateFlagReleasePlanRequest) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *CreateFlagReleasePlanRequest) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *CreateFlagReleasePlanRequest) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *CreateFlagReleasePlanRequest) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.

### GetDefault

`func (o *CreateFlagReleasePlanRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CreateFlagReleasePlanRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CreateFlagReleasePlanRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CreateFlagReleasePlanRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnvironments

`func (o *CreateFlagReleasePlanRequest) GetEnvironments() []FlagReleasePlanEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *CreateFlagReleasePlanRequest) GetEnvironmentsOk() (*[]FlagReleasePlanEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *CreateFlagReleasePlanRequest) SetEnvironments(v []FlagReleasePlanEnvironment)`

SetEnvironments sets Environments field to given value.


### GetKey

`func (o *CreateFlagReleasePlanRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateFlagReleasePlanRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateFlagReleasePlanRequest) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


