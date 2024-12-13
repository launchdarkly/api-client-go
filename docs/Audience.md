# Audience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**EnvironmentSummary**](EnvironmentSummary.md) |  | [optional] 
**Name** | **string** | The release phase name | 
**Configuration** | Pointer to [**AudienceConfiguration**](AudienceConfiguration.md) |  | [optional] 
**SegmentKeys** | Pointer to **[]string** | A list of segment keys | [optional] 

## Methods

### NewAudience

`func NewAudience(name string, ) *Audience`

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

### HasEnvironment

`func (o *Audience) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

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


### GetConfiguration

`func (o *Audience) GetConfiguration() AudienceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Audience) GetConfigurationOk() (*AudienceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Audience) SetConfiguration(v AudienceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Audience) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetSegmentKeys

`func (o *Audience) GetSegmentKeys() []string`

GetSegmentKeys returns the SegmentKeys field if non-nil, zero value otherwise.

### GetSegmentKeysOk

`func (o *Audience) GetSegmentKeysOk() (*[]string, bool)`

GetSegmentKeysOk returns a tuple with the SegmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeys

`func (o *Audience) SetSegmentKeys(v []string)`

SetSegmentKeys sets SegmentKeys field to given value.

### HasSegmentKeys

`func (o *Audience) HasSegmentKeys() bool`

HasSegmentKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


