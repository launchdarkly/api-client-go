# PutReleasePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**ReleasePolicyScope**](ReleasePolicyScope.md) |  | [optional] 
**ReleaseMethod** | [**ReleaseMethod**](ReleaseMethod.md) |  | 
**GuardedReleaseConfig** | Pointer to [**GuardedReleaseConfig**](GuardedReleaseConfig.md) |  | [optional] 
**ProgressiveReleaseConfig** | Pointer to **map[string]interface{}** | Configuration for progressive releases | [optional] 
**Name** | **string** | The name of the release policy | 

## Methods

### NewPutReleasePolicyRequest

`func NewPutReleasePolicyRequest(releaseMethod ReleaseMethod, name string, ) *PutReleasePolicyRequest`

NewPutReleasePolicyRequest instantiates a new PutReleasePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutReleasePolicyRequestWithDefaults

`func NewPutReleasePolicyRequestWithDefaults() *PutReleasePolicyRequest`

NewPutReleasePolicyRequestWithDefaults instantiates a new PutReleasePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *PutReleasePolicyRequest) GetScope() ReleasePolicyScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PutReleasePolicyRequest) GetScopeOk() (*ReleasePolicyScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PutReleasePolicyRequest) SetScope(v ReleasePolicyScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PutReleasePolicyRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetReleaseMethod

`func (o *PutReleasePolicyRequest) GetReleaseMethod() ReleaseMethod`

GetReleaseMethod returns the ReleaseMethod field if non-nil, zero value otherwise.

### GetReleaseMethodOk

`func (o *PutReleasePolicyRequest) GetReleaseMethodOk() (*ReleaseMethod, bool)`

GetReleaseMethodOk returns a tuple with the ReleaseMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMethod

`func (o *PutReleasePolicyRequest) SetReleaseMethod(v ReleaseMethod)`

SetReleaseMethod sets ReleaseMethod field to given value.


### GetGuardedReleaseConfig

`func (o *PutReleasePolicyRequest) GetGuardedReleaseConfig() GuardedReleaseConfig`

GetGuardedReleaseConfig returns the GuardedReleaseConfig field if non-nil, zero value otherwise.

### GetGuardedReleaseConfigOk

`func (o *PutReleasePolicyRequest) GetGuardedReleaseConfigOk() (*GuardedReleaseConfig, bool)`

GetGuardedReleaseConfigOk returns a tuple with the GuardedReleaseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardedReleaseConfig

`func (o *PutReleasePolicyRequest) SetGuardedReleaseConfig(v GuardedReleaseConfig)`

SetGuardedReleaseConfig sets GuardedReleaseConfig field to given value.

### HasGuardedReleaseConfig

`func (o *PutReleasePolicyRequest) HasGuardedReleaseConfig() bool`

HasGuardedReleaseConfig returns a boolean if a field has been set.

### GetProgressiveReleaseConfig

`func (o *PutReleasePolicyRequest) GetProgressiveReleaseConfig() map[string]interface{}`

GetProgressiveReleaseConfig returns the ProgressiveReleaseConfig field if non-nil, zero value otherwise.

### GetProgressiveReleaseConfigOk

`func (o *PutReleasePolicyRequest) GetProgressiveReleaseConfigOk() (*map[string]interface{}, bool)`

GetProgressiveReleaseConfigOk returns a tuple with the ProgressiveReleaseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveReleaseConfig

`func (o *PutReleasePolicyRequest) SetProgressiveReleaseConfig(v map[string]interface{})`

SetProgressiveReleaseConfig sets ProgressiveReleaseConfig field to given value.

### HasProgressiveReleaseConfig

`func (o *PutReleasePolicyRequest) HasProgressiveReleaseConfig() bool`

HasProgressiveReleaseConfig returns a boolean if a field has been set.

### GetName

`func (o *PutReleasePolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutReleasePolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutReleasePolicyRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


