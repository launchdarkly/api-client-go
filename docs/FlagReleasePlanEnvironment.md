# FlagReleasePlanEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentKey** | **string** |  | 
**EnvironmentReleasePlanKey** | Pointer to **string** | The environment release plan to run. Omitted for an immediate release. | [optional] 
**IsImmediate** | Pointer to **bool** |  | [optional] 
**TriggerType** | [**TriggerType**](TriggerType.md) |  | 

## Methods

### NewFlagReleasePlanEnvironment

`func NewFlagReleasePlanEnvironment(environmentKey string, triggerType TriggerType, ) *FlagReleasePlanEnvironment`

NewFlagReleasePlanEnvironment instantiates a new FlagReleasePlanEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagReleasePlanEnvironmentWithDefaults

`func NewFlagReleasePlanEnvironmentWithDefaults() *FlagReleasePlanEnvironment`

NewFlagReleasePlanEnvironmentWithDefaults instantiates a new FlagReleasePlanEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentKey

`func (o *FlagReleasePlanEnvironment) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *FlagReleasePlanEnvironment) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *FlagReleasePlanEnvironment) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetEnvironmentReleasePlanKey

`func (o *FlagReleasePlanEnvironment) GetEnvironmentReleasePlanKey() string`

GetEnvironmentReleasePlanKey returns the EnvironmentReleasePlanKey field if non-nil, zero value otherwise.

### GetEnvironmentReleasePlanKeyOk

`func (o *FlagReleasePlanEnvironment) GetEnvironmentReleasePlanKeyOk() (*string, bool)`

GetEnvironmentReleasePlanKeyOk returns a tuple with the EnvironmentReleasePlanKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentReleasePlanKey

`func (o *FlagReleasePlanEnvironment) SetEnvironmentReleasePlanKey(v string)`

SetEnvironmentReleasePlanKey sets EnvironmentReleasePlanKey field to given value.

### HasEnvironmentReleasePlanKey

`func (o *FlagReleasePlanEnvironment) HasEnvironmentReleasePlanKey() bool`

HasEnvironmentReleasePlanKey returns a boolean if a field has been set.

### GetIsImmediate

`func (o *FlagReleasePlanEnvironment) GetIsImmediate() bool`

GetIsImmediate returns the IsImmediate field if non-nil, zero value otherwise.

### GetIsImmediateOk

`func (o *FlagReleasePlanEnvironment) GetIsImmediateOk() (*bool, bool)`

GetIsImmediateOk returns a tuple with the IsImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImmediate

`func (o *FlagReleasePlanEnvironment) SetIsImmediate(v bool)`

SetIsImmediate sets IsImmediate field to given value.

### HasIsImmediate

`func (o *FlagReleasePlanEnvironment) HasIsImmediate() bool`

HasIsImmediate returns a boolean if a field has been set.

### GetTriggerType

`func (o *FlagReleasePlanEnvironment) GetTriggerType() TriggerType`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *FlagReleasePlanEnvironment) GetTriggerTypeOk() (*TriggerType, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *FlagReleasePlanEnvironment) SetTriggerType(v TriggerType)`

SetTriggerType sets TriggerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


