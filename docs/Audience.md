# Audience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**EnvironmentSummary**](EnvironmentSummary.md) |  | 
**Name** | **string** | The release phase name | 

## Methods

### NewAudience

`func NewAudience(environment EnvironmentSummary, name string, ) *Audience`

NewAudience instantiates a new Audience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceWithDefaults

`func NewAudienceWithDefaults() *Audience`

NewAudienceWithDefaults instantiates a new Audience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *Audience) GetEnvironment() EnvironmentSummary`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Audience) GetEnvironmentOk() (*EnvironmentSummary, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Audience) SetEnvironment(v EnvironmentSummary)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *Audience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Audience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Audience) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


