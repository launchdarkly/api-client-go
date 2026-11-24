# PostEnvironmentReleasePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for this release plan | 
**Key** | **string** | Unique key identifying this release plan within the project. Immutable once created. Unique only among active (non-deleted) plans — a soft-deleted plan&#39;s key can be reused by a new plan. | 
**Description** | Pointer to **string** | Human-readable description of this release plan&#39;s purpose | [optional] 
**ReleaseMethod** | Pointer to **string** | Discriminates which nested config object is required | [optional] [default to "progressive_release"]
**ProgressiveRelease** | [**EnvironmentReleasePlanProgressiveRelease**](EnvironmentReleasePlanProgressiveRelease.md) |  | 

## Methods

### NewPostEnvironmentReleasePlanRequest

`func NewPostEnvironmentReleasePlanRequest(name string, key string, progressiveRelease EnvironmentReleasePlanProgressiveRelease, ) *PostEnvironmentReleasePlanRequest`

NewPostEnvironmentReleasePlanRequest instantiates a new PostEnvironmentReleasePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEnvironmentReleasePlanRequestWithDefaults

`func NewPostEnvironmentReleasePlanRequestWithDefaults() *PostEnvironmentReleasePlanRequest`

NewPostEnvironmentReleasePlanRequestWithDefaults instantiates a new PostEnvironmentReleasePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostEnvironmentReleasePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostEnvironmentReleasePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostEnvironmentReleasePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PostEnvironmentReleasePlanRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostEnvironmentReleasePlanRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostEnvironmentReleasePlanRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetDescription

`func (o *PostEnvironmentReleasePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostEnvironmentReleasePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostEnvironmentReleasePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostEnvironmentReleasePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReleaseMethod

`func (o *PostEnvironmentReleasePlanRequest) GetReleaseMethod() string`

GetReleaseMethod returns the ReleaseMethod field if non-nil, zero value otherwise.

### GetReleaseMethodOk

`func (o *PostEnvironmentReleasePlanRequest) GetReleaseMethodOk() (*string, bool)`

GetReleaseMethodOk returns a tuple with the ReleaseMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseMethod

`func (o *PostEnvironmentReleasePlanRequest) SetReleaseMethod(v string)`

SetReleaseMethod sets ReleaseMethod field to given value.

### HasReleaseMethod

`func (o *PostEnvironmentReleasePlanRequest) HasReleaseMethod() bool`

HasReleaseMethod returns a boolean if a field has been set.

### GetProgressiveRelease

`func (o *PostEnvironmentReleasePlanRequest) GetProgressiveRelease() EnvironmentReleasePlanProgressiveRelease`

GetProgressiveRelease returns the ProgressiveRelease field if non-nil, zero value otherwise.

### GetProgressiveReleaseOk

`func (o *PostEnvironmentReleasePlanRequest) GetProgressiveReleaseOk() (*EnvironmentReleasePlanProgressiveRelease, bool)`

GetProgressiveReleaseOk returns a tuple with the ProgressiveRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveRelease

`func (o *PostEnvironmentReleasePlanRequest) SetProgressiveRelease(v EnvironmentReleasePlanProgressiveRelease)`

SetProgressiveRelease sets ProgressiveRelease field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


