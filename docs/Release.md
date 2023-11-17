# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Name** | **string** | The release pipeline name | 
**ReleasePipelineKey** | **string** | The release pipeline key | 
**ReleasePipelineDescription** | **string** | The release pipeline description | 
**Phases** | [**[]ReleasePhase**](ReleasePhase.md) | An ordered list of the release pipeline phases | 
**Version** | **int32** | The release version | 

## Methods

### NewRelease

`func NewRelease(name string, releasePipelineKey string, releasePipelineDescription string, phases []ReleasePhase, version int32, ) *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Release) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Release) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Release) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Release) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Release) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Release) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Release) SetName(v string)`

SetName sets Name field to given value.


### GetReleasePipelineKey

`func (o *Release) GetReleasePipelineKey() string`

GetReleasePipelineKey returns the ReleasePipelineKey field if non-nil, zero value otherwise.

### GetReleasePipelineKeyOk

`func (o *Release) GetReleasePipelineKeyOk() (*string, bool)`

GetReleasePipelineKeyOk returns a tuple with the ReleasePipelineKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePipelineKey

`func (o *Release) SetReleasePipelineKey(v string)`

SetReleasePipelineKey sets ReleasePipelineKey field to given value.


### GetReleasePipelineDescription

`func (o *Release) GetReleasePipelineDescription() string`

GetReleasePipelineDescription returns the ReleasePipelineDescription field if non-nil, zero value otherwise.

### GetReleasePipelineDescriptionOk

`func (o *Release) GetReleasePipelineDescriptionOk() (*string, bool)`

GetReleasePipelineDescriptionOk returns a tuple with the ReleasePipelineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePipelineDescription

`func (o *Release) SetReleasePipelineDescription(v string)`

SetReleasePipelineDescription sets ReleasePipelineDescription field to given value.


### GetPhases

`func (o *Release) GetPhases() []ReleasePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *Release) GetPhasesOk() (*[]ReleasePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *Release) SetPhases(v []ReleasePhase)`

SetPhases sets Phases field to given value.


### GetVersion

`func (o *Release) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Release) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Release) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


