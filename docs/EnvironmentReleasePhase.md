# EnvironmentReleasePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PhaseIndex** | **int32** | Zero-based index of this phase. | [readonly] 
**RuleId** | **NullableString** | ID of the flag rule for this phase. New public-API releases always have a rule ID. &#x60;null&#x60; is retained only when reading a historical release that ran on the default rule.  | 
**Description** | Pointer to **string** | The description of the flag rule when this release was created. Omitted for default-rule phases.  | [optional] [readonly] 
**Status** | [**ReleaseStatus**](ReleaseStatus.md) |  | 
**StatusReason** | Pointer to [**ReleaseStatusReason**](ReleaseStatusReason.md) |  | [optional] 
**Steps** | [**[]ReleaseStep**](ReleaseStep.md) |  | 
**CurrentStepIndex** | Pointer to **int32** | Zero-based index of the step the phase is on. Omitted until the first step starts. Once the phase completes or reverts, it stays at the last step that started, unlike &#x60;currentPhaseIndex&#x60;, which becomes &#x60;null&#x60; when the release reaches a terminal status.  | [optional] [readonly] 
**StartedAt** | Pointer to **int64** | Start of the phase in Unix milliseconds. | [optional] [readonly] 
**EndedAt** | Pointer to **int64** | End of the phase in Unix milliseconds. | [optional] [readonly] 
**MetricStatuses** | [**map[string]ReleaseMetricStatus**](ReleaseMetricStatus.md) | Per-phase metric status keyed by metric key. Empty for progressive releases. Metric definitions and monitoring preferences live once in the parent release&#39;s &#x60;metrics&#x60; array.  | [readonly] 

## Methods

### NewEnvironmentReleasePhase

`func NewEnvironmentReleasePhase(id string, phaseIndex int32, ruleId NullableString, status ReleaseStatus, steps []ReleaseStep, metricStatuses map[string]ReleaseMetricStatus, ) *EnvironmentReleasePhase`

NewEnvironmentReleasePhase instantiates a new EnvironmentReleasePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePhaseWithDefaults

`func NewEnvironmentReleasePhaseWithDefaults() *EnvironmentReleasePhase`

NewEnvironmentReleasePhaseWithDefaults instantiates a new EnvironmentReleasePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentReleasePhase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentReleasePhase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentReleasePhase) SetId(v string)`

SetId sets Id field to given value.


### GetPhaseIndex

`func (o *EnvironmentReleasePhase) GetPhaseIndex() int32`

GetPhaseIndex returns the PhaseIndex field if non-nil, zero value otherwise.

### GetPhaseIndexOk

`func (o *EnvironmentReleasePhase) GetPhaseIndexOk() (*int32, bool)`

GetPhaseIndexOk returns a tuple with the PhaseIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseIndex

`func (o *EnvironmentReleasePhase) SetPhaseIndex(v int32)`

SetPhaseIndex sets PhaseIndex field to given value.


### GetRuleId

`func (o *EnvironmentReleasePhase) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EnvironmentReleasePhase) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EnvironmentReleasePhase) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### SetRuleIdNil

`func (o *EnvironmentReleasePhase) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *EnvironmentReleasePhase) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetDescription

`func (o *EnvironmentReleasePhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReleasePhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReleasePhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReleasePhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentReleasePhase) GetStatus() ReleaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentReleasePhase) GetStatusOk() (*ReleaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentReleasePhase) SetStatus(v ReleaseStatus)`

SetStatus sets Status field to given value.


### GetStatusReason

`func (o *EnvironmentReleasePhase) GetStatusReason() ReleaseStatusReason`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *EnvironmentReleasePhase) GetStatusReasonOk() (*ReleaseStatusReason, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *EnvironmentReleasePhase) SetStatusReason(v ReleaseStatusReason)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *EnvironmentReleasePhase) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetSteps

`func (o *EnvironmentReleasePhase) GetSteps() []ReleaseStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *EnvironmentReleasePhase) GetStepsOk() (*[]ReleaseStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *EnvironmentReleasePhase) SetSteps(v []ReleaseStep)`

SetSteps sets Steps field to given value.


### GetCurrentStepIndex

`func (o *EnvironmentReleasePhase) GetCurrentStepIndex() int32`

GetCurrentStepIndex returns the CurrentStepIndex field if non-nil, zero value otherwise.

### GetCurrentStepIndexOk

`func (o *EnvironmentReleasePhase) GetCurrentStepIndexOk() (*int32, bool)`

GetCurrentStepIndexOk returns a tuple with the CurrentStepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepIndex

`func (o *EnvironmentReleasePhase) SetCurrentStepIndex(v int32)`

SetCurrentStepIndex sets CurrentStepIndex field to given value.

### HasCurrentStepIndex

`func (o *EnvironmentReleasePhase) HasCurrentStepIndex() bool`

HasCurrentStepIndex returns a boolean if a field has been set.

### GetStartedAt

`func (o *EnvironmentReleasePhase) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EnvironmentReleasePhase) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EnvironmentReleasePhase) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EnvironmentReleasePhase) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *EnvironmentReleasePhase) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *EnvironmentReleasePhase) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *EnvironmentReleasePhase) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *EnvironmentReleasePhase) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetMetricStatuses

`func (o *EnvironmentReleasePhase) GetMetricStatuses() map[string]ReleaseMetricStatus`

GetMetricStatuses returns the MetricStatuses field if non-nil, zero value otherwise.

### GetMetricStatusesOk

`func (o *EnvironmentReleasePhase) GetMetricStatusesOk() (*map[string]ReleaseMetricStatus, bool)`

GetMetricStatusesOk returns a tuple with the MetricStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricStatuses

`func (o *EnvironmentReleasePhase) SetMetricStatuses(v map[string]ReleaseMetricStatus)`

SetMetricStatuses sets MetricStatuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


