# AudiencePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentKey** | **string** | A project-unique key for the environment. | 
**Name** | **string** | The audience name | 
**SegmentKeys** | Pointer to **[]string** | Segments targeted by this audience. | [optional] 
**Configuration** | Pointer to [**AudienceConfiguration**](AudienceConfiguration.md) |  | [optional] 

## Methods

### NewAudiencePost

`func NewAudiencePost(environmentKey string, name string, ) *AudiencePost`

NewAudiencePost instantiates a new AudiencePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiencePostWithDefaults

`func NewAudiencePostWithDefaults() *AudiencePost`

NewAudiencePostWithDefaults instantiates a new AudiencePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentKey

`func (o *AudiencePost) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *AudiencePost) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *AudiencePost) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetName

`func (o *AudiencePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudiencePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudiencePost) SetName(v string)`

SetName sets Name field to given value.


### GetSegmentKeys

`func (o *AudiencePost) GetSegmentKeys() []string`

GetSegmentKeys returns the SegmentKeys field if non-nil, zero value otherwise.

### GetSegmentKeysOk

`func (o *AudiencePost) GetSegmentKeysOk() (*[]string, bool)`

GetSegmentKeysOk returns a tuple with the SegmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeys

`func (o *AudiencePost) SetSegmentKeys(v []string)`

SetSegmentKeys sets SegmentKeys field to given value.

### HasSegmentKeys

`func (o *AudiencePost) HasSegmentKeys() bool`

HasSegmentKeys returns a boolean if a field has been set.

### GetConfiguration

`func (o *AudiencePost) GetConfiguration() AudienceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AudiencePost) GetConfigurationOk() (*AudienceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AudiencePost) SetConfiguration(v AudienceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AudiencePost) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


