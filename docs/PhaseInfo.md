# PhaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The phase ID | 
**Name** | **string** | The release phase name | 
**ReleaseCount** | **int32** | The number of active releases in this phase | 

## Methods

### NewPhaseInfo

`func NewPhaseInfo(id string, name string, releaseCount int32, ) *PhaseInfo`

NewPhaseInfo instantiates a new PhaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseInfoWithDefaults

`func NewPhaseInfoWithDefaults() *PhaseInfo`

NewPhaseInfoWithDefaults instantiates a new PhaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhaseInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhaseInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhaseInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PhaseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhaseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhaseInfo) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseCount

`func (o *PhaseInfo) GetReleaseCount() int32`

GetReleaseCount returns the ReleaseCount field if non-nil, zero value otherwise.

### GetReleaseCountOk

`func (o *PhaseInfo) GetReleaseCountOk() (*int32, bool)`

GetReleaseCountOk returns a tuple with the ReleaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCount

`func (o *PhaseInfo) SetReleaseCount(v int32)`

SetReleaseCount sets ReleaseCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


