# FlagReleasePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Id** | **string** |  | 
**Version** | **int64** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Maintainer** | Pointer to [**MaintainerResponse**](MaintainerResponse.md) |  | [optional] 
**ExecutionMode** | Pointer to [**ExecutionMode**](ExecutionMode.md) |  | [optional] [default to SEQUENTIAL]
**ApplicationKeys** | Pointer to **[]string** |  | [optional] 
**Default** | **bool** | Whether this plan is the project&#39;s default. At most one active plan per project is the default. | [default to false]
**Environments** | [**[]FlagReleasePlanEnvironment**](FlagReleasePlanEnvironment.md) |  | 

## Methods

### NewFlagReleasePlan

`func NewFlagReleasePlan(key string, id string, version int64, name string, default_ bool, environments []FlagReleasePlanEnvironment, ) *FlagReleasePlan`

NewFlagReleasePlan instantiates a new FlagReleasePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagReleasePlanWithDefaults

`func NewFlagReleasePlanWithDefaults() *FlagReleasePlan`

NewFlagReleasePlanWithDefaults instantiates a new FlagReleasePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlagReleasePlan) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagReleasePlan) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagReleasePlan) SetKey(v string)`

SetKey sets Key field to given value.


### GetId

`func (o *FlagReleasePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagReleasePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagReleasePlan) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *FlagReleasePlan) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlagReleasePlan) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlagReleasePlan) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetName

`func (o *FlagReleasePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagReleasePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagReleasePlan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FlagReleasePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagReleasePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagReleasePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagReleasePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainer

`func (o *FlagReleasePlan) GetMaintainer() MaintainerResponse`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *FlagReleasePlan) GetMaintainerOk() (*MaintainerResponse, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *FlagReleasePlan) SetMaintainer(v MaintainerResponse)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *FlagReleasePlan) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetExecutionMode

`func (o *FlagReleasePlan) GetExecutionMode() ExecutionMode`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *FlagReleasePlan) GetExecutionModeOk() (*ExecutionMode, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *FlagReleasePlan) SetExecutionMode(v ExecutionMode)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *FlagReleasePlan) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetApplicationKeys

`func (o *FlagReleasePlan) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *FlagReleasePlan) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *FlagReleasePlan) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *FlagReleasePlan) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.

### GetDefault

`func (o *FlagReleasePlan) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FlagReleasePlan) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FlagReleasePlan) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEnvironments

`func (o *FlagReleasePlan) GetEnvironments() []FlagReleasePlanEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *FlagReleasePlan) GetEnvironmentsOk() (*[]FlagReleasePlanEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *FlagReleasePlan) SetEnvironments(v []FlagReleasePlanEnvironment)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


