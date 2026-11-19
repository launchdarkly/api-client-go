# EnvironmentReleasePlanProgressiveRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RandomizationUnit** | **string** | Context kind used to bucket contexts for allocation | 
**AdvancementMode** | Pointer to **string** | Whether a release started from this plan advances to the next phase automatically once the current phase&#39;s steps complete, or waits for a user/API action | [optional] [default to "manual"]
**CleanupBehavior** | Pointer to **string** | What happens to the flag&#39;s rules once a release started from this plan completes. Defaults to &#x60;remove_rules_and_update_default_rule&#x60;, which removes the release&#39;s targeting rules and updates the flag&#39;s default rule to serve the target variation; &#x60;leave_rules_in_place&#x60; leaves the targeting rules untouched. | [optional] [default to "remove_rules_and_update_default_rule"]
**Phases** | [**[]EnvironmentReleasePlanPhase**](EnvironmentReleasePlanPhase.md) | Ordered list of phases this release plan configures | 
**Metrics** | Pointer to [**map[string]EnvironmentReleasePlanMetric**](EnvironmentReleasePlanMetric.md) | Metrics/metric groups to monitor during rollout, keyed by metric/group key | [optional] 

## Methods

### NewEnvironmentReleasePlanProgressiveRelease

`func NewEnvironmentReleasePlanProgressiveRelease(randomizationUnit string, phases []EnvironmentReleasePlanPhase, ) *EnvironmentReleasePlanProgressiveRelease`

NewEnvironmentReleasePlanProgressiveRelease instantiates a new EnvironmentReleasePlanProgressiveRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePlanProgressiveReleaseWithDefaults

`func NewEnvironmentReleasePlanProgressiveReleaseWithDefaults() *EnvironmentReleasePlanProgressiveRelease`

NewEnvironmentReleasePlanProgressiveReleaseWithDefaults instantiates a new EnvironmentReleasePlanProgressiveRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRandomizationUnit

`func (o *EnvironmentReleasePlanProgressiveRelease) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *EnvironmentReleasePlanProgressiveRelease) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *EnvironmentReleasePlanProgressiveRelease) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.


### GetAdvancementMode

`func (o *EnvironmentReleasePlanProgressiveRelease) GetAdvancementMode() string`

GetAdvancementMode returns the AdvancementMode field if non-nil, zero value otherwise.

### GetAdvancementModeOk

`func (o *EnvironmentReleasePlanProgressiveRelease) GetAdvancementModeOk() (*string, bool)`

GetAdvancementModeOk returns a tuple with the AdvancementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancementMode

`func (o *EnvironmentReleasePlanProgressiveRelease) SetAdvancementMode(v string)`

SetAdvancementMode sets AdvancementMode field to given value.

### HasAdvancementMode

`func (o *EnvironmentReleasePlanProgressiveRelease) HasAdvancementMode() bool`

HasAdvancementMode returns a boolean if a field has been set.

### GetCleanupBehavior

`func (o *EnvironmentReleasePlanProgressiveRelease) GetCleanupBehavior() string`

GetCleanupBehavior returns the CleanupBehavior field if non-nil, zero value otherwise.

### GetCleanupBehaviorOk

`func (o *EnvironmentReleasePlanProgressiveRelease) GetCleanupBehaviorOk() (*string, bool)`

GetCleanupBehaviorOk returns a tuple with the CleanupBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupBehavior

`func (o *EnvironmentReleasePlanProgressiveRelease) SetCleanupBehavior(v string)`

SetCleanupBehavior sets CleanupBehavior field to given value.

### HasCleanupBehavior

`func (o *EnvironmentReleasePlanProgressiveRelease) HasCleanupBehavior() bool`

HasCleanupBehavior returns a boolean if a field has been set.

### GetPhases

`func (o *EnvironmentReleasePlanProgressiveRelease) GetPhases() []EnvironmentReleasePlanPhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *EnvironmentReleasePlanProgressiveRelease) GetPhasesOk() (*[]EnvironmentReleasePlanPhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *EnvironmentReleasePlanProgressiveRelease) SetPhases(v []EnvironmentReleasePlanPhase)`

SetPhases sets Phases field to given value.


### GetMetrics

`func (o *EnvironmentReleasePlanProgressiveRelease) GetMetrics() map[string]EnvironmentReleasePlanMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EnvironmentReleasePlanProgressiveRelease) GetMetricsOk() (*map[string]EnvironmentReleasePlanMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EnvironmentReleasePlanProgressiveRelease) SetMetrics(v map[string]EnvironmentReleasePlanMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EnvironmentReleasePlanProgressiveRelease) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


