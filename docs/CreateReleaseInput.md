# CreateReleaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseVariationId** | Pointer to **string** | The variation id to release to across all phases | [optional] 
**ReleasePipelineKey** | **string** | The key of the release pipeline to attach the flag to | 

## Methods

### NewCreateReleaseInput

`func NewCreateReleaseInput(releasePipelineKey string, ) *CreateReleaseInput`

NewCreateReleaseInput instantiates a new CreateReleaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleaseInputWithDefaults

`func NewCreateReleaseInputWithDefaults() *CreateReleaseInput`

NewCreateReleaseInputWithDefaults instantiates a new CreateReleaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseVariationId

`func (o *CreateReleaseInput) GetReleaseVariationId() string`

GetReleaseVariationId returns the ReleaseVariationId field if non-nil, zero value otherwise.

### GetReleaseVariationIdOk

`func (o *CreateReleaseInput) GetReleaseVariationIdOk() (*string, bool)`

GetReleaseVariationIdOk returns a tuple with the ReleaseVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVariationId

`func (o *CreateReleaseInput) SetReleaseVariationId(v string)`

SetReleaseVariationId sets ReleaseVariationId field to given value.

### HasReleaseVariationId

`func (o *CreateReleaseInput) HasReleaseVariationId() bool`

HasReleaseVariationId returns a boolean if a field has been set.

### GetReleasePipelineKey

`func (o *CreateReleaseInput) GetReleasePipelineKey() string`

GetReleasePipelineKey returns the ReleasePipelineKey field if non-nil, zero value otherwise.

### GetReleasePipelineKeyOk

`func (o *CreateReleaseInput) GetReleasePipelineKeyOk() (*string, bool)`

GetReleasePipelineKeyOk returns a tuple with the ReleasePipelineKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePipelineKey

`func (o *CreateReleaseInput) SetReleasePipelineKey(v string)`

SetReleasePipelineKey sets ReleasePipelineKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


