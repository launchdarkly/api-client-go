# PipelineRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Name** | **string** | The release pipeline name | 
**ReleasePipelineKey** | **string** | The release pipeline key | 
**ReleasePipelineDescription** | **string** | The release pipeline description | 
**Phases** | [**[]PipelineReleasePhase**](PipelineReleasePhase.md) | An ordered list of the release pipeline phases | 
**Version** | **int32** | The release version | 
**ReleaseVariationId** | Pointer to **string** | The chosen release variation ID to use across all phases of a release | [optional] 
**CanceledAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewPipelineRelease

`func NewPipelineRelease(name string, releasePipelineKey string, releasePipelineDescription string, phases []PipelineReleasePhase, version int32, ) *PipelineRelease`

NewPipelineRelease instantiates a new PipelineRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReleaseWithDefaults

`func NewPipelineReleaseWithDefaults() *PipelineRelease`

NewPipelineReleaseWithDefaults instantiates a new PipelineRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PipelineRelease) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PipelineRelease) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PipelineRelease) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PipelineRelease) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *PipelineRelease) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineRelease) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineRelease) SetName(v string)`

SetName sets Name field to given value.


### GetReleasePipelineKey

`func (o *PipelineRelease) GetReleasePipelineKey() string`

GetReleasePipelineKey returns the ReleasePipelineKey field if non-nil, zero value otherwise.

### GetReleasePipelineKeyOk

`func (o *PipelineRelease) GetReleasePipelineKeyOk() (*string, bool)`

GetReleasePipelineKeyOk returns a tuple with the ReleasePipelineKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePipelineKey

`func (o *PipelineRelease) SetReleasePipelineKey(v string)`

SetReleasePipelineKey sets ReleasePipelineKey field to given value.


### GetReleasePipelineDescription

`func (o *PipelineRelease) GetReleasePipelineDescription() string`

GetReleasePipelineDescription returns the ReleasePipelineDescription field if non-nil, zero value otherwise.

### GetReleasePipelineDescriptionOk

`func (o *PipelineRelease) GetReleasePipelineDescriptionOk() (*string, bool)`

GetReleasePipelineDescriptionOk returns a tuple with the ReleasePipelineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePipelineDescription

`func (o *PipelineRelease) SetReleasePipelineDescription(v string)`

SetReleasePipelineDescription sets ReleasePipelineDescription field to given value.


### GetPhases

`func (o *PipelineRelease) GetPhases() []PipelineReleasePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *PipelineRelease) GetPhasesOk() (*[]PipelineReleasePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *PipelineRelease) SetPhases(v []PipelineReleasePhase)`

SetPhases sets Phases field to given value.


### GetVersion

`func (o *PipelineRelease) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PipelineRelease) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PipelineRelease) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetReleaseVariationId

`func (o *PipelineRelease) GetReleaseVariationId() string`

GetReleaseVariationId returns the ReleaseVariationId field if non-nil, zero value otherwise.

### GetReleaseVariationIdOk

`func (o *PipelineRelease) GetReleaseVariationIdOk() (*string, bool)`

GetReleaseVariationIdOk returns a tuple with the ReleaseVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVariationId

`func (o *PipelineRelease) SetReleaseVariationId(v string)`

SetReleaseVariationId sets ReleaseVariationId field to given value.

### HasReleaseVariationId

`func (o *PipelineRelease) HasReleaseVariationId() bool`

HasReleaseVariationId returns a boolean if a field has been set.

### GetCanceledAt

`func (o *PipelineRelease) GetCanceledAt() int64`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *PipelineRelease) GetCanceledAtOk() (*int64, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *PipelineRelease) SetCanceledAt(v int64)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *PipelineRelease) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


