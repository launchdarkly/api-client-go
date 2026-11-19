# AIConfigTargetingEnvironmentRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The rule ID. Pass it as ruleId in targeting instructions that change this rule. | 
**Disabled** | Pointer to **bool** | Whether the rule is disabled. A disabled rule doesn&#39;t match any contexts. | [optional] 
**Description** | Pointer to **string** | The rule description | [optional] 
**Clauses** | [**[]AIConfigTargetingEnvironmentRuleClause**](AIConfigTargetingEnvironmentRuleClause.md) | The clauses a context must match | 
**Variation** | Pointer to **int32** | Index of the variation served by this rule | [optional] 
**Rollout** | Pointer to [**AIConfigTargetingEnvironmentFallthroughRollout**](AIConfigTargetingEnvironmentFallthroughRollout.md) |  | [optional] 
**TrackEvents** | **bool** | Whether to send detailed events for evaluations that match this rule | 

## Methods

### NewAIConfigTargetingEnvironmentRule

`func NewAIConfigTargetingEnvironmentRule(id string, clauses []AIConfigTargetingEnvironmentRuleClause, trackEvents bool, ) *AIConfigTargetingEnvironmentRule`

NewAIConfigTargetingEnvironmentRule instantiates a new AIConfigTargetingEnvironmentRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIConfigTargetingEnvironmentRuleWithDefaults

`func NewAIConfigTargetingEnvironmentRuleWithDefaults() *AIConfigTargetingEnvironmentRule`

NewAIConfigTargetingEnvironmentRuleWithDefaults instantiates a new AIConfigTargetingEnvironmentRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AIConfigTargetingEnvironmentRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIConfigTargetingEnvironmentRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIConfigTargetingEnvironmentRule) SetId(v string)`

SetId sets Id field to given value.


### GetDisabled

`func (o *AIConfigTargetingEnvironmentRule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AIConfigTargetingEnvironmentRule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AIConfigTargetingEnvironmentRule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AIConfigTargetingEnvironmentRule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *AIConfigTargetingEnvironmentRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIConfigTargetingEnvironmentRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIConfigTargetingEnvironmentRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AIConfigTargetingEnvironmentRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClauses

`func (o *AIConfigTargetingEnvironmentRule) GetClauses() []AIConfigTargetingEnvironmentRuleClause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *AIConfigTargetingEnvironmentRule) GetClausesOk() (*[]AIConfigTargetingEnvironmentRuleClause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *AIConfigTargetingEnvironmentRule) SetClauses(v []AIConfigTargetingEnvironmentRuleClause)`

SetClauses sets Clauses field to given value.


### GetVariation

`func (o *AIConfigTargetingEnvironmentRule) GetVariation() int32`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *AIConfigTargetingEnvironmentRule) GetVariationOk() (*int32, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *AIConfigTargetingEnvironmentRule) SetVariation(v int32)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *AIConfigTargetingEnvironmentRule) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetRollout

`func (o *AIConfigTargetingEnvironmentRule) GetRollout() AIConfigTargetingEnvironmentFallthroughRollout`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *AIConfigTargetingEnvironmentRule) GetRolloutOk() (*AIConfigTargetingEnvironmentFallthroughRollout, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *AIConfigTargetingEnvironmentRule) SetRollout(v AIConfigTargetingEnvironmentFallthroughRollout)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *AIConfigTargetingEnvironmentRule) HasRollout() bool`

HasRollout returns a boolean if a field has been set.

### GetTrackEvents

`func (o *AIConfigTargetingEnvironmentRule) GetTrackEvents() bool`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *AIConfigTargetingEnvironmentRule) GetTrackEventsOk() (*bool, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *AIConfigTargetingEnvironmentRule) SetTrackEvents(v bool)`

SetTrackEvents sets TrackEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


