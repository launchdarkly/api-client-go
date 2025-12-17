# ReleasePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**ReleasePoliciesAccessRep**](ReleasePoliciesAccessRep.md) |  | [optional] 
**Id** | **string** | The unique identifier of the release policy | 
**Scope** | Pointer to [**ReleasePolicyScope**](ReleasePolicyScope.md) |  | [optional] 
**Rank** | **int32** | The rank/priority of the release policy | 
**ReleaseMethod** | [**ReleaseMethod**](ReleaseMethod.md) |  | 
**GuardedReleaseConfig** | Pointer to [**GuardedReleaseConfig**](GuardedReleaseConfig.md) |  | [optional] 
**ProgressiveReleaseConfig** | Pointer to [**ProgressiveReleaseConfig**](ProgressiveReleaseConfig.md) |  | [optional] 
**Name** | **string** | The name of the release policy | 
**Key** | **string** | The human-readable key of the release policy | 

## Methods

### NewReleasePolicy

`func NewReleasePolicy(id string, rank int32, releaseMethod ReleaseMethod, name string, key string, ) *ReleasePolicy`

NewReleasePolicy instantiates a new ReleasePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePolicyWithDefaults

`func NewReleasePolicyWithDefaults() *ReleasePolicy`

NewReleasePolicyWithDefaults instantiates a new ReleasePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ReleasePolicy) GetAccess() ReleasePoliciesAccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ReleasePolicy) GetAccessOk() (*ReleasePoliciesAccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ReleasePolicy) SetAccess(v ReleasePoliciesAccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ReleasePolicy) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetId

`func (o *ReleasePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleasePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleasePolicy) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *ReleasePolicy) GetScope() ReleasePolicyScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ReleasePolicy) GetScopeOk() (*ReleasePolicyScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ReleasePolicy) SetScope(v ReleasePolicyScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ReleasePolicy) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRank

`func (o *ReleasePolicy) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ReleasePolicy) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ReleasePolicy) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetReleaseMethod

`func (o *ReleasePolicy) GetReleaseMethod() ReleaseMethod`

GetReleaseMethod returns the ReleaseMethod field if non-nil, zero value otherwise.

### GetReleaseMethodOk

`func (o *ReleasePolicy) GetReleaseMethodOk() (*ReleaseMethod, bool)`

GetReleaseMethodOk returns a tuple with the ReleaseMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMethod

`func (o *ReleasePolicy) SetReleaseMethod(v ReleaseMethod)`

SetReleaseMethod sets ReleaseMethod field to given value.


### GetGuardedReleaseConfig

`func (o *ReleasePolicy) GetGuardedReleaseConfig() GuardedReleaseConfig`

GetGuardedReleaseConfig returns the GuardedReleaseConfig field if non-nil, zero value otherwise.

### GetGuardedReleaseConfigOk

`func (o *ReleasePolicy) GetGuardedReleaseConfigOk() (*GuardedReleaseConfig, bool)`

GetGuardedReleaseConfigOk returns a tuple with the GuardedReleaseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardedReleaseConfig

`func (o *ReleasePolicy) SetGuardedReleaseConfig(v GuardedReleaseConfig)`

SetGuardedReleaseConfig sets GuardedReleaseConfig field to given value.

### HasGuardedReleaseConfig

`func (o *ReleasePolicy) HasGuardedReleaseConfig() bool`

HasGuardedReleaseConfig returns a boolean if a field has been set.

### GetProgressiveReleaseConfig

`func (o *ReleasePolicy) GetProgressiveReleaseConfig() ProgressiveReleaseConfig`

GetProgressiveReleaseConfig returns the ProgressiveReleaseConfig field if non-nil, zero value otherwise.

### GetProgressiveReleaseConfigOk

`func (o *ReleasePolicy) GetProgressiveReleaseConfigOk() (*ProgressiveReleaseConfig, bool)`

GetProgressiveReleaseConfigOk returns a tuple with the ProgressiveReleaseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveReleaseConfig

`func (o *ReleasePolicy) SetProgressiveReleaseConfig(v ProgressiveReleaseConfig)`

SetProgressiveReleaseConfig sets ProgressiveReleaseConfig field to given value.

### HasProgressiveReleaseConfig

`func (o *ReleasePolicy) HasProgressiveReleaseConfig() bool`

HasProgressiveReleaseConfig returns a boolean if a field has been set.

### GetName

`func (o *ReleasePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleasePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleasePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ReleasePolicy) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReleasePolicy) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReleasePolicy) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


