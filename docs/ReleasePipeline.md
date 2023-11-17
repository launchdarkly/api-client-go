# ReleasePipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the release pipeline was created | 
**Description** | Pointer to **string** | The release pipeline description | [optional] 
**Key** | **string** | The release pipeline key | 
**Name** | **string** | The release pipeline name | 
**Phases** | [**[]Phase**](Phase.md) | An ordered list of the release pipeline phases. Each phase is a logical grouping of one or more environments that share attributes for rolling out changes. | 
**Tags** | Pointer to **[]string** | A list of the release pipeline&#39;s tags | [optional] 
**Version** | Pointer to **int32** | The release pipeline version | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 

## Methods

### NewReleasePipeline

`func NewReleasePipeline(createdAt time.Time, key string, name string, phases []Phase, ) *ReleasePipeline`

NewReleasePipeline instantiates a new ReleasePipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleasePipelineWithDefaults

`func NewReleasePipelineWithDefaults() *ReleasePipeline`

NewReleasePipelineWithDefaults instantiates a new ReleasePipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ReleasePipeline) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReleasePipeline) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReleasePipeline) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *ReleasePipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReleasePipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReleasePipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReleasePipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *ReleasePipeline) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReleasePipeline) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReleasePipeline) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ReleasePipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleasePipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleasePipeline) SetName(v string)`

SetName sets Name field to given value.


### GetPhases

`func (o *ReleasePipeline) GetPhases() []Phase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *ReleasePipeline) GetPhasesOk() (*[]Phase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *ReleasePipeline) SetPhases(v []Phase)`

SetPhases sets Phases field to given value.


### GetTags

`func (o *ReleasePipeline) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReleasePipeline) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReleasePipeline) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReleasePipeline) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ReleasePipeline) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReleasePipeline) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReleasePipeline) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReleasePipeline) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccess

`func (o *ReleasePipeline) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ReleasePipeline) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ReleasePipeline) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ReleasePipeline) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


