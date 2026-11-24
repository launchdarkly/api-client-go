# EnvironmentReleasePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The release plan&#39;s unique identifier | 
**Key** | **string** | Unique key identifying this release plan within the project | 
**Name** | **string** | Human-readable name for this release plan | 
**Description** | Pointer to **string** | Human-readable description of this release plan&#39;s purpose | [optional] 
**EnvironmentKey** | **string** | The environment this release plan is scoped to | 
**ReleaseMethod** | **string** | Discriminates which nested config object is present | 
**Version** | **int64** | The release plan&#39;s version, starting at 1 and incremented on every update | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Default** | Pointer to **bool** | Whether this plan is the default for its environment. At most one active plan per environment is the default. Managed via the environment&#39;s default-release-plan sub-resource, not this request body. | [optional] [readonly] 
**ProgressiveRelease** | [**EnvironmentReleasePlanProgressiveReleaseRep**](EnvironmentReleasePlanProgressiveReleaseRep.md) |  | 

## Methods

### NewEnvironmentReleasePlan

`func NewEnvironmentReleasePlan(id string, key string, name string, environmentKey string, releaseMethod string, version int64, progressiveRelease EnvironmentReleasePlanProgressiveReleaseRep, ) *EnvironmentReleasePlan`

NewEnvironmentReleasePlan instantiates a new EnvironmentReleasePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanWithDefaults

`func NewEnvironmentReleasePlanWithDefaults() *EnvironmentReleasePlan`

NewEnvironmentReleasePlanWithDefaults instantiates a new EnvironmentReleasePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReleasePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReleasePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReleasePlan) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *EnvironmentReleasePlan) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentReleasePlan) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentReleasePlan) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *EnvironmentReleasePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentReleasePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentReleasePlan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnvironmentReleasePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReleasePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReleasePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReleasePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *EnvironmentReleasePlan) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *EnvironmentReleasePlan) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *EnvironmentReleasePlan) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetReleaseMethod

`func (o *EnvironmentReleasePlan) GetReleaseMethod() string`

GetReleaseMethod returns the ReleaseMethod field if non-nil, zero value otherwise.

### GetReleaseMethodOk

`func (o *EnvironmentReleasePlan) GetReleaseMethodOk() (*string, bool)`

GetReleaseMethodOk returns a tuple with the ReleaseMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMethod

`func (o *EnvironmentReleasePlan) SetReleaseMethod(v string)`

SetReleaseMethod sets ReleaseMethod field to given value.


### GetVersion

`func (o *EnvironmentReleasePlan) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentReleasePlan) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentReleasePlan) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *EnvironmentReleasePlan) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentReleasePlan) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentReleasePlan) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentReleasePlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentReleasePlan) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentReleasePlan) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentReleasePlan) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentReleasePlan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *EnvironmentReleasePlan) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *EnvironmentReleasePlan) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *EnvironmentReleasePlan) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *EnvironmentReleasePlan) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetProgressiveRelease

`func (o *EnvironmentReleasePlan) GetProgressiveRelease() EnvironmentReleasePlanProgressiveReleaseRep`

GetProgressiveRelease returns the ProgressiveRelease field if non-nil, zero value otherwise.

### GetProgressiveReleaseOk

`func (o *EnvironmentReleasePlan) GetProgressiveReleaseOk() (*EnvironmentReleasePlanProgressiveReleaseRep, bool)`

GetProgressiveReleaseOk returns a tuple with the ProgressiveRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveRelease

`func (o *EnvironmentReleasePlan) SetProgressiveRelease(v EnvironmentReleasePlanProgressiveReleaseRep)`

SetProgressiveRelease sets ProgressiveRelease field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


