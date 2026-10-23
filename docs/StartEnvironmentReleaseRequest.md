# StartEnvironmentReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeforeRuleId** | Pointer to **string** | Insert the newly created rule or phase rules before this rule. If omitted, append them to the end of the rule list.  | [optional] 
**OriginalVariationId** | **string** | ID of the variation the release starts from. The release gradually moves traffic from it to the target variation, and moves traffic back to it if the release reverts.  | 
**TargetVariationId** | **string** | ID of the variation being released. The release gradually moves traffic to it from the original variation, and moves traffic back to the original variation if the release reverts.  | 
**ReleasePlanKey** | Pointer to **string** | Key of an environment release plan in this environment. The plan supplies &#x60;randomizationUnit&#x60;, &#x60;phases&#x60;, &#x60;metrics&#x60;, &#x60;advancementMode&#x60;, and &#x60;cleanupBehavior&#x60;, so omit them. Omit this field to configure the release inline.  | [optional] 
**RandomizationUnit** | Pointer to **string** | Key of an existing, non-archived context kind in the project that the release randomizes on, such as &#x60;user&#x60;. Each phase&#39;s rule must include the clause &#x60;{\&quot;contextKind\&quot;: \&quot;&lt;unit&gt;\&quot;, \&quot;attribute\&quot;: \&quot;kind\&quot;, \&quot;op\&quot;: \&quot;in\&quot;, \&quot;values\&quot;: [\&quot;&lt;unit&gt;\&quot;]}&#x60;. Required unless &#x60;releasePlanKey&#x60; is set.  | [optional] 
**AdvancementMode** | Pointer to [**EnvironmentReleaseAdvancementMode**](EnvironmentReleaseAdvancementMode.md) | How the release moves between phases. Required when &#x60;phases&#x60; has more than one entry; a single-phase release ignores it and may omit it. Omit when &#x60;releasePlanKey&#x60; is set.  | [optional] 
**Phases** | Pointer to [**[]EnvironmentReleasePhaseInput**](EnvironmentReleasePhaseInput.md) | Ordered phases of the release. Required unless &#x60;releasePlanKey&#x60; is set. | [optional] 
**Metrics** | Pointer to [**[]EnvironmentReleaseMetricInput**](EnvironmentReleaseMetricInput.md) | Metrics shared by every phase. Omit or pass an empty array for a progressive release. Metric groups cannot be listed inline; a release plan&#39;s metric groups expand into their member metrics at start. Omit when &#x60;releasePlanKey&#x60; is set.  | [optional] 
**BypassHealthChecks** | Pointer to [**[]ReleaseHealthCheckKind**](ReleaseHealthCheckKind.md) | Health check kinds that do not block this request. LaunchDarkly still runs and enforces every health check kind not listed here.  | [optional] 
**CleanupBehavior** | Pointer to [**CleanupBehavior**](CleanupBehavior.md) | What happens to the release rules after the release completes. Omit for &#x60;leave_rules_in_place&#x60;. Omit when &#x60;releasePlanKey&#x60; is set.  | [optional] 
**Comment** | Pointer to **string** | Optional comment written to the flag&#39;s audit-log entry. | [optional] 

## Methods

### NewStartEnvironmentReleaseRequest

`func NewStartEnvironmentReleaseRequest(originalVariationId string, targetVariationId string, ) *StartEnvironmentReleaseRequest`

NewStartEnvironmentReleaseRequest instantiates a new StartEnvironmentReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartEnvironmentReleaseRequestWithDefaults

`func NewStartEnvironmentReleaseRequestWithDefaults() *StartEnvironmentReleaseRequest`

NewStartEnvironmentReleaseRequestWithDefaults instantiates a new StartEnvironmentReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeforeRuleId

`func (o *StartEnvironmentReleaseRequest) GetBeforeRuleId() string`

GetBeforeRuleId returns the BeforeRuleId field if non-nil, zero value otherwise.

### GetBeforeRuleIdOk

`func (o *StartEnvironmentReleaseRequest) GetBeforeRuleIdOk() (*string, bool)`

GetBeforeRuleIdOk returns a tuple with the BeforeRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeRuleId

`func (o *StartEnvironmentReleaseRequest) SetBeforeRuleId(v string)`

SetBeforeRuleId sets BeforeRuleId field to given value.

### HasBeforeRuleId

`func (o *StartEnvironmentReleaseRequest) HasBeforeRuleId() bool`

HasBeforeRuleId returns a boolean if a field has been set.

### GetOriginalVariationId

`func (o *StartEnvironmentReleaseRequest) GetOriginalVariationId() string`

GetOriginalVariationId returns the OriginalVariationId field if non-nil, zero value otherwise.

### GetOriginalVariationIdOk

`func (o *StartEnvironmentReleaseRequest) GetOriginalVariationIdOk() (*string, bool)`

GetOriginalVariationIdOk returns a tuple with the OriginalVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVariationId

`func (o *StartEnvironmentReleaseRequest) SetOriginalVariationId(v string)`

SetOriginalVariationId sets OriginalVariationId field to given value.


### GetTargetVariationId

`func (o *StartEnvironmentReleaseRequest) GetTargetVariationId() string`

GetTargetVariationId returns the TargetVariationId field if non-nil, zero value otherwise.

### GetTargetVariationIdOk

`func (o *StartEnvironmentReleaseRequest) GetTargetVariationIdOk() (*string, bool)`

GetTargetVariationIdOk returns a tuple with the TargetVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVariationId

`func (o *StartEnvironmentReleaseRequest) SetTargetVariationId(v string)`

SetTargetVariationId sets TargetVariationId field to given value.


### GetReleasePlanKey

`func (o *StartEnvironmentReleaseRequest) GetReleasePlanKey() string`

GetReleasePlanKey returns the ReleasePlanKey field if non-nil, zero value otherwise.

### GetReleasePlanKeyOk

`func (o *StartEnvironmentReleaseRequest) GetReleasePlanKeyOk() (*string, bool)`

GetReleasePlanKeyOk returns a tuple with the ReleasePlanKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasePlanKey

`func (o *StartEnvironmentReleaseRequest) SetReleasePlanKey(v string)`

SetReleasePlanKey sets ReleasePlanKey field to given value.

### HasReleasePlanKey

`func (o *StartEnvironmentReleaseRequest) HasReleasePlanKey() bool`

HasReleasePlanKey returns a boolean if a field has been set.

### GetRandomizationUnit

`func (o *StartEnvironmentReleaseRequest) GetRandomizationUnit() string`

GetRandomizationUnit returns the RandomizationUnit field if non-nil, zero value otherwise.

### GetRandomizationUnitOk

`func (o *StartEnvironmentReleaseRequest) GetRandomizationUnitOk() (*string, bool)`

GetRandomizationUnitOk returns a tuple with the RandomizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizationUnit

`func (o *StartEnvironmentReleaseRequest) SetRandomizationUnit(v string)`

SetRandomizationUnit sets RandomizationUnit field to given value.

### HasRandomizationUnit

`func (o *StartEnvironmentReleaseRequest) HasRandomizationUnit() bool`

HasRandomizationUnit returns a boolean if a field has been set.

### GetAdvancementMode

`func (o *StartEnvironmentReleaseRequest) GetAdvancementMode() EnvironmentReleaseAdvancementMode`

GetAdvancementMode returns the AdvancementMode field if non-nil, zero value otherwise.

### GetAdvancementModeOk

`func (o *StartEnvironmentReleaseRequest) GetAdvancementModeOk() (*EnvironmentReleaseAdvancementMode, bool)`

GetAdvancementModeOk returns a tuple with the AdvancementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancementMode

`func (o *StartEnvironmentReleaseRequest) SetAdvancementMode(v EnvironmentReleaseAdvancementMode)`

SetAdvancementMode sets AdvancementMode field to given value.

### HasAdvancementMode

`func (o *StartEnvironmentReleaseRequest) HasAdvancementMode() bool`

HasAdvancementMode returns a boolean if a field has been set.

### GetPhases

`func (o *StartEnvironmentReleaseRequest) GetPhases() []EnvironmentReleasePhaseInput`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *StartEnvironmentReleaseRequest) GetPhasesOk() (*[]EnvironmentReleasePhaseInput, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *StartEnvironmentReleaseRequest) SetPhases(v []EnvironmentReleasePhaseInput)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *StartEnvironmentReleaseRequest) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetMetrics

`func (o *StartEnvironmentReleaseRequest) GetMetrics() []EnvironmentReleaseMetricInput`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *StartEnvironmentReleaseRequest) GetMetricsOk() (*[]EnvironmentReleaseMetricInput, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *StartEnvironmentReleaseRequest) SetMetrics(v []EnvironmentReleaseMetricInput)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *StartEnvironmentReleaseRequest) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetBypassHealthChecks

`func (o *StartEnvironmentReleaseRequest) GetBypassHealthChecks() []ReleaseHealthCheckKind`

GetBypassHealthChecks returns the BypassHealthChecks field if non-nil, zero value otherwise.

### GetBypassHealthChecksOk

`func (o *StartEnvironmentReleaseRequest) GetBypassHealthChecksOk() (*[]ReleaseHealthCheckKind, bool)`

GetBypassHealthChecksOk returns a tuple with the BypassHealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassHealthChecks

`func (o *StartEnvironmentReleaseRequest) SetBypassHealthChecks(v []ReleaseHealthCheckKind)`

SetBypassHealthChecks sets BypassHealthChecks field to given value.

### HasBypassHealthChecks

`func (o *StartEnvironmentReleaseRequest) HasBypassHealthChecks() bool`

HasBypassHealthChecks returns a boolean if a field has been set.

### GetCleanupBehavior

`func (o *StartEnvironmentReleaseRequest) GetCleanupBehavior() CleanupBehavior`

GetCleanupBehavior returns the CleanupBehavior field if non-nil, zero value otherwise.

### GetCleanupBehaviorOk

`func (o *StartEnvironmentReleaseRequest) GetCleanupBehaviorOk() (*CleanupBehavior, bool)`

GetCleanupBehaviorOk returns a tuple with the CleanupBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupBehavior

`func (o *StartEnvironmentReleaseRequest) SetCleanupBehavior(v CleanupBehavior)`

SetCleanupBehavior sets CleanupBehavior field to given value.

### HasCleanupBehavior

`func (o *StartEnvironmentReleaseRequest) HasCleanupBehavior() bool`

HasCleanupBehavior returns a boolean if a field has been set.

### GetComment

`func (o *StartEnvironmentReleaseRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StartEnvironmentReleaseRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StartEnvironmentReleaseRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StartEnvironmentReleaseRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


