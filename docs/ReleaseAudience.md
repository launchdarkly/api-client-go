# ReleaseAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The audience ID | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Environment** | [**EnvironmentSummary**](EnvironmentSummary.md) |  | 
**Name** | **string** | The release phase name | 
**Configuration** | Pointer to [**AudienceConfiguration**](AudienceConfiguration.md) |  | [optional] 
**SegmentKeys** | Pointer to **[]string** | A list of segment keys | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RuleIds** | Pointer to **[]string** | The rules IDs added or updated by this audience | [optional] 

## Methods

### NewReleaseAudience

`func NewReleaseAudience(id string, environment EnvironmentSummary, name string, ) *ReleaseAudience`

NewReleaseAudience instantiates a new ReleaseAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseAudienceWithDefaults

`func NewReleaseAudienceWithDefaults() *ReleaseAudience`

NewReleaseAudienceWithDefaults instantiates a new ReleaseAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseAudience) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseAudience) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseAudience) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *ReleaseAudience) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReleaseAudience) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReleaseAudience) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReleaseAudience) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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


### GetConfiguration

`func (o *ReleaseAudience) GetConfiguration() AudienceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ReleaseAudience) GetConfigurationOk() (*AudienceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ReleaseAudience) SetConfiguration(v AudienceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ReleaseAudience) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetSegmentKeys

`func (o *ReleaseAudience) GetSegmentKeys() []string`

GetSegmentKeys returns the SegmentKeys field if non-nil, zero value otherwise.

### GetSegmentKeysOk

`func (o *ReleaseAudience) GetSegmentKeysOk() (*[]string, bool)`

GetSegmentKeysOk returns a tuple with the SegmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeys

`func (o *ReleaseAudience) SetSegmentKeys(v []string)`

SetSegmentKeys sets SegmentKeys field to given value.

### HasSegmentKeys

`func (o *ReleaseAudience) HasSegmentKeys() bool`

HasSegmentKeys returns a boolean if a field has been set.

### GetStatus

`func (o *ReleaseAudience) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReleaseAudience) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReleaseAudience) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReleaseAudience) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRuleIds

`func (o *ReleaseAudience) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *ReleaseAudience) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *ReleaseAudience) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *ReleaseAudience) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


