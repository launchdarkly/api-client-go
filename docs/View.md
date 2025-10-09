# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**ViewsAccessRep**](ViewsAccessRep.md) |  | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 
**Id** | **string** | Unique ID of this view | 
**AccountId** | **string** | ID of the account that owns this view | 
**ProjectId** | **string** | ID of the project this view belongs to | 
**ProjectKey** | **string** | Key of the project this view belongs to | 
**Key** | **string** | Unique key for the view within the account/project | 
**Name** | **string** | Human-readable name for the view | 
**Description** | **string** | Optional detailed description of the view | 
**GenerateSdkKeys** | **bool** | Whether to generate SDK keys for this view. Defaults to false. | 
**Version** | **int32** | Version number for tracking changes | 
**Tags** | **[]string** | Tags associated with this view | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**Archived** | **bool** | Whether this view is archived | [default to false]
**ArchivedAt** | Pointer to **int64** |  | [optional] 
**DeletedAt** | Pointer to **int64** |  | [optional] 
**Deleted** | **bool** | Whether this view is deleted | [default to false]
**Maintainer** | Pointer to [**Maintainer**](Maintainer.md) |  | [optional] 
**FlagsSummary** | Pointer to [**FlagsSummary**](FlagsSummary.md) |  | [optional] 
**SegmentsSummary** | Pointer to [**SegmentsSummary**](SegmentsSummary.md) |  | [optional] 
**MetricsSummary** | Pointer to [**MetricsSummary**](MetricsSummary.md) |  | [optional] 
**AiConfigsSummary** | Pointer to [**AIConfigsSummary**](AIConfigsSummary.md) |  | [optional] 
**ResourceSummary** | Pointer to [**ResourceSummary**](ResourceSummary.md) |  | [optional] 
**FlagsExpanded** | Pointer to [**ExpandedLinkedFlags**](ExpandedLinkedFlags.md) |  | [optional] 
**SegmentsExpanded** | Pointer to [**ExpandedLinkedSegments**](ExpandedLinkedSegments.md) |  | [optional] 
**MetricsExpanded** | Pointer to [**ExpandedLinkedMetrics**](ExpandedLinkedMetrics.md) |  | [optional] 
**AiConfigsExpanded** | Pointer to [**ExpandedLinkedAIConfigs**](ExpandedLinkedAIConfigs.md) |  | [optional] 
**ResourcesExpanded** | Pointer to [**ExpandedLinkedResources**](ExpandedLinkedResources.md) |  | [optional] 

## Methods

### NewView

`func NewView(id string, accountId string, projectId string, projectKey string, key string, name string, description string, generateSdkKeys bool, version int32, tags []string, createdAt int64, updatedAt int64, archived bool, deleted bool, ) *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *View) GetAccess() ViewsAccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *View) GetAccessOk() (*ViewsAccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *View) SetAccess(v ViewsAccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *View) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLinks

`func (o *View) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *View) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *View) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *View) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *View) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *View) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *View) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *View) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *View) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *View) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProjectId

`func (o *View) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *View) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *View) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProjectKey

`func (o *View) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *View) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *View) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetKey

`func (o *View) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *View) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *View) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *View) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *View) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *View) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGenerateSdkKeys

`func (o *View) GetGenerateSdkKeys() bool`

GetGenerateSdkKeys returns the GenerateSdkKeys field if non-nil, zero value otherwise.

### GetGenerateSdkKeysOk

`func (o *View) GetGenerateSdkKeysOk() (*bool, bool)`

GetGenerateSdkKeysOk returns a tuple with the GenerateSdkKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSdkKeys

`func (o *View) SetGenerateSdkKeys(v bool)`

SetGenerateSdkKeys sets GenerateSdkKeys field to given value.


### GetVersion

`func (o *View) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *View) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *View) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetTags

`func (o *View) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *View) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *View) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreatedAt

`func (o *View) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *View) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *View) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *View) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *View) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *View) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchived

`func (o *View) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *View) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *View) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArchivedAt

`func (o *View) GetArchivedAt() int64`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *View) GetArchivedAtOk() (*int64, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *View) SetArchivedAt(v int64)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *View) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *View) GetDeletedAt() int64`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *View) GetDeletedAtOk() (*int64, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *View) SetDeletedAt(v int64)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *View) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *View) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *View) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *View) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetMaintainer

`func (o *View) GetMaintainer() Maintainer`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *View) GetMaintainerOk() (*Maintainer, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *View) SetMaintainer(v Maintainer)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *View) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetFlagsSummary

`func (o *View) GetFlagsSummary() FlagsSummary`

GetFlagsSummary returns the FlagsSummary field if non-nil, zero value otherwise.

### GetFlagsSummaryOk

`func (o *View) GetFlagsSummaryOk() (*FlagsSummary, bool)`

GetFlagsSummaryOk returns a tuple with the FlagsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsSummary

`func (o *View) SetFlagsSummary(v FlagsSummary)`

SetFlagsSummary sets FlagsSummary field to given value.

### HasFlagsSummary

`func (o *View) HasFlagsSummary() bool`

HasFlagsSummary returns a boolean if a field has been set.

### GetSegmentsSummary

`func (o *View) GetSegmentsSummary() SegmentsSummary`

GetSegmentsSummary returns the SegmentsSummary field if non-nil, zero value otherwise.

### GetSegmentsSummaryOk

`func (o *View) GetSegmentsSummaryOk() (*SegmentsSummary, bool)`

GetSegmentsSummaryOk returns a tuple with the SegmentsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsSummary

`func (o *View) SetSegmentsSummary(v SegmentsSummary)`

SetSegmentsSummary sets SegmentsSummary field to given value.

### HasSegmentsSummary

`func (o *View) HasSegmentsSummary() bool`

HasSegmentsSummary returns a boolean if a field has been set.

### GetMetricsSummary

`func (o *View) GetMetricsSummary() MetricsSummary`

GetMetricsSummary returns the MetricsSummary field if non-nil, zero value otherwise.

### GetMetricsSummaryOk

`func (o *View) GetMetricsSummaryOk() (*MetricsSummary, bool)`

GetMetricsSummaryOk returns a tuple with the MetricsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSummary

`func (o *View) SetMetricsSummary(v MetricsSummary)`

SetMetricsSummary sets MetricsSummary field to given value.

### HasMetricsSummary

`func (o *View) HasMetricsSummary() bool`

HasMetricsSummary returns a boolean if a field has been set.

### GetAiConfigsSummary

`func (o *View) GetAiConfigsSummary() AIConfigsSummary`

GetAiConfigsSummary returns the AiConfigsSummary field if non-nil, zero value otherwise.

### GetAiConfigsSummaryOk

`func (o *View) GetAiConfigsSummaryOk() (*AIConfigsSummary, bool)`

GetAiConfigsSummaryOk returns a tuple with the AiConfigsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigsSummary

`func (o *View) SetAiConfigsSummary(v AIConfigsSummary)`

SetAiConfigsSummary sets AiConfigsSummary field to given value.

### HasAiConfigsSummary

`func (o *View) HasAiConfigsSummary() bool`

HasAiConfigsSummary returns a boolean if a field has been set.

### GetResourceSummary

`func (o *View) GetResourceSummary() ResourceSummary`

GetResourceSummary returns the ResourceSummary field if non-nil, zero value otherwise.

### GetResourceSummaryOk

`func (o *View) GetResourceSummaryOk() (*ResourceSummary, bool)`

GetResourceSummaryOk returns a tuple with the ResourceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSummary

`func (o *View) SetResourceSummary(v ResourceSummary)`

SetResourceSummary sets ResourceSummary field to given value.

### HasResourceSummary

`func (o *View) HasResourceSummary() bool`

HasResourceSummary returns a boolean if a field has been set.

### GetFlagsExpanded

`func (o *View) GetFlagsExpanded() ExpandedLinkedFlags`

GetFlagsExpanded returns the FlagsExpanded field if non-nil, zero value otherwise.

### GetFlagsExpandedOk

`func (o *View) GetFlagsExpandedOk() (*ExpandedLinkedFlags, bool)`

GetFlagsExpandedOk returns a tuple with the FlagsExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsExpanded

`func (o *View) SetFlagsExpanded(v ExpandedLinkedFlags)`

SetFlagsExpanded sets FlagsExpanded field to given value.

### HasFlagsExpanded

`func (o *View) HasFlagsExpanded() bool`

HasFlagsExpanded returns a boolean if a field has been set.

### GetSegmentsExpanded

`func (o *View) GetSegmentsExpanded() ExpandedLinkedSegments`

GetSegmentsExpanded returns the SegmentsExpanded field if non-nil, zero value otherwise.

### GetSegmentsExpandedOk

`func (o *View) GetSegmentsExpandedOk() (*ExpandedLinkedSegments, bool)`

GetSegmentsExpandedOk returns a tuple with the SegmentsExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsExpanded

`func (o *View) SetSegmentsExpanded(v ExpandedLinkedSegments)`

SetSegmentsExpanded sets SegmentsExpanded field to given value.

### HasSegmentsExpanded

`func (o *View) HasSegmentsExpanded() bool`

HasSegmentsExpanded returns a boolean if a field has been set.

### GetMetricsExpanded

`func (o *View) GetMetricsExpanded() ExpandedLinkedMetrics`

GetMetricsExpanded returns the MetricsExpanded field if non-nil, zero value otherwise.

### GetMetricsExpandedOk

`func (o *View) GetMetricsExpandedOk() (*ExpandedLinkedMetrics, bool)`

GetMetricsExpandedOk returns a tuple with the MetricsExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsExpanded

`func (o *View) SetMetricsExpanded(v ExpandedLinkedMetrics)`

SetMetricsExpanded sets MetricsExpanded field to given value.

### HasMetricsExpanded

`func (o *View) HasMetricsExpanded() bool`

HasMetricsExpanded returns a boolean if a field has been set.

### GetAiConfigsExpanded

`func (o *View) GetAiConfigsExpanded() ExpandedLinkedAIConfigs`

GetAiConfigsExpanded returns the AiConfigsExpanded field if non-nil, zero value otherwise.

### GetAiConfigsExpandedOk

`func (o *View) GetAiConfigsExpandedOk() (*ExpandedLinkedAIConfigs, bool)`

GetAiConfigsExpandedOk returns a tuple with the AiConfigsExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiConfigsExpanded

`func (o *View) SetAiConfigsExpanded(v ExpandedLinkedAIConfigs)`

SetAiConfigsExpanded sets AiConfigsExpanded field to given value.

### HasAiConfigsExpanded

`func (o *View) HasAiConfigsExpanded() bool`

HasAiConfigsExpanded returns a boolean if a field has been set.

### GetResourcesExpanded

`func (o *View) GetResourcesExpanded() ExpandedLinkedResources`

GetResourcesExpanded returns the ResourcesExpanded field if non-nil, zero value otherwise.

### GetResourcesExpandedOk

`func (o *View) GetResourcesExpandedOk() (*ExpandedLinkedResources, bool)`

GetResourcesExpandedOk returns a tuple with the ResourcesExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesExpanded

`func (o *View) SetResourcesExpanded(v ExpandedLinkedResources)`

SetResourcesExpanded sets ResourcesExpanded field to given value.

### HasResourcesExpanded

`func (o *View) HasResourcesExpanded() bool`

HasResourcesExpanded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


