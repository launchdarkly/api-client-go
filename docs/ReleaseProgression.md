# ReleaseProgression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**CompletedAt** | Pointer to **int64** |  | [optional] 
**FlagKey** | **string** | The flag key | 
**ActivePhaseId** | Pointer to **string** | The ID of the currently active release phase | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewReleaseProgression

`func NewReleaseProgression(createdAt int64, flagKey string, links map[string]Link, ) *ReleaseProgression`

NewReleaseProgression instantiates a new ReleaseProgression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseProgressionWithDefaults

`func NewReleaseProgressionWithDefaults() *ReleaseProgression`

NewReleaseProgressionWithDefaults instantiates a new ReleaseProgression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ReleaseProgression) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReleaseProgression) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReleaseProgression) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *ReleaseProgression) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ReleaseProgression) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ReleaseProgression) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ReleaseProgression) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetFlagKey

`func (o *ReleaseProgression) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *ReleaseProgression) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *ReleaseProgression) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetActivePhaseId

`func (o *ReleaseProgression) GetActivePhaseId() string`

GetActivePhaseId returns the ActivePhaseId field if non-nil, zero value otherwise.

### GetActivePhaseIdOk

`func (o *ReleaseProgression) GetActivePhaseIdOk() (*string, bool)`

GetActivePhaseIdOk returns a tuple with the ActivePhaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePhaseId

`func (o *ReleaseProgression) SetActivePhaseId(v string)`

SetActivePhaseId sets ActivePhaseId field to given value.

### HasActivePhaseId

`func (o *ReleaseProgression) HasActivePhaseId() bool`

HasActivePhaseId returns a boolean if a field has been set.

### GetLinks

`func (o *ReleaseProgression) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReleaseProgression) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReleaseProgression) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


