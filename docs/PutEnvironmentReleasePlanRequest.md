# PutEnvironmentReleasePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for this release plan | 
**Version** | Pointer to **int64** | Optional. The plan version this request was based on, from a previous read. Plans start at version 1. When present it must match the plan&#39;s current version — on mismatch the request is rejected with 409 and no changes are applied, so a stale copy can&#39;t silently overwrite a concurrent edit. When omitted, the plan&#39;s version is not checked. | [optional] 
**Description** | Pointer to **string** | Human-readable description of this release plan&#39;s purpose | [optional] 
**ReleaseMethod** | Pointer to **string** | Discriminates which nested config object is required | [optional] [default to "progressive_release"]
**ProgressiveRelease** | [**EnvironmentReleasePlanProgressiveRelease**](EnvironmentReleasePlanProgressiveRelease.md) |  | 

## Methods

### NewPutEnvironmentReleasePlanRequest

`func NewPutEnvironmentReleasePlanRequest(name string, progressiveRelease EnvironmentReleasePlanProgressiveRelease, ) *PutEnvironmentReleasePlanRequest`

NewPutEnvironmentReleasePlanRequest instantiates a new PutEnvironmentReleasePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutEnvironmentReleasePlanRequestWithDefaults

`func NewPutEnvironmentReleasePlanRequestWithDefaults() *PutEnvironmentReleasePlanRequest`

NewPutEnvironmentReleasePlanRequestWithDefaults instantiates a new PutEnvironmentReleasePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutEnvironmentReleasePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutEnvironmentReleasePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutEnvironmentReleasePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *PutEnvironmentReleasePlanRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PutEnvironmentReleasePlanRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PutEnvironmentReleasePlanRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PutEnvironmentReleasePlanRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *PutEnvironmentReleasePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutEnvironmentReleasePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutEnvironmentReleasePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutEnvironmentReleasePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReleaseMethod

`func (o *PutEnvironmentReleasePlanRequest) GetReleaseMethod() string`

GetReleaseMethod returns the ReleaseMethod field if non-nil, zero value otherwise.

### GetReleaseMethodOk

`func (o *PutEnvironmentReleasePlanRequest) GetReleaseMethodOk() (*string, bool)`

GetReleaseMethodOk returns a tuple with the ReleaseMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMethod

`func (o *PutEnvironmentReleasePlanRequest) SetReleaseMethod(v string)`

SetReleaseMethod sets ReleaseMethod field to given value.

### HasReleaseMethod

`func (o *PutEnvironmentReleasePlanRequest) HasReleaseMethod() bool`

HasReleaseMethod returns a boolean if a field has been set.

### GetProgressiveRelease

`func (o *PutEnvironmentReleasePlanRequest) GetProgressiveRelease() EnvironmentReleasePlanProgressiveRelease`

GetProgressiveRelease returns the ProgressiveRelease field if non-nil, zero value otherwise.

### GetProgressiveReleaseOk

`func (o *PutEnvironmentReleasePlanRequest) GetProgressiveReleaseOk() (*EnvironmentReleasePlanProgressiveRelease, bool)`

GetProgressiveReleaseOk returns a tuple with the ProgressiveRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveRelease

`func (o *PutEnvironmentReleasePlanRequest) SetProgressiveRelease(v EnvironmentReleasePlanProgressiveRelease)`

SetProgressiveRelease sets ProgressiveRelease field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


