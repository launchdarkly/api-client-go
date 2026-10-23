# PipelineReleaseAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The audience ID | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Environment** | Pointer to [**EnvironmentSummary**](EnvironmentSummary.md) |  | [optional] 
**Name** | **string** | The release phase name | 
**Configuration** | Pointer to [**AudienceConfiguration**](AudienceConfiguration.md) |  | [optional] 
**SegmentKeys** | Pointer to **[]string** | A list of segment keys | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RuleIds** | Pointer to **[]string** | The rules IDs added or updated by this audience | [optional] 

## Methods

### NewPipelineReleaseAudience

`func NewPipelineReleaseAudience(id string, name string, ) *PipelineReleaseAudience`

NewPipelineReleaseAudience instantiates a new PipelineReleaseAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReleaseAudienceWithDefaults

`func NewPipelineReleaseAudienceWithDefaults() *PipelineReleaseAudience`

NewPipelineReleaseAudienceWithDefaults instantiates a new PipelineReleaseAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineReleaseAudience) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineReleaseAudience) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineReleaseAudience) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *PipelineReleaseAudience) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PipelineReleaseAudience) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PipelineReleaseAudience) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PipelineReleaseAudience) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *PipelineReleaseAudience) GetEnvironment() EnvironmentSummary`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PipelineReleaseAudience) GetEnvironmentOk() (*EnvironmentSummary, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PipelineReleaseAudience) SetEnvironment(v EnvironmentSummary)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PipelineReleaseAudience) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *PipelineReleaseAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineReleaseAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineReleaseAudience) SetName(v string)`

SetName sets Name field to given value.


### GetConfiguration

`func (o *PipelineReleaseAudience) GetConfiguration() AudienceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PipelineReleaseAudience) GetConfigurationOk() (*AudienceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PipelineReleaseAudience) SetConfiguration(v AudienceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PipelineReleaseAudience) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetSegmentKeys

`func (o *PipelineReleaseAudience) GetSegmentKeys() []string`

GetSegmentKeys returns the SegmentKeys field if non-nil, zero value otherwise.

### GetSegmentKeysOk

`func (o *PipelineReleaseAudience) GetSegmentKeysOk() (*[]string, bool)`

GetSegmentKeysOk returns a tuple with the SegmentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeys

`func (o *PipelineReleaseAudience) SetSegmentKeys(v []string)`

SetSegmentKeys sets SegmentKeys field to given value.

### HasSegmentKeys

`func (o *PipelineReleaseAudience) HasSegmentKeys() bool`

HasSegmentKeys returns a boolean if a field has been set.

### GetStatus

`func (o *PipelineReleaseAudience) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PipelineReleaseAudience) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PipelineReleaseAudience) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PipelineReleaseAudience) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRuleIds

`func (o *PipelineReleaseAudience) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *PipelineReleaseAudience) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *PipelineReleaseAudience) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *PipelineReleaseAudience) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


