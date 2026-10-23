# EnvironmentReleasePhaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** | Existing rule to release on. The rule must already include the &#x60;kind&#x60; clause for &#x60;randomizationUnit&#x60; described in &#x60;startEnvironmentRelease&#x60;.  | 
**Steps** | [**[]EnvironmentReleaseStepInput**](EnvironmentReleaseStepInput.md) |  | 
**Clauses** | [**[]EnvironmentReleaseClause**](EnvironmentReleaseClause.md) | Clauses for a new rule. Must include the &#x60;kind&#x60; clause for &#x60;randomizationUnit&#x60; described in &#x60;startEnvironmentRelease&#x60;. A new rule that exactly duplicates an existing rule on the flag returns &#x60;400&#x60; with &#x60;new rule is exact duplicate of an existing rule at index &lt;n&gt;&#x60;.  | 
**Description** | Pointer to **string** | Description for a new rule. Must differ from the description of every other rule on the flag, including rules an earlier release left in place, or the request returns &#x60;400&#x60; with &#x60;duplicate rule description not allowed&#x60;.  | [optional] 

## Methods

### NewEnvironmentReleasePhaseInput

`func NewEnvironmentReleasePhaseInput(ruleId string, steps []EnvironmentReleaseStepInput, clauses []EnvironmentReleaseClause, ) *EnvironmentReleasePhaseInput`

NewEnvironmentReleasePhaseInput instantiates a new EnvironmentReleasePhaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleasePhaseInputWithDefaults

`func NewEnvironmentReleasePhaseInputWithDefaults() *EnvironmentReleasePhaseInput`

NewEnvironmentReleasePhaseInputWithDefaults instantiates a new EnvironmentReleasePhaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *EnvironmentReleasePhaseInput) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EnvironmentReleasePhaseInput) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EnvironmentReleasePhaseInput) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetSteps

`func (o *EnvironmentReleasePhaseInput) GetSteps() []EnvironmentReleaseStepInput`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *EnvironmentReleasePhaseInput) GetStepsOk() (*[]EnvironmentReleaseStepInput, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *EnvironmentReleasePhaseInput) SetSteps(v []EnvironmentReleaseStepInput)`

SetSteps sets Steps field to given value.


### GetClauses

`func (o *EnvironmentReleasePhaseInput) GetClauses() []EnvironmentReleaseClause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *EnvironmentReleasePhaseInput) GetClausesOk() (*[]EnvironmentReleaseClause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *EnvironmentReleasePhaseInput) SetClauses(v []EnvironmentReleaseClause)`

SetClauses sets Clauses field to given value.


### GetDescription

`func (o *EnvironmentReleasePhaseInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReleasePhaseInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReleasePhaseInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReleasePhaseInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


