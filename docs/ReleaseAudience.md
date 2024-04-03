# ReleaseAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**EnvironmentSummary**](EnvironmentSummary.md) |  | 
**Name** | **string** | The release phase name | 

## Methods

### NewReleaseAudience

`func NewReleaseAudience(environment EnvironmentSummary, name string, ) *ReleaseAudience`

NewReleaseAudience instantiates a new ReleaseAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseAudienceWithDefaults

`func NewReleaseAudienceWithDefaults() *ReleaseAudience`

NewReleaseAudienceWithDefaults instantiates a new ReleaseAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ReleaseAudience) GetEnvironment() EnvironmentSummary`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReleaseAudience) GetEnvironmentOk() (*EnvironmentSummary, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReleaseAudience) SetEnvironment(v EnvironmentSummary)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *ReleaseAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseAudience) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


