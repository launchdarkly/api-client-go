# MetricGroupRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this metric group | 
**Key** | **string** | A unique key to reference the metric group | 
**Name** | **string** | A human-friendly name for the metric group | 
**Kind** | **string** | The type of the metric group | 
**Description** | Pointer to **string** | Description of the metric group | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Tags** | **[]string** | Tags for the metric group | 
**CreationDate** | **int64** |  | 
**LastModified** | **int64** |  | 
**Maintainer** | [**MaintainerRep**](MaintainerRep.md) |  | 
**Metrics** | [**[]MetricInGroupRep**](MetricInGroupRep.md) | An ordered list of the metrics in this metric group | 
**Version** | **int32** | The version of this metric group | 
**Experiments** | Pointer to [**[]DependentExperimentRep**](DependentExperimentRep.md) |  | [optional] 
**ExperimentCount** | Pointer to **int32** | The number of experiments using this metric group | [optional] 
**ActiveExperimentCount** | Pointer to **int32** | The number of active experiments using this metric group | [optional] 
**ActiveGuardedRolloutCount** | Pointer to **int32** | The number of active guarded rollouts using this metric group | [optional] 
**TotalConnectionsCount** | Pointer to **int32** | The total number of connections using this metric group | [optional] 
**TotalActiveConnectionsCount** | Pointer to **int32** | The total number of active connections using this metric group | [optional] 

## Methods

### NewMetricGroupRep

`func NewMetricGroupRep(id string, key string, name string, kind string, links map[string]Link, tags []string, creationDate int64, lastModified int64, maintainer MaintainerRep, metrics []MetricInGroupRep, version int32, ) *MetricGroupRep`

NewMetricGroupRep instantiates a new MetricGroupRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGroupRepWithDefaults

`func NewMetricGroupRepWithDefaults() *MetricGroupRep`

NewMetricGroupRepWithDefaults instantiates a new MetricGroupRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricGroupRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricGroupRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricGroupRep) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *MetricGroupRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricGroupRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricGroupRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MetricGroupRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricGroupRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricGroupRep) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *MetricGroupRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricGroupRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricGroupRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *MetricGroupRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricGroupRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricGroupRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricGroupRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *MetricGroupRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricGroupRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricGroupRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetAccess

`func (o *MetricGroupRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MetricGroupRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MetricGroupRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MetricGroupRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetTags

`func (o *MetricGroupRep) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricGroupRep) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricGroupRep) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreationDate

`func (o *MetricGroupRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MetricGroupRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MetricGroupRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *MetricGroupRep) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *MetricGroupRep) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *MetricGroupRep) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetMaintainer

`func (o *MetricGroupRep) GetMaintainer() MaintainerRep`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *MetricGroupRep) GetMaintainerOk() (*MaintainerRep, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *MetricGroupRep) SetMaintainer(v MaintainerRep)`

SetMaintainer sets Maintainer field to given value.


### GetMetrics

`func (o *MetricGroupRep) GetMetrics() []MetricInGroupRep`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricGroupRep) GetMetricsOk() (*[]MetricInGroupRep, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricGroupRep) SetMetrics(v []MetricInGroupRep)`

SetMetrics sets Metrics field to given value.


### GetVersion

`func (o *MetricGroupRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetricGroupRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetricGroupRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetExperiments

`func (o *MetricGroupRep) GetExperiments() []DependentExperimentRep`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *MetricGroupRep) GetExperimentsOk() (*[]DependentExperimentRep, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *MetricGroupRep) SetExperiments(v []DependentExperimentRep)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *MetricGroupRep) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetExperimentCount

`func (o *MetricGroupRep) GetExperimentCount() int32`

GetExperimentCount returns the ExperimentCount field if non-nil, zero value otherwise.

### GetExperimentCountOk

`func (o *MetricGroupRep) GetExperimentCountOk() (*int32, bool)`

GetExperimentCountOk returns a tuple with the ExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentCount

`func (o *MetricGroupRep) SetExperimentCount(v int32)`

SetExperimentCount sets ExperimentCount field to given value.

### HasExperimentCount

`func (o *MetricGroupRep) HasExperimentCount() bool`

HasExperimentCount returns a boolean if a field has been set.

### GetActiveExperimentCount

`func (o *MetricGroupRep) GetActiveExperimentCount() int32`

GetActiveExperimentCount returns the ActiveExperimentCount field if non-nil, zero value otherwise.

### GetActiveExperimentCountOk

`func (o *MetricGroupRep) GetActiveExperimentCountOk() (*int32, bool)`

GetActiveExperimentCountOk returns a tuple with the ActiveExperimentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExperimentCount

`func (o *MetricGroupRep) SetActiveExperimentCount(v int32)`

SetActiveExperimentCount sets ActiveExperimentCount field to given value.

### HasActiveExperimentCount

`func (o *MetricGroupRep) HasActiveExperimentCount() bool`

HasActiveExperimentCount returns a boolean if a field has been set.

### GetActiveGuardedRolloutCount

`func (o *MetricGroupRep) GetActiveGuardedRolloutCount() int32`

GetActiveGuardedRolloutCount returns the ActiveGuardedRolloutCount field if non-nil, zero value otherwise.

### GetActiveGuardedRolloutCountOk

`func (o *MetricGroupRep) GetActiveGuardedRolloutCountOk() (*int32, bool)`

GetActiveGuardedRolloutCountOk returns a tuple with the ActiveGuardedRolloutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGuardedRolloutCount

`func (o *MetricGroupRep) SetActiveGuardedRolloutCount(v int32)`

SetActiveGuardedRolloutCount sets ActiveGuardedRolloutCount field to given value.

### HasActiveGuardedRolloutCount

`func (o *MetricGroupRep) HasActiveGuardedRolloutCount() bool`

HasActiveGuardedRolloutCount returns a boolean if a field has been set.

### GetTotalConnectionsCount

`func (o *MetricGroupRep) GetTotalConnectionsCount() int32`

GetTotalConnectionsCount returns the TotalConnectionsCount field if non-nil, zero value otherwise.

### GetTotalConnectionsCountOk

`func (o *MetricGroupRep) GetTotalConnectionsCountOk() (*int32, bool)`

GetTotalConnectionsCountOk returns a tuple with the TotalConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConnectionsCount

`func (o *MetricGroupRep) SetTotalConnectionsCount(v int32)`

SetTotalConnectionsCount sets TotalConnectionsCount field to given value.

### HasTotalConnectionsCount

`func (o *MetricGroupRep) HasTotalConnectionsCount() bool`

HasTotalConnectionsCount returns a boolean if a field has been set.

### GetTotalActiveConnectionsCount

`func (o *MetricGroupRep) GetTotalActiveConnectionsCount() int32`

GetTotalActiveConnectionsCount returns the TotalActiveConnectionsCount field if non-nil, zero value otherwise.

### GetTotalActiveConnectionsCountOk

`func (o *MetricGroupRep) GetTotalActiveConnectionsCountOk() (*int32, bool)`

GetTotalActiveConnectionsCountOk returns a tuple with the TotalActiveConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveConnectionsCount

`func (o *MetricGroupRep) SetTotalActiveConnectionsCount(v int32)`

SetTotalActiveConnectionsCount sets TotalActiveConnectionsCount field to given value.

### HasTotalActiveConnectionsCount

`func (o *MetricGroupRep) HasTotalActiveConnectionsCount() bool`

HasTotalActiveConnectionsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


