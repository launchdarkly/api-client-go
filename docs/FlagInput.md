# FlagInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** | The ID of the variation or rollout of the flag to use. Use \&quot;fallthrough\&quot; for the default targeting behavior when the flag is on. | 
**FlagConfigVersion** | **int32** | The flag version | 
**NotInExperimentVariationId** | Pointer to **string** | The ID of the variation to route traffic not part of the experiment analysis to. Defaults to variation ID of baseline treatment, if set. | [optional] 

## Methods

### NewFlagInput

`func NewFlagInput(ruleId string, flagConfigVersion int32, ) *FlagInput`

NewFlagInput instantiates a new FlagInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagInputWithDefaults

`func NewFlagInputWithDefaults() *FlagInput`

NewFlagInputWithDefaults instantiates a new FlagInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *FlagInput) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *FlagInput) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *FlagInput) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetFlagConfigVersion

`func (o *FlagInput) GetFlagConfigVersion() int32`

GetFlagConfigVersion returns the FlagConfigVersion field if non-nil, zero value otherwise.

### GetFlagConfigVersionOk

`func (o *FlagInput) GetFlagConfigVersionOk() (*int32, bool)`

GetFlagConfigVersionOk returns a tuple with the FlagConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagConfigVersion

`func (o *FlagInput) SetFlagConfigVersion(v int32)`

SetFlagConfigVersion sets FlagConfigVersion field to given value.


### GetNotInExperimentVariationId

`func (o *FlagInput) GetNotInExperimentVariationId() string`

GetNotInExperimentVariationId returns the NotInExperimentVariationId field if non-nil, zero value otherwise.

### GetNotInExperimentVariationIdOk

`func (o *FlagInput) GetNotInExperimentVariationIdOk() (*string, bool)`

GetNotInExperimentVariationIdOk returns a tuple with the NotInExperimentVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInExperimentVariationId

`func (o *FlagInput) SetNotInExperimentVariationId(v string)`

SetNotInExperimentVariationId sets NotInExperimentVariationId field to given value.

### HasNotInExperimentVariationId

`func (o *FlagInput) HasNotInExperimentVariationId() bool`

HasNotInExperimentVariationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


