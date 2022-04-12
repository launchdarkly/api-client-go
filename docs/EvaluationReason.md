# EvaluationReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Describes the general reason that LaunchDarkly selected this variation. | 
**RuleIndex** | Pointer to **int32** | The positional index of the matching rule if the kind is &#39;RULE_MATCH&#39;. The index is 0-based. | [optional] 
**RuleID** | Pointer to **string** | The unique identifier of the matching rule if the kind is &#39;RULE_MATCH&#39;. | [optional] 
**PrerequisiteKey** | Pointer to **string** | The key of the flag that failed if the kind is &#39;PREREQUISITE_FAILED&#39;. | [optional] 
**InExperiment** | Pointer to **bool** | Indicates whether the user was evaluated as part of an experiment. | [optional] 
**ErrorKind** | Pointer to **string** | The specific error type if the kind is &#39;ERROR&#39;. | [optional] 

## Methods

### NewEvaluationReason

`func NewEvaluationReason(kind string, ) *EvaluationReason`

NewEvaluationReason instantiates a new EvaluationReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationReasonWithDefaults

`func NewEvaluationReasonWithDefaults() *EvaluationReason`

NewEvaluationReasonWithDefaults instantiates a new EvaluationReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *EvaluationReason) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EvaluationReason) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EvaluationReason) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRuleIndex

`func (o *EvaluationReason) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *EvaluationReason) GetRuleIndexOk() (*int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *EvaluationReason) SetRuleIndex(v int32)`

SetRuleIndex sets RuleIndex field to given value.

### HasRuleIndex

`func (o *EvaluationReason) HasRuleIndex() bool`

HasRuleIndex returns a boolean if a field has been set.

### GetRuleID

`func (o *EvaluationReason) GetRuleID() string`

GetRuleID returns the RuleID field if non-nil, zero value otherwise.

### GetRuleIDOk

`func (o *EvaluationReason) GetRuleIDOk() (*string, bool)`

GetRuleIDOk returns a tuple with the RuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleID

`func (o *EvaluationReason) SetRuleID(v string)`

SetRuleID sets RuleID field to given value.

### HasRuleID

`func (o *EvaluationReason) HasRuleID() bool`

HasRuleID returns a boolean if a field has been set.

### GetPrerequisiteKey

`func (o *EvaluationReason) GetPrerequisiteKey() string`

GetPrerequisiteKey returns the PrerequisiteKey field if non-nil, zero value otherwise.

### GetPrerequisiteKeyOk

`func (o *EvaluationReason) GetPrerequisiteKeyOk() (*string, bool)`

GetPrerequisiteKeyOk returns a tuple with the PrerequisiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteKey

`func (o *EvaluationReason) SetPrerequisiteKey(v string)`

SetPrerequisiteKey sets PrerequisiteKey field to given value.

### HasPrerequisiteKey

`func (o *EvaluationReason) HasPrerequisiteKey() bool`

HasPrerequisiteKey returns a boolean if a field has been set.

### GetInExperiment

`func (o *EvaluationReason) GetInExperiment() bool`

GetInExperiment returns the InExperiment field if non-nil, zero value otherwise.

### GetInExperimentOk

`func (o *EvaluationReason) GetInExperimentOk() (*bool, bool)`

GetInExperimentOk returns a tuple with the InExperiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInExperiment

`func (o *EvaluationReason) SetInExperiment(v bool)`

SetInExperiment sets InExperiment field to given value.

### HasInExperiment

`func (o *EvaluationReason) HasInExperiment() bool`

HasInExperiment returns a boolean if a field has been set.

### GetErrorKind

`func (o *EvaluationReason) GetErrorKind() string`

GetErrorKind returns the ErrorKind field if non-nil, zero value otherwise.

### GetErrorKindOk

`func (o *EvaluationReason) GetErrorKindOk() (*string, bool)`

GetErrorKindOk returns a tuple with the ErrorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorKind

`func (o *EvaluationReason) SetErrorKind(v string)`

SetErrorKind sets ErrorKind field to given value.

### HasErrorKind

`func (o *EvaluationReason) HasErrorKind() bool`

HasErrorKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


