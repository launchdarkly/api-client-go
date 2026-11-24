# EnvironmentReleasePlanProgressiveReleaseRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RandomizationUnit** | Pointer to **string** | Context kind used to bucket contexts for allocation. Required when you create or update a plan. | [optional] 
**AdvancementMode** | Pointer to **string** | Whether a release started from this plan advances to the next phase automatically once the current phase&#39;s steps complete, or waits for a user/API action | [optional] 
**CleanupBehavior** | Pointer to **string** | What happens to the flag&#39;s rules once a release started from this plan completes. &#x60;remove_rules_and_update_default_rule&#x60; removes the release&#39;s targeting rules and updates the flag&#39;s default rule to serve the target variation; &#x60;leave_rules_in_place&#x60; leaves the targeting rules untouched. | [optional] 
**Phases** | [**[]EnvironmentReleasePlanPhase**](EnvironmentReleasePlanPhase.md) | Ordered list of phases this release plan configures | 
**Metrics** | Pointer to [**map[string]EnvironmentReleasePlanMetric**](EnvironmentReleasePlanMetric.md) | Metrics/metric groups to monitor during rollout, keyed by metric/group key | [optional] 

## Methods

### NewEnvironmentReleasePlanProgressiveReleaseRep

`func NewEnvironmentReleasePlanProgressiveReleaseRep(phases []EnvironmentReleasePlanPhase, ) *EnvironmentReleasePlanProgressiveReleaseRep`

NewEnvironmentReleasePlanProgressiveReleaseRep instantiates a new EnvironmentReleasePlanProgressiveReleaseRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanProgressiveReleaseRepWithDefaults

`func NewEnvironmentReleasePlanProgressiveReleaseRepWithDefaults() *EnvironmentReleasePlanProgressiveReleaseRep`

NewEnvironmentReleasePlanProgressiveReleaseRepWithDefaults instantiates a new EnvironmentReleasePlanProgressiveReleaseRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRandomizationUnit

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetAdvancementMode

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetAdvancementMode() string`

GetAdvancementMode returns the AdvancementMode field if non-nil, zero value otherwise.

### GetAdvancementModeOk

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetAdvancementModeOk() (*string, bool)`

GetAdvancementModeOk returns a tuple with the AdvancementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancementMode

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) SetAdvancementMode(v string)`

SetAdvancementMode sets AdvancementMode field to given value.

### HasAdvancementMode

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) HasAdvancementMode() bool`

HasAdvancementMode returns a boolean if a field has been set.

### GetCleanupBehavior

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetCleanupBehavior() string`

GetCleanupBehavior returns the CleanupBehavior field if non-nil, zero value otherwise.

### GetCleanupBehaviorOk

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetCleanupBehaviorOk() (*string, bool)`

GetCleanupBehaviorOk returns a tuple with the CleanupBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupBehavior

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) SetCleanupBehavior(v string)`

SetCleanupBehavior sets CleanupBehavior field to given value.

### HasCleanupBehavior

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) HasCleanupBehavior() bool`

HasCleanupBehavior returns a boolean if a field has been set.

### GetPhases

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetPhases() []EnvironmentReleasePlanPhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetPhasesOk() (*[]EnvironmentReleasePlanPhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) SetPhases(v []EnvironmentReleasePlanPhase)`

SetPhases sets Phases field to given value.


### GetMetrics

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetMetrics() map[string]EnvironmentReleasePlanMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) GetMetricsOk() (*map[string]EnvironmentReleasePlanMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) SetMetrics(v map[string]EnvironmentReleasePlanMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EnvironmentReleasePlanProgressiveReleaseRep) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


