# EnvironmentRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | [**ReleaseStatus**](ReleaseStatus.md) |  | 
**StatusReason** | Pointer to [**ReleaseStatusReason**](ReleaseStatusReason.md) |  | [optional] 
**ProjectKey** | **string** |  | 
**EnvironmentKey** | **string** |  | 
**FlagKey** | **string** |  | 
**OriginalVariationId** | **string** | ID of the variation the release started from. | 
**OriginalVariationValue** | **interface{}** | The value of the variation. This may be a string, boolean, number, JSON object, or array. | 
**TargetVariationId** | **string** | ID of the variation being released. | 
**TargetVariationValue** | **interface{}** | The value of the variation. This may be a string, boolean, number, JSON object, or array. | 
**RandomizationUnit** | **string** | Key of the context kind the release randomizes on. | 
**CleanupBehavior** | [**CleanupBehavior**](CleanupBehavior.md) |  | 
**Kind** | [**ReleaseType**](ReleaseType.md) |  | 
**Metrics** | [**[]ReleaseMetric**](ReleaseMetric.md) |  | 
**Phases** | [**[]EnvironmentReleasePhase**](EnvironmentReleasePhase.md) |  | 
**CurrentPhaseIndex** | **NullableInt32** | Zero-based index of the active phase. Null before the first phase, between phases, and after the release reaches a terminal status. | 
**AdvancementMode** | [**EnvironmentReleaseAdvancementMode**](EnvironmentReleaseAdvancementMode.md) |  | 
**CreatedAt** | **int64** | Creation time in Unix milliseconds. | 
**StartedAt** | Pointer to **int64** | Start of the first phase in Unix milliseconds. Absent until a phase has started. | [optional] [readonly] 
**EndedAt** | Pointer to **int64** | End of the last started phase in Unix milliseconds. Present only after the release reaches a terminal status of &#x60;completed&#x60; or &#x60;reverted&#x60;. | [optional] [readonly] 
**Links** | Pointer to [**map[string]EnvironmentReleaseLink**](EnvironmentReleaseLink.md) |  | [optional] 

## Methods

### NewEnvironmentRelease

`func NewEnvironmentRelease(id string, status ReleaseStatus, projectKey string, environmentKey string, flagKey string, originalVariationId string, originalVariationValue interface{}, targetVariationId string, targetVariationValue interface{}, randomizationUnit string, cleanupBehavior CleanupBehavior, kind ReleaseType, metrics []ReleaseMetric, phases []EnvironmentReleasePhase, currentPhaseIndex NullableInt32, advancementMode EnvironmentReleaseAdvancementMode, createdAt int64, ) *EnvironmentRelease`

NewEnvironmentRelease instantiates a new EnvironmentRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleaseWithDefaults

`func NewEnvironmentReleaseWithDefaults() *EnvironmentRelease`

NewEnvironmentReleaseWithDefaults instantiates a new EnvironmentRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentRelease) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *EnvironmentRelease) GetStatus() ReleaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentRelease) GetStatusOk() (*ReleaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentRelease) SetStatus(v ReleaseStatus)`

SetStatus sets Status field to given value.


### GetStatusReason

`func (o *EnvironmentRelease) GetStatusReason() ReleaseStatusReason`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *EnvironmentRelease) GetStatusReasonOk() (*ReleaseStatusReason, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *EnvironmentRelease) SetStatusReason(v ReleaseStatusReason)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *EnvironmentRelease) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetProjectKey

`func (o *EnvironmentRelease) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *EnvironmentRelease) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *EnvironmentRelease) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.


### GetEnvironmentKey

`func (o *EnvironmentRelease) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *EnvironmentRelease) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *EnvironmentRelease) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.


### GetFlagKey

`func (o *EnvironmentRelease) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *EnvironmentRelease) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *EnvironmentRelease) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetOriginalVariationId

`func (o *EnvironmentRelease) GetOriginalVariationId() string`

GetOriginalVariationId returns the OriginalVariationId field if non-nil, zero value otherwise.

### GetOriginalVariationIdOk

`func (o *EnvironmentRelease) GetOriginalVariationIdOk() (*string, bool)`

GetOriginalVariationIdOk returns a tuple with the OriginalVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVariationId

`func (o *EnvironmentRelease) SetOriginalVariationId(v string)`

SetOriginalVariationId sets OriginalVariationId field to given value.


### GetOriginalVariationValue

`func (o *EnvironmentRelease) GetOriginalVariationValue() interface{}`

GetOriginalVariationValue returns the OriginalVariationValue field if non-nil, zero value otherwise.

### GetOriginalVariationValueOk

`func (o *EnvironmentRelease) GetOriginalVariationValueOk() (*interface{}, bool)`

GetOriginalVariationValueOk returns a tuple with the OriginalVariationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVariationValue

`func (o *EnvironmentRelease) SetOriginalVariationValue(v interface{})`

SetOriginalVariationValue sets OriginalVariationValue field to given value.


### SetOriginalVariationValueNil

`func (o *EnvironmentRelease) SetOriginalVariationValueNil(b bool)`

 SetOriginalVariationValueNil sets the value for OriginalVariationValue to be an explicit nil

### UnsetOriginalVariationValue
`func (o *EnvironmentRelease) UnsetOriginalVariationValue()`

UnsetOriginalVariationValue ensures that no value is present for OriginalVariationValue, not even an explicit nil
### GetTargetVariationId

`func (o *EnvironmentRelease) GetTargetVariationId() string`

GetTargetVariationId returns the TargetVariationId field if non-nil, zero value otherwise.

### GetTargetVariationIdOk

`func (o *EnvironmentRelease) GetTargetVariationIdOk() (*string, bool)`

GetTargetVariationIdOk returns a tuple with the TargetVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVariationId

`func (o *EnvironmentRelease) SetTargetVariationId(v string)`

SetTargetVariationId sets TargetVariationId field to given value.


### GetTargetVariationValue

`func (o *EnvironmentRelease) GetTargetVariationValue() interface{}`

GetTargetVariationValue returns the TargetVariationValue field if non-nil, zero value otherwise.

### GetTargetVariationValueOk

`func (o *EnvironmentRelease) GetTargetVariationValueOk() (*interface{}, bool)`

GetTargetVariationValueOk returns a tuple with the TargetVariationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVariationValue

`func (o *EnvironmentRelease) SetTargetVariationValue(v interface{})`

SetTargetVariationValue sets TargetVariationValue field to given value.


### SetTargetVariationValueNil

`func (o *EnvironmentRelease) SetTargetVariationValueNil(b bool)`

 SetTargetVariationValueNil sets the value for TargetVariationValue to be an explicit nil

### UnsetTargetVariationValue
`func (o *EnvironmentRelease) UnsetTargetVariationValue()`

UnsetTargetVariationValue ensures that no value is present for TargetVariationValue, not even an explicit nil
### GetRandomizationUnit

`func (o *EnvironmentRelease) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *EnvironmentRelease) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *EnvironmentRelease) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.


### GetCleanupBehavior

`func (o *EnvironmentRelease) GetCleanupBehavior() CleanupBehavior`

GetCleanupBehavior returns the CleanupBehavior field if non-nil, zero value otherwise.

### GetCleanupBehaviorOk

`func (o *EnvironmentRelease) GetCleanupBehaviorOk() (*CleanupBehavior, bool)`

GetCleanupBehaviorOk returns a tuple with the CleanupBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupBehavior

`func (o *EnvironmentRelease) SetCleanupBehavior(v CleanupBehavior)`

SetCleanupBehavior sets CleanupBehavior field to given value.


### GetKind

`func (o *EnvironmentRelease) GetKind() ReleaseType`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnvironmentRelease) GetKindOk() (*ReleaseType, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnvironmentRelease) SetKind(v ReleaseType)`

SetKind sets Kind field to given value.


### GetMetrics

`func (o *EnvironmentRelease) GetMetrics() []ReleaseMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EnvironmentRelease) GetMetricsOk() (*[]ReleaseMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EnvironmentRelease) SetMetrics(v []ReleaseMetric)`

SetMetrics sets Metrics field to given value.


### GetPhases

`func (o *EnvironmentRelease) GetPhases() []EnvironmentReleasePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *EnvironmentRelease) GetPhasesOk() (*[]EnvironmentReleasePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *EnvironmentRelease) SetPhases(v []EnvironmentReleasePhase)`

SetPhases sets Phases field to given value.


### GetCurrentPhaseIndex

`func (o *EnvironmentRelease) GetCurrentPhaseIndex() int32`

GetCurrentPhaseIndex returns the CurrentPhaseIndex field if non-nil, zero value otherwise.

### GetCurrentPhaseIndexOk

`func (o *EnvironmentRelease) GetCurrentPhaseIndexOk() (*int32, bool)`

GetCurrentPhaseIndexOk returns a tuple with the CurrentPhaseIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhaseIndex

`func (o *EnvironmentRelease) SetCurrentPhaseIndex(v int32)`

SetCurrentPhaseIndex sets CurrentPhaseIndex field to given value.


### SetCurrentPhaseIndexNil

`func (o *EnvironmentRelease) SetCurrentPhaseIndexNil(b bool)`

 SetCurrentPhaseIndexNil sets the value for CurrentPhaseIndex to be an explicit nil

### UnsetCurrentPhaseIndex
`func (o *EnvironmentRelease) UnsetCurrentPhaseIndex()`

UnsetCurrentPhaseIndex ensures that no value is present for CurrentPhaseIndex, not even an explicit nil
### GetAdvancementMode

`func (o *EnvironmentRelease) GetAdvancementMode() EnvironmentReleaseAdvancementMode`

GetAdvancementMode returns the AdvancementMode field if non-nil, zero value otherwise.

### GetAdvancementModeOk

`func (o *EnvironmentRelease) GetAdvancementModeOk() (*EnvironmentReleaseAdvancementMode, bool)`

GetAdvancementModeOk returns a tuple with the AdvancementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancementMode

`func (o *EnvironmentRelease) SetAdvancementMode(v EnvironmentReleaseAdvancementMode)`

SetAdvancementMode sets AdvancementMode field to given value.


### GetCreatedAt

`func (o *EnvironmentRelease) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentRelease) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentRelease) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *EnvironmentRelease) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EnvironmentRelease) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EnvironmentRelease) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EnvironmentRelease) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *EnvironmentRelease) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *EnvironmentRelease) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *EnvironmentRelease) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *EnvironmentRelease) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetLinks

`func (o *EnvironmentRelease) GetLinks() map[string]EnvironmentReleaseLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentRelease) GetLinksOk() (*map[string]EnvironmentReleaseLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentRelease) SetLinks(v map[string]EnvironmentReleaseLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentRelease) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


