# EnvironmentReleaseExistingRulePhaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** | Existing rule to release on. The rule must already include the &#x60;kind&#x60; clause for &#x60;randomizationUnit&#x60; described in &#x60;startEnvironmentRelease&#x60;.  | 
**Steps** | [**[]EnvironmentReleaseStepInput**](EnvironmentReleaseStepInput.md) |  | 

## Methods

### NewEnvironmentReleaseExistingRulePhaseInput

`func NewEnvironmentReleaseExistingRulePhaseInput(ruleId string, steps []EnvironmentReleaseStepInput, ) *EnvironmentReleaseExistingRulePhaseInput`

NewEnvironmentReleaseExistingRulePhaseInput instantiates a new EnvironmentReleaseExistingRulePhaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReleaseExistingRulePhaseInputWithDefaults

`func NewEnvironmentReleaseExistingRulePhaseInputWithDefaults() *EnvironmentReleaseExistingRulePhaseInput`

NewEnvironmentReleaseExistingRulePhaseInputWithDefaults instantiates a new EnvironmentReleaseExistingRulePhaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *EnvironmentReleaseExistingRulePhaseInput) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EnvironmentReleaseExistingRulePhaseInput) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EnvironmentReleaseExistingRulePhaseInput) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetSteps

`func (o *EnvironmentReleaseExistingRulePhaseInput) GetSteps() []EnvironmentReleaseStepInput`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *EnvironmentReleaseExistingRulePhaseInput) GetStepsOk() (*[]EnvironmentReleaseStepInput, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *EnvironmentReleaseExistingRulePhaseInput) SetSteps(v []EnvironmentReleaseStepInput)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


